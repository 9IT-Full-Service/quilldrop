---
title: 'HomeLab mit mehreren Clustern hinter einer IP'
date: 2025-03-16 23:00:00
author: ruediger
cover: "/images/posts/2025/03/haproxy.webp"
featureImage: "/images/posts/2025/03/haproxy.webp"
tags: [HomeLab, Kubernetes, NFS, HAProxy]
categories: 
    - Kubernetes
preview: "HomeLab mit mehreren Clustern hinter einer IP."
draft: false
top: false
type: post
hide: false
toc: false
---

![HAProxy](/images/posts/2025/03/haproxy.webp)

# Homelab Übersicht 

Mein kleines Kubernetes Cluster Setup sieht aktuell wie folgt aus. 


<pre class="mermaid">
graph TB;
    Glasfaser-->UDM-Pro;
    UDM-Pro-->MacMini;
    MacMini-->Multipass-VM1;
    MacMini-->Multipass-VM2;
    MacMini-->Multipass-VM3;
    MacMini-->Multipass-VM4;
    MacMini-->Multipass-VM5;
    MacMini-->Multipass-VM6;
    MacMini-->Multipass-VM7;
    MacMini-->Multipass-VM8;
    MacMini-->Multipass-VM9;
    Multipass-VM7-->HaProxy
    HaProxy-->K8s-Controlplane1;
    HaProxy-->K8s-Controlplane2;
    HaProxy-->K8s-Controlplane3;
    HaProxy-->K8s-Worker1;
    HaProxy-->K8s-Worker2;
    HaProxy-->K8s-Worker3;

    subgraph Cluster01;
        subgraph Controlplanes;
            Multipass-VM1-->K8s-Controlplane1;
            Multipass-VM2-->K8s-Controlplane2;
            Multipass-VM3-->K8s-Controlplane3;
        end;
        subgraph Worker;
            Multipass-VM4-->K8s-Worker1;
            Multipass-VM5-->K8s-Worker2;
            Multipass-VM6-->K8s-Worker3;
        end;
    end;

    subgraph Cluster02;
        Multipass-VM8-->K8sdev-Controlplane1;
        Multipass-VM9-->K8sdev-Worker1;
    end;

    NFS-Server-->K8s-Controlplane1;
    NFS-Server-->K8s-Controlplane2;
    NFS-Server-->K8s-Controlplane3;
    NFS-Server-->K8s-Worker1;
    NFS-Server-->K8s-Worker2;
    NFS-Server-->K8s-Worker3;

    NFS-Server-->K8sdev-Controlplane1;
    NFS-Server-->K8sdev-Worker1;
</pre>

Alle VMs sind auf einem MacMini mit Multipass virtualisiert. Cluster01 hat 3 Controlplanes und 3 Worker. Der Cluster02 ist für DEV und Tests, hat aktuell nur ein Controlplane und einen Worker. Beide Cluster können aber bei Bedarf jederzeit vergrössert werden. Aktuell sind sogar noch 3 weitere Cluster aktiv, da aktuell viel getestet und ausprobiert wird. 

Auf der Multipass VM 07 ist ein HA-Proxy installiert, der sich um die Loadbalancing Frontends und Backends kümmert. 
Hier ist auch je Cluster eine VIP für die Kubernetes API. Über die kann die Kubernetes API über eine IP angesprochen werden und landet dann auf eines der Controlplanes. 

<pre class="mermaid">
graph TD;
    subgraph API-k8s;
        HaProxyAPI[HAProxy, VIP: 192.168.67.200, Port 6443]-->K8s-Controlplane1;
        HaProxyAPI[HAProxy, VIP: 192.168.67.200, Port 6443]-->K8s-Controlplane2;
        HaProxyAPI[HAProxy, VIP: 192.168.67.200, Port 6443]-->K8s-Controlplane3;
    end;
</pre>

Das gleiche gilt für die HTTP Zugriffe über Port 80, hier wird aber eine andere VIP genommen. Der Grund ist, das hier über Metallb andere oder auch weitere IPS für LoadbalancerIP oder einem weiteren IngressController im Kubernetes genutzt werden könnten. Aber auch für Migrationen des kompletten Clusters oder einzelner Services auf andere Cluster. 

<pre class="mermaid">
graph TD;
    subgraph Ingress-http;
        HaProxyhttp[HAProxy, VIP: 192.168.67.17, Port 80]-->K8s-Controlplane1-http;
        HaProxyhttp[HAProxy, VIP: 192.168.67.17, Port 80]-->K8s-Controlplane2-http;
        HaProxyhttp[HAProxy, VIP: 192.168.67.17, Port 80]-->K8s-Controlplane3-http;
        HaProxyhttp[HAProxy, VIP: 192.168.67.17, Port 80]-->K8s-Worker1-http;
        HaProxyhttp[HAProxy, VIP: 192.168.67.17, Port 80]-->K8s-Worker2-http;
        HaProxyhttp[HAProxy, VIP: 192.168.67.17, Port 80]-->K8s-Worker3-http;
    end;
</pre>

Auch für https sind im HA Proxy die Controlplanes und Worker des Clusters eingetragen. Die IP 192.168.67.17 ist die gleiche, da ich hier keine Trennung vornehme. Es währe aber möglich HTTPs Traffik nur auf bestimmte Nodes zu leiten oder gar schon vorne am HA-Proxy auf eine dedizierte IP. Da kommt es immer auf das Setup und den Service an den man bereitstellen möchte. Da Option besteht aber und man könnte auf stärkere Hardware TLS laufen lassen und auf den schwachen Nodes HTTP-Only. 

<pre class="mermaid">
graph TD;
    subgraph Ingress-https;
        HaProxyhttps[HAProxy, VIP: 192.168.67.17, Port 443]-->K8s-Controlplane1-https;
        HaProxyhttps[HAProxy, VIP: 192.168.67.17, Port 443]-->K8s-Controlplane2-https;
        HaProxyhttps[HAProxy, VIP: 192.168.67.17, Port 443]-->K8s-Controlplane3-https;
        HaProxyhttps[HAProxy, VIP: 192.168.67.17, Port 443]-->K8s-Worker1-https;
        HaProxyhttps[HAProxy, VIP: 192.168.67.17, Port 443]-->K8s-Worker2-https;
        HaProxyhttps[HAProxy, VIP: 192.168.67.17, Port 443]-->K8s-Worker3-https;
    end;
</pre>

# VIP IPs mit Keepalived 

Die VIPs werden auf der VM07 (HAProxy) per keepdalived hochgefahren. 

```
global_defs {
    enable_script_security
    script_user root
}

vrrp_script chk_haproxy {
    script 'killall -0 haproxy'
    interval 2
}

vrrp_instance cluster01-api-vip {
    interface enp0s1
    state MASTER
    priority 200
    virtual_router_id 51

    virtual_ipaddress {
        192.168.64.200/24
    }

    }
vrrp_instance cluster01-ingress-vip {
    interface enp0s1
    state MASTER
    priority 200
    virtual_router_id 52

    virtual_ipaddress {
        192.168.67.17/24
    }

    }
```

Die Instancen werden einfach hinzugefügt und dabei die `virtual_router_id` hochgezählt. Die müssen eindeutig sein. 
Nach einem Restart von Keepalived sind nach ein paar Sekunden die zusätzlichen IPs auf dem Interface hochgefahren: 

```
systemctl restart keepalived.service

2: enp0s1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 52:54:00:73:a6:5c brd ff:ff:ff:ff:ff:ff
    inet 192.168.64.92/24 metric 100 brd 192.168.64.255 scope global dynamic enp0s1
       valid_lft 2061sec preferred_lft 2061sec
    inet 192.168.67.17/24 scope global enp0s1
       valid_lft forever preferred_lft forever
    inet 192.168.64.200/24 scope global secondary enp0s1
       valid_lft forever preferred_lft forever
```

# HAProxy Konfiguration

Jetzt kann der HAProxy konfiguriert werden und die einzelnen Services auf die VIPs gebunden werden. 

## Global und Default 

Der Global und Default Block in der haproxy.cfg sieht ao aus:

```
global
        log /dev/log    local0
        log /dev/log    local1 notice
        chroot /var/lib/haproxy
        # stats socket /run/haproxy/admin.sock mode 660 level admin
        stats socket /var/run/haproxy.sock mode 600 level admin
        stats timeout 30s
        user haproxy
        group haproxy
        daemon

        # Default SSL material locations
        ca-base /etc/ssl/certs
        crt-base /etc/ssl/private

        # See: https://ssl-config.mozilla.org/#server=haproxy&server-version=2.0.3&config=intermediate
        ssl-default-bind-ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384
        ssl-default-bind-ciphersuites TLS_AES_128_GCM_SHA256:TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256
        ssl-default-bind-options ssl-min-ver TLSv1.2 no-tls-tickets

defaults
        log     global
        mode    http
        option  httplog
        option  dontlognull
        timeout connect 5000
        timeout client  50000
        timeout server  50000
        errorfile 400 /etc/haproxy/errors/400.http
        errorfile 403 /etc/haproxy/errors/403.http
        errorfile 408 /etc/haproxy/errors/408.http
        errorfile 500 /etc/haproxy/errors/500.http
        errorfile 502 /etc/haproxy/errors/502.http
        errorfile 503 /etc/haproxy/errors/503.http
        errorfile 504 /etc/haproxy/errors/504.http
```

## Das erste Frontend 

Anschliessend werden die Frontends konfiguiert. Frontends sind bei HAProxy die Services, die auf einer IP auf einem bestimmten Port listenen und auf ein Backend verweisen. Zu Backends dann weiter unten mehr. 

Als erstes das Frontend für die Kubernetes API auf Port 6443. 

```
frontend k3s-api-frontend
    bind 192.168.64.200:6443
    mode tcp
    option tcplog
    default_backend k3s-backend
```

##  Das erste Backend

Das Backend dazu enthält jetzt aber nicht nur einen Eintrag, sondern bekommt alle Controlplanes des Clusters. 

```
backend k3s-backend
    mode tcp
    option tcp-check
    balance roundrobin
    default-server inter 10s downinter 5s
    server controlplane1 192.168.64.91:6443 check
    server controlplane2 192.168.64.92:6443 check
    server controlplane3 192.168.64.93:6443 check
```

Damit sind die drei Controlplanes eingetragen und ein Aufruf auf `192.168.64.200:6443` gibt die Ausgabe der Kubernetes API zurück. Diese IP und den Port kann auch so in die Kubeconfig eingetragen werden. 
Das hat den Vorteil das wir jederzeit einen oder mehrere Controlplanes offline nehmen können und weiter auf die API zugreifen können. 

## HTTP/HTTPS Frontend/Backend

Das Frontend für HTTP und HTTPs sieht genau so aus. Hier einmal wie es für einen Cluster sein könnte. Anschliessend aber mit weiteren Optionen, da wir ein Frontend haben wollen, welches auf mehrere Backends zeigt. 
Wir wollen ja extern nur eine IP benutzen. Gerade an einem Anschluss fürs HomeLab mit DSL oder Glasfaser hat man ja nur eine IP. Und IPv6 halten ja zu viele noch für Teufelswerk. 


```
frontend k8s_http_lb_frontend
    bind 192.168.64.17:80
    mode http
    option httplog
    log-format "%ci:%cp [%t] %ft %b %s %TR/%Tw/%Tc/%Tr/%Tt %ST %B %CC %CS %tsc %ac/%fc/%bc/%sc/%rc %sq/%bq %hr %hs %{+Q}r"
    default_backend k8s_http_lb_backend

frontend k8s_https_lb_frontend
    bind 192.168.64.17:443 ssl crt /etc/haproxy/certs/ alpn h2,http/1.1
    mode http
    option httplog
    log-format "%ci:%cp [%t] %ft %b %s %TR/%Tw/%Tc/%Tr/%Tt %ST %B %CC %CS %tsc %ac/%fc/%bc/%sc/%rc %sq/%bq %hr %hs %{+Q}r"
    default_backend k8s_https_lb_backend
```

Und die Backends dazu: 

```
backend k8s_http_lb_backend
    mode tcp
    option tcp-check
    balance roundrobin
    default-server inter 10s downinter 5s
    server controlplane1 192.168.64.91:80 check
    server controlplane2 192.168.64.92:80 check
    server controlplane3 192.168.64.93:80 check
    server worker1 192.168.64.94:80 check
    server worker2 192.168.64.95:80 check
    server worker3 192.168.64.96:80 check

backend k8s_https_lb_backend
    mode http
    balance roundrobin
    default-server inter 10s downinter 5s
    server controlplane1 192.168.64.91:443 check
    server controlplane2 192.168.64.92:443 check
    server controlplane3 192.168.64.93:443 check
    server worker1 192.168.64.94:443 check
    server worker2 192.168.64.95:443 check
    server worker3 192.168.64.96:443 check
```

Damit wäre das ganze lauffähig und wir hätten die API, HTTP und HTTPs Ingress erreichbar. 

## Weitere Backends

Aber es sind ja mehrere Cluster und die externe IP am Anschluss soll ja auf den HAProxy geleitet werden und dann weiter an den entsprechenden Cluster. 

Cluster01 hat die Domain k8s01.example.net und entsprechende Subdomains web1.k8s01.example.net, web2.k8s01.example.net usw. 

Cluster01 DEV hat die Domain k8s01-dev.example.net und entsprechende Subdomains web1.k8s01-dev.example.net, web2.k8s01-dev.example.net usw. 

Cluster02 DEV hat die Domain k8s02-dev.example.net und entsprechende Subdomains web1.k8s02-dev.example.net, web2.k8s02-dev.example.net usw. 

Das ganze kann beliebig weiter getrieben werden und weitere Cluster hinzugüfgt werden. 

## Match Domains 

Damit das ganze funktioniert müssen wir als erstes das Frontent für HTTP und HTTPs erweitern. 

```
frontend k8s_http_lb_frontend
    bind 192.168.64.17:80
    mode http
    option httplog
    log-format "%ci:%cp [%t] %ft %b %s %TR/%Tw/%Tc/%Tr/%Tt %ST %B %CC %CS %tsc %ac/%fc/%bc/%sc/%rc %sq/%bq %hr %hs %{+Q}r"
    acl is_k8s01_dev hdr(host) -m reg ^[a-zA-Z0-9-]+\.k8s01-dev\.example\.net$
    acl is_k8s02_dev hdr(host) -m reg ^[a-zA-Z0-9-]+\.k8s02-dev\.example\.net$
    use_backend http_k8s01_dev_backend  if is_k8s01_dev
    use_backend http_k8s02_dev_backend  if is_k8s02_dev
    default_backend k8s_http_lb_backend

frontend k8s_https_lb_frontend
    bind 192.168.64.17:443 ssl crt /etc/haproxy/certs/ alpn h2,http/1.1
    mode http
    option httplog
    log-format "%ci:%cp [%t] %ft %b %s %TR/%Tw/%Tc/%Tr/%Tt %ST %B %CC %CS %tsc %ac/%fc/%bc/%sc/%rc %sq/%bq %hr %hs %{+Q}r"
    acl is_k8s01_dev hdr(host) -m reg ^[a-zA-Z0-9-]+\.k8s01-dev\.example\.net$
    acl is_k8s01_dev req.ssl_sni -m reg ^[a-zA-Z0-9-]+\.k8s01-dev\.example\.net$
    acl is_k8s02_dev hdr(host) -m reg ^[a-zA-Z0-9-]+\.k8s02-dev\.example\.net$
    acl is_k8s02_dev req.ssl_sni -m reg ^[a-zA-Z0-9-]+\.k8s02-dev\.example\.net$

    use_backend https_k8s01_dev_backend if is_k8s01_dev
    use_backend https_k8s02_dev_backend if is_k8s02_dev

    default_backend k8s_https_lb_backend
```

## Die zusätzlichen Backends 

Die beiden Backends der weiteren Cluster anlegen:

```
backend http_k8s01_dev_backend
    mode tcp
    option tcp-check
    balance roundrobin
    default-server inter 10s downinter 5s
    server controlplane1 192.168.64.100:80 check
    server worker1 192.168.64.101:80 check

backend https_k8s01_dev_backend
    mode http
    balance roundrobin
    default-server inter 10s downinter 5s
    server controlplane1 192.168.64.100:443 check
    server worker1 192.168.64.101:443 check

backend http_k8s02_dev_backend
    mode tcp
    option tcp-check
    balance roundrobin
    default-server inter 10s downinter 5s
    server controlplane1 192.168.64.110:80 check
    server worker1 192.168.64.111:80 check

backend https_k8s02_dev_backend
    mode http
    balance roundrobin
    default-server inter 10s downinter 5s
    server controlplane1 192.168.64.110:443 check
    server worker1 192.168.64.111:443 check
   
```

Damit würde das ganze jetzt mit dem Production, Dev 1 und Dev 2 funktionieren und die Requests werden je nach Domain auf den richtigen Cluster geleitet. 

# TLS-Termination mit Letsencrypt

Wer genau gegeuckt hat sieht in den Frontend Blöcken bei der Option `bind` beim Frontent für https `ssl crt /etc/haproxy/certs/ alpn h2,http/1.1`. Der HAProxy macht die TLS-termination, also auch das handling der HTTPs Verbindungen, inklusive der Zertifikaten. Ich habe meine Domains bei Cloudflare und habe dort einen API-Key angelegt, mit dem ich DNS Zonen lesen und bearbeiten kann. 

Cludflare DNS auch aus dem Grund, weil man so auch Wildcard Zertifikate erstellen kann, was die ganze Sache sehr viel einfacher macht. Man muss nicht für jeden neuen Host-Eintrag ein neues Zertifikat erstellen. 

acme.sh installieren: 

```
curl https://get.acme.sh | sh -s email=youraddress@example.com
```

Anscliessend noch den API-Key von Cloudflare mit dem passenden Account in die Env Vars per expose packen:

```
export CF_Key="763eac4f1bcebd8b5c95e9fc50d010b4"
export CF_Email="alice@example.com"
```

Die benötigten Zertifikate erstellen: 

```
acme.sh --issue --dns dns_cf -d k8s01.example.net -d '*.k8s01.example.net' --server letsencrypt
acme.sh --issue --dns dns_cf -d k8s01-dev.example.net -d '*.k8s01-dev.example.net' --server letsencrypt
acme.sh --issue --dns dns_cf -d k8s02-dev.example.net -d '*.k8s02-dev.example.net' --server letsencrypt
```

Die Zertifikate sind damit auch erledigt und man muss diese jetzt nur noch für HaProxy bereitstellen:

```
cat /root/.acme.sh/k8s01.example.net_ecc/fullchain.cer /root/.acme.sh/k8s01.example.net_ecc/k8s01.example.net.key > /etc/haproxy/certs/k8s01.example.net.pem
cat /root/.acme.sh/k8s01-dev.example.net_ecc/fullchain.cer /root/.acme.sh/k8s01-dev.example.net_ecc/k8s01-dev.example.net.key > /etc/haproxy/certs/k8s01-dev.example.net.pem
cat /root/.acme.sh/k8s02-dev.example.net_ecc/fullchain.cer /root/.acme.sh/k8s02-dev.example.net_ecc/k8s02-dev.example.net.key > /etc/haproxy/certs/k8s02-dev.example.net.pem
```

Wenn man noch mehr Cluster und Domains hat muss man die beiden letzten Schritte auch für diese Domains ausführen. 

`systemctl restart haproxy.service` und prüfen ob alles ok ist. Dazu kann man auch die Stats Option im HAProxy aktivieren und auch da mit prüfen ob alle Backend eingetragen sind. 

# HAProxy Stats 

```
listen stats
    bind *:9000
    stats enable
    stats uri /stats
    stats refresh 5s
    stats realm HomeLab\ Statistics
    stats auth youruser:yoursecretpassword
```

Die IP mit Port 9000 im Browser öffnen und die Stats Seite von HAProxy zeigt alle Frontends und Backends an. 

# Trotz TLS in HAProxy kann Cert-Manager genutzt werden

Hier wird das SSL-Offloading vom HAProxy gemacht. Das ist nötig, damit HAProxy auch in die Verbindungen gucken kann und anhad ` Host` im Header unterscheiden kann, zu welchem Backend der Request gesendet werden muss. 

Ich teste und nutze in Kubernetes auch den Cert-Manager. Das möchte ich auch weiter nutzen. Das ist aber auch kein Problem, da ein Ingress mit TLS, über den Issuer/ClusterIssuer genau so den Cert-Manager triggert wie bis her. 
Der legt dann den `http-solver` an und den passenden Ingress, der über den Pfad auf den `http-solver` zeigt.

Der Request geht dann ganz normal an das externe Interface, zum MacMini, HAProxy und wird dann per HTTP an den richtigen Cluster geleitet. 

Sollte ich also den HAProxy für Wartungen offline nehmen müssen, dann ist für mich aktuell nur der Production Cluster relevant. Und dafür kann ich dann einfach in der UDM-Pro einfach auch direkt auf den Cluster leiten lassen. 
Die Zertifikate sind dort dann auch schon vorhanden und es gibt keine Probleme oder Ausfälle. 

# Weitere Anregungen zu HAProxy 

Wer nur einen Cluster betreibt muss den SSL Part nicht machen. In dem Fall kann man auch weiter alle Zertifikate nur im Kubernetes mit dem Cert-Manager erstellen und stellt die Verbindungen im Backend, bei der Option `mode`, einfach auf `tcp`. Das `httplog` muss dann auch entfernt werden. 

Für Wartungen an Kubernetes Nodes können diese auch einfach ausgetragen werden oder hinter dem `check` trägt man einfach `disabled` zusätzlich ein und macht einen reload. Das setzt setzt den Server auf Maintainance. 

HAProxy ist sehr mächtig. Zum Beispiel wird hier einfach nur Roundrobin genutzt. Es Loadbalancer nach roundrobin, leastconn, source IP oder weiteren Loadbalancing-Algorithmen genutzt werden. Genau so kann Sticky Sessions mit Sticky Tables genutzt werden. 

Die Beispiele oben leiten bestimmte Sub-Sub-Domains auf Backends um, da auf die Subdomain gematched wird. 
Wer jetzt überlegt: 

> "Ja, aber was mache ich wenn ich alle Domains umziehen möchte, aber ich kann nicht alle auf einmal umziehen?" 

Das geht auch sehr einfach und war auch am Anfang meine Überlegung es komplett so zu machen. Dann hätte ich aber jeden Hostname (Alos Sub-Sub-Domain) beim anlegen händisch im HAProxy hinzufügen müssen. Das will ich nicht, da ich einfach schnell Dinge im Kubernetes anlegen will und sie dann sofort funktionieren. 

Aber man kann anstatt einem Regex auch mit Domainlisten arbeiten. Oder auch HAProxy Maps einsetzen:

Dafür eine Datei anlegen `/etc/haproxy/maps/hosts.map` 

```
#domainname                           backendname
nginx.k8s02-dev.example.net           https_k8s02_dev_backend
nginx2.k8s02-dev.example.net          https_k8s02_dev_backend
nginx.k8s01-dev.example.net           https_k8s01_dev_backend
ngin2.k8s01-dev.example.net           https_k8s01_dev_backend
nginx.k8s01.example.net               https_k8s01_backend
# [...]
api.k8s01.example.net                 https_k8s01_backend
```

Die Frontend Konfiguration dafür würde dann so ausehen:

```
frontend default
   bind :80
   use_backend %[req.hdr(host),lower,map_dom(/etc/haproxy/maps/hosts.map,be_default)]
```

Plus ggf. die IP fürs Bind, Loadbalancing-Algorithmen und weitere Optionen. 
Domain und Backend Listen können auch im Zusammenspiel mit ALCs genutzt werden. 
Es können auch bestimmte Routen auf andere Backend umgeleitet werden:

`/etc/haproxy/maps/routes.map`

```
/api    be_api
/login  be_auth
```

Frontend: 

```
frontend www
  bind :80
  use_backend %[path,map_beg(/etc/haproxy/maps/routes.map,be_default)]
```

Man sieht, HAProxy ist sehr mächtig und kann sehr viel. Was heute nur noch auf HAProxy Installationen zum einsatz kommt, die man eigentlich abschalten sollte, ist die Weiche für Desktop und Mobile Webseiten. Ja das haben einige früher so gemacht und ich kenne noch ein paar die das immer noch einsetzen müssen, da deren CMS weiter betrieben werden muss, weil keiner die Eier hat auch mal alten Schrott abzuschalten. ;-) 

