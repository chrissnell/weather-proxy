package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	cfgFile := flag.String("config", "config.yaml", "Path to config file (default: ./config.yaml)")
	flag.Parse()

	// Read our server configuration
	filename, _ := filepath.Abs(*cfgFile)
	cfg, err := NewConfig(filename)
	if err != nil {
		log.Fatalln("Error reading config file.  Did you pass the -config flag?  Run with -h for help.\n", err)
	}

	http.HandleFunc("/live", func(w http.ResponseWriter, req *http.Request) { liveWX(w, req, cfg) })
	http.HandleFunc("/last-day-rain", func(w http.ResponseWriter, req *http.Request) { lastDayRain(w, req, cfg) })
	http.HandleFunc("/day", func(w http.ResponseWriter, req *http.Request) { dayWX(w, req, cfg) })
	http.HandleFunc("/twodays", func(w http.ResponseWriter, req *http.Request) { twoDaysWX(w, req, cfg) })
	http.HandleFunc("/week", func(w http.ResponseWriter, req *http.Request) { weekWX(w, req, cfg) })
	http.HandleFunc("/month", func(w http.ResponseWriter, req *http.Request) { monthWX(w, req, cfg) })
	http.HandleFunc("/year", func(w http.ResponseWriter, req *http.Request) { yearWX(w, req, cfg) })

	if cfg.Service.Cert != "" && cfg.Service.Key != "" {
		log.Fatal(http.ListenAndServeTLS(fmt.Sprint(cfg.Service.ListenAddr, ":", cfg.Service.ListenPort), cfg.Service.Cert, cfg.Service.Key, nil))
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprint(cfg.Service.ListenAddr, ":", cfg.Service.ListenPort), nil))
	}
}
