---
title: Github workflow trigger in other repository
date: 2023-09-21 21:00:00
author: ruediger
cover: "/images/posts/github-actions-banner.webp"
tags: [GitHub, Action, Workflow, Trigger]
categories: 
    - Internet
preview: "Im ersten Repository wird ein Workflow `trigger_on_push.yaml` erstellt, der in einem anderen Repo einen Worklow `trigger_prod.yaml` triggert. Dieser kann sogar in einem anderen GitHub Account liegen. Man benötigt nur einen gültigen Token. " 
draft: false
top: false
type: post
hide: false
toc: false
---

[English Version](/posts/2023-09-21-github-workflow-trigger-in-other-repository-en.html)

![Github Actions](/images/posts/github-actions-banner.webp)

## Eine Github Action in einem anderem Repository triggern

Im ersten Repository wird ein Workflow `trigger_on_push.yaml` erstellt, der in einem anderen Repo einen Worklow `trigger_prod.yaml` triggert. Dieser kann sogar in einem anderen GitHub Account liegen. Man benötigt nur einen gültigen Token. 

Aber wieso über eine dazwischen geschaltete Action? Ganz einfach, das 2. Repository bekommt keine Pushes, ausser es würde sich etwas in diesem Repository ändern. Da aber die Änderungen nur im 1. Repository gemacht werden und das zweite nur die Software für das Generieren der Page und der Docker Images enthält, wird sich da nichts ändern. Also kann Semantic Release auch keine Version hochzählen. 

Daher die Trigger Action im 2. Repo. Die ändert nur eine Datei im Repo und pushed es in sich selbst. Danach Triggert sie die eigentliche Build Action `main.yaml`. Damit ist eine Änderung vorhanden und Semantic Release findet Commits, die dann in dieses Release mit rein kommen. 

Repo 1 Workflow `trigger_on_push.yaml`: 

```
name: Build Develop

on:
    push:
        branches:
            - develop

jobs:
    trigger-workflow:
        runs-on: ubuntu-latest
        steps:
            - name: Trigger build Repo
              uses: convictional/trigger-workflow-and-wait@v1.6.1
              with:
                owner: ruedigerp
                repo: ink-blog
                github_token: ${{ secrets.GH_TOKEN }}
                workflow_file_name: trigger_prod.yaml 
                ref: develop 
                wait_interval: 10 
                propagate_failure: true 
                trigger_workflow: true
                wait_workflow: true
```

Im 2. Repo `trigger_prod.yaml`

```
name: Trigger Build Prod

on:
  workflow_dispatch:
    inputs:
        workflow_02:
            description: 'ًWorkflow 2 which will be triggered'
            required: true
            default: 'trigger-workflow'

        workflow2_github_account:
            description: 'GitHub Account Owner'
            required: true
            default: 'youraccount'

        workflow2_repo_github:
            description: 'repo-name'
            required: true
            ref: main
            default: 'repo1'


jobs:
  trigger-build: 
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      - name: add Trigger page
        run: |
          ls -la 
          echo "build $(date)" > buildit.txt
          ls -la
          git remote -v
          export GH_TOKEN="${{ secrets.DOCKHUB_TOKEN }}"
          git config --global user.email "your@main.com"
          git config --global user.name "Your Name"
          git add .
          git commit -m "fix: new build $(date)"
          git push

  trigger-workflow:
    runs-on: ubuntu-latest
    steps:
        - name: Trigger build Repo
          uses: convictional/trigger-workflow-and-wait@v1.6.1
          with:
            owner: ruedigerp
            repo: repo2
            github_token: ${{ secrets.DOCKERHUB_TOKEN }}
            workflow_file_name: main.yaml 
            ref: main 
            wait_interval: 10 
            propagate_failure: true 
            trigger_workflow: true
            wait_workflow: true   
```

Die `main.yaml` in Repo2: 

```
name: Build Prod

on:
  push:
    branches:
      - 'main'
  workflow_dispatch:
    inputs:
        workflow_02:
            description: 'ًWorkflow 2 which will be triggered'
            required: true
            default: 'trigger-workflow'

        workflow2_github_account:
            description: 'GitHub Account Owner'
            required: true
            default: 'youraccount'

        workflow2_repo_github:
            description: 'repo-name'
            required: true
            default: 'repo2'

env:
  IMAGE_NAME: ghcr.io/youraccount/imagename
  DOCKERFILE: Dockerfile

jobs:
  prepare:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v2
        with:
          fetch-depth: 0

      - name: Create Sematic Release Version
        env:
          GITHUB_TOKEN: ${{ secrets.DOCKERHUB_TOKEN }}
        run: |
          export HOME=/tmp
          git config --global user.email "your@main.com"
          git config --global user.name "youraccount"
          git config --global credential.helper cache
          npx semantic-release

  build-amd64:
    needs: [ prepare ]
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v2
        with:
          fetch-depth: 0
      
      - name: Login to GitHub Container Registry
        uses: docker/login-action@v2
        with:
          registry: "ghcr.io"
          username: ${{ secrets.GH_USERNAME }}
          password: ${{ secrets.GH_TOKEN }}

      - name: Build and push AMD64 image
        run: |
          # ...
          # build stuff 
          # ...
          docker build -f ${{ env.DOCKERFILE }} --build-arg GOOS=linux --build-arg GOARCH=amd64 --build-arg NEXT_VERSION="$VERSION" --build-arg BUILD_TIMESTAMP="$BUILD_TIMESTAMP" --build-arg GIT_BRANCH="$GIT_BRANCH" --build-arg COMMIT_HASH="$COMMIT_HASH" -t ${{ env.IMAGE_NAME }}:$VERSION-amd64 -t ${{ env.IMAGE_NAME }}:latest-amd64 .
          docker push ${{ env.IMAGE_NAME }}:$VERSION-amd64          
          # ... 

  merge-images: 
     # ... 
  other-cool-stuff: 
     # ... 
```




