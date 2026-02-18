---
title: 'Mein Blog Kommentarsystem Self-Hosted'
date: 2025-06-25 08:00:00
update: 2025-06-25 08:50:00
author: ruediger
cover: "/images/posts/2025/06/comments.webp"
featureImage: "/images/posts/2025/06/comments.webp"
tags: [Kommentare, Blog, Static, InkProjek, Jekyll, Hugo, Gatsby]
categories: 
    - Development
preview: "Externe Kommentarsysteme wie Disqus bringen durchaus Vorteile mit sich: Sie sind schnell eingerichtet, bieten umfangreiche Funktionen und kümmern sich um Spam-Schutz und Moderation. Dennoch haben sie entscheidende Nachteile ..."
draft: false
top: false
type: post
hide: false
toc: false
---

![Automatisierte Kubernetes Volume-Backups](/images/posts/2025/06/comments.webp)

# TLDR; 

[Github Repo](https://github.com/ruedigerp/comments)


# Ein eigenes Kommentarsystem für statische Blogs: Warum ich auf Go setze

Statische Website-Generatoren wie Hugo, Jekyll oder Gatsby erfreuen sich großer Beliebtheit - und das zu Recht. Sie bieten schnelle Ladezeiten, hohe Sicherheit und einfaches Hosting. Doch wenn es um interaktive Funktionen wie Kommentare geht, stoßen sie schnell an ihre Grenzen. Während WordPress-Nutzer auf bewährte Plugins zurückgreifen können, müssen Betreiber statischer Blogs oft auf externe Anbieter wie Disqus ausweichen.

## Das Problem mit externen Kommentaranbietern

Externe Kommentarsysteme wie Disqus bringen durchaus Vorteile mit sich: Sie sind schnell eingerichtet, bieten umfangreiche Funktionen und kümmern sich um Spam-Schutz und Moderation. Dennoch haben sie entscheidende Nachteile:

**Datenschutz und Kontrolle:** Externe Anbieter sammeln oft umfangreiche Nutzerdaten und zeigen Werbung an. Als Website-Betreiber hat man wenig Kontrolle über diese Aspekte und muss sich auf die Datenschutzrichtlinien Dritter verlassen.

**Abhängigkeiten:** Was passiert, wenn der Anbieter seinen Service einstellt oder die Preise drastisch erhöht? Alle Kommentare könnten verloren gehen.

<!-- **Performance:** Externe Skripte können die Ladezeit der Website negativ beeinflussen und zusätzliche HTTP-Requests verursachen. -->

**Design-Integration:** Oft lassen sich externe Kommentarsysteme nur begrenzt an das eigene Website-Design anpassen.

## Die Lösung: Ein eigenes Kommentarsystem in Go

Aus diesen Gründen habe ich mich entschieden, ein eigenes Kommentarsystem zu entwickeln. Die Wahl fiel auf Go als Backend-Sprache, da sie sich hervorragend für Web-Services eignet und sowohl performant als auch ressourcenschonend ist.

### Architektur und Design-Entscheidungen

Das System basiert auf einer klaren Trennung zwischen dem statischen Blog und der Kommentarfunktionalität:

**Separate Domain:** Das Kommentarsystem läuft unter einer eigenen Domain, getrennt vom Hauptblog. Diese Architektur bietet mehrere Vorteile: Bessere Skalierbarkeit, einfachere Wartung und die Möglichkeit, das System für mehrere Websites zu nutzen.

**Go-Backend:** Der Server wurde in Go implementiert und stellt eine REST-API zur Verfügung. Go eignet sich perfekt für solche Aufgaben - es ist schnell, hat eine hervorragende Standard-Bibliothek für Web-Services und benötigt nur minimale Systemressourcen.

**JavaScript-Frontend:** Die Integration in den Blog erfolgt über ein JavaScript-Widget, das die Kommentare dynamisch lädt und darstellt. Dies ermöglicht eine nahtlose Integration in jedes statische Website-System.

### Integration in das Blog-System

Die Einbindung könnte nicht einfacher sein. Im Blog-Theme wird im Header einfach das JavaScriipt vom Kommentar Server geladen und lediglich ein `<div>`-Container mit einer eindeutigen ID platziert:

```html
<html>
    <head>
        <script src="https://comments.example.com/static/js/comment-widget.js"></script>
    </head>
...

```html
# Date + Title
<div data-comment-post-id="2025-06-19-git-merge-script"></div>
# or Title
<div data-comment-post-id="git-merge-script"></div>
# or URL Path 
<div data-comment-post-id="posts/2025-06-19-git-merge-script.html"></div>
```

Mit Blog-ID bei der Nutzung auf mehreren Seite (siehe oben): 

```html
# Blogname + Date + Title
<div data-comment-post-id="blog-example-net/2025-06-19-git-merge-script"></div>
# or Blogname + Title
<div data-comment-post-id="blog-example-net/git-merge-script"></div>
# or Blogname + URL Path 
<div data-comment-post-id="blog-example-net/posts/2025-06-19-git-merge-script.html"></div>
```

oder: 

```html
# Blog-ID + Date + Title
<div data-comment-post-id="123/2025-06-19-git-merge-script"></div>
# or Blog-ID + Title
<div data-comment-post-id="123/git-merge-script"></div>
# or Blog-ID + URL Path 
<div data-comment-post-id="123/posts/2025-06-19-git-merge-script.html"></div>
```

Das JavaScript-Widget erkennt diesen Container automatisch und lädt die entsprechenden Kommentare. Dabei kann die Integration sowohl post-spezifisch als auch global im Theme erfolgen - je nach gewünschter Flexibilität.

Wenn Kommentare global in allen Artikeln aktivieren möchte, kann man auch einfach folgenden Code an der entsprechenden Stelle im Theme einfügen. 

Das `{{.Link}}` ist hier speziefisch für [InkPaper, a static blog generator](https://github.com/InkProject/ink) . Je nach eingesetztem CMS muss diese Variable angepasst werden. 

```html
<script>
    CommentWidget.init({{.Link}});
</script>
```

Oder auch hier wieder mit eine Blog-ID: 

```html
<script>
    CommentWidget.init(blog-example-net/{{.Link}});
</script>
```

Kommentar Form: 
![1](/images/posts/2025/06/comments-1.png)

### Technische Vorteile

**Performance:** Da das System speziell für die eigenen Anforderungen entwickelt wurde, ist es schlank und schnell. Keine unnötigen Features bedeuten weniger Code und bessere Performance.

**Datenschutz:** Alle Daten bleiben unter eigener Kontrolle. Es werden nur die notwendigen Informationen gespeichert, und die Einhaltung der DSGVO liegt in den eigenen Händen.

**Anpassbarkeit:** Das Design lässt sich vollständig an die Website anpassen. CSS-Styles können frei definiert werden, ohne auf die Vorgaben eines externen Anbieters angewiesen zu sein.

**Skalierbarkeit:** Go ist bekannt für seine hervorragende Concurrent-Performance. Das System kann problemlos viele gleichzeitige Anfragen verarbeiten.

## Funktionsumfang und Features

Das selbst entwickelte System muss nicht weniger können als kommerzielle Alternativen. Typische Features umfassen:

- **Moderation:** Administrative Oberfläche zur Verwaltung und Freischaltung von Kommentaren

Adminpanel: 
![2](/images/posts/2025/06/comments-2.png)

### ToDos: 

   * **Spam-Schutz:** Implementierung eigener Spam-Filter oder Integration bestehender Lösungen
   * **Antwort-Funktionen:** Verschachtelte Kommentare und Antworten auf bestehende Beiträge
   * **Benachrichtigungen:** E-Mail-Benachrichtigungen bei neuen Kommentaren
   * **Rate-Limiting:** Schutz vor Spam durch Begrenzung der Kommentar-Frequenz

## Deployment und Betrieb

Ein weiterer Vorteil von Go ist die einfache Deployment-Strategie. Go-Programme werden zu einzelnen, ausführbaren Dateien kompiliert, die keine zusätzlichen Dependencies benötigen. Das macht das Deployment auf jedem Server unkompliziert.

Das System kann auf einem einfachen VPS betrieben werden und benötigt nur minimale Ressourcen. Für die Datenspeicherung wir ValKey (Redis) eingesetzt. 

Im Github Repository kann jederzeit das aktuelle [Release](https://github.com/ruedigerp/comments/releases) herunterladen werden. Die Version wird auch im Helm Chart immer aktualisiert, das das ein `helm upgrade ...` ausreicht für ein update oder bei FluxCD einfach automatisch per ImageUpdateAutomation. 

## Dokumentation und Installation

Installation mit einem Binary, als Docker Container, Helm, FluxCD ist unter den folgenden Links beschrieben.
Genau so wie die API Dokumentation und hilfreiche Redis Befehle, falls man mal etwas debuggen möchte. 

   * Installation: [Install doc](https://github.com/ruedigerp/comments/blob/main/docs/README.md)
   * Docker: [Docker-compose](https://github.com/ruedigerp/comments/blob/main/docs/docker-compose/README.md)
   * Kuberenetes (Noch in Arbeit) 
   * Helm: [helm](https://github.com/ruedigerp/comments/blob/main/docs/helm/README.md)
   * FluxCD: [FluxCD Installation](https://github.com/ruedigerp/comments/blob/main/docs/fluxcd/)
   * API: [API Docs](https://github.com/ruedigerp/comments/blob/main/docs/api/README.md)
   * Redis: [Redis Commands](https://github.com/ruedigerp/comments/blob/main/docs/redis/README.md)

## Fazit

Die Entwicklung eines eigenen Kommentarsystems mag zunächst nach Mehraufwand aussehen, bietet aber langfristig entscheidende Vorteile. Vollständige Kontrolle über Daten und Funktionalität, bessere Performance und die Unabhängigkeit von externen Anbietern rechtfertigen den initialen Entwicklungsaufwand.

Go erweist sich dabei als ideale Wahl für das Backend - die Sprache ist nicht nur performant und ressourcenschonend, sondern auch gut zu lernen und zu warten. In Kombination mit einem flexiblen JavaScript-Frontend entsteht so ein System, das sowohl technisch überzeugt als auch den eigenen Anforderungen perfekt entspricht.

Für Betreiber statischer Blogs, die Wert auf Datenschutz, Performance und Kontrolle legen, ist ein eigenes Kommentarsystem eine durchaus lohnenswerte Alternative zu externen Lösungen.

