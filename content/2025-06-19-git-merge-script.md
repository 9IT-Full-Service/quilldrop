---
title: 'Git Merge Script für Deployment Workflow'
date: 2025-06-19 08:00:00
update: 2025-06-19 10:50:00
author: ruediger
cover: "/images/posts/2025/06/git-merge-script.webp"
featureImage: "/images/posts/2025/06/git-merge-script.webp"
tags: [Git, Merge, Workflow, Merge. Braches]
categories: 
    - Development
preview: "Dieses Script automatisiert den gesamten Merge-Prozess zwischen verschiedenen Branches in einem Git-Repository und nutzt dabei die GitHub CLI für Pull Request Management. Es ist speziell für Teams entwickelt, die einen strukturierten Deployment-Workflow verwenden."
draft: false
top: false
type: post
hide: false
toc: false
---

![Automatisierte Kubernetes Volume-Backups](/images/posts/2025/06/git-merge-script.webp)

# Git Merge Automation Script

Ein Bash-Script zur Automatisierung von Git-Merge-Workflows über GitHub Pull Requests.

## 🎯 Zweck

Dieses Script automatisiert den gesamten Merge-Prozess zwischen verschiedenen Branches in einem Git-Repository und nutzt dabei die GitHub CLI für Pull Request Management. Es ist speziell für Teams entwickelt, die einen strukturierten Deployment-Workflow verwenden.

## 📋 Voraussetzungen

- **Git** installiert und konfiguriert
- **GitHub CLI (gh)** installiert und authentifiziert
- Berechtigung zum Erstellen und Mergen von Pull Requests im Repository
- Bash-Shell (Linux/macOS/WSL)

### GitHub CLI Installation

```bash
# macOS (Homebrew)
brew install gh

# Ubuntu/Debian
sudo apt install gh

# Windows (Chocolatey)
choco install gh
```

### GitHub CLI Authentifizierung

```bash
gh auth login
```

## 🚀 Installation

* Script herunterladen und ausführbar machen:

<script src="https://gist.github.com/ruedigerp/810aa03b35118f14d7c8d9b8cd99e955.js"></script>

```bash
chmod +x merge.sh
```

* Optional: In ein Verzeichnis im PATH verschieben:

```bash
sudo mv merge.sh /usr/local/bin/merge
```

## 📖 Verwendung

### Grundlegende Syntax

```bash
./merge.sh <command>
```

### Verfügbare Commands

#### 1. `dev-stage` - Development zu Staging

Merged den `dev` Branch in den `stage` Branch.

```bash
./merge.sh dev-stage
```

**Was passiert:**
- Erstellt einen Pull Request von `dev` → `stage`
- PR-Titel: "Deploy dev to stage - YYYY-MM-DD"
- Merged den PR automatisch nach erfolgreichen Checks

#### 2. `stage-main` - Staging zu Production

Merged den `stage` Branch in den `main` Branch (Production).

```bash
./merge.sh stage-main
```

**Was passiert:**
- Erstellt einen Pull Request von `stage` → `main`
- PR-Titel: "Deploy stage to main - YYYY-MM-DD"
- Merged den PR automatisch für Production-Release

#### 3. `back-merge` - Rück-Synchronisation

Synchronisiert den `main` Branch zurück zu `stage` und `dev` nach einem Production-Release.

```bash
./merge.sh back-merge
```

**Was passiert:**
1. Holt den neuesten Tag (Release-Version)
2. Erstellt PR: `main` → `stage` (Back-merge)
3. Erstellt PR: `main` → `dev` (Back-merge)
4. Beide PRs werden mit `[skip ci]` Flag gemerged

#### 4. `anywhere` - Flexibler Merge

Merged einen beliebigen Branch in einen anderen.

```bash
./merge.sh anywhere <target-branch> <source-branch>
```

**Beispiele:**
```bash
./merge.sh anywhere stage dev        # dev → stage
./merge.sh anywhere main hotfix     # hotfix → main
./merge.sh anywhere feature dev     # dev → feature
```

## 🔄 Typischer Workflow

### Standard Development Cycle

```bash
# 1. Development → Staging
./merge.sh dev-stage

# 2. Testing auf Staging...

# 3. Staging → Production
./merge.sh stage-main

# 4. Back-merge nach Release
./merge.sh back-merge
```

### Hotfix Workflow

```bash
# 1. Hotfix direkt zu Production
./merge.sh anywhere main hotfix-branch

# 2. Back-merge
./merge.sh back-merge
```

## 🎛️ Branch-Struktur

Das Script ist für folgende Branch-Struktur optimiert:

```
main (Production)
├── stage (Staging/Testing)
└── dev (Development)
    ├── feature/xyz
    ├── bugfix/abc
    └── hotfix/urgent
```

## ⚙️ Funktionsweise

### Pull Request Erstellung

- **Automatische Titel**: Datum-basierte PR-Titel
- **Beschreibungen**: Vordefinierte, aussagekräftige Beschreibungen
- **Auto-Merge**: PRs werden automatisch gemerged wenn alle Checks bestehen

### Back-Merge Besonderheiten

- **Tag-Detection**: Automatische Erkennung des neuesten Release-Tags
- **CI Skip**: Back-Merges verwenden `[skip ci]` um unnötige Builds zu vermeiden
- **Synchronisation**: Stellt sicher, dass alle Branches auf dem gleichen Stand sind

## 🛠️ Fehlerbehebung

### Häufige Probleme

#### GitHub CLI nicht authentifiziert
```bash
gh auth status
gh auth login
```

#### Merge-Konflikte
- Das Script stoppt bei Konflikten
- Löse Konflikte manuell im GitHub Web-Interface
- Oder löse sie lokal und push erneut

#### Fehlende Berechtigung
```bash
# Prüfe Repository-Berechtigung
gh repo view
```

#### Branch existiert nicht
```bash
# Verfügbare Branches anzeigen
git branch -a
```

### Debug-Modus

Für detaillierte Ausgabe:
```bash
bash -x ./merge.sh dev-stage
```

## 🔒 Sicherheitshinweise

- **Branch Protection**: Aktiviere Branch Protection Rules für `main` und `stage`
- **Required Reviews**: Konfiguriere erforderliche Code-Reviews
- **Status Checks**: Stelle sicher, dass CI/CD-Checks aktiviert sind
- **Auto-Merge**: Funktioniert nur wenn alle konfigurierten Checks bestehen

## 📝 Anpassungen

### Custom Branch Namen

Passe die Branch-Namen im Script an deine Naming-Convention an:

```bash
# Ändere im Script:
--base stage    # zu deinem Staging-Branch
--head dev      # zu deinem Development-Branch
```

### Custom PR-Titel

Ändere die PR-Titel-Templates:

```bash
--title "Deploy dev to stage - $(date +%Y-%m-%d)"
```

### Custom Commit Messages

Ändere die Merge-Commit-Messages:

```bash
--subject "Release stage"
```

## 🤝 Best Practices

1. **Teste immer auf Staging** bevor du zu Production merged
2. **Führe Back-Merges regelmäßig durch** um Branches synchron zu halten
3. **Prüfe CI/CD-Status** vor dem Merge
4. **Verwende aussagekräftige Commit-Messages** in deinen Feature-Branches
5. **Erstelle Tags** für Production-Releases

## 📊 Monitoring

### PR-Status prüfen

```bash
# Alle offenen PRs anzeigen
gh pr list

# Specific PR-Status
gh pr view <PR-NUMBER>
```

### Recent Merges

```bash
# Letzten 5 Commits auf main
git log --oneline -5 main

# Tags anzeigen
git tag --sort=-version:refname | head -5
```

## 🆘 Support

Bei Problemen:

1. Prüfe die Logs des Scripts
2. Verifiziere GitHub CLI Authentication
3. Prüfe Repository-Berechtigungen


<!-- <div x_data-comment-post-id="2025-06-19-git-merge-script"></div> -->
