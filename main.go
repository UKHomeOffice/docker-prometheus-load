package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics(metrics []prometheus.Counter, sleepTime int) {
    go func() {
        for {
            for _, metric := range metrics {
                metric.Inc()
            }
            time.Sleep(time.Duration(sleepTime) * time.Second)
        }
    }()
}

func main() {
    addr := flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
    metricName := flag.String("metric-name", "prometheus_load_fake", "The base name of the fake metrics to be published")
    metricsCount := flag.Int("metric-count", 10, "The number of metrics to publish")
    sleepTime := flag.Int("sleep-time", 2, "The number of seconds to sleep before incrementing the fake metrics")

    flag.Parse()

    metrics := make([]prometheus.Counter, *metricsCount)
    for i := 0; i < *metricsCount; i++ {
        metrics[i] = promauto.NewCounter(prometheus.CounterOpts{
            Name: fmt.Sprintf("%s_%d_total", *metricName, i),
            Help: fmt.Sprintf("Fake metric %d from the prometheus-load application", i),
        })
    }
    
    recordMetrics(metrics, *sleepTime)
    
    http.Handle("/metrics", promhttp.Handler())

    fmt.Printf("Publishing %d fake Prometheus metrics at address %s with base name of %s and refresh interval of %d seconds", *metricsCount, *addr, *metricName, *sleepTime)
    log.Fatal(http.ListenAndServe(*addr, nil))
}