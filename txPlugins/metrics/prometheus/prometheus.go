package main

/*
 * Capture Prometheus metrics
 */

import (
	"fmt"
	"net/http"
	"time"

	pluginMeta "github.com/keruzu/trapmux/txPlugins"

	"github.com/rs/zerolog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type prometheusStats struct {
	main_log *zerolog.Logger

	listenAddress string
	endpoint      string

	counters []prometheus.Counter
}

func (p *prometheusStats) Configure(pluginLog *zerolog.Logger, args map[string]string, metric_definitions []pluginMeta.MetricDef) error {
	p.main_log = pluginLog
	listenIP := args["listen_ip"]
	listenPort := args["listen_port"]
	p.listenAddress = listenIP + ":" + listenPort
	p.endpoint = args["endpoint"]

	for i, definition := range metric_definitions {
		p.counters[i] = promauto.NewCounter(prometheus.CounterOpts{
			Name: definition.Name,
			Help: definition.Help,
		})
	}

	exporter := fmt.Sprintf("http://%s/%s", p.listenAddress, p.endpoint)
	p.main_log.Info().Str("endpoint", exporter).Msg("Prometheus metrics exporter")

	go exposeMetrics(p.main_log, p.endpoint, p.listenAddress)

	return nil
}

func (p prometheusStats) Inc(metricIndex int) {

	p.counters[metricIndex].Inc()

}

func (p prometheusStats) Report() (string, error) {
	return "", nil
}

// exposeMetrics
// Allow Prometheus to gather current performance metrics via /metrics URL
func exposeMetrics(pluginLog *zerolog.Logger, endpoint string, listenAddress string) {
	mux := http.NewServeMux()
	mux.Handle(endpoint, promhttp.Handler())
	server := &http.Server{
		Addr:              listenAddress,
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       90 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		pluginLog.Error().Str("endpoint", endpoint).Str("listen_address", listenAddress).Msg("Prometheus metrics exporter unable to start HTTP service")
	}
}

var MetricPlugin prometheusStats
