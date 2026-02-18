package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type MenuItem struct {
	Label    string     `yaml:"label"`
	URL      string     `yaml:"url,omitempty"`
	Children []MenuItem `yaml:"children,omitempty"`
}

type Config struct {
	Title        string     `yaml:"title"`
	Description  string     `yaml:"description"`
	Author       string     `yaml:"author"`
	BaseURL      string     `yaml:"baseURL"`
	Port         int        `yaml:"port"`
	PostsPerPage int        `yaml:"postsPerPage"`
	ContentDir   string     `yaml:"contentDir"`
	SitesDir     string     `yaml:"sitesDir"`
	OutputDir    string     `yaml:"outputDir"`
	Menu         []MenuItem `yaml:"menu"`
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
		OutputDir:    "output",
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
