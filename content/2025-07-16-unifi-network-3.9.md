---
title: 'UniFi Network 9.3: Was ist neu und lohnt sich das Update?'
date: 2025-07-16 09:00:00
update: 2025-07-16 09:00:00
author: ruediger
cover: "/images/posts/2025/07/unifi-dashboard.webp"
featureImage: "/images/posts/2025/07/unifi-dashboard.webp"
tags: [Unifi, UDM, Network, Update]
categories: 
    - Network
preview: "Ubiquiti hat UniFi Network 9.3 veröffentlicht und ich hab mir die neuen Features mal genauer angeschaut. Spoiler: Es gibt einige wirklich nützliche Verbesserungen, auch wenn nicht alles revolutionär ist."
draft: false
top: false
type: post
hide: false
toc: true
---

![Links of the week](/images/posts/2025/07/unifi-dashboard.webp)


Ubiquiti hat UniFi Network 9.3 veröffentlicht und ich hab mir die neuen Features mal genauer angeschaut. Spoiler: Es gibt einige wirklich nützliche Verbesserungen, auch wenn nicht alles revolutionär ist.

## Die neue Client-Tabelle - endlich übersichtlich

Das erste, was mir aufgefallen ist: Die Client-Tabelle wurde komplett überarbeitet. Wer schon mal versucht hat, in einem größeren Netzwerk einen bestimmten Client zu finden, weiß wie nervig das bisher war.

**Was sich geändert hat:**
- Filtern nach Broadcast-Typ, Access Point, Funkband, WiFi-Generation oder Herstellern
- Echtzeitaktualisierungen (endlich!)
- Deutlich schneller, auch bei vielen Clients

Besonders praktisch finde ich die Herstellerfilterung. Wenn du wissen willst, welche Apple-Geräte gerade im Netz sind oder alle Samsung-Smartphones auf einmal anzeigen möchtest, geht das jetzt mit einem Klick.

Die Performance ist tatsächlich spürbar besser geworden. Selbst in meinem Testnetzwerk mit über 200 Clients lädt die Tabelle schnell und reagiert flüssig. Das war früher definitiv ein Schwachpunkt.

## DHCP Manager - überfällige Verbesserung

Der neue DHCP Manager ist ehrlich gesagt längst überfällig gewesen. Bisher musstest du für jeden VLAN separat schauen, welche IP-Adressen vergeben sind. Jetzt siehst du alles zentral an einem Ort.

**Was mir gut gefällt:**
- Alle aktiven Leases in einer Übersicht
- Einfache Verwaltung von statischen Zuweisungen
- Funktioniert auch bei mehreren Netzwerken problemlos

Ich nutze das hauptsächlich, um zu schauen, welche Geräte welche IPs haben und um bei Bedarf statische Zuweisungen zu machen. Geht jetzt deutlich schneller als vorher.

## Alarm Manager - nützlich, aber nicht perfekt

Der neue Alarm Manager ist eine interessante Ergänzung. Du kannst jetzt spezifische Geräte überwachen und dir Benachrichtigungen schicken lassen, wenn bestimmte Bedingungen erfüllt sind.

**Was funktioniert gut:**
- Überwachung einzelner Geräte mit eigenen Regeln
- Verschiedene Benachrichtigungsarten (E-Mail, etc.)
- Export zu externen Systemen möglich

**Wo noch Luft nach oben ist:**
Die Konfiguration ist etwas umständlich und die Dokumentation könnte besser sein. Für einfache Anwendungsfälle (Server offline, hohe Bandbreitennutzung) funktioniert es aber gut.

## System-Logs - deutlich besser geworden

Die System-Logs wurden komplett überarbeitet und sind jetzt wesentlich brauchbarer. Früher war das oft ein Krampf, relevante Informationen zu finden.

**Verbesserungen:**
- Viel mehr Details zu Ereignissen
- Bessere Suchfunktion
- CEF-Format für SIEM-Integration (falls du sowas nutzt)

Besonders für die Fehlersuche ist das eine echte Verbesserung. Die Logs sind jetzt strukturierter und man findet schneller, was man sucht.

## Content-Filterung - endlich granular

Die Content-Filterung wurde erweitert und ist jetzt deutlich flexibler. Du kannst pro Netzwerksegment verschiedene Filter erstellen und sogar Zeitpläne definieren.

**Neue Möglichkeiten:**
- Unbegrenzte Filter pro Netzwerk
- Integrierte Ad-Blocking-Funktion
- Zeitbasierte Regeln

Das ist besonders praktisch, wenn du unterschiedliche Benutzergruppen hast. Gäste-WLAN kann andere Regeln haben als das Mitarbeiter-Netzwerk, und das lässt sich jetzt viel einfacher umsetzen.

## CyberSecure Protection - Marketing oder Mehrwert?

Ubiquiti bewirbt die erweiterten Sicherheitsfeatures ziemlich stark. Die Realität ist: Es gibt tatsächlich Verbesserungen, aber die Grundfunktionen waren auch vorher schon solide.

**Was neu ist:**
- Proofpoint IDS/IPS Integration
- Cloudflare-basierte Echtzeit-Filterung
- Kontinuierliche Threat-Updates

Ob das in der Praxis einen großen Unterschied macht, hängt stark von deinem Anwendungsfall ab. Für die meisten kleineren Netzwerke ist es nice-to-have, aber kein Game-Changer.

## Multi-WAN Support - für die, die es brauchen

Die Multi-WAN-Funktionen wurden erweitert. Du kannst jetzt SLAs definieren und intelligenteres Load-Balancing betreiben.

**Neue Features:**
- Anpassbare SLAs für verschiedene Uplinks
- Policy-basiertes Routing
- Bessere Failover-Mechanismen

Das ist hauptsächlich für Leute interessant, die mehrere Internetverbindungen haben. Für normale Setups mit einem Provider ist das weniger relevant.

## Mein Fazit nach dem Testen

UniFi Network 9.3 ist ein solides Update mit einigen wirklich nützlichen Verbesserungen. Die überarbeitete Client-Tabelle allein macht das Update schon lohnenswert, wenn du regelmäßig mit der Verwaltung zu tun hast.

**Was mir besonders gut gefällt:**
- Client-Tabelle ist endlich brauchbar
- DHCP Manager spart Zeit
- System-Logs sind deutlich informativer

**Was noch verbesserungswürdig ist:**
- Alarm Manager könnte einfacher zu konfigurieren sein
- Manche Features fühlen sich noch etwas "beta" an
- Dokumentation ist teilweise lückenhaft

**Solltest du updaten?**
Ja, aber nicht sofort. Ich würde empfehlen, noch ein paar Wochen zu warten, bis die ersten Kinderkrankheiten ausgebügelt sind. Dann ist es definitiv ein lohnenswertes Update.

Aber die meisten die ich kenne werden es eh jetzt installieren oder schon installiert haben. 

Die meisten Verbesserungen sind praktische Alltagserleichterungen und keine revolutionären Neuerungen. Aber genau das macht ein gutes Update aus - es macht die tägliche Arbeit einfacher, ohne dabei neue Probleme zu schaffen.

Übrigens: Das Update ist kostenlos und funktioniert mit der bestehenden Hardware. Keine versteckten Kosten oder Abo-Fallen, was bei Ubiquiti auch nicht anders zu erwarten war.

Was denkst du über die neuen Features? Hast du schon Erfahrungen mit 9.3 gemacht? Lass es mich in den Kommentaren wissen!

