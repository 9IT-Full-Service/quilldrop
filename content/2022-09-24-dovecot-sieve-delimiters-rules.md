---
title: "Dovecot Sieve Delimiters Rules"
date: 2022-09-24 10:00:00
lastmod: 2022-09-24 10:00:00
author: ruediger
cover: "/images/cat/technik.webp"
tags:
  - Dovecot
  - Sieve
  - Delimiters
  - Rules
preview: Ich benutze seit einigen Jahren für manche Mailaccounts Delimiter im Local-Part der Adressen. So kann ich für bestimmte Seite immer, ohne erst einen Account oder eine Weiterleitung einrichten zu müssen, extra E-Mailadressen benutzen.
categories: 
  - Technik
type: post
toc: false
hide: false
draft: false
---


### Was sind Delimiter Mailadressen?

Ich benutze seit einigen Jahren für manche Mailaccounts Delimiter im Local-Part der Adressen. So kann ich für bestimmte Seite immer, ohne erst einen Account oder eine Weiterleitung einrichten zu müssen, extra E-Mailadressen benutzen.

Die E-Mailadressen sind immer nach folgendem Schema aufgebaut.

    username+order1.order2@domain.de

ordner1 ist dabei dann z.B. shops, social usw.
ordner2 ist z.B. bei shops dann amazon, ebay usw.

So könnte es dann folgende E-Mailadressen für die einzelnen Seiten geben:

   * username+shops.amazon@domain.de
   * username+shops.ebay@domain.de
   * username+social.facebook@domain.de
   * username+social.instagram@domain.de
   * username+social.twitter@domain.de

So hat jeder Shop und jedes soziale Netzwerk und viele weitere Portale und Services immer eine eigene E-Mailadresse.

### Automatische Sortierung

Die Mail können anschliessend per Filterregeln natürlich passend einsortiert werden. Wer seine Domains und Mailadressen auf einem Server mit Sieve als Filter, kann diese Mailadressen auch automatisch in die Ordner ablegen lassen.

Dafür gibt es einen Ordner im Postfach `Autosort` in dem dann alle Mails in die entsprechenden Ordnern speichert.

```
.Autosort.shops.amazon
.Autosort.shops.ebay
.Autosort.social.facebook
.Autosort.social.instagram
.Autosort.social.twitter
```

Im Mailclient wird dann unter dem Ordner Autosort die Ordner shops und social angelegt und dort drin dann Amazon und Ebay, bzw. Facebook, Instagram und Twitter angelegt, in denen dann die Mails jeweils einsortiert werden.

{{< postimage "Autosort Folders Mailpostfach" "folsers.png" "folders.png" >}}

![Autosort Folders Mailpostfach](/images/posts/folders.webp)


### Sieve Filterregeln

Damit das ganze funktioniert muss natürlich auch ein passender Filter im Sieve hinterlegt werden.

```
require ["reject","fileinto","imap4flags","body","vacation","copy","variables","regex","envelope"];

# rule:[autodelemiter]
if header :regex "Delivered-To" "username\\+([^.]*)\\.?([^.]*)\\.?([^.]*)\\.?([^.]*)\\.?([^.]*)@.*$" {
  if string :is "${1}" "" {} else { set :lower "part1" "${1}"; }
  if string :is "${2}" "" {} else { set :lower "part2" "${2}"; }
  if string :is "${3}" "" {} else { set :lower "part3" "${3}"; }
  if string :is "${4}" "" {} else { set :lower "part4" "${4}"; }
  if string :is "${5}" "" {} else { set :lower "part5" "${5}"; }
  if string :is "${6}" "" {} else { set :lower "part6" "${6}"; }

  set "targetfolder" "";
  if string :is "${part1}" "" {} else { set "targetfolder" "${part1}"; }
  if string :is "${2}" "" {} else { set "targetfolder" "${targetfolder}/${part2}"; }
  if string :is "${3}" "" {} else { set "targetfolder" "${targetfolder}/${part3}"; }
  if string :is "${4}" "" {} else { set "targetfolder" "${targetfolder}/${part4}"; }
  if string :is "${5}" "" {} else { set "targetfolder" "${targetfolder}/${part5}"; }
  if string :is "${6}" "" {} else { set "targetfolder" "${targetfolder}/${part6}"; }

  fileinto "Autosort/${targetfolder}";
}
```

Damit können die Mailadressen einfach angegeben werden und sie werden automatisch sortiert. Das ganze ist hier für bis zu 6 Ordnern ausgelegt. Damit wäre also folgende Mailadresse möglich `username+ordner1.ordner2.order3.order4.order5.order6@domain.tld` möglich.

Die Ordner dazu hätten dann folgende Struktur im Postfach:
```
Autosort
+-ordner1
  +-ordner2
    +-order3
      +-order4
        +-order5
          +-order6
```

Ich selbst habe bis jetzt glaube ich maximal 4 Ordner benutzt. Aber so ist halt noch Luft nach oben. ;-)

Das ist der Grund wieso ich 2789 Ordner in meinen Mailpostfächern habe. Und 97% habe ich nicht einmal selbst angelegt. Das passiert alles automatisch und die die Mails landen auch noch automatisch in diesen Ordnern.

So sind auch nur die wichtigsten E-Mail in der Inbox.

### Wieso Delimiter Adresse noch sinnvoll sind

Es kam in den letzten Jahren mehrfach zu Datenschutzvorfällen bei einigen Anbietern und Portalen. Werden solche E-Mailadresse durch einen Hack einer Seite abgezogen und eine Spamwelle kommt, schaltet man diese einzelne Delimiter Mailadresse einfach mit einer Filterregel ab. Also einfach löschen.
In ein paar Fällen war den Betreibern auch nicht bewusst das Daten von dem Portal abgezogen wurden. Durch die Erklärung und teils sogar recht kryptischen Mailadressen von mir war ihnen schnell klar das diese Spam/Pishing E-Mails nicht durch Zufall bei mir angekommen sind. Es wurde dann überprüft und durch den Hinweis konnten manche Vorfälle dann auch bestätigt und andere Benutzer direkt gewarnt werden.
