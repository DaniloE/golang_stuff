Just a litle programm to create statistics from a logfile. 
It was a little competition exercise, so notheing special ;)

If you want to use it, load the current GeoIP-CSV-database from https://dev.maxmind.com/geoip/legacy/geolite/.
The log should have the ip in the second column, otherwise it won't work.


Use it like:

`./sort_geoip -file=httpsd_21_01_2015.log --geoip GeoIPCountryWhois.csv`

example output from a apache logfile with ~500k lines on a 2,2Ghz virtual server(not really performant):


```
IPs in logfile ausgewertet:  639.431732ms

GeoIP-csv eingelesen:  480.477676ms
Parsing der IPs und Laender beendet:  324.551261ms
Top Ten der Zugriffe pro Land:
--------------
CN      296676
DE      128862
US      18501
AT      7442
FR      2954
CH      2279
HK      1869
JP      1799
CA      1644
TW      1410

Top Ten der zugreifenden IPs:
--------------
87.170.190.7    12058
::1             8840
113.200.204.235 2134
46.5.184.23     1807
213.174.47.79   1367
116.207.30.108  1331
87.154.172.131  1311
79.238.179.156  1307
141.8.189.125   1057
162.243.92.45   1002
Gesamtlaufzeit:  1.488752685s
```


