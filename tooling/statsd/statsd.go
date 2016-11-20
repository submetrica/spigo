package statsd

import (
	"github.com/cactus/go-statsd-client/statsd"

	"github.com/adrianco/spigo/tooling/archaius"
	"time"
)

var StatsdClient *statsd.Client
var apiKey string

func Setup(statsdUrl string, statsdPort string, statsdApiKey string) error {
	client, err := statsd.NewClient(statsdUrl+":"+statsdPort, "simianviz-client")

	if err == nil {
		StatsdClient = client.(*statsd.Client)
	}

	apiKey = statsdApiKey

	return err
}

func addTagsToMetric(metricName string, tags map[string]string) string {
	for k, v := range tags {
		metricName += "," + k + "=" + v
	}
	return metricName
}

func addApiKeyToMetric(metricName string, apiKey string) string {
	if apiKey != "" {
		return metricName + ",apikey=" + apiKey
	}
	return metricName
}

func Counter(metricName string, tags map[string]string, metricValue int64) {
	if !archaius.Conf.StatsdEnabled {
		return
	}

	var metric = addApiKeyToMetric(metricName, apiKey)
	metric = addTagsToMetric(metric, tags)

	StatsdClient.Inc(metric, metricValue, 1.0)
}

func TimingDuration(metricName string, tags map[string]string, metricValue time.Duration) {
	if !archaius.Conf.StatsdEnabled {
		return
	}

	var metric = addApiKeyToMetric(metricName, apiKey)
	metric = addTagsToMetric(metric, tags)

	StatsdClient.TimingDuration(metric, metricValue, 1.0)
}
