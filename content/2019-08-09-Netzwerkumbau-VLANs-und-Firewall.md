---
title: "Netzwerkumbau VLANs und Firewall"
date: 2019-08-09 22:52:27
update: 2019-08-09 22:52:27
author: ruediger
cover: "/images/posts/2019/08/09/network.webp"
tags:
    - Netzwerk
    - Internet
    - TCP
    - IP
    - VLAN
    - Switch
    - Router
    - Routing
    - Security
preview: "Oder auch erst einmal nicht. Ich wusste vom letzten Test die IP vom Ubiquiti EdgeRouterX nicht mehr. Dann halt mal wieder IPv6 zur Hilfe nehmen"
categories: 
    - Internet
toc: false
hide: false
type: post
---

# Dann wollen wir mal

Oder auch erst einmal nicht. Ich wusste vom letzten Test die IP vom Ubiquiti EdgeRouterX nicht mehr.
Dann halt mal wieder IPv6 zur Hilfe nehmen :smirk:

Ich habe ja eine Liste alle bekannten Adressen in dem Segment, also einfach `ping6` auf das interface `en0` und `ff02::1` und wir sollten die IPv6 Adresse von dem Teil haben.
<!--more-->
```
pi01:~# ping6 -I en0 ff02::1
16 bytes from fe80::7a8a:20ff:fe07:7a94%en0, icmp_seq=17 hlim=64 time=6.148 ms
16 bytes from fe80::7a8a:20ff:fe07:7a94%en0, icmp_seq=18 hlim=64 time=11.023 ms
16 bytes from fe80::7a8a:20ff:fe07:7a94%en0, icmp_seq=19 hlim=64 time=3.578 ms
```
Bingo, wir haben die IPv6 Adresse und so auch die Mac Adresse von dem Teil. Jetzt noch die ARP Table checken und wir sollten die IP haben.

```
pi01:~# arp -a -n | grep "78:8a:20:07:7a:94"
pi01:~# arp -a -n | grep "78:8a"
? (192.168.179.204) auf 78:8a:20:07:7a:95 [ether] auf eth0
```

Okay die Mac-Adresse passte dann nicht so ganz, aber da IPv6 von einem anderen Interface kommt, als nachher die IP-Adresse, einfach mit weniger String der Mac Adresse noch mal ARP fragen und zack da ist sie.

Also Netzwerkkabel von dem Teil raus und den Ping starten, um zu checken ob es wirklich der Ubiquiti EdgeRouterX ist.

```
pi01:~# ping 192.168.179.204
PING 192.168.179.204 (192.168.179.204) 56(84) bytes of data.
From 192.168.176.2 icmp_seq=1 Destination Host Unreachable
From 192.168.176.2 icmp_seq=2 Destination Host Unreachable
...
64 bytes from 192.168.179.204: icmp_req=19 ttl=64 time=1992 ms
64 bytes from 192.168.179.204: icmp_req=20 ttl=64 time=991 ms
64 bytes from 192.168.179.204: icmp_req=21 ttl=64 time=0.858 ms
```
Alles klar, der Abend war gerettet und nach dem einloggen erst mal checken was ich da vor ein paar Monaten verbrochen habe.
Mit der Konfig noch mal etwas ausprobiert und dann noch mal von vorne.

# Dann wird jetzt konfiguriert

Mit der vorhandenen Konfiguration weiter gemacht und auf den Interfaces
`eth1 (VLAN2)`, `eth2 (3)`, `eth3 (4)` und `eth4 (5)` die VLANs konfiguriert.
IP-Adressen jeweils aus einem Netz drauf gelegt und in jedem VLAN einen DHCP Server spendiert.

VLAN 2 auf dem Cisco Switch auf `gi8` und `gi10` getagged. Auf `gi10` hängt ein Accesspoint der das VLAN braucht, da er eine eigene SSID zum testen für die neuen Netze hat.

Das iPhone ins neue WiFi geworfen. Es gekommt eine IP wie gewünscht. Nur war trotzdem alles irgendwie hakelig.
Routing, Firewall usw. alles gechecked und ausprobiert. So will man das nicht haben.
Da ich mir eh gerade, durch ein blöden Konfigurationsfehler, die IP weg gezogen hatte und ich resetten musste, konnte ich auch gleich noch mal was anders konfigurieren.

Anstatt auf dem EdgeRouterX auf den Interfaces, jetzt auf dem EdgeRouterX Switch die VLANs eingerichtet. Um es kurz zu machen das wollte so gar nicht. Trotz der ganzen Anleitungen im Netz und genauen Beschreibungen das andere es so einsetzen. Es wollte einfach nicht.

Also noch einmal von vorne wie am Anfang schon. Nur dieses mal nach einem Reset des Routers. Ich hatte eh schon das Gefühl das ich da mal was drauf gemacht hatte was jetzt störte.

VLANs und VLAN-Interfaces anlegen, IPs setzen, Routing usw. und jetzt klappt es. An eth0 ist der Uplink zum Switch, hinter dem dann auch irgendwo das DSL Modem hängt. Switche, DSL und noch 1-2 andere Sachen werden in dem alten Netz bleiben. Das wird jetzt das Management Netz. Alle anderen VLANs bekommen jetzt eigene IP-Bereiche, inklusive DHCP Server.

Zwischen den Netzen Inter-VLAN-Routing und Firewall macht alles dicht bis auf das was durch soll. Gerade der Punkt war wichtig, da wir bei uns den einzigen Drucker im Haus haben den alle benutzen. Der kann jetzt auch weiter freigegegen werden. Genau wie unsere Synology. In die andere Richtung müssen wir auch mal auf andere Geräte in den anderen Wohnungen zugreifen. Daher müssen einzelne Geräte freigeschaltet werden können.

# So sieht die Konfiguration bis jetzt aus

Auf dem Ubiquiti sind die Interfaces jetzt eingerichtet.

{{< postimage "Ubiquiti Interfaces" "2019-08-10-Netzwerkumbau-VLANs-und-Firewall-1.webp" "2019-08-10-Netzwerkumbau-VLANs-und-Firewall-1.webp" >}}

![Ubiquiti Interfaces](/images/posts/2019-08-10-Netzwerkumbau-VLANs-und-Firewall-1.webp)


Auf dem AccessPoint ist eine Test-SSID im Vlan '2' und auf dem Port am Switch hat der AccessPoint auch das VLAN getagged bekommen.

Der Sitch an dem das ganze Zeug häng ist ein Cisco SG300-10 10-Port Gigabit Managed Switch. Der hat jetzt erst einmal 4 neue VLANs bekommen.

Auf dem Switch sieht die Config jetzt also für die VLANs so aus:

```
sw-02-1#show vlan tag 2
Created by: D-Default, S-Static, G-GVRP, R-Radius Assigned VLAN, V-Voice VLAN

Vlan       Name           Tagged Ports      UnTagged Ports      Created by
---- ----------------- ------------------ ------------------ ----------------
 2       VLAN2.OG           gi8,gi10             gi2                S

sw-02-1#show vlan tag 3
Created by: D-Default, S-Static, G-GVRP, R-Radius Assigned VLAN, V-Voice VLAN

Vlan       Name           Tagged Ports      UnTagged Ports      Created by
---- ----------------- ------------------ ------------------ ----------------
 3       VLAN1.OG           gi4,gi10                                S

sw-02-1#show vlan tag 4
Created by: D-Default, S-Static, G-GVRP, R-Radius Assigned VLAN, V-Voice VLAN

Vlan       Name           Tagged Ports      UnTagged Ports      Created by
---- ----------------- ------------------ ------------------ ----------------
 4        VLANEG                                                    S

sw-02-1#show vlan tag 5
Created by: D-Default, S-Static, G-GVRP, R-Radius Assigned VLAN, V-Voice VLAN

Vlan       Name           Tagged Ports      UnTagged Ports      Created by
---- ----------------- ------------------ ------------------ ----------------
 5        Server                                                    S
```

Hauptsächlich ist `VLAN 2` jetzt konfiguriert. Ist zwar auch noch nicht komplett fertig, aber dafür muss ich mir nachher erst einmal genau angucken welcher Port wo verkabelt ist und wie sich die beiden 16 Port Switche verhalten. Die beiden sind keine Cisco Switche sondern TP-Link Managed Switche. Ich weiß nicht wieso managed, weil naja die verrichten ihren Dienst, aber so viel können die auch nicht und das Webinterface ist einfach nur mies. Hätte ich das vorher gewusst wären das auch Cisco Switche geworden.

Sollten die mich nachher ärgern fliegen die raus und werden heute noch gegen Cisco Switche ersetzt. Dann heisst es dann wenigstens auch dabei: "On, Plug , Config und fertig."

Die TP-Link-Switche könnten noch ein Problem werden. Ich müsste eigentlich Tagged zu denen rübber und die VLANs dann UnTagged auf die Ports geben. Das war letzten schon nicht so erfolgreich.

# Die anderen Etagen sind dann auch noch dran

Das wird sehr schnell gehen. VLANs an den Uplink mit anlegen und dann alle Ports die dort für die die Etage sind mit dem VLAN Tag versehen. Auf dem AccessPoint das VLAN setzen und einfach nur alle Geräte kurz vor die Tür setzen damit sie wieder neu reinkommen. Danach sind sie in einer frisch renovierten Wohnung ... ähm ... Netzwerkumgebung.

Dann das gleiche noch auf der letzten Etage und dann ist das Thema durch.

# Das war ein Urlaubsprojekt für die nächsten 3 Wochen

Tja, erster Abend und eigentlich schon fertig. Der Rest ist jetzt nur noch stumpfes Port konfigurieren.

Aber es gibt ja noch genug für die 3 Wochen. Zum Beispiel da wir 2 DSL Anschlüsse haben auch jetzt wieder das Wechseln nach Last oder bei Ausfällen auf den anderen DSL Anschuss. Zutun gibt es immer etwas :laughing:.
