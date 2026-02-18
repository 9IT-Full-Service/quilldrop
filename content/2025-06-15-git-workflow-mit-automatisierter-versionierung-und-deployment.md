---
title: 'Moderne Blog-Entwicklung: Ein durchgängiger Git-Workflow mit automatisierter Versionierung und Deployment'
date: 2025-06-15 20:00:00
update: 2025-06-15 20:10:00
author: ruediger
cover: "/images/posts/2025/06/IMG_0971.webp"
featureImage: "/images/posts/2025/06/IMG_0971.webp"
tags: [Git, GitHub, Workflow]
categories: 
    - Git
preview: "In der modernen Softwareentwicklung ist ein sauberer Deployment-Workflow essentiell für die Qualitätssicherung und effiziente Zusammenarbeit. In diesem Artikel stelle ich meinen bewährten Workflow für die Blog-Entwicklung vor, der drei Stages nutzt und durch automatisierte Versionierung sowie GitOps-Prinzipien unterstützt wird."
series: ["FluxCD"] 
draft: false
top: false
type: post
hide: false
toc: false
---

![Links of the week](/images/posts/2025/06/IMG_0971.webp)

# Moderne Blog-Entwicklung: Ein durchgängiger Git-Workflow mit automatisierter Versionierung und Deployment

In der modernen Softwareentwicklung ist ein sauberer Deployment-Workflow essentiell für die Qualitätssicherung und effiziente Zusammenarbeit. In diesem Artikel stelle ich meinen bewährten Workflow für die Blog-Entwicklung vor, der drei Stages nutzt und durch automatisierte Versionierung sowie GitOps-Prinzipien unterstützt wird.

## Die drei Stages: Dev, Stage und Prod

Mein Blog-Setup basiert auf drei klar getrennten Umgebungen:

- **Dev**: Die Entwicklungsumgebung für neue Features und Experimente
- **Stage**: Die Staging-Umgebung für finale Tests unter produktionsähnlichen Bedingungen
- **Prod**: Die Live-Produktionsumgebung für Endnutzer

Diese Trennung spiegelt sich direkt in der Git-Branch-Struktur wider. Ich verwende entsprechend benannte Branches `dev`, `stage` und `main` (für Prod), wobei `main` als Hauptbranch fungiert.

## Automatisierte Versionierung mit Semantic Versioning

Ein Kernstück meines Workflows ist die automatisierte Versionierung durch GitHub Actions. Dabei nutze ich Semantic Versioning (SemVer) in Kombination mit aussagekräftigen Commit-Messages.

### Der Entwicklungsprozess

Alle Änderungen beginnen im `dev`-Branch. Wenn ich neue Features entwickle oder Bugs behebe, verwende ich spezifische Präfixe in meinen Commit-Messages:

```bash
git commit -m "fix: Behebung des Responsive-Problems im Header"
# oder
git commit -m "feat: Neue Kommentarfunktion hinzugefügt"
```

Diese strukturierten Commit-Messages sind nicht nur für Menschen lesbar, sondern triggern auch automatisch die Versionierung in meinen GitHub Actions. Jeder Commit mit `fix:` oder `feat:` erstellt automatisch einen neuen Git-Tag im Format `v1.9.40-dev.1`, `v1.9.40-dev.2`, usw.

### Von Dev zu Stage: Der erste Qualitätsfilter

Sobald alle Entwicklungsarbeiten in `dev` abgeschlossen und getestet sind, erstelle ich einen Pull Request von `dev` nach `stage`. Dieser Merge-Vorgang ist ein bewusster Schritt, der signalisiert: “Diese Änderungen sind bereit für die finale Testphase.”

Der Merge nach `stage` triggert automatisch die Erstellung eines neuen Tags im Format `v1.9.40-stage.1`. Falls weitere Änderungen aus `dev` folgen, werden diese als `v1.9.40-stage.2`, `v1.9.40-stage.3` usw. versioniert.

### Der finale Schritt: Stage zu Prod

Nach erfolgreichen Tests in der Staging-Umgebung erfolgt der finale Merge von `stage` nach `prod` (main). Dieser Schritt erstellt die finale Produktionsversion, zum Beispiel `v1.9.41` - eine saubere, produktionsreife Versionsnummer ohne Zusätze.

### Backward-Merge: Der Kreis schließt sich

Ein oft übersehener, aber wichtiger Schritt ist der Backward-Merge. Nach dem Release führe ich einen Merge von `prod` zurück nach `stage` und `dev` durch. Dies stellt sicher, dass alle Branches synchron bleiben und eventuelle Hotfixes oder produktionsrelevante Anpassungen in alle Umgebungen übernommen werden.

Nach diesem Backward-Merge beginnt der Zyklus von neuem, jetzt mit der Basis-Version `v1.9.41-dev.1`.

## Docker-Images und Container-Orchestrierung

Parallel zur Git-Versionierung erstellen meine GitHub Actions automatisch Docker-Images für jede Stage. Diese Images werden entsprechend der Branch-Namen getaggt:

- `myblog:v1.9.40-dev.1` für Entwicklungsversionen
- `myblog:v1.9.40-stage.1` für Staging-Versionen
- `myblog:v1.9.41` für Produktionsversionen

Diese konsistente Tagging-Strategie ermöglicht es, jederzeit nachzuvollziehen, welche Version in welcher Umgebung läuft.

## GitOps mit FluxCD: Automatisierte Deployments

Der finale Baustein meines Workflows ist die automatisierte Deployment-Pipeline mit FluxCD. Durch ImageUpdateAutomations überwacht FluxCD kontinuierlich meine Container-Registry auf neue Images.

Sobald ein neues Image für eine Stage verfügbar ist, aktualisiert FluxCD automatisch die entsprechenden Deployments im Kubernetes-Cluster. Dies bedeutet:

- Neue Dev-Images werden automatisch in die Entwicklungsumgebung deployed
- Stage-Images landen automatisch in der Staging-Umgebung
- Produktions-Images werden nach erfolgreichem Merge automatisch live geschaltet

## Vorteile dieses Workflows

### Nachvollziehbarkeit

Jede Änderung ist durch die semantische Versionierung klar nachverfolgbar. Ein Blick auf die Git-Tags zeigt sofort, welche Version welche Features oder Fixes enthält.

### Qualitätssicherung

Der mehrstufige Prozess stellt sicher, dass nur durchgetestete Änderungen in die Produktion gelangen. Die Staging-Umgebung fungiert als wichtige Barriere vor dem Live-System.

### Automatisierung reduziert Fehler

Durch die Automatisierung von Versionierung und Deployment werden menschliche Fehler minimiert. Vergessene Tags oder manuelle Deployment-Fehler gehören der Vergangenheit an.

### Schnelle Iteration

Entwickler können sich auf die eigentliche Arbeit konzentrieren, während der Workflow im Hintergrund für konsistente Deployments sorgt.

### Rollback-Fähigkeit

Durch die klare Versionierung ist es jederzeit möglich, zu einer vorherigen Version zurückzukehren, falls Probleme auftreten.

## Fazit

Dieser Workflow hat sich in meiner Blog-Entwicklung als äußerst effizient erwiesen. Die Kombination aus strukturierten Git-Branches, automatisierter semantischer Versionierung und GitOps-Prinzipien schafft einen robusten, nachvollziehbaren und wartungsarmen Entwicklungsprozess.

Der Schlüssel liegt in der Konsistenz: Jeder Schritt folgt klaren Regeln, jede Version ist eindeutig identifizierbar, und jede Umgebung spiegelt exakt den gewünschten Zustand wider. So kann ich mich auf das Wesentliche konzentrieren - das Erstellen großartiger Inhalte für meinen Blog.

