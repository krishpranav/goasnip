package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type sliceVal []string

func (s sliceVal) String() string {
	var str string
	for _, i := range s {
		str += fmt.Sprintf("%s\n", i)
	}
	return str
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("[!] Couldn't create file: %s\n", err.Error())
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintf(w, line)
	}
	return w.Flush()
}

func httpRequest(URI string) string {
	response, errGet := http.Get(URI)
	if errGet != nil {
		log.Fatalf("[!] Error Sending request: %s\n", errGet.Error())
	}
	responseText, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Fatalf("[!] Error reading response: %s\n", errRead.Error())
	}
	defer response.Body.Close()
	return string(responseText)
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func IsIPv6CIDR(cidr string) bool {
	ip, _, _ := net.ParseCIDR(cidr)
	return ip != nil && ip.To4() == nil
}

func CIDRToIP(cidr string) []string {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalf("[!] Failed to convert CIDR to IP: %s\n", err.Error())
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incrementIP(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) >= 3 {
		return ips[1 : len(ips)-1]
	} else {
		return ips
	}
}

func getIP(ipdomain string) []net.IP {
	ip, err := net.LookupIP(ipdomain)
	if err != nil {
		log.Fatalf("[!] Error looking up domain: %s\n", err.Error())
	}
	return ip
}

func parseResponse(response string) []string {
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
	arr := r.FindAllString(response, -1)
	return arr
}

func checkAPIRateLimit(response string) {
	if strings.Contains(response, "API") {
		fmt.Println("[!] The Hacker Target API limit was reached, exiting.....")
		os.Exit(0)
	}
}
