---
title: "Smarthome ioBroker Zigbee"
date: 2022-07-02 10:45:00
lastmod: 2022-07-02 10:45:00
author: ruediger
cover: "/images/cat/technik.webp"
tags:
  - ZigBee
  - Gateway
  - RaspbeeII
  - Raspberry
  - Phoscon
  - ioBroker
preview: "Das Smarthome wurde eine zeitlang mit FHEM, HA-Bridge bzw. HomeBridge betrieben. Es lief auch recht lange sehr gut. Irgend wann wollte aber HA-Bridge nicht mehr funktionieren. Alle versuche, selbst komplette Neuinstallation konnte HA-Bridge nicht mehr zum laufen bewegen. HA-Bridge ist dann erst einmal rausgeflogen und die meisten Sachen wurde einige Zeit direkt über den Echo direkt angebunden."
categories: 
  - Internet
toc: false
hide: false
type: post
draft: false
---

## Smarthome

Das Smarthome wurde eine zeitlang mit FHEM, HA-Bridge bzw. HomeBridge betrieben. Es lief auch recht lange sehr gut. Irgend wann wollte aber HA-Bridge nicht mehr funktionieren. Alle versuche, selbst komplette Neuinstallation konnte HA-Bridge nicht mehr zum laufen bewegen. HA-Bridge ist dann erst einmal rausgeflogen und die meisten Sachen wurde einige Zeit direkt über den Echo direkt angebunden.

![ioBroker HABPanel Keller](/images/posts/ioBroker-keller.webp)

Funktioniert auch gut, aber man muss alles mögliche, was vorher über HA-Bridge oder FHEM gemacht wurde über Routinen erstellen.
Kann man machen, ist dann aber oft sehr sperrig.

Viele der vorher gemachten Dinge können aber überhaupt nicht abgebildet werden. Vor allem nicht so schnell und einfach wie vorher. In der Alexa App mal eine kleine Routine anlegen ok. Aber z.B. bestimmte Zustände von Lichtern, Steckdosen, Multimedia Geräten zum Beispiel in Abhängkeit von 1-n an-/abwesenden Personen, ist sehr aufwändig bis gar nicht möglich.

## HomeAssistant kennengelernt

Vor ein paar Monaten habe ich dann bei einem Bekannten ein HomeAssistant (HA) migriert und konnte da mal ein wenig reingucken. Das hat mich natürlich wieder getriggert. Also habe ich mir HA einmal etwas angeguckt. Da ich aber hier keinen großen Server hinstellen möchte für ein "bisschen" SmartHome und HA offiziell nicht mit Docker supported wird, hatte sich das leider erledigt.

Also noch einmal nach HA-Bridge und anderen Alternativen geschaut. Dabei dann auch mal einen genaueren Blick auf ioBroker geworfen. Die Installation war schon mal sehr simple und schnell erledigt.

![ioBroker HABPanel Wohnzimmer](/images/posts/ioBroker-wohnzimmer.webp)

## ioBroker installieren

Das Raspbian auf eine SD-Karte schmeißen und anschliessend ioBroker installieren:

    curl -sL https://iobroker.net/install.sh | bash -

Dieser Befehl startet die gesamte ioBroker-Installation.
Am Ende der Installation wird dir die URL angezeigt, wie der ioBroker zu erreichen ist.

Aufrufen von ioBroker über die Web-Oberfläche

Gehe an deinen PC/Mac und öffne die Adresse, die am Ende des Setups zu sehen war. Folge den Anweisungen im ioBroker (Lizenzbestimmungen, Grundeinstellungen)

Passwortänderung des Benutzers “pi”

Das war es dann schon.

Im ioBroker kann man sich dann erst einmal umschauen und sich einen Überlick verschaffen was ioBroker überhaupt alles kann. Und das ist eine Menge. Auch hier wie immer erst einmal alles möglich getestet und ausprobiert.

## ioBroker Adapter

Bei ioBroker nennen sich zusätzliche Tools und Anbindungen an z.B. Alexa, VW-Coonect usw. Adapter.

Für mich natürlich interessant Alexa, VW-Connect, Adapter für die Visualisierung und viele andere mehr.

### Smarthome: Alexa2 Adapter.

Der Alexa2 Adapter sorgt nicht nur dafür um Smarte Geräte im ioBroker hinzuzufügen und sie dann auf den Echo Devices zu finden. Man bekommt auch sehr viel Information über alle möglichen Alexa Geräte und über Geräte die noch über eine Alexa, bzw. den eingebauten Hub verbunden sind.

Man kann auch Sprachausgaben an Echo Geräte schicken. Und nicht nur Ankündigungen die erst durch "Alexa, habe ich neue Benachrichtigungen" ausgegeben werden. Sie können auch direkt ausgegegeben werden.
Anwendung dafür wäre z.B. eine smarte Türklingel, die ab 20 Uhr keine Töne mehr von sich gibt, sondern nur noch im Wohnzimmer eine Sprachausgabe macht. Oder per Sprache und/oder Textausgabe auf einem FireTV.

### VW-Connect

In der Fülle an Adapter ist mir dann der VW-Connect Adapter aufgefallen. Den habe ich direkt mal installiert und getestet. Da ich einen Skoda mit Skoda Connect und war sehr überrascht was man mit diesem Adapter alles machen kann. Von KM-Stand (Gesamt, letzte Fahrt), Verbrauch, bis zum Status aller Fenster/Türen (offen, geschlossen, verriegelt). Und vieles andere mehr.

In Verbindung mit Datenpunkten und eGraph können so auch schöne Graphen zu allen möglichen Daten erstellt werden.
Dazu aber mal in einem anderen Beitrag mehr Info.

### Weitere Adapter

Alle hier aktuell installierten Adapter einmal in einer Liste. Wie schon oben bei egraph geschrieben, werde ich in anderen Beiträgen bestimmt noch mal genauer auf einzelne Adapter eingehen.

Die Liste erhält man mit: `iobroker list adapters`

   * admin
   * alexa2
   * backitup
   * cloud
   * daswetter
   * deconz
   * devices
   * discovery
   * dwd
   * echarts
   * firetv
   * habpanel
   * history
   * hue
   * hue-extended
   * icons-open-icon-library-png:
   * iot
   * javascript
   * net-tools
   * openweathermap
   * pi-hole
   * ping
   * smartthings
   * socketio
   * synology
   * tado
   * time-switch
   * unifi
   * vis
   * vis-weather
   * vw-connect
   * web
   * ws
   * yahka

### Wichtigsten Adapter

Die für mich erst einmal wichtigsten Adapter sind Alexa2, deConz ZigBee, HABpanel, Philips Hue-Bridge und Extended, Javascript, Pi-Hole, Samsung Smartthings, Tado, Unifiy Network.
Damit konnten dann alle Smarthome Geräte eingebunden, gesteuert und mit den JavaScript Adapter umfangreich programmiert werden.

Die Adapter DeConz und habpanel sind für das im [letzten Artikel](/posts/2022-06-29-zigbee-gateway-raspbee-ii/) erwähnte ZigBee Gateway und für die Viualisierung der einzelnen Räume auf dem jeweiligen Tablet an der Wand. Zu den Tablets in den Räumen mit der jeweiligen Ansicht der Geräte und einer View für Multimedia werde ich auch noch einen Artikel schreiben.

![ioBroker HABPanel Küche](/images/posts/ioBroker-kueche.webp)
