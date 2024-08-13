//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:1
)

import (
	"fmt"
	"strings"
)

type (
	AclOperation	int

	AclPermissionType	int

	AclResourceType	int

	AclResourcePatternType	int
)

// ref: https://github.com/apache/kafka/blob/trunk/clients/src/main/java/org/apache/kafka/common/acl/AclOperation.java
const (
	AclOperationUnknown	AclOperation	= iota
	AclOperationAny
	AclOperationAll
	AclOperationRead
	AclOperationWrite
	AclOperationCreate
	AclOperationDelete
	AclOperationAlter
	AclOperationDescribe
	AclOperationClusterAction
	AclOperationDescribeConfigs
	AclOperationAlterConfigs
	AclOperationIdempotentWrite
)

func (a *AclOperation) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:35
	_go_fuzz_dep_.CoverTab[97373]++
											mapping := map[AclOperation]string{
		AclOperationUnknown:		"Unknown",
		AclOperationAny:		"Any",
		AclOperationAll:		"All",
		AclOperationRead:		"Read",
		AclOperationWrite:		"Write",
		AclOperationCreate:		"Create",
		AclOperationDelete:		"Delete",
		AclOperationAlter:		"Alter",
		AclOperationDescribe:		"Describe",
		AclOperationClusterAction:	"ClusterAction",
		AclOperationDescribeConfigs:	"DescribeConfigs",
		AclOperationAlterConfigs:	"AlterConfigs",
		AclOperationIdempotentWrite:	"IdempotentWrite",
	}
	s, ok := mapping[*a]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:52
		_go_fuzz_dep_.CoverTab[97375]++
												s = mapping[AclOperationUnknown]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:53
		// _ = "end of CoverTab[97375]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:54
		_go_fuzz_dep_.CoverTab[97376]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:54
		// _ = "end of CoverTab[97376]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:54
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:54
	// _ = "end of CoverTab[97373]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:54
	_go_fuzz_dep_.CoverTab[97374]++
											return s
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:55
	// _ = "end of CoverTab[97374]"
}

// MarshalText returns the text form of the AclOperation (name without prefix)
func (a *AclOperation) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:59
	_go_fuzz_dep_.CoverTab[97377]++
											return []byte(a.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:60
	// _ = "end of CoverTab[97377]"
}

// UnmarshalText takes a text reprentation of the operation and converts it to an AclOperation
func (a *AclOperation) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:64
	_go_fuzz_dep_.CoverTab[97378]++
											normalized := strings.ToLower(string(text))
											mapping := map[string]AclOperation{
		"unknown":		AclOperationUnknown,
		"any":			AclOperationAny,
		"all":			AclOperationAll,
		"read":			AclOperationRead,
		"write":		AclOperationWrite,
		"create":		AclOperationCreate,
		"delete":		AclOperationDelete,
		"alter":		AclOperationAlter,
		"describe":		AclOperationDescribe,
		"clusteraction":	AclOperationClusterAction,
		"describeconfigs":	AclOperationDescribeConfigs,
		"alterconfigs":		AclOperationAlterConfigs,
		"idempotentwrite":	AclOperationIdempotentWrite,
	}
	ao, ok := mapping[normalized]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:82
		_go_fuzz_dep_.CoverTab[97380]++
												*a = AclOperationUnknown
												return fmt.Errorf("no acl operation with name %s", normalized)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:84
		// _ = "end of CoverTab[97380]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:85
		_go_fuzz_dep_.CoverTab[97381]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:85
		// _ = "end of CoverTab[97381]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:85
	// _ = "end of CoverTab[97378]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:85
	_go_fuzz_dep_.CoverTab[97379]++
											*a = ao
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:87
	// _ = "end of CoverTab[97379]"
}

// ref: https://github.com/apache/kafka/blob/trunk/clients/src/main/java/org/apache/kafka/common/acl/AclPermissionType.java
const (
	AclPermissionUnknown	AclPermissionType	= iota
	AclPermissionAny
	AclPermissionDeny
	AclPermissionAllow
)

func (a *AclPermissionType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:98
	_go_fuzz_dep_.CoverTab[97382]++
											mapping := map[AclPermissionType]string{
		AclPermissionUnknown:	"Unknown",
		AclPermissionAny:	"Any",
		AclPermissionDeny:	"Deny",
		AclPermissionAllow:	"Allow",
	}
	s, ok := mapping[*a]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:106
		_go_fuzz_dep_.CoverTab[97384]++
												s = mapping[AclPermissionUnknown]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:107
		// _ = "end of CoverTab[97384]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:108
		_go_fuzz_dep_.CoverTab[97385]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:108
		// _ = "end of CoverTab[97385]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:108
	// _ = "end of CoverTab[97382]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:108
	_go_fuzz_dep_.CoverTab[97383]++
											return s
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:109
	// _ = "end of CoverTab[97383]"
}

// MarshalText returns the text form of the AclPermissionType (name without prefix)
func (a *AclPermissionType) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:113
	_go_fuzz_dep_.CoverTab[97386]++
											return []byte(a.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:114
	// _ = "end of CoverTab[97386]"
}

// UnmarshalText takes a text reprentation of the permission type and converts it to an AclPermissionType
func (a *AclPermissionType) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:118
	_go_fuzz_dep_.CoverTab[97387]++
											normalized := strings.ToLower(string(text))
											mapping := map[string]AclPermissionType{
		"unknown":	AclPermissionUnknown,
		"any":		AclPermissionAny,
		"deny":		AclPermissionDeny,
		"allow":	AclPermissionAllow,
	}

	apt, ok := mapping[normalized]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:128
		_go_fuzz_dep_.CoverTab[97389]++
												*a = AclPermissionUnknown
												return fmt.Errorf("no acl permission with name %s", normalized)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:130
		// _ = "end of CoverTab[97389]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:131
		_go_fuzz_dep_.CoverTab[97390]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:131
		// _ = "end of CoverTab[97390]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:131
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:131
	// _ = "end of CoverTab[97387]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:131
	_go_fuzz_dep_.CoverTab[97388]++
											*a = apt
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:133
	// _ = "end of CoverTab[97388]"
}

// ref: https://github.com/apache/kafka/blob/trunk/clients/src/main/java/org/apache/kafka/common/resource/ResourceType.java
const (
	AclResourceUnknown	AclResourceType	= iota
	AclResourceAny
	AclResourceTopic
	AclResourceGroup
	AclResourceCluster
	AclResourceTransactionalID
	AclResourceDelegationToken
)

func (a *AclResourceType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:147
	_go_fuzz_dep_.CoverTab[97391]++
											mapping := map[AclResourceType]string{
		AclResourceUnknown:		"Unknown",
		AclResourceAny:			"Any",
		AclResourceTopic:		"Topic",
		AclResourceGroup:		"Group",
		AclResourceCluster:		"Cluster",
		AclResourceTransactionalID:	"TransactionalID",
		AclResourceDelegationToken:	"DelegationToken",
	}
	s, ok := mapping[*a]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:158
		_go_fuzz_dep_.CoverTab[97393]++
												s = mapping[AclResourceUnknown]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:159
		// _ = "end of CoverTab[97393]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:160
		_go_fuzz_dep_.CoverTab[97394]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:160
		// _ = "end of CoverTab[97394]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:160
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:160
	// _ = "end of CoverTab[97391]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:160
	_go_fuzz_dep_.CoverTab[97392]++
											return s
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:161
	// _ = "end of CoverTab[97392]"
}

// MarshalText returns the text form of the AclResourceType (name without prefix)
func (a *AclResourceType) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:165
	_go_fuzz_dep_.CoverTab[97395]++
											return []byte(a.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:166
	// _ = "end of CoverTab[97395]"
}

// UnmarshalText takes a text reprentation of the resource type and converts it to an AclResourceType
func (a *AclResourceType) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:170
	_go_fuzz_dep_.CoverTab[97396]++
											normalized := strings.ToLower(string(text))
											mapping := map[string]AclResourceType{
		"unknown":		AclResourceUnknown,
		"any":			AclResourceAny,
		"topic":		AclResourceTopic,
		"group":		AclResourceGroup,
		"cluster":		AclResourceCluster,
		"transactionalid":	AclResourceTransactionalID,
		"delegationtoken":	AclResourceDelegationToken,
	}

	art, ok := mapping[normalized]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:183
		_go_fuzz_dep_.CoverTab[97398]++
												*a = AclResourceUnknown
												return fmt.Errorf("no acl resource with name %s", normalized)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:185
		// _ = "end of CoverTab[97398]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:186
		_go_fuzz_dep_.CoverTab[97399]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:186
		// _ = "end of CoverTab[97399]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:186
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:186
	// _ = "end of CoverTab[97396]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:186
	_go_fuzz_dep_.CoverTab[97397]++
											*a = art
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:188
	// _ = "end of CoverTab[97397]"
}

// ref: https://github.com/apache/kafka/blob/trunk/clients/src/main/java/org/apache/kafka/common/resource/PatternType.java
const (
	AclPatternUnknown	AclResourcePatternType	= iota
	AclPatternAny
	AclPatternMatch
	AclPatternLiteral
	AclPatternPrefixed
)

func (a *AclResourcePatternType) String() string {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:200
	_go_fuzz_dep_.CoverTab[97400]++
											mapping := map[AclResourcePatternType]string{
		AclPatternUnknown:	"Unknown",
		AclPatternAny:		"Any",
		AclPatternMatch:	"Match",
		AclPatternLiteral:	"Literal",
		AclPatternPrefixed:	"Prefixed",
	}
	s, ok := mapping[*a]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:209
		_go_fuzz_dep_.CoverTab[97402]++
												s = mapping[AclPatternUnknown]
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:210
		// _ = "end of CoverTab[97402]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:211
		_go_fuzz_dep_.CoverTab[97403]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:211
		// _ = "end of CoverTab[97403]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:211
	// _ = "end of CoverTab[97400]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:211
	_go_fuzz_dep_.CoverTab[97401]++
											return s
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:212
	// _ = "end of CoverTab[97401]"
}

// MarshalText returns the text form of the AclResourcePatternType (name without prefix)
func (a *AclResourcePatternType) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:216
	_go_fuzz_dep_.CoverTab[97404]++
											return []byte(a.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:217
	// _ = "end of CoverTab[97404]"
}

// UnmarshalText takes a text reprentation of the resource pattern type and converts it to an AclResourcePatternType
func (a *AclResourcePatternType) UnmarshalText(text []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:221
	_go_fuzz_dep_.CoverTab[97405]++
											normalized := strings.ToLower(string(text))
											mapping := map[string]AclResourcePatternType{
		"unknown":	AclPatternUnknown,
		"any":		AclPatternAny,
		"match":	AclPatternMatch,
		"literal":	AclPatternLiteral,
		"prefixed":	AclPatternPrefixed,
	}

	arpt, ok := mapping[normalized]
	if !ok {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:232
		_go_fuzz_dep_.CoverTab[97407]++
												*a = AclPatternUnknown
												return fmt.Errorf("no acl resource pattern with name %s", normalized)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:234
		// _ = "end of CoverTab[97407]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:235
		_go_fuzz_dep_.CoverTab[97408]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:235
		// _ = "end of CoverTab[97408]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:235
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:235
	// _ = "end of CoverTab[97405]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:235
	_go_fuzz_dep_.CoverTab[97406]++
											*a = arpt
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:237
	// _ = "end of CoverTab[97406]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:238
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/acl_types.go:238
var _ = _go_fuzz_dep_.CoverTab
