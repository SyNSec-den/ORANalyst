//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:1
)

import (
	"fmt"
	"strings"

	"github.com/rcrowley/go-metrics"
)

// Use exponentially decaying reservoir for sampling histograms with the same defaults as the Java library:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:10
// 1028 elements, which offers a 99.9% confidence level with a 5% margin of error assuming a normal distribution,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:10
// and an alpha factor of 0.015, which heavily biases the reservoir to the past 5 minutes of measurements.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:10
// See https://github.com/dropwizard/metrics/blob/v3.1.0/metrics-core/src/main/java/com/codahale/metrics/ExponentiallyDecayingReservoir.java#L38
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:14
const (
	metricsReservoirSize	= 1028
	metricsAlphaFactor	= 0.015
)

func getOrRegisterHistogram(name string, r metrics.Registry) metrics.Histogram {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:19
	_go_fuzz_dep_.CoverTab[104239]++
											return r.GetOrRegister(name, func() metrics.Histogram {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:20
		_go_fuzz_dep_.CoverTab[104240]++
												return metrics.NewHistogram(metrics.NewExpDecaySample(metricsReservoirSize, metricsAlphaFactor))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:21
		// _ = "end of CoverTab[104240]"
	}).(metrics.Histogram)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:22
	// _ = "end of CoverTab[104239]"
}

func getMetricNameForBroker(name string, broker *Broker) string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:25
	_go_fuzz_dep_.CoverTab[104241]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:28
	return fmt.Sprintf(name+"-for-broker-%d", broker.ID())
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:28
	// _ = "end of CoverTab[104241]"
}

func getMetricNameForTopic(name string, topic string) string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:31
	_go_fuzz_dep_.CoverTab[104242]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:34
	return fmt.Sprintf(name+"-for-topic-%s", strings.Replace(topic, ".", "_", -1))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:34
	// _ = "end of CoverTab[104242]"
}

func getOrRegisterTopicMeter(name string, topic string, r metrics.Registry) metrics.Meter {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:37
	_go_fuzz_dep_.CoverTab[104243]++
											return metrics.GetOrRegisterMeter(getMetricNameForTopic(name, topic), r)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:38
	// _ = "end of CoverTab[104243]"
}

func getOrRegisterTopicHistogram(name string, topic string, r metrics.Registry) metrics.Histogram {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:41
	_go_fuzz_dep_.CoverTab[104244]++
											return getOrRegisterHistogram(getMetricNameForTopic(name, topic), r)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:42
	// _ = "end of CoverTab[104244]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/metrics.go:43
var _ = _go_fuzz_dep_.CoverTab
