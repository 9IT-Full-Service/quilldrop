---
title: "Blog Theme noch einmal angepasst"
date: 2019-08-20 21:27:00
update: 2019-08-20 21:27:00
author: ruediger
cover: "/images/posts/2019/08/12/vpn.webp"
tags:
    - Hugo
    - CMS
    - Theme
preview: "Auch heute wurde noch etwas weiter am Theme geschraubt. Es waren noch ein paar Baustellen was die Darstellung von Bildern betrifft."
categories: 
    - Technik
toc: false
hide: false
type: post
---


Auch heute wurde noch etwas weiter am Theme geschraubt. Es waren noch ein paar
Baustellen was die Darstellung von Bildern betrifft.
Jetzt passt erst einmal alles wie gewünscht. Und gerade wurde auch ein großer Test mit
Apple iPhones (4s bis XR), iPads (Alle von 1er bis Pro), Samsumg Phones und Tables (Alles mögliche), LG, Nexus (6,9), HTC One, Sony, Kindle usw.

<!--more-->

Eigentlich alles was man so draussen erwarten kann. Passt alles und die nächste Zeit jetzt nur noch Kleinigkeiten verbessern.

Da wären z.B. die Codeblöcke. Die gefallen mir noch nicht so gut, da ich aktuell keinen Hintergrund setzen kann. Code der auf der Seite angezeigt wird passt von der Breite zwar jetzt, aber die Box im Hintergrund geht nur ca. 60-70 der verfügbaren Breite.
Sieht dann halt doof aus wenn der Code dann drüber hinaus geht.

Bilder waren heute mittag noch mal nervig. Alles hat gepasst und ein Ende war in Sicht. Bis ich dann auch wieder alle anderen  Browser mit getestet habe. Und der Firefox hat alles komplett ~~beschis~~ zerschossen angezeigt.

Die Bilder liegen jetzt in einem CDN und können da auch schnell deployed werden. Die fliegen auch noch aus dem Hugo Ordner raus. Die kommen ja erst später in die Seite.
Das Theme und der restliche Content sind im Git auch schon getrennt, so das alles einzeln angefasst werden kann und erst in der Gitlab Pipeline wird alles zusammen gesteckt und die Seite generiert.

Für die Bilder habe ich jetzt ein kleines Script was die Bilder verkleinert. Da wird noch was verbessert. Da jetzt das Theme  fertig ist kann ich auch mal geziehlt mit den Grössen gucken. Die Bilder dann damit auch in die passende Grösse automatisch anpassen und ins CDN deployen, fertig.
