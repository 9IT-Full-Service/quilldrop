---
title: 'Static InkProject Blog Automation API'
date: 2024-01-11 16:00:00
author: ruediger
cover: "/images/posts/2024/01/static-blog.webp"
tags: [Blog, InkProjekt, API]
categories: 
    - Internet
preview: "" 
draft: true
top: false
type: post
hide: false
toc: false
---

<!-- [Englisch Version](/posts/2023-09-21-kubernetes-network-policies-en.html)
-->

![Static Weblog](/images/posts/2024/01/static-blog.webp)

Mein Blog ist jetzt seit ein past Jahren ein Static Html Blog. Benutzt hstte ich Jekyll, Octopress und etwas lönger Hugo. 
Seit ein paar Monaten benutze ich dafür InkProject. Mach der Migration der Blogposts, den ersten manuellen durchläufen und manuellem 'docker build' für dem nginx mit dem generierten HTML, ist das ganze dann auch recht schnell auf github gelandet. Mit ein paar Github Actions Zeilen wurde das Blog so scgon mal nach Änderungen und einem Push automatisch zusammen gebaut. 

Im nächsten Schritt wurde alles aufgeteilt. Das ink Binary und die Config selbst, der Content, die Images und das Theme sind in eigene Git Repository getrennt wurden. Gleichzeitig wurde der git Branch develop erstellt, so das es auch eine Testversion gibt. 

Im Content Repository wurde ein Workflow hinzugefügt, der bei commits automatisch das ink-Blog repository triggert und dort den Workflow aufruft um das Blog zu generieren.

Dabei wird eine neue Version mit semantic-Release getagged. Die Repository content, images und theme werden ausgechecked. Das Blog wird generiert und das fertig geberierte statische HTML wird in ein repository 'html' gepushed. 

Um dann alles zu deployen wird ein docker image erstellt und das vorher generierte HTML aus dem Repository ausgechecked und hinzugefügt. Das Docker Image in die Docker Registry gepushed und kann dann deployed werden. 

Für das Blog gibt es einen Helm Chart. Da alles automatisiert ist wird auch dieser im GitHub Actions Worklow mit erstellt. 
Also auch hier den Code auschecken, Helm package mit '--Version' und der vorher erstellten semantic-Release Version aufgerufen. Der fertige Helm Chart ins Chartmuseum gepushed und schon kann der nächste Step ein 'helm upgrade ...' ausführen. 

Die neue Blog Version ist online. Läuft auch sehr zuverlässig und stabil. Doch vom Push bis zur aktuellen Version im Browser dauert dann doch schon so 3-4 Minuten. 

Da musste noch etwas gehen. Und es ging noch etwas. Denn die Idee war dann ein Docker Image zu haben, welches durchgehend im Kubernetes löuft. Es hat alle nötigen Binary und alle Repository, die bei Bedarf aktualisiert werden. Das Blog wird generiert und auf jeweils für Prod und DEV eingehangene Mounts geschrieben. Diese werden dann für Kubernetes Deployments als ReadOnly Volumes genutzt. In den Deployments können dann so auch schnell weitere Pods skaliert werden. 

Jetzt möchte ich aber nicht jedesmal in den Pod im Kubernetes manuell ein Update anstoßen, vorallem mobil manchmal recht aufwändig am iPhone. Daher musste etwas her was das erledigt. 

Das Docker Image für das generieren hat einen kleinen API Service bekommen. Damit kann ich verschiedene URLs aufrufen, die dann verschiedene Sachen für das Blog in Prod oder DEV erledigen. 
Ok, die API macht das nicht selbst, die macht keine Logik oder für andere Commands aus. Sie schreibt Jobs als Textfiles in ein Verzeichnis und löscht sie dann auch gleich wieder. Mehr macht sie nicht. 

Denn ein 2. Service im ink Builder Docker Image ist auch wieder ein kleines, in go geschrienes Programm, was nur eine Aufgabe hat. Es überwacht ein Verzeichnis und so bald eine Datei für einen Job angelegt wurde, wird dieser Job ausgeführt. 
Das kann der Build vom Prod Blog oder für Dev sein. Aktualisieren der Images, das Theme oder andere Aufgaben. 

Ein Commit wie dieser Blog Post ins Git ist dann eine Aktualisierte Blog Version innerhalb von 20-30 Sekunden. 

Da ich eine andere Seite demnächst auch auf inkProject umstellen möchte, werde ich die API noch erweitern und kann dann alle Seiten damit erstellen lassen. 




