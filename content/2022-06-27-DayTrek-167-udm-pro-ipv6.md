---
title: "DrayTek 167 UDM Pro IPv6"
date: 2022-06-27 12:00:00
lastmod: 2022-06-27 12:00:00
author: ruediger
cover: "/images/cat/technik.webp"
tags:
  - Ubiquiti
  - Unfiy
  - UDM
  - DreamMachine
  - DrayTek
  - Vigor
  - 167
  - IPv6
preview: "Der etwas längere Weg zu IPv6 am Draytek Vigor 167 und der Ubiquiti DreamMachine Pro."
categories: 
  - Internet
toc: false
hide: false
draft: false
type: post
---


Der etwas längere Weg zu IPv6 am Draytek Vigor 167 und der Ubiquiti DreamMachine Pro.

![Ubiquiti DreamMachine Pro](/images/posts/unifiy-pro.webp)


[Der Chris](https://twitter.com/lelei) hatte mich letztes Jahr auf die neue Firmware hingewiesen, in der
IPv6 funktionieren soll. Bei ihm hatte es auch an seinem Anschluss auch sofort geklappt.

Hier funktionierte es aber überhaupt nicht. Selbst alles mögliche aus Foren hatte keinen Erfolg.
IPv6 Prefix setzen, WAN/Lan IPv6 Firewall regeln und vieles mehr. Selbst VLAN7 in der DreamMachine ausschalten und von dem Vigor übernehmen lassen, oder anders herum. Nichts hat funktioniert.

Immer wieder nach neuem Update der UDM Firmware habe ich dann getestet und probiert. Bis dann in einem Forum in einem Beitrag auf ein Firmware Update für das DrayTek Vigor 176 Modem gibt, welche das Problem beheben sollte.

Und siehe da, mit der neuen Firmware Version ist jetzt auch eine IPv6 Adresse verfügbar. Aber man sieht sie nicht im Webinterface der DreamMachine. Aber die Seite [ipv6-test.com](https://ipv6-test.com) zeigt IPv4 und IPv6 an.

IPv6-Test.com zeigt aber erst einmal einen Score mit 18 von 20 an. Weil ICMP blockiert wird.
Das kann behoben werden indem man Firewall Rules dafür anlegt.

    Type: Internet v6 Local
    IPv6 Protocol: ICMPv6
    IPv6 ICMP Type Name: Any
    Action: Accept

    Type: Internet v6 Local
    IPv6 Protocol: IPv6-ICMP
    Match all protocols except for this: enable

Zusätzlich kann man jetzt auch die IPv6 Prefixe bei den internen Netzen hinzufügen.

    IPv6 Interface Type: Prefix Delegation
    Router Advertisement (RA): Enable
    RA Priority: high
    DHCPv6 Range start: ::2
    DHCPv6 Range stop: ::7d1
    DHCPv6/RDNSS DNS Control: auto

Diese Einstellungen einfach bei allen internen Netzen hinzufügen und jedes Netz bekommt seinen eigenen IPv6 Prefix.

Damit haben alle Clients IPv6 Adressen bekommen. Der nächste Schritt war dann auch für IPv4 DHCP über die DreamMachine zu machen. Das wurde bis jetzt immer von einem der beiden Ubiquiti EgdeRouterX SPF erledigt.
Netzbereiche eingetragen aktiviert und die DHCP Services im EgdeRouterX deaktiviert. Und dann auch gleich die interfaces deaktiviert.

Jetzt macht die DreamMachine DHCP, IPv6 und alles was sie auch schon vorher gemacht hatte.

DreamMachine, ein Träumchen. ;-)
