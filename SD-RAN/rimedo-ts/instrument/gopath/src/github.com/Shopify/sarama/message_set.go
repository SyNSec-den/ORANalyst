//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:1
)

type MessageBlock struct {
	Offset	int64
	Msg	*Message
}

// Messages convenience helper which returns either all the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:8
// messages that are wrapped in this block
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:10
func (msb *MessageBlock) Messages() []*MessageBlock {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:10
	_go_fuzz_dep_.CoverTab[104008]++
											if msb.Msg.Set != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:11
		_go_fuzz_dep_.CoverTab[104010]++
												return msb.Msg.Set.Messages
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:12
		// _ = "end of CoverTab[104010]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:13
		_go_fuzz_dep_.CoverTab[104011]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:13
		// _ = "end of CoverTab[104011]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:13
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:13
	// _ = "end of CoverTab[104008]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:13
	_go_fuzz_dep_.CoverTab[104009]++
											return []*MessageBlock{msb}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:14
	// _ = "end of CoverTab[104009]"
}

func (msb *MessageBlock) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:17
	_go_fuzz_dep_.CoverTab[104012]++
											pe.putInt64(msb.Offset)
											pe.push(&lengthField{})
											err := msb.Msg.encode(pe)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:21
		_go_fuzz_dep_.CoverTab[104014]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:22
		// _ = "end of CoverTab[104014]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:23
		_go_fuzz_dep_.CoverTab[104015]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:23
		// _ = "end of CoverTab[104015]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:23
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:23
	// _ = "end of CoverTab[104012]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:23
	_go_fuzz_dep_.CoverTab[104013]++
											return pe.pop()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:24
	// _ = "end of CoverTab[104013]"
}

func (msb *MessageBlock) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:27
	_go_fuzz_dep_.CoverTab[104016]++
											if msb.Offset, err = pd.getInt64(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:28
		_go_fuzz_dep_.CoverTab[104021]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:29
		// _ = "end of CoverTab[104021]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:30
		_go_fuzz_dep_.CoverTab[104022]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:30
		// _ = "end of CoverTab[104022]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:30
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:30
	// _ = "end of CoverTab[104016]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:30
	_go_fuzz_dep_.CoverTab[104017]++

											lengthDecoder := acquireLengthField()
											defer releaseLengthField(lengthDecoder)

											if err = pd.push(lengthDecoder); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:35
		_go_fuzz_dep_.CoverTab[104023]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:36
		// _ = "end of CoverTab[104023]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:37
		_go_fuzz_dep_.CoverTab[104024]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:37
		// _ = "end of CoverTab[104024]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:37
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:37
	// _ = "end of CoverTab[104017]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:37
	_go_fuzz_dep_.CoverTab[104018]++

											msb.Msg = new(Message)
											if err = msb.Msg.decode(pd); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:40
		_go_fuzz_dep_.CoverTab[104025]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:41
		// _ = "end of CoverTab[104025]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:42
		_go_fuzz_dep_.CoverTab[104026]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:42
		// _ = "end of CoverTab[104026]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:42
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:42
	// _ = "end of CoverTab[104018]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:42
	_go_fuzz_dep_.CoverTab[104019]++

											if err = pd.pop(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:44
		_go_fuzz_dep_.CoverTab[104027]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:45
		// _ = "end of CoverTab[104027]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:46
		_go_fuzz_dep_.CoverTab[104028]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:46
		// _ = "end of CoverTab[104028]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:46
	// _ = "end of CoverTab[104019]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:46
	_go_fuzz_dep_.CoverTab[104020]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:48
	// _ = "end of CoverTab[104020]"
}

type MessageSet struct {
	PartialTrailingMessage	bool	// whether the set on the wire contained an incomplete trailing MessageBlock
	OverflowMessage		bool	// whether the set on the wire contained an overflow message
	Messages		[]*MessageBlock
}

func (ms *MessageSet) encode(pe packetEncoder) error {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:57
	_go_fuzz_dep_.CoverTab[104029]++
											for i := range ms.Messages {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:58
		_go_fuzz_dep_.CoverTab[104031]++
												err := ms.Messages[i].encode(pe)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:60
			_go_fuzz_dep_.CoverTab[104032]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:61
			// _ = "end of CoverTab[104032]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:62
			_go_fuzz_dep_.CoverTab[104033]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:62
			// _ = "end of CoverTab[104033]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:62
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:62
		// _ = "end of CoverTab[104031]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:63
	// _ = "end of CoverTab[104029]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:63
	_go_fuzz_dep_.CoverTab[104030]++
											return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:64
	// _ = "end of CoverTab[104030]"
}

func (ms *MessageSet) decode(pd packetDecoder) (err error) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:67
	_go_fuzz_dep_.CoverTab[104034]++
											ms.Messages = nil

											for pd.remaining() > 0 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:70
		_go_fuzz_dep_.CoverTab[104036]++
												magic, err := magicValue(pd)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:72
			_go_fuzz_dep_.CoverTab[104039]++
													if err == ErrInsufficientData {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:73
				_go_fuzz_dep_.CoverTab[104041]++
														ms.PartialTrailingMessage = true
														return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:75
				// _ = "end of CoverTab[104041]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:76
				_go_fuzz_dep_.CoverTab[104042]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:76
				// _ = "end of CoverTab[104042]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:76
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:76
			// _ = "end of CoverTab[104039]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:76
			_go_fuzz_dep_.CoverTab[104040]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:77
			// _ = "end of CoverTab[104040]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:78
			_go_fuzz_dep_.CoverTab[104043]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:78
			// _ = "end of CoverTab[104043]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:78
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:78
		// _ = "end of CoverTab[104036]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:78
		_go_fuzz_dep_.CoverTab[104037]++

												if magic > 1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:80
			_go_fuzz_dep_.CoverTab[104044]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:81
			// _ = "end of CoverTab[104044]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:82
			_go_fuzz_dep_.CoverTab[104045]++
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:82
			// _ = "end of CoverTab[104045]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:82
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:82
		// _ = "end of CoverTab[104037]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:82
		_go_fuzz_dep_.CoverTab[104038]++

												msb := new(MessageBlock)
												err = msb.decode(pd)
												switch err {
		case nil:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:87
			_go_fuzz_dep_.CoverTab[104046]++
													ms.Messages = append(ms.Messages, msb)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:88
			// _ = "end of CoverTab[104046]"
		case ErrInsufficientData:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:89
			_go_fuzz_dep_.CoverTab[104047]++

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:92
			if msb.Offset == -1 {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:92
				_go_fuzz_dep_.CoverTab[104050]++

														ms.OverflowMessage = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:94
				// _ = "end of CoverTab[104050]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:95
				_go_fuzz_dep_.CoverTab[104051]++
														ms.PartialTrailingMessage = true
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:96
				// _ = "end of CoverTab[104051]"
			}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:97
			// _ = "end of CoverTab[104047]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:97
			_go_fuzz_dep_.CoverTab[104048]++
													return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:98
			// _ = "end of CoverTab[104048]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:99
				_go_fuzz_dep_.CoverTab[104049]++
														return err
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:100
			// _ = "end of CoverTab[104049]"
		}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:101
		// _ = "end of CoverTab[104038]"
	}
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:102
	// _ = "end of CoverTab[104034]"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:102
	_go_fuzz_dep_.CoverTab[104035]++

												return nil
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:104
	// _ = "end of CoverTab[104035]"
}

func (ms *MessageSet) addMessage(msg *Message) {
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:107
	_go_fuzz_dep_.CoverTab[104052]++
												block := new(MessageBlock)
												block.Msg = msg
												ms.Messages = append(ms.Messages, block)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:110
	// _ = "end of CoverTab[104052]"
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:111
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/message_set.go:111
var _ = _go_fuzz_dep_.CoverTab
