---
title: "Rasbian ohne Monitor und Tastatur installieren"
date: 2019-09-03 18:37:42
update: 2019-09-03 18:37:42
author: ruediger
cover: "/images/cat/technik.webp"
tags:
  - RaspberryPi
  - Raspbian
  - Installation
  - WiFi
  - Headless
  - silentinstall
preview: "Einen neuen Raspbarry mit Rasbian ohne Monitor und Tastatur einrichten"
categories: 
  - Internet
toc: false
hide: false
type: post
---


Ich bestelle immer RaspberryPi Bundles wie das hier [UCreate Raspberry Pi 3 Model B+ Desktop Starter Kit (16 GB, schwarz)](https://www.amazon.de/gp/product/B07BNPZVR7/). Eigentlich immer mit einem vorinstallierten Rasbian. Dieses mal war aber kein Rasbian vorinstalliert sondern Noobs auf der SD-Card.

Nur stecke ich immer ein Netzwerkkabel an und Strom, warte kurz bis der RaspberryPi im Netz ist und logge mich dann per SSH dauf dem RaspberryPi ein. Das einzige was ich sonst immer mache ist im ROOT auf der SD-Card eine leere Datei `ssh` anlegen.

Um jetzt Headless zu installieren geht man wie folgt vor:

* In das Verzeichnis `/os` wechseln und alle Distributions Verzeichnisse löschen die man nicht benötigt. In diesem Fall bleibt dann nur das Verzeichnis `/os/Rasbian_Full` übrig.
* Im Root Verzeichnis die Datei `recovery.cmdline` öffnen und an das Ende `silentinstall` anhängen.
* Im Root Verzeichnis: `touch ssh`

Vorher:
```
runinstaller quiet vt.cur_default=1 coherent_pool=6M elevator=deadline
```
Nachher:
```
runinstaller quiet vt.cur_default=1 coherent_pool=6M elevator=deadline silentinstall
```

SD-Card einstecken und Strom anschliessen. Dann ist Kaffee holen angesagt. Die Installation dauert ein paar Minuten.
Wenn der RaspberryPi fertig ist sollte er wie gewohnt im Netz auftauchen und man kann sich drauf einloggen.

# Kleines WLAN Problem

Dieser RaspberryPi soll später mobil per WiFi über einen LTE Router ins Internet und per Kabel soll ein Kassendrucker angeschlossen werden. Da wpa_supplicant noch nicht konfiguriert war, aber schon gestartet ist hatte das noch kurz Probleme gemacht.

Die Datei `/etc/wpa_supplicant/wpa_supplicant.conf` wurde angepasst:

```
cat /etc/wpa_supplicant/wpa_supplicant.conf
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
ap_scan=1
update_config=1
network={
	ssid="Section3"
  psk="1234123412341234"
}
```

Kurzer Test mit laufendem `tail -f /var/log/message &` hat erst einmal überhaupt nicht funktioniert.

```
iw wlan0 info
Interface wlan0
	ifindex 3
	wdev 0x1
	addr b8:27:eb:87:71:dc
	type managed
	wiphy 0
	channel 11 (2462 MHz), width: 40 MHz, center1: 2452 MHz
	txpower 31.00 dBm
```

Läuft doch. Aber wieso klappt die Verbindung nicht? Ein Scan mit `iwlist wlan0 scan` hat auch geklappt, die gewohnt lange Liste der SSID hier wird ausgegeben.  

```
# > wpa_supplicant -i wlan0 -Dnl80211 -c /etc/wpa_supplicant/wpa_supplicant.conf
Successfully initialized wpa_supplicant
Failed to create interface p2p-dev-wlan0: -16 (Device or resource busy)
nl80211: Failed to create a P2P Device interface p2p-dev-wlan0
P2P: Failed to enable P2P Device interface
Sep  3 20:27:17 raspberrypi kernel: [ 2046.293855] brcmfmac: brcmf_cfg80211_add_iface:
iface validation failed: err=-16
Sep  3 20:27:17 raspberrypi kernel: [ 2046.293855] brcmfmac: brcmf_cfg80211_add_iface:
iface validation failed: err=-16
wlan0: Trying to associate with 90:de:d0:d0:f1:98 (SSID='Section3' freq=2462 MHz)
wlan0: Associated with 90:de:d0:d0:f1:98
wlan0: CTRL-EVENT-DISCONNECTED bssid=90:de:d0:d0:f1:98 reason=0 locally_generated=1
wlan0: WPA: 4-Way Handshake failed - pre-shared key may be incorrect
wlan0: CTRL-EVENT-SSID-TEMP-DISABLED id=0 ssid="Section3" auth_failures=1 duration=10
reason=WRONG_KEY
wlan0: CTRL-EVENT-REGDOM-CHANGE init=CORE type=WORLD
wlan0: CTRL-EVENT-REGDOM-CHANGE init=USER type=COUNTRY alpha2=US
wlan0: CTRL-EVENT-SSID-REENABLED id=0 ssid="Section3"
wlan0: Trying to associate with 90:de:d0:d0:f1:98 (SSID='Section3' freq=2462 MHz)
Sep  3 20:27:31 raspberrypi kernel: [ 2060.046763] brcmfmac: brcmf_cfg80211_escan:
Connecting: status (3)
Sep  3 20:27:31 raspberrypi kernel: [ 2060.046777] brcmfmac: brcmf_cfg80211_scan:
scan error (-11)
Sep  3 20:27:31 raspberrypi kernel: [ 2060.046763] brcmfmac: brcmf_cfg80211_escan:
Connecting: status (3)
Sep  3 20:27:31 raspberrypi kernel: [ 2060.046777] brcmfmac: brcmf_cfg80211_scan:
scan error (-11)
```

Alles noch mal überprüft. Die SSID stimmt, PSK stimmt, PSK auch noch mal verschlüsselt hinterlegt. Keine Verbindung mit dem WiFi möglich.

Dann mal checken ob vielleicht noch etwas läuft:

```
#> ps fauxww | grep wpa
root       468  0.0  0.3  10156  2864 ?        Ss   19:53   0:02 wpa_supplicant -B
-c/etc/wpa_supplicant/wpa_supplicant.conf -iwlan0 -Dnl80211,wext
root      2454  0.0  0.0   4372   572 pts/0    S+   20:38   0:00                  
        \_ grep wpa
```

Ah, ok. Da ist noch ein alter wpa_supplicant gestartet, der den Treiber falsch hatte.
Der wird dann einfach mal gekillt.
```
# > kill -9 468
```

Check ob er wirklich beendet ist:
```
# > ps fauxww | grep wpa
root      2456  0.0  0.0   4372   564 pts/0    S+   20:38   0:00                   
       \_ grep wpa
```

Alles klar und jetzt wpa_supplicant noch mal starten. Und siehe da, jetzt klappt es.

```
# > wpa_supplicant -i wlan0 -Dnl80211 -c /etc/wpa_supplicant/wpa_supplicant.conf
Successfully initialized wpa_supplicant
wlan0: Trying to associate with 90:de:d0:d0:f1:98 (SSID='Section3' freq=2462 MHz)
wlan0: Associated with 90:de:d0:d0:f1:98
wlan0: WPA: Key negotiation completed with 90:de:d0:d0:f1:98 [PTK=CCMP GTK=TKIP]
wlan0: CTRL-EVENT-CONNECTED - Connection to 90:de:d0:d0:f1:98 completed [id=0 id_str=]
Sep  3 20:38:28 raspberrypi kernel: [ 2717.234843] IPv6: ADDRCONF(NETDEV_CHANGE):
wlan0: link becomes ready
Sep  3 20:38:28 raspberrypi dhcpcd[414]: wlan0: carrier acquired
Sep  3 20:38:28 raspberrypi dhcpcd[414]: wlan0: IAID eb:87:71:dc
Sep  3 20:38:28 raspberrypi dhcpcd[414]: wlan0: adding address fe80::997b:1be8:e6e0:f5c8
Sep  3 20:38:28 raspberrypi dhcpcd[414]: wlan0: soliciting a DHCP lease
Sep  3 20:38:29 raspberrypi dhcpcd[414]: wlan0: soliciting an IPv6 router
Sep  3 20:38:29 raspberrypi dhcpcd[414]: wlan0: offered 10.0.2.125 from 10.0.2.1
Sep  3 20:38:29 raspberrypi dhcpcd[414]: wlan0: probing address 10.0.2.125/24
Sep  3 20:38:30 raspberrypi avahi-daemon[318]: Joining mDNS multicast group on interface
wlan0.IPv6 with address fe80::997b:1be8:e6e0:f5c8.
Sep  3 20:38:30 raspberrypi avahi-daemon[318]: New relevant interface wlan0.IPv6 for mDNS.
Sep  3 20:38:30 raspberrypi avahi-daemon[318]: Registering new address record for
fe80::997b:1be8:e6e0:f5c8 on wlan0.*.
Sep  3 20:38:34 raspberrypi dhcpcd[414]: wlan0: leased 10.0.2.125 for 3600 seconds
Sep  3 20:38:34 raspberrypi avahi-daemon[318]: Joining mDNS multicast group on interface
wlan0.IPv4 with address 10.0.2.125.
Sep  3 20:38:34 raspberrypi avahi-daemon[318]: New relevant interface wlan0.IPv4 for mDNS.
Sep  3 20:38:34 raspberrypi dhcpcd[414]: wlan0: adding route to 10.0.2.0/24
Sep  3 20:38:34 raspberrypi avahi-daemon[318]: Registering new address record for
10.0.2.125 on wlan0.IPv4.
Sep  3 20:38:34 raspberrypi dhcpcd[414]: wlan0: adding default route via 10.0.2.1
```

Damit das ganze auch noch nach dem Start funktioniert auch noch systemd konfiguriert.

`vim /lib/systemd/system/wpa_supplicant@wlan0.service`

```
[Unit]
Description=WPA-Supplicant-Daemon (wlan0)
Requires=sys-subsystem-net-devices-wlan0.device
BindsTo=sys-subsystem-net-devices-wlan0.device
After=sys-subsystem-net-devices-wlan0.device
Before=network.target
Wants=network.target

[Service]
Type=simple
RemainAfterExit=yes
ExecStart=/sbin/wpa_supplicant -qq -c/etc/wpa_supplicant/wpa_supplicant.conf -Dnl80211
-iwlan0
Restart=on-failure

[Install]
Alias=multi-user.target.wants/wpa_supplicant@wlan0.service
```

Und anschliessend aktivieren:

```
# > systemctl daemon-reload
# > systemctl enable wpa_supplicant@wlan0.service
# > systemctl start wpa_supplicant@wlan0.service
# > reboot
```

Das Netzwerkkabel entfernt und der RaspberryPi kam wieder über WiFi online.
Per SSH eingeloggt und den Rest gemacht. User pi deaktiviert, neuen User angelegt und weitere Software installiert.
Das was man halt so üblicherweise macht um den PI Safe und fertig zu bekommen.

Um das WiFi Password nicht im Klartext in der Konfiguration zu haben:

```
# > wpa_passphrase "WLAN-NAME" "123412341234" | grep -v "#"
network={
	ssid="WLAN-NAME"
	psk=c6e5b342b0bc6fe1aff18ee420ee5adbed3e8bfa1a7d5da9a7c7585fe0446fc0
}
# also:
# > wpa_passphrase "WLAN-NAME" "123412341234" | grep -v "#" >> \
/etc/wpa_supplicant/wpa_supplicant.conf
```

Wer WiFi schon bei der Headless Installation konfigurieren möchte lege die Datei
`wpa_supplicant.conf` einfach in das `/boot` Verzeichnis der SD-Card vor dem ersten Start.

```
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
ap_scan=1
update_config=1
country=DE
network={
	ssid="WLAN-NAME"
  psk=c6e5b342b0bc6fe1aff18ee420ee5adbed3e8bfa1a7d5da9a7c7585fe0446fc0
}
```

Dann ist WiFi auch nach der Installation sofort fertig konfiguriert.

# Power Management ausschalten

Da der RaspberryPi das WiFi Interface nicht abschalten soll muss noch das Power-Management
ausgeschaltet werden.

```
# > iwconfig wlan0 | grep Power
          Bit Rate=135 Mb/s   Tx-Power=31 dBm
          Power Management:on
```

Dazu in der Datei `/etc/rc.local` einfach vor dem `exit 0;` folgendes hinzufügen:

```
iwconfig wlan0 power off
```

Nach dem Reboot sollte auch das Power-Management ausgeschaltet sein:

```
> # iwconfig wlan0 | grep Power
          Bit Rate=135 Mb/s   Tx-Power=31 dBm
          Power Management:off
```



