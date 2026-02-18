package templates

import (
	"encoding/xml"
	"io"
	"time"

	"github.com/ruedigerp/newblog/internal/content"
)

type rssChannel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	Language      string    `xml:"language"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

func RenderRSS(w io.Writer, site SiteData, posts []*content.Post) error {
	items := make([]rssItem, 0, len(posts))
	for _, p := range posts {
		link := site.BaseURL + "/posts/" + p.Slug + "/"
		desc := p.GetPreview()
		items = append(items, rssItem{
			Title:       p.Title,
			Link:        link,
			Description: desc,
			PubDate:     p.Date.Format(time.RFC1123Z),
			GUID:        link,
		})
	}

	buildDate := time.Now().Format(time.RFC1123Z)
	if len(posts) > 0 {
		buildDate = posts[0].Date.Format(time.RFC1123Z)
	}

	feed := rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:         site.Title,
			Link:          site.BaseURL,
			Description:   site.Description,
			Language:      "de-de",
			LastBuildDate: buildDate,
			Items:         items,
		},
	}

	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(feed)
}
