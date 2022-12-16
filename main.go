package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
	println("Ok")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprintln(w, "", header)
	var acc []string = header["Accept"]
	for _, n := range acc {
		println(n)
	}
	acc = header["Accept-Language"]
	for _, n := range acc {
		println(n)
	}

	//VERSION := os.Getenv("NUMBER_OF_PROCESSORS")
	VERSION := os.Getenv("VERSION")
	println(VERSION)
	fmt.Fprintln(w, "", VERSION)

	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip == "" {
		ip = strings.TrimSpace(r.Header.Get("X-Real_Ip"))
		if ip == "" {
			ipR, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
			if err == nil {
				ip = ipR
			}
		}
	}

	io.WriteString(w, ip)
	io.WriteString(w, string(200))

	w.WriteHeader(200)
}
