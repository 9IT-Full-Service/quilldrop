---
title: 'FluxCD Globale Variabeln für GetHomepage'
date: 2025-10-16 20:00:00
update: 2025-10-16 20:00:00
author: ruediger
cover: "/images/posts/2025/10/fluxcd-globale-vars-fuer-gethomepage.webp"
featureImage: "/images/posts/2025/10/fluxcd-globale-vars-fuer-gethomepage.webp"
# images: 
#   - /images/posts/2025/08/telekom-mail-fail.webp
tags: [Kubernetes, FluxCD, GlobalVars, GetHomepage]
categories: 
  - Kubernetes
preview: "Globale Variabeln in FluxCD für Ressourcen nutzen. Zum Beispiel um in Ingress Annotations für GetHomepage zu setzen."
series: ["FluxCD"] 
draft: false
top: false
type: post
hide: false
toc: false
---

Damit GetHomepage Services anzeigt und in Gruppen zuordnet kann man im Ingress dafür Annotations setzen. 
Änderungen sollen nicht zu Aufwendig sein, daher habe ich sie als Variabeln in FluxCD gesetzt. 
Meine Defaults sind enabled, group, pod_selector und icon. Die werden immer gesetzt. 
Weitere werden in den Applications nachher noch bei Bedarf per Patch hinzugefügt. Oder die Defaults überschrieben. 

Aber erst einmal der Reige nach. Ein Ingress Object sieht erst einmal so aus. der Clusterissuer und die Middleware gehören nicht zu GetHomepage, werden aber auch gleich mit gemacht. So können diese auch gleich schnell geändert werden wenn sich z.B. der ClusterIssuer ändert. 

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
  clusterissuer: "cf-letsencrypt-prod"
  middlewares: "kube-system-redirect-scheme@kubernetescrd"
  gethomepage_enabled: '"true"'
  gethomepage_group: "Default"
  gethomepage_pod_selector: "''"
  gethomepage_icon: "homepage"
...
```

Diese Annotations, eher die Values werden einfach in einer ConfigMap gespeichert. Auf diese wird später im Ingress verwiesen. 

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: global-vars
  namespace: flux-system
data:
  clusterissuer: "cf-letsencrypt-prod"
  middlewares: "kube-system-redirect-scheme@kubernetescrd"
  gethomepage_enabled: '"true"'
  gethomepage_group: "Default"
  gethomepage_pod_selector: "''"
  gethomepage_icon: "homepage"
``` 

Damit die Configmap genutzt wir muss noch die kustomization.yaml in clusters/prod/flux-system/ angepasst werden. 

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- gotk-components.yaml
- gotk-sync.yaml
- gitops-repos.yaml 
- global_vars.yaml
...
```

Jetzt müssen wir nur noch die einzelnen Applications anpassen. Sie liegen bei mir im GitOps Repository in /apps/<appname>/base/ingress.yaml 

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: blog
  annotations:
    cert-manager.io/cluster-issuer: ${clusterissuer}
    traefik.ingress.kubernetes.io/router.middlewares: ${middlewares}
    gethomepage.dev/enabled: ${gethomepage_enabled}
    gethomepage.dev/group: ${gethomepage_group}
    gethomepage.dev/pod-selector: ${gethomepage_pod_selector}
spec:
  ingressClassName: traefik
  rules:
...
```

In der Kustomization der Application muss jetzt noch noch die global_vars.yaml in der App hinzugefügt werden. 

```yaml
apiVersion: kustomize.toolkit.fluxcd.io/v1
kind: Kustomization
metadata:
  name: blog
  namespace: flux-system
spec:
  interval: 2m
  prune: false
  path: ./apps/blog/prod
  sourceRef:
    kind: GitRepository
    name: flux-system
  targetNamespace: blog
  decryption:
    provider: sops
    secretRef:
      name: sops-age
  postBuild:
    substituteFrom:
      - kind: ConfigMap
        name: global-vars
    substitute:
      ingress_host: blog.kuepper.nrw
      tls_secret_name: blog.kuepper.nrw-tls
      middlewares: "kube-system-redirect-scheme@kubernetescrd,blog-redirect-blog-feed@kubernetescrd"
```

Die ConfigMap `global_vars` wird mit substituteFrom geladen und die Variabeln sind verfügbar. In diesem Fall wird die Middleware auch gleich überschrieben. Denn das Blog ein paar Redirects, die aktiviert werden müssen. 

Wenn jetzt alles ins Repo gepushed wird, werden nach kurzer Zeit die Annotations in dem Ingress auftauchen. Die Seite wird dann in GetHomepage angezeigt in der Gruppe Default. 

Möchte man jetzt ein paar Sachen Ändern oder erweitern geht man in die Kustomization der Application um sie anzupassen. 

Ich will beim Blog die Description, die URL für den link und das Icon anpassen.

```yaml
...
  patches:
    - patch: |-
        - op: add
          path: /metadata/annotations/gethomepage.dev~1description
          value: 'Blog Küpper'
        - op: add
          path: /metadata/annotations/gethomepage.dev~1href
          value: 'https://blog.kuepper.nrw'
        - op: add
          path: /metadata/annotations/gethomepage.dev~1icon
          value: 'mdi-web'
      target:
        kind: Ingress
        name: blog
  postBuild:
...
```

Möchte man die Group `Default` mit `Homepage` überschrieben fügt man einfach den entsprechenden Patch dazu:

```yaml
        - op: add
          path: /metadata/annotations/gethomepage.dev~1group
          value: 'Homepages'
```

Alles eingecheckt und FluxCD hat es ausgerollt, sieht der fertige Ingress im Cluster so aus:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    cert-manager.io/cluster-issuer: cf-letsencrypt-prod
    gethomepage.dev/description: Blog Küpper
    gethomepage.dev/enabled: "true"
    gethomepage.dev/group: Homepage
    gethomepage.dev/href: https://blog.kuepper.nrw
    gethomepage.dev/icon: mdi-web
    gethomepage.dev/pod-selector: ""
    traefik.ingress.kubernetes.io/router.middlewares: kube-system-redirect-scheme@kubernetescrd,blog-redirect-blog-feed@kubernetescrd
  labels:
...
```

Die Variabeln können auch auf mehrere Dateien aufgeteilt werden. Je nachdem wofür sie genutzt werden hat man sie dann auch suaber getrennt. Es ist auch möglich Variabeln die nichz alle User eines Repos sehen sollen aus einem anderen Repo zu laden. Als Beispiel wäre der ClusterIssuer, der nicht angepasst werden darf. Oder Daten die z.B. für einen SecretStore, DNS Api Token für external DNS oder den Certmanager. Die sollten zwar eh mit Age oder anderen Tools verschlüsselt sein, aber angenommen das wäre nicht der Fall, würde man sie einfach an einem anderen Ort ablegen. 

```yaml
apiVersion: source.toolkit.fluxcd.io/v1
kind: GitRepository
metadata:
  name: flux-vars-repo
  namespace: flux-system
spec:
  interval: 5m
  url: ssh://git@github.com/username/gitops-vars
  ref:
    branch: main
  secretRef:
    name: gitops-testapps-auth
```

Das ganze muss natürlich noch genutzt werden: 

```yaml
apiVersion: kustomize.toolkit.fluxcd.io/v1
kind: Kustomization
metadata:
  name: flux-vars
  namespace: flux-system
spec:
  interval: 10m
  sourceRef:
    kind: GitRepository
    name: flux-vars-repo
  path: ./production
  prune: true
```

```yaml
apiVersion: kustomize.toolkit.fluxcd.io/v1
kind: Kustomization
metadata:
  name: globalVars
  namespace: flux-system
spec:
  interval: 10m
  prune: true
  path: ./vars
  sourceRef:
    kind: GitRepository
    name: flux-vars-repo
  targetNamespace: flux-system
```

In dem repo `github.com/username/gitops-vars` legt man den Ordner `vars` an und legt dort einfach eine kustomization.yaml an mit folgendem Inhalt: 

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - global_vars.yaml

namespace: flux-system
```

Die `global_vars.yaml` kopiert man jetzt ins neue Repo unter `vars/global_vars.yaml` und checkt alles ein. 

FluxCD fügt nach dem Reconcil das Git Repo hinzu, legt entsprechend alles auf dem Cluster an und die Variabeln kommen aus einem externem Repo. Im Grunde das gleiche wie auch alle anderen Sachen in FluxCD von externen Git Repos genutzt werden können. Das ist ja das schöne an FluxCD. Man kann den Code so anpassen wie man es benötigt. Gerade wenn man immer die gleichen Teile auf mehreren Clustern oder Umgebungen wieder verwendet ist eine Trennung diese Teile sinnvoll. Man muss sie nur einmal anfassen und sie werden auf allen Clustern angewendet. 
Oder wenn man Infrastruktur, Tools von anderen Sachen trennen möchte. Wenn z.B. Entwickler Ihre Sachen selbst machen sollen, dann haben sie für Ihr Zeug ein eigenes Repo und können dort alles machen was sie wollen. 
