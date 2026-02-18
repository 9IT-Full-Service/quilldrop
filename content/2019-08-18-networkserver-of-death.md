---
title: "Networkserver of Death"
date: 2019-08-18 22:30:00
update: 2019-08-18 22:30:00
author: ruediger
cover: "/images/posts/2019/08/18/serverofdeath.webp"
tags:
    - Storage
    - NAS
    - Router
    - UPnP
    - AVM
    - Fritzbox
preview: "Server immer wieder von extern nicht mehr erreichbar"
categories: 
    - Technik
toc: false
hide: false
type: post
---

# Server immer wieder von extern nicht mehr erreichbar

Für das Projekt [Essenz - Rock Dein Block](https://www.beone-projects.com/projekte/essenz/ "Rock Dein Block Projekseite") haben wir hier auf einem Synology NAS mehrere Dienste laufen damit sich die Coaches und Teilnehmer austauschen können.

Es gab immer mal wieder Probleme mit dem Zugriffen auf das NAS. Und das Monitoring (Icinga) hat auch immer wieder alamiert. Nach dem umfangreichem Netzwerkumbau in den letzten Tagen hatte ich gestern Abend teilweise die Änderungen am Netz als Ursache nicht mehr ausschliessen können.
<!--more-->
Vorteil der Änderungen war aber ,durch die strickte Trennung mit VLANs und den Routern zwischen den Netzen, kann man jetzt gezielt nach Fehlern suchen. Es ist nicht mehr alles eine große Suppe IP Traffic von über 90 Geräten.

Das ganze könnte aber auch die Ursache des ganzen Übels sein. Daher habe ich heute die komplette Netzwerkkonfiguration noch mal auf links gedreht.

Auch weil, immer wenn das Problem aufgetaucht ist, der Traffik zwischen den Netzen und Switchen auf einmal massiv angestiegen ist. So stark das ich das eigentiche Problem da noch gar nicht auf dem Schirm hatte. Dazu war der Traffic auf den Ports viel zu hoch, als das man auf UPnP kommen könnte.

Der Traffik blieb selbst wenn dann das Portforwarding komplett ausgestiegen ist und sogar aus den Routern auf einmal verschwunden ist. Manuell gesetzt, Problem taucht auf, Portforwarding nicht mehr in der Konfiguration.

Und selbst dann blieb der Traffic intern weiter sehr hoch. Da geht man von einem Problem in der Netzwerkkonfiguration oder an den Geräten im Netz aus.

# Das NAS ist im Netzwerk noch mal umgezogen

Um einige Sachen ausschliessen zu können ist das NAS noch mal umgezogen. Also es blieb da stehen wo es ist, aber es wurden an den Switchen ein paar Ports umkonfiguriert, so das es so gesehen wieder direkt hinter den Routern angeordnet war. Das Problem war ... immer noch vorhanden.

Um das Problem an einem der Switche auszuschliessen wurde ein frischer Ubiquiti EdgeRouterX SPF aus dem Karton genommen und das NAS da angeschlossen. Da zum Switch, an dem das NAS eigentlich hängt, 4 Leitungen als LAG aus der anderen Etage kommen um 4 GigaBit/s hochschieben zu können, einfach eine davon gezogen und auch an den EdgeRouterX geklemmt. Die andere Seite unten vom Switch auch gezogen und direkt auf den Router zum Internet geklemmt.

> Internetrouter -> EdgeRouterX -> NAS

Portforwarding konfiguriert und getestet. Online war das NAS wieder auch wieder ohne Probleme. So bald wieder Requests rein kommen ist das Portforwarding weg und auf dem EdgeRouterX sieht man massig Traffik. Also alles andere im Netzwerk kann jetzt schon einmal ausgeschlossen werden.

Aber woher kommt es und was ist die Ursache?

# DNS gändert und Ruhe

Nach langem Suchen nach der Ursache und nicht wirklich der Ursache näher kommend noch mal überlegt was nicht ausgeschlossen wurde. Das war zu dem Zeitpunkt noch der Zugriff von außen auf das NAS. Die Domain dafür wird bestimmt schon bei einigen <s>BadBoys</s> SkriptKiddies rumgeistern. Also Attacken kann man nicht ausschliessen.

Also DNS Record löschen und auf `127.0.0.1` setzen.

```
dnsmngt -d -z domain.de -s subdomain
dnsmngt -a -z domain.de -s subdomain -i 127.0.0.1
```

Danach war auch erst einmal alles ruhig und es gab auch keine Probleme mehr im Netzwerk.

Das ganze wieder zurück und den DNS Record auf den alten Wert setzen. Kurze Zeit später war das Problem wieder da. Portforwarding zeigt den Server dahinter wieder als Offline an. Netzwerktraffik ging wieder hoch. Es besteht weiter.

Jetzt eben eine neue Subdomain anlegen und damit testen.

```
dnsmngt -a -z domain.de -s testsubdomain -i 111.222.333.444
```

Portforwarding wieder konfiguriert. Aufgerufen und es ist alles damit online. ein wenig getestet und auch jetzt wieder alles weg. Traffik ... ihr wisst schon.
 Bber die neue Subdomain kann keiner kennen. Dazu ist sie noch zu frisch.

Es muss also noch etwas anderes sein. Um zu prüfen ob die Fritzbox vielleicht mal wieder eine kaputte Konfiguration hat musste sie ausgeschlossen werden. Weil die Box Resetten, komplett auf Werkseinstellung und neu einrichten wäre nicht so der Aufwand. Aber wieso wenn es nicht nötig ist.

# Weg über den 2. Anschluss testen

Wir haben ja zwei Anschlüsse und daher auch zwei Fritzboxen. Auf der anderen Fritzbox das Portforwarding eingerichtet. Andere Fritzbox, anderer Anschluss. Wenn es da auch so ist, dann ist es nicht die Fritzbox.

Damit das ganze überhaupt funktioniert musste das NAS einfach nur als Standard Gateway die andere Fritzbox IP bekommen.

Parallel wurden ein paar Geräte in anderen Etage dazu angehalten den anderen Anschluss zu benutzen. So wurden alle anderen Geräte nicht gestört und über die Leitung ging nur noch das NAS online.

Default Routen werden hier auf den Routern eh immer beide gesetzt. Aber so das immer das andere Gateway als Fallback dient wenn ein Anschluss gestört ist.
So kann man auch mal eine Fritzbox vom Netz nehmen, ohne das groß jemand etwas merkt.

NAS online, alles ok. Doch auch da dann wieder Probleme. Da sogar so heftig, das die Fritzbox sogar komplett jedesmal das DSL verloren hat. Das ganze konnte ich 5 mal reproduzieren.

# UPnP noch mal gechecked

In den Fritzboxen noch einmal überprüft ob das UPnP aus ist. Bei dem Gerät war es auch nicht an. Das wurde jetzt schon zig mal überprüft, aber trotzdem noch mal gucken kann ja nicht schaden.

Auf dem NAS sollte es auch nicht an sein, das wird nicht benutzt und soll auch nicht aktiv sein. Aber ein Blick in die Config zeigt: da ist etwas konfiguriert ist. Also noch mal checken ob das auch aktiv ist.

Und siehe da, der Mist ist da aktiv.

Der ganze Aufriss, auf den Cisco Switchen alles möglich checken und tracen, auf den anderen Switche, auf dem EdgeRoutern, auf den AccessPoints, auf allen möglichen Geräten, weil das komplette Netzwerk auf einmal Traffic machte, der weder von innen nach draussen, noch von aussen nach drinnen so massiv kommen konnte. Tcpdump, traceroutes, mtr usw. um das Problem zu finden. Und dann ist es dieses fucking UPnP.

Portforwarding aktiviert, NAS wieder online. Also Tcpdump noch mal gezielt an einigen Punkten im Netzwerk auf UPnP Zeug angesetzt.

```
tcpdump -i eth0 udp and port 1900 and dst 239.255.255.250 -s0 -w UPnP.pcap
```

Das Netzwerk war die ganze Zeit ruhig. Bis man die App am Smartphone aus dem mobilen Netz aufruft und Daten vom NAS abruft und kopiert. Einfach auf das iPhone kopieren oder selbst auf dem NAS an einen anderen Ort kopieren.

In dem Moment brach jedesmal die Multicast Hölle aus. Das NAS hat jedesmal die Router zugebommt mit Multicast Paketen bis die eine das Portforwarding eingestellt hatte. Oder bei der anderen ja sogar so weit das DSL komplett ausgestiegen ist. Der Router musste jedesmal DSL neu synchronisieren.

# WTF

UPnP für internen Kram wie Media Server, Media Printer, Netzlaufwerke usw. ist ja toll und schön. Das benutzen wir hier mit einigen Geräten auch. Der Octagon z.B. stellt Streaming, Mediaplayer für Fotos und Videos damit bereit.

Aber an den Router soll so ein Mist bitte nicht. Das wurde 2013 schon sehr gut gezeigt. Als man ca. 50 Millionen Geräte im Netz gefunden hatte, die darüber von aussen übernommen werden konnten.

Das, also UPnP, ist wieder so ein tolles Beispiel für: `Dem Benutzer alles abnehmen, weil es so für ihn einfacher ist.`

Ja einfach, aber macht es nur schlimmer. Die Leute stellen sich Geräte zuhause hin die dann mal eben so das komplette Leben (Daten auf den Geräten) frei ins Internet stellen. Weil irgend ein Gerät meint dem Router über UPnP zu sagen: `Mach mal Port auf und schick alle rein zu mir, ich habe die Daten.`

> Ja, kannste schon so machen, ist dann halt kacke!

Wieso versucht man den Leuten immer alles abzunehmen? Wer ein Geräte hinter seinem "sicheren" Internetanschluss freigeben möchte, der soll das gefälligst selbst einrichten. Und ja dazu gehört auch: `Beschäftige Dich mit dem Scheiss`. Man sollte zumindest ein wenig verstehen was man da macht. Vor allem was man da mit seinen Daten im schlimmsten Fall machen. Oder man macht sich wenigstens kurz Gedanken darüber ob man das auch wirklich so will und nicht besser noch andere Sicherheitsmassnahmen dazu schaltet.

# Aber auch die Hersteller könnten mal ...

1. Wieso ballert da eine Synology einfach so massiv Multicast ins Netzwerk und zwar so das man an alles mögliche denkt. Nur nicht an UPnP.

Leute, wir haben hier zwischen den Wohnungen jeweils 2x1 GB im LAG. Und wir hier in unserer Wohnung nach oben ins Büro 4x1 GB im LAG. Und der UPnP-Traffik war zwischen dem ganzen Streaming und Backups die in der Zeit liefen zu sehen. Also es war nicht der UPnP-Traffik selbst zusehen. Danach wurde ja noch nicht geziehlt geguckt. Der Traffik war halt so viel das er in den Graphen und Statistiken nur auffallen konnte.

Und wenn etwas, wie in dem Fall UPnP, nicht funktioniert, wo war da eine Mail oder eine Meldung in der Oberfläche? Jeder Mist wird einem zig mal benachrichtigt, aber das nicht? Aber dafür das Netz mal komplett zu müllen. Ausserdem kann man einen Service auch sagen er soll ruhig sein wenn etwas nicht klappt und nicht einfach noch lauter ins Netz brüllen lassen.
Meinem Sohn sage ich auch:

> Es bringt nichts wenn Du lauter wirst, dadurch höre ich Dir nicht besser zu wenn ich Dir nicht gerade zuhören kann. Warte kurz und sei bitte kurz leise.

Und ja, wenn er dann nicht hört bekommt auch mein Sohn ein `FIN-RST` von mir und darf sich erst einmal abkühlen gehen.

2. AVM macht ja eigentlich immer einen guten Job. Aber UPnP deaktiveren haben sie ja jetzt so gut versteckt und ist auch nicht mehr so gut zu erkennen ob es aktiviert ist oder nicht. Das könnte bitte besser sein.

Und leider kann man auch nie gut erkennen, was, wann, wie schief läuft oder gelaufen ist. Man kann auch nicht so gut auf den Dingern gucken was in so einer Situation gerade passiert. Da wären ein wenig mehr Infos schon nicht schlecht.

Ausserdem wäre eine Funktion in der Fritzbox nicht schlecht, die einfach mal sagt: `Device XY, du kommst hier gerade mit ein paar tausend UPnP Zeug rein, das ist hier nicht, geh sterben` und dann strickt wegblockt. Fertig. Dazu dann noch in der Oberfläche eine Meldung, das da gerade ein Device scheisse baut.

# UPnP wird jetzt weggeblockt

Nachdem jetzt alles wieder zurück gebaut wurde. Also eher zurück-zurück ;-). Wird jetzt die Firewall zwischen den VLANS und zu den Routern noch um ein UPnP DROP erweitert.

Das ein Gerät meint es darf sich hier alles erlauben soll ja nicht noch einmal vorkommen. Ausserdem will ich sehen wenn so etwas noch einmal passiert und dann auch benachrichtigt werden.

UPnP vom Synology aus sollte eigentlich ausgeschaltet gewesen sein. Ich kann mir auch nicht vorstellen das irgend ein Paket, was installiert wurde, das selbstständig gemacht hat. Das ist eigentlich nicht Synology-Style

Das jemand beim Einrichten eines Services gedacht hat man muss das in den Einstellungen auch noch unter "Routerkonfigurieren" aktivieren, will ich nicht ausschliessen.

Ist schon heftig, wenn so etwas solche Auswirkungen hat. Und wird so auch nicht mehr passieren. Dafür wird es schöne Rules in den Firewall geben.
