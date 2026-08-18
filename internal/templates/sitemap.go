package templates

import (
	"encoding/xml"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ruedigerp/newblog/internal/content"
)

const sitemapNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

type sitemapURL struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	LastMod    string   `xml:"lastmod,omitempty"`
	ChangeFreq string   `xml:"changefreq,omitempty"`
	Priority   string   `xml:"priority,omitempty"`
}

type sitemapURLSet struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

// RenderSitemap writes a sitemaps.org 0.9 XML sitemap covering the homepage,
// its pagination pages, all posts, all static pages and the tag/category
// listings — i.e. every URL the generator produces as HTML.
func RenderSitemap(w io.Writer, site SiteData, posts []*content.Post, pages []*content.Page, postsPerPage int) error {
	base := strings.TrimRight(site.BaseURL, "/")

	// Newest post date doubles as lastmod for the index pages.
	siteMod := ""
	if len(posts) > 0 {
		siteMod = sitemapDate(postModTime(posts[0]))
	}

	urls := []sitemapURL{{
		Loc:        base + "/",
		LastMod:    siteMod,
		ChangeFreq: "daily",
		Priority:   "1.0",
	}}

	// Pagination: /page/2/, /page/3/, ... (page 1 is the homepage)
	if postsPerPage > 0 {
		totalPages := (len(posts) + postsPerPage - 1) / postsPerPage
		for page := 2; page <= totalPages; page++ {
			urls = append(urls, sitemapURL{
				Loc:        base + "/page/" + strconv.Itoa(page) + "/",
				LastMod:    siteMod,
				ChangeFreq: "weekly",
				Priority:   "0.4",
			})
		}
	}

	// Posts
	for _, p := range posts {
		urls = append(urls, sitemapURL{
			Loc:        base + "/posts/" + escapePath(p.Slug) + "/",
			LastMod:    sitemapDate(postModTime(p)),
			ChangeFreq: "monthly",
			Priority:   "0.8",
		})
	}

	// Static pages
	for _, pg := range pages {
		urls = append(urls, sitemapURL{
			Loc:        base + "/sites/" + escapePath(pg.Slug) + "/",
			ChangeFreq: "monthly",
			Priority:   "0.6",
		})
	}

	// Tag and category listings
	tagMap := content.CollectTags(posts)
	catMap := content.CollectCategories(posts)
	if len(tagMap) > 0 {
		urls = append(urls, sitemapURL{
			Loc:        base + "/tags/",
			LastMod:    siteMod,
			ChangeFreq: "weekly",
			Priority:   "0.5",
		})
		urls = append(urls, taxonomyURLs(base, "tags", tagMap)...)
	}
	if len(catMap) > 0 {
		urls = append(urls, sitemapURL{
			Loc:        base + "/categories/",
			LastMod:    siteMod,
			ChangeFreq: "weekly",
			Priority:   "0.5",
		})
		urls = append(urls, taxonomyURLs(base, "categories", catMap)...)
	}

	set := sitemapURLSet{Xmlns: sitemapNS, URLs: urls}

	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(set); err != nil {
		return err
	}
	_, err := w.Write([]byte("\n"))
	return err
}

// taxonomyURLs builds the per-tag / per-category URLs, sorted for stable output.
func taxonomyURLs(base, prefix string, m map[string][]*content.Post) []sitemapURL {
	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)

	out := make([]sitemapURL, 0, len(names))
	for _, name := range names {
		lastMod := ""
		for _, p := range m[name] {
			if d := sitemapDate(postModTime(p)); d > lastMod {
				lastMod = d
			}
		}
		out = append(out, sitemapURL{
			// Links in the themes use the lowercased name, same as the generator's dirs.
			Loc:        base + "/" + prefix + "/" + escapePath(strings.ToLower(name)) + "/",
			LastMod:    lastMod,
			ChangeFreq: "weekly",
			Priority:   "0.4",
		})
	}
	return out
}

// postModTime prefers the update date over the publication date.
func postModTime(p *content.Post) time.Time {
	if !p.Update.Time.IsZero() {
		return p.Update.Time
	}
	return p.Date.Time
}

// sitemapDate formats a W3C date (YYYY-MM-DD); zero times yield "".
func sitemapDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}

// escapePath percent-encodes each segment of a slug so tags with spaces or
// umlauts stay valid URLs.
func escapePath(slug string) string {
	parts := strings.Split(slug, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	return strings.Join(parts, "/")
}
