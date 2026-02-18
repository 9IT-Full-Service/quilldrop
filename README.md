# QuillDrop

**QuillDrop** ist ein modernes, minimalistisches Blog-CMS, geschrieben in Go. Es kombiniert die Geschwindigkeit eines Static Site Generators mit der Flexibilität eines dynamischen HTTP-Servers - ohne externe Datenbank, ohne JavaScript-Frameworks, ohne Overhead.

## Philosophie

> Write. Save. Published.

QuillDrop folgt dem Prinzip der maximalen Einfachheit: Markdown-Dateien schreiben, speichern - fertig. Kein Build-Tool-Chaos, kein Node.js, keine Datenbank. Ein einzelnes Go-Binary erledigt alles.

## Features

### Dual-Mode Betrieb

QuillDrop unterstützt zwei Betriebsmodi in einem einzigen Binary:

- **`quilldrop serve`** - Startet einen dynamischen HTTP-Server für lokale Entwicklung und Vorschau. Ideal zum Schreiben und sofortigen Testen neuer Posts.
- **`quilldrop generate`** - Generiert eine komplette statische Website als HTML-Dateien. Perfekt für Deployment auf Nginx, Apache, CDN oder GitHub Pages.

### Markdown mit YAML-Frontmatter

Posts und Seiten werden als einfache Markdown-Dateien mit YAML-Frontmatter geschrieben:

```yaml
---
title: "Mein neuer Blogpost"
date: 2025-11-06 12:00:00
author: "Max Mustermann"
cover: "/images/posts/2025/11/cover.webp"
tags: [Kubernetes, DevOps, Self-Hosted]
categories: [Technik]
preview: "Kurze Vorschau des Posts..."
draft: false
toc: true
---

# Hier beginnt der Post

Normales Markdown mit allen Extras...
```

Unterstützte Frontmatter-Felder:

| Feld | Beschreibung |
|------|-------------|
| `title` | Titel des Posts |
| `date` | Veröffentlichungsdatum (mehrere Formate unterstützt) |
| `update` | Letzte Aktualisierung |
| `author` | Autor des Posts |
| `cover` / `featureImage` | Cover-Bild (mit Fallback) |
| `tags` | Liste von Tags |
| `categories` | Liste von Kategorien |
| `preview` | Benutzerdefinierte Vorschau (sonst automatisch aus erstem Absatz) |
| `draft` | Entwurf - wird nicht veröffentlicht |
| `toc` | Inhaltsverzeichnis automatisch generieren |
| `hide` | Post verstecken |
| `top` | Post oben anpinnen |

### Erweitertes Markdown-Rendering

QuillDrop nutzt [Goldmark](https://github.com/yuin/goldmark) als Markdown-Engine mit folgenden Erweiterungen:

- **GitHub Flavored Markdown (GFM)** - Tabellen, Strikethrough, Autolinks, Task-Listen
- **Syntax Highlighting** - Über 200 Programmiersprachen mit dem Dracula-Theme via [Chroma](https://github.com/alecthomas/chroma)
- **Emoji-Support** - Shortcodes wie `:rocket:`, `:tada:`, `:satellite:`
- **Automatische Heading-IDs** - Für Ankerverlinkung und Inhaltsverzeichnis
- **Raw HTML** - Einbettung von HTML direkt im Markdown
- **Hugo-Kompatibilität** - `{{</* rawhtml */>}}` Shortcodes werden automatisch verarbeitet

### Responsives Design mit Dark/Light Theme

Das mitgelieferte Theme bietet:

- **Dark Mode als Default** mit einem hellen Alternativ-Theme
- **Theme Toggle** mit localStorage-Persistenz (bleibt nach Reload erhalten)
- **Futuristisches Design** - Dunkle Hintergrunde, Cyan-Akzente, subtile Glow-Effekte
- **Responsive Layout** - Mobile-first, optimiert für alle Bildschirmgrößen
- **Hamburger-Navigation** auf mobilen Geräten mit Fullscreen-Overlay
- **Dropdown-Menus** für verschachtelte Navigation
- **Typographie** - Inter als Textfont, JetBrains Mono für Code und Metadaten

### Navigation und Menü

Das Navigationsmenü wird vollständig über die `config.yaml` konfiguriert und unterstützt verschachtelte Dropdown-Menüs:

```yaml
menu:
  - label: "Home"
    url: "/"
  - label: "Projekte"
    children:
      - label: "VM-Manager"
        url: "/sites/projekte/vm-manager"
      - label: "VM-Tracker"
        url: "/sites/projekte/vm-tracker"
      - label: "QuillDrop"
        url: "/sites/projekte/quilldrop"
  - label: "Über mich"
    url: "/sites/ueber-mich"
  - label: "Tags"
    url: "/tags"
```

Neue Menüpunkte und Untermenüs können jederzeit durch einfaches Erweitern der YAML-Konfiguration hinzugefügt werden.

### Pagination

Die Startseite zeigt eine konfigurierbare Anzahl von Posts pro Seite (Standard: 5). Die Pagination bietet:

- **Intelligente Seitennummerierung** - Zeigt erste und letzte Seite, plus ein Fenster um die aktuelle Seite herum
- **Ellipsis** bei vielen Seiten (1 ... 10 11 **12** 13 14 ... 23)
- **Neuere/Ältere Buttons** für schnelle Navigation
- **Pretty URLs** - `/page/2`, `/page/3`, etc.
- SEO-freundlich: `/page/1` wird automatisch auf `/` umgeleitet (301)

### Tags und Kategorien

- **Tag-Übersicht** unter `/tags` mit Anzahl der Posts pro Tag
- **Tag-Seiten** unter `/tags/kubernetes` mit allen Posts eines Tags
- **Tag-Badges** auf Post-Cards und Einzelseiten

### Statische Seiten

Neben Blog-Posts unterstützt QuillDrop statische Seiten für:

- Impressum, Datenschutzerklärung
- Über mich / About
- Projektseiten (mit Unterseiten)
- Beliebige weitere Seiten

Seiten werden als Markdown-Dateien im `sites/`-Verzeichnis abgelegt. Verschachtelte Verzeichnisse werden automatisch erkannt - z.B. wird `sites/projekte/vm-tracker/index.md` unter `/sites/projekte/vm-tracker` erreichbar.

### RSS Feed

Automatisch generierter RSS 2.0 Feed unter `/feed.xml` mit:

- Den letzten 20 Posts
- Titel, Link, Vorschau und Veröffentlichungsdatum
- RSS-Autodiscovery im HTML-Head
- RSS-Icon in der Navigation

### Cover-Bilder

Posts können ein Cover-Bild definieren, das sowohl auf der Startseite (als Post-Card) als auch auf der Einzelansicht angezeigt wird:

- **21:9 Aspect Ratio** auf Post-Cards mit Zoom-on-Hover Effekt
- **Volle Breite** auf der Einzelpost-Seite
- **Lazy Loading** für optimale Performance
- **Fallback** von `cover` auf `featureImage`

## Architektur

### Projektstruktur

```
quilldrop/
├── main.go                          # CLI Entry Point
├── config.yaml                      # Konfiguration
├── content/                         # Blog-Posts (Markdown)
│   ├── 2025-11-06-mein-post.md
│   └── ...
├── sites/                           # Statische Seiten
│   ├── ueber-mich.md
│   ├── impressum.md
│   └── projekte/
│       └── mein-projekt/
│           └── index.md
├── static/                          # Statische Assets
│   ├── css/style.css
│   ├── js/theme.js
│   └── images/
├── internal/
│   ├── config/config.go             # YAML Config Loader
│   ├── content/
│   │   ├── post.go                  # Post Struct + FlexTime
│   │   ├── parser.go                # Markdown + Frontmatter Parser
│   │   └── page.go                  # Statische Seiten Parser
│   ├── server/server.go             # HTTP Server
│   ├── generator/generator.go       # Static Site Generator
│   └── templates/
│       ├── render.go                # Template Engine + Functions
│       ├── rss.go                   # RSS Feed Generator
│       ├── base.html                # Base Layout
│       ├── home.html                # Homepage + Pagination
│       ├── post.html                # Einzelner Post
│       ├── page.html                # Statische Seite
│       ├── tags.html                # Tag-Übersicht
│       └── tag.html                 # Tag-Seite
└── output/                          # Generierte statische Dateien
```

### Technologie-Stack

| Komponente | Technologie |
|-----------|-------------|
| Sprache | Go (Standard Library + minimale Dependencies) |
| HTTP Server | `net/http` (Go Standard Library) |
| Templates | `html/template` mit `embed.FS` |
| Markdown | Goldmark + GFM + Emoji + Chroma |
| Konfiguration | YAML via `gopkg.in/yaml.v3` |
| Syntax Highlighting | Chroma (Dracula Theme) |
| Fonts | Inter + JetBrains Mono (Google Fonts) |
| CSS | Vanilla CSS mit Custom Properties |
| JavaScript | Vanilla JS (kein Framework) |

### Dependencies

QuillDrop hat bewusst minimale Abhängigkeiten - **kein Web-Framework**, **kein CSS-Framework**, **kein JS-Framework**:

- `github.com/yuin/goldmark` - Markdown Parser (CommonMark-konform)
- `github.com/yuin/goldmark-emoji` - Emoji Shortcodes
- `github.com/yuin/goldmark-highlighting/v2` - Syntax Highlighting
- `github.com/alecthomas/chroma/v2` - Syntax Highlighting Engine
- `gopkg.in/yaml.v3` - YAML Parser

### Embedded Assets

Alle HTML-Templates werden via Go's `//go:embed` Directive direkt in das Binary eingebettet. Das bedeutet:

- **Einzelnes Binary** - Keine externen Template-Dateien nötig
- **Schneller Start** - Kein Dateisystem-Zugriff für Templates
- **Einfaches Deployment** - Ein Binary + Config + Content = fertig

## Konfiguration

Die gesamte Konfiguration erfolgt über eine einzige `config.yaml`:

```yaml
title: "Mein Blog"
description: "Tech Blog - DevOps, Kubernetes, Self-Hosted"
author: "Max Mustermann"
baseURL: "https://mein-blog.de"
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
  - label: "Über mich"
    url: "/sites/ueber-mich"
```

| Option | Default | Beschreibung |
|--------|---------|-------------|
| `title` | - | Titel der Website |
| `description` | - | Beschreibung (Meta-Tag + Hero) |
| `author` | - | Autor der Website |
| `baseURL` | - | Basis-URL für RSS und absolute Links |
| `port` | `8080` | Port für den dynamischen Server |
| `postsPerPage` | `5` | Anzahl Posts pro Seite |
| `contentDir` | `content` | Verzeichnis für Blog-Posts |
| `sitesDir` | `sites` | Verzeichnis für statische Seiten |
| `outputDir` | `output` | Ausgabeverzeichnis für statische Generierung |
| `menu` | `[]` | Navigationsmenü mit optionalen Untermenüs |

## Schnellstart

### Installation

```bash
# Repository klonen
git clone https://github.com/ruedigerp/quilldrop.git
cd quilldrop

# Dependencies laden
go mod download

# Binary bauen
go build -o quilldrop .
```

### Neuen Post erstellen

Eine neue Markdown-Datei im `content/`-Verzeichnis anlegen:

```bash
touch content/2025-12-01-mein-erster-post.md
```

```markdown
---
title: "Mein erster Post"
date: 2025-12-01 10:00:00
author: "Max Mustermann"
tags: [Blog, QuillDrop]
preview: "Das ist mein erster Post mit QuillDrop!"
toc: false
---

# Willkommen

Das ist mein erster Post mit **QuillDrop**.

## Code-Beispiel

```go
fmt.Println("Hello QuillDrop!")
```
```

### Lokale Vorschau

```bash
# Dynamischen Server starten
./quilldrop serve

# Oder direkt mit Go
go run . serve
```

Dann im Browser: [http://localhost:8080](http://localhost:8080)

### Statische Seite generieren

```bash
# HTML-Dateien generieren
./quilldrop generate

# Generierte Dateien befinden sich in output/
ls output/
```

Die generierten Dateien im `output/`-Verzeichnis können direkt auf einen Webserver (Nginx, Apache, Caddy) oder CDN deployed werden.

## URL-Schema

| URL | Beschreibung |
|-----|-------------|
| `/` | Startseite (letzte N Posts) |
| `/page/2` | Seite 2 der Post-Liste |
| `/posts/2025-11-06-mein-post` | Einzelner Blog-Post |
| `/tags` | Tag-Übersicht |
| `/tags/kubernetes` | Posts mit Tag "Kubernetes" |
| `/sites/ueber-mich` | Statische Seite |
| `/sites/projekte/vm-tracker` | Verschachtelte Projektseite |
| `/feed.xml` | RSS Feed |
| `/static/css/style.css` | Statische Assets |
| `/images/posts/2025/11/cover.webp` | Bilder |

## Warum QuillDrop?

- **Keine Datenbank** - Dateisystem als einzige Datenquelle
- **Keine Build-Pipeline** - Ein `go build` und fertig
- **Keine JS-Frameworks** - Vanilla JavaScript, unter 90 Zeilen
- **Minimale Dependencies** - 5 Go-Packages, alle fokussiert auf Markdown
- **Blitzschnell** - Generiert 100+ Posts in unter 3 Sekunden
- **Einzelnes Binary** - Templates eingebettet, kein Runtime-Setup
- **Hugo-kompatibel** - Bestehende Hugo-Posts mit Frontmatter funktionieren
- **Dual-Mode** - Entwicklung mit Server, Produktion mit Static Generator

## Lizenz

QuillDrop ist Open Source.
