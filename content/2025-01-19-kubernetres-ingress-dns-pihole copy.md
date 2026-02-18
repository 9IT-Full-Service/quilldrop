---
title: 'Kubernetres Ingress DNS PiHole'
date: 2025-01-19 0:00:00
author: ruediger
cover: "/images/posts/2025/01/pihole-local-dns.webp"
featureImage: "/images/posts/2025/01/pihole-local-dns.webp"
tags: [Kubernetes, HomeLab, external-DNS, PiHole]
categories: 
    - Kubernetes
preview: "Home Lab mit PiHole lokalen DNS Records."
draft: false
top: false
type: post
hide: false
toc: false
---

![piHole local dns](/images/posts/2025/01/pihole-local-dns.webp)

## Lokales Lab auf dem Mac mit DNS Recordes

Ich habe auf dem Mac Mini mit Multipass mehrere Server, die mit Kubernetes mehrere Dienste bereitstellen. Damit diese auch per Domain angesprochen werden benutze ich die lokalen DNS Records im PiHole. 

Wenn ich einen Service per Ingress oder Service bereitstelle wird automatisch ein DNS Record angelegt. Dazu benutze ich External-DNS. Ich kann damit entweder im PiHole die DNS Records anlegen oder bei Cloudflare in einer Zone. 
External-DNS kann aber auch mit [sehr vielen](https://github.com/kubernetes-sigs/external-dns?tab=readme-ov-file#the-latest-release) anderen DNS-Providern genutzt werden. 

In diesem Post beschreibe ich das einrichten im PiHole. 

## Konfiguration 

Namespace anlegen:

    kubectl create ns external-dns

Secret anlegen:

    kubectl create secret generic pihole-password --from-literal EXTERNAL_DNS_PIHOLE_PASSWORD=YourSecret -n external-dns


YourSecret ist das Passwort für den Login in Deinem PiHole. 

pihole.yaml

    ---
    apiVersion: v1
    kind: ServiceAccount
    metadata:
    name: external-dns
    namespace: external-dns
    ---
    apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRole
    metadata:
    name: external-dns
    namespace: external-dns
    rules:
    - apiGroups: [""]
    resources: ["services","endpoints","pods"]
    verbs: ["get","watch","list"]
    - apiGroups: ["extensions","networking.k8s.io"]
    resources: ["ingresses"]
    verbs: ["get","watch","list"]
    - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["list","watch"]
    ---
    apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRoleBinding
    metadata:
    name: external-dns-viewer
    namespace: external-dns
    roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: external-dns
    subjects:
    - kind: ServiceAccount
    name: external-dns
    namespace: external-dns 
    ---
    apiVersion: apps/v1
    kind: Deployment
    metadata:
    name: external-dns
    namespace: external-dns
    spec:
    strategy:
        type: Recreate
    selector:
        matchLabels:
        app: external-dns
    template:
        metadata:
        labels:
            app: external-dns
        spec:
        serviceAccountName: external-dns
        containers:
        - name: external-dns
            image: registry.k8s.io/external-dns/external-dns:v0.14.1
            # If authentication is disabled and/or you didn't create
            # a secret, you can remove this block.
            envFrom:
            - secretRef:
                # Change this if you gave the secret a different name
                name: pihole-password
            args:
            - --source=service
            - --source=ingress
            - --domain-filter=kuepper.lab
            # Pihole only supports A/CNAME records so there is no mechanism to track ownership.
            # You don't need to set this flag, but if you leave it unset, you will receive warning
            # logs when ExternalDNS attempts to create TXT records.
            - --registry=noop
            # IMPORTANT: If you have records that you manage manually in Pi-hole, set
            # the policy to upsert-only so they do not get deleted.
            - --policy=upsert-only
            - --provider=pihole
            # Change this to the actual address of your Pi-hole web server
            - --pihole-server=http://10.0.2.240
        securityContext:
            fsGroup: 65534 # For ExternalDNS to be able to read Kubernetes token files

In dem `pihole.yaml` muss die IP für den piHole angepasst werden: 

    - --pihole-server=http://10.0.2.240

Zusätzlich sollte die Domain angepasst werden. 

Die IP 10.0.2.240 mit der IP Deines PiHole ersetzen, da Du sehr wahrscheinlich nicht `kuepper.lab` benutzen willst.

    - --domain-filter=yourdomain.tld 

Das yaml-File in Deinem Kubernetes installieren und schon werden die Records automatisch angelegt.

    kubectl apply -f pihole.yaml


## Ingress anlegen und Tests

ingress.yaml

    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
    name: ink-blog
    namespace: ink
    spec:
    ingressClassName: traefik
    rules:
    - host: ink.kuepper.lab
        http:
        paths:
        - backend:
            service:
                name: ink-blog
                port:
                number: 80
            path: /
            pathType: Prefix

Den Ingress in Deinem Cluster anlegen:

    kubectl apply -f ingress.yaml 

Im Logfile vom external DNS sollte 1-2 Minuten nach dem apply folgende zeile im Log auftauchen:

    external-dns-786bc76c79-hbw84 time="2025-01-18T23:39:56Z" level=info msg="add ink.kuepper.lab IN A -> 192.168.64.6"

Die IP `192.168.64.6` ist die IP von meinem Cluster und der Domainname (fqdn) ink.kuepper.lab wird im PiHole eingetragen.


![piHole local dns](/images/posts/2025/01/pihole-local-dns.png)

Überprüfen wir das anschliessend, wird der hostname richtige aufgelöst:

    host ink.kuepper.lab
    ink.kuepper.lab has address 192.168.64.6

Dieses Setup ist für lokale Tests, kann aber auch für Services genutzt werden die extern erreichbar sind, so das die benötigten DNS Record automatisch bei Cloudflare oder anderen Providern eingerichtet werden. Dabei kann jeder Service auch Proxied bei Cloudflare eingerichtet werden. Also auch it SSL Zertifikaten. Im Free Account für Cloudflare können keine sub-sub Domains also xzy.foobar.yourdomain.tld mit einem Zertifikat ausgestattet werden. Diese kann man dann einfach ohne Cloudflare Proxy einrichten und dann lokal per Cert-Manager und Letsencrypt einfach mit Zertifikate ausstatten. 

Das beschreibe ich dann in einem weiterem Post. In diesem werde ich dann zeigen wie man das Zertifikat pro Domain oder einfach ein Wildcard Zertifikat erstellt. Je nachdem was man benötigt und wie Dein Setup ist. 
