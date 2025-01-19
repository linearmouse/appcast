package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "appcast_requests_total",
	}, []string{"app_version"})

	uaAppVersionRe = regexp.MustCompile(`(?:^|\s+)LinearMouse/(\d+\.\d+\.\d+(?:-beta\.\d+)?)`)
)

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/appcast.xml" {
		http.NotFound(w, r)
		return
	}

	if handleSparkle2(w, r) {
		return
	}

	match := uaAppVersionRe.FindStringSubmatch(r.UserAgent())
	if match != nil {
		appVersion := match[1]
		if appVersion != "" {
			requestsTotal.With(prometheus.Labels{"app_version": appVersion}).Inc()
		}
	}

	appcast, err := getAppCast()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/xml")
	w.Write(appcast)
}

func main() {
	go startMetricsServer()

	http.HandleFunc("/", handle)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func startMetricsServer() {
	metricsMux := http.NewServeMux()

	metricsMux.Handle("/metrics", promhttp.Handler())

	log.Fatalln(http.ListenAndServe(":9100", metricsMux))
}
