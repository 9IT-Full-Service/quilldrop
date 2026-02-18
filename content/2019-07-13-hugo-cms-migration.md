---
title: 'Hugo CMS migration'
date: 2019-07-13 17:03:03
update: 2019-07-13 17:03:03
author: ruediger
cover: "/images/posts/2019/08/04/programming.webp"
tags:
    - Webpage
    - CMS
    - Hugo
    - Generator
    - Pipeline
    - Deployment
    - Automatisierung
preview: "*Notiz an mich, um nich noch einmal suchen zu müssen*  Um im BGP manche Netze nicht zu erlauben:"
categories: 
    - Technik
toc: false
hide: false
type: post
---


## Migration nach Hugo CMS

Die Seite ist jetzt zu einem Hugo CMS migriert worden. Hugo ist ein Static Page Generator.
Im gegensatz zu Wordpress werden die Seiten nicht bei jedem Aufruf neu generiert, sondern nur nach Änderungen.
Die fertig Seiten werden dann auf dem Server bereitgestellt.

Das macht die Seite sehr schell und ich kann sie in Zukunft auch auf sehr vielen Servern, in Docker oder Kubernetes Clustern verteilen.
Also auch sehr grosse Lastspitzen locker abfangen.
<!--more-->
## Hugo - CMS

> Als statischer Websitegenerator werden von Hugo die HTML-Dateien – im Gegensatz zu dynamischen Websitegeneratoren – nicht jedes Mal, wenn die Webseite aufgerufen wird, neu generiert, sondern nur, wenn sich der Inhalt der jeweiligen Seite ändert. Insbesondere ermöglicht es Hugo, dass nur diejenigen HTML-Dokumente der jeweiligen Webseite neu gebaut werden müssen, in denen Änderungen auftraten. Hierdurch sollen die Ressourcen des Servers geschont und eine hohe Effizienz von diesem erreicht werden.[3] Nach einer nicht-repräsentativen Benchmark generiert Hugo Webseiten 75-mal schneller als der ebenfalls statische Websitegenerator Middleman.[4]

> Hugo unterstützt nativ neben HTML auch die Darstellung von Texten, die in Markdown verfasst wurden. Mit Hilfe externer Anwendungen kann diese Unterstützung auf AsciiDoc und reStructuredText erweitert werden. Auch YAML, JSON und TOML werden unterstützt. Mittels der sogenannten „LiveReload“-Funktion können Änderungen an den Dokumenten zeitgleich auf der Webseite übernommen werden. Die graphische Darstellung der Inhalte kann mittels verschiedener Themenvorlagen geregelt werden.[4][5] Dabei wird zwischen drei verschiedenen Grundtypen unterschieden: Single, List und Homepage. Die Nutzung der Themen erfolgt mittels der Template-Engine von Go. Hugo ermöglicht es zusätzlich, Inhalte der Webseiten mittels Schlüsselwörtern zu kategorisieren.

> Eine Besonderheit von Hugo ist, dass es einen eigenen HTTP-Server mitliefert. Hierdurch sind Anwender nicht auf z. B. nginx oder den Apache HTTP Server angewiesen, wodurch Abhängigkeiten verhindert werden. Auch bestimmte Laufzeitumgebungen und Datenbanken wie Ruby, PHP oder MySQL werden zur Nutzung nicht benötigt.

## Format von manchen Blog Artikeln noch defekt

In manchen Artikeln wurden Plugins für Tabellen, Soundcloud einbindung und andere benutzt.
Daher können manche der Artikel aktuell noch nicht richtig angezeigt werden.
Bei den Artikeln muss ich noch etwas nacharbeiten.

## Seite generieren auf GitLab

Beschreibung wie die Seite mit Gitlab Pipeline generiert wird ...

Um die Seite neu zu generieren werden Änderungen in den Master Branch gepushed.
Dadurch wird eine Pipeline getriggert die sich ein Hugo Docker Image holt und die Seite wird damit generiert (build).
Ist der Build fertig und ok wird ein ssh+rsync Docker Image geladen und die Seite wird auf den/die Webserver kopiert (deploy).

Die Konfiguration `.gitlab-ci.yml` der Pipeline sieht so aus:

```
stages:
  - build
  - deploy
build:
  stage: build
  image: registry.gitlab.com/ruedigerp/hugoci:latest
  script:
  - git submodule update --init --recursive
  - hugo -b "${BLOG_URL}"
  artifacts:
    paths:
    - public
    expire_in: 1 hour
  only:
  - master
build:
  stage: build
  image: registry.gitlab.com/ruedigerp/hugoci:latest
  script:
  - git submodule update --init --recursive
  - hugo -b "${BLOG_URL}"
  artifacts:
    paths:
    - public
    expire_in: 1 hour
  only:
  - master
deploy:
  stage: deploy
  image: registry.gitlab.com/ruedigerp/ci-deploy-rsync-ssh
  script:
  - echo "${SSH_PRIVATE_KEY}" > ${HOME}/id_rsa
  - chmod 400 ${HOME}/id_rsa
  - mkdir "${HOME}/.ssh"
  - echo "${SSH_KNOWN_HOSTS}" > "${HOME}/.ssh/known_hosts"
  - rsync -at --quiet --delete --delete-delay --delay-updates --exclude=_ --include=.well-known -e "ssh -i ~/id_rsa -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -p ${DEPLOY_PORT}" public/ ${DEPLOY_USER}@${DEPLOY_HOST}:${DEPLOY_DIR}
  variables:
    GIT_STRATEGY: none
  only:
  - master
```

Beim Build wird Die Seite  nach `./Public` generiert und als Artifakt für 1 Stunde gespeichert.
Das wird im Deploy Schritt dann wieder abgerufen und kann so dann in diesem Schritt benutzt werden.

Damit die Pipeline funktioniert müssen in den Settings noch die Variabeln hinterlegt werden.

* BLOG_URL
* SSH_PRIVATE_KEY
* SSH_KNOWN_HOSTS
* DEPLOY_PORT
* DEPLOY_USER
* DEPLOY_HOST
* DEPLOY_DIR

Den SSH Host Key bekommt man mit:

    ssh-keyscan -p $PORT dein.ssh.host.de

Wer noch keinen `ssh-key` hat oder einen neuen genieren will:

    ssh-keygen -t rsa -b 4096 [-f output_keyfile] [-C ci-cd-deploment]


~~Geplant sind noch weitere Stages um die Seite nicht nur lokal zu testen, sondern bevor sie Live jetzt auch im Internet in einer Testumgebung testen zu können.~~
Gerade wenn man Seiten mit mehreren Leuten betreut kann auf `staging` deployed werden und alle können überprüfen ob alles ok ist.
Erst danach wird `staging` in `master` gemerged und automatisch in `PROD` deployed.

Das ist jetzt auch umgesetzt mit folgender `.gitlab-ci.yml`.

```
stages:
  - build
  - deploy-dev
  - deploy
build:
  stage: build
  image: registry.gitlab.com/ruedigerp/hugoci:latest
  script:
  - git submodule update --init --recursive
  - hugo -b "${BLOG_URL}"
  artifacts:
    paths:
    - public
    expire_in: 1 hour
  only:
  - master
  - dev

dev:
  stage: deploy-dev
  image: registry.gitlab.com/ruedigerp/ci-deploy-rsync-ssh
  script:
  - echo "${SSH_PRIVATE_KEY}" > ${HOME}/id_rsa
  - chmod 400 ${HOME}/id_rsa
  - mkdir "${HOME}/.ssh"
  - echo "${SSH_KNOWN_HOSTS}" > "${HOME}/.ssh/known_hosts"
  - rsync -at --quiet --delete --delete-delay --delay-updates --exclude=_ --include=.well-known -e "ssh -i ~/id_rsa -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -p ${DEPLOY_PORT}" public/ ${DEPLOY_USER}@${DEPLOY_HOST}:${DEPLOY_DIR}dev
  variables:
    GIT_STRATEGY: none
  only:
  - dev

deploy:
  stage: deploy
  image: registry.gitlab.com/ruedigerp/ci-deploy-rsync-ssh
  script:
  - echo "${SSH_PRIVATE_KEY}" > ${HOME}/id_rsa
  - chmod 400 ${HOME}/id_rsa
  - mkdir "${HOME}/.ssh"
  - echo "${SSH_KNOWN_HOSTS}" > "${HOME}/.ssh/known_hosts"
  - rsync -at --quiet --delete --delete-delay --delay-updates --exclude=_ --include=.well-known -e "ssh -i ~/id_rsa -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -p ${DEPLOY_PORT}" public/ ${DEPLOY_USER}@${DEPLOY_HOST}:${DEPLOY_DIR}dev
  variables:
    GIT_STRATEGY: none
  only:
  - master
  - dev
  when: manual

```

Push in `dev` triggert die Pipeline. Die macht den Build und das Deploy dann in die Testumgebung.
Danach hält die Pipeline an und man kann alles testen. Wenn man dann nach `live` deployen will klickt man in der Pipeline und sie läuft weiter.

### Hugo Docker Image für GitLab Pipeline

Um die Seiten zu generieren ....

hugoci
Dockerfile:

```
FROM alpine:3.7

RUN apk add --update \
      git && \
    rm -rf /var/cache/apk/*

ENV HUGO_VERSION 0.42.2
ENV HUGO_RESOURCE hugo_${HUGO_VERSION}_Linux-64bit

ADD https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/${HUGO_RESOURCE}.tar.gz /tmp/

RUN mkdir /tmp/hugo && \
    tar -xvzf /tmp/${HUGO_RESOURCE}.tar.gz -C /tmp/hugo/ && \
    mv /tmp/hugo/hugo /usr/bin/hugo && \
    rm -rf /tmp/hugo*
```

.gitlab-ci.yml
```
image: docker:latest

services:
  - docker:dind

stages:
- build

variables:
  DOCKER_IMAGE_TAG: registry.gitlab.com/ruedigerp/ci-build-hugo

before_script:
  # - echo $CI_BUILD_TOKEN | docker login --username gitlab-ci-token --password-stdin registry.gitlab.com
  - echo "AAA-BBBBBBBBBBBBB" | docker login --username gitlab-ci-token --password-stdin registry.gitlab.com

build:
  stage: build
  script:
    - echo "AAA-BBBBBBBBBBBBB" | docker login --username gitlab-ci-token --password-stdin registry.gitlab.com
    - docker build --pull -t $DOCKER_IMAGE_TAG .
    - docker push $DOCKER_IMAGE_TAG
```


### Sync Deploy Docker Image

Um die Seiten zu deployen ...

Docker image: registry.gitlab.com/ruedigerp/ci-deploy-rsync-ssh
```
# folgt noch
```

.gitlab-ci.yml

```
# folgt noch
```
