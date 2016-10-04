package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config is the base configuraiton object
type Config struct {
	Service  ServiceConfig  `yaml:"service"`
	InfluxDB InfluxDBConfig `yaml:"influxdb"`
}

// ServiceConfig holds configuration specific to running the web service
type ServiceConfig struct {
	ListenAddr string `yaml:"listenaddr"`
	ListenPort string `yaml:"listenport"`
	Cert       string `yaml:"cert,omitempty"`
	Key        string `yaml:"key,omitempty"`
}

// InfluxDBConfig holds the configuration for the InfluxDB storage backend.
type InfluxDBConfig struct {
	Scheme   string `yaml:"scheme"`
	Host     string `yaml:"host"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database"`
	Port     int    `yaml:"port,omitempty"`
}

// NewConfig creates an new config object from the given filename.
func NewConfig(filename string) (Config, error) {
	cfgFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	c := Config{}
	err = yaml.Unmarshal(cfgFile, &c)
	if err != nil {
		return Config{}, err
	}

	if c.InfluxDB.Host == "" {
		log.Fatalln("Config file error: InfluxDB host must be provided.")
	}

	if c.InfluxDB.Scheme == "" {
		log.Fatalln("Config file error: InfluxDB protocol scheme must be provided.")
	}
	if c.InfluxDB.Database == "" {
		log.Fatalln("Config file error: InfluxDB database name must be provided.")
	}

	return c, nil
}
