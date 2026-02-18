---
title: 'Wasser predigen, Wein trinken: Telekoms Mailserver-Doppelmoral'
date: 2025-08-15 12:00:00
update: 2025-08-15 12:00:00
author: ruediger
cover: "/images/posts/2025/08/telekom-mail-fail.webp"
featureImage: "/images/posts/2025/08/telekom-mail-fail.webp"
# images: 
#   - /images/posts/2025/08/telekom-mail-fail.webp
tags: [Telekom, mailserver, reputation, fail]
categories: 
    - Mailserver
preview: "Wasser predigen, Wein trinken: Telekoms Mailserver-Doppelmoral. Und die Telekom E-Mail Engineers können anscheinend nicht mal `dig` bedienen."
draft: false
top: false
type: post
hide: false
toc: false
---

<!-- 
![Telekom Mail Fail](/images/posts/2025/08/telekom-mail-fail.webp) 
-->


Und die Telekom E-Mail Engineers können anscheinend nicht mal `dig` bedienen.

Wenn ich Spam oder Attacken auf meine Mailserver feststelle, dann ist das Erste, was ich mache, natürlich blocken. Anschließend wird geguckt, woher der Mailserver kommt, also Provider, Serverbesitzer, Domainbesitzer. Oft kommt es ja von einer Domain, auf der auch eine Internetseite betrieben wird. Dann kann man halt über das Impressum gehen.

Wenn das aber nicht so einfach zu finden ist, dann geht man halt über [postmaster@domain.tld](mailto:postmaster@domain.tld), die sollten immer erreichbar sein. Reagiert da keiner, hilft ein `dig soa domain.tld` und man schreibt den Zonemaster an, der ja wissen sollte, welchem Kunden der MX gehört. Bei großen Providern gibt es eine Abuse-Abteilung, die man kontaktieren kann.

Jetzt meinte die Telekom vorletzte Woche, eine Mail von mir abzuweisen. Ich maile nicht oft an t-online.de-Adressen, weil die meisten Gott sei Dank keine mehr haben. Aber es kam halt mal wieder vor.

```bash
<[**********@t-online.de](mailto:S-W-B@t-online.de)>: host [mx03.t-online.de](http://mx03.t-online.de/)[194.25.134.73] refused to talk to  
   me: 554 IP=23.88.46.137 - Dialup/transient IP not allowed. Use a  
   mailgateway or contact [toda@rx.t-online.de](mailto:toda@rx.t-online.de) if obsolete. (DIAL)  
Reporting-MTA: dns; [mail01.9it.de](http://mail01.9it.de/)  
X-Postfix-Queue-ID: B72DF44A35  
X-Postfix-Sender: rfc822; [ruediger@kuepper.nrw](mailto:ruediger@kuepper.nrw)  
Arrival-Date: Tue, 12 Aug 2025 09:12:53 +0000 (UTC)  
  
Final-Recipient: rfc822; [**********@t-online.de](mailto:**********@t-online.de)  
Original-Recipient: rfc822;[**********@t-online.de](mailto:**********@t-online.de)  
Action: failed  
Status: 4.0.0  
Remote-MTA: dns; [mx03.t-online.de](http://mx03.t-online.de/)  
Diagnostic-Code: smtp; 554 IP=23.88.46.137 - Dialup/transient IP not allowed.  
   Use a mailgateway or contact [toda@rx.t-online.de](mailto:toda@rx.t-online.de) if obsolete. (DIAL)
```
Der MX mx03.t-online.de meint also, es ist eine Dialup-IP, also von einem Einwahl-DSL-, Kabel- oder Glasfaseranschluss. Nein, ist sie nicht. Es ist eine IP bei einem großen deutschen Hostinganbieter.

Ich habe dann erst einmal selbst ein paar Sachen geprüft. Hätte ich nicht machen brauchen, weil ich wusste schon vorher, dass alles passt.

   * Vom Mailserver passte im DNS der A-Record, 
   * bei der IP die zurückgeliefert wurde passte der PTR genau auf den A-Record vom Mailserver.
   * TXT-Records für SPF: Correct.
   * DKIM: Correct
   * DMARC: Correct 

Die Telekom verwies in der Antwort auf: postmaster.t-online.de

Alles, was sie verlangten, war ok, außer dass keine Kontaktdaten hinterlegt sind. Was man aber mit `dig SOA 9it.de` oder einer einfachen Mail an [postmaster@9it.de](mailto:postmaster@9it.de) hätte gar nicht benötigt. Damit aber bald wieder, wenn doch mal wieder nötig, an eine t-online.de-Mailadresse gemailt werden könnte, habe ich die Mailadressen auf die Seite gepackt. War auch nicht genug, also ist da jetzt ein Link zu einem Impressum.

Lustigerweise meinte einer der Engineers der Telekom zu bemängeln, dass es ja auch nicht bei der mail01.9it.de gegeben sei. Jetzt mal im Ernst. mail01.9it.de und andere Hosts von mir sind reine Mailserver. Genau so betreibt es auch die Telekom und alle anderen Provider. Reine Mailserver, da ist kein Webserver, kein DNS-Server drauf. Sie machen reine Mailserver.

Sollen jetzt mal alle Mailadmins sich einfach mal genau so verhalten und es genau so begründen? Dann kann die Telekom die nächsten Wochen gerne auch mal damit rumschlagen, um ihre Mails loszuwerden.

```bash
mail03:/etc/nginx# dig +short MX rx.t-online.de
10 rx.t-online.de.
mail03:/etc/nginx# dig +short MX t-online.de
10 mx01.t-online.de.
10 mx02.t-online.de.
10 mx03.t-online.de.
10 mx00.t-online.de.
mail03:/etc/nginx# dig +short MX telekom.de
100 mailin12.telekom.de.
100 mailout32.telekom.de.
100 mailin42.telekom.de.
100 mailin32.telekom.de.
100 mailin22.telekom.de.
```
Alle Hostnames für diese MX sind alle nicht per HTTP erreichbar.

Nimmt man sich jetzt einfach mal eine Mail der Telekom:
```bash
2025-07-24T08:50:26.239134+00:00 localhost postfix/smtpd[1207115]: 3A5444167F: client=awmail151.telekom.de[194.25.225.223]
2025-07-24T08:50:26.277338+00:00 localhost postfix/cleanup[1207118]: 3A5444167F: message-id=<-2014779266.1753347025112.JavaMail.rechnung-online@telekom.de>
2025-07-24T08:50:26.591807+00:00 localhost postfix/qmgr[960119]: 3A5444167F: from=<rechnungonline@telekom.de>, size=162701, nrcpt=1 (queue active)
2025-07-24T08:50:26.640058+00:00 localhost postfix/lmtp[1207119]: 3A5444167F: to=<********@9it-server.de>, relay=mail01.9it.de[private/dovecot-lmtp], delay=0.43, delays=0.38/0/0.02/0.03, dsn=2.0.0, status=sent (250 2.0.0 <=<********@9it-server.de> oGl1JNLzgWhpaxIA0J78UA Saved)
2025-07-24T08:50:26.640423+00:00 localhost postfix/qmgr[960119]: 3A5444167F: 
```

Man kann Glück haben und man erreicht durch Zufall den richtigen Mitarbeiter, der es an die richtige Abteilung leitet. Denn man könnte ja dabei einfach auf telekom.de gehen, das Impressum anklicken und sein Glück versuchen. Denn die Kontaktdaten dort sind alles Kontaktdaten, Formulare und Chats, die für Endkunden ausgelegt sind. Kommt da mal mit Mailserver-Logs und Hinweisen zu Spam-Mails oder sonstigen Themen, die technisch sind und nichts mit "Mein DSL ist kaputt" zu tun haben. Eine direkte Kontaktmöglichkeit zu den E-Mail-Engineers findet man auf der Seite nirgendwo.

Andersherum ballern hier auch täglich sehr viele Mails auf meine Mailserver ein, die über Telekom-Infrastruktur versendet werden, die unerwünscht sind und sogar reine Spam-Mails sind. Dabei ist ein Marketinganbieter, bei dem die Telekom auch Technologiepartner ist. Aber auch Versender, die die Infrastruktur der Telekom nutzen. Die Domains, die genutzt werden, sind oft überhaupt nicht erreichbar. Da muss auch jeder andere Mailadmin auch mal `host` und `dig` anwerfen, um herauszufinden, dass ein MX bei der Telekom betrieben wird.

Der Mailserver, bzw. die IP, soll jetzt wieder resettet werden.

`Wir werden veranlassen, dass die Reputation dieser IP-Adresse bei unserem System resettet wird.`

Interessant war in einer vorherigen Mail dieser Satz:

`Von der genannten IP-Adresse war lange Zeit keine Aktivität bei uns feststellbar. Aus Sicherheitsgründen nehmen unsere Systeme von solchen IPs erst nach Prüfung und Reset der Reputation E-Mails entgegen.`

Richtig, es werden kaum Mails an Telekom-Mailadressen versendet. Das sollte ja ein Zeichen dafür sein, dass meine MX sauber sind und die Reputation sollte gut sein. Heißt das jetzt, ich muss einfach regelmäßig Mails an die Telekom-Adressen schicken, um nicht wieder aus deren Liste zu fliegen und ich mich somit wieder unnötig mit der Telekom auseinandersetzen muss?

Ok, als Kunde der Telekom kann man sich ja eine Mailadresse anlegen und einfach jeden Tag eine sinnlose Mail hinschicken lassen.

Obwohl ich ja kurz überlegt hatte, einfach alles, was an Schrott von der Telekom kommt und von meinen Mailservern geblockt oder als Spam erkannt wird, einfach jedes Mal bei der Telekom selbst zu reporten. Die Auswertungen laufen eh automatisiert und es wäre nur eine kleine Erweiterung der Tools, damit es einfach an [abuse@telekom.de](mailto:abuse@telekom.de) geht.

Aber sehr wahrscheinlich würde dann wieder jemand bei der Telekom auf seinem hohen Ross einfach nach ein paar Tagen alles löschen und einfach die IP komplett auf eine Blacklist packen. Selbst Spammern eine Plattform liefern und dann so reagieren, aber wenn andere eigentlich alles richtig machen, kommen sie mit ihren Reputationslisten und wird bestraft, wenn man eben keine Mails und somit Spam einliefert.

Das ist das gleiche Verhalten, wie es Jahre lang von Hotmail.com betrieben wurde. Jeden Tag sind von denen Millionen von Spam-Mails eingetroffen, aber selbst aufführen, als wäre man die Internet-Polizei und man kann anderen vorschreiben, was zu machen ist.

Kriegt erst einmal alle eure eigenen Sachen in den Griff. Dann können wir gerne über Mailsecurity reden.


