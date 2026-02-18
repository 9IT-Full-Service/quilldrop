---
title: "VLANS und Firewall"
date: 2019-08-12 17:35:27
update: 2019-08-12 17:35:27
author: ruediger
cover: "/images/posts/2019/08/12/vpn.webp"
tags:
    - Netzwerk
    - IP
    - VPN
    - VLAN
    - IPSec
    - OpenVPN
preview: "Wenn man das komplette Netzwerk endlich mal umkrämpelt und in den Zustand bringt den man schon lange haben wollte sind manche Fails beim Einrichten auch etwas gutes um zu sehen das alles funktioniert."
categories: 
    - Technik
toc: false
hide: false
type: post
---

Wenn man das komplette Netzwerk endlich mal umkrämpelt und in den Zustand bringt den man schon lange haben wollte sind manche Fails beim Einrichten auch etwas gutes um zu sehen das alles funktioniert.

Damit auch eine Kommunikation zwischen den Netzen funktioniert muss zwischen den VLANS geroutet werden. Damit dann aber nicht alles mit jedem telefonieren kann braucht es dann auch eine(n) ~~Firewall~~ Paketfilter.

<!--more-->
Gestern sind hier alle Smarthome Geräte in ein neues VLAN umgezogen. Beim einrichten wollte direkt zu Beginn der Tradfri Gateway nicht mit dem neuen Netz funktionieren.

WiFi A soll VLAN 10 sein. WiFi B soll Vlan 20 sein. Wenn man aber die VLANs vertauscht und dem AP die WiFi SSID mit dem falschem VLAN konfiguriert klappt das auch nicht nicht so wie man es will.

Als die 1. Etage umgezogen ist wurden auf dem Cisco Switch dort alle Ports mit dem VLAN versorgt und den UpLink das VLAN getagged. Wollte auch erst einmal nicht.
Man sollte auch auf dem anderen Switch auf dem UpLink Port zur 1. Etage das VLAN mit auf den Port schmeissen. Und zack, die ersten Geräte tauchen auch schon auf.

Oh ein Firmware Update verfügbar für den Switch da unten, also das auch noch direkt mit drauf auf den Switch. Dann muss der ja eh einen Restart machen und das könnte für die Geräte die noch nichts mitbekommen haben reichen um die Netzwerkports in den Geräte dazu zu veranlassen sich einmal auf down zu setzen und wenn die Switch Ports wieder hoch kommen das Netzwerk wieder up.

3 Etagen umgezogen und sauber getrennt. Fehlt nur noch eine. Die darf aber vielleicht noch 1-2 Tage warten.

In der Zwischenzeit wird die Firewall noch etwas feinjustiert. Wir haben 5 Amazon Echo/EchoDot, von denen 4 ohne Probleme funktioniert haben. Nur der Echo in der Küche. Der ist einer der ersten Generation und irgend etwas können die neuen, was das alte Schätzcken nicht kann. Denn der Echo konnte heute morgen Streamen, Radio Essen und EinsLive liefen ohne Probleme.

Linda wollte aber vorhin bei der Hausarbeit in der Küche Hörspiele hören. iPhone auf und mit dem Echo verbinden und es passierte nichts. Erst dachte ich, ok das iPhone ist ja noch im alten Netz, ziehste sie eben um. QRCode gescanned fertig.
30 Sekunden später steht sie wieder da: `Geht immer noch nicht.`

Alle anderen Echo Devices ohne Probleme. Amazon Music App auf, Hörspiel an und Echo auswählen klappt. Nur der in der Küche nicht.

Der Echo wurde gerade etwas länger bespasst:

* Es wurde alles mögliche gemacht.
* Er wurde ins alte Media WiFi zurück gepackt.
* Das iPhone hinterher.
* Echo zwischen den WiFi Netzen hin und her.
* iPhone per Bluetooth mit dem Echo gekoppelt.
* iPhone wieder getrennt.
* Alle Geräte gelöscht.
* Echo komplett zurückgesetzt.
* Echo lange vom Strom.

Es wurde wirklich alles ausprobiert. Das komische war es gab bei einigen der Aktionen oben aus der Liste Situationen bei denen es z.B. bei mir kurz klappte und wenn wir Lindas iPhone verbunden haben ging wieder nichts. Zwischendurch klappte auch mal kurz wieder an ihrem iPhone. Dann wieder nicht.

Die ganze Situation war sehr strange, da alles was man Alexa gefragt hat, wie Wetter, Verkehr usw. klappt. Also nur alles wo Alexa Text aus dem Netz abfragt und als Sprache aus gibt.

"Alexa, spiele Musik.". Kurze Zeit später sagt Alexa auch dabei sie kann die Musik nicht abspielen. Also Online ist sie, aber kann nicht alles machen.

Noch mal alles gechecked, AccessPoint, Switche, Ports, VLANs, IPs, Gateway, aber alles hat gepasst. Und mit anderen Echos klappt es ja auch. Also noch mal Firewall, tcpdump und Log auf dem EdgeRouter checken. Die Default Policy hatte Log an und hat das Logfile bei der Masse an Daten schon dicht gemacht. Da war erst einmal nichts mehr zu sehen, da die Platte voll war. Wenn vorher eine Rule etwas blockt konnte man das in dem Moment nicht sehen. Ein `echo "" > /var/log/messages` und noch mal getestet.

`[2OG-50-A]IN=eth0 OUT=eth1.2 MAC=78:8a:20:07:7a:94:38:10:d5:76:31:46:08:00 SRC=192.168.176.2 DST=10.0.2.115 LEN=540 TOS=0x00 PREC=0x00 TTL=62 ID=23432 PROTO=UDP SPT=53 DPT=53080 LEN=520`

Ok, DNS kommt nicht durch. Das wird dann mal freigeschaltet. Weiter Logfile geguckt und noch weiter Ports und Protokolle freigegeben. Ist zwar aufwändig bis alles läuft, aber ich will das so haben. Das die anderen Echos das Problem nicht hatten werde ich mir die Tage noch genauer angucken.

Alles wird geblockt, alles freigegeben was man braucht und der Rest der dann noch fehlt fällt auf wenn es nicht klappt.

Es gibt Leute die meinen alles offen und dann nur das verbieten was man nicht haben will. Bei vielen Sachen weiß man aber nicht welche Protokolle und Ports es benutzt. Das kann man nicht sicher machen, weil man nicht alles kennt was irgendwann mal durchs Netz schwirrt.

Die vertauschten VLANS am AccessPoint und die fehlenden Firewallfreischaltungen haben gezeigt: `Alles richtig gemacht und genau so sollte es werden`.
Das fühlt sich so einfach einfach richtig an.
