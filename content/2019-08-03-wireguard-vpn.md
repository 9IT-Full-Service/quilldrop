---
title: "Wireguard VPN"
date: 2019-08-03 14:53:26
update: 2019-08-03 14:53:26
author: ruediger
cover: "/images/posts/2019/08/03/network.webp"
tags:
    - Netzwerk
    - IP
    - IPSec
    - VPN
    - Tunnel
    - Routing
    - Internet
preview: "Seit Jahren benutze ich OpenVPN, das wird auch so bleiben. Gerade für mobile Geräte ist aber auch Racoon für IPSec im Einsatz. IPSec aus einem Grund: Authentifizieren mit Zertifikaten."
categories: 
    - Internet
toc: false
hide: false
type: post
---

### WireGuard VPN

Seit Jahren benutze ich OpenVPN, das wird auch so bleiben. Gerade für mobile Geräte ist aber auch Racoon für IPSec im Einsatz. IPSec aus einem Grund: Authentifizieren mit Zertifikaten.

Denn OnDemand Verbindungen zu bestimmten Zielen im internen Netz bekommt man auf dem iPhone nur per Profilen hin wenn man Zertifikate benutzt.

Also habe ich meine eigene CA mit der ich die Benutzer mit den entsprechenden Zertifikaten verwalten kann. Funktioniert auch sehr gut, ist aber auch etwas aufwändiger.
<!--more-->
Da ich privat und auch mit Arbeitskollegen Server, Rechner und VPN Verbindungen betreibe um Daten auszutauschen ist der Aufwand mit IPSec immer etwas mehr. Daher wird da meistens OpenVPN genommen.

Hat aber wieder das Problem das OnDemand und manche Client Konfiguration komplizierter ist.

### WireGuard - VPN in einfach

Seit längerem schwirrt im Netz WireGuard herum. Ich hatte mir das auch schon vor ein paar Monaten angeguckt. Damals zum testen auf einem RaspberryPI und auch als Docker Container. Die Tests waren nicht so erfolgreich.

Da sich ein Kollege jetzt auch mit WireGuard beschäftigt hat kam das Thema auch bei mir wieder auf. Und siehe da! Es hat sich einiges getan und die Entwicklung hat große Schritte gemacht.

Der Server läuft und die Konfiguration ist simple. Und das trifft auch auf die für die Clients zu.

Auf beiden Seiten wird ein Schlüsselpaar generiert. Jetzt werden einige sagen: aber das kann man in der Familie gar nicht benutzen. Das kapiert doch keiner.

Doch kann man. Ich hatte vorher auch die Zertifikate, ob OpenVPN oder IPSEC, generiert. Das Schlüsselpaar wird auch einfach von mir generiert und damit die Konfiguration erstellt.

### Konfiguration in 3 Minuten

Die Konfiguration besteht immer aus 2 Teilen. Einmal das Interface und dem Peer, dem Client.
Beim Server sind es halt mehrere Peers.

Für den Anfang erst einmal nur Server und einen Client.
Dann hat man jeweils 3-5 Zeilen und das war es.

Jetzt hat man eine simple Datei mit ca 10 Zeilen und kann sie an den Benutzer verteilen.

Jetzt kann man die Datei client.conf schicken. Oder man schickt die Datei einfach einmal durch qrencode und erhält einen QRCode.

### Client installiert und QRCode scannen

Der Benutzer installiert sich jetzt einfach nur noch die WireGuard App. Nach dem öffnen Tunnel hinzufügen und einmal auf QRCode scannen tippen.

Den QRCode einscannen und schon ist alles fertig. Der Client kann sich sofort verbinden und bei Bedarf auch einstellen das in bestimmten Wifi Netzen und/oder mobil immer das VPN aufgebaut werden soll.

### Schritt für schritt Konfiguration

Eine Anleitung zu WireGuard wird es die Tage noch geben. Dann wird es auch Config Beispiele geben.

Auch wie das mit dem Routing in das Netz zuhause oder zu anderen geht.
Wir benutzen in unserem kleinem Mesh Netz für jeden Teilnehmer eigene Netze und Routen diese mit Hilfe von BGP.

Viele Sachen davon waren immer kompliziert. Mit WireGuard ist vieles aber einfacher geworden und es kann jeder jetzt Schell und einfach sichere Verbindungen aufbauen.
