---
title: "Caddyserver als nginx Alternative"
date: 2019-08-29 08:15:00
update: 2019-08-29 08:15:00
author: ruediger
cover: "/images/posts/2019/08/29/caddyserver.webp"
tags:
    - caddy
    - webserver
    - nginx
    - alternative
preview: "Gerade für Testseiten bietet sich der kleine Webserver an. Er kann lokal einfach gestartet werden und man kann seine Seiten testen."
categories: 
    - Internet
toc: false
hide: false
type: post
---

Ich benutze sehr gerne den Nginx Webserver und das wird auch so bleiben.
Aber für keine Projekte habe ich jetzt einen Webserver gefunden der einfach und schnell benutzt werden kann.

<!--more-->

Die rede ist vom [Caddyserver](https://caddyserver.com).

Das nette ist:

* Konfigfile mit 3 Zeilen und läuft
* Läuft auf allen Plattformen
* Macht auf wunsch auch Letsencrypt SSL Certs On-The-Fly

Gerade für Testseiten bietet sich der kleine Webserver an. Er kann lokal einfach gestartet werden und man kann seine Seiten testen.

Auf der [Downloadseite](https://caddyserver.com/download) kann der Caddyserver für alle Plattformen heruntergeladen werden.
Dabei kann man auch PlugIns für sehr viele DNS Anbieter hinzufügen, um DNS Einträge für das generieren der SSL Certs zu erstellen.
Die Liste enthält alle möglichen Anbieter wie Cloudflare, Route53 (AWS), Azure, DYN usw.
In der PlugIn Liste sind sehr viele Plugins z.B. Proxy, HTTP-Auth, Geo-IP, IP-Filter und viele andere mehr.

Für Docker gibt es ein fertiges [Image](https://hub.docker.com/r/abiosoft/caddy/).

Das configfile ist wie gesagt sehr simple:

```
test1.homepage.net
browse
tls off
```

Möchte man PHP-FPM benutzen reichen 2 weitere Zeilen:

```
fastcgi / 127.0.0.1:9000 php # php variant
on startup php-fpm7 # php variant only
```

Die zweite Zeile kümmert sich sogar gleich darum das php-fpm7 vor dem Start des Webserver mit gestartet wird.
Hier kann man beliebige so genannte `RUN COMMANDS` ausführen um Dienste oder Befehle auszuführen die vor dem Start benötigt werden.

Ich werde den Webserver in nächster Zeit noch weiter testen. Hauptsächlich erst einmal für Testumgebungen. Mal sehen wie er sich so macht.
