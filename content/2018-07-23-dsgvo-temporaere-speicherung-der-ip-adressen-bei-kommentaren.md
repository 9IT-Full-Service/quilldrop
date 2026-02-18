---
title: 'DSGVO: temporäre Speicherung der IP Adressen bei Kommentaren'
date: 2018-07-23 12:15:41
update: 2018-07-23 12:15:41
author: ruediger
cover: "/images/cat/internet.webp"
tags:
    - Blog
    - DSGvO
    - Internet
    - IP
    - IP-Adresse
    - Kommentare
    - Spam
    - Wordpress
preview: "Mit der DSGVO ist die IP-Adresse zu schützen. Vorher wurde die IP auch nur temporär gespeichert."
categories: 
    - Technik
toc: false
hide: false
draft: false
type: post
---

Mit der DSGVO ist die IP-Adresse zu schützen. Vorher wurde die IP auch nur temporär gespeichert. Nach der Freischaltung eines Kommentars wurden die IP Adressen bei den Kommentaren gelöscht. Mit der DSGVO habe ich einfach mal die IP-Adresse bei den Kommentaren komplett eliminiert und jetzt knapp 2 Monate versucht ohne auszukommen. Ergebnis: Scheisse. Früher war es so das Spammer versucht haben ihre Scheisse los zu werden. Dabei sind sie meistens wie folgt vorgegangen:
<!--more-->
1.  Kommentar gesendet. (z.B. von 123.123.123.123)
2.  Minuten oder Stunden später kam dann der nächste Spam-Kommentar. wieder mit der gleichen IP Adresse.
3.  Einige Zeit später ein weiterer Kommentar.
4.  Spätestens jetzt habe ich dann die IP genommen, gegen ein Tool von mir geworfen und die IP landete in direkt in der Firewall in einer Blacklist.
5.  Ruhe
6.  Ohne geht das jetzt weiter bis dann 30 und mehr Spam-Kommentare das Blog zu müllen.

Daher werden hier jetzt bei den Kommentaren die IP-Adressen wieder gespeichert und nach Freigabe oder Löschung der Kommentare die IP-Adressen aus der Datenbank entfernt. Das ganze ist natürlich in der Datenschutzerklärung abgedeckt.

> ### Kommentare und Beiträge
>
> Wenn Nutzer Kommentare oder sonstige Beiträge hinterlassen, können ihre IP-Adressen auf Grundlage unserer berechtigten Interessen im Sinne des Art. 6 Abs. 1 lit. f. DSGVO für 7 Tage gespeichert werden.

Zusätzlich werden Besucher beim Schreiben der Kommentare auf die Speicherung der IP-Adresse hingewiesen. 

Icons made by [Freepik](http://www.freepik.com "Freepik") from [www.flaticon.com](https://www.flaticon.com/ "Flaticon") is licensed by [CC 3.0 BY](http://creativecommons.org/licenses/by/3.0/ "Creative Commons BY 3.0")
