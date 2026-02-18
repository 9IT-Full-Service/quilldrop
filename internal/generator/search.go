package generator

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ruedigerp/newblog/internal/content"
)

type searchEntry struct {
	Title      string   `json:"title"`
	Slug       string   `json:"slug"`
	Date       string   `json:"date"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
	Preview    string   `json:"preview"`
	Cover      string   `json:"cover"`
	URL        string   `json:"url"`
}

func GenerateSearchIndex(posts []*content.Post, outputPath string) {
	entries := make([]searchEntry, 0, len(posts))

	months := []string{
		"", "Januar", "Februar", "März", "April", "Mai", "Juni",
		"Juli", "August", "September", "Oktober", "November", "Dezember",
	}

	for _, p := range posts {
		date := p.Date.Time
		dateStr := date.Format("02.") + " " + months[date.Month()] + " " + date.Format("2006")

		tags := p.Tags
		if tags == nil {
			tags = []string{}
		}
		cats := p.Categories
		if cats == nil {
			cats = []string{}
		}

		entries = append(entries, searchEntry{
			Title:      p.Title,
			Slug:       p.Slug,
			Date:       dateStr,
			Tags:       tags,
			Categories: cats,
			Preview:    p.GetPreview(),
			Cover:      p.GetCover(),
			URL:        "/posts/" + p.Slug + "/",
		})
	}

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		log.Printf("Error marshaling search index: %v", err)
		return
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		log.Printf("Error writing search index: %v", err)
	}
}
