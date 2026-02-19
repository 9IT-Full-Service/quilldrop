---
title: 'QuillDrop - Mein eigenes Blog-CMS ist fertig'
date: 2026-06-19 10:00:00
author: ruediger
cover: "/images/posts/2026/02/quilldrop.webp"
tags: [Go, Blog, CMS, QuillDrop, Open-Source, Self-Hosted]
categories:
  - Internet
  - Projekte
preview: "Nach etlichen CMS- und Static-Site-Generator-Abenteuern mit Jekyll, Hugo und InkProject habe ich mein Blog jetzt komplett auf mein eigenes CMS migriert: QuillDrop. Geschrieben in Go, ohne Frameworks, ohne Datenbank - ein einziges Binary für alles."
draft: false
top: false
type: post
hide: false
toc: true
---

# QuillDrop - Mein eigenes Blog-CMS ist fertig

Jekyll, Octopress, Hugo, InkProject und dann noch mein eigenes CMS in Ruby - die Liste der Systeme, auf denen dieser Blog im Laufe der Jahre gelaufen ist, ist lang. Jetzt kommt ein neues dazu, aber dieses Mal ist es anders. Dieses Mal habe ich das Ding komplett selbst geschrieben: **QuillDrop**.

## Warum schon wieder ein neues CMS?

Die Kurzfassung: Weil keines genau das gemacht hat, was ich wollte - und gleichzeitig nichts, was ich nicht wollte.

Hugo ist großartig, aber die Template-Sprache ist eine Welt für sich und das Debugging bei Problemen kann schnell frustrierend werden. Mein vorheriges Ruby-basiertes CMS hat funktioniert, war aber mit der Zeit zu komplex geworden. Was ich wollte, war simpel: Markdown-Dateien schreiben, ein Binary starten, fertig. Keine Node-Module, keine Build-Pipeline, keine externe Datenbank.

Also habe ich QuillDrop gebaut.

## Was ist QuillDrop?

QuillDrop ist ein Blog-CMS, geschrieben in Go. Ein einzelnes Binary, das zwei Dinge kann:

- **`quilldrop serve`** startet einen lokalen HTTP-Server zum Entwickeln und Vorschauen
- **`quilldrop generate`** generiert eine komplett statische Website für das Deployment

Das war es. Keine Plugins, keine Themes zum Installieren, kein Package-Manager. Die Templates sind direkt im Binary eingebettet (`go:embed`), das CSS ist Vanilla CSS, das JavaScript ist Vanilla JavaScript. Die einzigen Go-Dependencies sind Goldmark für Markdown-Rendering und ein YAML-Parser für die Konfiguration.

## Die komplette Migration

Der Blog ist jetzt vollständig migriert. Alle Posts, alle statischen Seiten, alle Bilder, alle Funktionen die vorher schon da waren - alles läuft jetzt auf QuillDrop. Und das sind nicht wenige:

### Posts und Inhalte

Alle bestehenden Markdown-Posts wurden 1:1 übernommen. Das YAML-Frontmatter ist kompatibel mit Hugo, sodass die Migration im Grunde nur ein Kopieren der Dateien war. Cover-Bilder, Tags, Kategorien, Entwürfe - alles wird unterstützt.

### Statische Seiten

Seiten wie "Über mich" und die Projektseiten (VM-Tracker, QuillDrop selbst) liegen als Markdown im `sites/`-Verzeichnis. Verschachtelte Verzeichnisse werden automatisch erkannt.

### Navigation mit Dropdown-Menüs

Die komplette Navigation wird über die `config.yaml` konfiguriert. Dropdown-Menüs für Unterseiten wie bei "Projekte" sind kein Problem - einfach `children` definieren und fertig.

### Tags und Kategorien

Tags hatte das vorherige System auch, Kategorien sind neu dazugekommen. Beides wird aus dem Frontmatter gelesen und bekommt eigene Übersichtsseiten unter `/tags/` und `/categories/`.

### Dark/Light Theme

Der Dark Mode ist Standard, aber es gibt einen Toggle für ein helles Theme. Die Auswahl wird im Browser gespeichert und bleibt nach einem Reload erhalten.

### Pagination

Bei mittlerweile über 80 Posts ist eine Pagination essentiell. QuillDrop zeigt eine intelligente Seitennummerierung mit Ellipsis: `1 ... 10 11 12 13 14 ... 23`. Die erste Seite wird SEO-freundlich von `/page/1/` auf `/` umgeleitet.

### RSS Feed

Der RSS Feed liegt unter `/index.xml` - die gleiche URL wie vorher, damit bestehende Abonnenten nichts umstellen müssen.

### Volltextsuche

Das war eines der Features, die ich vom vorherigen Blog unbedingt mitnehmen wollte. Die Suche funktioniert komplett client-seitig: Beim Generieren wird eine `search-index.json` erstellt, die dann im Browser durchsucht wird. Lazy Loading sorgt dafür, dass der Index erst geladen wird, wenn die Suche geöffnet wird. Unterstützt werden mehrere Suchbegriffe (UND-Verknüpfung), und mit `Ctrl+K` kann man die Suche direkt öffnen.

### Inhaltsverzeichnis

Posts mit `toc: true` im Frontmatter bekommen ein automatisch generiertes Inhaltsverzeichnis. Das TOC erkennt H1-, H2- und H3-Überschriften und rückt relativ ein.

### Artikel-Navigation

Am Ende jedes Posts gibt es jetzt Links zum neueren und älteren Artikel. Beim neuesten Post wird nur "Älterer Artikel" angezeigt, beim ältesten nur "Neuerer Artikel".

## Technik unter der Haube

Der Stack ist bewusst minimal gehalten:

- **Go** mit der Standard Library für HTTP-Server und Templates
- **Goldmark** für Markdown-Rendering mit GFM, Emoji-Support und Syntax-Highlighting über Chroma
- **Vanilla CSS** mit Custom Properties für das Theming
- **Vanilla JavaScript** für Theme-Toggle, Suche und TOC-Generierung
- **Keine Datenbank** - das Dateisystem ist die einzige Datenquelle

Die gesamte Konfiguration läuft über eine einzige `config.yaml`. Keine versteckten Konfigurationsdateien, keine Environment-Variables die man vergessen könnte.

## Dual-Mode: Entwicklung und Produktion

Für die lokale Entwicklung starte ich `quilldrop serve` und sehe sofort jede Änderung im Browser. Für das Deployment auf den Webserver läuft `quilldrop generate` und spuckt statische HTML-Dateien aus, die direkt auf Nginx, Apache oder einem CDN liegen können.

Das ist das Schöne an einem Static Site Generator: Die fertige Seite ist nur noch HTML, CSS und ein bisschen JavaScript. Kein PHP, kein Ruby, kein Node.js auf dem Server. Einfach Dateien ausliefern.

## Fazit

Nach Jahren mit verschiedenen CMS-Systemen läuft der Blog jetzt auf meiner eigenen Software. Alles ist migriert, alle Funktionen sind da und das System macht genau das, was es soll - nicht mehr und nicht weniger.

Der Code ist Open Source und auf [GitHub](https://github.com/9it-full-service/quilldrop) [^verfügbar]. Wer einen schnellen, minimalistischen Blog ohne Overhead sucht, kann sich QuillDrop gerne anschauen.

> Write. Save. Published.

[^verfügbar]: Noch nicht, das kommt die Tage noch. Es gibt ein Update, so bald das Repo öffentlich ist.