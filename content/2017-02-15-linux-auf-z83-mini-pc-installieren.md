---
title: Linux auf Z83 Mini PC installieren
date: 2017-02-15 15:09:31
update: 2017-02-15 15:09:31
author: ruediger
cover: "/images/cat/technik.webp"
tags:
    - Atom
    - Backup
    - Fhem
    - Homebridge
    - Homekit
    - Installation
    - Intel
    - Linux
    - Mini-PC
    - Restore
    - Smart-Home
    - Technik
    - UEFI
    - USB-Stick
    - Z5
    - Z5-8300
    - Z83
preview: "Der Raspberry PI zuhause ist noch ein altes 1er Model. Für FHEM und Homebridge für Apple HomeKit musste eigentlich ein neuer Raspberry her. Denn Homebridge auf dem alten Schätzchen installieren klappt nicht. Vielleicht hätte man das mit einigen Stunts hinbekommen. Aber es soll ja Stabil laufen."
categories: 
    - Technik
toc: false
hide: false
type: post
---

Der Raspberry PI zuhause ist noch ein altes 1er Model. Für FHEM und Homebridge für Apple HomeKit musste eigentlich ein neuer Raspberry her. Denn Homebridge auf dem alten Schätzchen installieren klappt nicht. Vielleicht hätte man das mit einigen Stunts hinbekommen. Aber es soll ja Stabil laufen.
<!--more-->

Wie der Zufall will hatten ich mit ein paar Kollegen eine Schulung und einer der Kollegen erzählte sehr begeistert von seinem neuen Spielzeug. Er hatte sich für knapp 100 € den Z83 Mini PC gekauft und Linux drauf laufen. Hörte sich interressant an. Kurz überlegt und entschlossen den auch zu bestellen. Raspberry + Gehäuse + Netzteil + SD Card, da ist der Z83 leistungsfähiger und vom Preis daher günstiger. Am nächsten Tag kam das Teil auch schon an und abends wurde der USB Stick erstellt. Der dann prompt nicht installieren wollte. Google, Google, Google ... ein paar Sachen schnell ausprobiert. Nix. Da das Teil am TV hing konnte auch nur in den Werbepausen getestet werden. Neben bei den Sohn bespassen, füttern und ins Bett bringen, da bleibt nicht viel Zeit für Frickeln am Rechner. Ok, abbrechen und morgen weiter machen.

Am nächsten Tag auch nicht wirklich dazu gekommen. Irgendwie ist mit Arbeit, Kind und Renovierungen wenig Zeit für so etwas. Also den Kollegen im Xing angeschrieben ob er seinen Stick mitbringen kann. Damit funktionierte das dann auch prompt. Falls ich den noch einmal brauche direkt einmal eine Sicherung gemacht und auf einen Stick geschrieben. Das Image kommt jetzt auch noch ins Netz für mich und wenn es jemand auch einmal gebrauchen kann. Hier also die Info zu der Hardware, Sichern und Wiederherstellen des Sticks und zur Installation. btw: Homebridge ist direkt installiert und läuft mit Eve und FHEM zusammen perfekt. Endlich SmartHome per Siri.

Die Hardware
------------

Z83 Mini PC Intel Atom x5-Z8300 Processor (2M Cache, up to 1.84 GHz)

*   Product model:Z83 Mini PC
*   Dimension:119.5_119.5_24mm
*   Processor Number:x5-Z8350
*   Cache:2 MB
*   Instruction Set:64-bit
*   System config
*   OS:Support Windows10
*   Language :Multi -language
*   Intel CPU:Intel Atom x5-Z8350 Processor (2M Cache, up to 1.84 GHz)
*   Processor Graphics:Intel HD Graphics
*   Installed RAM:DDR3 2GB
*   System Disk:Windows(C:) 32GB
*   Ethernet: 1000Mbps LAN
*   WIFI: IEEE 802.11a/b/g/n，2.4G+5.8G
*   Bluetooth: BT 4.0
*   Antenna: Built-in antenna for WIFI
*   Expand Memory:SD Card (Support 128GB)
*   \# of Cores:4
*   \# of Threads:4
*   Processor Base Frequency:1.44 GHz
*   Burst Frequency:1.84 GHz
*   Scenario Design Power (SDP):2 W
*   Cache:2 MB
*   Instruction Set:64-bit
*   Button&ports Button:1\*Power Button
*   DC-in:1\*DC in Port
*   USB3.0:1\* Standard USB Port
*   HD:1\*HD A Type Port
*   RJ45:1\*RJ45 (1000Mbps network connection)
*   Headphone microphone:1\*Headphone microphone jack
*   SD Card:1\* SD card slot
*   USB2.0:2\* Standard USB Port

Stick Backup
============

Den Stick mit der funktionierden Installation einstecken und sichern.

```
#> dmesg| grep -A 10 "USB device" | tail -n 15
[ 2512.633014] usb 2-1.7: new high-speed USB device number 8 using ehci-pci
[ 2512.728550] usb 2-1.7: New USB device found, idVendor=090c, idProduct=1000
[ 2512.728558] usb 2-1.7: New USB device strings: Mfr=1, Product=2, SerialNumber=3
[ 2512.728562] usb 2-1.7: Product: USB DISK
[ 2512.728566] usb 2-1.7: Manufacturer: SMI Corporation
[ 2512.728569] usb 2-1.7: SerialNumber: AA04012700007674
[ 2512.729344] usb-storage 2-1.7:1.0: USB Mass Storage device detected
[ 2512.730417] scsi host8: usb-storage 2-1.7:1.0
[ 2513.734712] scsi 8:0:0:0: Direct-Access     USB      Stick 2.0 ME     1100 PQ: 0 ANSI: 0 CCS
[ 2513.735369] sd 8:0:0:0: Attached scsi generic sg2 type 0
[ 2513.738041] sd 8:0:0:0: [sdc] 1981440 512-byte logical blocks: (1.01 GB/968 MiB)
[ 2513.739134] sd 8:0:0:0: [sdc] Write Protect is off
[ 2513.739143] sd 8:0:0:0: [sdc] Mode Sense: 43 00 00 00
```
In diesem Fall ist es das Device /dev/sdb und hat 1GB.

```
#> dd if=/dev/sdb of=z83\_install.iso
1981440+0 Datensätze ein
1981440+0 Datensätze aus
1014497280 bytes (1,0 GB, 968 MiB) copied, 88,4281 s, 11,5 MB/s
```

Stick Restore
=============

Neuen Stick einstecken und das Image auf den Stick mit `dd` kopieren: Darauf achten das richtige Device zu wählen:

```
#> dmesg | grep -A 10 "Attached"
[1452.680096] sd 7:0:0:0: Attached scsi generic sg2 type 0
[1452.680659] sd 7:0:0:0: [sdb] 3948544 512-byte logical blocks: (2.02 GB/1.88 GiB)
[1452.681279] sd 7:0:0:0: [sdb] Write Protect is off
[1452.681282] sd 7:0:0:0: [sdb] Mode Sense: 00 00 00 00
[1452.681907] sd 7:0:0:0: [sdb] Asking for cache data failed
[1452.681911] sd 7:0:0:0: [sdb] Assuming drive cache: write through
[1452.823221] sdb: sdb1
[1452.826012] sd 7:0:0:0: [sdb] Attached SCSI removable disk
```

In diesem Fall ist es /dev/sdb. Wenn noch Daten auf dem Stick sind müssen diese nicht gelöscht werden. Sie werden eh gleich überschrieben. Aber man kann sie gut dafür nutzen um sicher zugehen ob es wirklich das richtige Device ist.

```
#> dd if=z83_install.iso of=/dev/sdb
1981440+0 Datensätze ein
1981440+0 Datensätze aus
1014497280 bytes (1,0 GB, 968 MiB) copied, 439,406 s, 2,3 MB/s
```
Überprüfen ob alles auf dem Stick ist:

```
#> fdisk -l /dev/sdb
Medium /dev/sdb: 1,9 GiB, 2021654528 Bytes, 3948544 Sektoren
Einheiten: sectors von 1 \* 512 = 512 Bytes
Sektorengröße (logisch/physisch): 512 Bytes / 512 Bytes
I/O Größe (minimal/optimal): 512 Bytes / 512 Bytes
Typ der Medienbezeichnung: dos
Medienkennung: 0x4455395b

Gerät      Boot  Start    Ende Sektoren Größe Id Typ
/dev/sdb1  \*         0 1482751  1482752  724M  0 Leer
/dev/sdb2       139820  144555     4736  2,3M ef EFI (FAT-12/16/32)
```

Z83 Mini PC starten und installieren
====================================

Jetzt den Stick einfach in den Z83 stecken und einschalten. Direkt nach dem drücken des Power-Buttons die Taste F7 gedrückt halten. Nach kurzer Zeit sollte die Auswahl des Boot Medium erscheinen. Dort den USB Stick auswählen und die Installation sollte starten. Have Fun.

Image des Installtions Sticks [https://drive.google.com/open?id=0B4MeHQoJL4f9bWFZa3pxTjZNUnc](https://drive.google.com/open?id=0B4MeHQoJL4f9bWFZa3pxTjZNUnc) 1 GB
