package main

import (
	"fmt"
	//"strings"
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"net"
	"os"
	"time"
)

type geonet struct {
	startip string
	endip   string
	land    string
}

func main() {

	start := time.Now()
	funcstart := time.Now()

	var logfile, geoipfile string
	flag.StringVar(&logfile, "file", "", "Logfile to analyse")
	flag.StringVar(&geoipfile, "geoip", "", "geoip-csv")
	flag.Parse()

	lfile, err := os.Open(logfile)
	if err != nil {
		fmt.Println("Konnte die Logdatei nicht öffnen")
		os.Exit(1)
	}
	gfile, err := os.Open(geoipfile)
	if err != nil {
		fmt.Println("Konnte die GeoIP-csv nicht öffnen")
		os.Exit(1)
	}

	ip_table := make(map[string]int)
	country_table := make(map[string]int)
	var geolist []geonet

	ra := []int{}
	var cn int
	scanner := bufio.NewScanner(lfile)
	for scanner.Scan() {
		ra := ra[:0]
		cn = 0
		line := scanner.Text()
		for i, a := range line {
			// if a is a whitespace
			if a == 32 {
				ra = append(ra, i)
				cn++
			}
			if cn == 2 {
				break
			}
		}
		ip := line[ra[0]+1 : ra[1]]
		ip_table[ip] += 1
	}
	fmt.Println("IPs in logfile ausgewertet: ", time.Since(funcstart))
	funcstart = time.Now()

	georeader := csv.NewReader(gfile)
	georeader.FieldsPerRecord = -1
	rawCSVdata, _ := georeader.ReadAll()
	for _, line := range rawCSVdata {
		geoeintrag := geonet{startip: line[0], endip: line[1], land: line[4]}
		geolist = append(geolist, geoeintrag)
	}
	fmt.Println("GeoIP-csv eingelesen: ", time.Since(funcstart))
	funcstart = time.Now()

	for ipkey_s, _ := range ip_table {
		ipkey := net.ParseIP(ipkey_s)
		range_start := 0
		range_end := len(geolist)
		for range_end-range_start > 1 {
			half := range_start + ((range_end - range_start) / 2)
			testv := geolist[half]
			startip := net.ParseIP(testv.startip)
			if bytes.Compare(ipkey, startip) >= 0 {
				range_start = half
			} else if bytes.Compare(ipkey, startip) < 0 {
				range_end = half
			}
		}
		for _, country := range geolist[range_start:range_end] {
			startip := net.ParseIP(country.startip)
			endip := net.ParseIP(country.endip)
			if bytes.Compare(ipkey, startip) >= 0 && bytes.Compare(ipkey, endip) <= 0 {
				if country_table[country.land] == 0 {
					country_table[country.land] = ip_table[ipkey_s]
				} else {
					country_table[country.land] += ip_table[ipkey_s]
				}
				break
			}
		}
	}
	fmt.Println("Parsing der IPs und Laender beendet: ", time.Since(funcstart))
	i := 0
	fmt.Printf("\nTop Ten der Zugriffe pro Land:\n--------------\n")
	for len(country_table) > 1 && i < 10 {
		max_value := 0
		for _, value := range country_table {
			act_value := value
			if act_value > max_value {
				max_value = act_value
			}
		}
		for key, value := range country_table {
			if value == max_value {
				fmt.Printf("%s\t%d\n", key, value)
				delete(country_table, key)
				i++
			}
		}
	}

	c := 0
	fmt.Printf("\nTop Ten der zugreifenden IPs:\n--------------\n")
	for len(ip_table) > 1 && c < 10 {
		max_value := 0
		for _, value := range ip_table {
			act_value := value
			if act_value > max_value {
				max_value = act_value
			}
		}
		for key, value := range ip_table {
			if value == max_value {
				if len(key) < 8 {
					fmt.Printf("%s\t\t%d\n", key, value)
				} else {
					fmt.Printf("%s\t%d\n", key, value)
				}
				delete(ip_table, key)
			}
		}
		c++
	}
	fmt.Println("Gesamtlaufzeit: ", time.Since(start))
}
