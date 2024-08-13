//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:1
)

type ListGroupsResponse struct {
	Err	KError
	Groups	map[string]string
}

func (r *ListGroupsResponse) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:8
	_go_fuzz_dep_.CoverTab[103790]++
												pe.putInt16(int16(r.Err))

												if err := pe.putArrayLength(len(r.Groups)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:11
		_go_fuzz_dep_.CoverTab[103793]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:12
		// _ = "end of CoverTab[103793]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:13
		_go_fuzz_dep_.CoverTab[103794]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:13
		// _ = "end of CoverTab[103794]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:13
	// _ = "end of CoverTab[103790]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:13
	_go_fuzz_dep_.CoverTab[103791]++
													for groupId, protocolType := range r.Groups {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:14
		_go_fuzz_dep_.CoverTab[103795]++
														if err := pe.putString(groupId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:15
			_go_fuzz_dep_.CoverTab[103797]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:16
			// _ = "end of CoverTab[103797]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:17
			_go_fuzz_dep_.CoverTab[103798]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:17
			// _ = "end of CoverTab[103798]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:17
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:17
		// _ = "end of CoverTab[103795]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:17
		_go_fuzz_dep_.CoverTab[103796]++
														if err := pe.putString(protocolType); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:18
			_go_fuzz_dep_.CoverTab[103799]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:19
			// _ = "end of CoverTab[103799]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:20
			_go_fuzz_dep_.CoverTab[103800]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:20
			// _ = "end of CoverTab[103800]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:20
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:20
		// _ = "end of CoverTab[103796]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:21
	// _ = "end of CoverTab[103791]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:21
	_go_fuzz_dep_.CoverTab[103792]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:23
	// _ = "end of CoverTab[103792]"
}

func (r *ListGroupsResponse) decode(pd packetDecoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:26
	_go_fuzz_dep_.CoverTab[103801]++
													kerr, err := pd.getInt16()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:28
		_go_fuzz_dep_.CoverTab[103806]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:29
		// _ = "end of CoverTab[103806]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:30
		_go_fuzz_dep_.CoverTab[103807]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:30
		// _ = "end of CoverTab[103807]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:30
	// _ = "end of CoverTab[103801]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:30
	_go_fuzz_dep_.CoverTab[103802]++

													r.Err = KError(kerr)

													n, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:35
		_go_fuzz_dep_.CoverTab[103808]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:36
		// _ = "end of CoverTab[103808]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:37
		_go_fuzz_dep_.CoverTab[103809]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:37
		// _ = "end of CoverTab[103809]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:37
	// _ = "end of CoverTab[103802]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:37
	_go_fuzz_dep_.CoverTab[103803]++
													if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:38
		_go_fuzz_dep_.CoverTab[103810]++
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:39
		// _ = "end of CoverTab[103810]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:40
		_go_fuzz_dep_.CoverTab[103811]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:40
		// _ = "end of CoverTab[103811]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:40
	// _ = "end of CoverTab[103803]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:40
	_go_fuzz_dep_.CoverTab[103804]++

													r.Groups = make(map[string]string)
													for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:43
		_go_fuzz_dep_.CoverTab[103812]++
														groupId, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:45
			_go_fuzz_dep_.CoverTab[103815]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:46
			// _ = "end of CoverTab[103815]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:47
			_go_fuzz_dep_.CoverTab[103816]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:47
			// _ = "end of CoverTab[103816]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:47
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:47
		// _ = "end of CoverTab[103812]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:47
		_go_fuzz_dep_.CoverTab[103813]++
														protocolType, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:49
			_go_fuzz_dep_.CoverTab[103817]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:50
			// _ = "end of CoverTab[103817]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:51
			_go_fuzz_dep_.CoverTab[103818]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:51
			// _ = "end of CoverTab[103818]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:51
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:51
		// _ = "end of CoverTab[103813]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:51
		_go_fuzz_dep_.CoverTab[103814]++

														r.Groups[groupId] = protocolType
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:53
		// _ = "end of CoverTab[103814]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:54
	// _ = "end of CoverTab[103804]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:54
	_go_fuzz_dep_.CoverTab[103805]++

													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:56
	// _ = "end of CoverTab[103805]"
}

func (r *ListGroupsResponse) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:59
	_go_fuzz_dep_.CoverTab[103819]++
													return 16
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:60
	// _ = "end of CoverTab[103819]"
}

func (r *ListGroupsResponse) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:63
	_go_fuzz_dep_.CoverTab[103820]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:64
	// _ = "end of CoverTab[103820]"
}

func (r *ListGroupsResponse) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:67
	_go_fuzz_dep_.CoverTab[103821]++
													return 0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:68
	// _ = "end of CoverTab[103821]"
}

func (r *ListGroupsResponse) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:71
	_go_fuzz_dep_.CoverTab[103822]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:72
	// _ = "end of CoverTab[103822]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/list_groups_response.go:73
var _ = _go_fuzz_dep_.CoverTab
