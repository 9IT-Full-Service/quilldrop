---
title: 'Schnell vorkonfigurierte VMs mit QEMU erstellen'
date: 2025-07-05
update: 2025-07-05
author: ruediger
cover: "/images/posts/2025/07/BC13E6DB-AC08-4445-A5E3-65CA6C1A2EE4.webp"
featureImage: "/images/posts/2025/07/BC13E6DB-AC08-4445-A5E3-65CA6C1A2EE4.webp"
tags: [VM, server, cloudinit]
categories: 
    - DevOps
preview: "Wer kennt das nicht? Man braucht mal eben eine saubere Testumgebung, will ein neues Tool ausprobieren oder ein Kubernetes-Cluster aufsetzen. Normalerweise bedeutet das: VM aufsetzen, OS installieren, Updates fahren, Tools installieren – und schon sind ein paar Stunden weg."
draft: false
top: false
type: post
hide: false
toc: false
---

![VM mit Qemu und Cloudinit erstellen](/images/posts/2025/07/BC13E6DB-AC08-4445-A5E3-65CA6C1A2EE4.webp)


Wer kennt das nicht? Man braucht mal eben eine saubere Testumgebung, will ein neues Tool ausprobieren oder ein Kubernetes-Cluster aufsetzen. Normalerweise bedeutet das: VM aufsetzen, OS installieren, Updates fahren, Tools installieren – und schon sind ein paar Stunden weg.

Mit QEMU und Cloud-Init geht das deutlich eleganter. Einmal konfiguriert, startet man eine vollständig vorkonfigurierte VM in wenigen Minuten. Hier zeige ich, wie das geht.

## Warum QEMU und Cloud-Init?

QEMU ist ein mächtiger Virtualisierer, der auf praktisch allen Plattformen läuft. Cloud-Init ist das Schweizer Taschenmesser für VM-Konfiguration – es kann beim ersten Boot automatisch User anlegen, SSH-Keys installieren, Software nachinstallieren und sogar komplette Skripte ausführen.

Die Kombination macht’s möglich: VM starten, kurz warten, fertig konfigurierte Umgebung nutzen.

## QEMU installieren

**macOS (mit Homebrew):**

```bash
brew install qemu
```

**Ubuntu/Debian:**

```bash
sudo apt update
sudo apt install qemu-system-aarch64 qemu-utils
```

**CentOS/RHEL/Fedora:**

```bash
sudo dnf install qemu-system-aarch64 qemu-img
```

## Das Base-Image besorgen

Wir verwenden ein fertiges Debian ARM64 Cloud-Image als Basis:

```bash
wget https://cloud.debian.org/images/cloud/bookworm/latest/debian-12-generic-arm64.qcow2
```

## Working Copy erstellen

Das Original-Image behalten wir als Template und erstellen eine Arbeitskopie:

```bash
cp debian-12-generic-arm64.qcow2 debian-testserver.qcow2
qemu-img resize debian-testserver.qcow2 20G
```

So kann man das Original-Image immer wieder für neue VMs verwenden.

## Cloud-Init konfigurieren

Cloud-Init braucht drei Dateien, die wir in einem eigenen Ordner sammeln:

```bash
mkdir cloud-init-data
cd cloud-init-data/
```

### user-data - Das Herzstück

Hier passiert die ganze Magie. Diese Datei definiert, wie die VM aussehen soll:

```yaml
cat <<EOF > user-data
#cloud-config

# User anlegen
users:
  - name: debian
    plain_text_passwd: testpass
    lock_passwd: false
    sudo: ALL=(ALL) NOPASSWD:ALL
    shell: /bin/bash
    ssh_authorized_keys:
      - ssh-ed25519 AAAAC3...DeinSSHKey...7iFVL

# SSH mit Passwort erlauben
ssh_pwauth: true

# Debug-Modus für Troubleshooting
debug: true

# Standard-Pakete installieren
packages:
  - htop
  - curl
  - git
  - vim
  - wget

# System updaten
package_update: true
package_upgrade: true

# Installationsskript vorbereiten
write_files:
  - path: /tmp/install-k3s.sh
    permissions: '0755'
    content: |
      #!/bin/bash
      set -e
      
      echo "K3s wird installiert..."
      curl -sfL https://get.k3s.io | sh -s - server --cluster-init
      echo "export KUBECONFIG=/etc/rancher/k3s/k3s.yaml" >> /root/.bashrc
      
      echo "k9s wird installiert..."
      wget https://github.com/derailed/k9s/releases/download/v0.32.7/k9s_Linux_arm64.tar.gz
      tar xzf k9s_Linux_arm64.tar.gz
      mv k9s /usr/local/bin/
      
      echo "Helm wird installiert..."
      curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
      chmod 700 get_helm.sh
      ./get_helm.sh
      
      echo "FluxCD CLI wird installiert..."
      curl -s https://fluxcd.io/install.sh | sudo bash
      
      echo "kubectl wird installiert..."
      curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl"
      sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
      
      # Praktische Aliases
      echo 'alias k="kubectl"' >> /root/.bashrc
      echo 'alias kgp="kubectl get pods"' >> /root/.bashrc
      
      # Warten bis Kubernetes läuft
      echo "Warte auf Kubernetes API..."
      while ! kubectl cluster-info &>/dev/null; do
        echo "Kubernetes startet noch..."
        sleep 10
      done
      
      echo "K3s Installation abgeschlossen!"

# Skript nach der Installation ausführen
runcmd:
  - /tmp/install-k3s.sh
EOF
```

### meta-data - Server-Metadaten

```bash
cat <<EOF > meta-data
instance-id: server-001
local-hostname: testserver
EOF
```

### network-config - Netzwerk festlegen

```bash
cat <<EOF > network-config
version: 2
ethernets:
  enp0s1:
    addresses: [192.168.74.10/24]
    gateway4: 192.168.74.1
    nameservers:
      addresses: [8.8.8.8, 1.1.1.1]
    dhcp4: false
EOF
```

## Cloud-Init ISO erstellen

Aus den drei Dateien basteln wir ein ISO-Image:

### macOS

```bash
cd ..
hdiutil makehybrid -iso -joliet -default-volume-name "cidata" -o seed.iso cloud-init-data/
```

### Linux

```bash
cd ..
# Mit genisoimage (meist vorinstalliert)
genisoimage -output seed.iso -volid cidata -joliet -rock cloud-init-data/

# Oder mit xorriso
xorriso -as mkisofs -V cidata -o seed.iso cloud-init-data/
```

## VM starten

Jetzt kommt der spannende Teil – die VM starten:

```bash
qemu-system-aarch64 \
    -name "testserver" \
    -machine type=virt,accel=hvf \
    -cpu cortex-a72 \
    -smp cores=4,threads=1 \
    -m 4G \
    -drive file=debian-testserver.qcow2,if=virtio,index=0,media=disk,format=qcow2 \
    -drive file=seed.iso,if=virtio,index=1,media=cdrom \
    -netdev user,id=net0,hostfwd=tcp::2222-:22,hostfwd=tcp::8080-:80 \
    -device virtio-net-pci,netdev=net0 \
    -bios /opt/homebrew/share/qemu/edk2-aarch64-code.fd \
    -nographic \
    -serial mon:stdio
```

**Was passiert hier?**

- `-accel=hvf`: Hardware-Beschleunigung (macOS), unter Linux nimmt man `kvm`
- `-m 4G`: 4 GB RAM
- `-smp cores=4`: 4 CPU-Kerne
- `hostfwd=tcp::2222-:22`: SSH über Port 2222 erreichbar
- `hostfwd=tcp::8080-:80`: HTTP über Port 8080 erreichbar
- `-nographic`: Läuft in der Konsole

## Netzwerk anpassen

Je nach Einsatzzweck kann man verschiedene Netzwerk-Modi nutzen:

### User-Modus (Standard)

```bash
-netdev user,id=net0,hostfwd=tcp::2222-:22
```

- Einfachste Variante
- VM kann ins Internet, ist aber von außen nicht direkt erreichbar
- Perfekt für Tests

### Bridge-Modus

```bash
-netdev bridge,id=net0,br=br0
```

- VM bekommt IP aus dem Host-Netzwerk
- Direkte Kommunikation möglich
- Braucht Bridge-Setup auf dem Host

### Host-Only

```bash
-netdev socket,id=net0,listen=:1234
```

- VM nur vom Host erreichbar
- Maximale Isolation
- Gut für Sicherheitstests

## Nach dem Start

### Per SSH einloggen

```bash
ssh -p 2222 debian@localhost
```

### Direkte Konsole nutzen

Falls SSH mal nicht klappt:

- `Ctrl+A, C` für QEMU-Monitor
- `info network` zeigt Netzwerk-Status
- `Ctrl+A, X` beendet QEMU

## Wenn’s nicht läuft

### Cloud-Init checken

```bash
# In der VM
sudo cloud-init status --wait
sudo cloud-init logs
```

### Netzwerk prüfen

```bash
# In der VM
ip addr show
ping google.com
```

### Services testen

```bash
# K3s Status
sudo systemctl status k3s
kubectl get nodes
```

## Erweiterte Tricks

### Mehrere VMs parallel

Einfach verschiedene Cloud-Init-Konfigurationen erstellen:

```bash
# Zweite VM
cp debian-12-generic-arm64.qcow2 debian-testserver-2.qcow2
# Neue cloud-init-data-2/ mit angepassten Einstellungen
# Andere Ports verwenden: 2223, 8081, etc.
```

### Automatisierung

Ein kleines Skript macht das Leben leichter:

```bash
#!/bin/bash
# vm-create.sh
VM_NAME=$1
SSH_PORT=$2
HTTP_PORT=$3

echo "Erstelle VM: $VM_NAME"
cp debian-12-generic-arm64.qcow2 $VM_NAME.qcow2

# Cloud-Init anpassen
mkdir cloud-init-$VM_NAME
sed "s/testserver/$VM_NAME/g" cloud-init-data/meta-data > cloud-init-$VM_NAME/meta-data
cp cloud-init-data/user-data cloud-init-$VM_NAME/
cp cloud-init-data/network-config cloud-init-$VM_NAME/

# ISO erstellen
hdiutil makehybrid -iso -joliet -default-volume-name "cidata" -o $VM_NAME-seed.iso cloud-init-$VM_NAME/

echo "VM $VM_NAME ist bereit!"
```

## Fazit

Mit dieser Methode hat man in wenigen Minuten eine vollständig konfigurierte VM am Start. Das Setup ist einmal Arbeit, aber dann kann man beliebig viele VMs aus dem Template erstellen.

Perfekt für:

- Schnelle Entwicklungsumgebungen
- CI/CD-Testing
- Kubernetes-Experimente
- Sicherheitstests
- Schulungen

Die ganze Konfiguration lässt sich versionieren und an verschiedene Projekte anpassen. Einmal erstellt, hat man immer eine saubere Testumgebung parat – ohne stundenlanges Setup.
