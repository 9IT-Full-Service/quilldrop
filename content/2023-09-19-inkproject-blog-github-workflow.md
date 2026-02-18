---
title: InkProject Blog Github Workflow
date: 2023-09-19 21:00:00
author: ruediger
cover: "/images/cat/technik.webp"
tags: [InkProjekt, Blog, GitHub, Action, Workflow, Trigger]
categories: 
    - Internet
preview: "Dieses Blog wird nicht mehr durch Hugo generiert. Der Hugo Static Content Generator war jetzt lange im Einsatz, doch seit ein paar Wochen habe ich mir alternativen angeguckt. Dabei bin ich mehrere durchgegangen und bin schließlich bei InkProjekt geblieben. " 
draft: false
top: false
type: post
hide: false
toc: false
---

[English Version](/posts/2023-09-19-inkproject-blog-github-workflow-en.html)

## Bey Bey Hugo, welcome Ink

Dieses Blog wird nicht mehr durch Hugo generiert. Der Hugo Static Content Generator war jetzt lange im Einsatz,
doch seit ein paar Wochen habe ich mir alternativen angeguckt. Dabei bin ich mehrere durchgegangen und bin schließlich bei InkProjekt geblieben. 

Nach den ersten Tests wurden dann auch nach und nach immer mehr Posts umgezogen. Dabei wurden dann auch gleich die ersten Schritte in Richtung Docker und Kubernetes gemacht. Also für Ink und dieses Blog. Docker und Kubernetes mache ich ja jetzt schom ein paar Jahre. 

## Schreiben, publish, build und deploy

Die Reihenfolge ist klar, wie immer halt. Ein neuer Post wird geschrieben, oder wie jetzt in Ink kopiert, anschliessend `ìnk publish` ausgeführt. Alle Files aus dem `public` nach `html` kopiert, so das ein Dockerfile für Nginx das html Verzeichnis mit dem neuen Content bespielen kann. Docker Push in die Docker Registry, gefolgt von einem `kubectl set image...`. Die neue Version war online. 

Das waren erst einmal noch manuelle Schritt. Ein paar Shell Scripte. 

## GitHub Actions

Nachdem das alles wunderbar lief war es an der Zeit alles automatisiert zu erledigen. Post oder Page anlegen, git add,commit,push und GitHub Action sollten sich um das zusammen bauen der Page, das Builden von Docker Images für alle Plattformen und den HelmChart kümmern. 

Das war alles erst einmal in einem Git Repo, da kam das als ersten der Helm Chart wieder raus in ein eigens. 
Der Workflow kümmerte sich nach dem Push dann um folgende Schritt: 

    1 repo Checkout
    2 Semantic Release Version hoch zählen
    3 Version in GitHub Env speichern
    4 Page im ink publish erstellen
    5 Seiten in Docker Images für alle Plattformen packen
    6 Docker Image zu einem MultiPlatform Image zusammenfügen
    7 Docker Image in die Registrie pushen
    8 Helm Chart Repo auschecken
    9 Mit der Semantiv Version Helm Chart erstellen
    10 Helm Chart ins Helm Chart Museum pushen
    11 Helm Chart in den Kubernetes Cluster deployen

## Da alles perfekt lieft, alles auseinander reißen

Da alles, bis auf der Helm Chart, in einem Repository gewesen ist konnte man Teile davon schlecht für andere Seiten oder Blogs wiederverwenden. Also wurde als erstes der Content in ein eigenes GitHub Repo gesteckt. 

Beim InkProject befindet sich in Source nicht nur die Markdown Files mit dem Content, auch die Images sind dort drin. Die wurden da auch direkt raus und auch in ein eigenes Repo kopiert. so das man wirklich nur Markdown Dateien für den Blog Content in einem Repo hat. 

Themes, die haben noch ein Helm Repository bekommen. Genau so wie die grundsätzlichen Teile von InkProject. Die Binary Files, ja Files, da ich direkt arm, amd usw dazu gepackt habe, falls die GitHub RUnner mal auf anderen Plattformen laufen, genau so wie die Config. 

## Der GitHub Workflow konnte dann angepasst werden. 

Der Workflow wird jetzt im Repo `ink-blog` gestartet. Dieser zieht sich nach einander die Repositories: 
    
    * ink-content
    * ink-images
    * ink-theme
    * ink-html

Die werden in den Ordnern `source`, `source/images`, `theme` und `html` abgelegt und somit ist die Ordner-Struktur von InkProject wie sie ein muss. 

Im nächsten Schritt wird der Blog erstellt: `ink publish` und dann die Dateien nach `html` kopiert. 
das Verzeichnis `html` ist ein GitHub Repo, da wird dann rein gewechselt und ein git add,commit,push gemacht. Damit habe ich dann schon einmal den neuen Content in einem eigenen Repo. 

Der Rest wird dann schon mal wieder im Workflow weggeschmissen und im nächsten Schritt werden jetzt die Docker Images für mehere Plattformen erstellt. Dazu wird dann das Repo `ink-html` benötigt und daher in den Steps ausgecheckt. 

Alle Docker Images werden wieder zu einem MultiPlatform Image zusammengefügt und in die Docker Registry geladen. 

Bei den ganzen Schritten natürlich wieder Semantoc Release, also Docker Images und Helm Chart automatisch mit der richtigen Version versehen. 

Der Helm Chart wird zum Schluß noch ausgecheckt und mit der Semantic Version erstellt, das Image über die `values.yaml` gesetzt. HelmChart ins Chart Museum und der Helm Chart in den Kubernetes Cluster installieren. 

## Workflow die 3. 

Der Workflow ist gut und man könnte ihn so lassen. Da ich aber den Content in einem eigenen Repo pflege, müsste man je immmer den Content erstellen oder ändern, dann pushen. Dann das gleiche beim `ink-blog` Repo auch noch machen, damit die Pipeline getriggert wird. 

Nö, das geht auch anders. 

## Content mit eigener Action 

Die Action in dem Content Repo macht nicht viel. Sie baut nicht mal irgend etwas zusammen. 
Sie triggert einfach nur die andere Pipeline im Reop `ink-blog`, welches dann alles weitere macht. 

Content triggert ink-blog, ink-blog checked `ink-content` und weitere Repos aus, baut alles zusammen, erstellt Images und Helm Chart und deployed alles in Kubernetes. 

So einfach ist das. 

Der Vorteil ist: Ich kann Content auch mal eben vom Handy aus erstellen und anpassen. Ich kann das ganze auch noch erweitern für andere Blogs und Homepages. Dann wird nur noch der Content für die einzelnen Seiten hinterlegt, der Build mit anderen Daten kommt immer aus den gleichen Repos. Das ist noch nicht eingerichtet, aber die Grundlage dafür ist schon mal angelegt. 



