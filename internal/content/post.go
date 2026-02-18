package content

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// FlexTime handles both "2025-01-02 15:04:05" and "2025-01-02" date formats.
type FlexTime struct {
	time.Time
}

func (ft *FlexTime) UnmarshalYAML(value *yaml.Node) error {
	layouts := []string{
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}
	for _, layout := range layouts {
		t, err := time.Parse(layout, value.Value)
		if err == nil {
			ft.Time = t
			return nil
		}
	}
	return fmt.Errorf("cannot parse date: %s", value.Value)
}

type Post struct {
	// Frontmatter
	Title        string   `yaml:"title"`
	Date         FlexTime `yaml:"date"`
	Update       FlexTime `yaml:"update"`
	Author       string   `yaml:"author"`
	Cover        string   `yaml:"cover"`
	FeatureImage string   `yaml:"featureImage"`
	Tags         []string `yaml:"tags"`
	Categories   []string `yaml:"categories"`
	Preview      string   `yaml:"preview"`
	Draft        bool     `yaml:"draft"`
	Top          bool     `yaml:"top"`
	Type         string   `yaml:"type"`
	Hide         bool     `yaml:"hide"`
	TOC          bool     `yaml:"toc"`

	// Derived
	Slug        string
	Content     string
	HTMLContent template.HTML
	Filename    string
}

// GetCover returns the cover image path. Falls back to featureImage if cover is empty.
func (p *Post) GetCover() string {
	if p.Cover != "" {
		return p.Cover
	}
	return p.FeatureImage
}

// GetPreview returns the preview text, falling back to the first paragraph of content.
func (p *Post) GetPreview() string {
	if p.Preview != "" {
		return p.Preview
	}
	lines := strings.Split(p.Content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "!") || strings.HasPrefix(line, "{{") || strings.HasPrefix(line, "```") {
			continue
		}
		if len(line) > 200 {
			return line[:200] + "..."
		}
		return line
	}
	return ""
}

// CollectTags builds a map of tag name -> posts for that tag.
func CollectTags(posts []*Post) map[string][]*Post {
	tags := make(map[string][]*Post)
	for _, p := range posts {
		for _, tag := range p.Tags {
			tags[tag] = append(tags[tag], p)
		}
	}
	return tags
}

// CollectCategories builds a map of category name -> posts for that category.
func CollectCategories(posts []*Post) map[string][]*Post {
	cats := make(map[string][]*Post)
	for _, p := range posts {
		for _, cat := range p.Categories {
			cats[cat] = append(cats[cat], p)
		}
	}
	return cats
}
