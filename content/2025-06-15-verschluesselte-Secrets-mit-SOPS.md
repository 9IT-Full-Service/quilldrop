---
title: 'SOPS-verschlüsselte Kubernetes Secrets'
date: 2025-06-15 13:00:00
author: ruediger
cover: "/images/posts/2025/06/sops-secret-encryption.webp"
featureImage: "/images/posts/2025/06/sops-secret-encryption.webp"
tags: [Kubernetes, Secrets, encrypt, SOPS, GitOps]
categories: 
    - Kubernetes
preview: "SOPS (Secrets OPerationS) ermöglicht die sichere Speicherung von Kubernetes Secrets im Git Repository durch selektive Verschlüsselung der sensitiven Datenblöcke. Mit dem Parameter --encrypted-regex '^(data|stringData)$' werden nur die Secret-Inhalte verschlüsselt, während die Kubernetes-Metadaten lesbar bleiben. Dies bietet eine GitOps-kompatible Alternative zu externen Secret Vaults und ermöglicht vollständige Versionskontrolle der Secret-Konfigurationen. Die AGE-Verschlüsselung stellt dabei sowohl Datenschutz als auch Integrität der gespeicherten Secrets sicher."
series: ["FluxCD"] 
draft: false
top: false
type: post
hide: false
toc: true
---

![SOPS Secret Encryption](/images/posts/2025/06/sops-secret-encryption.webp)

# Überblick

Bei der Verwaltung von Kubernetes Secrets ohne Verwendung eines Secret Vaults (wie HashiCorp Vault oder Azure Key Vault) ist die Verschlüsselung der Secrets vor der Speicherung im Git Repository essentiell. SOPS (Secrets OPerationS) bietet hierfür eine robuste Lösung.

# Verschlüsselungsansatz

Da Kubernetes Secrets ihre Metadaten (apiVersion, kind, metadata, type) für die korrekte Funktionalität benötigen, erfolgt die Verschlüsselung selektiv nur auf die sensitiven Datenblöcke:

   * data: Base64-kodierte Secret-Werte
   * stringData: Klartext Secret-Werte (werden automatisch zu data konvertiert)

Die Kubernetes-Metadaten bleiben unverschlüsselt, damit das YAML weiterhin als gültiges Kubernetes-Manifest erkannt wird.

# SOPS-Verschlüsselungskommando

```bash
sops --encrypt --in-place --encrypted-regex '^(data|stringData)$' \
    --age <SOPS-AGE-KEY> path/to/docker-secret.yaml
```

## Parameter-Erklärung:

   * --encrypt: Verschlüsselungsmodus
   * --in-place: Überschreibt die ursprüngliche Datei
   * --encrypted-regex '^(data|stringData)$': Regulärer Ausdruck zur Bestimmung der zu verschlüsselnden Felder
   * --age <SOPS-AGE-KEY>: Verwendung des AGE-Verschlüsselungsschlüssels

# Struktur eines verschlüsselten Docker Registry Secrets

```yaml
apiVersion: v1
data:
    .dockerconfigjson: ENC[AES256_GCM,data:EXAMPLE_ENCRYPTED_DATA_PLACEHOLDER_DO_NOT_USE_IN_PRODUCTION,iv:EXAMPLE_IV_PLACEHOLDER,tag:EXAMPLE_TAG_PLACEHOLDER,type:str]
kind: Secret
metadata:
    name: my-docker-registry-secret
    namespace: my-namespace
type: kubernetes.io/dockerconfigjson
sops:
    age:
        - recipient: age1example123456789abcdefghijklmnopqrstuvwxyz0123456789abcdef
          enc: |
            -----BEGIN AGE ENCRYPTED FILE-----
            EXAMPLE_AGE_ENCRYPTED_CONTENT_PLACEHOLDER
            DO_NOT_USE_THIS_IN_PRODUCTION
            THIS_IS_ONLY_AN_EXAMPLE
            -----END AGE ENCRYPTED FILE-----
    lastmodified: "2025-01-15T10:30:00Z"
    mac: ENC[AES256_GCM,data:EXAMPLE_MAC_PLACEHOLDER_FOR_DEMONSTRATION_PURPOSES_ONLY,iv:EXAMPLE_MAC_IV,tag:EXAMPLE_MAC_TAG,type:str]
    encrypted_regex: ^(data|stringData)$
    version: 3.10.2
```

## SOPS-Metadaten Erklärung

Nach der Verschlüsselung fügt SOPS automatisch einen sops-Block hinzu:
AGE-Verschlüsselung

   * recipient: Der öffentliche AGE-Schlüssel für die Verschlüsselung
   * enc: Der verschlüsselte Master-Schlüssel im AGE-Format

## Integritätssicherung

   * mac: Message Authentication Code zur Verifikation der Datenintegrität
   * lastmodified: Zeitstempel der letzten Änderung
   * encrypted_regex: Bestätigung des verwendeten Verschlüsselungsmusters
   * version: SOPS-Version

# Vorteile dieses Ansatzes

   1. GitOps-Kompatibilität: Secrets können sicher im Git Repository gespeichert werden
   2. Kubernetes-Kompatibilität: Metadaten bleiben lesbar für Kubernetes
   3. Selektive Verschlüsselung: Nur sensitive Daten werden verschlüsselt
   4. Audit-Trail: Git-History für Secret-Änderungen verfügbar
   5. Schlüsselrotation: AGE-Schlüssel können rotiert werden

# Entschlüsselung und Anwendung

## Entschlüsselung zur Ansicht

```bash
sops --decrypt path/to/docker-secret.yaml
```

## Direkte Anwendung auf Kubernetes

```bash
sops --decrypt path/to/docker-secret.yaml | kubectl apply -f -
```

# Sicherheitshinweise

   * AGE-Private-Keys müssen sicher außerhalb des Repositories gespeichert werden
   * Regelmäßige Rotation der Verschlüsselungsschlüssel
   * Verwendung unterschiedlicher Schlüssel für verschiedene Umgebungen (Dev/Staging/Prod)
   * Backup-Strategie für Verschlüsselungsschlüssel implementieren

