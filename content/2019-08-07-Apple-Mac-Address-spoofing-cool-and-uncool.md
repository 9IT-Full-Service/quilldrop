---
title: "Apple Mac Address Spoofing cool und uncool"
date: 2019-08-07 23:35:47
update: 2019-08-07 23:35:47
author: ruediger
cover: "/images/posts/2019/08/07/vpn.webp"
tags:
    - Apple
    - AppleTV
    - TimeCapsule
    - Network
    - VLAN
    - ARP
    - TCP
    - IP
    - Spoofing
preview: "Kriege ich gerade aus der 1. Etage hier im Haus eine iMessage: Hast du eine Ahnung warum das Internet so langsam ist? Streaming über Lan stoppt nach einer halben Sekunde."
categories: 
    - Technik
toc: false
hide: false
type: post
---

# Alarm aus der 1. Etage

Kriege ich gerade aus der 1. Etage hier im Haus eine iMessage:

> Hast du eine Ahnung warum das Internet so langsam ist? Streaming über Lan stoppt nach einer halben Sekunde.

Kurz gegen gechecked und Youtube auf dem iPhone hat keine Probleme. Also selbst mal den AppleTV anschmeissen und da überprüfen. ARD Mediathek ok, Youtube OK.
FireTV mit Netflix und Co getesten. Selbst FireTV und DreamTV den Stream Haus intern vom Octagon kann ohne Probleme HD abrufen. Kein ruckeln, keine Abbrüche.
<!--more-->
# Kurzes debugging
Kurzer Blick auf dem RaspberryPI der DHCP gemacht und da viele DHCP Requests von dem AppleTV gesehen der als Problemgerät gemeldet wurde. Mutt auf dem RaspberryPI aufgemacht und nach ArpWatch Mails geschaut. Treffer, sehr viele Mails mit "flip flop" Meldungen passend zum Gerät.

```
60114 N + Aug 07 Arpwatch pi01.9 (  10) flip flop (appletv-01-2.intern.pretzlaff.co) eth0
60115 N + Aug 07 Arpwatch pi01.9 (  10) flip flop (appletv-01-2.intern.pretzlaff.co) eth0
60116   + Aug 07 Arpwatch pi01.9 (  10) flip flop (appletv-01-2.intern.pretzlaff.co) eth0
```

Dann pingen wir mal das Gerät über die IP an:

```
ping 192.168.178.108
PING 192.168.178.108 (192.168.178.108) 56(84) bytes of data.
From 192.168.176.106: icmp_seq=1 Redirect Host(New nexthop: 192.168.178.108)
From 192.168.176.106: icmp_seq=2 Redirect Host(New nexthop: 192.168.178.108)
```

# Erste Analyse

Sieht komisch aus und mir fällt auch prompt etwas ein was die Ursache sein kann. Mac-Address Spoofing der Apple TimeCapsule Geräten, wenn sie andere Apple Geräte offline gehen sehen oder meinen die Geräte sind offline.

Eurer AppleTV ist ausgeschaltet und trotzdem zeigt ein iPhone, iPad oder Mac das Gerät trotzdem sofort an wenn Ihr auf AirPlay geht? Genau das ist gemeint.

Apple Geräte wie die TimeCapsule oder AppleTV gucken welche anderen Geräte noch im Netzwerk sind und geht eines der Geräte offline, zum Beispiel in den Standby, melden sie sich im Netzwerk ab und ein anderes Gerät wie die TimeCapsule nimmt sich auf sein Netzwerkinterface die Mac-Adresse des Gerätes welches in Standby gegangen ist.

Dadurch sieht das Gerät immer noch als online im Netzwerk aus. Das gleiche wird auch so beim Mac / Macbook gemacht.

Fragen jetzt andere Geräte im Netzwerk eines der Geräte im Standby an antwortet die TimeCapsule auch brav und sendet im Hintergrund ein Datenpaket an das Standby Gerät, welches dadurch aufgeweckt wird.

Praktisch wenn man Back to my Mac, VPN oder einfach lokal mal auf einen Rechner zugreifen will. Oder wenn man AirPlay machen möchte. Die Geräte wachen wie von selbst auf.

# Kann aber auch mal doof sein

Zurück zum Problem. Checken wir mal den Arp Cache:

```
arp -a| grep 78:ca:39:ff:ee:3c
...
appletv-01-2.intern.pretzlaff.co (192.168.178.108) auf 78:ca:39:ff:ee:3c [ether] auf eth0
appletv-0-1.intern.pretzlaff.co (192.168.177.3) auf 78:ca:39:ff:ee:3c [ether] auf eth0
tc02.intern.pretzlaff.co (192.168.176.106) auf 78:ca:39:ff:ee:3c [ether] auf eth0
...
```
Das ganze noch mal mit -n um schneller und übersichtlicher die Ergebnisse zu sehen. Denn im Netzwerk sind 92 Geräte zu dem Zeitpunkt aktiv gewesen und ohne -n dauert das schon etwas länger.

```
arp -a -n | grep 78:ca:39:ff:ee:3c
? (192.168.176.106) auf 78:ca:39:ff:ee:3c [ether] auf eth0
? (192.168.178.108) auf 78:ca:39:ff:ee:3c [ether] auf eth0
? (192.168.177.3) auf 78:ca:39:ff:ee:3c [ether] auf eth0
```

Zm Vergleich die Mac Adressen die eigentlich hinter IP `192.168.178.108` und `192.168.177.3` stecken:

```
arp -a -n | grep 192.168.178.108
? (192.168.178.108) auf 34:c0:59:31:c7:92 [ether] auf eth0
arp -a | grep 192.168.178.108
appletv-01-2.intern.pretzlaff.co (192.168.178.108) auf 34:c0:59:31:c7:92 [ether] auf eth0

arp -a -n | grep 192.168.177.3
? (192.168.177.3) auf 34:c0:59:31:fe:46 [ether] auf eth0
arp -a | grep 192.168.177.3
appletv-0-1.intern.pretzlaff.co (192.168.177.3) auf 34:c0:59:31:fe:46 [ether] auf eth0
```

Man sieht hier das die IPs bei anderen Geräten benutzt werden. Aber eben genau die Geräte haben aber jetzt die Mac Adresse `78:ca:39:ff:ee:3c`.

# Ursache ist eigentlich eine Super Funktion

Aber hier macht sie gerade ein Problem. Apple Geräte wie die TimeCapsule bekommen mit wenn sich andere Apple Geräte im Netz befinden. Geht jetzt der betroffene AppleTV in den Standby nimmt sich die TimeCapsule die Mac Adresse vom AppleTV und simuliert die Funktionen des AppleTV. Für andere Clients wirkt es als wäre der AppleTV online. AirPlay zeigt den AppleTV als verfügbar an und man kann diesen auswählen.

Verbindet sich ein Client jetzt mit dem AppleTV wird er per WakeOnLan Magic Paket aufgeweckt und die TimeCapsule entfernt die Mac Addresse wieder von seinem Netzwerk Interface. Der AppleTV übernimmt sie wieder und bekommt seine IP.

> The sleep proxy service responds to address resolution protocol requests on behalf of the low-power-mode device

> When a sleep proxy sees an IPv4 ARP or IPv6 ND Request for one of the sleeping device's addresses, it answers on behalf of the sleeping device, without waking it up, giving its own MAC address as the current (temporary) owner of that address.

Wir haben hier mehrere TimeCapsule und 4 aktive AppleTV. Der 5te ist aktuell nicht in Benutzung. Das Verhalten war aber jetzt auch nur von der TimeCapsule bekannt. Die beiden AppleTV die hier aber auch die Mac Adresse übernommen hatten waren AppleTV der neuesten Generation.

Es scheint also so zu sein das jetzt auch der AppleTV als Bonjour Sleep Proxy eingesetzt wird. Das würde auch die sporadischen komischen Netzwerkprobleme erklären die immer mal wieder im Haus aufkommen.

TimeCapsule und AppleTV kümmern sich also drum das iTunes Sharing, File Sharing, Druckerfreigaben auch dann erreichbar sind, bzw. die entsprechenden Geräte auch im Standby angesprochen werden können.

Und ja, da Apple sich immer mehr aus dem Routermarkt verabschiedet ist die Funktion in den AppleTV bzw. TVOS gewandert.

> Mit der Funktion „Ruhezustand bei Bedarf beenden“ (auf Ihrem Mac) und mit Bonjour Sleep Proxy (durch ein AirPort-Gerät oder Apple TV bereitgestellt) können Sie Energie sparen und Kosten senken, während gleichzeitig der Zugriff auf alle freigegebenen Dienste sichergestellt bleibt. Zudem können Sie über Zugang zu meinem Mac auch aus der Ferne über das Internet auf die freigegebenen Dienste zugreifen. Die Funktion „Ruhezustand bei Bedarf beenden“ wird zusammen mit Bonjour Sleep Proxy auf Ihrer AirPort-Basisstation, Time Capsule oder Apple TV (wenn sich keine AirPort-Basisstation oder Time Capsule im Netzwerk befindet) ausgeführt. Hinweis: Apple TV agiert auch im Ruhemodus als Bonjour Sleep Proxy.

# Apple hat wohl nicht an die Masse an Geräte gedacht

Wie erwähnt haben wir mehrere AppleTV, TimeCapsule und eine grössere Menge MacOS und iOS Geräten im Hause. Da kommt bei 3 Haushalten schon was zusammen.
Jede Wohnung bzw. Etage hat hier einen eigenen Switch und AccessPoint. Per Vlans getrennt. Bis auf das Zeug was Media angeht. Die AppleTV und TimeCapsule Geräte sind in einem Netz zusammen.

Jetzt geht der AppleTV von meinem Bruder offline und die TimeCapsules meinen jetzt beide sie müssten im Netz behaupten sie seien der AppleTV. Die anderen AppleTV behaupten das gleiche. Der AppleTV wird dann wieder aufgeweckt und dann fängt die ARP schlacht im Netzwerk an.

Eines der Geräte hat das Datenpacket für den AppleTV bekommen, sagt dem Apple TV bescheid und lässt die Mac Adresse wieder frei, was dann aber für eines der anderen TimeCapsules oder AppleTV als "Da geht ein Device offline ich mach mal Bonjour Sleep Proxy", fragt beim DHCP nach IP und schreit das dann auch so ins Netz damit alle bescheid wissen. Der eigentliche AppleTV ist irritiert und meldet sich im Netz "Ey bin doch online, ich brauche eine IP."

Jetzt fängt das gleiche von vorne an. Mac Adresse wird vom AppleTV aktiviert und ins Netz geblasen und von einem anderen Device freigegeben und das nächste meint da geht was offline und macht einen auf Bonjour Sleep Proxy. usw. usw.

# Ok, im Urlaub auch Media-Netze mit Vlans trennen

Wir haben hier zuhause 4 Etagen. EG und 1. OG sind jeweils 1 Wohnung und wir hier oben haben 2. und 3. Etage. Zwischen den einzelnen Wohungen sind immer 2 Verbindungen, die als LAG geschaltet sind. Also 1 <-> 2 <-> 3 und noch mal 1 <-> 3 Etagen. Im schlimmsten Fall könnte man also 4 Gigabit Bandbreite benutzen. Alles mit Cisco Switchen. Sollte das mal nicht mehr ausreichen gehen wir einfach über Glasfaser und könnten auf 10 Gigabit oder mit LAG auf 20 Gigabit.

Die AccessPoints haben mehrere SSID konfiguriert, die aber auf allen Etagen. So das man hier von einer Wohnung in die nächste gehen kann, Verbindungen bestehen bleiben und man überall sein WiFi Netz hat. Zusätzlich noch Media WiFi und Gäste WiFi.

Alles schön getrennt mit VLANs. Bis auf das Media-Netz. Da wird dann ab nächster Woche dran geschraubt und auch da alles schön getrennt. Dann kommen auch die beiden Ubiquiti Edge Router X zum Einsatz.

Ich werde auch mal im Detail den Aufbau des Netzes hier beschreiben. Das steht schon länger auf er ToDo und wird auch mal Zeit.
