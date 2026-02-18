---
title: 'E-Mails verwalten'
date: 2016-01-08 02:00:51
update: 2016-01-08 02:00:51
author: ruediger
cover: "/images/cat/organisation.webp"
tags:
    - E-Mail
    - Feed
    - Filter
    - IMAP
    - imapdir
    - Internet
    - Maildir
    - Organisation
    - Organisieren
    - RSS
    - Rule
    - Sieve
preview: "In diesem [Google+](https://plus.google.com/+RüdigerPretzlaff/posts/9v1WyUG1NM5) Post hatte ich das Thema ja mal anschnitten und jetzt endlich mal angefangen genauer zu beschreiben. Hier aber erst noch mal das von Google+"
categories: 
    - Organisation
toc: false
hide: false
type: post
---

In diesem [Google+](https://plus.google.com/+RüdigerPretzlaff/posts/9v1WyUG1NM5) Post hatte ich das Thema ja mal anschnitten und jetzt endlich mal angefangen genauer zu beschreiben. Hier aber erst noch mal das von Google+:

> A: wie viele Mail bekommst du eigentlich so am Tag.

> Ich: puh, kann ich gar nicht sagen. Müsste ich mal über eines der Logfiles das Script jagen.
<!--more-->
> A: ähm Moment, wieso sind das bei dir so vielen Ordner im Postfach. Äh, wie viele sind das denn?

> Ich: puh, kann ich gar nicht sagen. Letzte Woche waren es noch etwas über 140. Aber die werden
automatisch angelegt, je nachdem was für eine Mail reinkommt

> A: wie hast Du für alles Regeln angelegt?

> Ich: nein, aber eine richtige. Das zu erklären dauert aber etwas länger. Ich will das aber eh mal
aufschreiben.Auch weil schon andere gefragt haben, weil viele schon mit viel weniger Mails
überfordert sind.Mit diesem Mailsetup ist mir auch eine Funktion bei iOS aufgefallen die
richtig schick ist. Es gibt Ordner die werden zwar per Notification angezeigt. Aber nicht
als Zahl auf dem Homescreen. Wichtige Mail aus der Inbox fallen auf, da die Anzahl zusehen
ist. Nicht wichtige, aber gut wenn du sie mal wahrgenommen hast werden notifiziert, aber
sonst nichts. Alles andere ist unwichtig und man sieht sie unterwegs nur wenn man in einen
Mailordner gehen würde. Aber genauer dann wenn ich alles fertig gesammelt und geschrieben habe.


Als ersten einmal die Anzahl der Mails: ca. 700 bis 1200 pro Tag.

Davon gelangen in den Posteingangsordnern gerade mal ca. 30-50 Mails pro Tag. Alle anderen Mails sind in verschiedene Ordner schon auf dem Mailserver sortiert worden. Dabei gibt es mehrere Kategorien.

1.  AutoSort
    1.  Accounts
    2.  Dienste (Serverzeug)
    3.  Shops
    4.  Verträge
    5.  Soziale Netzwerke
    6.  Fritzboxen
2.  Feeds \*(Dazu später mehr \[rss2mail\])
3.  Office \*(Shared Folders)

_**Feeds**_ Fangen wir als erstes mit den einfachen Dingen an. Als erstes etwas zu der Kategorie **Feeds**. Diese E-Mails kommen nicht von externen Server, sondern werden auf dem Server selbst erstellt. Dazu verwende ich [rss2email](http://www.allthingsrss.com/rss2email/) und habe in der Konfiguration RSS-**Feeds** von Newsseiten und Blogs eingetragen, die an verschiedene Accounts gesendet werden. Unterkategorien sind:

1.  Apple
2.  Blogs
3.  News
4.  Städte

In den beiten Kategorien "**News**" und Städte sind z.B. die News Feeds von DW (Deutsche Welle) und bei Städte die Feeds von der-westen.de (Stadt Essen), Ka-news.de und Karlsruhe-insider.de. Die entsprechenden Feeds werden dann per Mail an folgende E-Mailadressen gesendet:

*   feeds+News.DW@example.net sortiert den Feed von dw.com in den Mailordner "Feeds.News.DW" (maildir ~/.Feeds.News.DW/)
*   feeds+Städte.Essen@example.net
*   feeds+Städte.Karlsruhe@example.net

Das ganze sortiert jetzt erst einmal überhaupt nichts. Der Server muss den Delimiter "+" unterstützen. Sollte er das nicht machen hat man trotzdem noch die Möglichkeit dies mit einem "Sieve-Filter" auf dem Server zu machen und einem Catch-All Account. Dazu und den Regeln auf dem Mailserver aber erst später. Vorher noch kurz zu den Kategorieren "**Office**" und "**AutoSort**". **AutoSort** ist die Kategorie um die es hier eigentlich geht. _**Office**_ Die Kategorie "**Office**" ist eigentlich am schnellsten erklärt. In dieser befinden sich nur Shared Folders. Also Ordner die man zwischen verschiedenen Accounts einfach freigeben kann. Das sind entweder Ordner für alle meine Accounts, damit ich z.B. alle Rechnungen, Bestellungen usw. von allen Account nach dem lesen in den jeweiligen Ordner verschiebe um alle Rechnungen von jedem Account aus schnell erreichen zu können. Die anderen Ordner sind dann auch für andere User auf dem Server freigeben. Projekte mit anderen werden damit organisiert, Daten ausgetauscht, archiviert usw. _**AutoSort**_ Im Google Post vom 30.10.2015 wurde eine Ordnerzahl von 140 genannt. Aktuell sind es jetzt 275 Ordner. Zur Erinnerung noch einmal die Liste der Unterkategorien:

1.  Accounts
2.  Dienste (Serverzeug)
3.  Shops
4.  Verträge
5.  Soziale Netzwerke
6.  Fritzboxen

In den Beispielen wird immer "user+\*@example.net" genannte. Die eigentliche Adresse ist user@example.net. "User" ist der Localpart der Mailadresse, bei den meisten also der Benutzername: cmueller@example.net.

_Hinweis:_

_Bei den Beispielen hier benutze ich bei den einzelnen Ordnernamen immer einen Grossbuchstaben am Anfang. Dies dient hier nur der besseren Lesbarkeit. Man kann das auch so machen, sollte sich aber bewusst sein das manche Anbieter das zwar so annehmen beim eintragen einer E-Mailadresse, teilweise auch in Einstellungen auch wieder so anzeigen, aber beim versenden der E-Mail alles in Kleinbuchstaben umwandeln. Das könnte dazu führen das auf einmal zwei Ordner für Onlineshops vorhanfen sind. "Shops" und "shops". Daher benutze ich nur noch Kleinbuchstaben._

In **Accounts** werden alle Accounts für alles mögliche gefiltert. Zum Beispiel dort in den Kategorien "Foren", "Musik", "Software", usw.

Melde ich mich in einem Forum an für Musik, landet es auch in Musik. Für die Anmeldung benutzt man z.B. die Mailadresse:

"user+Accounts.Foren.Gitarrenforum@example.com" -> landen im Mailordner "AutoSort.Foren.Musik.Gitarrenforum"

"user+Foren.Musik.BandsInNrw@example.net" -> Mailordner:   AutoSort.Foren.Musik.BandsInNrw

usw.

Für unseren Beispiel Kai wäre das dann: cmueller+Foren.Musik.BandsInNrw@example.net. In den weitern Beispielen werde ich für die Übersichtlichkeit immer "user+...." benutzen.

Bei **Dienste** filter ich alles in verschiedene Ordner für z.B.: "Cron", "Certs" (noch weiter unterteilt für die einzelnen Anbieter), SA-Learn (Spamlernen), usw.

Cronjobs auf dem dem Server werden an "user+AutoSort.Server.Cron@example.net" geschickt. Wenn man mehrere Server betreibt oder Mails aus verschiednen Jobs noch weiter filtern möchte hängt man einfach dementsprechend ".ServerA", ".ServerB" oder ".Backups", "Updates" an die Adresse an.

*   user+Server.Cron@example.net
*   user+Server.Cron.ServerA@example.net
*   user+Server.Cron.ServerA@example.net
*   user+Server.Cron.Backups@example.net
*   user+Server.Cron.Updates@example.net

Immer wird in den Cron oder den Unterordner in Cron gespeichert. Wem nur Cron reicht, der benutzt halt nur die erste Adresse. Falls man später merkt das dieser Ordner nicht reicht ändert man nur die Adresse für einen Server oder Job in der Konfiguration und schon werden die nächsten Mails neu sortiert eintreffen.

Dazu muss man keinen Filter anpassen. Zu den Filtern kommen wir aber später. Erst noch ein paar Beispiele um verständlich zu machen wie der Aufbau der E-Mailadressen sind.

**Shops** und **Soziale Netzwerke**

Amazon, Ebay, Wein Müller, Musikladen Mustermann. Alles Shops bei denen man sich so anmeldet. Die ersten beiden kennt jeder, die letzten beiden existieren glaube ich gar nicht. Aber sind ja auch nur Beispiele.

Bei allen Diensten meldet man sich ab jetzt nur noch mit Extraadressen für jeden Shop an. Oder man ändert die Mailadressen dort ab.

*   Amazon: user+Shops.Amazon@example.net
*   Ebay: user+Shops.Ebay@example.net
*   Wein Müller: user+Shops.WeinMueller@example.net
*   Musikladen Mustermann: user+Shops.MusikMustermann@example.net

Wer jetzt viele Musikläden und Getränkelieferanten in der Liste hat kann auch hier wieder unterteilen und sortieren:

*   user+Shops.Musik."AnbieterA"@....
*   user+Shops.Musik."AnbieterB"@....
*   usw.

Auch hier gilt wieder: Weitere Untergliederung ist immer möglich. Einfach die Mailadresse beim Shop/Anbieter ändern auf die neue (hier ".Musik." dazwischen packen) Adresse und fertig. Keine Änderungen am Mailfilter nötig. Die **Sozialen Netzwerke** landen bei mit im Ordner "**SN**", ich bin ab und zu Tippfaul und lange Ordnernamen will ich auch nicht. Aber auch noch ein anderer Punkt hat gerade bei den sozialen Netzwerken zu einer Verkürzung bei der Ordnerbezeichnung geführt. Es gab da einen Anbieter der nur einen nicht gerade lange localpart vor dem @-Zeichen oder die komplette Adresse nicht zu lang sein durfte. Es war ausgerechnet ein nicht mal alter Dienst. Eher ein neue Fancy Service, so ein Hipsterzeugs. Da fand ich es sehr lächerlich nur sehr kurze Adressen zulässig waren. Da wurde sehr wahrscheinlich das Formular schnell zusammen kopiert und Hipstermässig ohne richtiger QA gearbeitet. **Fritzboxen** und **Vertäge** Für die einzelnen Fritzboxen in der Familie bekomme ich die Mails für Updates, Fehler beim DynDNS, VPN Status usw. Verträge werden für Gas, Strom, Internetanschluss, Handytarife usw. angelegt. Hier jetzt nur noch die Beispiele dazu:

*   user+Fritzbox.Karlsruhe.rs32@example.net
*   user+Fritzbox.Essen.FS371@example.net
*   user+Fritzbox.Essen.FB39@example.net
*   user+Fritzbox.Essen.BS206@example.net
*   user+Contracts.Strom@example.net
*   user+Contracts.Gas@example.net
*   user+Contracts.Telekom@example.net
*   user+Contracts.T-Mobile@example.net

**Lists** Bei Mailinglisten wird genau das gleiche gemacht:

*   user+Lists.debian@example.net
*   user+Lists.debian.updates@example.net
*   user+Lists.debian.security@example.net
*   user+Lists.redhat@example.net
*   user+Lists.linuxkernel@example.net

**_Die Filterregel(n)_** Damit das ganze funktioniert braucht man nur noch den/die Filter anlegen. Eigentlich benötigt man nur einen Filter. Aber in bestimmten Situationen sind 3 oder 4 Filter sinnvoll. Dazu gleich aber noch mehr. Als erstes benötigt man im Sieve Filter einen weiteren Eintrag ("variables") bei "require".

```
require ["reject","fileinto","imap4flags","body","vacation","copy","variables"];
```

Jetzt fügt man nur eine Regel im Filter dazu:

```
rule:[AutoSort1]
if header :matches "X-Original-To" "user+\*@example.net" { fileinto "AS.${1}"; stop; }
```

Diese Regel bewirkt:

1.  alles was zwischen "user+" und dem "@" steht wird in eine Variable gespeichert.
2.  Als Zielordner wird beim fileinto die Variable an den AutoSort Ordner (AS) gehangen.

Mail an "user+Shops.Amazon@example.net" wird zu: fileinto "AS.Shops.Amazon"; **Jetzt die Ausnahmen,** falls man noch weitere "Probleme" umgehen oder Anforderungen haben sollte. _Probleme:_ Es gibt Anbieter die kein "+"-Zeichen in der E-Mailadresse zulassen. Entweder können sie im Backend damit nicht umgehen oder wie ich bei einigen erfahren konnte, können sie das speichern und auch benutzen. Nur ist der Check im Onlineformular nicht dafür ausgelegt und zeigt einfach einen Fehler an. Wer eine Domain alleine benutz kann einfach einen Catch-All Account auf seine Mailadresse anlegen. Oder er nicht die komplette Domain als Catch-All benutzen möchte kann je nach Mailserver und Konfiguration auch einen Catch-All auf den User anlegen: "user%@example.net" weitergeleitet auf "user@example.net" Anschliessen kann man anstatt "user+Shops.Shopname@example.net" die Adresse "userShops.Shopname@example.net" benutzen. Als Filter dafür fügt man dann folgendes dazu:

```
rule:[AutoSort2]
if header :matches "X-Original-To" "user\*@example.net" { fileinto "AS.${1}"; stop; }
```

Wer mehrere verschiedene Domains auf einen Account zusammenfügt kann folgende Regel benutzen: \[cc lang=c\]

```
rule:[delimiter1]
if header :matches "X-Original-To" "user+_@_" { fileinto "AS.${1}"; stop; }
```
**Die Feeds** Vor den AutoSort Regeln habe ich noch die letzte Rule eingetragen. Da ich die RSS-Feeds nicht mit im Ordner AutoSort haben möchte filter ich diese vorher schon weg. Beispiel von oben: "feeds+News.DW@example.net"

```
rule:[feeds]
if header :matches "X-Original-To" "feeds+\*@example.net" { fileinto "Feeds.${1}"; stop; }
```
Damit ist auch schon alles erledigt. Ich muss nie wieder einen Filter anlegen für neuen Accounts, Shops, Servermails o.a. Ich muss nur die richtige Mailadresse bei der Anmeldung/Einrichtung eintragen. Alle wichtigen Mails kann ich weiter auf der Hauptadresse eintreffen lassen und sehe sie direkt im Posteingang. Der Rest bei Bedarf. Ich kann mir auch so sehr gut meine Zeit planen wann ich welche Mails lese. Foren, Shops, Feeds stehen bei mir ganz hinten auf der Todo. Die sieht ungefähr so aus, variiert aber immer mal ein wenig. Wobei auch nicht alle Kategorieren hier nicht aufgelistet sind.

1.  Posteingang
2.  Fritzbox
3.  Server
4.  Mailinglisten
5.  Soziale Netzwerke
6.  Foren
7.  Feeds

_**Benachrichtigungen am iPhone**_ Im Google+ Beitrag hatte ich die Notifikationen am iPhone erwähnt. Wenn alles weggefiltert wird sieht man die Mails erst wenn man in einen der vielen Ordner guckt. Bei iOS kann man in der Liste mit der "Inbox", "Papierkorb", "Entwürfe" usw. auf "Bearbeiten" und ganz unten auf "Ordner hinzufügen" tippen. In der Ordnerliste kann man jetzt die gewünschten Ordner auswählen. Bei mir sind das ein paar News und Blog Feeds, der Nagios Monitoring Ordner und alles was wichtig ist. Insgesamt maximal 10 Ordner. Der schöne Nebeneffekt ist die Notifizierung und die Anzeige der neuen Mails bei geschlossener Mail.App. Alle Mails die in der Inbox und den gerade hinzufügten Ordnern werden per iOS Notifikation auf dem iPhone und der Apple Watch angezeigt. In der Mitteilungszentrale sieht man sie auch. Aber bei der geschlossenen Mail.app werden nur die E-Mails aus der Inbox angezeigt. So bekomme ich mit wenn Nagios Dienste alarmiert, sehe sie und kann sie zu Kenntnis nehmen und reagieren. Kommt dann eine wichtige Mail in die Inbox und gucke nur auf den den entsperrten Homescreen, sehe ich an der Mail.app nur eine "1" und nicht auch noch alle ungelesenen Nagios Meldungen. Genau das gleiche gilt für alle hinzugefügten Ordner News, Blogs u.a., sie werden zwar benachrichtigt, aber stören nicht mehr. Damit sind das schon mal 2 Stufen der Benachrichtigung. Aber wir haben ja jetzt auch noch die Ordner die nicht in der Liste hinzugefügt wurden. Damit sind es jetzt also 3 Stufen was die Priorisierung der E-Mails angeht.

1.  Inbox = wichtig
2.  Ordner mit Benachrichtigung
3.  Ordner ohne, die werden gelesen wenn die Zeit dafür übrig ist.

Die grosse Anzahl der Ordner wirkt erst einmal Heftig. Ich benutze das jetzt aber schon etwas länger und kann sagen der Arbeitsablauf "E-Mails lesen" hat sich dadurch erleichtert. Ich kann mich einfach, sinnvoll und mich dann um die Sachen kümmern wie es gerade passend ist. Im Zug schnell die Inbox checken, danach ist Zeit für News und Blogs. Nach dem dort der erste Überlick über neue Artikel gemacht wurde kommen kurz Server Mails dran. Nichts auffälliges oder aktuell nichts machen kann, aber weiß was ich im Büro dann als erstes mache, geht es dann erst einmal zurück zu den Artikeln. Man ist einfach schnell in die einzelnen Bereiche der Kategorien gesprungen, checken, reagieren, planen, konsumieren usw.

**Kurz zu rss2email**

Ich benutze [rss2email](http://www.allthingsrss.com/rss2email/) (früher auch mal [feed2imap](http://home.gna.org/feed2imap/)) gerne für RSS-Feeds, da ich mit RSS-Readern nie so richtig warm geworden bin. RSS-Reader gibt es zwar auch richtig gute. Aber immer wenn ich mich gerade an einen gewohnt habe wurde die Entwicklung eingestampft. Wieder einen neuen suchen. Kurz drauf wurde die Entwicklung eingestellt. Das war mir dann irgend wann zu lästig. Dabei waren dann auch mal RSS-Reader die zwar von anderen Readern importieren konnten, aber nicht exportieren. Das war dann richtig übel, weil man die Feeds dann einzeln da raus kopieren musste und im neuen Eintragen. Wer also einen RSS-Reader sucht, eh gerne mit einem Mailclient arbeitet und dort auch viel mit Mailordnern macht sollte sich [rss2email](http://www.allthingsrss.com/rss2email/) oder [feed2imap](http://home.gna.org/feed2imap/) einmal angucken.
