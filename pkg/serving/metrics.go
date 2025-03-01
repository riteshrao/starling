package serving

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
)

// StartMetrics starts serving metrics for prometheus server scrape.
func StartMetrics(cfg *Config) {
	log.Info().Msgf("serving prometheus metrics at http://localhost:%d/metrics", cfg.MetricsPort)
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(fmt.Sprintf(":%d", cfg.MetricsPort), nil)
}
