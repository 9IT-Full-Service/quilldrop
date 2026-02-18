---
title: "Netzwerk Umbaulabor"
date: 2019-08-09 07:32:59
update: 2019-08-09 07:32:59
author: ruediger
cover: "/images/posts/2019/08/09/network.webp"
tags:
    - Internet
    - Netzwerk
    - Router
    - VLAN
    - IP
preview: "Oder auch erst einmal nicht. Ich wusste vom letzten Test die IP vom Ubiquiti EdgeRouterX nicht mehr. Dann halt mal wieder IPv6 zur Hilfe nehmen"
categories: 
    - Internet
toc: false
hide: false
type: post
---

# Testlabor für den Netzwerkumbau

In den nächsten 3 Wochen ist Urlaub angesagt. Das heisst jetzt für die Family: Es wird umfangreiche Änderungen am Netzwerk geben.

Das ist schon länger geplant, aber die Zeit ist immer so eine Sache. Das Netzwerk hatte ich ja schon angesprochen. 4 Etagen, 3 Wohnungen, 5 Switche, 3 WiFi AccessPoints, jede Menge Smarthome, smarte Geräte und Zeug drum her rum.
<!--more-->
# Ubiquiti Edge Router X als Testlabor

Ich habe mit einem der beiden Router schon ein paar Sachen gemacht. Jetzt wird auch der wieder zum testen genommen. Ein zweiter liegt auch schon etwas länger bereit zum Einsatz daneben.

Die beiden werden jetzt dafür genommen sie als Gateways zum Internet und für die Interne Netztrennung einzusetzen.

Im ersten Schritt kommen beide erst einmal an den Switch ins vorhandene Netz. Auf der anderen Seite kommen dann die ersten Geräte zum testen. Getrennt vom Rest und damit wird dann alles so vorbereitet wie es nachher sein soll.

Funktioniert das wird ein zweites VLAN erstellt und so das nächste Etagen Netz getestet. Kommt das ins Netz und kann man bei Bedarf einzelne Geräte mit anderen im anderen Netz verbinden wird noch das 3. Netz genau so vorbereitet.

# Edge Router vom Labor hinten nach ganz vorne stellen.

Wenn die Punkte oben alle rund laufen werden die beiden bzw. einer davon vom Testlabor nach vorne hinter die Fritzbox gepackt. Ok, das passiert nicht per Kabel, es wird eher an den Switchports anders konfiguriert.

Ziel ist das alles vorzubereiten und dann an einem Tag einfach per Schalter umzustellen. Es soll keiner im Haus davon etwas bemerken. Die IPs werden sich ändern, aber das im Hintergrund.

Die Geräte per Kabel bekommen an allen Switchen kurz ein Port Down und Up. Die sollten das alle mitbekommen und das Interface neu konfigurieren. Sie sind dann in einem anderen VLAN und bekommen die neuen IPs.

Geräte per WiFi werden ganze einfach neu bespasst. Alle AccessPoints werden einfach kurz resettet, Clients werden neu verbunden und danach auch mit anderer Netzwerkkonfiguration wieder online kommen.

# So der Plan

Wir werden sehen wir schnell das alles klappt und wann alles fertig ist.
Erst mal im kleinen testen und dann weiter machen.

Ich werde dann mal die nächsten Tage und Wochen berichten.
