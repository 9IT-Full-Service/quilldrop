---
title: 'Automatisierte Kubernetes Volume-Backups'
date: 2025-06-18 12:00:00
update: 2025-06-18 12:00:00
author: ruediger
cover: "/images/posts/2025/06/k8s-pv-backup.webp"
featureImage: "/images/posts/2025/06/k8s-pv-backup.webp"
tags: [Kubernetes, Volumes, Backup]
categories: 
    - Kubernetes
preview: ""
draft: false
top: false
type: post
hide: false
toc: true
---

![Automatisierte Kubernetes Volume-Backups](/images/posts/2025/06/k8s-pv-backup.webp)

Ich betreibe meine Kubernetes Cluster bei Hetzner und habe in den Projekten zusätzlichen einen NFS-Server, der in Kubernetes die Volumes mit dem `nfs-subdir-external-provisioner` bereitstellt. 

Das gleiche Setup habe ich auch zuhause in meheren VMs. So das ich auch dort das gleiche Setup habe und lokale Services oder auch zum testen vorhanden habe wie im Live-Betrieb. 

Ich kann so zwischen Dev, Stage und Prod Daten bei Hetzner schnell hin und her kopieren. Oder über VPN auch mal die Daten nach Hause ziehen um Migrationen oder Änderungen zu testen. 

Daten können dann so auch schnell wieder hergestellt werden. 

Durch dieses Setup kann ich auch jeden Cluster einfach neu aufsetzen, mit FluxCD alle Kubernetes Resourcen installieren lassen und die Daten kommen von den NFS-Servern. 

Diese Daten kann ich an andere Orte synchronisieren und habe damit auch das Backup. 
Trotzdem sichere ich die Daten für die jeweiligen Custer noch auf den jeweiligen Kubernetes Clustern noch einmal auf den NFS-Servern. Das Backup ist für schnelles Recovery und dann im Einsatz wenn etwas getestet oder aktualisiert wird. Geht einmal etwas schief, kann ich die Daten direkt zurück spielen. 

# Der Backup Job

Das mache ich mit einem Backup Cronjob den ich vorher einmal manuell triggern kann. Er sichert dann alle PV Daten in einem Backup-Job. 

Da ich NFS mit dem  `nfs-subdir-external-provisioner` nutze und die PV/PVC mit dem AccessMode ReadWriteMany eingerichtet sind  können diese Volumes einfach vom Backup-Job gemounted werden. 

# Der Backup Job mit ReadWriteOnce

Das Problem sind immer Volumes mit dem AccessMode ReadWriteOnce. Das Volume kann nicht von anderen Pods gemounted werden. 

Um diese Volumes zu sichern muss 

* Der Pod der Application über das Deployment, Statefulset oder Daemonset einmal herunter skaliert werden.
* Ein Pod gestartet werden, der das Volume und das Backupziel mounted.
* Das Backup erstellt werden.
* Backup Pod stoppen.
* Deployment, Statefulset oder Daemonset wieder hoch skalieren.

Das kann man für 1-2 Deyployments manuell machen. Aber wenn im Cluster viele Services laufen wird diese Aufgabe recht langwierig. Das möchte man weg automatisieren. 

# Mein k8s-pv-backup Tool

Ich möchte mich um das Backup nicht kümmern. Auch möchte ich nicht nach jeder Installation von neuen Services auch noch die Backup Config anpassen müssen. Genau das gleiche, wenn ein Service nicht mehr benutzt wird und vom Cluster entfernt wurde. 

k8s-pv-backup läuft als Cronjob im Cluster und prüft als erstes alle PV im Cluster und sucht dazu das passende PVC. Dabei wird auch noch der AccessMode geprüft. 

Anhand des AccessModes wird entschieden ob ein Deployment, Statefulset oder Daemonset auf 0 skaliert werden muss oder nicht (ReadWriteMany).

Dann wird erst einmal RBAC entsprechend konfiguriert und ein PVC in dem Namespace des zu sichernden Service erstellt, mit dem PV der Application. Jetzt wird noch ein Job für das eigentliche Backup angelegt, der dann mit dem neuen PVC und dem Backup-NFS PVC das Backup erstellt. 

Ist der Job fertig wird für diesen einen Service das RBAC und PVC wieder gelöscht und der Job für den nächste Service wird angelegt. RBAC/PVC erstellen, Backup erstellen und aufräumen. 

Beim herunter und rauf skalieren der Deployment, Statefulset oder Daemonset wird immer auf den Wert gesetzt der vorher gesetzt gewesen ist. 

So hat man komplett automatisiert ein Backup. Zero-Downtime bei ReadWriteMany und kurze Downtime beim ReadWriteOnce. 

# k8s-pv-backup benutzen

Installation einfach per Helm. Es kann aber auch lokal ausgeführt werden. 


```bash
helm repo add ruedigerp https://ruedigerp.github.io/helm-charts/
helm repo update ruedigerp 

helm upgrade --install k8s-pv-backup ruedigerp/k8s-pv-backup --namespace backup-system --create-namespace --wait -f values.yaml
```

Beispiel values.yaml 

```yaml
nfs_server: "10.0.10.7"
nfs_path: "/srv/nfs/k8s-pv/production/k8s-backup"
storage_class: "nfs-client"
```

`nfs_server` und `nfs_path` entsprechend anpassen. 
`storage_class` ist aktuell für "nfs-client" ausgelegt. Das wird sich in Zukunft aber noch ändern, da natürlich auch andere Backupziele möglich sein sollen. 



