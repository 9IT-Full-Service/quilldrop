---
title: 'Kubernetes Network Policies'
date: 2023-09-21 22:00:00
author: ruediger
cover: "/images/posts/2023/09/kubernetes-netpol.webp"
tags: [Kubernetes, NetworkPolicy]
categories: 
    - Internet
preview: "In der Welt der Container-Orchestrierung spielt Kubernetes eine führende Rolle bei der Verwaltung und Automatisierung von containerisierten Anwendungen. Eine der Schlüsselkomponenten in Kubernetes ist die NetworkPolicy, ein wichtiges Werkzeug zur Kontrolle der Kommunikation zwischen Pods. In diesem Artikel werden wir uns mit den Grundlagen von NetworkPolicies befassen und wie sie zur Sicherung von Kubernetes-Clustern eingesetzt werden können." 
draft: false
top: false
type: post
hide: false
toc: false
---

[Englisch Version](/posts/2023-09-21-kubernetes-network-policies-en.html)

![Github Actions](/images/posts/2023/09/kubernetes-netpol.webp)


## Einleitung

In der Welt der Container-Orchestrierung spielt Kubernetes eine führende Rolle bei der Verwaltung und Automatisierung von containerisierten Anwendungen. Eine der Schlüsselkomponenten in Kubernetes ist die NetworkPolicy, ein wichtiges Werkzeug zur Kontrolle der Kommunikation zwischen Pods. In diesem Artikel werden wir uns mit den Grundlagen von NetworkPolicies befassen und wie sie zur Sicherung von Kubernetes-Clustern eingesetzt werden können.

## Grundlagen von NetworkPolicies

NetworkPolicies sind spezifische Objekte in Kubernetes, die festlegen, wie Gruppen von Pods miteinander und mit anderen Netzwerkendpunkten kommunizieren dürfen. Sie verwenden Labels, um Pods zu identifizieren und Regeln zu definieren, die den Datenverkehr zwischen diesen Pods steuern.

## Arten von NetworkPolicies

  1. Eingehende (Ingress) Regeln:
    * Diese Regeln steuern den eingehenden Datenverkehr zu Pods.
    * Sie können spezifizieren, welche Quellen (IP-Adressen oder Pods) erlaubt sind, auf einen Pod zuzugreifen.
  2. Ausgehende (Egress) Regeln:
    * Diese Regeln steuern den ausgehenden Datenverkehr von Pods.
    * Sie können festlegen, zu welchen Zielen (IP-Adressen oder Pods) ein Pod kommunizieren darf.

## Erstellung einer NetworkPolicy

Um eine NetworkPolicy zu erstellen, definiert man eine YAML-Datei mit den gewünschten Regeln und wendet sie auf den Kubernetes-Cluster an. Hier ist ein einfaches Beispiel für eine Ingress NetworkPolicy:

```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: example-networkpolicy
  namespace: default
spec:
  podSelector:
    matchLabels:
      app: myapp
  policyTypes:
  - Ingress
  ingress:
  - from:
    - ipBlock:
        cidr: 172.17.0.0/16

```

Diese Policy erlaubt den eingehenden Datenverkehr zu Pods mit dem Label app: myapp nur aus dem IP-Bereich 172.17.0.0/16.

## Best Practices

  1. Gewähren Sie nur die minimal notwendigen Berechtigungen.
    * Verweigern Sie standardmäßig alle Verbindungen und erlauben Sie nur spezifische.
  2. Explizite Egress-Regeln:
    * Definieren Sie klare Egress-Regeln, um den ausgehenden Datenverkehr zu kontrollieren.
  3. Verwendung von Labels:
    * Verwenden Sie Labels zur Identifizierung von Pods, um die Verwaltung von NetworkPolicies zu erleichtern.

## Fazit

NetworkPolicies sind ein unverzichtbares Werkzeug zur Sicherung von Kubernetes-Clustern. Sie ermöglichen feingranulare Kontrolle über die Kommunikation zwischen Pods und helfen, das Prinzip der geringsten Privilegien durchzusetzen. Durch den effektiven Einsatz von NetworkPolicies können Organisationen ihre Anwendungen in Kubernetes sicher und zuverlässig betreiben.

