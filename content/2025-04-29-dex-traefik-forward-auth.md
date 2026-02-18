---
title: 'Dex SSO mit Traefik und forward-auth'
date: 2025-04-29 15:00:00
author: ruediger
cover: "/images/posts/2025/04/dex-sso.webp"
featureImage: "/images/posts/2025/04/dex-sso.webp"
tags: [SSO, Login, Traefik, Forward-Auth]
categories: 
    - SSO
    - Login
preview: "SSO mit Dex für Google, Github und andere Accounts."
draft: false
top: false
type: post
hide: false
toc: false
---

![Dex SSO-Login](/images/posts/2025/04/dex-sso.webp)

Ich benutze Dex für den SSO Login auf mehreren Seiten von mir. Ich habe E-Mail Logins und Github aktuell bei mir aktiviert. Alle anderen Möglichkeiten die Dex bietet, wie LinkedIn, Facebook, Google uns viele mehr habe ich auch alle einmal getestet. 

Für die Nutzung von GitHub muss im Github Account eine [OAuth App](https://github.com/settings/developers) angelegt werden. Damit bekommen wir eine Client-ID und ein Client-Secret. Die beiden trägt man in die `values.yaml` für den Dex Helm Chart ein. 

    helm repo add dex https://charts.dexidp.io
    helm repo update dex 

Hier eine Beispiel `values.yaml` in der man Domains, Client-ID und Secret noch anpassen muss. 

Beispiel values.yaml für Dex:

    ingress:
    enabled: true
    className: traefik
    annotations:
        cert-manager.io/cluster-issuer: letsencrypt-cluster-issuer
        traefik.ingress.kubernetes.io/router.entrypoints: websecure
        traefik.ingress.kubernetes.io/router.tls: "true"
    hosts:
        - host: dex.example.com
        paths:
            - path: /
            pathType: ImplementationSpecific
    tls:
        - secretName: dex.example.com-cert
        hosts:
            - dex.example.com
    config:
    connectors:
    - config:
        clientID: 5J3FvbjYDQm4XeofPKtc
        clientSecret: iLuvTwVgbtrcLftWjqy2cIWVb0PVRGZG6F39XCEi
        loadAllGroups: true
        orgs:
        - name: exmaple-Org
        redirectURI: https://dex.example/callback
        teamNameField: slug
        useLoginAsID: true
        id: github
        name: GitHub
        type: github
    enablePasswordDB: true
    frontend:
        issuer: kuepper.nrw
        logoURL: https://blog.example.com/images/avatar.png
        theme: dark
    issuer: https://dex.example.com
    logger:
        level: debug
    oauth2:
        alwaysShowLoginScreen: true
        skipApprovalScreen: true
    staticClients:
    - id: github
        name: Traefik Forward Auth OIDC Dex App
        public: false
        redirectURIs:
        - https://login.example.com/_oauth
        - https://dex.example.com/_oauth
        - https://intern.example.com/_oauth
        secret: iLuvTwVgbtrcLftWjqy2cIWVb0PVRGZG6F39XCEi
    staticPasswords:
    - email: user1@example.com
        hash: $2y$10$D6QwNsN5j7hWOHOEPsXMcu77FJOT1Ae2bkfsGKjZLyx26kwZ9UA.S
        userID: 08a8684b-db88-4b73-90a9-3cd1661f5466
        username: user1
    storage:
        config:
        inCluster: true
        type: kubernetes

Mit dieser values.yaml wird jetzt erst einmal Dex installiert.

    helm upgrade -i dex dex/dex -n dex -f values.yaml --create-namespace --namespace dex --wait

Damit nachher auch Seiten umgeleitet werden können brauchen wir noch einen forward-auth Service im Cluster. 
Der kümmert sich um die Umleitung auf dex und den Callback nach dem Login. 

Hier hatte ich mir oauth2-proxy einmal angeguckt. Der muss aber dann in jedem Cluster eingerichtet und Konfiguriert werden. Ausserdem muss er in jeden Namespace rein damit es auch mit allen Services funktioniert. 

forward-auth ist da viel kleiner und einfacher zu konfigurieren. 

    helm repo add kuepper https://helm.9it.eu/
    helm repo update kuepper 

Hier eine Beispiel values.yaml für den forward-auth:

    deployment:
      args:
        - --secret=iLuvTwVgbtrcLftWjqy2cIWVb0PVRGZG6F39XCEi
        - --auth-host=login.example.com
        - --cookie-domain=example.com
        - --default-provider=oidc
        - --providers.oidc.issuer-url=https://dex.example.com
        - --providers.oidc.client-id=github
        - --providers.oidc.client-secret=iLuvTwVgbtrcLftWjqy2cIWVb0PVRGZG6F39XCEi

    ingress:
      domains:
        - name: auth.example.com
        tls: true
        servicename: forward-auth
        port: 4181
        path: /
        pathtype: Prefix

Die values.yaml hat noch mehr Values die angepasse werden können. Ein Liste mit allen Möglichkeiten kann mit `helm values ...` abgerufen werden:

    helm show values kuepper/traefik-forward-auth

Wenn alles installiert ist kann man einen Service, der mit einem Login versehen werden soll einfach per Treafik Ingress Annotation auf den forward-auth und so auf Dex umleiten. 

In einem Ingress einfach folgende Annotaions hinzufügen oder einen Middleware Eintrag erweitern:

    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      annotations:
        traefik.ingress.kubernetes.io/router.middlewares: kube-system-forward-auth-dex@kubernetescrd
      name: gethomepage
      namespace: gethomepage
    ...

Ruft man jetzt die Seite auf sollte ein Redirect auf Dex erfolgen. Dann kann direkt der Login mit github getestet werden. 


