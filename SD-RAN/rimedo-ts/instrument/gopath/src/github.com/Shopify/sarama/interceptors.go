//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:1
)

// ProducerInterceptor allows you to intercept (and possibly mutate) the records
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:3
// received by the producer before they are published to the Kafka cluster.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:3
// https://cwiki.apache.org/confluence/display/KAFKA/KIP-42%3A+Add+Producer+and+Consumer+Interceptors#KIP42:AddProducerandConsumerInterceptors-Motivation
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:6
type ProducerInterceptor interface {

	// OnSend is called when the producer message is intercepted. Please avoid
	// modifying the message until it's safe to do so, as this is _not_ a copy
	// of the message.
	OnSend(*ProducerMessage)
}

// ConsumerInterceptor allows you to intercept (and possibly mutate) the records
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:14
// received by the consumer before they are sent to the messages channel.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:14
// https://cwiki.apache.org/confluence/display/KAFKA/KIP-42%3A+Add+Producer+and+Consumer+Interceptors#KIP42:AddProducerandConsumerInterceptors-Motivation
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:17
type ConsumerInterceptor interface {

	// OnConsume is called when the consumed message is intercepted. Please
	// avoid modifying the message until it's safe to do so, as this is _not_ a
	// copy of the message.
	OnConsume(*ConsumerMessage)
}

func (msg *ProducerMessage) safelyApplyInterceptor(interceptor ProducerInterceptor) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:25
	_go_fuzz_dep_.CoverTab[103549]++
												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:26
		_go_fuzz_dep_.CoverTab[103551]++
													if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:27
			_go_fuzz_dep_.CoverTab[103552]++
														Logger.Printf("Error when calling producer interceptor: %s, %w\n", interceptor, r)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:28
			// _ = "end of CoverTab[103552]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:29
			_go_fuzz_dep_.CoverTab[103553]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:29
			// _ = "end of CoverTab[103553]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:29
		// _ = "end of CoverTab[103551]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:30
	// _ = "end of CoverTab[103549]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:30
	_go_fuzz_dep_.CoverTab[103550]++

												interceptor.OnSend(msg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:32
	// _ = "end of CoverTab[103550]"
}

func (msg *ConsumerMessage) safelyApplyInterceptor(interceptor ConsumerInterceptor) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:35
	_go_fuzz_dep_.CoverTab[103554]++
												defer func() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:36
		_go_fuzz_dep_.CoverTab[103556]++
													if r := recover(); r != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:37
			_go_fuzz_dep_.CoverTab[103557]++
														Logger.Printf("Error when calling consumer interceptor: %s, %w\n", interceptor, r)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:38
			// _ = "end of CoverTab[103557]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:39
			_go_fuzz_dep_.CoverTab[103558]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:39
			// _ = "end of CoverTab[103558]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:39
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:39
		// _ = "end of CoverTab[103556]"
	}()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:40
	// _ = "end of CoverTab[103554]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:40
	_go_fuzz_dep_.CoverTab[103555]++

												interceptor.OnConsume(msg)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:42
	// _ = "end of CoverTab[103555]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:43
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/interceptors.go:43
var _ = _go_fuzz_dep_.CoverTab
