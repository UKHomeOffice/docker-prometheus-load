# Prometheus Metrics Load Application

Simple Go application publishing fake Prometheus metrics for load testing.

The number of metrics as well as other parameters can be configured from the command line:

```shell script
Usage of /app/prometheus-load:
  -listen-address string
    	The address to listen on for HTTP requests. (default ":8080")
  -metric-count int
    	The number of metrics to publish (default 10)
  -metric-name string
    	The base name of the fake metrics to be published (default "prometheus_load_fake")
  -sleep-time int
    	The number of seconds to sleep before incrementing the fake metrics (default 2)```
