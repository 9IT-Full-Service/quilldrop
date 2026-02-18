---
title: 'Wer hat das ICE WiFi kaputt gemacht?'
date: 2016-01-16 14:00:23
update: 2016-01-16 14:00:23
author: ruediger
cover: "/images/cat/internet.webp"
tags:
    - Internet
    - WiFi
preview: "Die Deutsche Bahn bietet auf mittlerweilen vielen Strecken WLAN an. In der 1. Klasse kostenlos, in der 2. Klasse kostenpflichtig."
categories: 
    - Internet
toc: false
hide: false
type: post
---


Die Deutsche Bahn bietet auf mittlerweilen vielen Strecken WLAN an. In der 1. Klasse kostenlos, in der 2. Klasse kostenpflichtig. Wenn es funktioniert ist alles gut. Aber leider ist das nicht immer der Fall. Ein Störfaktor ist aber öfters zu erkennen. Andere WLAN Accespoints. Entweder hat jemand sein Smartphone asl Accespoint für einen Laptop oder neuerdings benutzen andere reisende im ICE immer mehr einen UMTS Accesspoint. 

<!--more-->

Anderes Problem sind die Telekom eigenen Accesspoints. Das kann jeder selbst checken wenn der ICE in den nächsten Bahnhof fährt.  Das Problem ist der Country Code. In den ICE Zügen ist dieser in den Accesspoints auf den Ländercode für Frankreich (FR) gestellt. Für die Telekom Hotspots in den Bahnhöfen auf Deutschland (DE).  Nur was passiert in solchen Momenten. Alle Accesspoints arbeiten mit 2,4 GHz, da sollte es ja kein Problem geben. Doch 811.2 hat die Eigenschaft das man auch die Frequenzen der jeweiligen Länder damit beachten kann. In einigen gibt es z.B. den Kanal 13, in anderen ist dieser nicht vorhanden. Weil der Frequenzbereich anderweitig benutzt wird.

![Wlan Konfikt Country Code](/images/posts/DB-Wifi-1.webp)

![Länderkennung FR](/images/posts/DB-Wifi-2.webp)

![Länderkennung FR](/images/posts/DB-Wifi-3.webp)

```
Dec 31 15:46:40 rprMacBookAir kernel[0]: en0: 802.11d country code set to 'FR'.
Dec 31 15:46:40 rprMacBookAir kernel[0]: en0: Supported channels 1 2 3 4 5 6
7 8 9 10 11 12 13 36 40 44 48 52 56 60 64 100 104 108 112 116 120 124 128 132
136 140 149 153 157 161 165
Dec 31 15:46:40 rprMacBookAir locationd[86]: NETWORK: no response from server,
reachability, 2, queryRetries, 5
```
