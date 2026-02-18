---
title: 'Debugging von Distroless-Containern in Kubernetes'
date: 2025-10-17 08:00:00
update: 2025-10-17 08:00:00
author: ruediger
cover: "/images/posts/2025/10/debugging-von-distroless-containern-in-kubernetes.webp"
featureImage: "/images/posts/2025/10/debugging-von-distroless-containern-in-kubernetes.webp"
# images: 
#   - /images/posts/2025/08/telekom-mail-fail.webp
tags: [Kubernetes, Debugging, Distroless-Containern]
categories: 
  - Kubernetes
preview: "Weil es keine Shell im Container gibt! Daher können Sie weder `ps`, `curl`, `netstat` noch andere nützliche Tools ausführen, um den Produktionsfehler zu debuggen. Lösung: Ephemeral Containers in Kubernetes"
draft: false
top: false
type: post
hide: false
toc: false
---

Distroless Container Images sind schnell, sicher und perfekt für die Produktion, können aber beim Debugging zum Albtraum werden.

In diesem Blog erfährst Du, warum Distroless Container Images so leistungsfähig sind, warum ihr Debugging schwierig ist und wie Ephemeral Containers in Kubernetes das Live-Debugging wieder einfach machen, ohne Ihre Anwendung neu zu starten.

# Was sind Distroless Images?

Distroless-Container sind Docker-Images, die keine Linux-Distribution (wie Ubuntu oder Alpine) enthalten. Sie beinhalten nur:

* Ihre Anwendung
* Deren Laufzeitumgebung (z.B. Java, Node.js, Go)

# Was fehlt darin?

* Keine Shell (sh, bash)
* Kein Paketmanager (apt, apk)
* Keine Debug-Tools (curl, ping, ps)

# Warum Distroless Images verwenden?

Distroless Images sind für Sicherheit, Performance und Einfachheit optimiert und werden aufgrund folgender Vorteile für den Einsatz in Produktionsumgebungen empfohlen:

* Schnellere Downloads und Deployments durch kleinere Docker-Image-Größe, wodurch Cloud-Infrastrukturkosten sowohl für  Netzwerk als auch Storage gespart werden
* Reduzierte Angriffsfläche durch weniger Binaries und Bibliotheken
* Kein Shell-Zugriff bedeutet, dass niemand per exec in den Container gelangen kann

# Das Debugging-Dilemma

Stellen Dir sich vor, Deine Node.js-App läuft in einem Distroless Container Image und wirft 500-Fehler. Um dies zu debuggen, haben Du versucht:

```bash
kubectl exec -it my-production-application -- sh
```

erhälst jedoch folgende Fehlermeldung:

```bash
error: unable to upgrade connection: container not found or shell not available
```

Warum? Weil es keine Shell im Container gibt! Daher können Sie weder `ps`, `curl`, `netstat` noch andere nützliche Tools ausführen, um den Produktionsfehler zu debuggen.

## Lösung: Ephemeral Containers in Kubernetes

Ephemeral Containers in Kubernetes ermöglichen es Dir, einen temporären Debug-Container in einen laufenden Pod zu injizieren, ohne Deine Anwendung zu stoppen, neu zu erstellen oder neu zu starten, um den Fehler im Anwendungs-Pod zu debuggen.

### Schritte:

1. Erstelle Dir einen Beispiel-Anwendungs-Pod mit dem Distroless Container Image

```yaml
apiVersion: v1
kind: Pod
metadata:
 name: my-app
spec:
 containers:
 — name: app
 image: gcr.io/distroless/static
 command: ["sleep", "3600"]
```

```bash
kubectl apply -f my-app.yaml
```

2. Debug-Container injizieren

```yaml
kubectl debug -it my-app \
 --target=app \
 --image=busybox \
 --name=debugger
```

Jetzt befindest Du Dich im `busybox`-Container, der die Netzwerk-, Prozess- und IPC-Namespaces mit Ihr Deinem  Anwendungs-Pod teilt, was Dir hilft, Dir Anwendungsprobleme zu debuggen.

```text
ps              # Prozesse des nginx-Container-Namespace auflisten
ls /            # Dateisystem durchsuchen
cat /etc/hosts  # Netzwerkkonfiguration anzeigen
ping 8.8.8.8    # Ausgehende Verbindung testen
wget <url>      # Falls unterstützt, Downloads testen
```

Und Deine Anwendung läuft ohne Neustart weiter.

