package main

import (
	"log"
	"net/http"
)

func liveWX(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT LAST(OutTemp), LAST(OutHumidity), LAST(Barometer), LAST(WindSpeed), LAST(WindDir) FROM wx_reading WHERE time > now() - 30m ORDER BY time DESC LIMIT 1")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
	w.Write(b)
}

func lastDayRain(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, " SELECT LAST(YearRain) - FIRST(YearRain) FROM wx_reading WHERE time > now() - 1d")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}

func dayWX(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT mean(OutTemp), mean(OutHumidity), mean(Barometer), max(WindSpeed), last(WindDir), difference(last(YearRain)) FROM wx_reading WHERE time > now() - 1d GROUP BY time(5m)")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}

func dayRain(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT difference(last(YearRain)) FROM wx_reading WHERE time > now() - 1d GROUP BY time(30m)")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}

func twoDaysWX(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT mean(OutTemp), mean(OutHumidity), mean(Barometer), max(WindSpeed), last(WindDir), difference(last(YearRain)) FROM wx_reading WHERE time > now() - 2d GROUP  BY time(5m)")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}

func twoDaysRain(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT difference(last(YearRain)) FROM wx_reading WHERE time > now() - 2d GROUP BY time(60m)")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}

func weekWX(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT mean(OutTemp), mean(OutHumidity), mean(Barometer), max(WindSpeed), last(WindDir), difference(last(YearRain)) FROM wx_reading WHERE time > now() - 7d GROUP BY time(30m)")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}

func monthWX(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT mean(OutTemp), mean(OutHumidity), mean(Barometer), max(WindSpeed), last(WindDir), difference(last(YearRain)) FROM wx_reading WHERE time > now() - 30d GROUP BY time(2h)")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}

func yearWX(w http.ResponseWriter, req *http.Request, c Config) {
	b := dbQuery(w, req, c, "SELECT mean(OutTemp), mean(OutHumidity), mean(Barometer), max(WindSpeed), last(WindDir), difference(last(YearRain)) FROM wx_reading WHERE time > now() - 365d GROUP BY time(24h)")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/javascript")
	w.Write(b)
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
}
