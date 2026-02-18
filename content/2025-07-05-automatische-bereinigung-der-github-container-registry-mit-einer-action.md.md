---
title: 'Automatische Bereinigung der GitHub Container Registry mit einer eigenen Action'
date: 2025-07-05 14:00:00
update: 2025-07-05 14:00:00
author: ruediger
cover: "/images/posts/2025/07/container-registry-cleanup.webp"
featureImage: "/images/posts/2025/07/container-registry-cleanup.webp"
tags: [github, ghcr, container, docker, images, registry, cleanup]
categories: 
    - Development
preview: "Die GitHub Container Registry (ghcr.io) ist ein praktischer Service zum Hosten von Docker Images direkt bei GitHub. Bei aktiver Entwicklung sammeln sich jedoch schnell hunderte oder sogar tausende alte Container-Versionen an, die wertvollen Speicherplatz verbrauchen und die Übersicht erschweren."
draft: false
top: false
type: post
hide: false
toc: false
---

![ghcr.io container registry cleanup](/images/posts/2025/07/container-registry-cleanup.webp)

Die GitHub Container Registry (ghcr.io) ist ein praktischer Service zum Hosten von Docker Images direkt bei GitHub. Bei aktiver Entwicklung sammeln sich jedoch schnell hunderte oder sogar tausende alte Container-Versionen an, die wertvollen Speicherplatz verbrauchen und die Übersicht erschweren. 

In diesem Artikel zeige ich, wie ich eine GitHub Action entwickelt habe, die automatisch alte Container Images bereinigt, dabei aber wichtige Versionen schützt.

## Das Problem: Explodierender Container-Speicher

Bei einem meiner Projekte hatte sich die Anzahl der Container-Versionen auf über 1000 angehäuft:

```
✓ Total versions: 1041
```

Jeder Push in verschiedene Branches erzeugte neue Images mit Tags wie:
- `v1.9.64-dev.5`
- `v1.9.64-stage.2`
- `v1.9.63-develop.1`
- Untagged Versionen von gescheiterten Builds

Während aktuelle Production-Tags wie `latest` oder `v1.9.66` natürlich erhalten bleiben sollen, können alte Development- und Staging-Versionen problemlos gelöscht werden.

## Die Lösung: Container Registry Cleanup Action

Ich habe eine GitHub Action entwickelt, die diese Aufgabe automatisiert. Die Action berücksichtigt dabei mehrere wichtige Aspekte:

### ✨ Hauptfunktionen

- **🗑️ Zeitbasierte Bereinigung**: Löscht Container-Versionen, die älter als X Tage sind
- **🛡️ Tag-Schutz**: Wichtige Tags wie `latest`, `main` oder Release-Versionen bleiben erhalten
- **📊 Mindestanzahl**: Behält immer eine konfigurierbare Mindestanzahl alter Versionen
- **⚡ Batch-Verarbeitung**: Löscht pro Lauf nur eine begrenzte Anzahl, um API-Limits zu respektieren
- **🔒 Sichere Löschung**: Detaillierte Logs und mehrfache Sicherheitsprüfungen

### 🔧 Konfigurierbare Parameter

Die Action bietet umfangreiche Konfigurationsmöglichkeiten:

| Parameter | Beschreibung | Standard | Beispiel |
|-----------|--------------|----------|----------|
| `package-name` | Name des Container-Packages | *erforderlich* | `my-app` |
| `token` | GitHub Token mit `packages:write` | *erforderlich* | `${{ secrets.PAT_TOKEN }}` |
| `days-old` | Lösche Versionen älter als X Tage | `21` | `14` |
| `min-versions-to-keep` | Mindestanzahl alter Versionen behalten | `3` | `5` |
| `max-versions-per-run` | Maximale Löschungen pro Lauf | `10` | `20` |
| `protected-tags` | Regex-Pattern für geschützte Tags | `latest\|main\|master\|develop\|dev` | `latest\|stable` |
| `delete-untagged-only` | Nur untagged Versionen löschen | `false` | `true` |

## 🛡️ Intelligenter Tag-Schutz

Das Herzstück der Action ist der intelligente Tag-Schutz. Über das `protected-tags` Parameter können Sie mit Regex-Patterns definieren, welche Tags niemals gelöscht werden sollen:

### Standard-Schutz
```yaml
protected-tags: 'latest|main|master|develop|dev'
```

### Erweiteter Schutz für Release-Versionen
```yaml
protected-tags: 'latest|main|v[0-9]+\\.[0-9]+\\.[0-9]+$'
```
Dies schützt Semantic Versioning Tags wie `v1.2.3`, aber nicht `v1.2.3-beta.1`.

### Minimaler Schutz
```yaml
protected-tags: 'latest|production'
```

## 📝 Praktische Anwendungsbeispiele

### Basis-Setup für tägliche Bereinigung


```yaml
name: Container Registry Cleanup

on:
  schedule:
    - cron: '0 2 * * *'  # Täglich um 2 Uhr
  workflow_dispatch:      # Manueller Start möglich

jobs:
  cleanup:
    runs-on: ubuntu-latest
    steps:
    - name: Cleanup old container images
      uses: ruedigerp/container-registry-cleanup@v1.1
      with:
        package-name: 'my-app'
        token: ${{ secrets.PAT_TOKEN }}
        days-old: 21
        min-versions-to-keep: 3
```


### Aggressive Bereinigung für Development

```yaml
- name: Cleanup development images
  uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app'
    token: ${{ secrets.PAT_TOKEN }}
    days-old: 7           # Nur 1 Woche behalten
    min-versions-to-keep: 1
    protected-tags: 'latest|production'
    max-versions-per-run: 50
```

### Sichere Bereinigung (nur untagged)

```yaml
- name: Safe cleanup - untagged only
  uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app'
    token: ${{ secrets.PAT_TOKEN }}
    delete-untagged-only: true
```

### Mehrere Packages bereinigen

```yaml
jobs:
  cleanup:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        package: ['frontend', 'backend', 'worker', 'database']
    steps:
    - name: Cleanup ${{ matrix.package }}
      uses: ruedigerp/container-registry-cleanup@v1.1
      with:
        package-name: ${{ matrix.package }}
        token: ${{ secrets.PAT_TOKEN }}
```

## 🔑 Token-Setup

Die Action benötigt einen Personal Access Token (PAT) mit entsprechenden Berechtigungen:

### 1. PAT erstellen
1. GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Erforderliche Scopes auswählen:
   - `read:packages`
   - `write:packages` 
   - `delete:packages`

### 2. Token als Repository Secret hinzufügen
1. Repository → Settings → Secrets and variables → Actions
2. New repository secret: `PAT_TOKEN`
3. Token-Wert einfügen

## 📊 Praktische Ergebnisse

Bei meinem Projekt mit über 1000 Container-Versionen:

**Vor der Bereinigung:**
```
✓ Total versions: 1041
Old versions found: 293
```

**Nach mehreren Läufen:**
```
✓ Total versions: 751
Successfully deleted: 20 versions (untagged + old tagged)
```

Die Action löscht systematisch alte Development-Tags wie:
- `v1.0.233-amd64` ✅ gelöscht
- `v1.0.232-stage.2` ✅ gelöscht  
- `v1.9.15-develop.1` ✅ gelöscht
- `latest` ❌ geschützt
- `v1.9.66` ❌ geschützt (aktuell)

## 🚀 Automatisierung und Best Practices

### Empfohlene Scheduler-Konfiguration

```yaml
on:
  schedule:
    - cron: '0 2 * * 0'  # Sonntags um 2 Uhr (wöchentlich)
  workflow_dispatch:
```

### Monitoring mit Outputs

```yaml
- name: Cleanup old images
  id: cleanup
  uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app'
    token: ${{ secrets.PAT_TOKEN }}

- name: Report results
  run: |
    echo "Deleted: ${{ steps.cleanup.outputs.deleted-count }} versions"
    echo "Total versions: ${{ steps.cleanup.outputs.total-versions }}"
    echo "Old versions found: ${{ steps.cleanup.outputs.old-versions }}"
```

### Staging/Production Unterscheidung

```yaml
# Staging - aggressiv
- name: Cleanup staging images
  if: github.ref == 'refs/heads/develop'
  uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app-staging'
    token: ${{ secrets.PAT_TOKEN }}
    days-old: 3
    min-versions-to-keep: 1

# Production - konservativ  
- name: Cleanup production images
  if: github.ref == 'refs/heads/main'
  uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app'
    token: ${{ secrets.PAT_TOKEN }}
    days-old: 30
    min-versions-to-keep: 10
    protected-tags: 'latest|stable|v[0-9]+\\.[0-9]+\\.[0-9]+$'
```

## 🔍 Troubleshooting

### Häufige Probleme

**403 Forbidden**: 
- Prüfen Sie die PAT-Berechtigungen
- Stellen Sie sicher, dass das Package zugänglich ist

**Keine Versionen gelöscht**:
- `min-versions-to-keep` könnte zu hoch sein
- Alle Versionen könnten durch `protected-tags` geschützt sein
- Keine Versionen älter als `days-old` vorhanden

### Debug-Modus

```yaml
- name: Debug package info
  run: |
    gh api /user/packages/container/my-app | jq
    gh api /user/packages/container/my-app/versions | jq '.[0:3]'
  env:
    GITHUB_TOKEN: ${{ secrets.PAT_TOKEN }}
```

## 💡 Fazit

Die Container Registry Cleanup Action automatisiert eine zeitaufwändige Maintenance-Aufgabe und bietet dabei maximale Flexibilität und Sicherheit. Durch die intelligente Tag-Filterung bleiben wichtige Versionen erhalten, während alte Development- und Staging-Versionen systematisch entfernt werden.

**Vorteile:**
- ✅ Reduzierter Storage-Verbrauch
- ✅ Bessere Übersicht in der Registry
- ✅ Automatisierte Wartung
- ✅ Schutz wichtiger Versionen
- ✅ Konfigurierbar für verschiedene Szenarien

Die Action ist Open Source verfügbar und kann direkt über den GitHub Marketplace eingebunden werden.

---

## 📚 Vollständige README

Hier die komplette Dokumentation der Action:

---

# Container Registry Cleanup Action

A GitHub Action to automatically cleanup old container images from GitHub Container Registry (ghcr.io).

## Features

- 🗑️ Delete old container images based on age
- 🛡️ Protect important tags (latest, main, master, etc.)
- 📊 Keep minimum number of versions
- ⚡ Configurable batch processing
- 🔒 Safe deletion with detailed logging
- 🎯 Support for both tagged and untagged versions

## Usage

### Basic Example

```yaml
name: Cleanup Container Registry

on:
  schedule:
    - cron: '0 2 * * *'  # Daily at 2 AM
  workflow_dispatch:

jobs:
  cleanup:
    runs-on: ubuntu-latest
    steps:
    - name: Cleanup old container images
      uses: ruedigerp/container-registry-cleanup@v1.1
      with:
        package-name: 'my-app'
        token: ${{ secrets.PAT_TOKEN }}
        days-old: 21
        min-versions-to-keep: 3
```

### Advanced Example

```yaml
name: Advanced Container Cleanup

on:
  workflow_dispatch:

jobs:
  cleanup:
    runs-on: ubuntu-latest
    steps:
    - name: Cleanup multiple packages
      uses: ruedigerp/container-registry-cleanup@v1.1
      with:
        package-name: 'my-app'
        token: ${{ secrets.PAT_TOKEN }}
        days-old: 14
        min-versions-to-keep: 5
        max-versions-per-run: 20
        protected-tags: 'latest|stable|v[0-9]+\\.[0-9]+\\.[0-9]+'
        delete-untagged-only: false
```

### Multiple Packages

```yaml
jobs:
  cleanup:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        package: ['frontend', 'backend', 'worker']
    steps:
    - name: Cleanup ${{ matrix.package }}
      uses: ruedigerp/container-registry-cleanup@v1.1
      with:
        package-name: ${{ matrix.package }}
        token: ${{ secrets.PAT_TOKEN }}
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `package-name` | Name of the container package to cleanup | Yes | - |
| `token` | GitHub token with `packages:write` permission | Yes | - |
| `days-old` | Delete versions older than this many days | No | `21` |
| `min-versions-to-keep` | Minimum number of old versions to keep | No | `3` |
| `max-versions-per-run` | Maximum versions to delete per run | No | `10` |
| `protected-tags` | Regex pattern for protected tags (pipe-separated) | No | `latest\|main\|master\|develop\|dev` |
| `delete-untagged-only` | Only delete untagged versions | No | `false` |

## Outputs

| Output | Description |
|--------|-------------|
| `deleted-count` | Number of versions deleted |
| `total-versions` | Total versions found |
| `old-versions` | Number of old versions found |

## Token Setup

1. Create a Personal Access Token (PAT):
   - Go to GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
   - Select scopes: `read:packages`, `write:packages`, `delete:packages`

2. Add the token as a repository secret:
   - Repository → Settings → Secrets and variables → Actions
   - Name: `PAT_TOKEN`
   - Value: Your created token

## Protected Tags

By default, these tags are protected and won't be deleted:
- `latest`
- `main`
- `master` 
- `develop`
- `dev`

You can customize this with the `protected-tags` input using regex patterns.

## Safety Features

- **Minimum versions**: Always keeps a minimum number of old versions
- **Batch processing**: Limits deletions per run to avoid overwhelming the API
- **Protected tags**: Prevents deletion of important tags
- **Detailed logging**: Shows exactly what's being deleted and why
- **Dry-run capability**: Set `max-versions-per-run: 0` to see what would be deleted

## Examples by Use Case

### Conservative Cleanup (Untagged Only)
```yaml
- uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app'
    token: ${{ secrets.PAT_TOKEN }}
    delete-untagged-only: true
```

### Aggressive Cleanup (Keep Only Latest Releases)
```yaml
- uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app'
    token: ${{ secrets.PAT_TOKEN }}
    days-old: 7
    min-versions-to-keep: 1
    protected-tags: 'latest|v[0-9]+\\.[0-9]+\\.[0-9]+'
```

### Large Repository Cleanup
```yaml
- uses: ruedigerp/container-registry-cleanup@v1.1
  with:
    package-name: 'my-app'
    token: ${{ secrets.PAT_TOKEN }}
    max-versions-per-run: 50
    days-old: 30
```

## Troubleshooting

### Common Issues

1. **403 Forbidden**: Check that your PAT has the correct permissions
2. **Package not found**: Ensure the package name is correct and accessible
3. **No versions deleted**: Check that versions exist and meet the age criteria

### Debug Mode

Add this step before the cleanup to debug:

```yaml
- name: Debug package info
  run: |
    gh api /user/packages/container/my-app | jq
    gh api /user/packages/container/my-app/versions | jq '.[0:3]'
  env:
    GITHUB_TOKEN: ${{ secrets.PAT_TOKEN }}
```

## License

MIT License - see [LICENSE](https://github.com/ruedigerp/container-registry-cleanup/blob/main/LICENSE) file for details.

## Contributing

Contributions welcome! Please read [CONTRIBUTING](https://github.com/ruedigerp/container-registry-cleanup) for guidelines.
