---
title: "WiFi QRCode mit qrencode erstellen"
date: 2019-08-10 20:18:49
update: 2019-08-10 20:18:49
author: ruediger
cover: "/images/posts/2019/08/10/wifi.webp"
tags:
    - QRCode
    - qrencode
    - brew
    - apt
    - WLAN
    - WiFi
preview: "Um ohne langes tippen in das WiFi Netz zu kommen erstelle ich immer QRCodes. Ich erstelle sie mit qrencode."
categories: 
    - Technik
toc: false
hide: false
type: post
---

# QRCode für WiFi Zugänge erstellen

Um ohne langes tippen in das WiFi Netz zu kommen erstelle ich immer QRCodes.
Ich erstelle sie mit qrencode.

```
# Debian
apt install -y qrencode
# MacOS
brew install qrencode
```
<!--more-->
Anschliessend kann man sehr einfach einen WiFi QRCode erstellen:

```
qrencode -o wifi-zugang.png "WIFI:S:UnserWlan;T:WPA2;P:strengeheim12345;;" --dpi=300 -s 100
```

![Der QRCode für das WiFi, generiert mit qrencode](/images/posts/wifi-zugang.png)

So kann für neue WiFi Netze schnell der Zugang weitergegeben werden. Wir benutzen das zuhause regelmässig für die WiFi-Gastzugänge, da auch da regelmässig die Passwörter geändert werden.

# Kann auch Adressdaten in einen QRCode packen

Addresse.txt
```
BEGIN:VCARD
VERSION:3.0
N:Nachname, Vorname
ORG:nachname.de
TITLE:Webmaster und Author
EMAIL;TYPE=PREF,INTERNET: info@nachname.de
END:VCARD
```

```
qrencode -o vcard-low.png < adresse.txt
```

![VCard QRCode](/images/posts/vcard-low.png)

```
qrencode -l H -o vcard.png < adresse.txt
```


![VCard mit mehr Fehlerkorrektur](/images/posts/vcard.png)
