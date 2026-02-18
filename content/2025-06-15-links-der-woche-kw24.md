---
title: 'Links der Woche KW 24'
date: 2025-06-15 15:54:00
update: 2025-06-15 15:54:00
author: ruediger
cover: "/images/posts/2025/06/links-of-the-week.webp"
featureImage: "/images/posts/2025/06/links-of-the-week.webp"
tags: [Links]
categories: 
    - Links
preview: "Ein paar der interessante Themen die mir die letzten Tage in die Browser Tabs gespült wurden oder mit denen ich mich beschäfftigt habe."
draft: false
top: false
type: post
hide: false
toc: false
---

![Links of the week](/images/posts/2025/06/links-of-the-week.webp)


# Tailscale Your legacy VPN belongs in the past
  
Link: [Tailscale](https://tailscale.com) 

Tailscale ist ein modernes VPN-System, das auf WireGuard basiert und Geräte automatisch über ein privates Netzwerk verbindet. Es erstellt ein verschlüsseltes Mesh-Netzwerk zwischen deinen Geräten - egal ob Computer, Smartphones oder Server - ohne komplizierte manuelle Konfiguration.

Die Hauptvorteile: Einfache Installation über Apps, automatische Peer-to-Peer-Verbindungen (wo möglich), Zero-Trust-Sicherheitsmodell und zentrale Verwaltung über eine Web-Oberfläche. Du meldest dich einfach mit deinem Google/Microsoft/GitHub-Account an, installierst die App auf deinen Geräten, und sie können sich sofort sicher miteinander verbinden - auch wenn sie sich in verschiedenen Netzwerken befinden.

Besonders praktisch für Homelab-Setups, Remote-Arbeit oder den sicheren Zugriff auf eigene Dienste von unterwegs.

# Kubernetes Cluster API

Link: [Kubernetes Cluster API](https://cluster-api.sigs.k8s.io/introduction)

Kubernetes Cluster API ist ein deklaratives Tool zur Verwaltung von Kubernetes-Clustern als Code. Es behandelt Cluster-Infrastruktur wie normale Kubernetes-Ressourcen - du definierst Cluster in YAML-Manifesten und die Cluster API erstellt, skaliert und verwaltet sie automatisch.

Kernkonzept: Management-Cluster verwaltet Workload-Cluster. Du beschreibst gewünschte Cluster-Konfiguration (Node-Anzahl, VM-Größen, Kubernetes-Version) und die API sorgt für die Umsetzung.

Provider-Support: Funktioniert mit AWS, Azure, GCP, vSphere, OpenStack und vielen anderen Infrastrukturen über spezielle Provider.

Hauptvorteile: Einheitliche API für alle Cloud-Provider, GitOps-Integration, automatische Lifecycle-Verwaltung (Updates, Patches), und Cluster-Templates für Standardisierung.
Besonders nützlich für Multi-Cloud-Umgebungen oder wenn du viele Kubernetes-Cluster automatisiert verwalten möchtest.


# CyberChef - Das Cyber-Schweizer-Taschenmesser

> **Hinweis:** Vergleichbar mit [IT-Tools](https://it-tools.tech) / [IT-Tools Github](https://github.com/CorentinTh/it-tools)

Link: [CyberChef Github](https://github.com/gchq/CyberChef)

Demo: [Demo Link](https://gchq.github.io/CyberChef/)

CyberChef ist eine einfache, intuitive Web-App für die Durchführung aller Arten von "Cyber"-Operationen innerhalb eines Webbrowsers. Diese Operationen umfassen einfache Kodierungen wie XOR und Base64, komplexere Verschlüsselungen wie AES, DES und Blowfish, das Erstellen von Binär- und Hex-Dumps, Komprimierung und Dekomprimierung von Daten, Berechnung von Hashes und Prüfsummen, IPv6- und X.509-Parsing, Änderung von Zeichenkodierungen und vieles mehr.

Das Tool ist darauf ausgelegt, sowohl technischen als auch nicht-technischen Analysten zu ermöglichen, Daten auf komplexe Weise zu manipulieren, ohne sich mit komplexen Tools oder Algorithmen auseinandersetzen zu müssen. Es wurde von einem Analysten in seiner 10%-Innovationszeit über mehrere Jahre hinweg konzipiert, entworfen, entwickelt und schrittweise verbessert.

