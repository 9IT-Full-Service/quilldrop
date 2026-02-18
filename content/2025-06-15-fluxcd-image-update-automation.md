---
title: 'Vereinfachte GitOps-Pipeline mit FluxCD und automatischen Image-Updates' 
date: 2025-06-15 10:00:00
xupdate: 2025-06-15 13:51:00
update: 2025-06-15 14:51:00
author: ruediger
cover: "/images/posts/2025/06/fluxcd-image-update-automation.webp"
featureImage: "/images/posts/2025/06/fluxcd-image-update-automation.webp"
tags: [Kubernetes, FluxCD, ImageUpdateAutomation, GitOps]
categories: 
  - Kubernetes
preview: "Die Lösung liegt in der Trennung von Build- und Deployment-Prozessen durch eine GitOps-Architektur. Anstatt alles in einer monolithischen Pipeline zu erledigen, beschränke ich die Build-Pipeline auf das Wesentliche: semantische Versionierung, Docker Image-Erstellung und Push in die Registry. Das Deployment wird komplett an FluxCD delegiert, das kontinuierlich das GitOps-Repository überwacht und automatisch neue Image-Versionen erkennt und ausrollt."
series: ["FluxCD"] 
draft: false
top: false
type: post
hide: false
toc: false
---

![FluxCD Image Update Automation](/images/posts/2025/06/fluxcd-image-update-automation.webp)

Traditionelle Build-Pipelines sind oft komplex und fehleranfällig. Sie müssen nicht nur den Anwendungscode verarbeiten, sondern auch Helm Charts verwalten, semantische Versionierung durchführen und das finale Deployment orchestrieren. Wenn dabei ein Schritt fehlschlägt – beispielsweise ein fehlerhafter Helm Chart – bricht die gesamte Pipeline ab, obwohl das Docker Image bereits erfolgreich erstellt wurde. Besonders bei größeren Projekten mit längeren Build-Zeiten ist das frustrierend und ineffizient.

Die Lösung liegt in der Trennung von Build- und Deployment-Prozessen durch eine GitOps-Architektur. Anstatt alles in einer monolithischen Pipeline zu erledigen, beschränke ich die Build-Pipeline auf das Wesentliche: semantische Versionierung, Docker Image-Erstellung und Push in die Registry. Das Deployment wird komplett an FluxCD delegiert, das kontinuierlich das GitOps-Repository überwacht und automatisch neue Image-Versionen erkennt und ausrollt.

Diese Architektur bietet mehrere Vorteile: Build-Pipelines werden einfacher und stabiler, Deployments erfolgen vollautomatisch ohne manuelle Eingriffe, und durch die stage-spezifische Konfiguration können verschiedene Umgebungen (dev, stage, prod) mit unterschiedlichen Image-Tags versorgt werden. FluxCD übernimmt dabei nicht nur das Deployment, sondern auch die automatische Aktualisierung der Image-Tags im GitOps-Repository.

Build Pipeline sehen oft wie folgt aus: 

![Build Pipeline sehen oft wie folgt aus](/images/posts/2025/06/build-pipeline.png)

Diese Piplelines müssen nicht nur den Code selbst auschecken, sie müssen zusätzlich auch noch den HelmChart auschecken könne. Liegen diese wo anders, benötigt man hier wieder einen Token für Pull und Push. 
Tritt beim Pull oder des Push des HelmCharts ein Problem auf, oder ist der HelmChart fehlerhaft, dann bricht die komplette Pipeline ab. Gerade bei großen Projekten, bei denen ein Build etwas länger dauert ist das sehr ärgerlich. Da das Image eigentlich schon fertig für das Deployment ist, aber ein Fehler im Chart das verhindert. 

Meine Pipeline macht nur noch SemVer und erstellt das Docker Image und push es in die Docker Registry. Sonst wird da nicht mehr gemacht. Der Rest wird an anderer Stelle erledigt. 

Ich installiere meine Kubernetes Cluster mit Cloudinit und mache so direkt ein `flux bootstrap`, damit alles in den Clustern installiert wird. 

```bash
# ...
# other stuff: install kubernete, helm, k9s ....
# ...
# FluxCD Bootstrap
export GITHUB_TOKEN="${github_token}"
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml

# Warten bis Kubernetes API verfügbar ist
while ! kubectl cluster-info &>/dev/null; do
	echo "Waiting for Kubernetes API..."
	sleep 10
done

# FluxCD Bootstrap ausführen
flux bootstrap github \
--owner=ruedigerp \
--repository=fluxcd \
--branch=main \
--path=clusters/production \
--personal \
--components-extra=image-reflector-controller,image-automation-controller
```

Im Repo `gitops` ist für jeden Cluster ein Verzeichnis für die jeweile Konfigurationen. Damit werden Helm Chart Repos, Helm Charts, Applications und weitere Stage spezifische Konfigurationen gemacht. 

Im repo ist auch ein Verzeichnis `/apps/, in dem sind Applications die installiert werden sollen. Wie zum Beispiel mein Blog. 

```bash
❯ tree apps/blog
apps/blog
├── base
│   ├── deployment.yaml
│   ├── ingress-http.yaml
│   ├── ingress-https.yaml
│   ├── kustomization.yaml
│   └── service.yaml
├── dev
│   ├── deployment-patch.yaml
│   └── kustomization.yaml
├── prod
│   ├── deployment-patch.yaml
│   └── kustomization.yaml
└── stage
    ├── deployment-patch.yaml
    └── kustomization.yaml
```

Im Base Verzeichnis sind alle Kubernetes YAML-Files die für die Application benötigt werden. Das Deployment enthält wie üblich ein Docker Image:
```bash
grep image apps/blog/base/deployment.yaml
        image: ghcr.io/ruedigerp/ink-blog.kuepper.nrw:v0.0.1-develop.1
```

Hier ist es egal welche Version eingetragen ist, da der Image Name und Tag je nach Stage ersetzt werden. 

Denn im Apps Verzeichnis sind für jeden Cluster/Stage ein weiteres Verzeichnis, welches dann die Konfigurationen enthält, die auf dem Ziel Cluster genutzt werden sollen. 
Wie auch das Patch für die Images.

```bash
❯ grep image apps/blog/prod/deployment-patch.yaml
        image: ghcr.io/ruedigerp/ink-blog.kuepper.nrw:v1.9.31 # {"$imagepolicy": "blog:blog-policy"}
❯ grep image apps/blog/stage/deployment-patch.yaml
        image: ghcr.io/ruedigerp/ink-blog.kuepper.nrw:v1.9.32-stage.1 # {"$imagepolicy": "blog:blog-policy"}
❯ grep image apps/blog/dev/deployment-patch.yaml
        image: ghcr.io/ruedigerp/ink-blog.kuepper.nrw:v1.9.32-develop.2 # {"$imagepolicy": "blog:blog-policy"}
```

Aufgerufen werden sie durch die jeweiligen kustomization.yaml Files in den Cluster/Stage Verzeichnissen. Hier das Beispiel von `prod`: 

```yaml
❯ cat apps/blog/prod/kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - ../base

patches:
  - path: deployment-patch.yaml

namespace: blog
```

Das sagt Flux an der Stelle es soll alles aus dem Verzeichnis `../base` anwenden, und das Patch-File für das Deployment anwenden. Hier können auch noch weitere Patches für alles mögliche, wie beispielsweise Änderungen am Ingress, Service, Secrets usw. 

Der Patch für das Deployment sieht wie folgt aus:

```yaml
❯ cat apps/blog/prod/deployment-patch.yaml
# apps/blog/stage/deployment-patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: blog
  namespace: blog
spec:
  template:
    spec:
      containers:
      - name: blog
        image: ghcr.io/ruedigerp/ink-blog.kuepper.nrw:v1.9.31 # {"$imagepolicy": "blog:blog-policy"}
```

Der Patch nimmt das Deployment `blog` im Namespace `blog` und ersetzt das image welches in `base/deployment.yaml` gesetzt ist. 

Wichtig ist in der `deployment.yaml` und `deployment-patch-yaml` der Kommentar hinter dem Image:  `# {"$imagepolicy": "blog:blog-policy"}`.

Damit weiß ImageUpdateAutomoation von FluxCD welches Image er ersetzten soll und mit welcher Policy. In diesem Fall `blog-policy` im Namespace `blog`. 

```yaml
# clusters/production/flux-system/kustomizations/private/blog.yaml
apiVersion: kustomize.toolkit.fluxcd.io/v1
kind: Kustomization
metadata:
  name: blog
  namespace: flux-system
spec:
  dependsOn:
    - name: infrastructure
  interval: 5m
  prune: false
  path: ./apps/blog/prod
  sourceRef:
    kind: GitRepository
    name: flux-system
  targetNamespace: blog
  decryption:
    provider: sops
    secretRef:
      name: sops-age
# weitere Sachen wie Patch der Replica, ingress Domains usw. 
# ... 
```

Über weitere Kustomizations lade ich im Verzeichnis `clusters/production` noch viele weitere Sachen wie Helm Repos, Helm Charts, weitere Apps und auch die Konfigurationen der imageUpdates. 

```bash
tree clusters/production/infrastructure/image-updater/
clusters/production/infrastructure/image-updater/
├── blog
│   ├── image-policy.yaml
│   ├── image-repository.yaml
│   ├── image-update-automation.yaml
│   └── kustomization.yaml
├── image-write-repo.yaml
├── kustomization.yaml
├── other-app
│   ├── image-policy.yaml
│   ├── image-repository.yaml
│   ├── image-update-automation.yaml
│   └── kustomization.yaml
└── second-other-app
    ├── image-policy.yaml
    ├── image-repository.yaml
    ├── image-update-automation.yaml
    └── kustomization.yaml
```

Die Image-Update-Automation benötigt für das prüfen auf neue Docker Images, das prüfen ob der Image-Tag für die Stage passt und die Automation selbst drei Konfigurationen. 
Das Image Repo konfigurieren: 

```yaml
# clusters/production/infrastructure/image-updater/blog/image-repository.yaml
apiVersion: image.toolkit.fluxcd.io/v1beta2
kind: ImageRepository
metadata:
  name: blog-repo
  namespace: blog
spec:
  image: ghcr.io/ruedigerp/ink-blog.kuepper.nrw
  interval: 10m
  provider: generic
```

Damit wird alle 10 Minuten geprüft, ob es ein neues Image gibt. Und die Policy kann da überprüfen welcher Tag auf die entsprechende Stage passt. 

```yaml
# clusters/production/infrastructure/image-updater/blog/image-policy.yaml
apiVersion: image.toolkit.fluxcd.io/v1beta2
kind: ImagePolicy
metadata:
  name: blog-policy
  namespace: blog  # Gleicher Namespace
spec:
  imageRepositoryRef:
    name: blog-repo  # Muss mit ImageRepository.name übereinstimmen
  filterTags:
    pattern: '^v[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$'
  policy:
    semver:
      range: '>=1.0.0'    
```

Der `filterTag` ist entsprechend der Stages jeweils angepasst. Dafür liegt entsprechend des Clusters, bzw. der Stage, in `/cluster/{prod,stage,dev}/infrastructure/image-updater/blog/image-policy.yaml` die passende Policy. 

Stage: 
```yaml
...
  filterTags:
    pattern: '^v[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$'
  policy:
    semver:
      range: '>=1.0.0-stage' 

...
```

Dev:
```yaml
  filterTags:
    pattern: '^v[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$'
  policy:
    semver:
      range: '>=1.0.0-dev' 
```

Die Dateien sind ansonsten identisch. 

Wenn es einen neuen Tag gibt, wird diese Änderung mit der `ImageUpdateAutomation` vorgenommen und ins GitOps Repo commited. 

```yaml
# clusters/production/infrastructure/image-updater/blog/image-update-automation.yaml
apiVersion: image.toolkit.fluxcd.io/v1beta2
kind: ImageUpdateAutomation
metadata:
  name: blog-automation
  namespace: blog  # Gleicher Namespace
spec:
  sourceRef:
    kind: GitRepository
    name: flux-system-write
    namespace: flux-system
  git:
    checkout:
      ref:
        branch: main
    commit:
      author:
        email: fluxcdbot@users.noreply.github.com
        name: fluxcdbot
      messageTemplate: |
        Automated image update

        Automation name: {{ .AutomationObject }}

        Images:
        {{- range .Updated.Images }}
        - {{.}}
        {{- end }}
    push:
      branch: main
  interval: 1m
  update:
    path: "./apps/blog/prod"  # Anpassen an deinen Pfad
    strategy: Setters
```

Damit wird dann in `./apps/blog/prod` das image im `deployment-patch.yaml` geändert und ins Git Repo gepusched. Der Patch ändert dann beim `reconcile` der Kustomization das Deployment und applied es in den Cluster. 

Ich muss daher keine Änderungen im GitOps Repo machen. Das passiert alles automatisch. 
Wenn man nach einem Docker Image Build nicht automatisch ausrollen möchte, kann man auch über einen PR die Änderungen machen lassen. So kann man neue Versionen erst ausrollen wenn man es möchte und hat mehr Kontrolle. 


