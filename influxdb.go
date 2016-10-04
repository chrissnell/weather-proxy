package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func dbQuery(w http.ResponseWriter, req *http.Request, c Config, query string) []byte {
	var u url.URL

	u.Scheme = c.InfluxDB.Scheme
	if c.InfluxDB.Port == 0 {
		u.Host = fmt.Sprint(c.InfluxDB.Host, ":8086")
	} else {
		u.Host = fmt.Sprint(c.InfluxDB.Host, ":", c.InfluxDB.Port)
	}
	u.Path = "/query"

	q := u.Query()

	q.Set("db", c.InfluxDB.Database)

	q.Set("q", query)
	q.Set("epoch", "s")

	u.RawQuery = q.Encode()

	fmt.Println("url:", u.String())

	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
	}
	//fmt.Println(string(body))
	return body
}
