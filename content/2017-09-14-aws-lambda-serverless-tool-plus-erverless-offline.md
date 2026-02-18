---
title: 'AWS Lambda serverless Framework und serverless-offline'
date: 2017-09-14 09:53:46
update: 2017-09-14 09:53:46
author: ruediger
cover: "/images/cat/internet.webp"
tags:
    - AWS
    - Cloud
    - Internet
    - Javascript
    - NodeJS
    - npm
    - Programming
    - Serverless
preview: "nodejs, npm, serverless und serverless-offline installieren"
categories: 
    - Internet
toc: false
hide: false
type: post
---


nodejs, npm, serverless und serverless-offline installieren
===========================================================

*   [serverless](https://github.com/serverless/serverless)
*   [serverless-offline](https://github.com/dherault/serverless-offline)
*   [NodeJS](https://nodejs.org/en/)

Installation nodejs
-------------------
<!--more-->

```
$ sudo apt-get install curl python-software-properties
$ curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash -
$ sudo apt-get install nodejs
$ node -v

v8.2.1

$ npm -v

5.3.0
```

Installation serverless
-----------------------

```
npm install -g serverless
```

serverless service erstellen
============================

Project erstellen
-----------------

```
# Ein neues Serverless Service/Project erstellen
serverless create --template aws-nodejs --path my-service
# In das neue Verzeichnis wechseln
cd my-service
```

AWS Access-Key und Secret
-------------------------

```
export AWS_ACCESS_KEY_ID=<your-key-here>
export AWS_SECRET_ACCESS_KEY=<your-secret-key-here>
serverless deploy
```
oder
```
serverless config credentials --provider aws --key AKIAIOSFODNN7EXAMPLE --secret wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

Deploy Service
--------------

```
serverless deploy -v
```

Die Logs der Function abrufen:
------------------------------

```
serverless invoke -f hello -l
```

Service in AWS entfernen
------------------------

```
serverless remove
```

mysql RSD DB mit NodeJS in Lambda benutzen
==========================================

Als erstes im Projectverzeichnis (my-project) das mysql plugin installieren
```
npm install mysql
```
Im Code mysql hinzufügen und den restlichen mysql java-script Code:
```
var mysql = require('mysql');
module.exports.view = (event, context, callback) => {
  const ip = event.requestContext.identity.sourceIp;
  var connection = mysql.createConnection({
    host: 'YourRdsDB.xxxxxxx.eu-central-1.rds.amazonaws.com',
    user: 'DBUser',
    password: 'DBPass',
    database: 'DBname'
  });
  connection.connect();
  // ....
}
```

Serverless offline ohne AWS benutzen
====================================

Serverless-offline installieren
-------------------------------

```
npm install serverless-offline --save-dev
```

### Hinter einem Cooperate-Proxy ggf. auch noch mit Auth?

Einfach npm mit config set die Parameter für https-proxy und proxy setzen.
```
npm config set proxy http://"username:gehe\!m"@proxy.example.com:3128
npm config set https-proxy http://"username:gehe\!m"@proxy.example.com:3128
```
Sonderzeichen wie das '!' hier im Beispiel müssen escaped werden. Und Benutzer und Password müssen komplett in " gesetzt werden. Das Verzeichnis "node\_modules" sollte jetzt ungefähr so aussehen und unter anderem serverless-offline auflisten:

```
ls node_modules/
accept            babel-helpers   balanced-match   content               esutils            home-or-tmp  json5             mimos          os-homedir           serverless-offline  supports-color
ammo              babel-messages  boom             convert-source-map    globals            invariant    jsonpath-plus     minimatch      os-tmpdir            shot                to-fast-properties
ansi-regex        babel-register  brace-expansion  core-js               h2o2               iron         js-string-escape  minimist       path-is-absolute     slash               topo
ansi-styles       babel-runtime   call             cryptiles             hapi               isemail      js-tokens         mkdirp         peekaboo             source-map          trim-right
b64               babel-template  catbox           crypto                hapi-cors-headers  is-finite    kilt              moment         pez                  source-map-support  velocityjs
babel-code-frame  babel-traverse  catbox-memory    debug                 has-ansi           items        lodash            ms             private              statehood           vise
babel-core        babel-types     chalk            detect-indent         heavy              joi          loose-envify      nigel          regenerator-runtime  strip-ansi          wreck
babel-generator   babylon         concat-map       escape-string-regexp  hoek               jsesc        mime-db           number-is-nan  repeating            subtext
```

PlugIn in der serverless.yml hinzufügen
---------------------------------------

```
plugins:
  - serverless-offline
```

Überprüfen ob das PlugIn verfügbar ist
--------------------------------------

```
serverless
```
Die Ausgabe sollte unter commands jetzt zusätzlich "offline" und "offline-start" auflisten
```
...
logs .......................... Output the logs of a deployed function
metrics ....................... Show metrics for a specific function
offline ....................... Simulates API Gateway to call your lambda functions offline.
offline start ................. Simulates API Gateway to call your lambda functions offline using backward compatible initialization.
package ....................... Packages a Serverless service
remove ........................ Remove Serverless service and all resources
...
```

In der letzten Zeile werden alle verfügbaren PlugIns aufgelistet und "Offline" sollte dort auch aufgelistet werden.
```
Plugins
AwsCommon, AwsCompileAlexaSkillEvents, AwsCompileApigEvents, AwsCompileCloudWatchEventEvents, AwsCompileCloudWatchLogEvents, AwsCompileCognitoUserPoolEvents, AwsCompileFunctions, AwsCompileIoTEvents, AwsCompileS3Events, AwsCompileSNSEvents, AwsCompileScheduledEvents, AwsCompileStreamEvents, AwsConfigCredentials, AwsDeploy, AwsDeployFunction, AwsDeployList, AwsInfo, AwsInvoke, AwsInvokeLocal, AwsLogs, AwsMetrics, AwsPackage, AwsProvider, AwsRemove, AwsRollback, AwsRollbackFunction, Config, Create, Deploy, Emit, Info, Install, Invoke, Login, Logout, Logs, Metrics, Offline, Package, Platform, Remove, Rollback, Run, SlStats
```

Projekt offline starten
-----------------------

`serverless offline start` or `sls offline start`.
```
serverless offline start
Serverless: Starting Offline: dev/us-east-1.

Serverless: Routes for hello:
Serverless: (none)

Serverless: Offline listening on http://localhost:3000
```

Service testen
--------------

Dabei an den Proxy denken und den Parameter `--noproxy` setzen:
```
curl --noproxy "127.0.0.1, localhost" http://localhost:3000
```

Parameter von serverless-offline
--------------------------------

```
serverless offline --help
--prefix                -p  Adds a prefix to every path, to send your requests to http://localhost:3000/[prefix]/[your_path] instead. E.g. -p dev
--location              -l  The root location of the handlers' files. Defaults to the current directory
--host                  -o  Host name to listen on. Default: localhost
--port                  -P  Port to listen on. Default: 3000
--stage                 -s  The stage used to populate your templates. Default: the first stage found in your project.
--region                -r  The region used to populate your templates. Default: the first region for the first stage found.
--noTimeout             -t  Disables the timeout feature.
--noEnvironment             Turns off loading of your environment variables from serverless.yml. Allows the usage of tools such as PM2 or docker-compose.
--resourceRoutes            Turns on loading of your HTTP proxy settings from serverless.yml.
--dontPrintOutput           Turns off logging of your lambda outputs in the terminal.
--httpsProtocol         -H  To enable HTTPS, specify directory (relative to your cwd, typically your project dir) for both cert.pem and key.pem files.
--skipCacheInvalidation -c  Tells the plugin to skip require cache invalidation. A script reloading tool like Nodemon might then be needed.
--corsAllowOrigin           Used as default Access-Control-Allow-Origin header value for responses. Delimit multiple values with commas. Default: '*'
--corsAllowHeaders          Used as default Access-Control-Allow-Headers header value for responses. Delimit multiple values with commas. Default: 'accept,content-type,x-api-key'
--corsDisallowCredentials   When provided, the default Access-Control-Allow-Credentials header value will be passed as 'false'. Default: true
--exec "<script>"           When provided, a shell script is executed when the server starts up, and the server will shut domn after handling this command.
```

Parameter in der serverless.yml setzen
--------------------------------------

Beispiel:
```
custom:
  serverless-offline:
    httpsProtocol: "dev-certs"
    port: 4000
    prefix: "dev"
    stage: "dev"
```
