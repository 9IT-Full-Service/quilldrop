---
title: 'FluxCD Image Update mit Pull Request'
date: 2025-06-16 01:00:00
xdate: 2025-06-16 01:00:00
update: 2025-06-16 06:40:00
author: ruediger
cover: "/images/posts/2025/06/fluxcd-image-update-automation.webp"
featureImage: "/images/posts/2025/06/fluxcd-image-update-automation.webp"
tags: [Kubernetes, FluxCD, ImageUpdateAutomation, GitOps]
categories: 
    - Kubernetes
preview: ""
series: ["FluxCD"] 
draft: false
top: false
type: post
hide: false
toc: true
---

![FluxCD Image Update Automation](/images/posts/2025/06/fluxcd-image-update-automation.webp)

In meinem vorherigen Artikel zu [FluxCD und Image Update Automation](https://blog.kuepper.nrw/posts/2025-06-15-fluxcd-image-update-automation) habe ich die direkte Implementierung von automatischen Container-Image-Updates demonstriert. Während diese Lösung für Development-Umgebungen praktikabel ist, erfordert eine Production-Umgebung typischerweise einen kontrollierten Review-Prozess vor dem Deployment.

# Problem der direkten Automation

Die ursprüngliche Konfiguration pushed Image-Updates direkt auf den main Branch, wodurch Änderungen sofort ohne manuellen Review-Prozess deployed werden. Für kritische Production-Workloads ist dieser Ansatz oft zu risikoreich.

Lösung: Pull Request-basierter Workflow
Durch eine geringfügige Modifikation der ImageUpdateAutomation-Konfiguration lässt sich ein Pull Request-basierter Approval-Workflow implementieren. Der entscheidende Parameter ist die Änderung des Target-Branch von main zu einem dedizierten Update-Branch:



```yaml
# Ursprünglich: direkter Push auf main
push:
  branch: main

# Modifiziert: Push auf separaten Branch für PR-Workflow  
push:
  branch: flux-updates-source
```

# Technische Implementierung

## Personal Access Token konfigurieren:

   * GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
   * Erforderliche Scopes: `repo` (Full control) und `workflow` (Update workflows)
   * Token als Repository Secret `PAT_TOKEN` hinterlegen

## GitHub Actions Workflow erstellen:
   * Datei: .github/workflows/create-pr-on-flux-updates.yml

```yaml
name: Create PR for Flux Image Updates

on:
  push:
    branches:
      - flux-updates-source

permissions:
  contents: write
  pull-requests: write

jobs:
  create-pr:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout flux-updates-source
        uses: actions/checkout@v4
        with:
          ref: flux-updates-source
          fetch-depth: 0
          token: ${{ secrets.PAT_TOKEN }}

      - name: Create and push PR branch
        run: |
          # Erstelle einen neuen Branch basierend auf flux-updates-source
          git checkout -b flux-image-updates-$(date +%Y%m%d-%H%M%S)
          
          # Push den neuen Branch
          git push origin HEAD
          
          # Speichere Branch-Name für nächsten Step
          echo "PR_BRANCH=$(git branch --show-current)" >> $GITHUB_ENV

      - name: Create Pull Request
        uses: actions/github-script@v7
        with:
          github-token: ${{ secrets.PAT_TOKEN }}
          script: |
            const { data: pullRequest } = await github.rest.pulls.create({
              owner: context.repo.owner,
              repo: context.repo.repo,
              title: '🤖 Automated Image Update',
              head: process.env.PR_BRANCH,
              base: 'main',
              body: `## Automated Image Update
              
              This PR was automatically created by Flux CD image automation.
              
              **Changes:**
              - Updated container images in production environment
              
              **Source:** flux-updates-source
              **Timestamp:** ${new Date().toISOString()}`,
            });
            
            // Labels hinzufügen
            await github.rest.issues.addLabels({
              owner: context.repo.owner,
              repo: context.repo.repo,
              issue_number: pullRequest.number,
              labels: ['automated', 'flux', 'image-update']
            });
            
            console.log(`Pull Request created: ${pullRequest.html_url}`);
```

## Automatisierter PR-Erstellungsprozess:

Der Workflow reagiert auf Push-Events zum `flux-updates-source` Branch und erstellt automatisch einen Pull Request mit den Image-Updates. Dies ermöglicht Code-Review, automatisierte Tests und kontrollierte Deployment-Zyklen.

## Branch-Management:

Optional kann automatisches Branch-Cleanup nach PR-Merge in den Repository-Settings aktiviert werden (Settings → General → "Automatically delete head branches").
Diese Implementierung kombiniert die Effizienz der automatischen Image-Erkennung mit den Governance-Anforderungen von Production-Umgebungen und bietet dabei vollständige Traceability über Git-History und PR-Metadaten.

