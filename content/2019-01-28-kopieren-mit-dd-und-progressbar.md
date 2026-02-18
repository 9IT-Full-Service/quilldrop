---
title: "Kopieren mit dd und Progressbar"
date: 2019-01-28 19:07:53
update: 2019-01-28 19:07:53
author: ruediger
cover: "/images/cat/technik.webp"
tags:
    - Boot
    - CleanInstall
    - Installation
    - Linux
    - MacOS
    - Partition
    - Restore
    - Technik
    - USB-Stick
preview: "Wer einen Rescue USB Stick erstellen will kann mit Hilfe von dd das .iso auf den Stick schreiben."
categories: 
    - Technik
toc: false
hide: false
type: post
---

Wer einen Rescue USB Stick erstellen will kann mit Hilfe von dd das .iso auf den Stick schreiben.

```
dd if=~/rescue.iso of=/dev/disk3 bs=1m
```

Dabei wird aber nicht angezeigt wie weit der Kopiervorgang ist. Bei den heutigen Betriebsystemen ist das Image aber mehrere Gigabytes groß und es dauert bei einem langsamen Stick sehr lange. Es könnte aber auch sein das ein Problem aufgetreten ist und man sieht nicht ob noch weiter kopiert wird. Daher kann man zwischen dem if (InputFile) und of (OutputFile) das Programm pv getrennt durch ein pipe setzen.
<!--more-->
```
dd if=~/rescue.iso | pv | dd of=/dev/disk3 bs=1m mb-rp# dd if=snow\ leopard\ install.iso | pv | dd of=/dev/disk3 bs=1m
731MiB 0:28:30 [   0 B/s] [             <=>            ]
```
So sieht man wie viel MB/GB schon kopiert wurden und wird nach Stunden warten nicht nervös und bricht das kopieren nicht vor lauter Verzweiflung auch noch unwissend bei 98% ab. ;-) **DMG to ISO umwandeln** Da ich selbst auch immer suchen muss halte ich hier auch gleich das dmg zu iso Image umwandeln fest.

```
hdiutil convert /path/imagefile.dmg -format UDTO -o /path/convertedimage.iso
```
