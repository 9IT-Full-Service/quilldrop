---
title: "Nginx Google Fonts Filtern"
date: 2022-09-25 10:00:00
lastmod: 2022-09-25 10:00:00
author: ruediger
cover: "/images/cat/technik.webp"
tags:
  - Nginx
  - GoogleFonts
  - Internet
preview: Wir ersetzen einfach die Domain für die Google Fonts 
categories: 
  - Technik
type: post
ShowToc: false
hide: false
draft: false
---


### Abmahnwelle wegen Google Fonts?

Seit ein paar Tagen finden einige wieder Briefe im Briefkasten wegen einer Abmahnwelle zu Google Fonts. Dieses mal lässt sich sogar ein Anwalt vor den Karren spannen und er versendet aktuell wohl tausende Abmahnungen.

Laut Post einer Kanzlei sind dort alleine am Freitag vorletzter Woche 300 dieser Schreiben eingetroffen die man bearbeitet.

Wer noch Google Fonts direkt von den Google Server einbunden hat, sollte diese lokal speichern und direkt selbst ausliefern.

### Was machen wenn Theme Updates Fonts wieder einbinden?

Es gibt z.B. Wordpress Themes bei denen man die Fonts zwar anpassen kann, aber nach einem Update des Themes wird wieder Google Fonts von den Servern bei Google geladen.

Anderes Problem sind Personen die nicht wissen was sie machen müssen.
Wer, egal wieso, die Nutzung von Google Fonts von Google Servern unterbinden möchte und selbst einen Nginx Webserver einsetzt, der kann das sehr leicht für alle Seiten vom Nginx ereldigen lassen.

In der Config vom Nginx einfach folgende Zeile mit eintragen.


Wir ersetzen einfach die Domain für die Google Fonts gegen:
`https://fonts.googleapis.com`.


```
sub_filter 'https://fonts.googleapis.com' 'https://disable-google-fonts';
```

Die Zeile könnte man jetzt auch gut dafür nutzen um alle Seiten zu finden die noch Google Fonts einbinden.

```
sub_filter 'https://fonts.googleapis.com' '/googlefonts/';
```

Damit wird der Font nicht auf dem Sevrer gefunden der die Seite ausliefert. Dafür kann man dann aber einfach in den Logfiles alles Seiten nach `/googlefonts/` suchen und hat dann schnell alle Seiten und Stellen wo man noch anpassen müsste.

```
grep '/googlefonts/' /path/to/logfiles/*.log
```

Sollten Einträge vorhanden sein weiß man schnell wo noch Google Fonts eingebunden sind.
