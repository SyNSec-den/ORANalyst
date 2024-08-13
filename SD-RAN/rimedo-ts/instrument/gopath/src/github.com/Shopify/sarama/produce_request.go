//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:1
)

import "github.com/rcrowley/go-metrics"

// RequiredAcks is used in Produce Requests to tell the broker how many replica acknowledgements
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:5
// it must see before responding. Any of the constants defined here are valid. On broker versions
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:5
// prior to 0.8.2.0 any other positive int16 is also valid (the broker will wait for that many
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:5
// acknowledgements) but in 0.8.2.0 and later this will raise an exception (it has been replaced
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:5
// by setting the `min.isr` value in the brokers configuration).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:10
type RequiredAcks int16

const (
	// NoResponse doesn't send any response, the TCP ACK is all you get.
	NoResponse	RequiredAcks	= 0
	// WaitForLocal waits for only the local commit to succeed before responding.
	WaitForLocal	RequiredAcks	= 1
	// WaitForAll waits for all in-sync replicas to commit before responding.
	// The minimum number of in-sync replicas is configured on the broker via
	// the `min.insync.replicas` configuration key.
	WaitForAll	RequiredAcks	= -1
)

type ProduceRequest struct {
	TransactionalID	*string
	RequiredAcks	RequiredAcks
	Timeout		int32
	Version		int16	// v1 requires Kafka 0.9, v2 requires Kafka 0.10, v3 requires Kafka 0.11
	records		map[string]map[int32]Records
}

func updateMsgSetMetrics(msgSet *MessageSet, compressionRatioMetric metrics.Histogram,
	topicCompressionRatioMetric metrics.Histogram) int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:32
	_go_fuzz_dep_.CoverTab[105661]++
												var topicRecordCount int64
												for _, messageBlock := range msgSet.Messages {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:34
		_go_fuzz_dep_.CoverTab[105663]++

													if messageBlock.Msg.Set != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:36
			_go_fuzz_dep_.CoverTab[105665]++
														topicRecordCount += int64(len(messageBlock.Msg.Set.Messages))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:37
			// _ = "end of CoverTab[105665]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:38
			_go_fuzz_dep_.CoverTab[105666]++

														topicRecordCount++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:40
			// _ = "end of CoverTab[105666]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:41
		// _ = "end of CoverTab[105663]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:41
		_go_fuzz_dep_.CoverTab[105664]++

													if messageBlock.Msg.compressedSize != 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:43
			_go_fuzz_dep_.CoverTab[105667]++
														compressionRatio := float64(len(messageBlock.Msg.Value)) /
				float64(messageBlock.Msg.compressedSize)

														intCompressionRatio := int64(100 * compressionRatio)
														compressionRatioMetric.Update(intCompressionRatio)
														topicCompressionRatioMetric.Update(intCompressionRatio)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:49
			// _ = "end of CoverTab[105667]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:50
			_go_fuzz_dep_.CoverTab[105668]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:50
			// _ = "end of CoverTab[105668]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:50
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:50
		// _ = "end of CoverTab[105664]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:51
	// _ = "end of CoverTab[105661]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:51
	_go_fuzz_dep_.CoverTab[105662]++
												return topicRecordCount
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:52
	// _ = "end of CoverTab[105662]"
}

func updateBatchMetrics(recordBatch *RecordBatch, compressionRatioMetric metrics.Histogram,
	topicCompressionRatioMetric metrics.Histogram) int64 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:56
	_go_fuzz_dep_.CoverTab[105669]++
												if recordBatch.compressedRecords != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:57
		_go_fuzz_dep_.CoverTab[105671]++
													compressionRatio := int64(float64(recordBatch.recordsLen) / float64(len(recordBatch.compressedRecords)) * 100)
													compressionRatioMetric.Update(compressionRatio)
													topicCompressionRatioMetric.Update(compressionRatio)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:60
		// _ = "end of CoverTab[105671]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:61
		_go_fuzz_dep_.CoverTab[105672]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:61
		// _ = "end of CoverTab[105672]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:61
	// _ = "end of CoverTab[105669]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:61
	_go_fuzz_dep_.CoverTab[105670]++

												return int64(len(recordBatch.Records))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:63
	// _ = "end of CoverTab[105670]"
}

func (r *ProduceRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:66
	_go_fuzz_dep_.CoverTab[105673]++
												if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:67
		_go_fuzz_dep_.CoverTab[105679]++
													if err := pe.putNullableString(r.TransactionalID); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:68
			_go_fuzz_dep_.CoverTab[105680]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:69
			// _ = "end of CoverTab[105680]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:70
			_go_fuzz_dep_.CoverTab[105681]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:70
			// _ = "end of CoverTab[105681]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:70
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:70
		// _ = "end of CoverTab[105679]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:71
		_go_fuzz_dep_.CoverTab[105682]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:71
		// _ = "end of CoverTab[105682]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:71
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:71
	// _ = "end of CoverTab[105673]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:71
	_go_fuzz_dep_.CoverTab[105674]++
												pe.putInt16(int16(r.RequiredAcks))
												pe.putInt32(r.Timeout)
												metricRegistry := pe.metricRegistry()
												var batchSizeMetric metrics.Histogram
												var compressionRatioMetric metrics.Histogram
												if metricRegistry != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:77
		_go_fuzz_dep_.CoverTab[105683]++
													batchSizeMetric = getOrRegisterHistogram("batch-size", metricRegistry)
													compressionRatioMetric = getOrRegisterHistogram("compression-ratio", metricRegistry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:79
		// _ = "end of CoverTab[105683]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:80
		_go_fuzz_dep_.CoverTab[105684]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:80
		// _ = "end of CoverTab[105684]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:80
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:80
	// _ = "end of CoverTab[105674]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:80
	_go_fuzz_dep_.CoverTab[105675]++
												totalRecordCount := int64(0)

												err := pe.putArrayLength(len(r.records))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:84
		_go_fuzz_dep_.CoverTab[105685]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:85
		// _ = "end of CoverTab[105685]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:86
		_go_fuzz_dep_.CoverTab[105686]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:86
		// _ = "end of CoverTab[105686]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:86
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:86
	// _ = "end of CoverTab[105675]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:86
	_go_fuzz_dep_.CoverTab[105676]++

												for topic, partitions := range r.records {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:88
		_go_fuzz_dep_.CoverTab[105687]++
													err = pe.putString(topic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:90
			_go_fuzz_dep_.CoverTab[105692]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:91
			// _ = "end of CoverTab[105692]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:92
			_go_fuzz_dep_.CoverTab[105693]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:92
			// _ = "end of CoverTab[105693]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:92
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:92
		// _ = "end of CoverTab[105687]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:92
		_go_fuzz_dep_.CoverTab[105688]++
													err = pe.putArrayLength(len(partitions))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:94
			_go_fuzz_dep_.CoverTab[105694]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:95
			// _ = "end of CoverTab[105694]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:96
			_go_fuzz_dep_.CoverTab[105695]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:96
			// _ = "end of CoverTab[105695]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:96
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:96
		// _ = "end of CoverTab[105688]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:96
		_go_fuzz_dep_.CoverTab[105689]++
													topicRecordCount := int64(0)
													var topicCompressionRatioMetric metrics.Histogram
													if metricRegistry != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:99
			_go_fuzz_dep_.CoverTab[105696]++
														topicCompressionRatioMetric = getOrRegisterTopicHistogram("compression-ratio", topic, metricRegistry)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:100
			// _ = "end of CoverTab[105696]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:101
			_go_fuzz_dep_.CoverTab[105697]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:101
			// _ = "end of CoverTab[105697]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:101
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:101
		// _ = "end of CoverTab[105689]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:101
		_go_fuzz_dep_.CoverTab[105690]++
													for id, records := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:102
			_go_fuzz_dep_.CoverTab[105698]++
														startOffset := pe.offset()
														pe.putInt32(id)
														pe.push(&lengthField{})
														err = records.encode(pe)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:107
				_go_fuzz_dep_.CoverTab[105701]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:108
				// _ = "end of CoverTab[105701]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:109
				_go_fuzz_dep_.CoverTab[105702]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:109
				// _ = "end of CoverTab[105702]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:109
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:109
			// _ = "end of CoverTab[105698]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:109
			_go_fuzz_dep_.CoverTab[105699]++
														err = pe.pop()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:111
				_go_fuzz_dep_.CoverTab[105703]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:112
				// _ = "end of CoverTab[105703]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:113
				_go_fuzz_dep_.CoverTab[105704]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:113
				// _ = "end of CoverTab[105704]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:113
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:113
			// _ = "end of CoverTab[105699]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:113
			_go_fuzz_dep_.CoverTab[105700]++
														if metricRegistry != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:114
				_go_fuzz_dep_.CoverTab[105705]++
															if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:115
					_go_fuzz_dep_.CoverTab[105707]++
																topicRecordCount += updateBatchMetrics(records.RecordBatch, compressionRatioMetric, topicCompressionRatioMetric)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:116
					// _ = "end of CoverTab[105707]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:117
					_go_fuzz_dep_.CoverTab[105708]++
																topicRecordCount += updateMsgSetMetrics(records.MsgSet, compressionRatioMetric, topicCompressionRatioMetric)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:118
					// _ = "end of CoverTab[105708]"
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:119
				// _ = "end of CoverTab[105705]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:119
				_go_fuzz_dep_.CoverTab[105706]++
															batchSize := int64(pe.offset() - startOffset)
															batchSizeMetric.Update(batchSize)
															getOrRegisterTopicHistogram("batch-size", topic, metricRegistry).Update(batchSize)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:122
				// _ = "end of CoverTab[105706]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:123
				_go_fuzz_dep_.CoverTab[105709]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:123
				// _ = "end of CoverTab[105709]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:123
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:123
			// _ = "end of CoverTab[105700]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:124
		// _ = "end of CoverTab[105690]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:124
		_go_fuzz_dep_.CoverTab[105691]++
													if topicRecordCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:125
			_go_fuzz_dep_.CoverTab[105710]++
														getOrRegisterTopicMeter("record-send-rate", topic, metricRegistry).Mark(topicRecordCount)
														getOrRegisterTopicHistogram("records-per-request", topic, metricRegistry).Update(topicRecordCount)
														totalRecordCount += topicRecordCount
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:128
			// _ = "end of CoverTab[105710]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:129
			_go_fuzz_dep_.CoverTab[105711]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:129
			// _ = "end of CoverTab[105711]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:129
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:129
		// _ = "end of CoverTab[105691]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:130
	// _ = "end of CoverTab[105676]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:130
	_go_fuzz_dep_.CoverTab[105677]++
												if totalRecordCount > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:131
		_go_fuzz_dep_.CoverTab[105712]++
													metrics.GetOrRegisterMeter("record-send-rate", metricRegistry).Mark(totalRecordCount)
													getOrRegisterHistogram("records-per-request", metricRegistry).Update(totalRecordCount)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:133
		// _ = "end of CoverTab[105712]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:134
		_go_fuzz_dep_.CoverTab[105713]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:134
		// _ = "end of CoverTab[105713]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:134
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:134
	// _ = "end of CoverTab[105677]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:134
	_go_fuzz_dep_.CoverTab[105678]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:136
	// _ = "end of CoverTab[105678]"
}

func (r *ProduceRequest) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:139
	_go_fuzz_dep_.CoverTab[105714]++
												r.Version = version

												if version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:142
		_go_fuzz_dep_.CoverTab[105721]++
													id, err := pd.getNullableString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:144
			_go_fuzz_dep_.CoverTab[105723]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:145
			// _ = "end of CoverTab[105723]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:146
			_go_fuzz_dep_.CoverTab[105724]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:146
			// _ = "end of CoverTab[105724]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:146
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:146
		// _ = "end of CoverTab[105721]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:146
		_go_fuzz_dep_.CoverTab[105722]++
													r.TransactionalID = id
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:147
		// _ = "end of CoverTab[105722]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:148
		_go_fuzz_dep_.CoverTab[105725]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:148
		// _ = "end of CoverTab[105725]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:148
	// _ = "end of CoverTab[105714]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:148
	_go_fuzz_dep_.CoverTab[105715]++
												requiredAcks, err := pd.getInt16()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:150
		_go_fuzz_dep_.CoverTab[105726]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:151
		// _ = "end of CoverTab[105726]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:152
		_go_fuzz_dep_.CoverTab[105727]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:152
		// _ = "end of CoverTab[105727]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:152
	// _ = "end of CoverTab[105715]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:152
	_go_fuzz_dep_.CoverTab[105716]++
												r.RequiredAcks = RequiredAcks(requiredAcks)
												if r.Timeout, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:154
		_go_fuzz_dep_.CoverTab[105728]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:155
		// _ = "end of CoverTab[105728]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:156
		_go_fuzz_dep_.CoverTab[105729]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:156
		// _ = "end of CoverTab[105729]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:156
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:156
	// _ = "end of CoverTab[105716]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:156
	_go_fuzz_dep_.CoverTab[105717]++
												topicCount, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:158
		_go_fuzz_dep_.CoverTab[105730]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:159
		// _ = "end of CoverTab[105730]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:160
		_go_fuzz_dep_.CoverTab[105731]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:160
		// _ = "end of CoverTab[105731]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:160
	// _ = "end of CoverTab[105717]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:160
	_go_fuzz_dep_.CoverTab[105718]++
												if topicCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:161
		_go_fuzz_dep_.CoverTab[105732]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:162
		// _ = "end of CoverTab[105732]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:163
		_go_fuzz_dep_.CoverTab[105733]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:163
		// _ = "end of CoverTab[105733]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:163
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:163
	// _ = "end of CoverTab[105718]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:163
	_go_fuzz_dep_.CoverTab[105719]++

												r.records = make(map[string]map[int32]Records)
												for i := 0; i < topicCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:166
		_go_fuzz_dep_.CoverTab[105734]++
													topic, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:168
			_go_fuzz_dep_.CoverTab[105737]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:169
			// _ = "end of CoverTab[105737]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:170
			_go_fuzz_dep_.CoverTab[105738]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:170
			// _ = "end of CoverTab[105738]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:170
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:170
		// _ = "end of CoverTab[105734]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:170
		_go_fuzz_dep_.CoverTab[105735]++
													partitionCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:172
			_go_fuzz_dep_.CoverTab[105739]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:173
			// _ = "end of CoverTab[105739]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:174
			_go_fuzz_dep_.CoverTab[105740]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:174
			// _ = "end of CoverTab[105740]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:174
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:174
		// _ = "end of CoverTab[105735]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:174
		_go_fuzz_dep_.CoverTab[105736]++
													r.records[topic] = make(map[int32]Records)

													for j := 0; j < partitionCount; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:177
			_go_fuzz_dep_.CoverTab[105741]++
														partition, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:179
				_go_fuzz_dep_.CoverTab[105746]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:180
				// _ = "end of CoverTab[105746]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:181
				_go_fuzz_dep_.CoverTab[105747]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:181
				// _ = "end of CoverTab[105747]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:181
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:181
			// _ = "end of CoverTab[105741]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:181
			_go_fuzz_dep_.CoverTab[105742]++
														size, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:183
				_go_fuzz_dep_.CoverTab[105748]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:184
				// _ = "end of CoverTab[105748]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:185
				_go_fuzz_dep_.CoverTab[105749]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:185
				// _ = "end of CoverTab[105749]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:185
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:185
			// _ = "end of CoverTab[105742]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:185
			_go_fuzz_dep_.CoverTab[105743]++
														recordsDecoder, err := pd.getSubset(int(size))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:187
				_go_fuzz_dep_.CoverTab[105750]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:188
				// _ = "end of CoverTab[105750]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:189
				_go_fuzz_dep_.CoverTab[105751]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:189
				// _ = "end of CoverTab[105751]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:189
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:189
			// _ = "end of CoverTab[105743]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:189
			_go_fuzz_dep_.CoverTab[105744]++
														var records Records
														if err := records.decode(recordsDecoder); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:191
				_go_fuzz_dep_.CoverTab[105752]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:192
				// _ = "end of CoverTab[105752]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:193
				_go_fuzz_dep_.CoverTab[105753]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:193
				// _ = "end of CoverTab[105753]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:193
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:193
			// _ = "end of CoverTab[105744]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:193
			_go_fuzz_dep_.CoverTab[105745]++
														r.records[topic][partition] = records
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:194
			// _ = "end of CoverTab[105745]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:195
		// _ = "end of CoverTab[105736]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:196
	// _ = "end of CoverTab[105719]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:196
	_go_fuzz_dep_.CoverTab[105720]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:198
	// _ = "end of CoverTab[105720]"
}

func (r *ProduceRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:201
	_go_fuzz_dep_.CoverTab[105754]++
												return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:202
	// _ = "end of CoverTab[105754]"
}

func (r *ProduceRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:205
	_go_fuzz_dep_.CoverTab[105755]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:206
	// _ = "end of CoverTab[105755]"
}

func (r *ProduceRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:209
	_go_fuzz_dep_.CoverTab[105756]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:210
	// _ = "end of CoverTab[105756]"
}

func (r *ProduceRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:213
	_go_fuzz_dep_.CoverTab[105757]++
												switch r.Version {
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:215
		_go_fuzz_dep_.CoverTab[105758]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:216
		// _ = "end of CoverTab[105758]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:217
		_go_fuzz_dep_.CoverTab[105759]++
													return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:218
		// _ = "end of CoverTab[105759]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:219
		_go_fuzz_dep_.CoverTab[105760]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:220
		// _ = "end of CoverTab[105760]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:221
		_go_fuzz_dep_.CoverTab[105761]++
													return V2_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:222
		// _ = "end of CoverTab[105761]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:223
		_go_fuzz_dep_.CoverTab[105762]++
													return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:224
		// _ = "end of CoverTab[105762]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:225
	// _ = "end of CoverTab[105757]"
}

func (r *ProduceRequest) ensureRecords(topic string, partition int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:228
	_go_fuzz_dep_.CoverTab[105763]++
												if r.records == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:229
		_go_fuzz_dep_.CoverTab[105765]++
													r.records = make(map[string]map[int32]Records)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:230
		// _ = "end of CoverTab[105765]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:231
		_go_fuzz_dep_.CoverTab[105766]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:231
		// _ = "end of CoverTab[105766]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:231
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:231
	// _ = "end of CoverTab[105763]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:231
	_go_fuzz_dep_.CoverTab[105764]++

												if r.records[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:233
		_go_fuzz_dep_.CoverTab[105767]++
													r.records[topic] = make(map[int32]Records)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:234
		// _ = "end of CoverTab[105767]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:235
		_go_fuzz_dep_.CoverTab[105768]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:235
		// _ = "end of CoverTab[105768]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:235
	// _ = "end of CoverTab[105764]"
}

func (r *ProduceRequest) AddMessage(topic string, partition int32, msg *Message) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:238
	_go_fuzz_dep_.CoverTab[105769]++
												r.ensureRecords(topic, partition)
												set := r.records[topic][partition].MsgSet

												if set == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:242
		_go_fuzz_dep_.CoverTab[105771]++
													set = new(MessageSet)
													r.records[topic][partition] = newLegacyRecords(set)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:244
		// _ = "end of CoverTab[105771]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:245
		_go_fuzz_dep_.CoverTab[105772]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:245
		// _ = "end of CoverTab[105772]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:245
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:245
	// _ = "end of CoverTab[105769]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:245
	_go_fuzz_dep_.CoverTab[105770]++

												set.addMessage(msg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:247
	// _ = "end of CoverTab[105770]"
}

func (r *ProduceRequest) AddSet(topic string, partition int32, set *MessageSet) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:250
	_go_fuzz_dep_.CoverTab[105773]++
												r.ensureRecords(topic, partition)
												r.records[topic][partition] = newLegacyRecords(set)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:252
	// _ = "end of CoverTab[105773]"
}

func (r *ProduceRequest) AddBatch(topic string, partition int32, batch *RecordBatch) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:255
	_go_fuzz_dep_.CoverTab[105774]++
												r.ensureRecords(topic, partition)
												r.records[topic][partition] = newDefaultRecords(batch)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:257
	// _ = "end of CoverTab[105774]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:258
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/produce_request.go:258
var _ = _go_fuzz_dep_.CoverTab
