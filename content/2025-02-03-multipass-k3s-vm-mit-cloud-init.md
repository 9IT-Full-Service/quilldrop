---
title: 'Multipass K3S VM mit Cloud-init'
date: 2025-02-03 21:00:00
author: ruediger
cover: "/images/posts/2024/12/tmux.webp"
featureImage: "/images/posts/2024/12/tmux.webp"
tags: [Kubernetes, k3s, HomeLab, multipass, Cloud-init]
categories: 
    - Kubernetes
preview: "Schnell eine k3s VM mit multipass erstellen "
draft: false
top: false
type: post
hide: false
toc: false
---

![k3s k9s](/images/posts/2024/12/tmux.webp)

# Multipass installieren

    ❯ brew install multipass

# Cloud-Init Yaml erstellen 

cloud-init.yaml

    packages:
    # - traceroute
    # - frr

    runcmd:
    - export HOME='/home/ubuntu'
    - export USER='ubuntu'
    - cd $HOME
    - curl -sfL https://get.k3s.io | sh -s - server --cluster-init --disable=servicelb --tls-san=192.168.64.251 --disable=traefik
    - echo "export KUBECONFIG=/etc/rancher/k3s/k3s.yaml" > /root/.bashrc
    - wget https://github.com/derailed/k9s/releases/download/v0.32.7/k9s_Linux_arm64.tar.gz; tar xzf k9s_Linux_arm64.tar.gz; mv k9s /usr/local/bin/;
    - curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3; chmod 700 get_helm.sh; ./get_helm.sh

# VM erstellen 

    ❯ multipass launch -c 1 -m 1G --disk 10G --network en0 --name controlplane-test --cloud-init cloud-init.yaml

    ❯ multipass list
    Name                    State             IPv4             Image
    cloud-init-test         Running           192.168.64.45    Ubuntu 24.04 LTS
                                            10.0.2.109
                                            10.42.0.0
                                            10.42.0.1


# In VM einloggen 

    ❯ multipass shell controlplane-test
    Welcome to Ubuntu 24.04.1 LTS (GNU/Linux 6.8.0-51-generic aarch64)

    Last login: Mon Feb  3 20:26:31 2025 from 192.168.64.1
    ubuntu@controlplane-test:~$ sudo su - 
    root@controlplane-test:~# kubectl get nodes
    NAME              STATUS   ROLES                       AGE   VERSION
    controlplane-test Ready    control-plane,etcd,master   7m    v1.31.5+k3s1
    root@controlplane-test:~# kubectl get pods -A
    NAMESPACE     NAME                                      READY   STATUS    RESTARTS        AGE
    kube-system   coredns-ccb96694c-knm8m                   1/1     Running   0               6m58s
    kube-system   local-path-provisioner-5cf85fd84d-h9rjg   1/1     Running   2 (6m49s ago)   6m58s
    kube-system   metrics-server-5985cbc9d7-cv25j           1/1     Running   2 (6m49s ago)   6m58s
    root@controlplane-test:~#

# VM löschen

    ❯ multipass delete controlplane-test
    ❯ multipass purge 
