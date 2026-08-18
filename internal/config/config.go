package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type MenuItem struct {
	Label    string     `yaml:"label"`
	URL      string     `yaml:"url,omitempty"`
	Children []MenuItem `yaml:"children,omitempty"`
}

// LLMsConfig toggles generation/serving of the llms.txt convention files.
// Both default to true so behaviour without the block stays unchanged.
type LLMsConfig struct {
	Enabled *bool `yaml:"enabled,omitempty"` // llms.txt
	Full    *bool `yaml:"full,omitempty"`    // llms-full.txt
}

// EnabledOrDefault returns whether llms.txt should be produced (default: true).
func (l LLMsConfig) EnabledOrDefault() bool {
	if l.Enabled == nil {
		return true
	}
	return *l.Enabled
}

// FullOrDefault returns whether llms-full.txt should be produced (default: true).
func (l LLMsConfig) FullOrDefault() bool {
	if l.Full == nil {
		return true
	}
	return *l.Full
}

// SitemapConfig toggles generation/serving of sitemap.xml for search engines.
// Defaults to true so behaviour without the block stays unchanged.
type SitemapConfig struct {
	Enabled *bool `yaml:"enabled,omitempty"` // sitemap.xml
}

// EnabledOrDefault returns whether sitemap.xml should be produced (default: true).
func (s SitemapConfig) EnabledOrDefault() bool {
	if s.Enabled == nil {
		return true
	}
	return *s.Enabled
}

// RobotsConfig toggles generation/serving of robots.txt and lets the site
// define its own Allow/Disallow rules. Defaults to true so behaviour without
// the block stays unchanged.
type RobotsConfig struct {
	Enabled   *bool    `yaml:"enabled,omitempty"`   // robots.txt
	UserAgent string   `yaml:"userAgent,omitempty"` // default: "*"
	Allow     []string `yaml:"allow,omitempty"`
	Disallow  []string `yaml:"disallow,omitempty"`
}

// EnabledOrDefault returns whether robots.txt should be produced (default: true).
func (r RobotsConfig) EnabledOrDefault() bool {
	if r.Enabled == nil {
		return true
	}
	return *r.Enabled
}

// UserAgentOrDefault returns the User-agent line's value (default: "*").
func (r RobotsConfig) UserAgentOrDefault() string {
	if r.UserAgent == "" {
		return "*"
	}
	return r.UserAgent
}

type Config struct {
	Title        string        `yaml:"title"`
	Description  string        `yaml:"description"`
	Author       string        `yaml:"author"`
	BaseURL      string        `yaml:"baseURL"`
	Port         int           `yaml:"port"`
	PostsPerPage int           `yaml:"postsPerPage"`
	ContentDir   string        `yaml:"contentDir"`
	SitesDir     string        `yaml:"sitesDir"`
	StaticDir    string        `yaml:"staticDir"`
	OutputDir    string        `yaml:"outputDir"`
	ThemesDir    string        `yaml:"themesDir"`
	Theme        string        `yaml:"theme"`
	Menu         []MenuItem    `yaml:"menu"`
	LLMs         LLMsConfig    `yaml:"llms"`
	Sitemap      SitemapConfig `yaml:"sitemap"`
	Robots       RobotsConfig  `yaml:"robots"`
}

// ThemeDir returns the directory of the active theme, e.g. "themes/default".
func (c *Config) ThemeDir() string {
	return filepath.Join(c.ThemesDir, c.Theme)
}

// ThemeTemplatesDir returns the template directory of the active theme.
func (c *Config) ThemeTemplatesDir() string {
	return filepath.Join(c.ThemeDir(), "templates")
}

// ThemeStaticDir returns the static asset directory of the active theme.
// The directory is optional; a theme may ship without own assets.
func (c *Config) ThemeStaticDir() string {
	return filepath.Join(c.ThemeDir(), "static")
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{
		Port:         8080,
		PostsPerPage: 5,
		ContentDir:   "content",
		SitesDir:     "sites",
		StaticDir:    "static",
		OutputDir:    "output",
		ThemesDir:    "themes",
		Theme:        "default",
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	// Empty values in config.yaml must not override the defaults.
	if cfg.StaticDir == "" {
		cfg.StaticDir = "static"
	}
	if cfg.ThemesDir == "" {
		cfg.ThemesDir = "themes"
	}
	if cfg.Theme == "" {
		cfg.Theme = "default"
	}
	return cfg, nil
}
