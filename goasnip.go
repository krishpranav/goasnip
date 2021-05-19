package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
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
