---
title: 'Twitter und Youtube/Vimeo Videos datenschutzgerecht  eingebunden'
date: 2018-06-06 21:10:02
update: 2018-06-06 21:10:02
author: ruediger
cover: "/images/cat/technik.webp"
tags: [datenschutz, dsgvo, Facebook, Internet, Social Network, Technik, Twitter, vimeo, wordpress, youtube]
preview: "Hier werden jetzt Videos von Youtube, Vimeo (Facebook würde auch gehen), sowie Tweets von Twitter datenschutzgerecht eingebunden."
categories: 
    - Technik
toc: false
hide: false
type: post
---

Hier werden jetzt Videos von Youtube, Vimeo (Facebook würde auch gehen), sowie Tweets von Twitter datenschutzgerecht eingebunden.  
[Heise.de](https://www.heise.de/newsticker/meldung/Embetty-Social-Media-Inhalte-datenschutzgerecht-einbinden-4060362.html) hat dafür [Embetty](https://github.com/heiseonline/embetty) und [Embetty Server](https://github.com/heiseonline/embetty-server) veröffentlicht.

<!--more-->

Beim Aufruf der Artikel werden die Videovorschau jetzt nicht mehr direkt bei Google und Co mit der IP der Besucher dieser Seite der abgerufen.
Das wird jetzt alles über den Embetty Server gemacht. Embetty hängt also als Proxy zwischen euch und Youtube, Vimeo, Facebook und Twitter.
Erst wenn ein Video angeklickt wird werden Daten an die Server der Anbieter gesendet.
Im Wordpress Header wurde jetzt einfach der embetty Server hinzugefügt und im Webroot der Seite das embetty.js abgelegt:

```
<meta data-embetty-server="https://blog.pretzlaff.info:8089">
<script async src="/embetty.js"></script>
```

Tweet im Artikel einbinden:
```
<embetty-tweet status="1000738984253943811"></embetty-tweet>
```

Ergebnis:

```
<embetty-video type="vimeo" video-id="91085172"></embetty-video>
```

Ergebnis:

```
git clone https://github.com/heiseonline/embetty-server.git
cd embetty-server
```

docker-compose.yml

```
version: '3.1'
services:
  server:
    image: heiseonline/embetty-server:latest
    ports:
      - 8089:8080
    environment:
      - VALID_ORIGINS=http://localhost
      - TWITTER_ACCESS_TOKEN_KEY=<YOURTOKENKEY>
      - TWITTER_ACCESS_TOKEN_SECRET=<YOURSECTRETKEY>
      - TWITTER_CONSUMER_KEY=<YOURCUNSOMERKEY>
      - TWITTER_CONSUMER_SECRET=<YOURCONSUMERSECRET>
```

 Speichern ..

```docker-compose up -d``` und der Server rennt.
