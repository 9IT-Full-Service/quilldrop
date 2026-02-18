---
title: 'AdBlock mit Pi-Hole ist Selbstverteidigung'
date: 2024-01-14 14:00:00
author: ruediger
cover: "/images/posts/2024/01/adblocker-pihole-protection.webp"
tags: [RaspberryPi, Pi-Hole]
categories: 
    - Internet
preview: "Pi-hole ist eine Netzwerk-basierte Anwendung zur Werbe- und Tracking-Blockierung. Es funktioniert, indem es als DNS-Server in Ihrem Netzwerk agiert und unerwünschte Inhalte, insbesondere Werbung und Tracker, filtert, bevor sie auf Ihre Geräte geladen werden. Hier sind die grundlegenden Schritte, um Pi-hole einzusetzen" 
draft: false
top: false
type: post
hide: false
toc: false
---

<!-- [Englisch Version](/posts/2023-09-21-kubernetes-network-policies-en.html)
-->

![Pi-Hole](/images/posts/2024/01/adblocker-pihole-protection.webp)


Pi-hole ist eine Netzwerk-basierte Anwendung zur Werbe- und Tracking-Blockierung. Es funktioniert, indem es als DNS-Server in Ihrem Netzwerk agiert und unerwünschte Inhalte, insbesondere Werbung und Tracker, filtert, bevor sie auf Ihre Geräte geladen werden. Hier sind die grundlegenden Schritte, um Pi-hole einzusetzen:

1. **Hardware-Voraussetzungen**: Pi-hole kann auf Hardware wie einem Raspberry Pi, einem anderen Einplatinencomputer oder sogar auf einem alten Computer installiert werden. 

2. **Betriebssystem**: Du musst ein kompatibles Betriebssystem auf Deiner Hardware installieren, oft ist dies eine Linux-Distribution.

3. **Installation von Pi-hole**: Die Installation erfolgt normalerweise über eine einfache Befehlszeile. Pi-hole bietet auf ihrer Website eine offizielle Installationsanleitung.

4. **Konfiguration Ihres Netzwerks**: Nach der Installation müssen Sie Ihren Router so konfigurieren, dass er Pi-hole als primären DNS-Server verwendet. Dadurch wird der gesamte Datenverkehr Ihres Netzwerks durch Pi-hole geleitet.

5. **Verwaltung und Wartung**: Pi-hole bietet ein Web-Dashboard, über das Sie die Einstellungen verwalten, Statistiken einsehen und die Blockierlisten anpassen können.

6. **Regelmäßige Updates**: Es ist wichtig, Pi-hole regelmäßig zu aktualisieren, um sicherzustellen, dass es effektiv bleibt und Sicherheitsrisiken minimiert werden.

Pi-hole ist besonders beliebt für seine Effizienz bei der Reduzierung von Werbung auf allen Geräten im Netzwerk, seine relativ einfache Einrichtung und geringe Hardwareanforderungen. Es bietet auch einen erhöhten Datenschutz, da es hilft, Online-Tracking zu reduzieren.

# Pi-Hole installieren

Die Installation von Pi-hole auf einem Raspberry Pi ist ein relativ einfacher Prozess. Hier sind die grundlegenden Schritte:

1. **Raspberry Pi vorbereiten**: 
   - Besorgen Dir einen Raspberry Pi und stellen sicher, dass Du eine SD-Karte, ein Netzteil und eine Netzwerkverbindung (entweder über Ethernet oder WLAN) hast.
   - Installiere ein Betriebssystem auf Deinem Raspberry Pi, üblicherweise wird Raspbian verwendet, das auf Debian basiert.

2. **Betriebssystem einrichten**:
   - Schreibe das Betriebssystem-Image auf die SD-Karte. Dazu kannst Du Software wie BalenaEtcher verwenden.
   - Lege die SD-Karte in den Raspberry Pi ein und starte ihn.
   - Führe die grundlegende Konfiguration durch, wie das Einrichten einer Netzwerkverbindung und das Ändern des Standardpassworts.

3. **Pi-hole installieren**:
   - Öffne einen Terminal auf dem Raspberry Pi.
   - Führe den folgenden Befehl aus: `curl -sSL https://install.pi-hole.net | bash`
   - Dieses Kommando lädt das Installationsskript von Pi-hole herunter und führt es aus.

4. **Pi-hole Konfiguration**:
   - Während der Installation wirst Du durch verschiedene Konfigurationsoptionen geführt, einschließlich der Auswahl eines DNS-Servers, der Einstellung von IPv4 oder IPv6 und der Festlegung von Filterlisten.
   - Nach Abschluss der Konfiguration zeigt das Installationsprogramm die IP-Adresse des Pi-hole-Servers sowie das Passwort für das Web-Administrationsinterface an.

5. **Router-Konfiguration**:
   - Um Pi-hole für Dein gesamtes Netzwerk zu verwenden, muss der Raspberry Pi als primären DNS-Server in Deinem Router eingerichtet werden.
   - Melde Dich sich bei Deinem Router an und suche die DNS-Einstellungen. Hier trägst Du die IP-Adresse des Raspberry Pi ein.

6. **Überprüfen und verwalten**:
   - Nachdem Du Pi-hole als DNS-Server eingestellt hast, kannst Du auf das Pi-hole-Administrations-Dashboard zugreifen, indem Du die IP-Adresse des Raspberry Pi in einem Webbrowser eingibst und Dich anmeldest.
   - Überprüfe, ob die Werbeblockierung funktioniert und passe bei Bedarf die Einstellungen im Dashboard an.

Vergewissere dich, dass regelmäßige Updates für sowohl das Betriebssystem als auch Pi-hole durchgeführt werden, um Sicherheit und Funktionalität zu gewährleisten.

# Sperrelisten hizufügen und verwalten

Um weitere Listen in Pi-hole hinzuzufügen, kkannst Du das Web-Administrationsinterface nutzen. Hier sind die Schritte, um zusätzliche Blocklisten hinzuzufügen:

1. **Pi-hole Webinterface öffnen**: 
   - Gebe die IP-Adresse Deines Raspberry Pi in einem Webbrowser ein, um auf das Pi-hole Dashboard zuzugreifen.
   - Melde Dich an, falls erforderlich.

2. **Zum Blocklisten-Management navigieren**:
   - Im Dashboard findest Du einen Abschnitt oder eine Registerkarte namens "Group Management" oder "Gruppenverwaltung".
   - Klicke darauf und wähle dann „Adlists“ oder „Blocklisten“.

3. **Neue Listen hinzufügen**:
   - Hier kannst Du die URLs der zusätzlichen Blocklisten eingeben. Blocklisten sind in der Regel als URL zu einer Liste im Internet verfügbar, die von verschiedenen Quellen gepflegt werden.
   - Gebe die URL der gewünschten Liste ein und klicke auf „Add“ oder „Hinzufügen“.

4. **Änderungen anwenden**:
   - Nachdem Du neue Listen hinzugefügt hast, ist es wichtig, die Gravity-Datenbank zu aktualisieren, damit die Änderungen wirksam werden. 
   - Dies können Du tun, indem Du im Hauptdashboard auf „Tools“ und dann auf „Update Gravity“ klickst.

5. **Überprüfen und Testen**:
   - Überprüfe nach dem Aktualisieren der Gravity-Datenbank, ob die neuen Listen korrekt hinzugefügt wurden und funktionieren.
   - Du kkannst dies testen, indem Du auf Websites zugreifst, die normalerweise Werbung anzeigen, und überprüfen, ob diese Werbung nun blockiert wird.

6. **Regelmäßige Wartung**:
   - Beachte, dass einige Listen möglicherweise regelmäßig aktualisiert werden müssen, um effektiv zu bleiben. 
   - Pi-hole führt standardmäßig regelmäßige Updates durch, aber Du kannst dies auch manuell im Webinterface unter „Tools“ > „Update Gravity“ durchführen.

Denke, dass das Hinzufügen vieler Listen die Leistung Deines Pi-hole beeinträchtigen kann, besonders wenn Dein Raspberry Pi über begrenzte Ressourcen verfügt. Es ist auch möglich, dass zu viele Listen zu falsch-positiven Ergebnissen führen, bei denen legitime Websites fälschlicherweise blockiert werden. Es ist daher wichtig, ein Gleichgewicht zu finden und Listen auszuwählen, die Deinen spezifischen Anforderungen und Präferenzen entsprechen.

# Listen mit denen gute Erfahrungen gemacht wurden

Ende 2022 habe ich von [Frank Plaschke](https://about.me/Frank.Plaschke) eine Seite geschickt bekommen. 
Auf [Firebog.net](https://firebog.net) sind eine Menge Sperrlisten für Pi-Hole gelistet, die Pi-Hole bei ihm sehr effizient gemacht haben. Ich selbst habe sie bei mir auch eingetragen und Pi-Hole wurde sehr viel besser. 

