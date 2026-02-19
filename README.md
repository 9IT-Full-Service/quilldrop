# QuillDrop

**QuillDrop** is a modern, minimalist blog CMS written in Go. It combines the speed of a static site generator with the flexibility of a dynamic HTTP server — no external database, no JavaScript frameworks, no overhead.

🇩🇪 [Deutsche Version](README-de.md)

## Philosophy

> Write. Save. Published.

QuillDrop follows the principle of maximum simplicity: write Markdown files, save them — done. No build tool chaos, no Node.js, no database. A single Go binary handles everything.

## Features

### Dual-Mode Operation

QuillDrop supports two modes of operation in a single binary:

- **`quilldrop serve`** — Starts a dynamic HTTP server for local development and preview. Ideal for writing and instantly testing new posts.
- **`quilldrop generate`** — Generates a complete static website as HTML files. Perfect for deployment on Nginx, Apache, CDN, or GitHub Pages.

### Markdown with YAML Frontmatter

Posts and pages are written as simple Markdown files with YAML frontmatter:

```yaml
---
title: "My New Blog Post"
date: 2025-11-06 12:00:00
author: "Jane Doe"
cover: "/images/posts/2025/11/cover.webp"
tags: [Kubernetes, DevOps, Self-Hosted]
categories: [Tech]
preview: "A short preview of the post..."
draft: false
toc: true
---

# Post starts here

Regular Markdown with all the extras...
```

Supported frontmatter fields:

| Field | Description |
|-------|-------------|
| `title` | Post title |
| `date` | Publication date (multiple formats supported) |
| `update` | Last updated |
| `author` | Post author |
| `cover` / `featureImage` | Cover image (with fallback) |
| `tags` | List of tags |
| `categories` | List of categories |
| `preview` | Custom preview (otherwise auto-generated from first paragraph) |
| `draft` | Draft — will not be published |
| `toc` | Automatically generate table of contents |
| `hide` | Hide post |
| `top` | Pin post to top |

### Extended Markdown Rendering

QuillDrop uses [Goldmark](https://github.com/yuin/goldmark) as its Markdown engine with the following extensions:

- **GitHub Flavored Markdown (GFM)** — Tables, strikethrough, autolinks, task lists
- **Syntax Highlighting** — Over 200 programming languages with the Dracula theme via [Chroma](https://github.com/alecthomas/chroma)
- **Emoji Support** — Shortcodes like `:rocket:`, `:tada:`, `:satellite:`
- **Automatic Heading IDs** — For anchor linking and table of contents
- **Raw HTML** — Embed HTML directly in Markdown
- **Hugo Compatibility** — `{{</* rawhtml */>}}` shortcodes are processed automatically

### Responsive Design with Dark/Light Theme

The included theme offers:

- **Dark Mode as default** with a light alternative theme
- **Theme Toggle** with localStorage persistence (survives page reloads)
- **Futuristic Design** — Dark backgrounds, cyan accents, subtle glow effects
- **Responsive Layout** — Mobile-first, optimized for all screen sizes
- **Hamburger Navigation** on mobile devices with fullscreen overlay and dedicated stacking context
- **Dropdown Menus** for nested navigation (touch-optimized on mobile)
- **Integrated Search** — Magnifying glass in the navbar with Ctrl+K shortcut
- **Typography** — Inter as body font, JetBrains Mono for code and metadata

### Navigation and Menu

The navigation menu is fully configured via `config.yaml` and supports nested dropdown menus:

```yaml
menu:
  - label: "Home"
    url: "/"
  - label: "Projects"
    children:
      - label: "VM-Manager"
        url: "/sites/projects/vm-manager"
      - label: "VM-Tracker"
        url: "/sites/projects/vm-tracker"
      - label: "QuillDrop"
        url: "/sites/projects/quilldrop"
  - label: "About"
    url: "/sites/about"
  - label: "Tags"
    url: "/tags"
```

New menu items and submenus can be added at any time by simply extending the YAML configuration.

### Pagination

The homepage displays a configurable number of posts per page (default: 5). Pagination features:

- **Smart page numbering** — Shows first and last page, plus a window around the current page
- **Ellipsis** for many pages (1 ... 10 11 **12** 13 14 ... 23)
- **Newer/Older buttons** for quick navigation
- **Pretty URLs** — `/page/2`, `/page/3`, etc.
- SEO-friendly: `/page/1` automatically redirects to `/` (301)

### Tags and Categories

QuillDrop supports both tags and categories for content organization:

- **Tag overview** at `/tags/` with post count per tag
- **Tag pages** at `/tags/kubernetes/` with all posts for a tag
- **Category overview** at `/categories/` with post count per category
- **Category pages** at `/categories/tech/` with all posts in a category
- **Tag and category badges** on post cards and detail pages
- Tags and categories are read from YAML frontmatter (`tags`, `categories`)

### Full-Text Search

QuillDrop includes an integrated client-side search that works entirely without a backend:

- **Search index** — A `search-index.json` with all posts is generated during build
- **Lazy loading** — The search index is only loaded when the search is first opened
- **Multi-term search** — Multiple search terms are combined with AND
- **Fields** — Searches title, preview, tags, and categories
- **Keyboard shortcut** — `Ctrl+K` / `Cmd+K` opens the search
- **Magnifying glass in the navbar** — Click the search icon to open the search field
- **Debounce** — Search results appear after 200ms typing delay
- **Maximum 8 results** with highlighting of search terms
- **Escape** or click outside closes the search
- No external service, no framework — pure vanilla JavaScript

### Article Navigation

At the end of each blog post, navigation to the previous and next article is displayed:

- **Newer article** (← left) — Links to the chronologically newer post
- **Older article** (→ right) — Links to the chronologically older post
- On the newest article, only "Older article" is shown
- On the oldest article, only "Newer article" is shown
- Displays the title of the linked article

### Table of Contents

Posts can activate an automatically generated table of contents:

- Enabled via `toc: true` in the frontmatter
- Supports **H1, H2, and H3** headings
- **Relative indentation** — The TOC detects the minimum heading level and indents relative to it
- Automatic anchor links to the respective headings
- Generated client-side for fast page load times

### Static Pages

In addition to blog posts, QuillDrop supports static pages for:

- Legal notice, privacy policy
- About me / About
- Project pages (with subpages)
- Any additional pages

Pages are stored as Markdown files in the `sites/` directory. Nested directories are automatically recognized — e.g., `sites/projects/vm-tracker/index.md` becomes accessible at `/sites/projects/vm-tracker`.

### RSS Feed

Automatically generated RSS 2.0 feed at `/index.xml` with:

- The latest 20 posts
- Title, link, preview, and publication date
- RSS autodiscovery in the HTML head
- RSS icon in the navigation
- URL `/index.xml` for compatibility with existing blog setups

### Cover Images

Posts can define a cover image that is displayed both on the homepage (as a post card) and on the detail view:

- **21:9 aspect ratio** on post cards with zoom-on-hover effect
- **Full width** on the single post page
- **Lazy loading** for optimal performance
- **Fallback** from `cover` to `featureImage`

## Architecture

### Project Structure

```
quilldrop/
├── main.go                          # CLI entry point
├── config.yaml                      # Configuration
├── content/                         # Blog posts (Markdown)
│   ├── 2025-11-06-my-post.md
│   └── ...
├── sites/                           # Static pages
│   ├── about.md
│   ├── legal.md
│   └── projects/
│       └── my-project/
│           └── index.md
├── static/                          # Static assets
│   ├── css/style.css
│   ├── js/
│   │   ├── theme.js                 # Dark/Light toggle + TOC generator
│   │   └── search.js                # Client-side full-text search
│   └── images/
├── internal/
│   ├── config/config.go             # YAML config loader
│   ├── content/
│   │   ├── post.go                  # Post struct + FlexTime + Tags/Categories
│   │   ├── parser.go                # Markdown + frontmatter parser
│   │   └── page.go                  # Static pages parser
│   ├── server/server.go             # HTTP server
│   ├── generator/
│   │   ├── generator.go             # Static site generator
│   │   └── search.go                # Search index generator (JSON)
│   └── templates/
│       ├── render.go                # Template engine + functions
│       ├── rss.go                   # RSS feed generator
│       ├── base.html                # Base layout + navbar + search
│       ├── home.html                # Homepage + pagination
│       ├── post.html                # Single post + prev/next navigation
│       ├── page.html                # Static page
│       ├── tags.html                # Tag overview
│       ├── tag.html                 # Tag page
│       ├── categories.html          # Category overview
│       └── category.html            # Category page
└── output/                          # Generated static files
```

### Technology Stack

| Component | Technology |
|-----------|------------|
| Language | Go (standard library + minimal dependencies) |
| HTTP Server | `net/http` (Go standard library) |
| Templates | `html/template` with `embed.FS` |
| Markdown | Goldmark + GFM + Emoji + Chroma |
| Configuration | YAML via `gopkg.in/yaml.v3` |
| Syntax Highlighting | Chroma (Dracula theme) |
| Fonts | Inter + JetBrains Mono (Google Fonts) |
| CSS | Vanilla CSS with custom properties |
| JavaScript | Vanilla JS — theme toggle, search, TOC (no framework) |

### Dependencies

QuillDrop intentionally has minimal dependencies — **no web framework**, **no CSS framework**, **no JS framework**:

- `github.com/yuin/goldmark` — Markdown parser (CommonMark compliant)
- `github.com/yuin/goldmark-emoji` — Emoji shortcodes
- `github.com/yuin/goldmark-highlighting/v2` — Syntax highlighting
- `github.com/alecthomas/chroma/v2` — Syntax highlighting engine
- `gopkg.in/yaml.v3` — YAML parser

### Embedded Assets

All HTML templates are embedded directly into the binary via Go's `//go:embed` directive. This means:

- **Single binary** — No external template files needed
- **Fast startup** — No filesystem access for templates
- **Easy deployment** — One binary + config + content = done

## Configuration

All configuration is done via a single `config.yaml`:

```yaml
title: "My Blog"
description: "Tech Blog - DevOps, Kubernetes, Self-Hosted"
author: "Jane Doe"
baseURL: "https://my-blog.com"
port: 8080
postsPerPage: 5
contentDir: "content"
sitesDir: "sites"
outputDir: "output"

menu:
  - label: "Home"
    url: "/"
  - label: "Tags"
    url: "/tags"
  - label: "About"
    url: "/sites/about"
```

| Option | Default | Description |
|--------|---------|-------------|
| `title` | — | Website title |
| `description` | — | Description (meta tag + hero) |
| `author` | — | Website author |
| `baseURL` | — | Base URL for RSS and absolute links |
| `port` | `8080` | Port for the dynamic server |
| `postsPerPage` | `5` | Number of posts per page |
| `contentDir` | `content` | Directory for blog posts |
| `sitesDir` | `sites` | Directory for static pages |
| `outputDir` | `output` | Output directory for static generation |
| `menu` | `[]` | Navigation menu with optional submenus |

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/ruedigerp/quilldrop.git
cd quilldrop

# Download dependencies
go mod download

# Build the binary
go build -o quilldrop .
```

### Create a New Post

Create a new Markdown file in the `content/` directory:

```bash
touch content/2025-12-01-my-first-post.md
```

```markdown
---
title: "My First Post"
date: 2025-12-01 10:00:00
author: "Jane Doe"
tags: [Blog, QuillDrop]
preview: "This is my first post with QuillDrop!"
toc: false
---

# Welcome

This is my first post with **QuillDrop**.

```

### Local Preview

```bash
# Start the dynamic server
./quilldrop serve

# Or run directly with Go
go run . serve
```

Then open in the browser: [http://localhost:8080](http://localhost:8080)

### Generate Static Site

```bash
# Generate HTML files
./quilldrop generate

# Generated files are in output/
ls output/
```

The generated files in the `output/` directory can be deployed directly to a web server (Nginx, Apache, Caddy) or CDN.

## URL Schema

All URLs consistently use trailing slashes to avoid server-side redirects:

| URL | Description |
|-----|-------------|
| `/` | Homepage (latest N posts) |
| `/page/2/` | Page 2 of the post list |
| `/posts/2025-11-06-my-post/` | Single blog post |
| `/tags/` | Tag overview |
| `/tags/kubernetes/` | Posts with tag "Kubernetes" |
| `/categories/` | Category overview |
| `/categories/tech/` | Posts in category "Tech" |
| `/sites/about/` | Static page |
| `/sites/projects/vm-tracker/` | Nested project page |
| `/index.xml` | RSS feed |
| `/search-index.json` | Search index (JSON) |
| `/static/css/style.css` | Static assets |
| `/images/posts/2025/11/cover.webp` | Images |

## Why QuillDrop?

- **No database** — The filesystem is the only data source
- **No build pipeline** — A single `go build` and you're done
- **No JS frameworks** — Vanilla JavaScript for theme, search, and TOC
- **Minimal dependencies** — 5 Go packages, all focused on Markdown
- **Blazing fast** — Generates 100+ posts in under 3 seconds
- **Single binary** — Templates embedded, no runtime setup needed
- **Hugo compatible** — Existing Hugo posts with frontmatter just work
- **Dual-mode** — Development with server, production with static generator

## License

QuillDrop is open source.
