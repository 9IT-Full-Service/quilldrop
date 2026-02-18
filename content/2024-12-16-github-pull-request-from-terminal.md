---
title: 'Github Pull Requests mit gh cli Tool'
date: 2024-12-15 20:00:00
author: ruediger
cover: "/images/posts/2024/12/tmux.webp"
featureImage: "/images/posts/2024/12/tmux.webp"
tags: [MacOS, iTerm, tmux]
categories: 
    - MacOS
preview: "Meine tmux konfiguration."
draft: true
top: false
type: post
hide: true
toc: false
---

![iTerm zsh](/images/posts/2024/12/github.webp)

Die Änderungen sind gemacht, der Commit erstellt und gepushed. Jetzt fehlt nur noch der Pull Request um den neuen Code online zu bekommen. 

Browser öffnen, Github URL eingeben, Projekt und Repo wählen, Pull Requests anklicken. 
Neuen Pull Request erstellen, den Branch auswählen, Zielbranch auswählen und Titel + Body eintragen. 

Nervig? Yep. Das geht auch per Terminal in der Shell. 

## gh cli Tool installieren

    brew install gh 

    gh auth login

    git checkout -b feature-branch
    git add .
    git commit -m "Deine Nachricht für den Commit"

    git push origin feature-branch

    gh pr create --base main --title "Titel des PR" --body "Beschreibung des PR"
    
