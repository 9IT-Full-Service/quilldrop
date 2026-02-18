---
title: 'VM-Tracker - Automatische Registration and Monitoring-System'
date: 2025-11-01 05:00:00
update: 2025-11-10 05:00:00
author: ruediger
cover: "/images/posts/2025/11/vm-tracker.webp"
# images: 
#   - /images/posts/2025/08/telekom-mail-fail.webp
featureImage: /images/posts/2025/11/vm-tracker.webp
tags: [Qemu, VM, Overview]
categories: 
  - Internet
preview: ""
draft: false
top: false
type: post
hide: false
toc: false
---

# :satellite: VM-Tracker - das GetHomepage für Server und VMs :sunglasses:

## VM-Manager - VM mit Cloud-Init Templates erstellen 

Um Qemu-VMs schnell aufsetzen zu können habe ich ein Tool geschrieben mit dem ich schnell neue Server für verschiedene Zwecke erstellen kann. Da ich für Kubernetes, GitHub Runner, Nginx Webserver und andere Server in dem Tool Templates habe sind neue Server schnell aufgesetzt. Die Templates können erweitert werden oder neue erstellt werden. 

Wenn Server erstellt werden wird immer eine Serielle-Console erstellt und über ein Webinterface erreichbar gemacht. Das ist natürlich nur intern erreichbar und wenn ich von unterwegs einen Server starte weiß ich die IP erst einmal nicht. Ich kann natürlich über VPN zugreifen, aber manchmal will man einfach etwas starten und dann einfach drauf zugreifen oder wenn es extern erreichbar sein soll, einfach die IP per Port-Forward freigeben. 

Der VM-Manger wird demnächst Online gestellt und kann dann von jedem genutzt werden. 

## VM-Tracker für bessere Rechner, Server und VM Übersicht :monocle_face:

Damit ich sehe welche VMs aktuell auf den Rechnern hier gestartet sind habe ich ein Tool geschrieben das:

  * Einfach das default Netzwerkinterface sucht, 
  * Die IP-Adresse ausliest 
  * Und dann mit dem Hostname an eine API sendet. 

Die API merkt sich alle Server die sich melden. Ein Healthcheck-Tracker prüft die Meldungen alle 30 Sekunden und meldet sich ein Server 60 Sekunden nicht wird er als Offline angezeigt. Die Server können auch aus der List gelöscht werden. Offline Server verschwinden komplett, sie melden sich ja nicht mehr. Server die Online sind melden sich ja regelmässig und werden nach kurzer Zeit wieder auftauchen. 

Das Webinterface zeigt alle Server in einer schönen Übersicht an: 

![VM-Tracker Overview](/images/posts/2025/11/vm-tracker-overview.webp)

## Installation 

Update: Links zu den Installationsanleitungen sind in diesem [Blogpost](/posts/2025-11-06-vm-tracker-client-und-api-server-jetzt-self-hosted/)

Den VM-Tracker-Client kann man einfach auf den Server kopieren und ausführen. Die Binaries sind für verschiedene Betriebsysteme und Architekturen verfügbar:

+ vm-tracker-client-386
+ vm-tracker-client-amd64
+ vm-tracker-client-arm64
+ vm-tracker-client-armv6
+ vm-tracker-client-armv7
+ vm-tracker-client-darwin
+ vm-tracker-client-darwin-amd64
+ vm-tracker-client-darwin-arm64
+ vm-tracker-client-freebsd-amd64
+ vm-tracker-client-windows-386.exe
+ vm-tracker-client-windows-amd64.exe

VM-Tracker Client Ausgeführt:

```bash
./vm-tracker-client-darwin-arm64 -api https://vm-tracker.kuepper.nrw -interface en0 -interval 30
2025/11/01 09:58:48 VM Tracker Client starting...
2025/11/01 09:58:48 API Server: https://vm-tracker.kuepper.nrw
2025/11/01 09:58:48 Interface: en0
2025/11/01 09:58:48 Update interval: 30s
2025/11/01 09:58:48 Successfully registered with API server (hostname: mbair-rk.local, IP: 10.0.2.206)
```

Das Interface automatisch setzen lassen:

MacOS: 
```bash
INTERFACE=$(route get 1.1.1.1 | awk '/interface:/ {print $2; exit}')

./vm-tracker-client-darwin-arm64 -api https://vm-tracker.kuepper.nrw -interface ${INTERFACE} -interval 30
```

Linux: 
```bash
INTERFACE=$(route get 1.1.1.1 | awk '/dev/ {print $5; exit}')

./vm-tracker-client-darwin-arm64 -api https://vm-tracker.kuepper.nrw -interface ${INTERFACE} -interval 30
```

Damit man aber nicht alles manuell kopieren und ausführen muss gibt es auch ein Script für die Installation.
Die Installation kann dann einfach schnell per SSH ausgeführt werden oder so wie bei mir in den Templates von meinem VM-Tool eingefügt werden. 

Manuell ausführen:
```bash
wget -O- https://vm-tracker.kuepper.nrw/download/install-tracker.sh | bash
```

Automatisiert in Cloud-Init:

```yaml { title = "cloudinit-templates/nginx/user-data/user-data.yaml" }
...
runcmd:
  - wget -O- https://vm-tracker.kuepper.nrw/download/install-tracker.sh | bash
  - systemctl enable qemu-guest-agent
  - systemctl start qemu-guest-agent
  - echo "Cloud-init Konfiguration abgeschlossen" > /var/log/cloudinit-done.log
...
```

Damit wird der VM-Tracker Client automatisch installiert. Es wird das Binary für die richtige Zielarchitektur heruntergeladen und an die richtige Stelle kopiert. 
Zusätzlich wird ein Systemd Service angelegt, aktiviert und gestartet. 

> Hinweis: Das Installationsscript funktinoiert aktuell nur auf Linus Systemen. MacOS und Windows müssen noch integriert werden. 

So bald auch MacOS hinzugefügt ist und weitere Änderungen gemacht sind wird der VM-Tracker auch Public gestellt. 
Das ist dann aber auch nicht nur der Client, auch die API mit dem Webinterface wird verfügbar sein. 
Dann kann jeder seinen eigenen Tracker betreiben und so schnell einen Überblick über seine Infrastruktur erhalten, Probleme mit Servern oder VMs sehen. 

## VM-Manager und Vm-Tracker in Aktion

Es sind zwei Webserver vorhanden, web001 und web002. Jetzt werden 2 weitere benötigt. 
Die Konfiguration dafür ist schon vorbereitet und die VMs können erstellt werden. Nach dem Start wird dann mit Cloud-Init komplett konfiguriert und auch der VM-Tracker gestartet. So sieht man dann auch nach kurzer Zeit die neuen Server in der Übersicht des VM-Trackers. 


{{< video "/images/posts/2025/11/vm-tracker.mp4" "my-5" >}}

{{< rawhtml >}} 

<video width=100% controls>
    <source src="/videos/vm-tracker.mp4"type="video/mp4">
    Your browser does not support the video tag.  
</video>

{{< /rawhtml >}}