package impl

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/graphite"
)

type Metrics struct {
	grap                *graphite.Graphite
	EntryRequestCount   metrics.Counter
	EntryRequestLatency metrics.Histogram
	EntryMetricFunction func(begin time.Time)

	MediaRequestCount   metrics.Counter
	MediaRequestLatency metrics.Histogram
	MediaMetricFunction func(begin time.Time)
}

var mets = Metrics{}

// network: tcp/udp
func InitMetrics(logger Logger, network string, address string) {
	if mets.grap == nil {
		mets.grap = graphite.New("strapi-webhook.", logger)
		mets.EntryRequestCount = mets.grap.NewCounter("Entry.RequestCount")
		mets.EntryRequestLatency = mets.grap.NewHistogram("Entry.RequestLatency", 50)
		mets.EntryMetricFunction = func(begin time.Time) {
			mets.EntryRequestCount.Add(1)
			mets.EntryRequestLatency.Observe(time.Since(begin).Seconds())
		}

		mets.MediaRequestCount = mets.grap.NewCounter("Media.RequestCount")
		mets.MediaRequestLatency = mets.grap.NewHistogram("Media.RequestLatency", 50)
		mets.MediaMetricFunction = func(begin time.Time) {
			mets.MediaRequestCount.Add(1)
			mets.MediaRequestLatency.Observe(time.Since(begin).Seconds())
		}

		t := time.NewTicker(time.Second * time.Duration(10))
		go mets.grap.SendLoop(context.Background(), t.C, network, address)
	}
}

func GetMetrics() Metrics {
	return mets
}
