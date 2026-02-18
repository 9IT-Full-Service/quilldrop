---
title: "Fremde Geräte im Wifi"
date: 2022-07-04 10:00:00
lastmod: 2022-07-04 10:00:00
author: ruediger
cover: "/images/posts/2019/08/03/network.webp"
tags:
  - Router
  - Wifi
  - WLAN
preview: "Fremde Geräte im WLAN? Chip.de, Telekom.de, pcwelt.de und ihre Tipps beim Thema WLAN. Es ist zum Haare raufen. Bei manchen Artikeln fast man sich nur noch an den Kopf und will eine Aspirin einwerfen."
categories: 
  - Internet
toc: false
hide: false
type: post 
draft: false
---

Fremde Geräte im WLAN?

Chip.de, Telekom.de, pcwelt.de und ihre Tipps beim Thema WLAN. Es ist zum Haare raufen. Bei manchen Artikeln fast man sich nur noch an den Kopf und will eine Aspirin einwerfen.

Ja, nutzt jemand das WiFi unberechtigt kann es zu einer langsamen Internetverbindung führen. Oder im schlimmsten Fall zu rechtlichen Problemen kommen. Dann, wenn jemand illegales Zeug mit dem Internetanschluss macht.

Aber was bitte werden in den Artikeln für Tipps gegeben? Es fängt damit an, daß man vielleicht vermutet jemand ist im WiFi und dann wird anhand vom Router erklärt wie man so einen Benutzer dort findet. Es stimmt das sehr viele wahrscheinlich eine Fritzbox haben. Also wird es anhand der Weboberfläche der Fritzbox erklärt. Nur was machen bitte alle Kunden der Telekom, die einen Speedport Router haben. Oder Netgear, TP-Link usw.?

Also, hat man den Verdacht, es befindet sich ein unbefugter Benutzer in WiFi, dann ist die Lösung sehr kurz und Knapp: Änder Dein WiFi-Passwort. Punkt.

Damit wäre jeder Artikel in 2-3 Sätzen schon erledigt. Aber das wichtigste fehlt in allen dieser Artikel. Keiner geht drauf ein was man sonst noch machen könnte. Es wird eigentlich in allen immer nur geraten den MAC-Adressen Filter zu aktivieren.

Damit, mit den Mac-Filter, dann auch viel Spaß. Denn z.B. iOS/iPadOS (Apple)  und auch einige Andoid Geräte haben die Möglichkeit die MAC-Adressen zu verschleiern. Bei Apple heißt es in den WLAN Einstellungen „Private WLAN-Adresse“. Das ist eigentlich als Standard immer aktiv.

Was heißt das genau?

Jedes Gerät hat bei Bluetooth und Netzwerkschnittstellen (Kabel oder WiFi) immer eine weltweit eindeutige MAC-Adresse, über die es identifizierbar ist. Im Heimnetzwerk und in Firmen bekommt jedes Gerät meistens seine IP beim Verbinden von einem DHCP-Server zugewiesen. Dazu sendet das Gerät die Anfrage für eine IP ins Netzwerk. In dieser Anfrage ist die Mac-Adresse enthalten. Die IP-Adressen werden aus einem Pool genommen und entsprechend eine freie an das Gerät vergeben. Verlässt man das Netzwerk oder schaltet das Gerät aus, geht dann nach kurzer Zeit wieder in Netzwerk, bekommt man meistens die gleiche IP zugewiesen.
Oft kann man sich trotz DHCP in den meisten Routern auch immer die gleiche IP Adresse geben lassen, indem man dem Gerät im Router die IP fest zuweisen kann.

Dabei wird immer die MAC-Adresse herangezogen. Wir erinnern uns an die Anfrage mit der Mac-Adresse weiter oben.

Jetzt aktivieren wir aber an einem iPhone „Private WLAN-Adresse“ und erhalten damit jedes mal sogar eine andere MAC-Adresse. Entsprechend ist das Gerät auch noch für den Router und dem darauf laufenden DHCP-Server ein neues Gerät. Also bekommen wir jedesmal eine andere IP-Adresse im Netzwerk.

Und jetzt schlägt dann der „Pro-Tipp“ großer IT Bereiche mancher Portale zu.
Die Empfehlung den MAC-Adressenfilter zu aktivieren.
Da sitzt jetzt der nicht Internet-Pro an seinem iPhone und macht sich Sorgen um sein WiFi. Irgend etwas stimmt damit nicht. Kommt jetzt auf einen der Artikel und in ihm kommt etwas Angst auf beim lesen. „Es könnte jemand Dein Netzwerk benutzen“. MAC-Adressenfilter aktivieren, super wird gemacht. Verstanden wozu und was er damit macht, keine Ahnung. Wird schon richtig sein. Es funktioniert auch prima, da der Filter aktiviert werden kann und sein iPhone, so wie das Gerät der Holden auch schön in der Liste eingetragen sind.

Bis er dann sein Gerät einige Zeit nicht mehr an hatte. Es lag aus auf dem Tisch. iPhone entsperren und surf…..
Es bleibt offline, keine App kann sich verbinden, keine Internetseite funktioniert mehr. Die Frau kann aber noch alles machen, sie war die ganze Zeit am Telefon.
Würde man jetzt mit dem noch funktionierenden Gerät auf dem Router noch mal den MAC-Adressenfilter überprüfen, würde man sehen:
   * Sein Gerät ist offline
   * Ein neues Gerät versucht zu verbinden, wird aber geblockt. 
   * Das „neue“ Gerät ist aber sein Gerät.

Da „Private WLAN-Adresse“ aktiv ist, hat er jedesmal eine neue MAC-Adresse. Daher ist es für den Router und seinen Mac-Adressenfilter auch jedesmal ein neues Gerät.

Im Normalfall geht der unbedarfte Benutzer nicht hin und deaktiviert den Filter jetzt oder fügt das „neue“ Gerät im Filter hinzu. Am besten deaktiviert er „Private WLAN-Adresse“ im iPhone.

Nein, es klappt bei Ihr und nicht mehr bei ihm. Also scheint wohl irgend etwas am WLAN zu sein. Also: „ich starte mal den Router neu, vielleicht geht es ja dann wieder.“

Der Supergau tritt ein. Ab jetzt, also der Fall wie er heutzutage oft üblich ist, wenn es nur noch Smartphone und/oder Tablets im Haushalt gibt, kommt kein Gerät mehr in das WiFi rein.
Alle haben die Einstellung für die „Private WLAN-Adresse“ aktiviert. Der MAC-Adressenfilter ist aktiv und sagt jetzt zu jedem Gerät: „ich kenn Dich nicht, du kommst hier nicht rein.

Und jetzt wird es für die beiden spaßig. Die kommen in Ihr WiFi nicht mehr rein. Sie können jetzt auch nicht mehr den Filter deaktivieren. Sie müssen das mit einem Laptop machen der mit einem Kabel an den Router angeschlossen wird.
Und selbst wenn sie den haben wird das für die meisten ein schwieriges Unterfangen. Denn oft wissen sie nicht mal die IP von ihrem Router.

Und alles nur weil man Probleme mit dem WiFi hatte. Auf ein paar Seiten nach der Lösung geguckt hat und dabei dann die Panik wegen fremden Nutzern aufkam. Dann der Tipp mit dem Mac-Adressenfilter, der heutzutage nicht nur unwillkommene Benutzer aussperrt, sondern direkt alle. Auch die, die rein dürfen, besser gesagt bis dahin durften.

Sollte ein WiFi langsam sein oder es öfters zu Abbrüchen kommen, ist es in den aller meisten Fällen kein fremder Benutzer. Das wird in sehr, wirklich sehr wenigen Fällen die Ursache sein.
In den meisten Fällen sind es bescheidene WiFi Router, die von der Hardware und Software nicht die tollsten Geräte sind.
Auch immer störend,
gerade in Städten, sind viele andere WiFi-Netzwerke, die sich die Frequenzen teilen und sich so gegenseitig stören.
Oder andere Geräte in der eigenen Wohnung und sehr oft der falsche Standort. Oft reicht es den Router nur etwas anders zu positionieren. Die meisten haben gefühlt ihre Router auf Hüfthöhe auf der Komode oder im schlimmsten Fall dahinter. Sieht ja nicht toll aus die Box. Einfach mal nach oben an die Wand hängen. Wirkt oft Wunder.

Aber was wäre wenn man sich jetzt einmal wirklich nicht sicher ist, ob nicht doch jemand einfach das WiFi mit benutzt? Wie oben beschrieben hilft da einfach am besten das WiFi Passwort zu ändern. Wer
möchte kann sich ja gerne die verbundene Geräte einmal auf dem Router angucken. Vielleicht hat ja doch jemand das WiFi Passwort erraten oder woher sich immer.
Aber wenn man sich unsicher ist, wie bei allen anderen Sachen auch: Passwort ändern.

Alles andere ist unnötig und kann zu mehr Problemen führen.

Wer das Passwort ändert und nicht auf anderen Geräten mühsam eingeben möchte:
   * iPhone und iPad können WiFi Passwörter teilen. Ein anderes Gerät will verbinden und ein anderen, welches schon damit verbunden ist kriegt eine Meldung und kann es dann einfach per antippen teilen.
   * Falls andere Geräte sich verbinden sollen kann man auch einfach einen QR-Code für WiFi-Name und WiFi-Passwort erstellen.

Bei der Wahl des WiFi-Passworts nicht sparsam sein. Es sollte schon sehr sicher sein. Man gibt es ja auch nicht jeden Tag ein. Es wird ja auch gespeichert.
Also ruhig auch mal nach Monaten oder 1-2 Jahren auch mal wieder ändern.
Gerade wenn man einen Router zum ersten Mal eingerichtet hat, ändern.
Oft sind die WiFi-Passwörter recht einfach. Teilweise haben manche Hersteller einfach die MAC-Adresse des Routers genommen. Oder die MAC-Adresse mit eigenen Tools in eine WiFi-Passwort umgewandelt. Was aber unsicher war und so von fremden schnell selbst berechnet werden konnte.
Also immer ruhig das WiFi-Passwort selbst einmal ändern.
