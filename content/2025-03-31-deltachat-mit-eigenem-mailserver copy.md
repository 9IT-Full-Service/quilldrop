---
title: 'Delta Chat mit eigenem Mailserver'
date: 2025-03-31 08:00:00
author: ruediger
cover: "/images/posts/2025/03/deltachat.webp"
featureImage: "/images/posts/2025/03/deltachat.webp"
tags: [DeltaChat, decentralized, secure, messenger, app]
categories: 
    - Messenger
    - decentralized
preview: "HomeLab mit mehreren Clustern hinter einer IP."
draft: false
top: false
type: post
hide: false
toc: false
---

![Delta Chat](/images/posts/2025/03/deltachat.webp)


[Alexander Lehmann](https://www.threads.net/@alexlehm) hat auf Threads über [Delta Chat](https://delta.chat/de/) [gepostet](https://www.threads.net/@alexlehm/post/DH0m29TtqrT). 

Ich dachte natürlich: "Noch ein Messenger?". Aber den Link musste ich anklicken und wenigstens einmal gucken. 
Ein dezentraler Messenger und sicher soll er sein. Das ist Threema, Singal und Co auch. Aber man kann sich bei einem Chatmail-Server von Delta Chat anmelden oder seinen eigenen Mailserver benutzen. 

Und da wurde es dann interessant. Ich habe einfach mal Mailaccounts von 2 Domains von mir in der App eingerichtet. 

Ich habe erst einmal einen Account direkt in der App angelegt und dabei aber den Button nicht gefunden (eher übersehen), mit dem man mit einem eigenen Mailserver verbinden kann. 

![Anderen Server verwenden](/images/posts/2025/03/deltachat_own_server.png)

Beim anlegen ist ganz unten der Punkt "Anderen Server verwenden". Auf diesen geklickt und anschliessend auf "Klassischen E-Mail-Login", erscheinen weitere Einstellungen zum Mailserver und Mailaccount. 

![Delta Chat Mail Settings](/images/posts/2025/03/deltachat-mail-settings.png)

Alle Daten vom Mailaccount eingetragen und ich könnte zwischen beiden Account Nachrichten senden. 
Die Nachrichten landeten im Posteingang und in den Chats in Delta Chat. 

Die Nachrichten wollte ich aber nicht im Posteingang haben, sondern direkt in den Ordner "DeltaChat" verschieben. Das macht es übersichtlicher und in den Settings von DeltaChat kann man auch einstellen, das nur Nachrichten aus diesem Ordner beachtet werden. 

Hier meine Sieve Mailfilter Rule: 

    # rule:[deltachat]
    if anyof (header :contains "Chat-Version" "1.0")
    {
            fileinto :create "DeltaChat";
            stop;
    }

