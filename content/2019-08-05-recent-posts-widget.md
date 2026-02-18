---
title: "Recent Posts Widget"
date: 2019-08-05 22:00:00
update: 2019-08-05 22:00:00
author: ruediger
cover: "/images/posts/2019/08/05/programming.webp"
tags:
    - Hugo
    - CMS
    - Sidebar
    - Widgets
preview: "In meinen Blogs hatte ich immer ein Widgets für die letzte Posts. Das habe ich gerade auch für das Hugo CMS erstellt."
categories: 
    - Technik
toc: false
hide: false
type: post
---

### Recent Posts Widget für die Sidebar

In meinen Blogs hatte ich immer ein Widgets für die letzte Posts. Das habe ich gerade auch für das Hugo CMS erstellt.

Im Theme unter Layouts -> Partials -> Widgets habe ich eine Datei `lastposts.html` erstellt:
<!--more-->

```
    <div class="panel-body">
      <ul class="nav nav-pills nav-stacked">
      {{ $count := .Site.Params.widgets.recent_posts }}
      {{ range first $count .Pages }}
        <li><a href="{{ .Permalink }}">{{ .Name }}</a></li>
      {{ end }}
      </ul>
    </div>
  </div>
```


Die Zeilen 1, 9-12 und 16 sind markiert. Das sind Zeilen die das Widget steuern bzw. die Anzahl an konfigurierten letzten Posts ausgeben.

* Zeile 1 überprüft ob das Widget aktiv ist.
* Zeile 9 setzt die Variable $count auf die konfiguriere Anzahl der posts
* Zeile 10 bis 12 durchläuft die letzte $count Posts und zeigt sie an.

Das funktioniert auf der Startseite gut, aber in den Artikeln bleibt die Liste leer.

Also wurde jetzt noch folgender Code ersetzt:
```
{{ $count := .Site.Params.widgets.recent_posts }}
{{ range first $count .Pages }}
  <li><a href="{{ .Permalink }}">{{ .Name }}</a></li>
{{ end }}
```

Durch:
```
{{ $pages := where .Site.RegularPages "Type" "in" .Site.Params.mainSections }}
{{ range first $count $pages }}
<li><a href="{{ .Permalink }}">{{ .Name }}</a></li>
{{ end }}
```

Jetzt wird das Widget auf allen Seiten befüllt.

### Konfiguration des Widget

Das Widget kann in der `config.toml` konfiguriert werden.

```
...
# Enable and disable widgets for the right sidebar
[params.widgets]
    categories = true
    tags = true
    search = true
    recent_posts = 10
...
```

Das Widget kann mit `false` deaktiviert werden.
Die Anzahl kann frei gewählt werden.
