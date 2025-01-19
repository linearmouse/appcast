package main

import (
	"net/http"
	"regexp"
)

var uaSparkle2Re = regexp.MustCompile(`\bSparkle\/2\.0\.[01]\b`)

// Sparkle 2.0.0 will check the certificate.
// After v0.6.1, LinearMouse uses a new certificate, so the early versions
// can only upgrade to v0.6.1 directly.
func handleSparkle2(w http.ResponseWriter, r *http.Request) bool {
	if !uaSparkle2Re.MatchString(r.UserAgent()) {
		return false
	}

	appcast := `<rss xmlns:sparkle="http://www.andymatuschak.org/xml-namespaces/sparkle" xmlns:dc="http://purl.org/dc/elements/1.1/" version="2.0">
  <channel>
    <title>LinearMouse</title>
    <link>https://github.com/linearmouse/linearmouse</link>
    <item>
      <title>v0.6.1</title>
      <sparkle:version>0.6.1</sparkle:version>
      <pubDate>Fri, 10 Jun 2022 06:09:49 GMT</pubDate>
      <link>https://github.com/linearmouse/linearmouse/releases/tag/v0.6.1</link>
      <description><![CDATA[
<h2 dir="auto">What's Changed</h2>
<h3 dir="auto">Bug fixes</h3>
<ul dir="auto">
<li>Fix the launch at login issue in some cases by <a class="user-mention notranslate" data-hovercard-type="user" data-hovercard-url="/users/lujjjh/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/lujjjh">@lujjjh</a> in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="1263448488" data-permission-text="Title is private" data-url="https://github.com/linearmouse/linearmouse/issues/132" data-hovercard-type="pull_request" data-hovercard-url="/linearmouse/linearmouse/pull/132/hovercard" href="https://github.com/linearmouse/linearmouse/pull/132">#132</a></li>
<li>Fix freezing after granting accessibility permission by <a class="user-mention notranslate" data-hovercard-type="user" data-hovercard-url="/users/lujjjh/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/lujjjh">@lujjjh</a> in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="1265513799" data-permission-text="Title is private" data-url="https://github.com/linearmouse/linearmouse/issues/136" data-hovercard-type="pull_request" data-hovercard-url="/linearmouse/linearmouse/pull/136/hovercard" href="https://github.com/linearmouse/linearmouse/pull/136">#136</a></li>
</ul>
<h3 dir="auto">Other changes</h3>
<ul dir="auto">
<li>Update translations (Italian and Portuguese, Brazilian) by <a class="user-mention notranslate" data-hovercard-type="user" data-hovercard-url="/users/LuigiPiccoli17/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/LuigiPiccoli17">@LuigiPiccoli17</a> in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="1260318189" data-permission-text="Title is private" data-url="https://github.com/linearmouse/linearmouse/issues/131" data-hovercard-type="pull_request" data-hovercard-url="/linearmouse/linearmouse/pull/131/hovercard" href="https://github.com/linearmouse/linearmouse/pull/131">#131</a></li>
<li>Universal back and forward: Ignore Dota 2 by <a class="user-mention notranslate" data-hovercard-type="user" data-hovercard-url="/users/aramann/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/aramann">@aramann</a> in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="1263521521" data-permission-text="Title is private" data-url="https://github.com/linearmouse/linearmouse/issues/133" data-hovercard-type="pull_request" data-hovercard-url="/linearmouse/linearmouse/pull/133/hovercard" href="https://github.com/linearmouse/linearmouse/pull/133">#133</a></li>
<li>Add a guide on how to grant Accessibility permission by <a class="user-mention notranslate" data-hovercard-type="user" data-hovercard-url="/users/lujjjh/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/lujjjh">@lujjjh</a> in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="1265978312" data-permission-text="Title is private" data-url="https://github.com/linearmouse/linearmouse/issues/137" data-hovercard-type="pull_request" data-hovercard-url="/linearmouse/linearmouse/pull/137/hovercard" href="https://github.com/linearmouse/linearmouse/pull/137">#137</a></li>
</ul>
<h2 dir="auto">New Contributors</h2>
<ul dir="auto">
<li><a class="user-mention notranslate" data-hovercard-type="user" data-hovercard-url="/users/aramann/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/aramann">@aramann</a> made their first contribution in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="1263521521" data-permission-text="Title is private" data-url="https://github.com/linearmouse/linearmouse/issues/133" data-hovercard-type="pull_request" data-hovercard-url="/linearmouse/linearmouse/pull/133/hovercard" href="https://github.com/linearmouse/linearmouse/pull/133">#133</a></li>
<li><a class="user-mention notranslate" data-hovercard-type="user" data-hovercard-url="/users/LuigiPiccoli17/hovercard" data-octo-click="hovercard-link-click" data-octo-dimensions="link_type:self" href="https://github.com/LuigiPiccoli17">@LuigiPiccoli17</a> made their first contribution in <a class="issue-link js-issue-link" data-error-text="Failed to load title" data-id="1260318189" data-permission-text="Title is private" data-url="https://github.com/linearmouse/linearmouse/issues/131" data-hovercard-type="pull_request" data-hovercard-url="/linearmouse/linearmouse/pull/131/hovercard" href="https://github.com/linearmouse/linearmouse/pull/131">#131</a></li>
</ul>
<p dir="auto"><strong>Full Changelog</strong>: <a class="commit-link" href="https://github.com/linearmouse/linearmouse/compare/v0.6.0...v0.6.1"><tt>v0.6.0...v0.6.1</tt></a></p>]]></description>
      <enclosure url="https://dl.linearmouse.org/v0.6.1/LinearMouse.dmg" type="application/octet-stream"/>
    </item>
  </channel>
</rss>`

	w.Header().Set("Content-Type", "text/xml")
	w.Header().Set("Cache-Control", "max-age=0")

	w.Write([]byte(appcast))

	return true
}
