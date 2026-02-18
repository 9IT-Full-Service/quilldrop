---
title: 'yazi lazygit neovim'
date: 2024-12-15 20:00:00
author: ruediger
cover: "/images/posts/2024/12/tmux.webp"
featureImage: "/images/posts/2024/12/tmux.webp"
tags: 
  - MacOS
  - iTerm
  - tmux
categories: 
  - MacOS
preview: "Meine tmux konfiguration."
draft: true
top: false
type: post
hide: true
toc: false
---

![iTerm zsh](/images/posts/2024/12/tmux.webp)

## gh cli Tool installieren

    brew install gh 

    gh auth login

    git checkout -b feature-branch
    git add .
    git commit -m "Deine Nachricht für den Commit"

    git push origin feature-branch

    gh pr create --base main --title "Titel des PR" --body "Beschreibung des PR"
    


## tmux installieren 

Installation von tmux mit homebrew 

    brew install tmux

## Konfiguration 

Damit Änderungen in der `tmux.conf` für alle Sessions, Windows und Panes neugeladen werden können wird als ersten `bind r ` konfiguriert:  

    unbind r
    bind r source-file ~/.tmux.conf

Der Standard Prefix ist default auf `control + b` und ich finde ihn auf `control +s` angenehmer. Was auch an meiner Zeit mit `screen` liegt. 

    set -g prefix C-s

Aktivieren der Mouse, damit auch die Panes mit der Maus angepasst werden können: 

    set -g mouse on

Um mit den Tasten `control + h,j,k,l` durch die Panes springen zu können: 

    setw -g mode-keys vi
    bind-key h select-pane -L
    bind-key j select-pane -D
    bind-key k select-pane -U
    bind-key l select-pane -R

Für die Plugins in tmux benutze ich den tmux plugin manager - tpm. 

    set -g @plugin 'tmux-plugins/tpm'

Für die Statusleiste hatte ich Dracula getestet, dann aber `tmux2k` gefunden. Das Plugin wird einfach mit folgender Zeile aktiviert. 

    set -g @plugin '2kabhishek/tmux2k'

Die Konfiguration für tmux2k: 

    set -g @tmux2k-theme 'onedark icons'
    set -g @tmux2k-icons-only true
    set -g @tmux2k-left-plugins "git cpu-usage ram-usage"
    set -g @tmux2k-right-plugins "battery cpu ram git time"
    set -g @tmux2k-network-name "en0"
    set -g @tmux2k-show-powerline true
    set -g @tmux2k-show-fahrenheit false
    set -g @tmux2k-military-time true
    set -g @tmux2k-border-contrast true

    # available colors: white, gray, dark_gray, light_purple, dark_purple, cyan, green, orange, red, pink, yellow
    # set -g @tmux2k-[plugin-name]-colors "[background] [foreground]"
    set -g @tmux2k-cpu-usage-colors "blue dark_gray"

    # it can accept `session`, `rocket`, `window`, or any character.
    set -g @tmux2k-show-left-icon ""

    # update powerline symbols
    set -g @tmux2k-show-left-sep ""
    set -g @tmux2k-show-right-sep ""

    # change refresh rate
    set -g @tmux2k-refresh-rate 5

Die Statusbar habe ich für Tmux gerne oben, da meine ZSH Statusbar immer unten ist wirkt das aufgräumter. 
Ausserdem sehe ich so auch oben immer die wichtigen Infos, wie welches Window aktiv ist, Git Branch usw. 

    set -g status-position top

Jetzt noch tpm aktivieren und schon kann man einen Reload der Konfiguration machen: 

    run '~/.tmux/plugins/tpm/tpm'

Wenn noch keine tmux Session vorhanden, einfach eine neue aufmachen. Ist schon eine vorhanden kann man jetzt mit `control + r` die Konfiguration laden und anschliessend noch tpm die Plugins installieren lassen mit `control + I`. 

Die Plugins werden installiert und die Statusleiste sollte jetzt aktiv und zu sehen sein. 

Die komplette Konfiguration: 

    unbind r
    bind r source-file ~/.tmux.conf

    set -g prefix C-s
    set -g mouse on
    setw -g mode-keys vi
    bind-key h select-pane -L
    bind-key j select-pane -D
    bind-key k select-pane -U
    bind-key l select-pane -R

    set -g @plugin 'tmux-plugins/tpm'
    set -g @plugin '2kabhishek/tmux2k'
    set -g @tmux2k-theme 'onedark icons'
    set -g @tmux2k-icons-only true
    set -g @tmux2k-left-plugins "git cpu-usage ram-usage"
    set -g @tmux2k-right-plugins "battery cpu ram git time"
    set -g @tmux2k-network-name "en0"
    set -g @tmux2k-show-powerline true
    set -g @tmux2k-show-fahrenheit false
    set -g @tmux2k-military-time true
    set -g @tmux2k-border-contrast true

    # available colors: white, gray, dark_gray, light_purple, dark_purple, cyan, green, orange, red, pink, yellow
    # set -g @tmux2k-[plugin-name]-colors "[background] [foreground]"
    set -g @tmux2k-cpu-usage-colors "blue dark_gray"

    # it can accept `session`, `rocket`, `window`, or any character.
    set -g @tmux2k-show-left-icon ""

    # update powerline symbols
    set -g @tmux2k-show-left-sep ""
    set -g @tmux2k-show-right-sep ""

    # change refresh rate
    set -g @tmux2k-refresh-rate 5

    set -g status-position top

    run '~/.tmux/plugins/tpm/tpm'


