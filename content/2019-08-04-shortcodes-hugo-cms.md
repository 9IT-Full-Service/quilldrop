---
title: "Eigene Shortcodes Hugo CMS"
date: 2019-08-04 01:00:00
update: 2019-08-04 01:00:00
author: ruediger
cover: "/images/posts/2019/08/04/programming.webp"
tags:
    - Webpage
    - CMS
    - Hugo
    - Generator
    - Pipeline
    - Deploymend
    - Automatisierung
preview: "Um Bilder in Seiten oder einem Blogartikel schnell einfügen zu können habe ich mit im Hugo CMS einen eigenen Shortcut erstellt."
categories: 
    - Technik
toc: false
hide: false
type: post
---

### Shortcode erstellen

Um Bilder in Seiten oder einem Blogartikel schnell einfügen zu können habe ich mit im Hugo CMS einen eigenen Shortcut erstellt.

Ich habe immer das originale Bild und ein kleiner gerechnetes Bild. Diese werden
im Ordner `/static/img/posts/` gespeichert.
<!--more-->
Datei im Ordner: `/layouts/shortcodes/postimage.html`

```
<div>
<a href="{{ $.Site.BaseURL}}/img/posts/{{ index .Params 2 }}"><img src="{{ $.Site.BaseURL}}/img/posts/{{ index .Params 1 }}" width="800"></a>
<p>{{ index .Params 0 }}</p>
</div>
```

In Seiten/Artikeln kann ich so jetzt einfach mitfolgendem Code einfach Bilder einbinden:

```
{{ < postimage "title" "image-original.webp" "Image-thumbnail.webp" >}}
```

Die Bilder liegen aktuell noch alle in `/static/img/posts/`. Das werde ich aber noch ändern und auch den Shortcode anpassen. Die Originale bleiben in `/static/img/posts/` aber die Thumbnails werden aber in `/static/img/posts/thumbs` landen. Denn ich möchte die Bilder auch noch automatisch generieren lassen ohne großen Aufwand.

Aktuell wird jedes Bild einzeln verkleinert.
```
nconvert -resize 800 -o DB-Wifi-1-800.png DB-Wifi-1.png
```

Das will ich gerne in folgendes ändern:
```
for FILE in $(find ${IMAGEDIR} -type f -maxdepth 1)
do
  nconvert -resize 800 -o ${IMAGEDIR}/thumbs/${FILE} ${FILE}
done
```

Für das Image Rezise habe ich jetzt ein Docker Image erstellt.

Dockerfile

```
FROM alpine:edge
MAINTAINER "Rüdiger Küpper <ruediger@kuepper.nrw>"
RUN apk update && apk add imagemagick bash
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
```

Entrypint.sh

```
#!/bin/bash
cd /posts
for FILE in $(find . ! -name "*.svg" -type f -maxdepth 1); do convert -resize 800 ${FILE} thumbs/${FILE}; done
```

Im Ordner `/static/img/imagesresize.sh`:

```
#!/bin/bash
docker run -it -v $(pwd)/posts:/posts imageresize
```

Ein `./imagereszie.sh` im Ordner `/static/img/` generiert jetzt alle Bilder als Thumbnails in fester Breite neu und legt sie im Ordner `thumbs` ab.

Den Shortcode habe ich jetzt noch wie folgt angepasst:

```
<div>
<a href="{{ $.Site.BaseURL}}/img/posts/{{ index .Params 1 }}"><img src="{{ $.Site.BaseURL}}/img/posts/thumbs/{{ index .Params 1 }}" width="800"></a>
<p>{{ index .Params 0 }}</p>
</div>
```

### Shortcode für Soundcloud

Um meine oder von anderen Soundcloud Tracks einzubinden habe ich mir auch einen Shortcode geschrieben.

Datei im Ordner: `/layouts/shortcodes/soundcloud.html`
```
<p>{{ index .Params 1 }}</p>
<p>
<iframe width="736" height="400" scrolling="no" frameborder="no" src="https://w.soundcloud.com/player/?visual=true&#038;url=https%3A%2F%2Fapi.soundcloud.com%2Ftracks%2F{{ index .Params 0 }}&#038;show_artwork=true&#038;maxwidth=736&#038;maxheight=1000"></iframe>
</p>
```

Eingebunden wird das ganze dann in Artikel oder Seiten mit:

```
{{ < soundcloud 235962771 "Rock Solo Guitar - New Version" >}}
```

<!-- 
Das Ergebnis sieht dann so aus:

{{< soundcloud 235962771 "Rock Solo Guitar - New Version" >}}
-->


### Shortcode für Soundcloud Alben
```
<p>{{ index .Params 1 }}</p>
<p>
<iframe width="100%" height="300" scrolling="no" frameborder="no" allow="autoplay" src="https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/playlists/{{ index .Params 0 }}&color=%23ff5500&auto_play=false&hide_related=false&show_comments=true&show_user=true&show_reposts=false&show_teaser=true&visual=true"></iframe>
</p>
```

Eingebunden wird das ganze dann in Artikel oder Seiten mit:
```
{{ < soundcloudalbum 298199575 "clockopera - Veen Album 2017" >}}
```

<!-- 
Das Ergebnis sieht so aus:

{{< soundcloudalbum 298199575 "clockopera - Veen Album 2017" >}}
-->
