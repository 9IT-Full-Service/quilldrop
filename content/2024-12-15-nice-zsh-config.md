---
title: 'Nice Terminal und zsh Konfiguration'
date: 2024-12-15 16:00:00
author: ruediger
cover: "/images/posts/2024/12/zsh.webp"
tags: [MacOS, iTerm, zsh, nerdfonts]
categories: 
    - MacOS
preview: "Nice Terminal und zsh Konfiguration mit oh-my-zsh, Powerlevel10k, Plugins und weiteren Tools."
draft: false
top: false
type: post
hide: false
toc: false
---

![iTerm zsh](/images/posts/2024/12/zsh.png)

## iTerm2 installieren und konfigurieren

Als erstes ersetzen wir den MacOS Terminal und installieren iterm2 

    brew install iterm2 

Für iTerm2 benutze ich das Gruvebox Colorscheme: 

    curl -so gruvebox.itermcolors https://raw.githubusercontent.com/mbadolato/iTerm2-Color-Schemes/master/schemes/GruvboxDark.itermcolors

Das Color-Scheme importiert man in den Settings von iTerm2 unter Profiles -> Colors -> Color Presets (unten rechts).

![Import colorscheme](/images/posts/2024/12/iterm-colors.png)

Anschliessend kann man dort dann "gruvebox" auswählen. 

Als nächstes brauchen wir noch Fonts. Die bekommt man auf Nerdfonts.com. Ich benutze den Font "Nerd Hack Fonts":

    curl -sLo Hack.zip https://github.com/ryanoasis/nerd-fonts/releases/download/v3.3.0/Hack.zip

Die Datei entpacken und den Font einfach mit Doppelklick öffnen. Unter MacOS ist damit die Font Installation schon erledigt. 

Den Font kann an dann in den Settings von iTerm2 direkt auswählen und aktivieren. Profiles -> Text -> Font.


## ZSH konfigurieren

Wer noch kein ZSH installiert hat kann ZSH mit `brew install zsh` installieren. 

oh-my-zsh managed die ZSH Konfiguration und bringt eine Menge an Plugins, Themes und noch mehr mit um ZSH so anpassen zu können wir man es benötigt. Egal Aussehen oder Funktionen. 

    sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

Das powerlevel10k Plugin installieren:

    git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k

In der .zshrc kann man das Theme auf anschliessend auf Powerlevel10k setzen:

    ZSH_THEME="powerlevel10k/powerlevel10k"

Jetzt noch etwas die zsh Shell pimpen. Dazu werden jetzt noch zsh-syntax-highlighting und zsh-autosuggestions installiert.

    brew install zsh-autosuggestions
    brew install zsh-syntax-highlighting

Beide aktiveren, in dem man einfach folgende Zeilen in der .zshrc am Ende hinzufügt: 

    source $(brew --prefix)/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
    source $(brew --prefix)/share/zsh-autosuggestions/zsh-autosuggestions.zsh

In der Datei `.zshrc` ist ziemlich am Anfang eine Variable `plugins`, bei dieser fügt man jetzt dir beiden Plugins hinzu.

Bei mir sieht die Zeile so aus:

    plugins=(command-not-found web-search git git-prompt systemadmin taskwarrior svn perl macos screen terminitor lol emoji-clock themes history zsh-syntax-highlighting zsh-autosuggestions)

Damit die File- und Ordnerliste so aussehen wir oben im Screenshot benutze ich colorls, welches mit Ruby gem installiert wird. 

    sudo gem install colorls

Ggf. muss vorher noch ein `gem install public_suffix -v 5.1.1` ausgeführt werden. Das wird aber auch in dem Fall beim install von colorls angezeigt. 

Damit ist dann auch `colorls` als Befehl nutzbar. Wenn man jetzt aber nicht immer colorls eingeben möchte, kann man sich einfach in der `.zshrc` einfach einen Alias dafür anlegen: 

    alias ls="colorls"

Jetzt kann man entweder einen neuen Termin öffnen oder man gibt einfach `source ~/.zshrc` ein um die Konfiguration einlesen zu lassen. 

## Noch mehr mit Powerlevel10k 

![zsh powerlevel10k](/images/posts/2024/12/p10k.png)

Hier sieht man meinen kompletten ZSH Prompt. Hier werden neben der Uhrzeit, Batterie, interne und externe IP, Akku jede Menge Sachen angezeigt. Wenn ich in einem Verzeichnis mit z.B. mit einer Kubeconfig bin und/oder einem Git Repository, wird man der Kuberentes Context, Git Branch usw. angezeigt. 

Die Konfig ist etwas größer, daher habe ich sie in einen [GitHub Gist](https://gist.github.com/ruedigerp/670abd2bfecba6c4cd481b7cf4352570) hinterlegt. 

Die Powerlevel10k Konfiguration habe ich in der Datei `~/.p10k.zsh` gespeichert. Und in der `.zshrc` wird dann anschliessend noch eingefügt: 

    # Enable Powerlevel10k instant prompt. Should stay close to the top of ~/.zshrc.
    # Initialization code that may require console input (password prompts, [y/n]
    # confirmations, etc.) must go above this block; everything else may go below.
    if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
        source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
    fi

    source $(brew --prefix)/share/powerlevel10k/powerlevel9k.zsh-theme

    # To customize prompt, run `p10k configure` or edit ~/.p10k.zsh.
    [[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh

    # typeset -g POWERLEVEL9K_INSTANT_PROMPT=quiet
    typeset -g POWERLEVEL9K_INSTANT_PROMPT=off

Jetzt noch mal einen neuen Terminal öffnen oder wieder ein `source ~/.zshrc`. 


