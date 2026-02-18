---
title: Kubernetes, no default StorageClass
date: 2023-09-11 12:15:00
update: 2023-09-11 12:15:00
cover: "/images/cat/technik.webp"
author: ruediger
draft: false
top: false
tags:
  - Kubernetes
  - StorageClass
categories: 
    - Internet
preview: "Kubernetes, no default StorageClass"
type: post
hide: false
toc: false
---

Beim Anlegen eines PVC, ohne Angabe einer StorageClass, wird das PV nicht angelegt, wenn im Cluster keine StorageClass als `default` ausgewählt ist. Das PVC steht anschliesend dauerhaft auf `pending`.

Mit `kubectl` die Liste der verfügbaren StorageClass abrufen und aus der Liste die gewünschte Storageclass heraus suchen. 

> When creating a PVC without specifying a StorageClass, the PV is not created if no StorageClass is selected as default in the cluster. The PVC then remains in a pending state indefinitely.

> Retrieve the list of available StorageClass using kubectl and search for the desired StorageClass from the list.

```
kubectl get storageclass
NAME                 PROVISIONER                     RECLAIMPOLICY   VOLUMEBINDINGMODE      ALLOWVOLUMEEXPANSION   AGE
csi-disk             everest-csi-provisioner         Delete          Immediate              true                   66d
csi-disk-topology    everest-csi-provisioner         Delete          WaitForFirstConsumer   true                   66d
...
```

Mit dem Namen der StorageClass ein Kubectl Path ausführen und dabei dann `is-default-class` auf `true` setzen.

> Execute a kubectl path with the name of the StorageClass and set is-default-class to true.

```
kubectl patch storageclass csi-disk -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
storageclass.storage.k8s.io/csi-disk patched
```


Die StorageClass wird anschliessend als Default StorageClass angezeigt. 

> The StorageClass will then be displayed as the default StorageClass.


```
kubectl get storageclass
NAME                 PROVISIONER                     RECLAIMPOLICY   VOLUMEBINDINGMODE      ALLOWVOLUMEEXPANSION   AGE
csi-disk (default)   everest-csi-provisioner         Delete          Immediate              true                   66d
csi-disk-topology    everest-csi-provisioner         Delete          WaitForFirstConsumer   true                   66d
...
```