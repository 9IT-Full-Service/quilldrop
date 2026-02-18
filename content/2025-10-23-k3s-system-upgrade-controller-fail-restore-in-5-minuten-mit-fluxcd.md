---
title: 'k3s System-Upgrade-Controller Fail, Restore in 5 Minuten mit FluxCD'
date: 2025-10-23 15:00:00
update: 2025-10-23 15:00:00
author: ruediger
cover: "/images/posts/2025/10/k3s-fluxcd-k3s.webp"
# images: 
#   - /images/posts/2025/08/telekom-mail-fail.webp
featureImage: /images/posts/2025/10/k3s-fluxcd-k3s.webp
tags: [Kubernetes, k3s, fluxcd, Recovery, GitOps]
categories: 
  - Kubernetes
preview: "Wenn der k3s System Upgrade Controller eine Recovery-Übung erzwingt: So wie vorletzte Woche, als ich eigentlich einfach nur ein Upgrade von k3s auf die aktuellste Version machen wollte. Nur ist das irgendwie etwas schiefgelaufen. Alle Services liefen noch, die Container waren alle noch da und alles war erreichbar. Das Einzige, was nicht gestartet ist, war k3s."
draft: false
top: false
type: post
hide: false
toc: false
---


Wenn der k3s System Upgrade Controller eine Recovery-Übung erzwingt: So wie vorletzte Woche, als ich eigentlich einfach nur ein Upgrade von k3s auf die aktuellste Version machen wollte.
Nur ist das irgendwie etwas schiefgelaufen. Alle Services liefen noch, die Container waren alle noch da und alles war erreichbar. Das Einzige, was nicht gestartet ist, war k3s.

Kurzes Checken der Logfiles zeigte, dass die Node-IPs nicht zu denen in der Cluster-Config passten. Es wurde versucht, mit der Loadbalancer-IP zu verbinden. Ändern in der Datenbank wollte irgendwie nicht funktionieren. Loadbalancer-IP kurz getrennt, sodass nur noch die Node-IP auf den einzelnen Nodes zu sehen war. Auch damit wollte k3s nicht starten.

Also den System-Upgrade-Controller einfach mal zurückrollen lassen. Was dann auch geklappt hat. Alles wieder ok, k3s wieder gestartet und nutzbar. Nur waren im Log noch ein paar Sachen, die etwas stutzig machten. Nichts Kritisches, aber es waren selbst noch diese Einträge zu sehen, nachdem ich eine der Applications, für die Fehler angezeigt wurden, auf 0 skaliert und sogar komplett entfernt hatte.

Der Cluster lief danach noch ein paar Tage ohne Probleme. Und da ich mit FluxCD jetzt alles in einem GitOps-GitHub-Repository habe, war eh schon länger geplant, auch meinen privaten Production-Kubernetes-Cluster neu aufzusetzen. Das ist mit den Clustern labor01, dev und stage schon mehrfach gemacht worden. Der Prod-Cluster war zwar schon die ganze Zeit mit im Repo und alle Komponenten migriert, aber der Teufel liegt ja immer im Detail. Habe ich wirklich an alles gedacht und ist wirklich alles im Git?

Außerdem: Was ist mit den Daten, die nicht auf dem NFS-Server liegen? Die Daten der Volumes liegen hauptsächlich auf NFS-Shares, außer dem vom Garage S3 Storage. Dafür habe ich dann auch gleich noch ein Backup- und Restore-Tool erstellt. Dazu die Tage noch mehr.

Nachdem alle Daten aus dem Garage auch noch gesichert wurden, war die Zeit für das Löschen des Kubernetes-Clusters gekommen. Gefolgt von Server Create plus Cloud-Init. Damit wurde alles wieder erstellt. Das Cloud-Init installiert dann alles, was benötigt wird: k3s, helm, k9s, FluxCD.

FluxCD wird dann per Bootstrap automatisch mit dem GitOps-Repo verbunden und installiert dann alle Ressourcen. Nach etwa 3-4 Minuten ist der Cluster wieder komplett gestartet und alles wieder installiert und aktiv. Während der Cert-Manager mit Cloudflare DNS API sich kurz noch um die Zertifikate gekümmert hat, habe ich mich kurz Garage gewidmet. Den Garage-INIT-Job getriggert, der kümmert sich darum, dass sich alle Garage-Nodes gegenseitig kennen. Anschließend das Restore für Garage ausgeführt und alle Buckets mit Daten und die Secrets waren wieder da.

Das Ganze war schon sehr gut und war genau so, wie es sein soll. Und ja, es gab 2-3 Kleinigkeiten, die waren noch nicht im Git. Die wurden dann auch gleich noch erledigt. So wird dann bei der nächsten Installation direkt alles fertig sein.

Ich kann auch den Cluster aufsetzen, aber beim ersten Start werden die Zertifikate noch gar nicht generiert. So kann ich erst einmal alles prüfen und ein paar Tests machen. Wenn alles ok ist, wird der Cert-Manager aktiv geschaltet und das Loadbalancer-Target auf den neuen geändert. So ist der Cluster innerhalb ein paar Sekunden ausgetauscht.

Dass der System-Upgrade-Controller fehlgeschlagen ist, war ärgerlich, aber das war auch irgendwie gut so. Eine Recovery war schon länger geplant, da es bei Test-Clustern im HomeLab und auch für die anderen Stages immer gut lief. Es sind immer ein paar Sachen aufgefallen und behoben worden. Da habe ich aber auch eine Neuinstallation einfach mal so zwischendurch gemacht. Beim Production-Cluster war immer etwas Respekt davor. Aber jetzt musste es einfach gemacht werden. Es lief zwar alles, aber es waren Fehler zu sehen, bei denen ich nicht wusste, woher sie kommen und wieso. Was macht das sonst noch im Hintergrund – könnte es irgendwann doch den ganzen Cluster einfach zerreißen und alles wäre offline?