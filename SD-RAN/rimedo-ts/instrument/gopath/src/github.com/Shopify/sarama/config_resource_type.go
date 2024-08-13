//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:1
)

// ConfigResourceType is a type for resources that have configs.
type ConfigResourceType int8

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:9
const (
	// UnknownResource constant type
	UnknownResource	ConfigResourceType	= 0
	// TopicResource constant type
	TopicResource	ConfigResourceType	= 2
	// BrokerResource constant type
	BrokerResource	ConfigResourceType	= 4
	// BrokerLoggerResource constant type
	BrokerLoggerResource	ConfigResourceType	= 8
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:18
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/config_resource_type.go:18
var _ = _go_fuzz_dep_.CoverTab
