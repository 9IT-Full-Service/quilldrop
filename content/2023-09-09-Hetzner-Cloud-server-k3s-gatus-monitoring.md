---
title: Hetzner-Cloud-server-k3s-gatus-monitoring
date: 2023-09-09 11:26:47
author: ruediger
tags:
  - Hetzner
  - k3s
  - monitoring
categories: 
  - Internet
draft: true
top: false
type: post
hide: true
toc: false
---


```
#!/bin/bash

# INSTALL K3S
curl -sfL https://get.k3s.io | sh
echo 'export KUBECONFIG=/etc/rancher/k3s/k3s.yaml' >> ~/.bashrc
source ~/.bashrc

# INSTALL HELM
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

# INSTALL K9S
# ARM 
curl -L -O https://github.com/derailed/k9s/releases/download/v0.27.4/k9s_Linux_arm64.tar.gz
tar -xzf k9s_Linux_arm64.tar.gz
# AMD
# curl -L -O https://github.com/derailed/k9s/releases/download/v0.27.4/k9s_Linux_amd64.tar.gz
# tar -xzf k9s_Linux_amd64.tar.gz
mv k9s /usr/local/bin/.

cat /etc/rancher/k3s/k3s.yaml

```

```
cat << 'EOF' | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    email: ruediger@kuepper.nrw
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-prod
    solvers:
    - http01:
        ingress:
          class: traefik
EOF          
```

```
hcloud server create --image ubuntu-22.04 --type cx11 --name Ubuntu22-server --location fsn1 --ssh-key ruediger@kuepper.nrw
```


```
export INSTALL_K3S_VERSION=v1.21.3+k3s1
export INSTALL_K3S_EXEC="server --disable traefik --disable servicelb --disable metrics-server --disable-cloud-controller \
       --kube-proxy-arg proxy-mode=ipvs --cluster-cidr=10.42.0.0/16,fd42::/48 --service-cidr=10.43.0.0/16,fd43::/112 \
       --disable-network-policy --flannel-backend=none --node-ip=162.55.33.231,2a01:4f8:c013:882::1"

wget https://get.k3s.io -O k3s.sh
less k3s.sh
bash k3s.sh
```

```
wget https://docs.projectcalico.org/manifests/calico.yaml
```

```
"ipam": {
    "type": "calico-ipam",
    "assign_ipv4": "true",
    "assign_ipv6": "true"
},
```

```
kubectl apply -f calico.yaml
```


```
wget https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/namespace.yaml -O metallb-namespace.yaml
wget https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/metallb.yaml -O metallb-0.10.2-manifest.yaml
kubectl apply -f metallb-namespace.yaml -f metallb-0.10.2-manifest.yaml
```

```
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: metallb-system
  name: config
data:
  config: |
    address-pools:
    - name: default
      protocol: layer2
      addresses:
      - 162.55.33.231/32
      - 2a01:4f8:c013:882::1/128
```

```
kubectl apply -f metallb-config.yaml
```

```
cat << 'EOF' | kubectl apply -f -
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: default-pool
  namespace: metallb-system
spec:
  addresses:
  - 162.55.33.231/32
  - 2a01:4f8:c013:882::/64
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: default
  namespace: metallb-system
spec:
  ipAddressPools:
  - default-pool
EOF
```


```
helm install cilium cilium/cilium \
--namespace kube-system \
--set ipv4.enabled=true \
--set ipv6.enabled=true \
--set ipam.mode=cluster-pool \
--set ipam.operator.clusterPoolIPv4PodCIDRList="10.96.0.0/16" \
--set ipam.operator.clusterPoolIPv6PodCIDRList="2a01:4f8:c013:882::/96" \
--set ipam.operator.clusterPoolIPv4MaskSize=24 \
--set ipam.operator.clusterPoolIPv6MaskSize=112 \
--set bpf.masquerade=true \
--set enableIPv6Masquerade=false
```

<!-- 
```
# First add metallb repository to your helm
helm repo add metallb https://metallb.github.io/metallb
# Check if it was found
helm search repo metallb
# Install metallb
helm upgrade --install metallb metallb/metallb --create-namespace \
--namespace metallb-system --wait
```

```
cat << 'EOF' | kubectl apply -f -
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: default-pool
  namespace: metallb-system
spec:
  addresses:
  - 188.34.183.120-188.34.183.120
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: default
  namespace: metallb-system
spec:
  ipAddressPools:
  - default-pool
EOF
```

```kubectl get pods -n metallb-system ```

```
root@control01:~/metallb# kubectl get pods -n metallb-system
NAME                         READY   STATUS    RESTARTS   AGE
controller-57fd9c5bb-rdl7v   1/1     Running   0          5m42s
speaker-h7chj                1/1     Running   0          5m42s
speaker-pg7kp                1/1     Running   0          5m42s
speaker-78pdz                1/1     Running   0          5m42s
speaker-ghpxz                1/1     Running   0          5m42s
speaker-8cf7k                1/1     Running   0          5m42s
speaker-2t6jp                1/1     Running   0          5m42s
speaker-cjcpn                1/1     Running   0          5m41s
speaker-mv7v4                1/1     Running   0          5m42s
```

```
kubectl get svc -n kube-system
NAME                   TYPE              CLUSTER-IP    EXTERNAL-IP       PORT(S)                                 AGE
kube-dns           ClusterIP        10.43.0.10           <none>           53/UDP,53/TCP,9153/TCP       3h45m
metrics-server   ClusterIP         10.43.254.144     <none>           443/TCP                                  3h45m
traefik                LoadBalancer 10.43.159.145      192.168.0.200        80:31771/TCP,443:30673/TCP 3h44m

```

```
kubectl get events -n kube-system --field-selector involvedObject.name=traefik
LAST SEEN TYPE REASON OBJECT MESSAGE
61s Normal IPAllocated service/traefik Assigned IP ["192.168.0.200"]
60s Normal nodeAssigned service/traefik announcing from node "cube02" with protocol "layer2"
```

 -->


