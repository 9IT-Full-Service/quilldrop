---
title: 'PHP in AWS Lambda Function ausführen'
date: 2017-09-19 12:30:14
update: 2017-09-19 12:30:14
author: ruediger
cover: "/images/cat/internet.webp"
tags: [Application, aws, Azure, Cloud, Google Cloud Computing, Internet, Internet, nodeJS, perl, php, programming]
preview: "Bei der Transformation in der Cloud möchte man manchmal schnell kleine Scripts in die Cloud bringen ohne erst Server oder Container in der Cloud zu installieren. In vielen Fällen reichen Funktionen wie AWS-Lambda-Functions, Azure Functions oder Google Cloud Functions."
categories: 
    - Internet
toc: false
hide: false
type: post
---

Bei der Transformation in der Cloud möchte man manchmal schnell kleine Scripts in die Cloud bringen ohne erst Server oder Container in der Cloud zu installieren. In vielen Fällen reichen Funktionen wie AWS-Lambda-Functions, Azure Functions oder Google Cloud Functions. AWS Lambda bietet dafür in Lambda folgende (Script)-Sprachen an:
<!--more-->

*   C#
*   Java
*   python
*   nodejs

Gerade in älteren Umgebungen werden oft Perl und PHP für kleinere Webscripte, die Applikationen zum Austausch von Informationen benutzen eingesetzt. Klar, diese Scripte sollte man teilweise überdenken und erneuern. Das kostet zum einen Geld, jemand muss es machen und manche sollen auch nur noch bei dem Sprung in die Cloud helfen und dann eh das zeitliche segnen. Daher wäre es nicht schlecht diese Tools schnell mit in die Cloud zu nehmen um andere Software so in die Cloud zu bekommen und die Tools nicht als Spassbremse zu haben.

nodejs und statisch kompliliertes php
-------------------------------------

Da Lambda nodejs anbietet kann man es auch dazu nötigen PHP auszuführen. Ja, würde ich es jetzt auch lesen oder es von jemanden hören, ich würde mich jetzt genau so schütteln. Aber es soll ja nur den Sprung in die Cloud ermöglichen und zeigen wie man solche Probleme umgehen kann.

PHP Binary erstellen
--------------------

Um das PHP statisch als ein Binary vorliegen zu haben bietet sich Docker an. build\_php\_7.sh
```
#!/bin/sh
PHP_VERSION_GIT_BRANCH=PHP-7.1.1
echo "Build PHP Binary from current branch '$PHP_VERSION_GIT_BRANCH' on https://github.com/php/php-src"
docker build --build-arg PHP_VERSION=$PHP_VERSION_GIT_BRANCH -t php-build -f Dockerfile.BuildPHP .
container=$(docker create php-build)
docker -D cp $container:/root/php7/usr/bin/php ./php
docker rm $container
```
Das stellt einen Container der PHP 7.7.1 zusammensetzt und anschliessend das fertige Binary aus den Docker Container kopiert. Wer eine andere Version benötigt kann die Version mit PHP\_VERSION\_GIT\_BRANCH setzen.

PHP Script
----------

```
<?php
echo "Hello world!";
var_dump($argv);
?>
```

index.js spawn für php
----------------------

Damit der PHP Interpreter und ein Script aufgerufen werden benötigt man nur noch etwas JavaScript:
```
'use strict';
var child_process = require('child_process');
exports.handler = function(event, context) {
  var strToReturn = '';
  var proc = child_process.spawn('./php', [ "index.php", JSON.stringify(event), { stdio: 'inherit' } ]);
  proc.stdout.on('data', function (data) {
    var dataStr = data.toString()
    console.log('stdout: ' + dataStr);
    strToReturn += dataStr
  });

  proc.on('close', function(code) {
    if(code !== 0) {
      return context.done(new Error("Process exited with non-zero status code"));
    }
    context.succeed(strToReturn);
  });
}
```

Alles einpacken und verschiffen
-------------------------------

Jetzt nur noch ein Zip-File `aws-lambda-php-example.zip` erstellen mit den Dateien:

*   index.js
*   index.php
*   php zip aws-lambda-php-example.zip index.js index.php php


In AWS eine neue Lambda Function erstellen mit folgenden Parametern:

*   Type: nodeJS
*   RAM: 128mb
*   Timeout: 3 seconds

Den JavaScript Code nicht in den Online Editor kopieren, sondern Upload auswählen und das komplette ZIP-File hochladen und an der Lambda Function anhängen. Alternativ im S3 ablegen und aus dem Bucket heraus laden lassen Benötigt man die Schnittstelle per HTTP von extern und/oder anderen Instanzen kann man auch noch das API-Gateway von Amazon hinzuziehen.

Spassbremse umgangen
--------------------

Das Script stört nicht mehr die weitere Transformation und alle anderen Softwarebrocken können ihren Weg in die Cloud beschreiten. Die Teams, die solche Softwarestückchen einmal in schön abliefern müssen, können dies dann noch später nachholen. Haben aber auch erst einmal Zeit für die grösseren Cloud-Projekte. Wer aus "Gründen" nicht sofort eine komplette Software in die Cloud bringen kann sollte sich das API-Gateway einmal genauer angucken. Path und Method lassen sich damit sehr gut trennen. Einzelne Aufrufe lassen sich so auf die alte Software oder die neue Software leiten. Natürlich nur, so lange so etwas mit einer Software möglich ist und intern keine Abhängigkeiten bestehen. Ist es möglich steht einer Migration einzelner Funktionen nichts im Wege. Rollback inklusive.
