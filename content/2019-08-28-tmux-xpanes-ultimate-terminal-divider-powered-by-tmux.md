---
title: "tmux-xpanes - Ultimate terminal divider powered by tmux"
date: 2019-08-28 08:15:00
update: 2019-08-28 08:15:00
author: ruediger
cover: "/images/posts/2019/08/28/ping_pane_title.webp"
tags:
    - Terminal
    - tmux
    - tmux-xpanes
preview: "Bei diesem Terminal-Multiplexer kann man sehr gut und schnell per Shortcut aktivieren das Befehle in allen Fenstern ausgeführt werden. Nach dem Wechsel auf den Mac habe ich immer eine brauchbare Alternative gesucht."
categories: 
    - Internet
toc: false
hide: false
type: post
---

Als ich noch eine Linux Workstation hatte, habe ich eine lange Zeit [Terminator](https://gnometerminator.blogspot.com) benutzt.
Bei diesem Terminal-Multiplexer kann man sehr gut und schnell per Shortcut aktivieren das Befehle in allen Fenstern ausgeführt werden.
Nach dem Wechsel auf den Mac habe ich immer eine brauchbare Alternative gesucht.
Jetzt bin ich über [tmux-xpanes](https://github.com/greymd/tmux-xpanes) gestolpert.

<!--more-->

![xpanes Ultimate terminal divider powered by tmux](/images/posts/2019/08/28/movie_v4.gif)


<!-- /img/posts/2019/08/28/movie_v4.gif -->

Mit [csshx](http://macappstore.org/csshx/) bin ich nie richtig warm geworden.

Bei [iterm2](https://iterm2.com) den ich benutze geht das auch mit
`send input to all tabs` . Aber irgendwie war das auch immer nicht so super.

Mit [tmux](https://github.com/tmux/tmux) kann man das aktivieren mit `:setw synchronize-panes` und wieder ausschalten mit `:setw synchronize-panes off`

[tmux-xpanes](https://github.com/greymd/tmux-xpanes) macht es aber irgend wie schicker. Alleine schon der Aufruf ist schon mal
sehr cool.

```
xpanes --log=~/log --ssh user1@host1 user2@host2 user2@host3
docker ps -q | xpanes -s -c "docker exec -it {} sh"
```

Weitere Beispiele sind auf der Github Seite von [tmux-xpanes](https://github.com/greymd/tmux-xpanes) beschrieben.

# Installieren

Mac:

```
brew install tmux-xpanes
```

CentOS, RHEL:

```
yum install \
https://github.com/greymd/tmux-xpanes/releases/download/v4.1.1/tmux-xpanes_v4.1.1.rpm
```

Ubuntu/Debian:

```
sudo apt install software-properties-common

sudo add-apt-repository ppa:greymd/tmux-xpanes
sudo apt update
sudo apt install tmux-xpanes
```
