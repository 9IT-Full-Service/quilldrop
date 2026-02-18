---
title: "Roundcube Webmail mit Zwei-Faktor Authentifizierung"
date: 2019-03-30 21:35:05
update: 2019-03-30 21:35:05
author: ruediger
cover: "/images/cat/internet.webp"
tags: ["Alle Beiträge", "Authentication", "Authentifizierung", "E-Mail", "Internet", "Internet", "Security", "Two-Factor", "Two-Factor Authentication", "Webmail", "Zwei-Faktor"]
preview: "Um die Sicherheit zu erhöhen habe ich vor ein paar Wochen das Two-Factor Gauthenticator Modul im Roundcube Webmail hinzugefügt. Damit können die User sich nicht mehr nur mit Benutzer und Passwort einloggen, sondern brauchen einen zweiten Faktor. Das ist dann ein 6-stelliger Code aus dem Google Authenticator auf dem Smartphone. Die Installation ist recht einfach."
categories: 
    - Internet
toc: false
hide: false
type: post
---

Um die Sicherheit zu erhöhen habe ich vor ein paar Wochen das Two-Factor Gauthenticator Modul im Roundcube Webmail hinzugefügt. Damit können die User sich nicht mehr nur mit Benutzer und Passwort einloggen, sondern brauchen einen zweiten Faktor. Das ist dann ein 6-stelliger Code aus dem Google Authenticator auf dem Smartphone. Die Installation ist recht einfach. Das Module wird in das Webverzeichnis in dem Roundcube installiert ist im Ordner "**plugins**" gespeichert. z.B. "**/var/www/webmail/plugins**"

<!--more-->
> ```
> cd /var/www/webmail/plugins
> git clone https://github.com/alexandregz/twofactor\_gauthenticator.git
> ```

Um das PlugIn zu aktiveren öffnet man die Roundcube Konfigurationsdatei und fügt das PlugIn in der Liste hinzu:

> ```
> $config\[‘plugins’\] = array(‘plugin1’, ‘plugin2’, ‘twofactor\_gauthenticator’);
> ```

Anschliessend kann man sich im Roundcube einloggen und in die Einstellungen unter dem Punkt "**Zwei-Faktor Authentifizierung**" die Einstellungen vornehmen und den Goolge Authenticator einrichten.

![Roundcube Webmail mit Two-Factor Authentication](/images/posts/2factor-1024x312.webp)


1.  Auf speichern klicken.
2.  Es erscheint eine Meldung das nicht alle Felder ausgefüllt sind.
3.  OK Bestätigen und es erscheint ein QR-Code.
4.  Den Google Authenticator öffnen.
5.  Hinzufügen auswählen und Barcode scannen.
6.  Den 6 stelligen Code aus der App eingeben und prüfen anklicken.
7.  Auf "Speichern" klicken.

Ab jetzt wird nach dem Login im Webmail nach dem zweiten Faktor gefragt. Die Wiederherstellungscodes sollte man speichern und sicher aufbewahren. Diese werden benötigt wenn man das Smartphone verliert oder wenn der Google Authenticator zum Beispiel bei einem defekten Gerät nicht mehr zur Verfügung stehen sollte.
