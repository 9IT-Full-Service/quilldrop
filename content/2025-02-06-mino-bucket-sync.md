---
title: 'Minio Bucket sync'
date: 2025-02-06 12:00:00
author: ruediger
cover: "/images/posts/2025/02/minio.webp"
featureImage: "/images/posts/2025/02/minio.webp"
tags: [Minio, Bucket, Backup, Restore]
categories: 
    - Kubernetes
preview: "Minio S3-Bucket synchronisieren."
draft: true
top: false
type: post
hide: true
toc: false
---

![minio](/images/posts/2025/02/minio.webp)


    brew install s3cmd
    brew install rclone
    s3cmd --configure
    s3cmd get s3://terraform-state/state/terraform.tfstate
    rclone config
    rclone lsd tfstate:
    rclone ls tfstate:terraform-state
    rclone sync tfstate:terraform-state ./

    brew install s3cmd
    brew install rclone
    rclone config
    rclone lsd ftstate2:
    rclone mkdir ftstate2:terraform-state
    rclone copy  state ftstate2:terraform-state
    rclone lsd ftstate2:

    rclone sync tfstate:terraform-state ./
    rclone sync state ftstate2:terraform-state

