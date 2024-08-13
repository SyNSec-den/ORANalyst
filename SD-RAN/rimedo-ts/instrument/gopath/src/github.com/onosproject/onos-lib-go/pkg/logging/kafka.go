// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:15
package logging

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:15
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:15
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:15
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:15
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:15
)

import (
	"net/url"
	"strings"

	kafka "github.com/Shopify/sarama"
	"go.uber.org/zap"
)

func init() {
	err := zap.RegisterSink("kafka", kafkaSinkFactory)
	if err != nil {
		panic(err)
	}
}

// kafkaSink is a Kafka sink
type kafkaSink struct {
	producer	kafka.SyncProducer
	topic		string
	key		string
}

// kafkaSinkFactory is a factory for the Kafka sink
func kafkaSinkFactory(u *url.URL) (zap.Sink, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:40
	_go_fuzz_dep_.CoverTab[132036]++
													topic := "kafka_default_topic"
													key := "kafka_default_key"
													m, _ := url.ParseQuery(u.RawQuery)
													if len(m["topic"]) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:44
		_go_fuzz_dep_.CoverTab[132040]++
														topic = m["topic"][0]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:45
		// _ = "end of CoverTab[132040]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:46
		_go_fuzz_dep_.CoverTab[132041]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:46
		// _ = "end of CoverTab[132041]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:46
	// _ = "end of CoverTab[132036]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:46
	_go_fuzz_dep_.CoverTab[132037]++

													if len(m["key"]) != 0 {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:48
		_go_fuzz_dep_.CoverTab[132042]++
														key = m["key"][0]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:49
		// _ = "end of CoverTab[132042]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:50
		_go_fuzz_dep_.CoverTab[132043]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:50
		// _ = "end of CoverTab[132043]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:50
	// _ = "end of CoverTab[132037]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:50
	_go_fuzz_dep_.CoverTab[132038]++

													brokers := strings.Split(u.Host, ",")
													config := kafka.NewConfig()
													config.Producer.Return.Successes = true

													producer, err := kafka.NewSyncProducer(brokers, config)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:57
		_go_fuzz_dep_.CoverTab[132044]++
														return kafkaSink{}, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:58
		// _ = "end of CoverTab[132044]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:59
		_go_fuzz_dep_.CoverTab[132045]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:59
		// _ = "end of CoverTab[132045]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:59
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:59
	// _ = "end of CoverTab[132038]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:59
	_go_fuzz_dep_.CoverTab[132039]++

													return kafkaSink{
		producer:	producer,
		topic:		topic,
		key:		key,
	}, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:65
	// _ = "end of CoverTab[132039]"
}

// Write implements zap.Sink Write function
func (s kafkaSink) Write(b []byte) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:69
	_go_fuzz_dep_.CoverTab[132046]++
													var returnErr error
													for _, topic := range strings.Split(s.topic, ",") {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:71
		_go_fuzz_dep_.CoverTab[132048]++
														if s.key != "" {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:72
			_go_fuzz_dep_.CoverTab[132049]++
															_, _, err := s.producer.SendMessage(&kafka.ProducerMessage{
				Topic:	topic,
				Key:	kafka.StringEncoder(s.key),
				Value:	kafka.ByteEncoder(b),
			})
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:78
				_go_fuzz_dep_.CoverTab[132050]++
																returnErr = err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:79
				// _ = "end of CoverTab[132050]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:80
				_go_fuzz_dep_.CoverTab[132051]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:80
				// _ = "end of CoverTab[132051]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:80
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:80
			// _ = "end of CoverTab[132049]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:81
			_go_fuzz_dep_.CoverTab[132052]++
															_, _, err := s.producer.SendMessage(&kafka.ProducerMessage{
				Topic:	topic,
				Value:	kafka.ByteEncoder(b),
			})
			if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:86
				_go_fuzz_dep_.CoverTab[132053]++
																returnErr = err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:87
				// _ = "end of CoverTab[132053]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:88
				_go_fuzz_dep_.CoverTab[132054]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:88
				// _ = "end of CoverTab[132054]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:88
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:88
			// _ = "end of CoverTab[132052]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:89
		// _ = "end of CoverTab[132048]"

	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:91
	// _ = "end of CoverTab[132046]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:91
	_go_fuzz_dep_.CoverTab[132047]++
													return len(b), returnErr
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:92
	// _ = "end of CoverTab[132047]"
}

// Sync implement zap.Sink func Sync
func (s kafkaSink) Sync() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:96
	_go_fuzz_dep_.CoverTab[132055]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:97
	// _ = "end of CoverTab[132055]"
}

// Close implements zap.Sink Close function
func (s kafkaSink) Close() error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:101
	_go_fuzz_dep_.CoverTab[132056]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:102
	// _ = "end of CoverTab[132056]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:103
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-lib-go@v0.8.13/pkg/logging/kafka.go:103
var _ = _go_fuzz_dep_.CoverTab
