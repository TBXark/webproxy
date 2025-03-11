package main

import (
	"flag"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	target := flag.String("target", "", "target url")
	socks5 := flag.String("socks5", "", "socks5 address")
	listen := flag.String("listen", ":8080", "listen address")
	flag.Parse()
	if *target == "" || *socks5 == "" {
		flag.Usage()
		return
	}
	err := run(*target, *socks5, *listen)
	if err != nil {
		log.Fatal(err)
	}
}

func run(target string, socks5 string, listen string) error {
	targetURL, err := url.Parse(target)
	if err != nil {
		return err
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	dialer, err := proxy.SOCKS5("tcp", socks5, nil, proxy.Direct)
	if err != nil {
		return err
	}

	reverseProxy.Transport = &http.Transport{
		Dial: dialer.Dial,
	}

	originalDirector := reverseProxy.Director
	reverseProxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Host = targetURL.Host
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	})

	log.Printf("Web proxy server started at %s, target: %s, socks5: %s", listen, target, socks5)
	err = http.ListenAndServe(listen, nil)
	if err != nil {
		return err
	}
	return nil
}
