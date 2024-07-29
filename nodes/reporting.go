package nodes

import (
	"os"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/metrics/influxdb"
)

var registry = metrics.NewRegistry()

func EnableMetrics(conf *Config) {
	if !conf.Metrics.Enabled {
		return
	}
	metrics.Enabled = true
	hn, err := os.Hostname()
	if err != nil {
		hn = "localhost"
	}
	tags := map[string]string{"host": hn}

	log.Info("Starting metrics", "url", conf.Metrics.Endpoint,
		"db", conf.Metrics.Database, "namespace", conf.Metrics.Namespace)

	go influxdb.InfluxDBWithTags(registry, 10*time.Second,
		conf.Metrics.Endpoint, conf.Metrics.Database,
		conf.Metrics.Username, conf.Metrics.Password, conf.Metrics.Namespace, tags)
}
