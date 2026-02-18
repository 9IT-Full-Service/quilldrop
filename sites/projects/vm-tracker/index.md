---
title: 'VM-Tracker'
author: ruediger
date: 2016-01-08 00:39:15 +0200
update: 2023-09-01 00:39:15 +0200
draft: false
tags: 
  - Projekte
  - VM-Tracker
categories: 
  - Projekte
type: page
---

# VM-Tracker

## VM-Tracker für bessere Rechner, Server und VM Übersicht :monocle_face:

Damit ich sehe welche VMs aktuell auf den Rechnern hier gestartet sind habe ich ein Tool geschrieben das:

  * Einfach das default Netzwerkinterface sucht, 
  * Die IP-Adresse ausliest 
  * Und dann mit dem Hostname an eine API sendet. 

Die API merkt sich alle Server die sich melden. Ein Healthcheck-Tracker prüft die Meldungen alle 30 Sekunden und meldet sich ein Server 60 Sekunden nicht wird er als Offline angezeigt. Die Server können auch aus der List gelöscht werden. Offline Server verschwinden komplett, sie melden sich ja nicht mehr. Server die Online sind melden sich ja regelmässig und werden nach kurzer Zeit wieder auftauchen. 

Das Webinterface zeigt alle Server in einer schönen Übersicht an: 

![VM-Tracker Overview](/images/posts/2025/11/vm-tracker-overview.webp)

# VM-Tracker Server installieren. 

Damit man selbst einen Endpunkt für die Clients, unter einer eigenen Domain hat, kann der API-Server jetzt auch Self-Hosted betrieben werden. 

Anleitungen gibt es für: 

* [Binary Installation](https://vm-tracker.kuepper.nrw/docs/de/api/installation_binary/)
* [Docker](https://vm-tracker.kuepper.nrw/docs/de/api/installation_docker/)
* [Kubernetes](https://vm-tracker.kuepper.nrw/docs/de/api/installation_kubernetes/)
* [Helm](https://vm-tracker.kuepper.nrw/docs/de/api/installation_helm/)
* [Kustomization](https://vm-tracker.kuepper.nrw/docs/de/api/installation_kustomization/)
* [FluxCD Kustomization](https://vm-tracker.kuepper.nrw/docs/de/api/installation_fluxcd_kustomization/)
* [FluxCD Helm Release](https://vm-tracker.kuepper.nrw/docs/de/api/installation_fluxcd_helm_release/)

Bei allen Installationen kann man diese ENV-Variabeln setzen:

* API_BASE_URL=https://vm-tracker.example.com
* BASE_URL=https://vm-tracker.example.com

Diese werden benutzt für z.B. Ingress, um die API über die Domain erreichbar zu machen. Ausserdem wird damit das Installations-Skript erstellt, damit die Clients sich mit der richtigen API verbinden. 
Dadurch kann der Client schnell und unkompliziert auf allen Sytemen installiert werden. 

# VM-Tracker Client installieren. 

Welche Möglichkeiten bei der Client Installation zur Verfügung stehen ist hier beschrieben:

* [Binary](https://vm-tracker.kuepper.nrw/docs/de/client/installation_binary/)
* [Skript](https://vm-tracker.kuepper.nrw/docs/de/client/installation_script/)
* [Systemd](https://vm-tracker.kuepper.nrw/docs/de/client/installation_systemd/)
* [Cloud-Init](https://vm-tracker.kuepper.nrw/docs/de/client/installation_cloud_init/)
* [Shelly Script](https://vm-tracker.kuepper.nrw/docs/de/client/installation_shelly/)

Die komplette Dokumentation in deutsch ist [hier](https://vm-tracker.kuepper.nrw/docs/de/) und die englische [hier](https://vm-tracker.kuepper.nrw/docs/en/).

