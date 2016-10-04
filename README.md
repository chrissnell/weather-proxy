# weather-proxy
A simple service to safely proxy live weather readings from InfluxDB.  This proxy is heavily tied to my [gopherwx](github.com/chrissnell/gopherwx) project, which gathers live weather readings from a Davis Instruments Vantage Pro2 weather station.  This proxy allows you to expose live weather data stored in InfluxDB without having to expose InfluxDB directly to the wild, wild Internet.

## Use
1. Configure weather-proxy by editing `config.yaml` and pointing it at your InfluxDB instance.  Provide weather-proxy with a TLS cert and key so that it can answer HTTPS requests.  If you don't provide a cert+key, it will default to speaking HTTP.

2. Expose weather-proxy to the public Internet if you wish.
