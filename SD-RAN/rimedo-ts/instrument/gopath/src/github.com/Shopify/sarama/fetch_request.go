//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:1
)

type fetchRequestBlock struct {
	Version			int16
	currentLeaderEpoch	int32
	fetchOffset		int64
	logStartOffset		int64
	maxBytes		int32
}

func (b *fetchRequestBlock) encode(pe packetEncoder, version int16) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:11
	_go_fuzz_dep_.CoverTab[102849]++
												b.Version = version
												if b.Version >= 9 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:13
		_go_fuzz_dep_.CoverTab[102852]++
													pe.putInt32(b.currentLeaderEpoch)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:14
		// _ = "end of CoverTab[102852]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:15
		_go_fuzz_dep_.CoverTab[102853]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:15
		// _ = "end of CoverTab[102853]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:15
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:15
	// _ = "end of CoverTab[102849]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:15
	_go_fuzz_dep_.CoverTab[102850]++
												pe.putInt64(b.fetchOffset)
												if b.Version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:17
		_go_fuzz_dep_.CoverTab[102854]++
													pe.putInt64(b.logStartOffset)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:18
		// _ = "end of CoverTab[102854]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:19
		_go_fuzz_dep_.CoverTab[102855]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:19
		// _ = "end of CoverTab[102855]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:19
	// _ = "end of CoverTab[102850]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:19
	_go_fuzz_dep_.CoverTab[102851]++
												pe.putInt32(b.maxBytes)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:21
	// _ = "end of CoverTab[102851]"
}

func (b *fetchRequestBlock) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:24
	_go_fuzz_dep_.CoverTab[102856]++
												b.Version = version
												if b.Version >= 9 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:26
		_go_fuzz_dep_.CoverTab[102861]++
													if b.currentLeaderEpoch, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:27
			_go_fuzz_dep_.CoverTab[102862]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:28
			// _ = "end of CoverTab[102862]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:29
			_go_fuzz_dep_.CoverTab[102863]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:29
			// _ = "end of CoverTab[102863]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:29
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:29
		// _ = "end of CoverTab[102861]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:30
		_go_fuzz_dep_.CoverTab[102864]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:30
		// _ = "end of CoverTab[102864]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:30
	// _ = "end of CoverTab[102856]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:30
	_go_fuzz_dep_.CoverTab[102857]++
												if b.fetchOffset, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:31
		_go_fuzz_dep_.CoverTab[102865]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:32
		// _ = "end of CoverTab[102865]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:33
		_go_fuzz_dep_.CoverTab[102866]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:33
		// _ = "end of CoverTab[102866]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:33
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:33
	// _ = "end of CoverTab[102857]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:33
	_go_fuzz_dep_.CoverTab[102858]++
												if b.Version >= 5 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:34
		_go_fuzz_dep_.CoverTab[102867]++
													if b.logStartOffset, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:35
			_go_fuzz_dep_.CoverTab[102868]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:36
			// _ = "end of CoverTab[102868]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:37
			_go_fuzz_dep_.CoverTab[102869]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:37
			// _ = "end of CoverTab[102869]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:37
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:37
		// _ = "end of CoverTab[102867]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:38
		_go_fuzz_dep_.CoverTab[102870]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:38
		// _ = "end of CoverTab[102870]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:38
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:38
	// _ = "end of CoverTab[102858]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:38
	_go_fuzz_dep_.CoverTab[102859]++
												if b.maxBytes, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:39
		_go_fuzz_dep_.CoverTab[102871]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:40
		// _ = "end of CoverTab[102871]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:41
		_go_fuzz_dep_.CoverTab[102872]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:41
		// _ = "end of CoverTab[102872]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:41
	// _ = "end of CoverTab[102859]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:41
	_go_fuzz_dep_.CoverTab[102860]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:42
	// _ = "end of CoverTab[102860]"
}

// FetchRequest (API key 1) will fetch Kafka messages. Version 3 introduced the MaxBytes field. See
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:45
// https://issues.apache.org/jira/browse/KAFKA-2063 for a discussion of the issues leading up to that.  The KIP is at
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:45
// https://cwiki.apache.org/confluence/display/KAFKA/KIP-74%3A+Add+Fetch+Response+Size+Limit+in+Bytes
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:48
type FetchRequest struct {
	MaxWaitTime	int32
	MinBytes	int32
	MaxBytes	int32
	Version		int16
	Isolation	IsolationLevel
	SessionID	int32
	SessionEpoch	int32
	blocks		map[string]map[int32]*fetchRequestBlock
	forgotten	map[string][]int32
	RackID		string
}

type IsolationLevel int8

const (
	ReadUncommitted	IsolationLevel	= iota
	ReadCommitted
)

func (r *FetchRequest) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:68
	_go_fuzz_dep_.CoverTab[102873]++
												pe.putInt32(-1)
												pe.putInt32(r.MaxWaitTime)
												pe.putInt32(r.MinBytes)
												if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:72
		_go_fuzz_dep_.CoverTab[102881]++
													pe.putInt32(r.MaxBytes)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:73
		// _ = "end of CoverTab[102881]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:74
		_go_fuzz_dep_.CoverTab[102882]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:74
		// _ = "end of CoverTab[102882]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:74
	// _ = "end of CoverTab[102873]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:74
	_go_fuzz_dep_.CoverTab[102874]++
												if r.Version >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:75
		_go_fuzz_dep_.CoverTab[102883]++
													pe.putInt8(int8(r.Isolation))
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:76
		// _ = "end of CoverTab[102883]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:77
		_go_fuzz_dep_.CoverTab[102884]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:77
		// _ = "end of CoverTab[102884]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:77
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:77
	// _ = "end of CoverTab[102874]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:77
	_go_fuzz_dep_.CoverTab[102875]++
												if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:78
		_go_fuzz_dep_.CoverTab[102885]++
													pe.putInt32(r.SessionID)
													pe.putInt32(r.SessionEpoch)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:80
		// _ = "end of CoverTab[102885]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:81
		_go_fuzz_dep_.CoverTab[102886]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:81
		// _ = "end of CoverTab[102886]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:81
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:81
	// _ = "end of CoverTab[102875]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:81
	_go_fuzz_dep_.CoverTab[102876]++
												err = pe.putArrayLength(len(r.blocks))
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:83
		_go_fuzz_dep_.CoverTab[102887]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:84
		// _ = "end of CoverTab[102887]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:85
		_go_fuzz_dep_.CoverTab[102888]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:85
		// _ = "end of CoverTab[102888]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:85
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:85
	// _ = "end of CoverTab[102876]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:85
	_go_fuzz_dep_.CoverTab[102877]++
												for topic, blocks := range r.blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:86
		_go_fuzz_dep_.CoverTab[102889]++
													err = pe.putString(topic)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:88
			_go_fuzz_dep_.CoverTab[102892]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:89
			// _ = "end of CoverTab[102892]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:90
			_go_fuzz_dep_.CoverTab[102893]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:90
			// _ = "end of CoverTab[102893]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:90
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:90
		// _ = "end of CoverTab[102889]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:90
		_go_fuzz_dep_.CoverTab[102890]++
													err = pe.putArrayLength(len(blocks))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:92
			_go_fuzz_dep_.CoverTab[102894]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:93
			// _ = "end of CoverTab[102894]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:94
			_go_fuzz_dep_.CoverTab[102895]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:94
			// _ = "end of CoverTab[102895]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:94
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:94
		// _ = "end of CoverTab[102890]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:94
		_go_fuzz_dep_.CoverTab[102891]++
													for partition, block := range blocks {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:95
			_go_fuzz_dep_.CoverTab[102896]++
														pe.putInt32(partition)
														err = block.encode(pe, r.Version)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:98
				_go_fuzz_dep_.CoverTab[102897]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:99
				// _ = "end of CoverTab[102897]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:100
				_go_fuzz_dep_.CoverTab[102898]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:100
				// _ = "end of CoverTab[102898]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:100
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:100
			// _ = "end of CoverTab[102896]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:101
		// _ = "end of CoverTab[102891]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:102
	// _ = "end of CoverTab[102877]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:102
	_go_fuzz_dep_.CoverTab[102878]++
												if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:103
		_go_fuzz_dep_.CoverTab[102899]++
													err = pe.putArrayLength(len(r.forgotten))
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:105
			_go_fuzz_dep_.CoverTab[102901]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:106
			// _ = "end of CoverTab[102901]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:107
			_go_fuzz_dep_.CoverTab[102902]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:107
			// _ = "end of CoverTab[102902]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:107
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:107
		// _ = "end of CoverTab[102899]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:107
		_go_fuzz_dep_.CoverTab[102900]++
													for topic, partitions := range r.forgotten {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:108
			_go_fuzz_dep_.CoverTab[102903]++
														err = pe.putString(topic)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:110
				_go_fuzz_dep_.CoverTab[102906]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:111
				// _ = "end of CoverTab[102906]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:112
				_go_fuzz_dep_.CoverTab[102907]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:112
				// _ = "end of CoverTab[102907]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:112
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:112
			// _ = "end of CoverTab[102903]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:112
			_go_fuzz_dep_.CoverTab[102904]++
														err = pe.putArrayLength(len(partitions))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:114
				_go_fuzz_dep_.CoverTab[102908]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:115
				// _ = "end of CoverTab[102908]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:116
				_go_fuzz_dep_.CoverTab[102909]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:116
				// _ = "end of CoverTab[102909]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:116
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:116
			// _ = "end of CoverTab[102904]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:116
			_go_fuzz_dep_.CoverTab[102905]++
														for _, partition := range partitions {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:117
				_go_fuzz_dep_.CoverTab[102910]++
															pe.putInt32(partition)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:118
				// _ = "end of CoverTab[102910]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:119
			// _ = "end of CoverTab[102905]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:120
		// _ = "end of CoverTab[102900]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:121
		_go_fuzz_dep_.CoverTab[102911]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:121
		// _ = "end of CoverTab[102911]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:121
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:121
	// _ = "end of CoverTab[102878]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:121
	_go_fuzz_dep_.CoverTab[102879]++
												if r.Version >= 11 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:122
		_go_fuzz_dep_.CoverTab[102912]++
													err = pe.putString(r.RackID)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:124
			_go_fuzz_dep_.CoverTab[102913]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:125
			// _ = "end of CoverTab[102913]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:126
			_go_fuzz_dep_.CoverTab[102914]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:126
			// _ = "end of CoverTab[102914]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:126
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:126
		// _ = "end of CoverTab[102912]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:127
		_go_fuzz_dep_.CoverTab[102915]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:127
		// _ = "end of CoverTab[102915]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:127
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:127
	// _ = "end of CoverTab[102879]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:127
	_go_fuzz_dep_.CoverTab[102880]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:129
	// _ = "end of CoverTab[102880]"
}

func (r *FetchRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:132
	_go_fuzz_dep_.CoverTab[102916]++
												r.Version = version

												if _, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:135
		_go_fuzz_dep_.CoverTab[102928]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:136
		// _ = "end of CoverTab[102928]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:137
		_go_fuzz_dep_.CoverTab[102929]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:137
		// _ = "end of CoverTab[102929]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:137
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:137
	// _ = "end of CoverTab[102916]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:137
	_go_fuzz_dep_.CoverTab[102917]++
												if r.MaxWaitTime, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:138
		_go_fuzz_dep_.CoverTab[102930]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:139
		// _ = "end of CoverTab[102930]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:140
		_go_fuzz_dep_.CoverTab[102931]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:140
		// _ = "end of CoverTab[102931]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:140
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:140
	// _ = "end of CoverTab[102917]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:140
	_go_fuzz_dep_.CoverTab[102918]++
												if r.MinBytes, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:141
		_go_fuzz_dep_.CoverTab[102932]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:142
		// _ = "end of CoverTab[102932]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:143
		_go_fuzz_dep_.CoverTab[102933]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:143
		// _ = "end of CoverTab[102933]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:143
	// _ = "end of CoverTab[102918]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:143
	_go_fuzz_dep_.CoverTab[102919]++
												if r.Version >= 3 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:144
		_go_fuzz_dep_.CoverTab[102934]++
													if r.MaxBytes, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:145
			_go_fuzz_dep_.CoverTab[102935]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:146
			// _ = "end of CoverTab[102935]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:147
			_go_fuzz_dep_.CoverTab[102936]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:147
			// _ = "end of CoverTab[102936]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:147
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:147
		// _ = "end of CoverTab[102934]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:148
		_go_fuzz_dep_.CoverTab[102937]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:148
		// _ = "end of CoverTab[102937]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:148
	// _ = "end of CoverTab[102919]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:148
	_go_fuzz_dep_.CoverTab[102920]++
												if r.Version >= 4 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:149
		_go_fuzz_dep_.CoverTab[102938]++
													isolation, err := pd.getInt8()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:151
			_go_fuzz_dep_.CoverTab[102940]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:152
			// _ = "end of CoverTab[102940]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:153
			_go_fuzz_dep_.CoverTab[102941]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:153
			// _ = "end of CoverTab[102941]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:153
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:153
		// _ = "end of CoverTab[102938]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:153
		_go_fuzz_dep_.CoverTab[102939]++
													r.Isolation = IsolationLevel(isolation)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:154
		// _ = "end of CoverTab[102939]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:155
		_go_fuzz_dep_.CoverTab[102942]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:155
		// _ = "end of CoverTab[102942]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:155
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:155
	// _ = "end of CoverTab[102920]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:155
	_go_fuzz_dep_.CoverTab[102921]++
												if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:156
		_go_fuzz_dep_.CoverTab[102943]++
													r.SessionID, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:158
			_go_fuzz_dep_.CoverTab[102945]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:159
			// _ = "end of CoverTab[102945]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:160
			_go_fuzz_dep_.CoverTab[102946]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:160
			// _ = "end of CoverTab[102946]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:160
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:160
		// _ = "end of CoverTab[102943]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:160
		_go_fuzz_dep_.CoverTab[102944]++
													r.SessionEpoch, err = pd.getInt32()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:162
			_go_fuzz_dep_.CoverTab[102947]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:163
			// _ = "end of CoverTab[102947]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:164
			_go_fuzz_dep_.CoverTab[102948]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:164
			// _ = "end of CoverTab[102948]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:164
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:164
		// _ = "end of CoverTab[102944]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:165
		_go_fuzz_dep_.CoverTab[102949]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:165
		// _ = "end of CoverTab[102949]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:165
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:165
	// _ = "end of CoverTab[102921]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:165
	_go_fuzz_dep_.CoverTab[102922]++
												topicCount, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:167
		_go_fuzz_dep_.CoverTab[102950]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:168
		// _ = "end of CoverTab[102950]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:169
		_go_fuzz_dep_.CoverTab[102951]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:169
		// _ = "end of CoverTab[102951]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:169
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:169
	// _ = "end of CoverTab[102922]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:169
	_go_fuzz_dep_.CoverTab[102923]++
												if topicCount == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:170
		_go_fuzz_dep_.CoverTab[102952]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:171
		// _ = "end of CoverTab[102952]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:172
		_go_fuzz_dep_.CoverTab[102953]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:172
		// _ = "end of CoverTab[102953]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:172
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:172
	// _ = "end of CoverTab[102923]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:172
	_go_fuzz_dep_.CoverTab[102924]++
												r.blocks = make(map[string]map[int32]*fetchRequestBlock)
												for i := 0; i < topicCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:174
		_go_fuzz_dep_.CoverTab[102954]++
													topic, err := pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:176
			_go_fuzz_dep_.CoverTab[102957]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:177
			// _ = "end of CoverTab[102957]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:178
			_go_fuzz_dep_.CoverTab[102958]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:178
			// _ = "end of CoverTab[102958]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:178
		// _ = "end of CoverTab[102954]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:178
		_go_fuzz_dep_.CoverTab[102955]++
													partitionCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:180
			_go_fuzz_dep_.CoverTab[102959]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:181
			// _ = "end of CoverTab[102959]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:182
			_go_fuzz_dep_.CoverTab[102960]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:182
			// _ = "end of CoverTab[102960]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:182
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:182
		// _ = "end of CoverTab[102955]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:182
		_go_fuzz_dep_.CoverTab[102956]++
													r.blocks[topic] = make(map[int32]*fetchRequestBlock)
													for j := 0; j < partitionCount; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:184
			_go_fuzz_dep_.CoverTab[102961]++
														partition, err := pd.getInt32()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:186
				_go_fuzz_dep_.CoverTab[102964]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:187
				// _ = "end of CoverTab[102964]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:188
				_go_fuzz_dep_.CoverTab[102965]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:188
				// _ = "end of CoverTab[102965]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:188
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:188
			// _ = "end of CoverTab[102961]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:188
			_go_fuzz_dep_.CoverTab[102962]++
														fetchBlock := &fetchRequestBlock{}
														if err = fetchBlock.decode(pd, r.Version); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:190
				_go_fuzz_dep_.CoverTab[102966]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:191
				// _ = "end of CoverTab[102966]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:192
				_go_fuzz_dep_.CoverTab[102967]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:192
				// _ = "end of CoverTab[102967]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:192
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:192
			// _ = "end of CoverTab[102962]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:192
			_go_fuzz_dep_.CoverTab[102963]++
														r.blocks[topic][partition] = fetchBlock
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:193
			// _ = "end of CoverTab[102963]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:194
		// _ = "end of CoverTab[102956]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:195
	// _ = "end of CoverTab[102924]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:195
	_go_fuzz_dep_.CoverTab[102925]++

												if r.Version >= 7 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:197
		_go_fuzz_dep_.CoverTab[102968]++
													forgottenCount, err := pd.getArrayLength()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:199
			_go_fuzz_dep_.CoverTab[102970]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:200
			// _ = "end of CoverTab[102970]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:201
			_go_fuzz_dep_.CoverTab[102971]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:201
			// _ = "end of CoverTab[102971]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:201
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:201
		// _ = "end of CoverTab[102968]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:201
		_go_fuzz_dep_.CoverTab[102969]++
													r.forgotten = make(map[string][]int32)
													for i := 0; i < forgottenCount; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:203
			_go_fuzz_dep_.CoverTab[102972]++
														topic, err := pd.getString()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:205
				_go_fuzz_dep_.CoverTab[102975]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:206
				// _ = "end of CoverTab[102975]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:207
				_go_fuzz_dep_.CoverTab[102976]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:207
				// _ = "end of CoverTab[102976]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:207
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:207
			// _ = "end of CoverTab[102972]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:207
			_go_fuzz_dep_.CoverTab[102973]++
														partitionCount, err := pd.getArrayLength()
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:209
				_go_fuzz_dep_.CoverTab[102977]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:210
				// _ = "end of CoverTab[102977]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:211
				_go_fuzz_dep_.CoverTab[102978]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:211
				// _ = "end of CoverTab[102978]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:211
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:211
			// _ = "end of CoverTab[102973]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:211
			_go_fuzz_dep_.CoverTab[102974]++
														r.forgotten[topic] = make([]int32, partitionCount)

														for j := 0; j < partitionCount; j++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:214
				_go_fuzz_dep_.CoverTab[102979]++
															partition, err := pd.getInt32()
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:216
					_go_fuzz_dep_.CoverTab[102981]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:217
					// _ = "end of CoverTab[102981]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:218
					_go_fuzz_dep_.CoverTab[102982]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:218
					// _ = "end of CoverTab[102982]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:218
				}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:218
				// _ = "end of CoverTab[102979]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:218
				_go_fuzz_dep_.CoverTab[102980]++
															r.forgotten[topic][j] = partition
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:219
				// _ = "end of CoverTab[102980]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:220
			// _ = "end of CoverTab[102974]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:221
		// _ = "end of CoverTab[102969]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:222
		_go_fuzz_dep_.CoverTab[102983]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:222
		// _ = "end of CoverTab[102983]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:222
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:222
	// _ = "end of CoverTab[102925]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:222
	_go_fuzz_dep_.CoverTab[102926]++

												if r.Version >= 11 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:224
		_go_fuzz_dep_.CoverTab[102984]++
													r.RackID, err = pd.getString()
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:226
			_go_fuzz_dep_.CoverTab[102985]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:227
			// _ = "end of CoverTab[102985]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:228
			_go_fuzz_dep_.CoverTab[102986]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:228
			// _ = "end of CoverTab[102986]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:228
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:228
		// _ = "end of CoverTab[102984]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:229
		_go_fuzz_dep_.CoverTab[102987]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:229
		// _ = "end of CoverTab[102987]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:229
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:229
	// _ = "end of CoverTab[102926]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:229
	_go_fuzz_dep_.CoverTab[102927]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:231
	// _ = "end of CoverTab[102927]"
}

func (r *FetchRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:234
	_go_fuzz_dep_.CoverTab[102988]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:235
	// _ = "end of CoverTab[102988]"
}

func (r *FetchRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:238
	_go_fuzz_dep_.CoverTab[102989]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:239
	// _ = "end of CoverTab[102989]"
}

func (r *FetchRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:242
	_go_fuzz_dep_.CoverTab[102990]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:243
	// _ = "end of CoverTab[102990]"
}

func (r *FetchRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:246
	_go_fuzz_dep_.CoverTab[102991]++
												switch r.Version {
	case 0:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:248
		_go_fuzz_dep_.CoverTab[102992]++
													return MinVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:249
		// _ = "end of CoverTab[102992]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:250
		_go_fuzz_dep_.CoverTab[102993]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:251
		// _ = "end of CoverTab[102993]"
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:252
		_go_fuzz_dep_.CoverTab[102994]++
													return V0_10_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:253
		// _ = "end of CoverTab[102994]"
	case 3:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:254
		_go_fuzz_dep_.CoverTab[102995]++
													return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:255
		// _ = "end of CoverTab[102995]"
	case 4, 5:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:256
		_go_fuzz_dep_.CoverTab[102996]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:257
		// _ = "end of CoverTab[102996]"
	case 6:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:258
		_go_fuzz_dep_.CoverTab[102997]++
													return V1_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:259
		// _ = "end of CoverTab[102997]"
	case 7:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:260
		_go_fuzz_dep_.CoverTab[102998]++
													return V1_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:261
		// _ = "end of CoverTab[102998]"
	case 8:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:262
		_go_fuzz_dep_.CoverTab[102999]++
													return V2_0_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:263
		// _ = "end of CoverTab[102999]"
	case 9, 10:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:264
		_go_fuzz_dep_.CoverTab[103000]++
													return V2_1_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:265
		// _ = "end of CoverTab[103000]"
	case 11:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:266
		_go_fuzz_dep_.CoverTab[103001]++
													return V2_3_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:267
		// _ = "end of CoverTab[103001]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:268
		_go_fuzz_dep_.CoverTab[103002]++
													return MaxVersion
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:269
		// _ = "end of CoverTab[103002]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:270
	// _ = "end of CoverTab[102991]"
}

func (r *FetchRequest) AddBlock(topic string, partitionID int32, fetchOffset int64, maxBytes int32) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:273
	_go_fuzz_dep_.CoverTab[103003]++
												if r.blocks == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:274
		_go_fuzz_dep_.CoverTab[103008]++
													r.blocks = make(map[string]map[int32]*fetchRequestBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:275
		// _ = "end of CoverTab[103008]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:276
		_go_fuzz_dep_.CoverTab[103009]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:276
		// _ = "end of CoverTab[103009]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:276
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:276
	// _ = "end of CoverTab[103003]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:276
	_go_fuzz_dep_.CoverTab[103004]++

												if r.Version >= 7 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:278
		_go_fuzz_dep_.CoverTab[103010]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:278
		return r.forgotten == nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:278
		// _ = "end of CoverTab[103010]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:278
	}() {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:278
		_go_fuzz_dep_.CoverTab[103011]++
													r.forgotten = make(map[string][]int32)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:279
		// _ = "end of CoverTab[103011]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:280
		_go_fuzz_dep_.CoverTab[103012]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:280
		// _ = "end of CoverTab[103012]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:280
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:280
	// _ = "end of CoverTab[103004]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:280
	_go_fuzz_dep_.CoverTab[103005]++

												if r.blocks[topic] == nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:282
		_go_fuzz_dep_.CoverTab[103013]++
													r.blocks[topic] = make(map[int32]*fetchRequestBlock)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:283
		// _ = "end of CoverTab[103013]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:284
		_go_fuzz_dep_.CoverTab[103014]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:284
		// _ = "end of CoverTab[103014]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:284
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:284
	// _ = "end of CoverTab[103005]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:284
	_go_fuzz_dep_.CoverTab[103006]++

												tmp := new(fetchRequestBlock)
												tmp.Version = r.Version
												tmp.maxBytes = maxBytes
												tmp.fetchOffset = fetchOffset
												if r.Version >= 9 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:290
		_go_fuzz_dep_.CoverTab[103015]++
													tmp.currentLeaderEpoch = int32(-1)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:291
		// _ = "end of CoverTab[103015]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:292
		_go_fuzz_dep_.CoverTab[103016]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:292
		// _ = "end of CoverTab[103016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:292
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:292
	// _ = "end of CoverTab[103006]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:292
	_go_fuzz_dep_.CoverTab[103007]++

												r.blocks[topic][partitionID] = tmp
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:294
	// _ = "end of CoverTab[103007]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:295
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/fetch_request.go:295
var _ = _go_fuzz_dep_.CoverTab
