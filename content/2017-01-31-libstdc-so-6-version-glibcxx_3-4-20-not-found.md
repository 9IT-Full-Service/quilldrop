---
title: 'libstdc++.so.6: version `GLIBCXX_3.4.20'' not found'
date: 2017-01-31 21:07:03
update: 2017-01-31 21:07:03
author: ruediger
cover: "/images/cat/technik.webp"
tags:
    - Arm
    - Fhem
    - Homekit
    - Libc6
    - Linux
    - nodeJS
    - npm
    - RaspberryPi
    - Technik
preview: "Da will man einmal kurz was neues installieren und dann das:"
categories: 
    - Technik
toc: false
hide: false
type: post
---

Da will man einmal kurz was neues installieren und dann das:

```
foo@pi01:~# node -v
foo@pi01:~# node: /usr/lib/arm-linux-gnueabihf/libstdc++.so.6: version
`GLIBCXX_3.4.20' not found (required by node)
foo@pi01:~# node: /lib/arm-linux-gnueabihf/libc.so.6: version `GLIBC_2.16'
not found (required by node)
```
<!--more-->

Na dann fixt man es halt eben schnell:

```
foo@pi01:~# sudo apt-get update
foo@pi01:~# sudo apt-get install gcc-4.8 g++-4.8
foo@pi01:~# sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.6 20
foo@pi01:~# sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.8 50
foo@pi01:~# sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-4.6 20
foo@pi01:~# sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-4.8 50
```
Da zum testen ein alter RaspberryPI B im Einsatz ist:

```
foo@pi01:~# apt-get purge node
foo@pi01:~# wget https://nodejs.org/download/release/v0.10.0/node-v0.10.0-linux-arm-pi.tar.gz
foo@pi01:~# cd /usr/local
foo@pi01:/usr/local# tar xzvf ~/node-v0.10.0-linux-arm-pi.tar.gz --strip=1
```

Fertig und weiter machen mit dem was man eigentlich vor hatte. ;-)

```
foo@pi01:/usr/local# node -v
v0.10.0
foo@pi01:/usr/local# npm -v
1.2.14
```
