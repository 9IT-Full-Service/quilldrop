---
title: 'Wordpress, Slack.com, Telegram, Icinga'
date: 2016-02-24 11:40:56
update: 2016-02-24 11:40:56
author: ruediger
cover: "/images/cat/internet.webp"
tags: [Internet, Wordpress, Icinga, Slack, Telegram]
preview: "Eher durch Zufall bin ich auf Slack.com gestoßen. Ich hatte nach einer Chatlösung für Seiten gesucht. Dabei bin ich auf Chatlio gestossen. Damit wurden die Anfragen per Chat in Slack gesendet. Antworten musste man dann aber leider über den Adminbereich vom Wordpress."
categories: 
    - internet
toc: false
hide: false
type: post
---


Eher durch Zufall bin ich auf Slack.com gestoßen. Ich hatte nach einer Chatlösung für Seiten gesucht. Dabei bin ich auf Chatlio gestossen. Damit wurden die Anfragen per Chat in Slack gesendet. Antworten musste man dann aber leider über den Adminbereich vom Wordpress. Dabei ist mir aber die Nagios/Icinga App im Slack aufgefallen. Installiert und im Icinga konfiguriert schickt es jetzt die Alarme in einen Slack Raum. slack\_nagios.cfg

<!--more-->

```
define contact {
      contact\_name                             slack
      alias                                    Slack
      service\_notification\_period              24x7
      host\_notification\_period                 24x7
      service\_notification\_options             w,u,c,r
      host\_notification\_options                d,r
      service\_notification\_commands            notify-service-by-slack
      host\_notification\_commands               notify-host-by-slack
}

define command {
      command\_name notify-service-by-slack
      command\_line /usr/local/bin/slack\_nagios.pl -field slack\_channel=#alerts -field HOSTALIAS="$HOSTNAME$" -field SERVICEDESC="$SERVICEDESC$" -field SERVICESTATE="$SERVICESTATE$" -field SERVICEOUTPUT="$SERVICEOUTPUT$" -field NOTIFICATIONTYPE="$NOTIFICATIONTYPE$"
}

define command {
      command\_name notify-host-by-slack
      command\_line /usr/local/bin/slack\_nagios.pl -field slack\_channel=#ops -field HOSTALIAS="$HOSTNAME$" -field HOSTSTATE="$HOSTSTATE$" -field HOSTOUTPUT="$HOSTOUTPUT$" -field NOTIFICATIONTYPE="$NOTIFICATIONTYPE$"
}
```

Das benötigte Pakete installieren:

```
sudo apt-get install libwww-perl libcrypt-ssleay-perl
```

In den Einstellungen auf Slack.com den Token erstellen. Das Skript für die Notifizierung über Slack.com herunterladen, ausführbar machen und den Token, sowie die Slack Team URL eintragen.

```
wget https://raw.github.com/tinyspeck/services-examples/master/nagios.pl
cp nagios.pl /usr/local/bin/slack\_nagios.pl
chmod 755 /usr/local/bin/slack\_nagios.pl

```

Team und Token in nagios.pl anpassen:

```
my $opt\_domain = "DeinSlackTeam.slack.com"; # Your team's domain
my $opt\_token = "HKhwerKJ72Kghhj23gsJG8"; # The token from your Nagios services page
```

Icinga neustarten und die nächsten Alarme landen im Slack Team Channel. Ich habe dafür einen neuen Raum #alerts im Slack angelegt.

### Chat auf der Homepage.

Chatlio war jetzt nicht so super. Man kann nur antworten wenn man selbst online im Desktop ist. Für Leute die unterwegs auch mit Kunden Kontakt aufnehmen wollen nicht brauchbar. Slack hat im App Verzeichnis einen Chatbot. Linked-Chat im Slack installieren und einfach Schritt für Schritt die angebenen Schritte durch gehen. Als erstes verbinden man einfach einen Raum im Slack mit Linked-Chat. Anschliessend noch Telegram. Bei beiden gibt es nur wenige Schritte um dies zu bewerkstelligen.

1.  Slack bzw. Telegram Link anklicken.
2.  Raum wählen
3.  den Befehl /link <linked-Chat-ID> senden.

Das war es dann auch schon. Jetzt kann man noch den Titel der Chatbox anpassen (Online/Offline). Farbe anpassen, Position des Chats auf der Seite und die Arbeitszeiten angeben. Durch die Arbeitszeiten wird festgelegt wann der Chat erreichbar ist und wann die Nachrichten nur mit angegebener E-Mailadresse gesendet werden. So kann man auch Nachts Anfragen annehmen und dem Kunden später auch noch antworten. Das ganze dann speichern und auf der Homepage einfach nur diesen Code mit einbauen:

Jetzt erscheint sofort auf der Internetseite unten eine kleine Box mit dem Chat und Besucher der Seite können schnell Anfragen stellen. Das ganze werde ich jetzt mal mit zwei Bekannten testen die schon mal genau wegen so etwas gefragt haben.
