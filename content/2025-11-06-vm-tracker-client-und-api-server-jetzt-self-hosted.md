---
title: 'VM-Tracker - Client und API-server jetzt Self-Hosted'
date: 2025-11-06 12:00:00
update: 2025-11-06 12:00:00
author: ruediger
cover: "/images/posts/2025/11/vm-tracker.webp"
# images: 
#   - /images/posts/2025/08/telekom-mail-fail.webp
featureImage: /images/posts/2025/11/vm-tracker.webp
tags: [Qemu, VM, Overview, Self-Hosted]
categories: 
  - Internet
preview: "Jetzt ist alles online und kann installiert werden. So das jeder den VM-Tracker selbst betreiben kann und die eigenen RaspberryPi, VMs, Dedicated-Server, NAS-Systeme oder Shellys Tracken kann."
draft: false
top: false
type: post
hide: false
toc: false
---

# :satellite: VM-Tracker - Downloads für Self-Hosted Installationen verfügbar :sunglasses:

Im [letzten Post](/posts/2025-11-01-vm-tracker-automatische-registration-und-monitoring-system/) hatte ich den VM-Tracker schon einmal vorgestellt und in den letzten Tagen noch selbst einiges getestet. 

Jetzt ist alles online und kann installiert werden. So das jeder den VM-Tracker selbst betreiben kann und die eigenen RaspberryPi, VMs, Dedicated-Server, NAS-Systeme oder Shellys Tracken kann. 

Ja, auch Shellys können sich registrieren und so Überwacht werden. 
Gehe zu Settings → Scripts → Add Script und erstelle:

*Den Domainnamen noch anpassen!*

```javascript
function registerDevice() {
  // Device Info abrufen
  Shelly.call("Shelly.GetDeviceInfo", {}, function(deviceInfo) {
    let hostname = deviceInfo.id || "unknown-shelly";  // z.B. "shellyplus1pm-a1b2c3d4"
    
    Shelly.call("Shelly.GetStatus", {}, function(result) {
      let ip_address = null;
      let interface_name = null;
      
      if (result.eth && result.eth.ip) {
        ip_address = result.eth.ip;
        interface_name = "eth0";
      } else if (result.wifi && result.wifi.sta_ip) {
        ip_address = result.wifi.sta_ip;
        interface_name = "wlan0";
      }
      
      if (!ip_address) {
        print("Keine IP gefunden!");
        return;
      }
      
      Shelly.call(
        "HTTP.POST",
        {
          url: "https://vm-tracker.example.com/api/register",
          content_type: "application/json",
          body: JSON.stringify({
            hostname: hostname,
            ip_address: ip_address,
            interface: interface_name
          })
        },
        function(result, error_code, error_message) {
          if (error_code === 0) {
            print("✓ Registriert: ", hostname, " - ", ip_address);
          } else {
            print("✗ Fehler: ", error_message);
          }
        }
      );
    });
  });
}

// Bei Netzwerk-Verbindung ausführen
Shelly.addEventHandler(function(event) {
  if (event.component === "wifi" && event.info.status === "got ip") {
    print("WiFi verbunden, registriere...");
    Timer.set(2000, false, registerDevice);  // 2s Verzögerung
  }
  if (event.component === "eth" && event.info.status === "up") {
    print("Ethernet verbunden, registriere...");
    Timer.set(2000, false, registerDevice);
  }
});

// Zusätzlich alle 10 Minuten
Timer.set(30000, true, registerDevice);
```

Das Script kann man einfach auf jeden Shelly erstellen und sie melden sich dann einfach an. 

![VM-Tracker registrierter Shelly](/images/posts/2025/11/vm-tracker-shelly.webp)

# VM-Tracker Server installieren. 

Damit man selbst einen Endpunkt für die Clients, unter einer eigenen Domain hat, kann der API-Server jetzt auch Self-Hosted betrieben werden. 

Anleitungen gibt es für: 

* [Binary Installation](https://vm-tracker.kuepper.nrw/docs/de/api/installation_binary/)
* [Docker](https://vm-tracker.kuepper.nrw/docs/de/api/installation_docker/)
* [Kubernetes](https://vm-tracker.kuepper.nrw/docs/de/api/installation_kubernetes/)
* [Helm](https://vm-tracker.kuepper.nrw/docs/de/api/installation_helm/)
* [Kustomization](https://vm-tracker.kuepper.nrw/docs/de/api/installation_kustomization/)
* [FluxCD Kustomization](https://vm-tracker.kuepper.nrw/docs/de/api/installation_fluxcd_kustomization/)
* [FluxCD Helm Release](https://vm-tracker.kuepper.nrw/docs/de/api/installation_fluxcd_helm_release/)

Bei allen Installationen kann man diese ENV-Variabeln setzen:

* API_BASE_URL=https://vm-tracker.example.com
* BASE_URL=https://vm-tracker.example.com

Diese werden benutzt für z.B. Ingress, um die API über die Domain erreichbar zu machen. Ausserdem wird damit das Installations-Skript erstellt, damit die Clients sich mit der richtigen API verbinden. 
Dadurch kann der Client schnell und unkompliziert auf allen Sytemen installiert werden. 

# VM-Tracker Client installieren. 

Welche Möglichkeiten bei der Client Installation zur Verfügung stehen ist hier beschrieben:

* [Binary](https://vm-tracker.kuepper.nrw/docs/de/client/installation_binary/)
* [Skript](https://vm-tracker.kuepper.nrw/docs/de/client/installation_script/)
* [Systemd](https://vm-tracker.kuepper.nrw/docs/de/client/installation_systemd/)
* [Cloud-Init](https://vm-tracker.kuepper.nrw/docs/de/client/installation_cloud_init/)
* [Shelly Script](https://vm-tracker.kuepper.nrw/docs/de/client/installation_shelly/)

Die komplette Dokumentation in deutsch ist [hier](https://vm-tracker.kuepper.nrw/docs/de/) und die englische [hier](https://vm-tracker.kuepper.nrw/docs/en/).


## VM-Manager und Vm-Tracker in Aktion

{{< video "/images/posts/2025/11/vm-tracker.mp4" "my-5" >}}

{{< rawhtml >}} 

<video width=100% controls>
    <source src="/videos/vm-tracker.mp4"type="video/mp4">
    Your browser does not support the video tag.  
</video>

{{< /rawhtml >}}