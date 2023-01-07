package main

import (
	"context"
	"encoding/xml"
	"errors"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var (
	githubClient *githubv4.Client

	appcastSingleFlightGroup     singleflight.Group
	appcastLastUpdatedAt         time.Time
	appcastMinimalUpdateInterval = time.Minute

	appcast atomic.Pointer[[]byte]
)

var errAppcastXMLNotReady = errors.New("appcast.xml not ready")

func init() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	githubClient = githubv4.NewClient(httpClient)
}

type Release struct {
	XMLName     xml.Name  `xml:"item"`
	Title       string    `xml:"title"`
	Version     string    `xml:"sparkle:version"`
	Channel     string    `xml:"sparkle:channel,omitempty"`
	PubDate     time.Time `xml:"pubDate"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Enclosure   struct {
		URL  string `xml:"url,attr"`
		Type string `xml:"type,attr"`
	} `xml:"enclosure"`
}

type Channel struct {
	XMLName  xml.Name `xml:"channel"`
	Title    string   `xml:"title"`
	Link     string   `xml:"link"`
	Releases []Release
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Sparkle string   `xml:"xmlns:sparkle,attr"`
	DC      string   `xml:"xmlns:dc,attr"`
	Version string   `xml:"version,attr"`
	Channel Channel
}

func getAppCast() ([]byte, error) {
	ch := appcastSingleFlightGroup.DoChan("", func() (any, error) {
		if time.Since(appcastLastUpdatedAt) < appcastMinimalUpdateInterval {
			return nil, nil
		}

		updateAppcast()

		appcastLastUpdatedAt = time.Now()

		return nil, nil
	})

	a := appcast.Load()
	if a == nil {
		<-ch
		a = appcast.Load()
		if a == nil {
			return nil, errAppcastXMLNotReady
		}
	}

	// In most cases, we won't wait until `ch` to be fulfilled.
	// It's okay not to consume the data in `ch` since the `ch` is a buffered channel.

	return *a, nil
}

func updateAppcast() {
	log.Println("Update appcast.xml")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query struct {
		Repository struct {
			Releases struct {
				Nodes []struct {
					Name            string
					TagName         string
					PublishedAt     time.Time
					DescriptionHTML string `graphql:"descriptionHTML"`
					ReleaseAssets   struct {
						Nodes []struct {
							DownloadURL string
						}
					} `graphql:"releaseAssets(first: 1)"`
				}
			} `graphql:"releases(first: 10)"`
		} `graphql:"repository(owner: \"linearmouse\", name: \"linearmouse\")"`
	}

	err := githubClient.Query(ctx, &query, nil)
	if err != nil {
		log.Printf("githubClient.Query: %v", err)
		return
	}

	var releases []Release

	for _, node := range query.Repository.Releases.Nodes {
		var channel string
		if strings.Contains(node.TagName, "-beta.") {
			channel = "beta"
		}

		release := Release{
			Title:       node.Name,
			Version:     strings.TrimPrefix(node.TagName, "v"),
			Channel:     channel,
			PubDate:     node.PublishedAt,
			Link:        "https://github.com/linearmouse/linearmouse/releases/tag/" + node.TagName,
			Description: node.DescriptionHTML,
		}

		release.Enclosure.Type = "application/octet-stream"
		release.Enclosure.URL = strings.Replace(node.ReleaseAssets.Nodes[0].DownloadURL,
			"https://github.com/linearmouse/linearmouse/releases/download/",
			"https://dl.linearmouse.org/", 1)

		releases = append(releases, release)
	}

	channel := Channel{
		Title:    "LinearMouse",
		Link:     "https://linearmouse.app",
		Releases: releases,
	}

	rss := RSS{
		Sparkle: "http://www.andymatuschak.org/xml-namespaces/sparkle",
		DC:      "http://purl.org/dc/elements/1.1/",
		Version: "2.0",
		Channel: channel,
	}

	a, err := xml.Marshal(rss)
	if err != nil {
		log.Printf("xml.Marshal: %v", err)
		return
	}

	appcast.Store(&a)
}
