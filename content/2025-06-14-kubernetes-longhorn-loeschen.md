---
title: 'FIX: Kubernetes Longhorn löschen hängt in Terminating fest'
date: 2025-06-14 15:00:00
author: ruediger
cover: "/images/posts/2025/06/fix-problem.webp"
featureImage: "/images/posts/2025/06/fix-problem.webp"
tags: [Kubernetes, Longhorn, uninstall]
categories: 
    - Kubernetes
    - Login
preview: "Longhorn getestet, gelöscht und Namespace hängt im Status Terminating."
draft: false
top: false
type: post
hide: false
toc: false
---

![Longhorn Namespace und CRDs unlöschbar](/images/posts/2025/06/fix-problem.webp)

### Longhorn-Deinstallationsproblem: Persistierende CRDs und Namespace

Bei einem Test von Longhorn als Kubernetes-Storage-Lösung traten Performance-Probleme auf, die zu einer erheblichen Cluster-Verlangsamung führten. Nach der Entscheidung zur Deinstallation wurden zwar Pods und Deployments erfolgreich entfernt, jedoch blieben Custom Resource Definitions (CRDs) und der zugehörige Namespace bestehen.

### Problemanalyse:

Namespace `longhorn` verbleibt im Status `Terminating`
CRDs lassen sich nicht über Standard-Löschbefehle entfernen
Namespace-Löschung wird durch bestehende CRDs blockiert

### Ursache:

Longhorn-CRDs enthalten Finalizers, die eine automatische Bereinigung verhindern. Dies ist ein dokumentiertes Problem bei der Longhorn-Deinstallation.

### Lösungsansatz:

Entfernung der Finalizers aus allen Longhorn-CRDs
Forcierte Löschung der CRDs
Bereinigung des Namespace

### Technische Umsetzung:

```bash
kubectl get crd | grep longhorn | awk '{print $1}' | xargs -I {} kubectl patch crd {} -p '{"metadata":{"finalizers":[]}}' --type=merge
customresourcedefinition.apiextensions.k8s.io/backups.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/engineimages.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/engines.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/nodes.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/replicas.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/sharemanagers.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/snapshots.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/volumeattachments.longhorn.io patched
customresourcedefinition.apiextensions.k8s.io/volumes.longhorn.io patched
```

Und jetzt noch den Namespace `longhorn`

```bash
kubectl patch namespace longhorn -p '{"metadata":{"finalizers":[]}}' --type=merge
namespace/foo patched

kubectl delete ns longhorn 

❯ kubectl get namespace | grep longhorn
❯

```

Diese Vorgehensweise löst das persistierende Deinstallationsproblem und ermöglicht eine vollständige Bereinigung der Longhorn-Komponenten.

Der Longhorn Namespace ist Geschichte. 


