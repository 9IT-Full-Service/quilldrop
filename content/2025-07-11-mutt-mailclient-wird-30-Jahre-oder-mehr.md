---
title: 'Der Mutt Mailclient wird 30 oder mehr Jahre alt'
date: 2025-07-11 17:00:00
update: 2025-07-11 17:00:00
author: ruediger
cover: "/images/posts/2025/07/mutt-30-years.webp"
featureImage: "/images/posts/2025/07/mutt-30-years.webp"
tags: [Mail, Mutt, IMAP, Client]
categories: 
    - Software
preview: "Happy Birthday Mutt Mail Client. 30 Jahre ohne Schnick-Schnack einfach Mails abrufen, lesen, schreiben und versenden."
draft: false
top: false
type: post
hide: false
toc: true
---

![Happy Birthday Mutt Mail Client](/images/posts/2025/07/mutt-30-years.webp)

## Mutt Mail client

Da ich vorhin den Mailclient Mutt in einem Artikel erwähnt habe und wissen wollte, wie alt Mutt überhaupt ist, habe ich etwas im Internet recherchiert. Wann wurde es veröffentlicht, wann gestartet?

1995 wurde Mutt veröffentlicht. Also ist Mutt jetzt offiziell 30 Jahre alt, wenn man nach dem Release-Jahr geht. Genau eingrenzen lässt es sich nicht, da nirgendwo ein genaues Datum steht. Auch keine Ankündigung zu finden. Oder das Internet hat tatsächlich mal etwas vergessen.

Interessant war jedoch die Reise durch das Git-Repository, denn dort kann man immer noch sehen, wie die Reise begonnen hat:
Vom ersten initialen Commit auf Github kann man jede Änderungen und Verbesserung seit dem anschauen. 
Ich meine vorher war es in einem SVN Repo. Das habe ich aber nicht mehr gefunden. Auch keine Hinweise in Mailinglisten oder auf Internetseiten warun dazu zu finden. 

```text
#### [Initial revision](https://github.com/muttmua/mutt/commit/1a5381e07e97fe482c2b3a7c75f99938f0b105d4 "Initial revision")

![author](https://github.githubassets.com/images/gravatars/gravatar-user-420.png?size=32)

Thomas Roessler

committed on Jun 8, 1998
```

Und die README.md ist auch noch sehr übersichtlich:

```markdown
README for mutt-0.90i
=======================

Installation instructions are detailed in ''INSTALL''.

The user manual is in doc/manual.txt.

PGP users please read doc/pgp-Notes.txt before proceeding.

For more information, see the Mutt home page,
http://www.cs.hmc.edu/~me/mutt/index.html.

The primary distribution point for Mutt is
ftp://ftp.cs.hmc.edu/pub/me/mutt.  See the home page for mirror sites.

Michael Elkins <me@cs.hmc.edu>, January 22, 1998
Thomas Roessler <roessler@guug.de>, February 3, 1998
```

## Mutt kann 1998 schon gpg/sMime

Wenn man das aus dem Jahr 1998 liest, fragt man sich: Warum ist das für Apple, Microsoft und alle anderen heute immer noch so schwer?

```text
- Other than multipart/mixed and PGP/MIME, Mutt should allow the user to
  specify what to do with other types of multipart messages (i.e., so a user
  can deal with S/MIME messages reasonably)
...
```

Ja, Mutt konnte damals schon mit Verschlüsselung und Signieren umgehen. Das Internet wäre so viel besser, wenn das Standard in allen Clients werden würde.

Jetzt komme bitte keiner mit "Aber das ist doch alles viel zu kompliziert". Kompliziert ist irgendwie alles. Passwörter landen in einer Notiz-App, Word oder Excel. Mehrfaktor-Authentifizierung für Logins oder Passkeys sind zu kompliziert.

Hört auf, ihr seid einfach nur zu faul und zu bequem. Dafür gibt es Apps, die euch das alles abnehmen. Nur komisch finde ich immer, wenn genau das "zu kompliziert" von genau den gleichen Leuten kommt, die Dinge einrichten und installieren, die aus meiner Sicht schon sehr hart an kompliziert sind. Da geht es dann um Accounts, Software, Lizenzen, die über dubiose Wege und mit viel Gefrickel eingerichtet werden müssen, und ich denke nur: "Respekt, aber Passwortmanager und Authenticator-App sind zu schwer."

## Wieder mehr Mutt benutzen

Aber zurück zu Mutt. Aktuell nutze ich Mutt nicht mehr so oft. Aber das wird sich bald wieder ändern. Gerade jetzt zum 30. Geburtstag wird Mutt einfach mal wieder reaktiviert. Gerade mit vielen Mails in vielen Ordnern ist Mutt einfach unschlagbar. Kein Scrollen durch Ordnerlisten für verschiedene Accounts. Einfach eine Taste jede halbe Stunde drücken und man sieht, in welchen Ordnern neue Mails sind. Man sieht nicht nur, in welchen Ordnern neue Mails sind – Mutt springt einfach mit einer Taste zum nächsten Ordner mit neuen Mails.

Mutt ist so mächtig. Und das vermisse ich auch irgendwie bei allen Mail-Clients, die es aktuell so gibt. Einfach mit einem Shortcut alle Mails aus der Inbox verschieben, die mit GitHub-PRs zu tun haben. Oder alle Incident-Mails. Oder man markiert mal eben per Regex alle Mails außer den 2-3, die man noch benötigt, und verschiebt dann einen Batch mit 10 oder mehr Mails.

## Meine Mutt Konfiguration

Hier einmal meine Konfiguration für 2 IMAP-Mailaccounts, sehr abgespeckt. Das Switchen zwischen den Accounts geht mit den Tasten 1 und 2. Falls die Sidebar mal nicht aktualisieren sollte: einfach 0 drücken.


### Installation

```bash
brew install mutt

# mutt_dotlock(1) has been installed, but does not have the permissions to lock
# spool files in /var/mail. To grant the necessary permissions, run:

sudo chgrp mail /opt/homebrew/Cellar/mutt/2.2.14/bin/mutt_dotlock
sudo chmod g+s /opt/homebrew/Cellar/mutt/2.2.14/bin/mutt_dotlock

# Alternatively, you may configure `spoolfile` in your .muttrc to a file inside
# your home directory.
```

### .muttrc Config File

```bash
cat <<EOF > ~/.muttrc
folder-hook 'account.com.example.user1' 'source ~/.mutt/account.com.example.user1'

folder-hook 'account.com.example.user2' 'source ~/.mutt/account.com.example.user2'

# Default account
source ~/.mutt/account.com.example.user1


macro index 1 '<sync-mailbox><enter-command>source ~/.mutt/account.com.example.user1<enter><enter-command>exec imap-logout-all<enter><enter-command>unset sidebar_visible<enter><enter-command>set sidebar_visible=yes<enter><change-folder>=INBOX<enter>' "Switch to User 1 account"

macro index 2 '<sync-mailbox><enter-command>source ~/.mutt/account.com.example.user2<enter><enter-command>exec imap-logout-all<enter><enter-command>unset sidebar_visible<enter><enter-command>set sidebar_visible=yes<enter><change-folder>=INBOX<enter>' "Switch to User 2 account"

macro index 0 '<enter-command>unset sidebar_visible<enter><enter-command>set sidebar_visible=yes<enter><enter-command>set sidebar_width=25<enter><refresh>' "Full sidebar reset"

# Fetch mail shortcut
bind index G imap-fetch-mail

set mail_check = 60 # Prüfe alle 60 Sekunden
set imap_check_subscribed = yes # Nur abonnierte Ordner prüfen
set imap_list_subscribed = yes # Nur abonnierte Ordner anzeigen
set imap_passive = no
set imap_idle = yes
set mail_check_stats = yes
set mail_check_stats_interval = 60

# Seitenleiste mit Ordnern (optional):

set sidebar_format = "%B%?F? [%F]?%* %?N?%N/?%S"
set folder_format = "%2C %t %N %F %2l %-8.8u %-8.8g %8s %d %f"
set sort_browser = alpha
set sidebar_visible = yes
set sidebar_width = 25
set sidebar_short_path = yes
# folder indent wer auch dort einrückungen möchte 
# set sidebar_folder_indent = yes
# set sidebar_delim_chars = "/."
set sidebar_new_mail_only = no
set sidebar_next_new_wrap = yes

bind index,pager K sidebar-prev # K = Nach oben
bind index,pager J sidebar-next # J = Nach unten
bind index,pager L sidebar-open # L = Öffnen
bind index,pager ü sidebar-page-up # ü = Seite hoch
bind index,pager ö sidebar-page-down # ü = Seite runter
bind index,pager B sidebar-toggle-visible
  
# Alle Header erstmal ausblenden

ignore *
# Nur wichtige Header anzeigen (in dieser Reihenfolge)
unignore from: to: cc: bcc: date: subject: reply-to:
unignore organization: user-agent: x-mailer:
unignore list-id: x-mailing-list:

# Header-Reihenfolge festlegen

hdr_order from: to: cc: bcc: subject: date: reply-to: organization:

# === PAGER-KONFIGURATION (E-Mail-Ansicht) ===

# Anzahl der Index-Zeilen über der E-Mail anzeigen

set pager_index_lines = 10 # Zeigt 10 E-Mails über der aktuellen

# Pager-Verhalten

set pager_stop = yes # Stoppe am Ende der Nachricht
set pager_context = 3 # 3 Zeilen Kontext beim Scrollen
set menu_scroll = yes # Scrollen statt Seitenweise
set smart_wrap = yes # Intelligenter Zeilenumbruch
set markers = no # Keine '+' Zeichen bei umgebrochenen Zeilen
set pager_index_lines = 10

# Aliases / Addressbook 
set alias_file = ~/.mutt/aliases
set sort_alias = alias
set reverse_alias = yes
set reverse_name = yes
set reverse_realname = yes
source ~/.mutt/aliases
set alias_format = "%4n %2f %t %-15a %-25r"


set charset = "utf-8"
set send_charset = "utf-8"
set assumed_charset = "utf-8"

# Für bessere Anzeige beim E-Mail schreiben:

set compose_format = "-- Mutt: Compose [Approx. msg size: %l Atts: %a]%>-"

# === SCHÖNERE E-MAIL-LISTE ===
# Index-Format (E-Mail-Liste) anpassen
set index_format = "%4C %Z %{%b %d} %-15.15L (%?l?%4l&%4c?) %s"

EOF
```


### Accounts anlegen 

```bash
mkdir -p ~/.mutt/cache/com.example.user1/{bodies,header}
mkdir -p ~/.mutt/cache/com.example.user2/{bodies,header}
```

#### User 1:

Datei: ~/.mutt/account.com.example.user1


```bash
cat <<EOF > .mutt/account.com.example.user1
unmailboxes *

set ssl_starttls=yes  
set ssl_force_tls=yes  
set from = 'user1@example.com'
set use_from = yes
set imap_user = 'user1@example.com'
set imap_pass = 'yourpass' 
set smtp_pass = 'yourpass'
set realname='Bob Mayer'  
set folder = "imaps://mail.example.com:993"
set spoolfile = "+INBOX"
set postponed = "+Drafts"

set record = "+Sent Messages"
set trash = "+Deleted Messages"
 
set header_cache = "~/.mutt/cache/com.example.user1/headers"  
set message_cachedir = "~/.mutt/cache/com.example.user1/bodies"  
set certificate_file = "~/.mutt/certificates"  
set smtp_url = 'smtp://user1@mail.example.com:587/'

set smtp_authenticators = "plain"
set move = no  
set imap_keepalive = 900

set ssl_use_sslv3 = no
set ssl_use_tlsv1 = no
set ssl_use_tlsv1_1 = no
set ssl_use_tlsv1_2 = yes
set ssl_use_tlsv1_3 = yes
set ssl_verify_dates = yes
set ssl_verify_host = no
EOF
```

#### User 2:

Datei: Datei: ~/.mutt/

```bash
cat <<EOF > .mutt/account.com.example.user1
unmailboxes *

set ssl_starttls=yes  
set ssl_force_tls=yes  
set from = 'user2@example.com'
set use_from = yes
set imap_user = 'user2@example.com'
set imap_pass = 'yourpass' 
set smtp_pass = 'yourpass'
set realname='Bob Mayer'  
set folder = "imaps://mail.example.com:993"
set spoolfile = "+INBOX"
set postponed = "+Drafts"

set record = "+Sent Messages"
set trash = "+Deleted Messages"
 
set header_cache = "~/.mutt/cache/com.example.user2/headers"  
set message_cachedir = "~/.mutt/cache/com.example.user2/bodies"  
set certificate_file = "~/.mutt/certificates"  
set smtp_url = 'smtp://user2@mail.example.com:587/'

set smtp_authenticators = "plain"
set move = no  
set imap_keepalive = 900

set ssl_use_sslv3 = no
set ssl_use_tlsv1 = no
set ssl_use_tlsv1_1 = no
set ssl_use_tlsv1_2 = yes
set ssl_use_tlsv1_3 = yes
set ssl_verify_dates = yes
set ssl_verify_host = no
EOF
```

### Mutt Alias / Adressbuch Datei

Alias / Adressbuch:

```bash
~/.mutt/aliases
alias bob "Bob Mayer" <user1@example.com>
alias Alice "Alice Müller" <user1@example.com>
```

