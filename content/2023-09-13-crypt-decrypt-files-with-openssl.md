---
title: Crypt/Decrypt files with openssl
date: 2023-09-13 13:12:00
author: ruediger
cover: "/images/cat/technik.webp"
tags: [Crpto, Openssl, Keys, Files, Encrypt, Decrypt]
categories: 
    - Internet
preview: "If you want to encrypt a file that should also be decrypted by several people, openssl is a very good choice. Anyone who wants to decrypt the file later must first create a key pair." 
draft: true
top: false
type: post
hide: false
toc: false
---

Wenn man eine Datei verschlüsseln möchte, die auch von mehreren Personen wieder entschlüsselt werden soll, 
bietet sich openssl sehr gut an. Jeder der die Datei später wieder entschlüsseln können soll, muss sich erst einmal ein Schlüsselpaar anlegen: 

> If you want to encrypt a file that should also be decrypted by several people, openssl is a very good choice. Anyone who wants to decrypt the file later must first create a key pair.


```
openssl req -x509 -newkey rsa:4096 -days 3650 -subj "/C=DE/ST=*/L=*/O=*/OU=*/CN=User1/" -keyout user1.key -out user1.pub
...+.......+.....+.+......+.....+...+.+..+.......+...+..+.......+...+......+............+.....+...+............+....+...+...........+......+.............+.....+.+............+..+.+.........+++++++++++++++++++++++++++++++++++++++++++++*...+......+.............+..+...+.+......+.....+...+.+...............+...............+..............+...+...+.......+.....+...+.+......+.................+.+.........+..+++++++++++++++++++++++++++++++++++++++++++++*..+..........+..+..........+.....+....+.....+...+.+..............+.....................+.+.....+.+.........+........+...+...+..........+............+...+.........+..+......+..........+........+.+...+.....+.............+...+.........+.....+...+...+......+.........+....+...+.................+++++
.........+...+.......+...+.....+...+.+...+..+...............+.+.....+++++++++++++++++++++++++++++++++++++++++++++*..+.....+.+............+.....+...+.+...+...+.........+.....+....+...........+.+..+......+..........+++++++++++++++++++++++++++++++++++++++++++++*.........+......+.+..+...+....+...+...+...+..+..........+..+...............+.......+...+............+.....+....+....................+.+.....+.........+......+......+.........+.+........................+...+...........+....+............+..............+....+.....................+..+...............+.......+..+........................+................+......+.....+.......+...........+......+.+........+............+....+..............+....+.....+.+...............+...........+...+...+.........+.+.........+.....+....+.........+......+..+....+...+...............+...+...+............+..+.+.....+..........+++++
Enter PEM pass phrase:
Verifying - Enter PEM pass phrase:
-----
openssl req -x509 -newkey rsa:4096 -days 3650 -subj "/C=DE/ST=*/L=*/O=*/OU=*/CN=User2/" -keyout user2.key -out user2.pub
...
```

Die beiden Secret User Keys bleiben bei den jeweiligen Usern, so das nur er diesen Schlüssel hat. Den Public Key kann man auch miteinander teilen, so das jeder Dateien für den anderen oder alle anderen Personen verschlüsseln kann. 

Möchte man jetzt eine Datei für beide User verschlüsseln übergibt man einfach die entsprechenden Public Key an openssl. 

> The two secret user keys remain with their respective users, so that only he has this key. The public key can be shared with each other, so that anyone can encrypt files for the other or all other people.

> If you now want to encrypt a file for both users, simply hand over the corresponding public key to openssl.

```
openssl smime -encrypt -aes256 -in test.pdf -out test.pdf.enc -binary -outform DER user1.pub user2.pub
ls -la *.pdf*
-rw-r--r--@ 1 rk  staff  240654 13 Sep 13:21 test.pdf
-rw-r--r--  1 rk  staff    6689 13 Sep 13:22 test.pdf.enc
```

Damit ist die Datei Verschlüsselt und beide User haben die Möglichkeit die Datei zu entschlüsseln. 

> With that, the file is encrypted, and both users have the ability to decrypt the file.

```
openssl smime -decrypt -in test.pdf.enc -inform DEM -inkey user1.key -out decrypt-test.pdf
ls -la *pdf*
-rw-r--r--@ 1 rk  staff  1355211 13 Sep 13:48 decrypt-test.pdf
-rw-r--r--@ 1 rk  staff  1355211 13 Sep 13:44 test.pdf
-rw-r--r--  1 rk  staff  1355942 13 Sep 13:45 test.pdf.enc
md5 test.pdf decrypt-test.pdf
MD5 (test.pdf) = 42f83d23f5d23cd474574c5bc2f93b8e
MD5 (decrypt-test.pdf) = 42f83d23f5d23cd474574c5bc2f93b8e
```

Möchet man mehrere Datein verschlüsseln nimmt man einfach `tar` und komprimiert alle Dateien. Anschliessend verschlüsselt man das komplette Tarball. 

> If you want to encrypt multiple files, simply use tar to compress all files. Afterwards, encrypt the entire tarball.