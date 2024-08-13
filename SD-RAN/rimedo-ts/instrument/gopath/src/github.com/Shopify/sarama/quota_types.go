//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:1
)

type (
	QuotaEntityType	string

	QuotaMatchType	int
)

// ref: https://github.com/apache/kafka/blob/trunk/clients/src/main/java/org/apache/kafka/common/quota/ClientQuotaEntity.java
const (
	QuotaEntityUser		QuotaEntityType	= "user"
	QuotaEntityClientID	QuotaEntityType	= "client-id"
	QuotaEntityIP		QuotaEntityType	= "ip"
)

// ref: https://github.com/apache/kafka/blob/trunk/clients/src/main/java/org/apache/kafka/common/requests/DescribeClientQuotasRequest.java
const (
	QuotaMatchExact	QuotaMatchType	= iota
	QuotaMatchDefault
	QuotaMatchAny
)

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:21
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/quota_types.go:21
var _ = _go_fuzz_dep_.CoverTab
