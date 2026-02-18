---
title: 'Quagga BGP prefix-list'
date: 2019-08-03 01:03:26
update: 2019-08-03 01:03:26
author: ruediger
cover: "/images/posts/2019/08/05/network.webp"
tags:
    - Netzwerk
    - IP
    - BGP
    - Routing
    - Internet
preview: "*Notiz an mich, um nich noch einmal suchen zu müssen*  Um im BGP manche Netze nicht zu erlauben:"
categories: 
    - Internet
toc: false
hide: false
type: post
---

*Notiz an mich, um nich noch einmal suchen zu müssen* :point_up:

Um im BGP manche Netze nicht zu erlauben:

<!--more-->
```
router bgp 65001
 bgp router-id 10.10.10.1
 network 10.101.0.0/16
 neighbor 10.11.0.1 remote-as 65002
 neighbor 10.11.0.1 description 65002
 neighbor 10.11.0.1 prefix-list icvpn4 in
 neighbor 10.11.0.1 prefix-list icvpn4 out
!
ip prefix-list icvpn4 description *** ICVPN prefix-list for internal and public IP address space ***
ip prefix-list icvpn4 seq 20 deny 10.101.0.0/16 le 24
```

Damit ist 10.101.0.0/16 und /24er aus dem Block nicht mehr erlaubt.
