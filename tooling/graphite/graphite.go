package graphite

import (
	"github.com/adrianco/spigo/tooling/archaius"
	"github.com/marpaia/graphite-golang"
	"log"
	"strconv"
)

var GraphiteServer *graphite.Graphite
var metricPrefix string

func Setup(graphiteUrl string, graphitePort string, graphitePrefix string) error {
	if graphitePrefix != "" {
		metricPrefix = graphitePrefix + "."
	}

	port, err := strconv.Atoi(graphitePort)
	if err != nil {
		log.Println("GRAPHITEPORT environment variable must be a valid Graphite remote port value")
		return err
	}

	GraphiteServer, err = graphite.NewGraphiteUDP(graphiteUrl, port)
	if err != nil {
		log.Println("Failed to create a Graphite client")
		return err
	}
	return nil
}

func SendMetric(metricName string, metricValue string) {
	if !archaius.Conf.GraphiteEnabled {
		return
	}
	GraphiteServer.SimpleSend(metricPrefix+metricName, metricValue)
}
