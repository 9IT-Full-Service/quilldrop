---
title: 'Neues, selbst gebautes Blog- und CMS'
date: 2024-02-08 18:00:00
author: ruediger
cover: "/images/posts/2024/02/cms.webp"
tags: [Homepage, CMS, Blog]
categories: 
    - Internet
preview: "Mit meinem neuen System können einzelne Bereiche oder Seiten auf separaten Instanzen betrieben werden, die nur bei Bedarf skaliert werden. Dies reduziert die Belastung für den Rest der Homepage, und selbst bei starker Last bleibt alles andere erreichbar."
draft: false
top: false
type: post
hide: true
toc: false
---

![Blog and CMS](/images/posts/2024/02/cms.webp)


Jekyll, Octopress, Hugo, InkProject ... Die Liste der CMS- und Static-Page-Generatoren neben Typo3, Joomla und WordPress, um nur einige zu nennen, ist bereits lang. Noch länger war der Plan, das eigene Projekt fertigzustellen. Es war sogar vorteilhaft, dass es länger dauerte als geplant, da mehrere Anläufe nötig waren. Die Erfahrungen mit den Static-Page-Generatoren waren jedoch sehr wertvoll, denn alle haben ihre Stärken und Schwächen.

Mit InkProject begann die Trennung von Software, Inhalten, Theme und der fertigen Seite. Dies führte zum aktuellen Ansatz, der noch weiter verfeinert wurde.

Auch der fertige Inhalt kann getrennt werden. Bei allen CMS oder Static-Page-Generatoren wird das gesamte System üblicherweise auf einem Server abgelegt, wobei eine horizontale Skalierung möglich ist. Dies betrifft dann jedoch alle Seiten. Wenn ein Bereich, wie der Blog, ein Ressort/Themenbereich oder eine einzelne Seite, stark belastet wird, musste bisher das gesamte System skaliert werden.


    cms_cli page --all
    cms_cli server --build --name default 
    cms_cli server --rollout --name default
 
    cms_cli page -s -f 2024-01-04-Warp-Editor
    cms_cli server --build --name technik
    cms_cli server --rollout --name technik


Mit meinem neuen System können einzelne Bereiche oder Seiten auf separaten Instanzen betrieben werden, die nur bei Bedarf skaliert werden. Dies reduziert die Belastung für den Rest der Homepage, und selbst bei starker Last bleibt alles andere erreichbar.

Inhaltlich wird alles im HTML-Ordner abgelegt, ähnlich wie bei anderen Systemen. Es besteht jedoch die Möglichkeit, Inhalte an zusätzlichen Orten zu speichern. Standardmäßig wird das Docker-Image unter dem Namen „default“ erstellt, neben zwei weiteren Images für CSS und Fonts, die unabhängig voneinander sind.

Wenn eine Seite oder ein Blog mehrere Bereiche hat, wie Allgemeines, Technik, Developer, Smart Home usw., könnte der Bereich Smart Home ausgelagert werden. Dies geschieht über die Konfiguration für den Inhalt.

Anschließend wird das entsprechende Docker-Image gebaut und ausgerollt.

Alle Schritte, vom Generieren aller Seiten, über das Erstellen der Docker-Images bis hin zum Rollout aller Container, können über die CLI erfolgen. Beiträge und Seiten können auch einzeln generiert werden, ohne dass sie auf der Webseite verlinkt werden. Sie werden erst durch das Generieren aller Seiten oder weiterer Seiten, wie Übersichtsseiten, zugänglich.

Durch diese Trennung können für einzelne Bereiche unterschiedliche Themes verwendet werden, was eine optische Unterscheidung ermöglicht. Bei verschiedenen Verantwortlichen für unterschiedliche Bereiche ermöglicht die Trennung ein paralleles Arbeiten und Veröffentlichen. Publishing und Deployment erfolgen nur für die jeweiligen Bereiche.

Das System unterstützt auch fortgeschrittene Szenarien, wie die Nutzung eines Gateways oder CDNs, und kann die Pfade für die Bereiche auf unterschiedliche Instanzen lenken, die sogar in einem Multi-Cloud-Setup bei AWS, Azure, GCP, OTC usw. laufen.

Der einzige fehlende Teil ist ein Kommentarsystem. Ein eigenes Kommentarsystem zu schreiben, wäre technisch machbar, aber die Integration in statische Seiten und die Bekämpfung von Spam erfordern zusätzlichen Aufwand.

Discus hatte ich schon mal früher im Einsatz. Ähnliche Systeme gibt es auch von anderen Anbietern. OpenSource und/oder Selfhosted Lösungen gibt es auch. Davon hatte ich mir über die Jahre auch immer mal wieder ein paar angeguckt. Leider waren alle nicht immer zufriedenstellend. Einige waren sehr aufwändig beim Einrichten und Betrieb. Manche hatten im Betrieb Probleme gemacht. Sei es nur wenn es nur um das Thema Updates ging. Es war einfach zu nervig andauernd eingreifen zu müssen. 

Klar man könnte jetzt Discus oder ein anderes System nehmen. Problem gelöst, um den Betrieb kümmert sich jemand anderes. Aber die Daten liegen bei jemanden anderen und nicht bei einem selbst. 

Also habe ich nach einiger Zeit mal wieder Tante Google befragt und ein wenig umgeguckt. Dabei bin ich auf Commento gestoßen und es sah auf dem ersten Blick schon sehr gut aus. Commento kann über die Seite von Commento benutzt werden, oder auch selfhosted auf dem eigenen Server installiert werden. 

Ich habe Commento für ARM64 kompiliert und in ein ARM64 Docker-Image integriert, inklusive PostgreSQL, Konfiguration und PersistentVolume im Kubernetes. Ein Helm-Chart und alles für das ARM-Image werde ich noch veröffentlichen, ebenso wie Commento, damit es direkt angeboten werden kann.

Wer den Footer dieses Blogs betrachtet, wird noch eine Weile das „Powered by InkProject“ sehen. Zuerst wird eine andere Seite auf mein System umgezogen, da dies dringender benötigt wird. Aber auch hier wird es nach und nach voran gehen, denn ich kann ja jetzt einzelne Bereiche nach und nach umziehen. 