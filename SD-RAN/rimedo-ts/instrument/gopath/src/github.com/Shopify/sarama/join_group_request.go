//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:1
)

type GroupProtocol struct {
	Name		string
	Metadata	[]byte
}

func (p *GroupProtocol) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:8
	_go_fuzz_dep_.CoverTab[103559]++
												p.Name, err = pd.getString()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:10
		_go_fuzz_dep_.CoverTab[103561]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:11
		// _ = "end of CoverTab[103561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:12
		_go_fuzz_dep_.CoverTab[103562]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:12
		// _ = "end of CoverTab[103562]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:12
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:12
	// _ = "end of CoverTab[103559]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:12
	_go_fuzz_dep_.CoverTab[103560]++
												p.Metadata, err = pd.getBytes()
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:14
	// _ = "end of CoverTab[103560]"
}

func (p *GroupProtocol) encode(pe packetEncoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:17
	_go_fuzz_dep_.CoverTab[103563]++
												if err := pe.putString(p.Name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:18
		_go_fuzz_dep_.CoverTab[103566]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:19
		// _ = "end of CoverTab[103566]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:20
		_go_fuzz_dep_.CoverTab[103567]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:20
		// _ = "end of CoverTab[103567]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:20
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:20
	// _ = "end of CoverTab[103563]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:20
	_go_fuzz_dep_.CoverTab[103564]++
												if err := pe.putBytes(p.Metadata); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:21
		_go_fuzz_dep_.CoverTab[103568]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:22
		// _ = "end of CoverTab[103568]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:23
		_go_fuzz_dep_.CoverTab[103569]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:23
		// _ = "end of CoverTab[103569]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:23
	// _ = "end of CoverTab[103564]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:23
	_go_fuzz_dep_.CoverTab[103565]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:24
	// _ = "end of CoverTab[103565]"
}

type JoinGroupRequest struct {
	Version			int16
	GroupId			string
	SessionTimeout		int32
	RebalanceTimeout	int32
	MemberId		string
	ProtocolType		string
	GroupProtocols		map[string][]byte	// deprecated; use OrderedGroupProtocols
	OrderedGroupProtocols	[]*GroupProtocol
}

func (r *JoinGroupRequest) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:38
	_go_fuzz_dep_.CoverTab[103570]++
												if err := pe.putString(r.GroupId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:39
		_go_fuzz_dep_.CoverTab[103576]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:40
		// _ = "end of CoverTab[103576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:41
		_go_fuzz_dep_.CoverTab[103577]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:41
		// _ = "end of CoverTab[103577]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:41
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:41
	// _ = "end of CoverTab[103570]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:41
	_go_fuzz_dep_.CoverTab[103571]++
												pe.putInt32(r.SessionTimeout)
												if r.Version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:43
		_go_fuzz_dep_.CoverTab[103578]++
													pe.putInt32(r.RebalanceTimeout)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:44
		// _ = "end of CoverTab[103578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:45
		_go_fuzz_dep_.CoverTab[103579]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:45
		// _ = "end of CoverTab[103579]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:45
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:45
	// _ = "end of CoverTab[103571]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:45
	_go_fuzz_dep_.CoverTab[103572]++
												if err := pe.putString(r.MemberId); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:46
		_go_fuzz_dep_.CoverTab[103580]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:47
		// _ = "end of CoverTab[103580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:48
		_go_fuzz_dep_.CoverTab[103581]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:48
		// _ = "end of CoverTab[103581]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:48
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:48
	// _ = "end of CoverTab[103572]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:48
	_go_fuzz_dep_.CoverTab[103573]++
												if err := pe.putString(r.ProtocolType); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:49
		_go_fuzz_dep_.CoverTab[103582]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:50
		// _ = "end of CoverTab[103582]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:51
		_go_fuzz_dep_.CoverTab[103583]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:51
		// _ = "end of CoverTab[103583]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:51
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:51
	// _ = "end of CoverTab[103573]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:51
	_go_fuzz_dep_.CoverTab[103574]++

												if len(r.GroupProtocols) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:53
		_go_fuzz_dep_.CoverTab[103584]++
													if len(r.OrderedGroupProtocols) > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:54
			_go_fuzz_dep_.CoverTab[103587]++
														return PacketDecodingError{"cannot specify both GroupProtocols and OrderedGroupProtocols on JoinGroupRequest"}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:55
			// _ = "end of CoverTab[103587]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:56
			_go_fuzz_dep_.CoverTab[103588]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:56
			// _ = "end of CoverTab[103588]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:56
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:56
		// _ = "end of CoverTab[103584]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:56
		_go_fuzz_dep_.CoverTab[103585]++

													if err := pe.putArrayLength(len(r.GroupProtocols)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:58
			_go_fuzz_dep_.CoverTab[103589]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:59
			// _ = "end of CoverTab[103589]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:60
			_go_fuzz_dep_.CoverTab[103590]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:60
			// _ = "end of CoverTab[103590]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:60
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:60
		// _ = "end of CoverTab[103585]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:60
		_go_fuzz_dep_.CoverTab[103586]++
													for name, metadata := range r.GroupProtocols {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:61
			_go_fuzz_dep_.CoverTab[103591]++
														if err := pe.putString(name); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:62
				_go_fuzz_dep_.CoverTab[103593]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:63
				// _ = "end of CoverTab[103593]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:64
				_go_fuzz_dep_.CoverTab[103594]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:64
				// _ = "end of CoverTab[103594]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:64
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:64
			// _ = "end of CoverTab[103591]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:64
			_go_fuzz_dep_.CoverTab[103592]++
														if err := pe.putBytes(metadata); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:65
				_go_fuzz_dep_.CoverTab[103595]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:66
				// _ = "end of CoverTab[103595]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:67
				_go_fuzz_dep_.CoverTab[103596]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:67
				// _ = "end of CoverTab[103596]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:67
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:67
			// _ = "end of CoverTab[103592]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:68
		// _ = "end of CoverTab[103586]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:69
		_go_fuzz_dep_.CoverTab[103597]++
													if err := pe.putArrayLength(len(r.OrderedGroupProtocols)); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:70
			_go_fuzz_dep_.CoverTab[103599]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:71
			// _ = "end of CoverTab[103599]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:72
			_go_fuzz_dep_.CoverTab[103600]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:72
			// _ = "end of CoverTab[103600]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:72
		// _ = "end of CoverTab[103597]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:72
		_go_fuzz_dep_.CoverTab[103598]++
													for _, protocol := range r.OrderedGroupProtocols {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:73
			_go_fuzz_dep_.CoverTab[103601]++
														if err := protocol.encode(pe); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:74
				_go_fuzz_dep_.CoverTab[103602]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:75
				// _ = "end of CoverTab[103602]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:76
				_go_fuzz_dep_.CoverTab[103603]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:76
				// _ = "end of CoverTab[103603]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:76
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:76
			// _ = "end of CoverTab[103601]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:77
		// _ = "end of CoverTab[103598]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:78
	// _ = "end of CoverTab[103574]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:78
	_go_fuzz_dep_.CoverTab[103575]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:80
	// _ = "end of CoverTab[103575]"
}

func (r *JoinGroupRequest) decode(pd packetDecoder, version int16) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:83
	_go_fuzz_dep_.CoverTab[103604]++
												r.Version = version

												if r.GroupId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:86
		_go_fuzz_dep_.CoverTab[103613]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:87
		// _ = "end of CoverTab[103613]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:88
		_go_fuzz_dep_.CoverTab[103614]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:88
		// _ = "end of CoverTab[103614]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:88
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:88
	// _ = "end of CoverTab[103604]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:88
	_go_fuzz_dep_.CoverTab[103605]++

												if r.SessionTimeout, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:90
		_go_fuzz_dep_.CoverTab[103615]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:91
		// _ = "end of CoverTab[103615]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:92
		_go_fuzz_dep_.CoverTab[103616]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:92
		// _ = "end of CoverTab[103616]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:92
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:92
	// _ = "end of CoverTab[103605]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:92
	_go_fuzz_dep_.CoverTab[103606]++

												if version >= 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:94
		_go_fuzz_dep_.CoverTab[103617]++
													if r.RebalanceTimeout, err = pd.getInt32(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:95
			_go_fuzz_dep_.CoverTab[103618]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:96
			// _ = "end of CoverTab[103618]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:97
			_go_fuzz_dep_.CoverTab[103619]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:97
			// _ = "end of CoverTab[103619]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:97
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:97
		// _ = "end of CoverTab[103617]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:98
		_go_fuzz_dep_.CoverTab[103620]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:98
		// _ = "end of CoverTab[103620]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:98
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:98
	// _ = "end of CoverTab[103606]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:98
	_go_fuzz_dep_.CoverTab[103607]++

												if r.MemberId, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:100
		_go_fuzz_dep_.CoverTab[103621]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:101
		// _ = "end of CoverTab[103621]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:102
		_go_fuzz_dep_.CoverTab[103622]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:102
		// _ = "end of CoverTab[103622]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:102
	// _ = "end of CoverTab[103607]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:102
	_go_fuzz_dep_.CoverTab[103608]++

												if r.ProtocolType, err = pd.getString(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:104
		_go_fuzz_dep_.CoverTab[103623]++
													return
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:105
		// _ = "end of CoverTab[103623]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:106
		_go_fuzz_dep_.CoverTab[103624]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:106
		// _ = "end of CoverTab[103624]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:106
	// _ = "end of CoverTab[103608]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:106
	_go_fuzz_dep_.CoverTab[103609]++

												n, err := pd.getArrayLength()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:109
		_go_fuzz_dep_.CoverTab[103625]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:110
		// _ = "end of CoverTab[103625]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:111
		_go_fuzz_dep_.CoverTab[103626]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:111
		// _ = "end of CoverTab[103626]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:111
	// _ = "end of CoverTab[103609]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:111
	_go_fuzz_dep_.CoverTab[103610]++
												if n == 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:112
		_go_fuzz_dep_.CoverTab[103627]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:113
		// _ = "end of CoverTab[103627]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:114
		_go_fuzz_dep_.CoverTab[103628]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:114
		// _ = "end of CoverTab[103628]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:114
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:114
	// _ = "end of CoverTab[103610]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:114
	_go_fuzz_dep_.CoverTab[103611]++

												r.GroupProtocols = make(map[string][]byte)
												for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:117
		_go_fuzz_dep_.CoverTab[103629]++
													protocol := &GroupProtocol{}
													if err := protocol.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:119
			_go_fuzz_dep_.CoverTab[103631]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:120
			// _ = "end of CoverTab[103631]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:121
			_go_fuzz_dep_.CoverTab[103632]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:121
			// _ = "end of CoverTab[103632]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:121
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:121
		// _ = "end of CoverTab[103629]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:121
		_go_fuzz_dep_.CoverTab[103630]++
													r.GroupProtocols[protocol.Name] = protocol.Metadata
													r.OrderedGroupProtocols = append(r.OrderedGroupProtocols, protocol)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:123
		// _ = "end of CoverTab[103630]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:124
	// _ = "end of CoverTab[103611]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:124
	_go_fuzz_dep_.CoverTab[103612]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:126
	// _ = "end of CoverTab[103612]"
}

func (r *JoinGroupRequest) key() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:129
	_go_fuzz_dep_.CoverTab[103633]++
												return 11
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:130
	// _ = "end of CoverTab[103633]"
}

func (r *JoinGroupRequest) version() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:133
	_go_fuzz_dep_.CoverTab[103634]++
												return r.Version
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:134
	// _ = "end of CoverTab[103634]"
}

func (r *JoinGroupRequest) headerVersion() int16 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:137
	_go_fuzz_dep_.CoverTab[103635]++
												return 1
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:138
	// _ = "end of CoverTab[103635]"
}

func (r *JoinGroupRequest) requiredVersion() KafkaVersion {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:141
	_go_fuzz_dep_.CoverTab[103636]++
												switch r.Version {
	case 2:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:143
		_go_fuzz_dep_.CoverTab[103637]++
													return V0_11_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:144
		// _ = "end of CoverTab[103637]"
	case 1:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:145
		_go_fuzz_dep_.CoverTab[103638]++
													return V0_10_1_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:146
		// _ = "end of CoverTab[103638]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:147
		_go_fuzz_dep_.CoverTab[103639]++
													return V0_9_0_0
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:148
		// _ = "end of CoverTab[103639]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:149
	// _ = "end of CoverTab[103636]"
}

func (r *JoinGroupRequest) AddGroupProtocol(name string, metadata []byte) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:152
	_go_fuzz_dep_.CoverTab[103640]++
												r.OrderedGroupProtocols = append(r.OrderedGroupProtocols, &GroupProtocol{
		Name:		name,
		Metadata:	metadata,
	})
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:156
	// _ = "end of CoverTab[103640]"
}

func (r *JoinGroupRequest) AddGroupProtocolMetadata(name string, metadata *ConsumerGroupMemberMetadata) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:159
	_go_fuzz_dep_.CoverTab[103641]++
												bin, err := encode(metadata, nil)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:161
		_go_fuzz_dep_.CoverTab[103643]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:162
		// _ = "end of CoverTab[103643]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:163
		_go_fuzz_dep_.CoverTab[103644]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:163
		// _ = "end of CoverTab[103644]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:163
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:163
	// _ = "end of CoverTab[103641]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:163
	_go_fuzz_dep_.CoverTab[103642]++

												r.AddGroupProtocol(name, bin)
												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:166
	// _ = "end of CoverTab[103642]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:167
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/join_group_request.go:167
var _ = _go_fuzz_dep_.CoverTab
