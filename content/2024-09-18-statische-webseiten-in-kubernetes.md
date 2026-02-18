---
title: 'Statische Webseiten in Kubernetes'
date: 2024-09-18 18:00:00
author: ruediger
cover: "/images/posts/2024/02/cms.webp"
tags: [Homepage, CMS, Blog, Kubernetes]
categories: 
    - Internet
preview: "Statische Webseiten in Kubernetes."
draft: false
top: false
type: post
hide: false
toc: false
---

![Blog and CMS](/images/posts/2024/02/cms.webp)

Ich habe mehrere Seiten und benutze dafür Hugo, InkProject und ein eigenes in Go geschriebenes CMS. 
Der Content wird in allen drei mit Markdown Files erstellt. Mit den drei Systemen wird dann die jeweilige Seite statisch generiert. 

## Github/Gitlab Pipelines? 

Benutze ich Github oder Gitlab Pipelines für das generieren? Nein.

Der Content, die CMS Config, Themes usw. liegen alle in Github. Aber Github Actions oder Pipelines benutze ich bei den Seiten nicht mehr. Das habe ich über Jahre so gemacht, aber der Aufwand ist mir einfach zu groß. Und das veröffentlichen, also das generieren der Seiten, erstellen der Images und Deployment der Images in Kubernetes dauert einfach zu lange. 

### Aufwand und Zeit.

Damit meine ich das Hugo, InkProject und mein eigens CMS bei aktualisieren neu zu kompilieren, Docker Images zu bauen und so bei Updates der Systeme mit der neuen Version die Seiten erstellen zu können. 

Auch das generieren der Seiten und das anschliessende erstellen der Images dauert einfach sehr lange. Da jeder Step ein Docker Image hoch fährt, seine Aufgabe erledigt und anschliessende die fertigen Daten weiter verarbeitet. Images, da ich nur noch Multi-Arch Images baue, in die Registry pushen und mit dem fertigen Image den Helm Chart mit der neuen Image Version erstellt und deployed.

Das dauerte ja nach Seite 3-5 Minuten. Was bei kleinen Änderungen wie Typos einfach nervig war. 

## Wie ich jetzt Seiten aktualisiere

Jede Seite ist lokal auf dem Rechner und ich füge neue Markdown Files für neuen Content dazu. Ich kann bei allen drei Systemen einfach `<CMSNAME> build` oder `<CMSNAME serve` benutzen. So kann ich einfach aus dem Ordner `public` die Setie aufrufen oder mit dem `serve` die Seite über `http://localhost:8000` aufrufen. So kann man schnell die fertige Seiten oder Posts angucken. 

In jedem Ordner ist ein Script `build.sh`, welches dann die Seite generiert und die Docker Images erstellt und dann ein Multi-Arch Manifest generiert. 
Dabei wird die Version für die Image-Tags immer hochgezählt.

Das Image kann so auch direkt im Cluster aktualisiert werden, wird aber nur bei Tests oder schnellen Änderungen genutzt. 

Das Rollout wird aber über den Helm-Chart gemacht. 

```
#!/bin/bash

CMS="ink"
# CMS="hugo"
# CMS="kube-cms"
REGISTRY="ghcr.io"
USERNAME="youruser"
IMAGE="myblog"

${CMS} build

perl -pe 's/(version: )(v\d\.)(\d\.)(\d+)/$1 . $2.$3.($4 + 1)/ge' docker.yaml > docker.yaml.tmp; 
mv docker.yaml.tmp docker.yaml
VERSION=$(perl -pe 's/(version: )(v\d\.)(\d\.)(\d+)/$2.$3.($4)/ge' docker.yaml)


docker buildx build --no-cache -t ${REGISTRY}/${USERNAME}/${IMAGE}:${VERSION}-arm64 .
docker buildx build --no-cache --platform linux/arm64 -t ${REGISTRY}/${USERNAME}/${IMAGE}:${VERSION}-arm64-linux .
docker buildx build --no-cache --platform linux/amd64 -t ${REGISTRY}/${USERNAME}/${IMAGE}:${VERSION}-amd64 .

docker push ${REGISTRY}/${USERNAME}/${IMAGE}:${VERSION}-arm64
docker push ${REGISTRY}/${USERNAME}/${IMAGE}:${VERSION}-arm64-linux
docker push ${REGISTRY}/${USERNAME}/${IMAGE}:${VERSION}-amd64

docker manifest create ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION \
    --amend ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION-amd64 \
    --amend ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION-arm64-linux \
    --amend ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION-arm64
docker manifest push ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION
docker manifest create ${REGISTRY}/${USERNAME}/${IMAGE}:latest \
    --amend ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION-amd64 \
    --amend ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION-arm64-linux \
    --amend ${REGISTRY}/${USERNAME}/${IMAGE}:$VERSION-arm64
docker manifest push ${REGISTRY}/${USERNAME}/${IMAGE}:latest
```

## Helm-Chart

Der Helm-Chart kümmert sich um das Deployment in den bzw. die Kubernetes Cluster. Das Deployment, Service, Ingress, TLS-Cert, Traefik Middlewares und vielem mehr. Änderungen an den Workloads können so schnell gemacht werden und Rollback ist schnell möglich. 

Auch hier wird, wie beim Docker Image, die Version immer hoch gezählt. Der aktuelle Image-Tag wird aus dem Verzeichnis des CMS ausgelesen und als Value beim Installaieren bzw. beim Upgrade des Helm-Charts im Cluster übergeben. 

Damit ich beim Rollout des Helm-Charts nicht zusätzlich immer erst in das CMS Verzeichnis wechseln muss, um ein neues Image bei Änderungen zu generieren, kann ich einfach beim Aufrufe vom Script `package.sh` einfach ein `docker` mit anhängen. 

So bin ich flexible und kann auch nur Änderungen am Helm-Chart ausrollen, ohne immer wieder neue Docker Images zu erstellen. 

```
#!/bin/bash

HOMEPAGEURL="blog.example.net"
HELMREPO="charts.mydomain.de"
CHARTNAME="myblog"
CHARTUSER="bob"
CHARTPASSWORD="mysecret"

if [ "$1" = "docker" ];
then
    cd ../../services/${HOMEPAGEURL}/
    ./build.sh
    cd -
fi


perl -pe 's/(version: )(v\d\.)(\d\.)(\d+)/$1 . $2.$3.($4 + 1)/ge' helm/Chart.yaml > helm/Chart.yaml.tmp;
mv helm/Chart.yaml.tmp helm/Chart.yaml

TAG=$(grep -a -m1 'version: ' ../../services/${HOMEPAGEURL}/docker.yaml | cut -d " " -f2)

helm package helm
helm cm-push -u ${CHARTUSER} -p ${CHARTPASSWORD} $(ls -tr -1 *.tgz | tail -n 1) ${HELMREPO}

helm repo update
helm upgrade ${CHARTNAME} ${HELMREPO}/${CHARTNAME} --namespace ${CHARTNAME} --set tag=${TAG}
```
