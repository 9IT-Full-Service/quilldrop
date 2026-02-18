---
title: 'Atom Editor hinter Proxy'
date: 2017-09-18 16:35:03
update: 2017-09-18 16:35:03
author: ruediger
cover: "/images/cat/internet.webp"
tags:
    - Free
    - Installation
    - Internet
    - Internet
    - Linux
    - MacOS
    - Programming
preview: "Da will man sich im [Atom Editor](https://atom.io) mal eben das Package script installieren, um Code zum testen direkt im Atom auszuführen, da stellt sich der Proxy mal wieder in den Weg."
categories: 
    - Internet
toc: false
hide: false
type: post
---

Da will man sich im [Atom Editor](https://atom.io) mal eben das Package "script" installieren, um Code zum testen direkt im Atom auszuführen, da stellt sich der Proxy mal wieder in den Weg.

<!--more-->

```
apm config set https-proxy http://proxy.example.com:3128
apm config set http-proxy http://proxy.example.com:3128
```

Atom wieder öffnen und Package installieren.
