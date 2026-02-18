---
title: 'DSL Graph mit Python'
date: 2018-01-02 16:14:52
update: 2018-01-02 16:14:52
author: ruediger
cover: "/images/cat/internet.webp"
tags:
    - Internet
    - Linux
    - MacOS
    - Programming
    - Python
preview: "Da heute ja [hier](https://blog.pretzlaff.info/2018/01/02/dsl-traffic-2017/) Graphen zum Traffic im Jahr 2016 veröffentlicht wurde hier auch gleich eines der Scripts."
categories: 
    - Technik
toc: false
hide: false
draft: false
type: post
---

Da heute ja [hier](https://blog.pretzlaff.info/2018/01/02/dsl-traffic-2017/) Graphen zum Traffic im Jahr 2016 veröffentlicht wurde hier auch gleich eines der Scripts.
Datenquelle: fb.csv

<!--more-->

```
Monat;Gesendet;Empfangen;Gesamt 
Jan;4124;16815;20939 
Feb;1078;403;1481 
Mar;1199;446;1645 
Apr;2464;36476;38940 
Mai;92979;615268;708247 
Jun;138402;664743;803145 
Jul;116406;507155;623561 
Aug;35654;471810;507464 
Sep;31362;428400;459762 
Okt;24072;549927;573999 
Nov;44095;914362;958457 
Dez;57889;1141699;1199588 
Gesamt:;549724;5347504;5897228
```

Benötigte python Pakete:

*   pygal
*   cairosvg

`pip install pygal pip install cairosvg` bzw. `apt-get install -y python-cairosvg python-pygal` Script welches die Daten ausliest und die Grafik erstellt:

```
#!/bin/bash
OUT=`for i in $(tail -n 13 fb.csv | head -n 12 | awk -F";" {'print $2'} ); do echo -n "$i "; done | sed -e 's/\ $//' | sed -e 's/\ /, /g'`
IN=`for i in $(tail -n 13 fb.csv | head -n 12 | awk -F";" {'print $3'} ); do echo -n "$i "; done | sed -e 's/\ $//' | sed -e 's/\ /, /g'`
SUM=`for i in $(tail -n 13 fb.csv | head -n 12 | awk -F";" {'print $4'} ); do echo -n "$i "; done | sed -e 's/\ $//' | sed -e 's/\ /, /g'`

echo " import pygal bar_chart = pygal.Bar() bar_chart.x_labels = 'Jan', 'Feb', 'Mar', 'Apr', 'Mai','Jun','Jul','Aug','Sep','Okt','Nov','Dez' bar_chart.add('Eingehend', [$IN ]) bar_chart.add('Ausgehend', [$OUT ]) bar_chart.add('Gesamt', [$SUM ]) bar_chart.render_to_file('output.svg') " > generate.py
```
Ausführen:
```
/usr/bin/python generate.py
```

oder ausführbar machen und ausführen:

```
chmod +x generate.sh
./generate.sh
```

Ergebnis: [![PyGal output](/images/posts/output.svg)](-/imgages/posts/output.svg)
