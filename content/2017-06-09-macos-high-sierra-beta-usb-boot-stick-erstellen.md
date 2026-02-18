---
title: 'MacOs High Sierra Beta USB Boot-Stick erstellen'
date: 2017-06-09 11:52:17
update: 2017-06-09 11:52:17
author: ruediger
cover: "/images/cat/technik.webp"
tags: 
    - Apple
    - Application
    - Boot
    - High Sierra
    - Installation
    - Internet
    - MacOS
    - mount
    - Partition
    - USB-Stick
preview: 'MacOs High Sierra Beta USB Boot-Stick erstellen'
categories: 
    - Technik
toc: false
hide: false
type: post
---


   * Lade das High Sierra Beta-Installationsprogramm herunter und stell sicher, dass es sich im / Applications-Ordner befindet. Dies ist der Standard-Download-Ort vom Mac App Store.
   * Einen >=8GB USB-Stick einstecken. Wenn der Stick nicht bereits als GUID Partition Map und Mac OS Extended (Journaled) formatiert ist, starte die Festplatten-Utility-Anwendung und formatiere den Stick. Dadurch werden alle Daten vom Laufwerk gelöscht. 
   <!--more-->
   * Öffne ein Terminal-Fenster und füge folgenden Befehl ein, um den Beta-Installer auf den USB zu verschieben und ihn bootfähig zu machen: 
   
```
sudo /Applications/Install\ macOS\ 10.13\ Beta.app/Contents/Resources/createinstallmedia --volume /Volumes/USB --applicationpath /Applications/Install\ macOS\ 10.13\ Beta.app --nointeraction
```
   * Gebe dein Ihr Benutzer-Passwort ein, wenn dazu aufgefordert wird. Das kopieren wird gestartet. Das USB-Laufwerk wird während des gesamten Prozesses aus gehangen und wird nicht auf dem Desktop angezeigt. Auf dem Terminal wird der Fortschritt angezeigt. Sobald die Dateien kopiert wurden und das Laufwerk bootfähig gemacht wurde, wird das Laufwerk wieder auf dem Desktop erscheinen und das Terminal wird angezeigt, dass der Prozess abgeschlossen ist.
