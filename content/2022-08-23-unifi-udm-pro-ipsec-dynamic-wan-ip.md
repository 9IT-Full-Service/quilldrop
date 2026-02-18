---
title: "Unifi UDM PRO IPSec dynamische WAN-IP"
date: 2022-09-23 15:00:00
update: 2022-09-23 15:00:00
author: ruediger
cover: "/images/cat/technik.webp"
tags:
  - ubiquiti
  - unifi
  - udmpro
  - VPN
  - Site2Site
  - WAN
  - IP
preview: "IPSec Site2Site Tunnel mit der Unifi Dream Machine Pro ist sehr einfach eingerichtet.
Bei einer statischen WAN-IP hat man auch keine Probleme damit. Wer aber eine UDM Pro mit einer dynamischen IP-Adresse benutzt wird schnell merken wenn der Tunnel nicht läuft, wenn durch einen Reconnect am Anschluss eine neue IP gesetzt ist."
categories: 
  - Technik
showToc: false
hide: false
type: post
draft: false
---

IPSec Site2Site Tunnel mit der Unifi Dream Machine Pro ist sehr einfach eingerichtet.
Bei einer statischen WAN-IP hat man auch keine Probleme damit. Wer aber eine UDM Pro mit einer dynamischen IP-Adresse benutzt wird schnell merken wenn der Tunnel nicht läuft, wenn durch einen Reconnect am Anschluss eine neue IP gesetzt ist.

Die UDM ändert leider nicht die Config automatisch ab. Aber auch auf der Gegenseite muss die Config angepasst werden.

Das ganze erledige ich jetzt mit Ansible. Ändert sich die IP wird die aktuelle IP per Ninja Template in die neue Config gegossen und auf den Remote-IP-Server kopiert.
Für die UDM kommt auch ein Script zum Einsatz. Es wird per Ansible kopiert und anschliessend auch ausgeführt.

Der Ansible Task für den Remote IPSec-Server holt sich erst einmal die aktuelle IP.
Damit wird dann das Template gefüllt, also an der Stelle für die LeftID (rightid=...) und auf den Remote Server kopiert.

```
---

- name: Get JSON from the Interwebs
  uri: url="https://ip.tytik.cloud/json" return_content=yes method="GET" body_format="json"
  register: json_response
  delegate_to: localhost

- name: set ip var
  set_fact: ip="{{ (json_response.content|from_json)['data'] }}"


- name: ipsec.conf
  template:
    src: ipsec.conf.j2
    dest: /etc/ipsec.conf
    owner: root
    group: root
    mode: 0640
```

Bei der UDM Pro wird ein Shell Script kopiert, welches auch die aktuelle IP Adresse per Curl abruft und dann per sed die alte IP gegen die neue ersetzt. Anschliessend wird IPSec neugestartet und der Tunnel aufgebaut.

```
---

- name: Copy file with owner and permissions
  copy:
    src: ip_wan_ip.sh
    dest: /root/ip_wan_ip.sh
    owner: root
    group: root
    mode: '0544'

- name: Change Remote ip
  command: /root/ip_wan_ip.sh
```

Da die IP bei `rightid` gesetzt ist und bei `right` der Subdomain und dieser auf die dynamische IP Zeigt, wird die Subdomain auch direkt noch aktualisert. Dazu habe ich ein Update Script, mit dem im BIND DNS Server der A-Record für diese Subdomain aktualisiert wird.

```
- name: Change Domain Record
  command: /root/nsupdate.sh -z domain.tld -t A -h ipsec -v {{ ip }}
  delegate_to: localhost
```

Damit kann regelmässig überprüft werden ob sich die IP geändert hat und es werden auf allen Seiten die nötigen Änderungen gemacht um den Tunnel wieder aufzubauen.

Das umständliche DNS Update und das herumgeklicke in der UDM Pro Weboberfläche fällt weg. Ausserdem kann das dann auch automatisiert gemacht werden und man muss sich auch nicht mehr drum kümmern.

Die Scripte und Ansible Roles sind fertig. Da noch ein paar Sachen auf das Setup hier angepasst sind und an mehreren Stellen etwas angepasst werden müsste, ist der Code noch nicht im Git.
Das werde ich noch nachholen und ein Update hier hinterlassen. So das es von jedem genutzt werden kann und nur an einer oder zwei Stellen dann die IPs, Hostnames, User usw. eingetragen werden müssen.
