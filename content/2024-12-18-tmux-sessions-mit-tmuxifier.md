---
title: 'tmux Sessions mit tmuxifier verwalten'
date: 2024-12-18 06:00:00
author: ruediger
cover: "/images/posts/2024/12/tmux.webp"
featureImage: "/images/posts/2024/12/tmux.webp"
tags: [MacOS, iTerm, tmux, tmuxifier]
categories: 
    - MacOS
preview: "Wer kennt es nicht, ein Kollege kommt an den Schreibtisch und beschreibt ein Problem auf einem Cluster.
Man öffnet ein Terminal und öffnet 4-5 Tabs, da man auf 2-3 Servern per SSH muss, 2-3 Shells für Config Verzeichnis und z.B. für Tests mit Curl."
draft: false
top: false
type: post
hide: false
toc: false
---

![iTerm zsh](/images/posts/2024/12/tmux.webp)

## tmux Session mit Tmuxifier

Wer kennt es nicht: Ein Kollege kommt an den Schreibtisch und beschreibt ein Problem auf einem Cluster.
Man öffnet ein Terminal und 4–5 Tabs, da man auf 2–3 Servern per SSH zugreifen muss, sowie 2–3 Shells für das Config-Verzeichnis und z. B. für Tests mit Curl.

Allein die Logins, Verzeichniswechsel und das Öffnen der Konfigurationen zum Vergleichen dauert, bis man endlich loslegen kann.

Ich benutze im Terminal Tmux und dafür das Tool Tmuxifier, um diese Schritte zu automatisieren.

[Tmuxify](https://github.com/jimeh/tmuxifier) dein Tmux!
Erstelle, bearbeite, verwalte und lade komplexe Tmux-Konfigurationen für Sitzungen, Fenster und Panes mit Leichtigkeit.

Kurz gesagt, Tmuxifier ermöglicht es dir, „Layout“-Dateien einfach zu erstellen, zu bearbeiten und zu laden. Diese Dateien sind einfache Shell-Skripte, in denen du den Tmux-Befehl sowie von Tmuxifier bereitgestellte Hilfsbefehle verwendest, um Tmux-Sitzungen und Fenster zu verwalten.


### Fenster-Layouts

Fenster-Layouts erstellen ein neues Tmux-Fenster, wobei optional der Fenstertitel und das Root-Verzeichnis festgelegt werden können, in dem sich alle Shells standardmäßig befinden. Sie ermöglichen es dir, ein Fenster einfach in spezifisch dimensionierte Panes zu unterteilen und es nach deinen Wünschen anzupassen.

Du kannst ein Fenster-Layout direkt in deiner aktuellen Tmux-Sitzung laden oder es in ein Sitzungs-Layout integrieren, sodass das Fenster zusammen mit der Sitzung erstellt wird.

### Sitzungs-Layouts

Sitzungs-Layouts erstellen eine neue Tmux-Sitzung und legen dabei optional einen Sitzungstitel und ein Root-Verzeichnis fest, in dem sich alle Shells der Sitzung standardmäßig befinden. Fenster können der Sitzung entweder durch das Laden bestehender Fenster-Layouts hinzugefügt oder direkt innerhalb der Sitzungs-Layout-Datei definiert werden.


## Tmuxifier installieren 

Installation von Tmuxifier mit homebrew: 

    git clone https://github.com/jimeh/tmuxifier.git ~/.tmuxifier

## Konfiguration 

Als Erstes erstellt man eine neue Session, die man in Zukunft einfach mit load-session immer wieder aufrufen kann:

    tmuxifier create-session example 

Dadurch wird die Datei `.tmuxifier/layouts/example.session.sh` angelegt. Diese wird automatisch geöffnet und kann anschließend bearbeitet werden.

    session_root "~/Code/example-cluster/"

    if initialize_session "example"; then

        # Create a new window inline within session layout definition.
        new_window "cluster1"
        new_window "cluster1"
        new_window "deploy"
        new_window "services"
        new_window "shell

        # Select the default active window on session creation.
        select_window 0
        run_cmd "KUBECONFIG=/Users/rk/Code/example-cluster/cluster1-kubeconfig.yaml"
        run_cmd "cluster1"
        select_window 1
        run_cmd "KUBECONFIG=/Users/rk/Code/example-cluster/cluster2-kubeconfig.yaml"
        run_cmd "cluster1"
        select_window 2
        run_cmd "cd deployments; nvim "
        select_window 3
        run_cmd "cd services; nvim "
        select_window 4 
        run_cmd "cd deployments"
        select window 0
    fi

    # Finalize session creation and switch/attach to it.
    finalize_and_go_to_session

Muss ich auf die beiden Cluster und die entsprechenden Verzeichnisse zugreifen, um etwas zu debuggen, auszurollen usw., kann ich einfach eine neue Tmux-Session erstellen, die dann 4 Tmux-Fenster öffnet.

In den ersten beiden Fenstern sind die beiden Kubernetes-Cluster mit k9s geöffnet.

Im 3. und 4. Fenster befinden sich die Deployments (Helm-Charts) bzw. Services (Code und Docker-Skripte) und sind direkt in nvim geöffnet.

Das 5. Fenster ist eine Shell, die direkt im Verzeichnis deployments startet, um schnell Pakete erstellen oder Deployments durchführen zu können.

Die Session starte ich dann mit: 

    tmuxifier load-session example 

Nach 1–2 Sekunden ist in Tmux eine neue Session mit dem Namen Example geöffnet – mit 5 Fenstern, die genau so gestartet und geöffnet wurden, wie in der Tmuxifier-Konfiguration angegeben.

Falls dich jemand fragt, warum ich die kubeconfig.yaml-Dateien pro Cluster habe und nicht in der globalen Benutzer-Kubeconfig-Context-Datei: Ich benutze ein Tool (direnv) für Environment-Variablen (Env-Vars) in der Shell, das in Verzeichnissen die entsprechenden Env-Vars setzt, die dort benötigt werden. Das werde ich in einem anderen Post genauer beschreiben.

Öffnet man jetzt die Session- und Fensterliste, sieht man die Session mit allen Fenstern und kann diese auswählen.
Bei mir gibt es aktuell drei Sessions mit mehreren Fenstern. Die Session aus der Konfiguration ist die Session `example` mit ihren fünf Tmux-Fenstern.

    (1)   + dev-cluster: 2 windows
    (2)   - example: 4 windows
    (3)   ├─> 0: cluster1*
    (4)   ├─> 1: cluster1
    (5)   ├─> 2: deploy
    (6)   └─> 3: services
    (7)   └─> 3: shell
    (8) + prod-cluster: 3 windows

### Tmuxifier erleichtert wiederkehrende Sessions

Die Aufgaben und die damit einhergehenden, immer wieder gleichen Schritte, die man täglich ausführt, lassen sich mit Tmuxifier schnell und einfach automatisieren. Die Konfiguration ist recht simpel, und wenn man – wie ich – täglich auf mehreren Servern und Clustern arbeitet, kann man sich hierdurch einige Tasks erheblich erleichtern.

Ich habe eine Konfiguration für einen Test- und einen Produktions-Cluster erstellt, sodass ich schnell die wichtigsten Dinge geöffnet habe. Zusätzlich nutze ich das oben beschriebene Beispiel für meinen Blog und andere Tools auf meinen privaten Kubernetes-Clustern. Dadurch kann ich mein Blog und andere Projekte schnell aktualisieren.

Für Migrationen erstelle ich ebenfalls eine Tmuxifier-Konfiguration. In den Tagen vor der Migration bin ich häufig auf den entsprechenden Servern eingeloggt, um alles vorzubereiten – oft auch mit Migrationstests, die den eigentlichen Migrationstag simulieren. Die Tmux-Fenster und Sessions kann ich nach der Vorbereitung schließen und dann am Tag der Migration einfach die Session in Tmux neu erstellen lassen.

