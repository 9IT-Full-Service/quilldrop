---
title: 'GitHub Actions: Das Triggern von Anderen Actions'
date: 2023-09-21 20:00:00
author: ruediger
cover: "/images/posts/github-actions-banner.webp"
tags: [GitHub, Action, Workflow, Trigger]
categories: 
    - Internet
preview: "GitHub Actions ist ein CI/CD-Service (Continuous Integration/Continuous Deployment) von GitHub, der es Entwicklern ermöglicht, Automatisierungsworkflows direkt in ihren GitHub-Repositories zu erstellen und zu teilen. Eine häufige Anforderung ist das Triggern einer Action aus einer anderen Action heraus. In diesem Artikel werden wir untersuchen, wie man GitHub Actions effektiv verkettet und so komplexere Automatisierungsprozesse erstellt." 
draft: false
top: false
type: post
hide: false
toc: false
---

[English Version](/posts/2023-09-21-github-actions-das-triggern-von-anderen-actions-en.html)

![Github Actions](/images/posts/github-actions-banner.webp)

## Einleitung

GitHub Actions ist ein CI/CD-Service (Continuous Integration/Continuous Deployment) von GitHub, der es Entwicklern ermöglicht, Automatisierungsworkflows direkt in ihren GitHub-Repositories zu erstellen und zu teilen. Eine häufige Anforderung ist das Triggern einer Action aus einer anderen Action heraus. In diesem Artikel werden wir untersuchen, wie man GitHub Actions effektiv verkettet und so komplexere Automatisierungsprozesse erstellt.

## Grundlagen von GitHub Actions

GitHub Actions basieren auf Workflows, die in YAML-Dateien innerhalb des .github/workflows-Verzeichnisses eines Repositories definiert sind. Ein Workflow besteht aus einem oder mehreren Jobs, und jeder Job besteht aus Schritten, die Befehle oder Actions sind.

## Triggern von Anderen Actions

Um eine GitHub Action aus einer anderen Action heraus zu triggern, gibt es verschiedene Ansätze:

Man kann ein Repository Dispatch Event von einer Action auslösen, um einen Workflow in demselben oder einem anderen Repository zu starten.
Dies erfordert ein Personal Access Token, um das Event zu authentifizieren.

  1. Workflow Dispatch Event:
    * Mit dem Workflow Dispatch Event kann man manuell einen Workflow starten.
    * Es kann auch durch einen API-Aufruf ausgelöst werden, was es von einer anderen Action aus startbar macht.
  2. Scheduled Workflows:
    * Workflows können so konfiguriert werden, dass sie zu bestimmten Zeiten ausgeführt werden.
    * Dies kann genutzt werden, um nachfolgende Workflows zu triggern.

## Beispiel: Repository Dispatch Event

Hier ist ein einfaches Beispiel, wie man ein Repository Dispatch Event verwendet, um einen Workflow zu triggern:

```
# .github/workflows/triggering-workflow.yml
name: Triggering Workflow
on: [push]
jobs:
  trigger:
    runs-on: ubuntu-latest
    steps:
      - name: Trigger Another Workflow
        run: |
          curl -XPOST -u "USERNAME:${{ secrets.PAT }}" \
            -H "Accept: application/vnd.github.everest-preview+json" \
            -H "Content-Type: application/json" \
            https://api.github.com/repos/OWNER/REPO/dispatches \
            --data '{"event_type": "triggered"}'

```

```
# .github/workflows/triggered-workflow.yml
name: Triggered Workflow
on:
  repository_dispatch:
    types: [triggered]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v2
      # Weitere Schritte...

```

In diesem Beispiel wird durch einen push Event der triggering-workflow.yml Workflow gestartet, der dann das triggered-workflow.yml durch ein Repository Dispatch Event auslöst.

### Fazit

Das Triggern von GitHub Actions aus anderen Actions heraus ermöglicht die Erstellung von komplexen und flexiblen Automatisierungsworkflows. Durch die Verwendung von Methoden wie dem Repository Dispatch Event oder dem Workflow Dispatch Event können Entwickler ihre CI/CD-Prozesse effizient gestalten und optimieren.