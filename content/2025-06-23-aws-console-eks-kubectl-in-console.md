---
title: 'AWS EKS kubectl in Web-Console'
date: 2025-06-23 11:00:00
update: 2025-06-23 11:00:00
author: ruediger
cover: "/images/posts/2025/06/aws-els-console-kubeconfig.webp"
featureImage: "/images/posts/2025/06/aws-els-console-kubeconfig.webp"
tags: [aws, eks, kubernetes, kubeconfig, console]
categories: 
    - Kubernetes
preview: "AWS Web Console und schnell mal eben auf einem Kubernetes Cluster etwas überprüfen."
draft: false
top: false
type: post
hide: false
toc: false
---

![Automatisierte Kubernetes Volume-Backups](/images/posts/2025/06/aws-els-console-kubeconfig.webp)

Da ich gerade genau das einmal wieder machen musste, nicht direkt an die KubeConfig gekommen bin, habe ich kurz über die Shell in der AWS Console nachgeschaut. 
Und da man so etwas immer wieder vergisst und suchen muss wie man es machen, einfach mal hier kurz festgehalten. 

In der Shell ist erst einmal keine `./kube/config` vorhanden. Diese kann man sich aber schnell in die Shell importieren. 

```bash
aws eks update-kubeconfig --region eu-central-1 --name my-cluster
```

Anschliessend kann man direkt mit kubectl auf den Cluster zugreifen. 

