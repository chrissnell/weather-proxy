#!/bin/bash -e

if [ -z "${WEATHER_PROXY_CONFIG}" ]; then
  echo The env var WEATHER_PROXY_CONFIG needs to be defined. 
  echo This variable points weather-proxy towards its config file.
  echo This image accepts a volume, /config, that you can
  echo use for passing in a config file externally.
  echo Exiting...
  exit 1
fi

# Use gosu to drop privileges
exec gosu nobody /weather-proxy -config=$WEATHER_PROXY_CONFIG
