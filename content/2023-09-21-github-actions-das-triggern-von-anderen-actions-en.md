---
title: 'GitHub Actions: Triggering Other Actions'
date: 2023-09-21 20:00:00
author: ruediger
cover: "/images/posts/github-actions-banner.webp"
tags: [GitHub, Action, Workflow, Trigger]
categories: 
    - Internet
preview: "GitHub Actions is a CI/CD service (Continuous Integration/Continuous Deployment) provided by GitHub, allowing developers to create and share automation workflows directly within their GitHub repositories. A common requirement is triggering one Action from another. In this article, we will explore how to effectively chain GitHub Actions to create more complex automation processes." 
draft: false
top: false
type: post
hide: true
toc: false
---

[English Version](/posts/2023-09-21-github-actions-das-triggern-von-anderen-actions-en.html)

![Github Actions](/images/posts/github-actions-banner.webp)

## Introduction

GitHub Actions is a CI/CD service (Continuous Integration/Continuous Deployment) provided by GitHub, allowing developers to create and share automation workflows directly within their GitHub repositories. A common requirement is triggering one Action from another. In this article, we will explore how to effectively chain GitHub Actions to create more complex automation processes.

## Basics of GitHub Actions

GitHub Actions are based on workflows, defined in YAML files within the .github/workflows directory of a repository. A workflow consists of one or more jobs, and each job contains steps, which are commands or Actions.

## Triggering Other Actions

To trigger a GitHub Action from another Action, several approaches are available:

1. Repository Dispatch Event:
   * One can trigger a Repository Dispatch Event from an Action to start a workflow in the same or a different repository.
   * This requires a Personal Access Token to authenticate the event.
2. Workflow Dispatch Event:
   * The Workflow Dispatch Event allows for manually triggering a workflow.
   * It can also be triggered by an API call, making it startable from another Action.
3. Scheduled Workflows:
   * Workflows can be configured to run at specific times.
   * This can be used to trigger subsequent workflows.



## Example: Repository Dispatch Event

Here is a simple example of using a Repository Dispatch Event to trigger a workflow:


```
# .github/workflows/triggering-workflow.yml
name: Triggering Workflow
on: [push]
jobs:
  trigger:
    runs-on: ubuntu-latest
    steps:
      - name: Trigger Another Workflow
        run: |
          curl -XPOST -u "USERNAME:${{ secrets.PAT }}" \
            -H "Accept: application/vnd.github.everest-preview+json" \
            -H "Content-Type: application/json" \
            https://api.github.com/repos/OWNER/REPO/dispatches \
            --data '{"event_type": "triggered"}'


```

```
# .github/workflows/triggered-workflow.yml
name: Triggered Workflow
on:
  repository_dispatch:
    types: [triggered]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v2
      # Additional steps...


```

In this example, a push event starts the triggering-workflow.yml workflow, which then triggers the triggered-workflow.yml through a Repository Dispatch Event.

## Conclusion

Triggering GitHub Actions from other Actions enables the creation of complex and flexible automation workflows. By using methods like the Repository Dispatch Event or the Workflow Dispatch Event, developers can efficiently design and optimize their CI/CD processes.