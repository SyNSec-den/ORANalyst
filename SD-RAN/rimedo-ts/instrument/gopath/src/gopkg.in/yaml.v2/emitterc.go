//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1
)

import (
	"bytes"
	"fmt"
)

// Flush the buffer if needed.
func flush(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:9
	_go_fuzz_dep_.CoverTab[124888]++
										if emitter.buffer_pos+5 >= len(emitter.buffer) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:10
		_go_fuzz_dep_.CoverTab[124890]++
											return yaml_emitter_flush(emitter)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:11
		// _ = "end of CoverTab[124890]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:12
		_go_fuzz_dep_.CoverTab[124891]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:12
		// _ = "end of CoverTab[124891]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:12
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:12
	// _ = "end of CoverTab[124888]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:12
	_go_fuzz_dep_.CoverTab[124889]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:13
	// _ = "end of CoverTab[124889]"
}

// Put a character to the output buffer.
func put(emitter *yaml_emitter_t, value byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:17
	_go_fuzz_dep_.CoverTab[124892]++
										if emitter.buffer_pos+5 >= len(emitter.buffer) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:18
		_go_fuzz_dep_.CoverTab[124894]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:18
		return !yaml_emitter_flush(emitter)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:18
		// _ = "end of CoverTab[124894]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:18
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:18
		_go_fuzz_dep_.CoverTab[124895]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:19
		// _ = "end of CoverTab[124895]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:20
		_go_fuzz_dep_.CoverTab[124896]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:20
		// _ = "end of CoverTab[124896]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:20
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:20
	// _ = "end of CoverTab[124892]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:20
	_go_fuzz_dep_.CoverTab[124893]++
										emitter.buffer[emitter.buffer_pos] = value
										emitter.buffer_pos++
										emitter.column++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:24
	// _ = "end of CoverTab[124893]"
}

// Put a line break to the output buffer.
func put_break(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:28
	_go_fuzz_dep_.CoverTab[124897]++
										if emitter.buffer_pos+5 >= len(emitter.buffer) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:29
		_go_fuzz_dep_.CoverTab[124900]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:29
		return !yaml_emitter_flush(emitter)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:29
		// _ = "end of CoverTab[124900]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:29
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:29
		_go_fuzz_dep_.CoverTab[124901]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:30
		// _ = "end of CoverTab[124901]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:31
		_go_fuzz_dep_.CoverTab[124902]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:31
		// _ = "end of CoverTab[124902]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:31
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:31
	// _ = "end of CoverTab[124897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:31
	_go_fuzz_dep_.CoverTab[124898]++
										switch emitter.line_break {
	case yaml_CR_BREAK:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:33
		_go_fuzz_dep_.CoverTab[124903]++
											emitter.buffer[emitter.buffer_pos] = '\r'
											emitter.buffer_pos += 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:35
		// _ = "end of CoverTab[124903]"
	case yaml_LN_BREAK:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:36
		_go_fuzz_dep_.CoverTab[124904]++
											emitter.buffer[emitter.buffer_pos] = '\n'
											emitter.buffer_pos += 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:38
		// _ = "end of CoverTab[124904]"
	case yaml_CRLN_BREAK:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:39
		_go_fuzz_dep_.CoverTab[124905]++
											emitter.buffer[emitter.buffer_pos+0] = '\r'
											emitter.buffer[emitter.buffer_pos+1] = '\n'
											emitter.buffer_pos += 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:42
		// _ = "end of CoverTab[124905]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:43
		_go_fuzz_dep_.CoverTab[124906]++
											panic("unknown line break setting")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:44
		// _ = "end of CoverTab[124906]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:45
	// _ = "end of CoverTab[124898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:45
	_go_fuzz_dep_.CoverTab[124899]++
										emitter.column = 0
										emitter.line++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:48
	// _ = "end of CoverTab[124899]"
}

// Copy a character from a string into buffer.
func write(emitter *yaml_emitter_t, s []byte, i *int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:52
	_go_fuzz_dep_.CoverTab[124907]++
										if emitter.buffer_pos+5 >= len(emitter.buffer) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:53
		_go_fuzz_dep_.CoverTab[124910]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:53
		return !yaml_emitter_flush(emitter)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:53
		// _ = "end of CoverTab[124910]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:53
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:53
		_go_fuzz_dep_.CoverTab[124911]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:54
		// _ = "end of CoverTab[124911]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:55
		_go_fuzz_dep_.CoverTab[124912]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:55
		// _ = "end of CoverTab[124912]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:55
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:55
	// _ = "end of CoverTab[124907]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:55
	_go_fuzz_dep_.CoverTab[124908]++
										p := emitter.buffer_pos
										w := width(s[*i])
										switch w {
	case 4:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:59
		_go_fuzz_dep_.CoverTab[124913]++
											emitter.buffer[p+3] = s[*i+3]
											fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:61
		// _ = "end of CoverTab[124913]"
	case 3:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:62
		_go_fuzz_dep_.CoverTab[124914]++
											emitter.buffer[p+2] = s[*i+2]
											fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:64
		// _ = "end of CoverTab[124914]"
	case 2:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:65
		_go_fuzz_dep_.CoverTab[124915]++
											emitter.buffer[p+1] = s[*i+1]
											fallthrough
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:67
		// _ = "end of CoverTab[124915]"
	case 1:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:68
		_go_fuzz_dep_.CoverTab[124916]++
											emitter.buffer[p+0] = s[*i+0]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:69
		// _ = "end of CoverTab[124916]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:70
		_go_fuzz_dep_.CoverTab[124917]++
											panic("unknown character width")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:71
		// _ = "end of CoverTab[124917]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:72
	// _ = "end of CoverTab[124908]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:72
	_go_fuzz_dep_.CoverTab[124909]++
										emitter.column++
										emitter.buffer_pos += w
										*i += w
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:76
	// _ = "end of CoverTab[124909]"
}

// Write a whole string into buffer.
func write_all(emitter *yaml_emitter_t, s []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:80
	_go_fuzz_dep_.CoverTab[124918]++
										for i := 0; i < len(s); {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:81
		_go_fuzz_dep_.CoverTab[124920]++
											if !write(emitter, s, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:82
			_go_fuzz_dep_.CoverTab[124921]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:83
			// _ = "end of CoverTab[124921]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:84
			_go_fuzz_dep_.CoverTab[124922]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:84
			// _ = "end of CoverTab[124922]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:84
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:84
		// _ = "end of CoverTab[124920]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:85
	// _ = "end of CoverTab[124918]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:85
	_go_fuzz_dep_.CoverTab[124919]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:86
	// _ = "end of CoverTab[124919]"
}

// Copy a line break character from a string into buffer.
func write_break(emitter *yaml_emitter_t, s []byte, i *int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:90
	_go_fuzz_dep_.CoverTab[124923]++
										if s[*i] == '\n' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:91
		_go_fuzz_dep_.CoverTab[124925]++
											if !put_break(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:92
			_go_fuzz_dep_.CoverTab[124927]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:93
			// _ = "end of CoverTab[124927]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:94
			_go_fuzz_dep_.CoverTab[124928]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:94
			// _ = "end of CoverTab[124928]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:94
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:94
		// _ = "end of CoverTab[124925]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:94
		_go_fuzz_dep_.CoverTab[124926]++
											*i++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:95
		// _ = "end of CoverTab[124926]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:96
		_go_fuzz_dep_.CoverTab[124929]++
											if !write(emitter, s, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:97
			_go_fuzz_dep_.CoverTab[124931]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:98
			// _ = "end of CoverTab[124931]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:99
			_go_fuzz_dep_.CoverTab[124932]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:99
			// _ = "end of CoverTab[124932]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:99
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:99
		// _ = "end of CoverTab[124929]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:99
		_go_fuzz_dep_.CoverTab[124930]++
											emitter.column = 0
											emitter.line++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:101
		// _ = "end of CoverTab[124930]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:102
	// _ = "end of CoverTab[124923]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:102
	_go_fuzz_dep_.CoverTab[124924]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:103
	// _ = "end of CoverTab[124924]"
}

// Set an emitter error and return false.
func yaml_emitter_set_emitter_error(emitter *yaml_emitter_t, problem string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:107
	_go_fuzz_dep_.CoverTab[124933]++
										emitter.error = yaml_EMITTER_ERROR
										emitter.problem = problem
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:110
	// _ = "end of CoverTab[124933]"
}

// Emit an event.
func yaml_emitter_emit(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:114
	_go_fuzz_dep_.CoverTab[124934]++
										emitter.events = append(emitter.events, *event)
										for !yaml_emitter_need_more_events(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:116
		_go_fuzz_dep_.CoverTab[124936]++
											event := &emitter.events[emitter.events_head]
											if !yaml_emitter_analyze_event(emitter, event) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:118
			_go_fuzz_dep_.CoverTab[124939]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:119
			// _ = "end of CoverTab[124939]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:120
			_go_fuzz_dep_.CoverTab[124940]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:120
			// _ = "end of CoverTab[124940]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:120
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:120
		// _ = "end of CoverTab[124936]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:120
		_go_fuzz_dep_.CoverTab[124937]++
											if !yaml_emitter_state_machine(emitter, event) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:121
			_go_fuzz_dep_.CoverTab[124941]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:122
			// _ = "end of CoverTab[124941]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:123
			_go_fuzz_dep_.CoverTab[124942]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:123
			// _ = "end of CoverTab[124942]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:123
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:123
		// _ = "end of CoverTab[124937]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:123
		_go_fuzz_dep_.CoverTab[124938]++
											yaml_event_delete(event)
											emitter.events_head++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:125
		// _ = "end of CoverTab[124938]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:126
	// _ = "end of CoverTab[124934]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:126
	_go_fuzz_dep_.CoverTab[124935]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:127
	// _ = "end of CoverTab[124935]"
}

// Check if we need to accumulate more events before emitting.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:130
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:130
// We accumulate extra
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:130
//   - 1 event for DOCUMENT-START
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:130
//   - 2 events for SEQUENCE-START
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:130
//   - 3 events for MAPPING-START
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:137
func yaml_emitter_need_more_events(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:137
	_go_fuzz_dep_.CoverTab[124943]++
										if emitter.events_head == len(emitter.events) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:138
		_go_fuzz_dep_.CoverTab[124948]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:139
		// _ = "end of CoverTab[124948]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:140
		_go_fuzz_dep_.CoverTab[124949]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:140
		// _ = "end of CoverTab[124949]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:140
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:140
	// _ = "end of CoverTab[124943]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:140
	_go_fuzz_dep_.CoverTab[124944]++
										var accumulate int
										switch emitter.events[emitter.events_head].typ {
	case yaml_DOCUMENT_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:143
		_go_fuzz_dep_.CoverTab[124950]++
											accumulate = 1
											break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:145
		// _ = "end of CoverTab[124950]"
	case yaml_SEQUENCE_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:146
		_go_fuzz_dep_.CoverTab[124951]++
											accumulate = 2
											break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:148
		// _ = "end of CoverTab[124951]"
	case yaml_MAPPING_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:149
		_go_fuzz_dep_.CoverTab[124952]++
											accumulate = 3
											break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:151
		// _ = "end of CoverTab[124952]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:152
		_go_fuzz_dep_.CoverTab[124953]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:153
		// _ = "end of CoverTab[124953]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:154
	// _ = "end of CoverTab[124944]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:154
	_go_fuzz_dep_.CoverTab[124945]++
										if len(emitter.events)-emitter.events_head > accumulate {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:155
		_go_fuzz_dep_.CoverTab[124954]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:156
		// _ = "end of CoverTab[124954]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:157
		_go_fuzz_dep_.CoverTab[124955]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:157
		// _ = "end of CoverTab[124955]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:157
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:157
	// _ = "end of CoverTab[124945]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:157
	_go_fuzz_dep_.CoverTab[124946]++
										var level int
										for i := emitter.events_head; i < len(emitter.events); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:159
		_go_fuzz_dep_.CoverTab[124956]++
											switch emitter.events[i].typ {
		case yaml_STREAM_START_EVENT, yaml_DOCUMENT_START_EVENT, yaml_SEQUENCE_START_EVENT, yaml_MAPPING_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:161
			_go_fuzz_dep_.CoverTab[124958]++
												level++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:162
			// _ = "end of CoverTab[124958]"
		case yaml_STREAM_END_EVENT, yaml_DOCUMENT_END_EVENT, yaml_SEQUENCE_END_EVENT, yaml_MAPPING_END_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:163
			_go_fuzz_dep_.CoverTab[124959]++
												level--
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:164
			// _ = "end of CoverTab[124959]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:164
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:164
			_go_fuzz_dep_.CoverTab[124960]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:164
			// _ = "end of CoverTab[124960]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:165
		// _ = "end of CoverTab[124956]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:165
		_go_fuzz_dep_.CoverTab[124957]++
											if level == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:166
			_go_fuzz_dep_.CoverTab[124961]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:167
			// _ = "end of CoverTab[124961]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:168
			_go_fuzz_dep_.CoverTab[124962]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:168
			// _ = "end of CoverTab[124962]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:168
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:168
		// _ = "end of CoverTab[124957]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:169
	// _ = "end of CoverTab[124946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:169
	_go_fuzz_dep_.CoverTab[124947]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:170
	// _ = "end of CoverTab[124947]"
}

// Append a directive to the directives stack.
func yaml_emitter_append_tag_directive(emitter *yaml_emitter_t, value *yaml_tag_directive_t, allow_duplicates bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:174
	_go_fuzz_dep_.CoverTab[124963]++
										for i := 0; i < len(emitter.tag_directives); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:175
		_go_fuzz_dep_.CoverTab[124965]++
											if bytes.Equal(value.handle, emitter.tag_directives[i].handle) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:176
			_go_fuzz_dep_.CoverTab[124966]++
												if allow_duplicates {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:177
				_go_fuzz_dep_.CoverTab[124968]++
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:178
				// _ = "end of CoverTab[124968]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:179
				_go_fuzz_dep_.CoverTab[124969]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:179
				// _ = "end of CoverTab[124969]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:179
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:179
			// _ = "end of CoverTab[124966]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:179
			_go_fuzz_dep_.CoverTab[124967]++
												return yaml_emitter_set_emitter_error(emitter, "duplicate %TAG directive")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:180
			// _ = "end of CoverTab[124967]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:181
			_go_fuzz_dep_.CoverTab[124970]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:181
			// _ = "end of CoverTab[124970]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:181
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:181
		// _ = "end of CoverTab[124965]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:182
	// _ = "end of CoverTab[124963]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:182
	_go_fuzz_dep_.CoverTab[124964]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:186
	tag_copy := yaml_tag_directive_t{
		handle:	make([]byte, len(value.handle)),
		prefix:	make([]byte, len(value.prefix)),
	}
										copy(tag_copy.handle, value.handle)
										copy(tag_copy.prefix, value.prefix)
										emitter.tag_directives = append(emitter.tag_directives, tag_copy)
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:193
	// _ = "end of CoverTab[124964]"
}

// Increase the indentation level.
func yaml_emitter_increase_indent(emitter *yaml_emitter_t, flow, indentless bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:197
	_go_fuzz_dep_.CoverTab[124971]++
										emitter.indents = append(emitter.indents, emitter.indent)
										if emitter.indent < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:199
		_go_fuzz_dep_.CoverTab[124973]++
											if flow {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:200
			_go_fuzz_dep_.CoverTab[124974]++
												emitter.indent = emitter.best_indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:201
			// _ = "end of CoverTab[124974]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:202
			_go_fuzz_dep_.CoverTab[124975]++
												emitter.indent = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:203
			// _ = "end of CoverTab[124975]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:204
		// _ = "end of CoverTab[124973]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:205
		_go_fuzz_dep_.CoverTab[124976]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:205
		if !indentless {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:205
			_go_fuzz_dep_.CoverTab[124977]++
												emitter.indent += emitter.best_indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:206
			// _ = "end of CoverTab[124977]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:207
			_go_fuzz_dep_.CoverTab[124978]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:207
			// _ = "end of CoverTab[124978]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:207
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:207
		// _ = "end of CoverTab[124976]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:207
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:207
	// _ = "end of CoverTab[124971]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:207
	_go_fuzz_dep_.CoverTab[124972]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:208
	// _ = "end of CoverTab[124972]"
}

// State dispatcher.
func yaml_emitter_state_machine(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:212
	_go_fuzz_dep_.CoverTab[124979]++
										switch emitter.state {
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:214
		_go_fuzz_dep_.CoverTab[124981]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:214
		// _ = "end of CoverTab[124981]"
	case yaml_EMIT_STREAM_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:215
		_go_fuzz_dep_.CoverTab[124982]++
											return yaml_emitter_emit_stream_start(emitter, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:216
		// _ = "end of CoverTab[124982]"

	case yaml_EMIT_FIRST_DOCUMENT_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:218
		_go_fuzz_dep_.CoverTab[124983]++
											return yaml_emitter_emit_document_start(emitter, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:219
		// _ = "end of CoverTab[124983]"

	case yaml_EMIT_DOCUMENT_START_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:221
		_go_fuzz_dep_.CoverTab[124984]++
											return yaml_emitter_emit_document_start(emitter, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:222
		// _ = "end of CoverTab[124984]"

	case yaml_EMIT_DOCUMENT_CONTENT_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:224
		_go_fuzz_dep_.CoverTab[124985]++
											return yaml_emitter_emit_document_content(emitter, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:225
		// _ = "end of CoverTab[124985]"

	case yaml_EMIT_DOCUMENT_END_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:227
		_go_fuzz_dep_.CoverTab[124986]++
											return yaml_emitter_emit_document_end(emitter, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:228
		// _ = "end of CoverTab[124986]"

	case yaml_EMIT_FLOW_SEQUENCE_FIRST_ITEM_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:230
		_go_fuzz_dep_.CoverTab[124987]++
											return yaml_emitter_emit_flow_sequence_item(emitter, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:231
		// _ = "end of CoverTab[124987]"

	case yaml_EMIT_FLOW_SEQUENCE_ITEM_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:233
		_go_fuzz_dep_.CoverTab[124988]++
											return yaml_emitter_emit_flow_sequence_item(emitter, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:234
		// _ = "end of CoverTab[124988]"

	case yaml_EMIT_FLOW_MAPPING_FIRST_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:236
		_go_fuzz_dep_.CoverTab[124989]++
											return yaml_emitter_emit_flow_mapping_key(emitter, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:237
		// _ = "end of CoverTab[124989]"

	case yaml_EMIT_FLOW_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:239
		_go_fuzz_dep_.CoverTab[124990]++
											return yaml_emitter_emit_flow_mapping_key(emitter, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:240
		// _ = "end of CoverTab[124990]"

	case yaml_EMIT_FLOW_MAPPING_SIMPLE_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:242
		_go_fuzz_dep_.CoverTab[124991]++
											return yaml_emitter_emit_flow_mapping_value(emitter, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:243
		// _ = "end of CoverTab[124991]"

	case yaml_EMIT_FLOW_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:245
		_go_fuzz_dep_.CoverTab[124992]++
											return yaml_emitter_emit_flow_mapping_value(emitter, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:246
		// _ = "end of CoverTab[124992]"

	case yaml_EMIT_BLOCK_SEQUENCE_FIRST_ITEM_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:248
		_go_fuzz_dep_.CoverTab[124993]++
											return yaml_emitter_emit_block_sequence_item(emitter, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:249
		// _ = "end of CoverTab[124993]"

	case yaml_EMIT_BLOCK_SEQUENCE_ITEM_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:251
		_go_fuzz_dep_.CoverTab[124994]++
											return yaml_emitter_emit_block_sequence_item(emitter, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:252
		// _ = "end of CoverTab[124994]"

	case yaml_EMIT_BLOCK_MAPPING_FIRST_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:254
		_go_fuzz_dep_.CoverTab[124995]++
											return yaml_emitter_emit_block_mapping_key(emitter, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:255
		// _ = "end of CoverTab[124995]"

	case yaml_EMIT_BLOCK_MAPPING_KEY_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:257
		_go_fuzz_dep_.CoverTab[124996]++
											return yaml_emitter_emit_block_mapping_key(emitter, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:258
		// _ = "end of CoverTab[124996]"

	case yaml_EMIT_BLOCK_MAPPING_SIMPLE_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:260
		_go_fuzz_dep_.CoverTab[124997]++
											return yaml_emitter_emit_block_mapping_value(emitter, event, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:261
		// _ = "end of CoverTab[124997]"

	case yaml_EMIT_BLOCK_MAPPING_VALUE_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:263
		_go_fuzz_dep_.CoverTab[124998]++
											return yaml_emitter_emit_block_mapping_value(emitter, event, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:264
		// _ = "end of CoverTab[124998]"

	case yaml_EMIT_END_STATE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:266
		_go_fuzz_dep_.CoverTab[124999]++
											return yaml_emitter_set_emitter_error(emitter, "expected nothing after STREAM-END")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:267
		// _ = "end of CoverTab[124999]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:268
	// _ = "end of CoverTab[124979]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:268
	_go_fuzz_dep_.CoverTab[124980]++
										panic("invalid emitter state")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:269
	// _ = "end of CoverTab[124980]"
}

// Expect STREAM-START.
func yaml_emitter_emit_stream_start(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:273
	_go_fuzz_dep_.CoverTab[125000]++
										if event.typ != yaml_STREAM_START_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:274
		_go_fuzz_dep_.CoverTab[125008]++
											return yaml_emitter_set_emitter_error(emitter, "expected STREAM-START")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:275
		// _ = "end of CoverTab[125008]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:276
		_go_fuzz_dep_.CoverTab[125009]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:276
		// _ = "end of CoverTab[125009]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:276
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:276
	// _ = "end of CoverTab[125000]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:276
	_go_fuzz_dep_.CoverTab[125001]++
										if emitter.encoding == yaml_ANY_ENCODING {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:277
		_go_fuzz_dep_.CoverTab[125010]++
											emitter.encoding = event.encoding
											if emitter.encoding == yaml_ANY_ENCODING {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:279
			_go_fuzz_dep_.CoverTab[125011]++
												emitter.encoding = yaml_UTF8_ENCODING
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:280
			// _ = "end of CoverTab[125011]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:281
			_go_fuzz_dep_.CoverTab[125012]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:281
			// _ = "end of CoverTab[125012]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:281
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:281
		// _ = "end of CoverTab[125010]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:282
		_go_fuzz_dep_.CoverTab[125013]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:282
		// _ = "end of CoverTab[125013]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:282
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:282
	// _ = "end of CoverTab[125001]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:282
	_go_fuzz_dep_.CoverTab[125002]++
										if emitter.best_indent < 2 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:283
		_go_fuzz_dep_.CoverTab[125014]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:283
		return emitter.best_indent > 9
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:283
		// _ = "end of CoverTab[125014]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:283
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:283
		_go_fuzz_dep_.CoverTab[125015]++
											emitter.best_indent = 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:284
		// _ = "end of CoverTab[125015]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:285
		_go_fuzz_dep_.CoverTab[125016]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:285
		// _ = "end of CoverTab[125016]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:285
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:285
	// _ = "end of CoverTab[125002]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:285
	_go_fuzz_dep_.CoverTab[125003]++
										if emitter.best_width >= 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:286
		_go_fuzz_dep_.CoverTab[125017]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:286
		return emitter.best_width <= emitter.best_indent*2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:286
		// _ = "end of CoverTab[125017]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:286
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:286
		_go_fuzz_dep_.CoverTab[125018]++
											emitter.best_width = 80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:287
		// _ = "end of CoverTab[125018]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:288
		_go_fuzz_dep_.CoverTab[125019]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:288
		// _ = "end of CoverTab[125019]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:288
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:288
	// _ = "end of CoverTab[125003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:288
	_go_fuzz_dep_.CoverTab[125004]++
										if emitter.best_width < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:289
		_go_fuzz_dep_.CoverTab[125020]++
											emitter.best_width = 1<<31 - 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:290
		// _ = "end of CoverTab[125020]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:291
		_go_fuzz_dep_.CoverTab[125021]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:291
		// _ = "end of CoverTab[125021]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:291
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:291
	// _ = "end of CoverTab[125004]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:291
	_go_fuzz_dep_.CoverTab[125005]++
										if emitter.line_break == yaml_ANY_BREAK {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:292
		_go_fuzz_dep_.CoverTab[125022]++
											emitter.line_break = yaml_LN_BREAK
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:293
		// _ = "end of CoverTab[125022]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:294
		_go_fuzz_dep_.CoverTab[125023]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:294
		// _ = "end of CoverTab[125023]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:294
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:294
	// _ = "end of CoverTab[125005]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:294
	_go_fuzz_dep_.CoverTab[125006]++

										emitter.indent = -1
										emitter.line = 0
										emitter.column = 0
										emitter.whitespace = true
										emitter.indention = true

										if emitter.encoding != yaml_UTF8_ENCODING {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:302
		_go_fuzz_dep_.CoverTab[125024]++
											if !yaml_emitter_write_bom(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:303
			_go_fuzz_dep_.CoverTab[125025]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:304
			// _ = "end of CoverTab[125025]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:305
			_go_fuzz_dep_.CoverTab[125026]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:305
			// _ = "end of CoverTab[125026]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:305
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:305
		// _ = "end of CoverTab[125024]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:306
		_go_fuzz_dep_.CoverTab[125027]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:306
		// _ = "end of CoverTab[125027]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:306
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:306
	// _ = "end of CoverTab[125006]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:306
	_go_fuzz_dep_.CoverTab[125007]++
										emitter.state = yaml_EMIT_FIRST_DOCUMENT_START_STATE
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:308
	// _ = "end of CoverTab[125007]"
}

// Expect DOCUMENT-START or STREAM-END.
func yaml_emitter_emit_document_start(emitter *yaml_emitter_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:312
	_go_fuzz_dep_.CoverTab[125028]++

										if event.typ == yaml_DOCUMENT_START_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:314
		_go_fuzz_dep_.CoverTab[125031]++

											if event.version_directive != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:316
			_go_fuzz_dep_.CoverTab[125041]++
												if !yaml_emitter_analyze_version_directive(emitter, event.version_directive) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:317
				_go_fuzz_dep_.CoverTab[125042]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:318
				// _ = "end of CoverTab[125042]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:319
				_go_fuzz_dep_.CoverTab[125043]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:319
				// _ = "end of CoverTab[125043]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:319
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:319
			// _ = "end of CoverTab[125041]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:320
			_go_fuzz_dep_.CoverTab[125044]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:320
			// _ = "end of CoverTab[125044]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:320
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:320
		// _ = "end of CoverTab[125031]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:320
		_go_fuzz_dep_.CoverTab[125032]++

											for i := 0; i < len(event.tag_directives); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:322
			_go_fuzz_dep_.CoverTab[125045]++
												tag_directive := &event.tag_directives[i]
												if !yaml_emitter_analyze_tag_directive(emitter, tag_directive) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:324
				_go_fuzz_dep_.CoverTab[125047]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:325
				// _ = "end of CoverTab[125047]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:326
				_go_fuzz_dep_.CoverTab[125048]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:326
				// _ = "end of CoverTab[125048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:326
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:326
			// _ = "end of CoverTab[125045]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:326
			_go_fuzz_dep_.CoverTab[125046]++
												if !yaml_emitter_append_tag_directive(emitter, tag_directive, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:327
				_go_fuzz_dep_.CoverTab[125049]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:328
				// _ = "end of CoverTab[125049]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:329
				_go_fuzz_dep_.CoverTab[125050]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:329
				// _ = "end of CoverTab[125050]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:329
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:329
			// _ = "end of CoverTab[125046]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:330
		// _ = "end of CoverTab[125032]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:330
		_go_fuzz_dep_.CoverTab[125033]++

											for i := 0; i < len(default_tag_directives); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:332
			_go_fuzz_dep_.CoverTab[125051]++
												tag_directive := &default_tag_directives[i]
												if !yaml_emitter_append_tag_directive(emitter, tag_directive, true) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:334
				_go_fuzz_dep_.CoverTab[125052]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:335
				// _ = "end of CoverTab[125052]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:336
				_go_fuzz_dep_.CoverTab[125053]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:336
				// _ = "end of CoverTab[125053]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:336
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:336
			// _ = "end of CoverTab[125051]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:337
		// _ = "end of CoverTab[125033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:337
		_go_fuzz_dep_.CoverTab[125034]++

											implicit := event.implicit
											if !first || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:340
			_go_fuzz_dep_.CoverTab[125054]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:340
			return emitter.canonical
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:340
			// _ = "end of CoverTab[125054]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:340
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:340
			_go_fuzz_dep_.CoverTab[125055]++
												implicit = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:341
			// _ = "end of CoverTab[125055]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:342
			_go_fuzz_dep_.CoverTab[125056]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:342
			// _ = "end of CoverTab[125056]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:342
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:342
		// _ = "end of CoverTab[125034]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:342
		_go_fuzz_dep_.CoverTab[125035]++

											if emitter.open_ended && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
			_go_fuzz_dep_.CoverTab[125057]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
			return (event.version_directive != nil || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
				_go_fuzz_dep_.CoverTab[125058]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
				return len(event.tag_directives) > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
				// _ = "end of CoverTab[125058]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
			// _ = "end of CoverTab[125057]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:344
			_go_fuzz_dep_.CoverTab[125059]++
												if !yaml_emitter_write_indicator(emitter, []byte("..."), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:345
				_go_fuzz_dep_.CoverTab[125061]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:346
				// _ = "end of CoverTab[125061]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:347
				_go_fuzz_dep_.CoverTab[125062]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:347
				// _ = "end of CoverTab[125062]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:347
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:347
			// _ = "end of CoverTab[125059]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:347
			_go_fuzz_dep_.CoverTab[125060]++
												if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:348
				_go_fuzz_dep_.CoverTab[125063]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:349
				// _ = "end of CoverTab[125063]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:350
				_go_fuzz_dep_.CoverTab[125064]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:350
				// _ = "end of CoverTab[125064]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:350
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:350
			// _ = "end of CoverTab[125060]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:351
			_go_fuzz_dep_.CoverTab[125065]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:351
			// _ = "end of CoverTab[125065]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:351
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:351
		// _ = "end of CoverTab[125035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:351
		_go_fuzz_dep_.CoverTab[125036]++

											if event.version_directive != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:353
			_go_fuzz_dep_.CoverTab[125066]++
												implicit = false
												if !yaml_emitter_write_indicator(emitter, []byte("%YAML"), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:355
				_go_fuzz_dep_.CoverTab[125069]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:356
				// _ = "end of CoverTab[125069]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:357
				_go_fuzz_dep_.CoverTab[125070]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:357
				// _ = "end of CoverTab[125070]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:357
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:357
			// _ = "end of CoverTab[125066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:357
			_go_fuzz_dep_.CoverTab[125067]++
												if !yaml_emitter_write_indicator(emitter, []byte("1.1"), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:358
				_go_fuzz_dep_.CoverTab[125071]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:359
				// _ = "end of CoverTab[125071]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:360
				_go_fuzz_dep_.CoverTab[125072]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:360
				// _ = "end of CoverTab[125072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:360
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:360
			// _ = "end of CoverTab[125067]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:360
			_go_fuzz_dep_.CoverTab[125068]++
												if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:361
				_go_fuzz_dep_.CoverTab[125073]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:362
				// _ = "end of CoverTab[125073]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:363
				_go_fuzz_dep_.CoverTab[125074]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:363
				// _ = "end of CoverTab[125074]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:363
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:363
			// _ = "end of CoverTab[125068]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:364
			_go_fuzz_dep_.CoverTab[125075]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:364
			// _ = "end of CoverTab[125075]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:364
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:364
		// _ = "end of CoverTab[125036]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:364
		_go_fuzz_dep_.CoverTab[125037]++

											if len(event.tag_directives) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:366
			_go_fuzz_dep_.CoverTab[125076]++
												implicit = false
												for i := 0; i < len(event.tag_directives); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:368
				_go_fuzz_dep_.CoverTab[125077]++
													tag_directive := &event.tag_directives[i]
													if !yaml_emitter_write_indicator(emitter, []byte("%TAG"), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:370
					_go_fuzz_dep_.CoverTab[125081]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:371
					// _ = "end of CoverTab[125081]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:372
					_go_fuzz_dep_.CoverTab[125082]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:372
					// _ = "end of CoverTab[125082]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:372
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:372
				// _ = "end of CoverTab[125077]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:372
				_go_fuzz_dep_.CoverTab[125078]++
													if !yaml_emitter_write_tag_handle(emitter, tag_directive.handle) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:373
					_go_fuzz_dep_.CoverTab[125083]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:374
					// _ = "end of CoverTab[125083]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:375
					_go_fuzz_dep_.CoverTab[125084]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:375
					// _ = "end of CoverTab[125084]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:375
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:375
				// _ = "end of CoverTab[125078]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:375
				_go_fuzz_dep_.CoverTab[125079]++
													if !yaml_emitter_write_tag_content(emitter, tag_directive.prefix, true) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:376
					_go_fuzz_dep_.CoverTab[125085]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:377
					// _ = "end of CoverTab[125085]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:378
					_go_fuzz_dep_.CoverTab[125086]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:378
					// _ = "end of CoverTab[125086]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:378
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:378
				// _ = "end of CoverTab[125079]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:378
				_go_fuzz_dep_.CoverTab[125080]++
													if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:379
					_go_fuzz_dep_.CoverTab[125087]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:380
					// _ = "end of CoverTab[125087]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:381
					_go_fuzz_dep_.CoverTab[125088]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:381
					// _ = "end of CoverTab[125088]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:381
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:381
				// _ = "end of CoverTab[125080]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:382
			// _ = "end of CoverTab[125076]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:383
			_go_fuzz_dep_.CoverTab[125089]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:383
			// _ = "end of CoverTab[125089]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:383
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:383
		// _ = "end of CoverTab[125037]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:383
		_go_fuzz_dep_.CoverTab[125038]++

											if yaml_emitter_check_empty_document(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:385
			_go_fuzz_dep_.CoverTab[125090]++
												implicit = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:386
			// _ = "end of CoverTab[125090]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:387
			_go_fuzz_dep_.CoverTab[125091]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:387
			// _ = "end of CoverTab[125091]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:387
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:387
		// _ = "end of CoverTab[125038]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:387
		_go_fuzz_dep_.CoverTab[125039]++
											if !implicit {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:388
			_go_fuzz_dep_.CoverTab[125092]++
												if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:389
				_go_fuzz_dep_.CoverTab[125095]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:390
				// _ = "end of CoverTab[125095]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:391
				_go_fuzz_dep_.CoverTab[125096]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:391
				// _ = "end of CoverTab[125096]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:391
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:391
			// _ = "end of CoverTab[125092]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:391
			_go_fuzz_dep_.CoverTab[125093]++
												if !yaml_emitter_write_indicator(emitter, []byte("---"), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:392
				_go_fuzz_dep_.CoverTab[125097]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:393
				// _ = "end of CoverTab[125097]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:394
				_go_fuzz_dep_.CoverTab[125098]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:394
				// _ = "end of CoverTab[125098]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:394
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:394
			// _ = "end of CoverTab[125093]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:394
			_go_fuzz_dep_.CoverTab[125094]++
												if emitter.canonical {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:395
				_go_fuzz_dep_.CoverTab[125099]++
													if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:396
					_go_fuzz_dep_.CoverTab[125100]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:397
					// _ = "end of CoverTab[125100]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:398
					_go_fuzz_dep_.CoverTab[125101]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:398
					// _ = "end of CoverTab[125101]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:398
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:398
				// _ = "end of CoverTab[125099]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:399
				_go_fuzz_dep_.CoverTab[125102]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:399
				// _ = "end of CoverTab[125102]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:399
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:399
			// _ = "end of CoverTab[125094]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:400
			_go_fuzz_dep_.CoverTab[125103]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:400
			// _ = "end of CoverTab[125103]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:400
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:400
		// _ = "end of CoverTab[125039]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:400
		_go_fuzz_dep_.CoverTab[125040]++

											emitter.state = yaml_EMIT_DOCUMENT_CONTENT_STATE
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:403
		// _ = "end of CoverTab[125040]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:404
		_go_fuzz_dep_.CoverTab[125104]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:404
		// _ = "end of CoverTab[125104]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:404
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:404
	// _ = "end of CoverTab[125028]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:404
	_go_fuzz_dep_.CoverTab[125029]++

										if event.typ == yaml_STREAM_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:406
		_go_fuzz_dep_.CoverTab[125105]++
											if emitter.open_ended {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:407
			_go_fuzz_dep_.CoverTab[125108]++
												if !yaml_emitter_write_indicator(emitter, []byte("..."), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:408
				_go_fuzz_dep_.CoverTab[125110]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:409
				// _ = "end of CoverTab[125110]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:410
				_go_fuzz_dep_.CoverTab[125111]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:410
				// _ = "end of CoverTab[125111]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:410
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:410
			// _ = "end of CoverTab[125108]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:410
			_go_fuzz_dep_.CoverTab[125109]++
												if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:411
				_go_fuzz_dep_.CoverTab[125112]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:412
				// _ = "end of CoverTab[125112]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:413
				_go_fuzz_dep_.CoverTab[125113]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:413
				// _ = "end of CoverTab[125113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:413
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:413
			// _ = "end of CoverTab[125109]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:414
			_go_fuzz_dep_.CoverTab[125114]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:414
			// _ = "end of CoverTab[125114]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:414
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:414
		// _ = "end of CoverTab[125105]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:414
		_go_fuzz_dep_.CoverTab[125106]++
											if !yaml_emitter_flush(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:415
			_go_fuzz_dep_.CoverTab[125115]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:416
			// _ = "end of CoverTab[125115]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:417
			_go_fuzz_dep_.CoverTab[125116]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:417
			// _ = "end of CoverTab[125116]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:417
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:417
		// _ = "end of CoverTab[125106]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:417
		_go_fuzz_dep_.CoverTab[125107]++
											emitter.state = yaml_EMIT_END_STATE
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:419
		// _ = "end of CoverTab[125107]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:420
		_go_fuzz_dep_.CoverTab[125117]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:420
		// _ = "end of CoverTab[125117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:420
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:420
	// _ = "end of CoverTab[125029]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:420
	_go_fuzz_dep_.CoverTab[125030]++

										return yaml_emitter_set_emitter_error(emitter, "expected DOCUMENT-START or STREAM-END")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:422
	// _ = "end of CoverTab[125030]"
}

// Expect the root node.
func yaml_emitter_emit_document_content(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:426
	_go_fuzz_dep_.CoverTab[125118]++
										emitter.states = append(emitter.states, yaml_EMIT_DOCUMENT_END_STATE)
										return yaml_emitter_emit_node(emitter, event, true, false, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:428
	// _ = "end of CoverTab[125118]"
}

// Expect DOCUMENT-END.
func yaml_emitter_emit_document_end(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:432
	_go_fuzz_dep_.CoverTab[125119]++
										if event.typ != yaml_DOCUMENT_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:433
		_go_fuzz_dep_.CoverTab[125124]++
											return yaml_emitter_set_emitter_error(emitter, "expected DOCUMENT-END")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:434
		// _ = "end of CoverTab[125124]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:435
		_go_fuzz_dep_.CoverTab[125125]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:435
		// _ = "end of CoverTab[125125]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:435
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:435
	// _ = "end of CoverTab[125119]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:435
	_go_fuzz_dep_.CoverTab[125120]++
										if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:436
		_go_fuzz_dep_.CoverTab[125126]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:437
		// _ = "end of CoverTab[125126]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:438
		_go_fuzz_dep_.CoverTab[125127]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:438
		// _ = "end of CoverTab[125127]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:438
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:438
	// _ = "end of CoverTab[125120]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:438
	_go_fuzz_dep_.CoverTab[125121]++
										if !event.implicit {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:439
		_go_fuzz_dep_.CoverTab[125128]++

											if !yaml_emitter_write_indicator(emitter, []byte("..."), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:441
			_go_fuzz_dep_.CoverTab[125130]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:442
			// _ = "end of CoverTab[125130]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:443
			_go_fuzz_dep_.CoverTab[125131]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:443
			// _ = "end of CoverTab[125131]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:443
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:443
		// _ = "end of CoverTab[125128]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:443
		_go_fuzz_dep_.CoverTab[125129]++
											if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:444
			_go_fuzz_dep_.CoverTab[125132]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:445
			// _ = "end of CoverTab[125132]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:446
			_go_fuzz_dep_.CoverTab[125133]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:446
			// _ = "end of CoverTab[125133]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:446
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:446
		// _ = "end of CoverTab[125129]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:447
		_go_fuzz_dep_.CoverTab[125134]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:447
		// _ = "end of CoverTab[125134]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:447
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:447
	// _ = "end of CoverTab[125121]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:447
	_go_fuzz_dep_.CoverTab[125122]++
										if !yaml_emitter_flush(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:448
		_go_fuzz_dep_.CoverTab[125135]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:449
		// _ = "end of CoverTab[125135]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:450
		_go_fuzz_dep_.CoverTab[125136]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:450
		// _ = "end of CoverTab[125136]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:450
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:450
	// _ = "end of CoverTab[125122]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:450
	_go_fuzz_dep_.CoverTab[125123]++
										emitter.state = yaml_EMIT_DOCUMENT_START_STATE
										emitter.tag_directives = emitter.tag_directives[:0]
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:453
	// _ = "end of CoverTab[125123]"
}

// Expect a flow item node.
func yaml_emitter_emit_flow_sequence_item(emitter *yaml_emitter_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:457
	_go_fuzz_dep_.CoverTab[125137]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:458
		_go_fuzz_dep_.CoverTab[125142]++
											if !yaml_emitter_write_indicator(emitter, []byte{'['}, true, true, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:459
			_go_fuzz_dep_.CoverTab[125145]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:460
			// _ = "end of CoverTab[125145]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:461
			_go_fuzz_dep_.CoverTab[125146]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:461
			// _ = "end of CoverTab[125146]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:461
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:461
		// _ = "end of CoverTab[125142]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:461
		_go_fuzz_dep_.CoverTab[125143]++
											if !yaml_emitter_increase_indent(emitter, true, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:462
			_go_fuzz_dep_.CoverTab[125147]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:463
			// _ = "end of CoverTab[125147]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:464
			_go_fuzz_dep_.CoverTab[125148]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:464
			// _ = "end of CoverTab[125148]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:464
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:464
		// _ = "end of CoverTab[125143]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:464
		_go_fuzz_dep_.CoverTab[125144]++
											emitter.flow_level++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:465
		// _ = "end of CoverTab[125144]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:466
		_go_fuzz_dep_.CoverTab[125149]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:466
		// _ = "end of CoverTab[125149]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:466
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:466
	// _ = "end of CoverTab[125137]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:466
	_go_fuzz_dep_.CoverTab[125138]++

										if event.typ == yaml_SEQUENCE_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:468
		_go_fuzz_dep_.CoverTab[125150]++
											emitter.flow_level--
											emitter.indent = emitter.indents[len(emitter.indents)-1]
											emitter.indents = emitter.indents[:len(emitter.indents)-1]
											if emitter.canonical && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:472
			_go_fuzz_dep_.CoverTab[125153]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:472
			return !first
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:472
			// _ = "end of CoverTab[125153]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:472
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:472
			_go_fuzz_dep_.CoverTab[125154]++
												if !yaml_emitter_write_indicator(emitter, []byte{','}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:473
				_go_fuzz_dep_.CoverTab[125156]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:474
				// _ = "end of CoverTab[125156]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:475
				_go_fuzz_dep_.CoverTab[125157]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:475
				// _ = "end of CoverTab[125157]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:475
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:475
			// _ = "end of CoverTab[125154]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:475
			_go_fuzz_dep_.CoverTab[125155]++
												if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:476
				_go_fuzz_dep_.CoverTab[125158]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:477
				// _ = "end of CoverTab[125158]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:478
				_go_fuzz_dep_.CoverTab[125159]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:478
				// _ = "end of CoverTab[125159]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:478
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:478
			// _ = "end of CoverTab[125155]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:479
			_go_fuzz_dep_.CoverTab[125160]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:479
			// _ = "end of CoverTab[125160]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:479
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:479
		// _ = "end of CoverTab[125150]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:479
		_go_fuzz_dep_.CoverTab[125151]++
											if !yaml_emitter_write_indicator(emitter, []byte{']'}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:480
			_go_fuzz_dep_.CoverTab[125161]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:481
			// _ = "end of CoverTab[125161]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:482
			_go_fuzz_dep_.CoverTab[125162]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:482
			// _ = "end of CoverTab[125162]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:482
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:482
		// _ = "end of CoverTab[125151]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:482
		_go_fuzz_dep_.CoverTab[125152]++
											emitter.state = emitter.states[len(emitter.states)-1]
											emitter.states = emitter.states[:len(emitter.states)-1]

											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:486
		// _ = "end of CoverTab[125152]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:487
		_go_fuzz_dep_.CoverTab[125163]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:487
		// _ = "end of CoverTab[125163]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:487
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:487
	// _ = "end of CoverTab[125138]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:487
	_go_fuzz_dep_.CoverTab[125139]++

										if !first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:489
		_go_fuzz_dep_.CoverTab[125164]++
											if !yaml_emitter_write_indicator(emitter, []byte{','}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:490
			_go_fuzz_dep_.CoverTab[125165]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:491
			// _ = "end of CoverTab[125165]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:492
			_go_fuzz_dep_.CoverTab[125166]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:492
			// _ = "end of CoverTab[125166]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:492
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:492
		// _ = "end of CoverTab[125164]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:493
		_go_fuzz_dep_.CoverTab[125167]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:493
		// _ = "end of CoverTab[125167]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:493
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:493
	// _ = "end of CoverTab[125139]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:493
	_go_fuzz_dep_.CoverTab[125140]++

										if emitter.canonical || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:495
		_go_fuzz_dep_.CoverTab[125168]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:495
		return emitter.column > emitter.best_width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:495
		// _ = "end of CoverTab[125168]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:495
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:495
		_go_fuzz_dep_.CoverTab[125169]++
											if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:496
			_go_fuzz_dep_.CoverTab[125170]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:497
			// _ = "end of CoverTab[125170]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:498
			_go_fuzz_dep_.CoverTab[125171]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:498
			// _ = "end of CoverTab[125171]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:498
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:498
		// _ = "end of CoverTab[125169]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:499
		_go_fuzz_dep_.CoverTab[125172]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:499
		// _ = "end of CoverTab[125172]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:499
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:499
	// _ = "end of CoverTab[125140]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:499
	_go_fuzz_dep_.CoverTab[125141]++
										emitter.states = append(emitter.states, yaml_EMIT_FLOW_SEQUENCE_ITEM_STATE)
										return yaml_emitter_emit_node(emitter, event, false, true, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:501
	// _ = "end of CoverTab[125141]"
}

// Expect a flow key node.
func yaml_emitter_emit_flow_mapping_key(emitter *yaml_emitter_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:505
	_go_fuzz_dep_.CoverTab[125173]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:506
		_go_fuzz_dep_.CoverTab[125180]++
											if !yaml_emitter_write_indicator(emitter, []byte{'{'}, true, true, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:507
			_go_fuzz_dep_.CoverTab[125183]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:508
			// _ = "end of CoverTab[125183]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:509
			_go_fuzz_dep_.CoverTab[125184]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:509
			// _ = "end of CoverTab[125184]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:509
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:509
		// _ = "end of CoverTab[125180]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:509
		_go_fuzz_dep_.CoverTab[125181]++
											if !yaml_emitter_increase_indent(emitter, true, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:510
			_go_fuzz_dep_.CoverTab[125185]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:511
			// _ = "end of CoverTab[125185]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:512
			_go_fuzz_dep_.CoverTab[125186]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:512
			// _ = "end of CoverTab[125186]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:512
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:512
		// _ = "end of CoverTab[125181]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:512
		_go_fuzz_dep_.CoverTab[125182]++
											emitter.flow_level++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:513
		// _ = "end of CoverTab[125182]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:514
		_go_fuzz_dep_.CoverTab[125187]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:514
		// _ = "end of CoverTab[125187]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:514
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:514
	// _ = "end of CoverTab[125173]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:514
	_go_fuzz_dep_.CoverTab[125174]++

										if event.typ == yaml_MAPPING_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:516
		_go_fuzz_dep_.CoverTab[125188]++
											emitter.flow_level--
											emitter.indent = emitter.indents[len(emitter.indents)-1]
											emitter.indents = emitter.indents[:len(emitter.indents)-1]
											if emitter.canonical && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:520
			_go_fuzz_dep_.CoverTab[125191]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:520
			return !first
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:520
			// _ = "end of CoverTab[125191]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:520
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:520
			_go_fuzz_dep_.CoverTab[125192]++
												if !yaml_emitter_write_indicator(emitter, []byte{','}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:521
				_go_fuzz_dep_.CoverTab[125194]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:522
				// _ = "end of CoverTab[125194]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:523
				_go_fuzz_dep_.CoverTab[125195]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:523
				// _ = "end of CoverTab[125195]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:523
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:523
			// _ = "end of CoverTab[125192]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:523
			_go_fuzz_dep_.CoverTab[125193]++
												if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:524
				_go_fuzz_dep_.CoverTab[125196]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:525
				// _ = "end of CoverTab[125196]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:526
				_go_fuzz_dep_.CoverTab[125197]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:526
				// _ = "end of CoverTab[125197]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:526
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:526
			// _ = "end of CoverTab[125193]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:527
			_go_fuzz_dep_.CoverTab[125198]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:527
			// _ = "end of CoverTab[125198]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:527
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:527
		// _ = "end of CoverTab[125188]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:527
		_go_fuzz_dep_.CoverTab[125189]++
											if !yaml_emitter_write_indicator(emitter, []byte{'}'}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:528
			_go_fuzz_dep_.CoverTab[125199]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:529
			// _ = "end of CoverTab[125199]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:530
			_go_fuzz_dep_.CoverTab[125200]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:530
			// _ = "end of CoverTab[125200]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:530
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:530
		// _ = "end of CoverTab[125189]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:530
		_go_fuzz_dep_.CoverTab[125190]++
											emitter.state = emitter.states[len(emitter.states)-1]
											emitter.states = emitter.states[:len(emitter.states)-1]
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:533
		// _ = "end of CoverTab[125190]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:534
		_go_fuzz_dep_.CoverTab[125201]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:534
		// _ = "end of CoverTab[125201]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:534
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:534
	// _ = "end of CoverTab[125174]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:534
	_go_fuzz_dep_.CoverTab[125175]++

										if !first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:536
		_go_fuzz_dep_.CoverTab[125202]++
											if !yaml_emitter_write_indicator(emitter, []byte{','}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:537
			_go_fuzz_dep_.CoverTab[125203]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:538
			// _ = "end of CoverTab[125203]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:539
			_go_fuzz_dep_.CoverTab[125204]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:539
			// _ = "end of CoverTab[125204]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:539
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:539
		// _ = "end of CoverTab[125202]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:540
		_go_fuzz_dep_.CoverTab[125205]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:540
		// _ = "end of CoverTab[125205]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:540
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:540
	// _ = "end of CoverTab[125175]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:540
	_go_fuzz_dep_.CoverTab[125176]++
										if emitter.canonical || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:541
		_go_fuzz_dep_.CoverTab[125206]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:541
		return emitter.column > emitter.best_width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:541
		// _ = "end of CoverTab[125206]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:541
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:541
		_go_fuzz_dep_.CoverTab[125207]++
											if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:542
			_go_fuzz_dep_.CoverTab[125208]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:543
			// _ = "end of CoverTab[125208]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:544
			_go_fuzz_dep_.CoverTab[125209]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:544
			// _ = "end of CoverTab[125209]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:544
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:544
		// _ = "end of CoverTab[125207]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:545
		_go_fuzz_dep_.CoverTab[125210]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:545
		// _ = "end of CoverTab[125210]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:545
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:545
	// _ = "end of CoverTab[125176]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:545
	_go_fuzz_dep_.CoverTab[125177]++

										if !emitter.canonical && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:547
		_go_fuzz_dep_.CoverTab[125211]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:547
		return yaml_emitter_check_simple_key(emitter)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:547
		// _ = "end of CoverTab[125211]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:547
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:547
		_go_fuzz_dep_.CoverTab[125212]++
											emitter.states = append(emitter.states, yaml_EMIT_FLOW_MAPPING_SIMPLE_VALUE_STATE)
											return yaml_emitter_emit_node(emitter, event, false, false, true, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:549
		// _ = "end of CoverTab[125212]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:550
		_go_fuzz_dep_.CoverTab[125213]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:550
		// _ = "end of CoverTab[125213]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:550
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:550
	// _ = "end of CoverTab[125177]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:550
	_go_fuzz_dep_.CoverTab[125178]++
										if !yaml_emitter_write_indicator(emitter, []byte{'?'}, true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:551
		_go_fuzz_dep_.CoverTab[125214]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:552
		// _ = "end of CoverTab[125214]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:553
		_go_fuzz_dep_.CoverTab[125215]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:553
		// _ = "end of CoverTab[125215]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:553
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:553
	// _ = "end of CoverTab[125178]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:553
	_go_fuzz_dep_.CoverTab[125179]++
										emitter.states = append(emitter.states, yaml_EMIT_FLOW_MAPPING_VALUE_STATE)
										return yaml_emitter_emit_node(emitter, event, false, false, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:555
	// _ = "end of CoverTab[125179]"
}

// Expect a flow value node.
func yaml_emitter_emit_flow_mapping_value(emitter *yaml_emitter_t, event *yaml_event_t, simple bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:559
	_go_fuzz_dep_.CoverTab[125216]++
										if simple {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:560
		_go_fuzz_dep_.CoverTab[125218]++
											if !yaml_emitter_write_indicator(emitter, []byte{':'}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:561
			_go_fuzz_dep_.CoverTab[125219]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:562
			// _ = "end of CoverTab[125219]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:563
			_go_fuzz_dep_.CoverTab[125220]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:563
			// _ = "end of CoverTab[125220]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:563
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:563
		// _ = "end of CoverTab[125218]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:564
		_go_fuzz_dep_.CoverTab[125221]++
											if emitter.canonical || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:565
			_go_fuzz_dep_.CoverTab[125223]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:565
			return emitter.column > emitter.best_width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:565
			// _ = "end of CoverTab[125223]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:565
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:565
			_go_fuzz_dep_.CoverTab[125224]++
												if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:566
				_go_fuzz_dep_.CoverTab[125225]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:567
				// _ = "end of CoverTab[125225]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:568
				_go_fuzz_dep_.CoverTab[125226]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:568
				// _ = "end of CoverTab[125226]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:568
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:568
			// _ = "end of CoverTab[125224]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:569
			_go_fuzz_dep_.CoverTab[125227]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:569
			// _ = "end of CoverTab[125227]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:569
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:569
		// _ = "end of CoverTab[125221]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:569
		_go_fuzz_dep_.CoverTab[125222]++
											if !yaml_emitter_write_indicator(emitter, []byte{':'}, true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:570
			_go_fuzz_dep_.CoverTab[125228]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:571
			// _ = "end of CoverTab[125228]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:572
			_go_fuzz_dep_.CoverTab[125229]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:572
			// _ = "end of CoverTab[125229]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:572
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:572
		// _ = "end of CoverTab[125222]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:573
	// _ = "end of CoverTab[125216]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:573
	_go_fuzz_dep_.CoverTab[125217]++
										emitter.states = append(emitter.states, yaml_EMIT_FLOW_MAPPING_KEY_STATE)
										return yaml_emitter_emit_node(emitter, event, false, false, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:575
	// _ = "end of CoverTab[125217]"
}

// Expect a block item node.
func yaml_emitter_emit_block_sequence_item(emitter *yaml_emitter_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:579
	_go_fuzz_dep_.CoverTab[125230]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:580
		_go_fuzz_dep_.CoverTab[125235]++
											if !yaml_emitter_increase_indent(emitter, false, emitter.mapping_context && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:581
			_go_fuzz_dep_.CoverTab[125236]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:581
			return !emitter.indention
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:581
			// _ = "end of CoverTab[125236]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:581
		}()) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:581
			_go_fuzz_dep_.CoverTab[125237]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:582
			// _ = "end of CoverTab[125237]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:583
			_go_fuzz_dep_.CoverTab[125238]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:583
			// _ = "end of CoverTab[125238]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:583
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:583
		// _ = "end of CoverTab[125235]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:584
		_go_fuzz_dep_.CoverTab[125239]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:584
		// _ = "end of CoverTab[125239]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:584
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:584
	// _ = "end of CoverTab[125230]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:584
	_go_fuzz_dep_.CoverTab[125231]++
										if event.typ == yaml_SEQUENCE_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:585
		_go_fuzz_dep_.CoverTab[125240]++
											emitter.indent = emitter.indents[len(emitter.indents)-1]
											emitter.indents = emitter.indents[:len(emitter.indents)-1]
											emitter.state = emitter.states[len(emitter.states)-1]
											emitter.states = emitter.states[:len(emitter.states)-1]
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:590
		// _ = "end of CoverTab[125240]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:591
		_go_fuzz_dep_.CoverTab[125241]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:591
		// _ = "end of CoverTab[125241]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:591
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:591
	// _ = "end of CoverTab[125231]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:591
	_go_fuzz_dep_.CoverTab[125232]++
										if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:592
		_go_fuzz_dep_.CoverTab[125242]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:593
		// _ = "end of CoverTab[125242]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:594
		_go_fuzz_dep_.CoverTab[125243]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:594
		// _ = "end of CoverTab[125243]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:594
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:594
	// _ = "end of CoverTab[125232]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:594
	_go_fuzz_dep_.CoverTab[125233]++
										if !yaml_emitter_write_indicator(emitter, []byte{'-'}, true, false, true) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:595
		_go_fuzz_dep_.CoverTab[125244]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:596
		// _ = "end of CoverTab[125244]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:597
		_go_fuzz_dep_.CoverTab[125245]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:597
		// _ = "end of CoverTab[125245]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:597
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:597
	// _ = "end of CoverTab[125233]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:597
	_go_fuzz_dep_.CoverTab[125234]++
										emitter.states = append(emitter.states, yaml_EMIT_BLOCK_SEQUENCE_ITEM_STATE)
										return yaml_emitter_emit_node(emitter, event, false, true, false, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:599
	// _ = "end of CoverTab[125234]"
}

// Expect a block key node.
func yaml_emitter_emit_block_mapping_key(emitter *yaml_emitter_t, event *yaml_event_t, first bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:603
	_go_fuzz_dep_.CoverTab[125246]++
										if first {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:604
		_go_fuzz_dep_.CoverTab[125252]++
											if !yaml_emitter_increase_indent(emitter, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:605
			_go_fuzz_dep_.CoverTab[125253]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:606
			// _ = "end of CoverTab[125253]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:607
			_go_fuzz_dep_.CoverTab[125254]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:607
			// _ = "end of CoverTab[125254]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:607
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:607
		// _ = "end of CoverTab[125252]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:608
		_go_fuzz_dep_.CoverTab[125255]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:608
		// _ = "end of CoverTab[125255]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:608
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:608
	// _ = "end of CoverTab[125246]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:608
	_go_fuzz_dep_.CoverTab[125247]++
										if event.typ == yaml_MAPPING_END_EVENT {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:609
		_go_fuzz_dep_.CoverTab[125256]++
											emitter.indent = emitter.indents[len(emitter.indents)-1]
											emitter.indents = emitter.indents[:len(emitter.indents)-1]
											emitter.state = emitter.states[len(emitter.states)-1]
											emitter.states = emitter.states[:len(emitter.states)-1]
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:614
		// _ = "end of CoverTab[125256]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:615
		_go_fuzz_dep_.CoverTab[125257]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:615
		// _ = "end of CoverTab[125257]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:615
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:615
	// _ = "end of CoverTab[125247]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:615
	_go_fuzz_dep_.CoverTab[125248]++
										if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:616
		_go_fuzz_dep_.CoverTab[125258]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:617
		// _ = "end of CoverTab[125258]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:618
		_go_fuzz_dep_.CoverTab[125259]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:618
		// _ = "end of CoverTab[125259]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:618
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:618
	// _ = "end of CoverTab[125248]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:618
	_go_fuzz_dep_.CoverTab[125249]++
										if yaml_emitter_check_simple_key(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:619
		_go_fuzz_dep_.CoverTab[125260]++
											emitter.states = append(emitter.states, yaml_EMIT_BLOCK_MAPPING_SIMPLE_VALUE_STATE)
											return yaml_emitter_emit_node(emitter, event, false, false, true, true)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:621
		// _ = "end of CoverTab[125260]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:622
		_go_fuzz_dep_.CoverTab[125261]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:622
		// _ = "end of CoverTab[125261]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:622
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:622
	// _ = "end of CoverTab[125249]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:622
	_go_fuzz_dep_.CoverTab[125250]++
										if !yaml_emitter_write_indicator(emitter, []byte{'?'}, true, false, true) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:623
		_go_fuzz_dep_.CoverTab[125262]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:624
		// _ = "end of CoverTab[125262]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:625
		_go_fuzz_dep_.CoverTab[125263]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:625
		// _ = "end of CoverTab[125263]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:625
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:625
	// _ = "end of CoverTab[125250]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:625
	_go_fuzz_dep_.CoverTab[125251]++
										emitter.states = append(emitter.states, yaml_EMIT_BLOCK_MAPPING_VALUE_STATE)
										return yaml_emitter_emit_node(emitter, event, false, false, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:627
	// _ = "end of CoverTab[125251]"
}

// Expect a block value node.
func yaml_emitter_emit_block_mapping_value(emitter *yaml_emitter_t, event *yaml_event_t, simple bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:631
	_go_fuzz_dep_.CoverTab[125264]++
										if simple {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:632
		_go_fuzz_dep_.CoverTab[125266]++
											if !yaml_emitter_write_indicator(emitter, []byte{':'}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:633
			_go_fuzz_dep_.CoverTab[125267]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:634
			// _ = "end of CoverTab[125267]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:635
			_go_fuzz_dep_.CoverTab[125268]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:635
			// _ = "end of CoverTab[125268]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:635
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:635
		// _ = "end of CoverTab[125266]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:636
		_go_fuzz_dep_.CoverTab[125269]++
											if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:637
			_go_fuzz_dep_.CoverTab[125271]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:638
			// _ = "end of CoverTab[125271]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:639
			_go_fuzz_dep_.CoverTab[125272]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:639
			// _ = "end of CoverTab[125272]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:639
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:639
		// _ = "end of CoverTab[125269]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:639
		_go_fuzz_dep_.CoverTab[125270]++
											if !yaml_emitter_write_indicator(emitter, []byte{':'}, true, false, true) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:640
			_go_fuzz_dep_.CoverTab[125273]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:641
			// _ = "end of CoverTab[125273]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:642
			_go_fuzz_dep_.CoverTab[125274]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:642
			// _ = "end of CoverTab[125274]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:642
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:642
		// _ = "end of CoverTab[125270]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:643
	// _ = "end of CoverTab[125264]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:643
	_go_fuzz_dep_.CoverTab[125265]++
										emitter.states = append(emitter.states, yaml_EMIT_BLOCK_MAPPING_KEY_STATE)
										return yaml_emitter_emit_node(emitter, event, false, false, true, false)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:645
	// _ = "end of CoverTab[125265]"
}

// Expect a node.
func yaml_emitter_emit_node(emitter *yaml_emitter_t, event *yaml_event_t,
	root bool, sequence bool, mapping bool, simple_key bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:650
	_go_fuzz_dep_.CoverTab[125275]++

										emitter.root_context = root
										emitter.sequence_context = sequence
										emitter.mapping_context = mapping
										emitter.simple_key_context = simple_key

										switch event.typ {
	case yaml_ALIAS_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:658
		_go_fuzz_dep_.CoverTab[125276]++
											return yaml_emitter_emit_alias(emitter, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:659
		// _ = "end of CoverTab[125276]"
	case yaml_SCALAR_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:660
		_go_fuzz_dep_.CoverTab[125277]++
											return yaml_emitter_emit_scalar(emitter, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:661
		// _ = "end of CoverTab[125277]"
	case yaml_SEQUENCE_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:662
		_go_fuzz_dep_.CoverTab[125278]++
											return yaml_emitter_emit_sequence_start(emitter, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:663
		// _ = "end of CoverTab[125278]"
	case yaml_MAPPING_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:664
		_go_fuzz_dep_.CoverTab[125279]++
											return yaml_emitter_emit_mapping_start(emitter, event)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:665
		// _ = "end of CoverTab[125279]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:666
		_go_fuzz_dep_.CoverTab[125280]++
											return yaml_emitter_set_emitter_error(emitter,
			fmt.Sprintf("expected SCALAR, SEQUENCE-START, MAPPING-START, or ALIAS, but got %v", event.typ))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:668
		// _ = "end of CoverTab[125280]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:669
	// _ = "end of CoverTab[125275]"
}

// Expect ALIAS.
func yaml_emitter_emit_alias(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:673
	_go_fuzz_dep_.CoverTab[125281]++
										if !yaml_emitter_process_anchor(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:674
		_go_fuzz_dep_.CoverTab[125283]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:675
		// _ = "end of CoverTab[125283]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:676
		_go_fuzz_dep_.CoverTab[125284]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:676
		// _ = "end of CoverTab[125284]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:676
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:676
	// _ = "end of CoverTab[125281]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:676
	_go_fuzz_dep_.CoverTab[125282]++
										emitter.state = emitter.states[len(emitter.states)-1]
										emitter.states = emitter.states[:len(emitter.states)-1]
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:679
	// _ = "end of CoverTab[125282]"
}

// Expect SCALAR.
func yaml_emitter_emit_scalar(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:683
	_go_fuzz_dep_.CoverTab[125285]++
										if !yaml_emitter_select_scalar_style(emitter, event) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:684
		_go_fuzz_dep_.CoverTab[125291]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:685
		// _ = "end of CoverTab[125291]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:686
		_go_fuzz_dep_.CoverTab[125292]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:686
		// _ = "end of CoverTab[125292]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:686
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:686
	// _ = "end of CoverTab[125285]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:686
	_go_fuzz_dep_.CoverTab[125286]++
										if !yaml_emitter_process_anchor(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:687
		_go_fuzz_dep_.CoverTab[125293]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:688
		// _ = "end of CoverTab[125293]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:689
		_go_fuzz_dep_.CoverTab[125294]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:689
		// _ = "end of CoverTab[125294]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:689
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:689
	// _ = "end of CoverTab[125286]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:689
	_go_fuzz_dep_.CoverTab[125287]++
										if !yaml_emitter_process_tag(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:690
		_go_fuzz_dep_.CoverTab[125295]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:691
		// _ = "end of CoverTab[125295]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:692
		_go_fuzz_dep_.CoverTab[125296]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:692
		// _ = "end of CoverTab[125296]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:692
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:692
	// _ = "end of CoverTab[125287]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:692
	_go_fuzz_dep_.CoverTab[125288]++
										if !yaml_emitter_increase_indent(emitter, true, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:693
		_go_fuzz_dep_.CoverTab[125297]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:694
		// _ = "end of CoverTab[125297]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:695
		_go_fuzz_dep_.CoverTab[125298]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:695
		// _ = "end of CoverTab[125298]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:695
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:695
	// _ = "end of CoverTab[125288]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:695
	_go_fuzz_dep_.CoverTab[125289]++
										if !yaml_emitter_process_scalar(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:696
		_go_fuzz_dep_.CoverTab[125299]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:697
		// _ = "end of CoverTab[125299]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:698
		_go_fuzz_dep_.CoverTab[125300]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:698
		// _ = "end of CoverTab[125300]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:698
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:698
	// _ = "end of CoverTab[125289]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:698
	_go_fuzz_dep_.CoverTab[125290]++
										emitter.indent = emitter.indents[len(emitter.indents)-1]
										emitter.indents = emitter.indents[:len(emitter.indents)-1]
										emitter.state = emitter.states[len(emitter.states)-1]
										emitter.states = emitter.states[:len(emitter.states)-1]
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:703
	// _ = "end of CoverTab[125290]"
}

// Expect SEQUENCE-START.
func yaml_emitter_emit_sequence_start(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:707
	_go_fuzz_dep_.CoverTab[125301]++
										if !yaml_emitter_process_anchor(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:708
		_go_fuzz_dep_.CoverTab[125305]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:709
		// _ = "end of CoverTab[125305]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:710
		_go_fuzz_dep_.CoverTab[125306]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:710
		// _ = "end of CoverTab[125306]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:710
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:710
	// _ = "end of CoverTab[125301]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:710
	_go_fuzz_dep_.CoverTab[125302]++
										if !yaml_emitter_process_tag(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:711
		_go_fuzz_dep_.CoverTab[125307]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:712
		// _ = "end of CoverTab[125307]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:713
		_go_fuzz_dep_.CoverTab[125308]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:713
		// _ = "end of CoverTab[125308]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:713
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:713
	// _ = "end of CoverTab[125302]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:713
	_go_fuzz_dep_.CoverTab[125303]++
										if emitter.flow_level > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		_go_fuzz_dep_.CoverTab[125309]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		return emitter.canonical
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		// _ = "end of CoverTab[125309]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		_go_fuzz_dep_.CoverTab[125310]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		return event.sequence_style() == yaml_FLOW_SEQUENCE_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		// _ = "end of CoverTab[125310]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		_go_fuzz_dep_.CoverTab[125311]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:714
		return yaml_emitter_check_empty_sequence(emitter)
											// _ = "end of CoverTab[125311]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:715
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:715
		_go_fuzz_dep_.CoverTab[125312]++
											emitter.state = yaml_EMIT_FLOW_SEQUENCE_FIRST_ITEM_STATE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:716
		// _ = "end of CoverTab[125312]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:717
		_go_fuzz_dep_.CoverTab[125313]++
											emitter.state = yaml_EMIT_BLOCK_SEQUENCE_FIRST_ITEM_STATE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:718
		// _ = "end of CoverTab[125313]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:719
	// _ = "end of CoverTab[125303]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:719
	_go_fuzz_dep_.CoverTab[125304]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:720
	// _ = "end of CoverTab[125304]"
}

// Expect MAPPING-START.
func yaml_emitter_emit_mapping_start(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:724
	_go_fuzz_dep_.CoverTab[125314]++
										if !yaml_emitter_process_anchor(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:725
		_go_fuzz_dep_.CoverTab[125318]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:726
		// _ = "end of CoverTab[125318]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:727
		_go_fuzz_dep_.CoverTab[125319]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:727
		// _ = "end of CoverTab[125319]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:727
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:727
	// _ = "end of CoverTab[125314]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:727
	_go_fuzz_dep_.CoverTab[125315]++
										if !yaml_emitter_process_tag(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:728
		_go_fuzz_dep_.CoverTab[125320]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:729
		// _ = "end of CoverTab[125320]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:730
		_go_fuzz_dep_.CoverTab[125321]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:730
		// _ = "end of CoverTab[125321]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:730
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:730
	// _ = "end of CoverTab[125315]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:730
	_go_fuzz_dep_.CoverTab[125316]++
										if emitter.flow_level > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		_go_fuzz_dep_.CoverTab[125322]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		return emitter.canonical
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		// _ = "end of CoverTab[125322]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		_go_fuzz_dep_.CoverTab[125323]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		return event.mapping_style() == yaml_FLOW_MAPPING_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		// _ = "end of CoverTab[125323]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		_go_fuzz_dep_.CoverTab[125324]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:731
		return yaml_emitter_check_empty_mapping(emitter)
											// _ = "end of CoverTab[125324]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:732
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:732
		_go_fuzz_dep_.CoverTab[125325]++
											emitter.state = yaml_EMIT_FLOW_MAPPING_FIRST_KEY_STATE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:733
		// _ = "end of CoverTab[125325]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:734
		_go_fuzz_dep_.CoverTab[125326]++
											emitter.state = yaml_EMIT_BLOCK_MAPPING_FIRST_KEY_STATE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:735
		// _ = "end of CoverTab[125326]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:736
	// _ = "end of CoverTab[125316]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:736
	_go_fuzz_dep_.CoverTab[125317]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:737
	// _ = "end of CoverTab[125317]"
}

// Check if the document content is an empty scalar.
func yaml_emitter_check_empty_document(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:741
	_go_fuzz_dep_.CoverTab[125327]++
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:742
	// _ = "end of CoverTab[125327]"
}

// Check if the next events represent an empty sequence.
func yaml_emitter_check_empty_sequence(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:746
	_go_fuzz_dep_.CoverTab[125328]++
										if len(emitter.events)-emitter.events_head < 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:747
		_go_fuzz_dep_.CoverTab[125330]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:748
		// _ = "end of CoverTab[125330]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:749
		_go_fuzz_dep_.CoverTab[125331]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:749
		// _ = "end of CoverTab[125331]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:749
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:749
	// _ = "end of CoverTab[125328]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:749
	_go_fuzz_dep_.CoverTab[125329]++
										return emitter.events[emitter.events_head].typ == yaml_SEQUENCE_START_EVENT && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:750
		_go_fuzz_dep_.CoverTab[125332]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:750
		return emitter.events[emitter.events_head+1].typ == yaml_SEQUENCE_END_EVENT
											// _ = "end of CoverTab[125332]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:751
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:751
	// _ = "end of CoverTab[125329]"
}

// Check if the next events represent an empty mapping.
func yaml_emitter_check_empty_mapping(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:755
	_go_fuzz_dep_.CoverTab[125333]++
										if len(emitter.events)-emitter.events_head < 2 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:756
		_go_fuzz_dep_.CoverTab[125335]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:757
		// _ = "end of CoverTab[125335]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:758
		_go_fuzz_dep_.CoverTab[125336]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:758
		// _ = "end of CoverTab[125336]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:758
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:758
	// _ = "end of CoverTab[125333]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:758
	_go_fuzz_dep_.CoverTab[125334]++
										return emitter.events[emitter.events_head].typ == yaml_MAPPING_START_EVENT && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:759
		_go_fuzz_dep_.CoverTab[125337]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:759
		return emitter.events[emitter.events_head+1].typ == yaml_MAPPING_END_EVENT
											// _ = "end of CoverTab[125337]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:760
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:760
	// _ = "end of CoverTab[125334]"
}

// Check if the next node can be expressed as a simple key.
func yaml_emitter_check_simple_key(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:764
	_go_fuzz_dep_.CoverTab[125338]++
										length := 0
										switch emitter.events[emitter.events_head].typ {
	case yaml_ALIAS_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:767
		_go_fuzz_dep_.CoverTab[125340]++
											length += len(emitter.anchor_data.anchor)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:768
		// _ = "end of CoverTab[125340]"
	case yaml_SCALAR_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:769
		_go_fuzz_dep_.CoverTab[125341]++
											if emitter.scalar_data.multiline {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:770
			_go_fuzz_dep_.CoverTab[125348]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:771
			// _ = "end of CoverTab[125348]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:772
			_go_fuzz_dep_.CoverTab[125349]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:772
			// _ = "end of CoverTab[125349]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:772
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:772
		// _ = "end of CoverTab[125341]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:772
		_go_fuzz_dep_.CoverTab[125342]++
											length += len(emitter.anchor_data.anchor) +
			len(emitter.tag_data.handle) +
			len(emitter.tag_data.suffix) +
			len(emitter.scalar_data.value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:776
		// _ = "end of CoverTab[125342]"
	case yaml_SEQUENCE_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:777
		_go_fuzz_dep_.CoverTab[125343]++
											if !yaml_emitter_check_empty_sequence(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:778
			_go_fuzz_dep_.CoverTab[125350]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:779
			// _ = "end of CoverTab[125350]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:780
			_go_fuzz_dep_.CoverTab[125351]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:780
			// _ = "end of CoverTab[125351]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:780
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:780
		// _ = "end of CoverTab[125343]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:780
		_go_fuzz_dep_.CoverTab[125344]++
											length += len(emitter.anchor_data.anchor) +
			len(emitter.tag_data.handle) +
			len(emitter.tag_data.suffix)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:783
		// _ = "end of CoverTab[125344]"
	case yaml_MAPPING_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:784
		_go_fuzz_dep_.CoverTab[125345]++
											if !yaml_emitter_check_empty_mapping(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:785
			_go_fuzz_dep_.CoverTab[125352]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:786
			// _ = "end of CoverTab[125352]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:787
			_go_fuzz_dep_.CoverTab[125353]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:787
			// _ = "end of CoverTab[125353]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:787
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:787
		// _ = "end of CoverTab[125345]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:787
		_go_fuzz_dep_.CoverTab[125346]++
											length += len(emitter.anchor_data.anchor) +
			len(emitter.tag_data.handle) +
			len(emitter.tag_data.suffix)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:790
		// _ = "end of CoverTab[125346]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:791
		_go_fuzz_dep_.CoverTab[125347]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:792
		// _ = "end of CoverTab[125347]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:793
	// _ = "end of CoverTab[125338]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:793
	_go_fuzz_dep_.CoverTab[125339]++
										return length <= 128
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:794
	// _ = "end of CoverTab[125339]"
}

// Determine an acceptable scalar style.
func yaml_emitter_select_scalar_style(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:798
	_go_fuzz_dep_.CoverTab[125354]++

										no_tag := len(emitter.tag_data.handle) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:800
		_go_fuzz_dep_.CoverTab[125363]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:800
		return len(emitter.tag_data.suffix) == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:800
		// _ = "end of CoverTab[125363]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:800
	}()
										if no_tag && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
		_go_fuzz_dep_.CoverTab[125364]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
		return !event.implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
		// _ = "end of CoverTab[125364]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
		_go_fuzz_dep_.CoverTab[125365]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
		return !event.quoted_implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
		// _ = "end of CoverTab[125365]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:801
		_go_fuzz_dep_.CoverTab[125366]++
											return yaml_emitter_set_emitter_error(emitter, "neither tag nor implicit flags are specified")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:802
		// _ = "end of CoverTab[125366]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:803
		_go_fuzz_dep_.CoverTab[125367]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:803
		// _ = "end of CoverTab[125367]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:803
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:803
	// _ = "end of CoverTab[125354]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:803
	_go_fuzz_dep_.CoverTab[125355]++

										style := event.scalar_style()
										if style == yaml_ANY_SCALAR_STYLE {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:806
		_go_fuzz_dep_.CoverTab[125368]++
											style = yaml_PLAIN_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:807
		// _ = "end of CoverTab[125368]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:808
		_go_fuzz_dep_.CoverTab[125369]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:808
		// _ = "end of CoverTab[125369]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:808
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:808
	// _ = "end of CoverTab[125355]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:808
	_go_fuzz_dep_.CoverTab[125356]++
										if emitter.canonical {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:809
		_go_fuzz_dep_.CoverTab[125370]++
											style = yaml_DOUBLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:810
		// _ = "end of CoverTab[125370]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:811
		_go_fuzz_dep_.CoverTab[125371]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:811
		// _ = "end of CoverTab[125371]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:811
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:811
	// _ = "end of CoverTab[125356]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:811
	_go_fuzz_dep_.CoverTab[125357]++
										if emitter.simple_key_context && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:812
		_go_fuzz_dep_.CoverTab[125372]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:812
		return emitter.scalar_data.multiline
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:812
		// _ = "end of CoverTab[125372]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:812
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:812
		_go_fuzz_dep_.CoverTab[125373]++
											style = yaml_DOUBLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:813
		// _ = "end of CoverTab[125373]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:814
		_go_fuzz_dep_.CoverTab[125374]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:814
		// _ = "end of CoverTab[125374]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:814
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:814
	// _ = "end of CoverTab[125357]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:814
	_go_fuzz_dep_.CoverTab[125358]++

										if style == yaml_PLAIN_SCALAR_STYLE {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:816
		_go_fuzz_dep_.CoverTab[125375]++
											if emitter.flow_level > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:817
			_go_fuzz_dep_.CoverTab[125378]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:817
			return !emitter.scalar_data.flow_plain_allowed
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:817
			// _ = "end of CoverTab[125378]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:817
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:817
			_go_fuzz_dep_.CoverTab[125379]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:817
			return emitter.flow_level == 0 && func() bool {
													_go_fuzz_dep_.CoverTab[125380]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:818
				return !emitter.scalar_data.block_plain_allowed
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:818
				// _ = "end of CoverTab[125380]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:818
			}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:818
			// _ = "end of CoverTab[125379]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:818
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:818
			_go_fuzz_dep_.CoverTab[125381]++
												style = yaml_SINGLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:819
			// _ = "end of CoverTab[125381]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:820
			_go_fuzz_dep_.CoverTab[125382]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:820
			// _ = "end of CoverTab[125382]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:820
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:820
		// _ = "end of CoverTab[125375]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:820
		_go_fuzz_dep_.CoverTab[125376]++
											if len(emitter.scalar_data.value) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
			_go_fuzz_dep_.CoverTab[125383]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
			return (emitter.flow_level > 0 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
				_go_fuzz_dep_.CoverTab[125384]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
				return emitter.simple_key_context
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
				// _ = "end of CoverTab[125384]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
			// _ = "end of CoverTab[125383]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:821
			_go_fuzz_dep_.CoverTab[125385]++
												style = yaml_SINGLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:822
			// _ = "end of CoverTab[125385]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:823
			_go_fuzz_dep_.CoverTab[125386]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:823
			// _ = "end of CoverTab[125386]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:823
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:823
		// _ = "end of CoverTab[125376]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:823
		_go_fuzz_dep_.CoverTab[125377]++
											if no_tag && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:824
			_go_fuzz_dep_.CoverTab[125387]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:824
			return !event.implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:824
			// _ = "end of CoverTab[125387]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:824
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:824
			_go_fuzz_dep_.CoverTab[125388]++
												style = yaml_SINGLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:825
			// _ = "end of CoverTab[125388]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:826
			_go_fuzz_dep_.CoverTab[125389]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:826
			// _ = "end of CoverTab[125389]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:826
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:826
		// _ = "end of CoverTab[125377]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:827
		_go_fuzz_dep_.CoverTab[125390]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:827
		// _ = "end of CoverTab[125390]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:827
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:827
	// _ = "end of CoverTab[125358]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:827
	_go_fuzz_dep_.CoverTab[125359]++
										if style == yaml_SINGLE_QUOTED_SCALAR_STYLE {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:828
		_go_fuzz_dep_.CoverTab[125391]++
											if !emitter.scalar_data.single_quoted_allowed {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:829
			_go_fuzz_dep_.CoverTab[125392]++
												style = yaml_DOUBLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:830
			// _ = "end of CoverTab[125392]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:831
			_go_fuzz_dep_.CoverTab[125393]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:831
			// _ = "end of CoverTab[125393]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:831
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:831
		// _ = "end of CoverTab[125391]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:832
		_go_fuzz_dep_.CoverTab[125394]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:832
		// _ = "end of CoverTab[125394]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:832
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:832
	// _ = "end of CoverTab[125359]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:832
	_go_fuzz_dep_.CoverTab[125360]++
										if style == yaml_LITERAL_SCALAR_STYLE || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:833
		_go_fuzz_dep_.CoverTab[125395]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:833
		return style == yaml_FOLDED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:833
		// _ = "end of CoverTab[125395]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:833
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:833
		_go_fuzz_dep_.CoverTab[125396]++
											if !emitter.scalar_data.block_allowed || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
			_go_fuzz_dep_.CoverTab[125397]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
			return emitter.flow_level > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
			// _ = "end of CoverTab[125397]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
			_go_fuzz_dep_.CoverTab[125398]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
			return emitter.simple_key_context
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
			// _ = "end of CoverTab[125398]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:834
			_go_fuzz_dep_.CoverTab[125399]++
												style = yaml_DOUBLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:835
			// _ = "end of CoverTab[125399]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:836
			_go_fuzz_dep_.CoverTab[125400]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:836
			// _ = "end of CoverTab[125400]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:836
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:836
		// _ = "end of CoverTab[125396]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:837
		_go_fuzz_dep_.CoverTab[125401]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:837
		// _ = "end of CoverTab[125401]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:837
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:837
	// _ = "end of CoverTab[125360]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:837
	_go_fuzz_dep_.CoverTab[125361]++

										if no_tag && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
		_go_fuzz_dep_.CoverTab[125402]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
		return !event.quoted_implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
		// _ = "end of CoverTab[125402]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
		_go_fuzz_dep_.CoverTab[125403]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
		return style != yaml_PLAIN_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
		// _ = "end of CoverTab[125403]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:839
		_go_fuzz_dep_.CoverTab[125404]++
											emitter.tag_data.handle = []byte{'!'}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:840
		// _ = "end of CoverTab[125404]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:841
		_go_fuzz_dep_.CoverTab[125405]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:841
		// _ = "end of CoverTab[125405]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:841
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:841
	// _ = "end of CoverTab[125361]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:841
	_go_fuzz_dep_.CoverTab[125362]++
										emitter.scalar_data.style = style
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:843
	// _ = "end of CoverTab[125362]"
}

// Write an anchor.
func yaml_emitter_process_anchor(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:847
	_go_fuzz_dep_.CoverTab[125406]++
										if emitter.anchor_data.anchor == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:848
		_go_fuzz_dep_.CoverTab[125410]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:849
		// _ = "end of CoverTab[125410]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:850
		_go_fuzz_dep_.CoverTab[125411]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:850
		// _ = "end of CoverTab[125411]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:850
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:850
	// _ = "end of CoverTab[125406]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:850
	_go_fuzz_dep_.CoverTab[125407]++
										c := []byte{'&'}
										if emitter.anchor_data.alias {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:852
		_go_fuzz_dep_.CoverTab[125412]++
											c[0] = '*'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:853
		// _ = "end of CoverTab[125412]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:854
		_go_fuzz_dep_.CoverTab[125413]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:854
		// _ = "end of CoverTab[125413]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:854
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:854
	// _ = "end of CoverTab[125407]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:854
	_go_fuzz_dep_.CoverTab[125408]++
										if !yaml_emitter_write_indicator(emitter, c, true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:855
		_go_fuzz_dep_.CoverTab[125414]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:856
		// _ = "end of CoverTab[125414]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:857
		_go_fuzz_dep_.CoverTab[125415]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:857
		// _ = "end of CoverTab[125415]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:857
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:857
	// _ = "end of CoverTab[125408]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:857
	_go_fuzz_dep_.CoverTab[125409]++
										return yaml_emitter_write_anchor(emitter, emitter.anchor_data.anchor)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:858
	// _ = "end of CoverTab[125409]"
}

// Write a tag.
func yaml_emitter_process_tag(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:862
	_go_fuzz_dep_.CoverTab[125416]++
										if len(emitter.tag_data.handle) == 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:863
		_go_fuzz_dep_.CoverTab[125419]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:863
		return len(emitter.tag_data.suffix) == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:863
		// _ = "end of CoverTab[125419]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:863
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:863
		_go_fuzz_dep_.CoverTab[125420]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:864
		// _ = "end of CoverTab[125420]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:865
		_go_fuzz_dep_.CoverTab[125421]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:865
		// _ = "end of CoverTab[125421]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:865
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:865
	// _ = "end of CoverTab[125416]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:865
	_go_fuzz_dep_.CoverTab[125417]++
										if len(emitter.tag_data.handle) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:866
		_go_fuzz_dep_.CoverTab[125422]++
											if !yaml_emitter_write_tag_handle(emitter, emitter.tag_data.handle) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:867
			_go_fuzz_dep_.CoverTab[125424]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:868
			// _ = "end of CoverTab[125424]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:869
			_go_fuzz_dep_.CoverTab[125425]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:869
			// _ = "end of CoverTab[125425]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:869
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:869
		// _ = "end of CoverTab[125422]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:869
		_go_fuzz_dep_.CoverTab[125423]++
											if len(emitter.tag_data.suffix) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:870
			_go_fuzz_dep_.CoverTab[125426]++
												if !yaml_emitter_write_tag_content(emitter, emitter.tag_data.suffix, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:871
				_go_fuzz_dep_.CoverTab[125427]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:872
				// _ = "end of CoverTab[125427]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:873
				_go_fuzz_dep_.CoverTab[125428]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:873
				// _ = "end of CoverTab[125428]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:873
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:873
			// _ = "end of CoverTab[125426]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:874
			_go_fuzz_dep_.CoverTab[125429]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:874
			// _ = "end of CoverTab[125429]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:874
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:874
		// _ = "end of CoverTab[125423]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:875
		_go_fuzz_dep_.CoverTab[125430]++

											if !yaml_emitter_write_indicator(emitter, []byte("!<"), true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:877
			_go_fuzz_dep_.CoverTab[125433]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:878
			// _ = "end of CoverTab[125433]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:879
			_go_fuzz_dep_.CoverTab[125434]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:879
			// _ = "end of CoverTab[125434]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:879
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:879
		// _ = "end of CoverTab[125430]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:879
		_go_fuzz_dep_.CoverTab[125431]++
											if !yaml_emitter_write_tag_content(emitter, emitter.tag_data.suffix, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:880
			_go_fuzz_dep_.CoverTab[125435]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:881
			// _ = "end of CoverTab[125435]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:882
			_go_fuzz_dep_.CoverTab[125436]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:882
			// _ = "end of CoverTab[125436]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:882
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:882
		// _ = "end of CoverTab[125431]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:882
		_go_fuzz_dep_.CoverTab[125432]++
											if !yaml_emitter_write_indicator(emitter, []byte{'>'}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:883
			_go_fuzz_dep_.CoverTab[125437]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:884
			// _ = "end of CoverTab[125437]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:885
			_go_fuzz_dep_.CoverTab[125438]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:885
			// _ = "end of CoverTab[125438]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:885
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:885
		// _ = "end of CoverTab[125432]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:886
	// _ = "end of CoverTab[125417]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:886
	_go_fuzz_dep_.CoverTab[125418]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:887
	// _ = "end of CoverTab[125418]"
}

// Write a scalar.
func yaml_emitter_process_scalar(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:891
	_go_fuzz_dep_.CoverTab[125439]++
										switch emitter.scalar_data.style {
	case yaml_PLAIN_SCALAR_STYLE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:893
		_go_fuzz_dep_.CoverTab[125441]++
											return yaml_emitter_write_plain_scalar(emitter, emitter.scalar_data.value, !emitter.simple_key_context)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:894
		// _ = "end of CoverTab[125441]"

	case yaml_SINGLE_QUOTED_SCALAR_STYLE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:896
		_go_fuzz_dep_.CoverTab[125442]++
											return yaml_emitter_write_single_quoted_scalar(emitter, emitter.scalar_data.value, !emitter.simple_key_context)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:897
		// _ = "end of CoverTab[125442]"

	case yaml_DOUBLE_QUOTED_SCALAR_STYLE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:899
		_go_fuzz_dep_.CoverTab[125443]++
											return yaml_emitter_write_double_quoted_scalar(emitter, emitter.scalar_data.value, !emitter.simple_key_context)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:900
		// _ = "end of CoverTab[125443]"

	case yaml_LITERAL_SCALAR_STYLE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:902
		_go_fuzz_dep_.CoverTab[125444]++
											return yaml_emitter_write_literal_scalar(emitter, emitter.scalar_data.value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:903
		// _ = "end of CoverTab[125444]"

	case yaml_FOLDED_SCALAR_STYLE:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:905
		_go_fuzz_dep_.CoverTab[125445]++
											return yaml_emitter_write_folded_scalar(emitter, emitter.scalar_data.value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:906
		// _ = "end of CoverTab[125445]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:906
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:906
		_go_fuzz_dep_.CoverTab[125446]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:906
		// _ = "end of CoverTab[125446]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:907
	// _ = "end of CoverTab[125439]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:907
	_go_fuzz_dep_.CoverTab[125440]++
										panic("unknown scalar style")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:908
	// _ = "end of CoverTab[125440]"
}

// Check if a %YAML directive is valid.
func yaml_emitter_analyze_version_directive(emitter *yaml_emitter_t, version_directive *yaml_version_directive_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:912
	_go_fuzz_dep_.CoverTab[125447]++
										if version_directive.major != 1 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:913
		_go_fuzz_dep_.CoverTab[125449]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:913
		return version_directive.minor != 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:913
		// _ = "end of CoverTab[125449]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:913
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:913
		_go_fuzz_dep_.CoverTab[125450]++
											return yaml_emitter_set_emitter_error(emitter, "incompatible %YAML directive")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:914
		// _ = "end of CoverTab[125450]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:915
		_go_fuzz_dep_.CoverTab[125451]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:915
		// _ = "end of CoverTab[125451]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:915
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:915
	// _ = "end of CoverTab[125447]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:915
	_go_fuzz_dep_.CoverTab[125448]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:916
	// _ = "end of CoverTab[125448]"
}

// Check if a %TAG directive is valid.
func yaml_emitter_analyze_tag_directive(emitter *yaml_emitter_t, tag_directive *yaml_tag_directive_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:920
	_go_fuzz_dep_.CoverTab[125452]++
										handle := tag_directive.handle
										prefix := tag_directive.prefix
										if len(handle) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:923
		_go_fuzz_dep_.CoverTab[125458]++
											return yaml_emitter_set_emitter_error(emitter, "tag handle must not be empty")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:924
		// _ = "end of CoverTab[125458]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:925
		_go_fuzz_dep_.CoverTab[125459]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:925
		// _ = "end of CoverTab[125459]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:925
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:925
	// _ = "end of CoverTab[125452]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:925
	_go_fuzz_dep_.CoverTab[125453]++
										if handle[0] != '!' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:926
		_go_fuzz_dep_.CoverTab[125460]++
											return yaml_emitter_set_emitter_error(emitter, "tag handle must start with '!'")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:927
		// _ = "end of CoverTab[125460]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:928
		_go_fuzz_dep_.CoverTab[125461]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:928
		// _ = "end of CoverTab[125461]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:928
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:928
	// _ = "end of CoverTab[125453]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:928
	_go_fuzz_dep_.CoverTab[125454]++
										if handle[len(handle)-1] != '!' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:929
		_go_fuzz_dep_.CoverTab[125462]++
											return yaml_emitter_set_emitter_error(emitter, "tag handle must end with '!'")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:930
		// _ = "end of CoverTab[125462]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:931
		_go_fuzz_dep_.CoverTab[125463]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:931
		// _ = "end of CoverTab[125463]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:931
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:931
	// _ = "end of CoverTab[125454]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:931
	_go_fuzz_dep_.CoverTab[125455]++
										for i := 1; i < len(handle)-1; i += width(handle[i]) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:932
		_go_fuzz_dep_.CoverTab[125464]++
											if !is_alpha(handle, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:933
			_go_fuzz_dep_.CoverTab[125465]++
												return yaml_emitter_set_emitter_error(emitter, "tag handle must contain alphanumerical characters only")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:934
			// _ = "end of CoverTab[125465]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:935
			_go_fuzz_dep_.CoverTab[125466]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:935
			// _ = "end of CoverTab[125466]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:935
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:935
		// _ = "end of CoverTab[125464]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:936
	// _ = "end of CoverTab[125455]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:936
	_go_fuzz_dep_.CoverTab[125456]++
										if len(prefix) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:937
		_go_fuzz_dep_.CoverTab[125467]++
											return yaml_emitter_set_emitter_error(emitter, "tag prefix must not be empty")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:938
		// _ = "end of CoverTab[125467]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:939
		_go_fuzz_dep_.CoverTab[125468]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:939
		// _ = "end of CoverTab[125468]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:939
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:939
	// _ = "end of CoverTab[125456]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:939
	_go_fuzz_dep_.CoverTab[125457]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:940
	// _ = "end of CoverTab[125457]"
}

// Check if an anchor is valid.
func yaml_emitter_analyze_anchor(emitter *yaml_emitter_t, anchor []byte, alias bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:944
	_go_fuzz_dep_.CoverTab[125469]++
										if len(anchor) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:945
		_go_fuzz_dep_.CoverTab[125472]++
											problem := "anchor value must not be empty"
											if alias {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:947
			_go_fuzz_dep_.CoverTab[125474]++
												problem = "alias value must not be empty"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:948
			// _ = "end of CoverTab[125474]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:949
			_go_fuzz_dep_.CoverTab[125475]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:949
			// _ = "end of CoverTab[125475]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:949
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:949
		// _ = "end of CoverTab[125472]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:949
		_go_fuzz_dep_.CoverTab[125473]++
											return yaml_emitter_set_emitter_error(emitter, problem)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:950
		// _ = "end of CoverTab[125473]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:951
		_go_fuzz_dep_.CoverTab[125476]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:951
		// _ = "end of CoverTab[125476]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:951
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:951
	// _ = "end of CoverTab[125469]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:951
	_go_fuzz_dep_.CoverTab[125470]++
										for i := 0; i < len(anchor); i += width(anchor[i]) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:952
		_go_fuzz_dep_.CoverTab[125477]++
											if !is_alpha(anchor, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:953
			_go_fuzz_dep_.CoverTab[125478]++
												problem := "anchor value must contain alphanumerical characters only"
												if alias {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:955
				_go_fuzz_dep_.CoverTab[125480]++
													problem = "alias value must contain alphanumerical characters only"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:956
				// _ = "end of CoverTab[125480]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:957
				_go_fuzz_dep_.CoverTab[125481]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:957
				// _ = "end of CoverTab[125481]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:957
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:957
			// _ = "end of CoverTab[125478]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:957
			_go_fuzz_dep_.CoverTab[125479]++
												return yaml_emitter_set_emitter_error(emitter, problem)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:958
			// _ = "end of CoverTab[125479]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:959
			_go_fuzz_dep_.CoverTab[125482]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:959
			// _ = "end of CoverTab[125482]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:959
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:959
		// _ = "end of CoverTab[125477]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:960
	// _ = "end of CoverTab[125470]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:960
	_go_fuzz_dep_.CoverTab[125471]++
										emitter.anchor_data.anchor = anchor
										emitter.anchor_data.alias = alias
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:963
	// _ = "end of CoverTab[125471]"
}

// Check if a tag is valid.
func yaml_emitter_analyze_tag(emitter *yaml_emitter_t, tag []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:967
	_go_fuzz_dep_.CoverTab[125483]++
										if len(tag) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:968
		_go_fuzz_dep_.CoverTab[125486]++
											return yaml_emitter_set_emitter_error(emitter, "tag value must not be empty")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:969
		// _ = "end of CoverTab[125486]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:970
		_go_fuzz_dep_.CoverTab[125487]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:970
		// _ = "end of CoverTab[125487]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:970
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:970
	// _ = "end of CoverTab[125483]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:970
	_go_fuzz_dep_.CoverTab[125484]++
										for i := 0; i < len(emitter.tag_directives); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:971
		_go_fuzz_dep_.CoverTab[125488]++
											tag_directive := &emitter.tag_directives[i]
											if bytes.HasPrefix(tag, tag_directive.prefix) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:973
			_go_fuzz_dep_.CoverTab[125489]++
												emitter.tag_data.handle = tag_directive.handle
												emitter.tag_data.suffix = tag[len(tag_directive.prefix):]
												return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:976
			// _ = "end of CoverTab[125489]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:977
			_go_fuzz_dep_.CoverTab[125490]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:977
			// _ = "end of CoverTab[125490]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:977
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:977
		// _ = "end of CoverTab[125488]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:978
	// _ = "end of CoverTab[125484]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:978
	_go_fuzz_dep_.CoverTab[125485]++
										emitter.tag_data.suffix = tag
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:980
	// _ = "end of CoverTab[125485]"
}

// Check if a scalar is valid.
func yaml_emitter_analyze_scalar(emitter *yaml_emitter_t, value []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:984
	_go_fuzz_dep_.CoverTab[125491]++
										var (
		block_indicators	= false
		flow_indicators		= false
		line_breaks		= false
		special_characters	= false

		leading_space	= false
		leading_break	= false
		trailing_space	= false
		trailing_break	= false
		break_space	= false
		space_break	= false

		preceded_by_whitespace	= false
		followed_by_whitespace	= false
		previous_space		= false
		previous_break		= false
	)

	emitter.scalar_data.value = value

	if len(value) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1006
		_go_fuzz_dep_.CoverTab[125502]++
											emitter.scalar_data.multiline = false
											emitter.scalar_data.flow_plain_allowed = false
											emitter.scalar_data.block_plain_allowed = true
											emitter.scalar_data.single_quoted_allowed = true
											emitter.scalar_data.block_allowed = false
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1012
		// _ = "end of CoverTab[125502]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1013
		_go_fuzz_dep_.CoverTab[125503]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1013
		// _ = "end of CoverTab[125503]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1013
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1013
	// _ = "end of CoverTab[125491]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1013
	_go_fuzz_dep_.CoverTab[125492]++

										if len(value) >= 3 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
		_go_fuzz_dep_.CoverTab[125504]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
		return ((value[0] == '-' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			_go_fuzz_dep_.CoverTab[125505]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			return value[1] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			// _ = "end of CoverTab[125505]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			_go_fuzz_dep_.CoverTab[125506]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			return value[2] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			// _ = "end of CoverTab[125506]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
		}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			_go_fuzz_dep_.CoverTab[125507]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			return (value[0] == '.' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
				_go_fuzz_dep_.CoverTab[125508]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
				return value[1] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
				// _ = "end of CoverTab[125508]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
				_go_fuzz_dep_.CoverTab[125509]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
				return value[2] == '.'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
				// _ = "end of CoverTab[125509]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
			// _ = "end of CoverTab[125507]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
		// _ = "end of CoverTab[125504]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1015
		_go_fuzz_dep_.CoverTab[125510]++
											block_indicators = true
											flow_indicators = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1017
		// _ = "end of CoverTab[125510]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1018
		_go_fuzz_dep_.CoverTab[125511]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1018
		// _ = "end of CoverTab[125511]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1018
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1018
	// _ = "end of CoverTab[125492]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1018
	_go_fuzz_dep_.CoverTab[125493]++

										preceded_by_whitespace = true
										for i, w := 0, 0; i < len(value); i += w {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1021
		_go_fuzz_dep_.CoverTab[125512]++
											w = width(value[i])
											followed_by_whitespace = i+w >= len(value) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1023
			_go_fuzz_dep_.CoverTab[125516]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1023
			return is_blank(value, i+w)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1023
			// _ = "end of CoverTab[125516]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1023
		}()

											if i == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1025
			_go_fuzz_dep_.CoverTab[125517]++
												switch value[i] {
			case '#', ',', '[', ']', '{', '}', '&', '*', '!', '|', '>', '\'', '"', '%', '@', '`':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1027
				_go_fuzz_dep_.CoverTab[125518]++
													flow_indicators = true
													block_indicators = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1029
				// _ = "end of CoverTab[125518]"
			case '?', ':':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1030
				_go_fuzz_dep_.CoverTab[125519]++
													flow_indicators = true
													if followed_by_whitespace {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1032
					_go_fuzz_dep_.CoverTab[125522]++
														block_indicators = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1033
					// _ = "end of CoverTab[125522]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1034
					_go_fuzz_dep_.CoverTab[125523]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1034
					// _ = "end of CoverTab[125523]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1034
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1034
				// _ = "end of CoverTab[125519]"
			case '-':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1035
				_go_fuzz_dep_.CoverTab[125520]++
													if followed_by_whitespace {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1036
					_go_fuzz_dep_.CoverTab[125524]++
														flow_indicators = true
														block_indicators = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1038
					// _ = "end of CoverTab[125524]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1039
					_go_fuzz_dep_.CoverTab[125525]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1039
					// _ = "end of CoverTab[125525]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1039
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1039
				// _ = "end of CoverTab[125520]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1039
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1039
				_go_fuzz_dep_.CoverTab[125521]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1039
				// _ = "end of CoverTab[125521]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1040
			// _ = "end of CoverTab[125517]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1041
			_go_fuzz_dep_.CoverTab[125526]++
												switch value[i] {
			case ',', '?', '[', ']', '{', '}':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1043
				_go_fuzz_dep_.CoverTab[125527]++
													flow_indicators = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1044
				// _ = "end of CoverTab[125527]"
			case ':':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1045
				_go_fuzz_dep_.CoverTab[125528]++
													flow_indicators = true
													if followed_by_whitespace {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1047
					_go_fuzz_dep_.CoverTab[125531]++
														block_indicators = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1048
					// _ = "end of CoverTab[125531]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1049
					_go_fuzz_dep_.CoverTab[125532]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1049
					// _ = "end of CoverTab[125532]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1049
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1049
				// _ = "end of CoverTab[125528]"
			case '#':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1050
				_go_fuzz_dep_.CoverTab[125529]++
													if preceded_by_whitespace {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1051
					_go_fuzz_dep_.CoverTab[125533]++
														flow_indicators = true
														block_indicators = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1053
					// _ = "end of CoverTab[125533]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1054
					_go_fuzz_dep_.CoverTab[125534]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1054
					// _ = "end of CoverTab[125534]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1054
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1054
				// _ = "end of CoverTab[125529]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1054
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1054
				_go_fuzz_dep_.CoverTab[125530]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1054
				// _ = "end of CoverTab[125530]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1055
			// _ = "end of CoverTab[125526]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1056
		// _ = "end of CoverTab[125512]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1056
		_go_fuzz_dep_.CoverTab[125513]++

											if !is_printable(value, i) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
			_go_fuzz_dep_.CoverTab[125535]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
			return !is_ascii(value, i) && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
				_go_fuzz_dep_.CoverTab[125536]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
				return !emitter.unicode
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
				// _ = "end of CoverTab[125536]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
			}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
			// _ = "end of CoverTab[125535]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1058
			_go_fuzz_dep_.CoverTab[125537]++
												special_characters = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1059
			// _ = "end of CoverTab[125537]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1060
			_go_fuzz_dep_.CoverTab[125538]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1060
			// _ = "end of CoverTab[125538]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1060
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1060
		// _ = "end of CoverTab[125513]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1060
		_go_fuzz_dep_.CoverTab[125514]++
											if is_space(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1061
			_go_fuzz_dep_.CoverTab[125539]++
												if i == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1062
				_go_fuzz_dep_.CoverTab[125543]++
													leading_space = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1063
				// _ = "end of CoverTab[125543]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1064
				_go_fuzz_dep_.CoverTab[125544]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1064
				// _ = "end of CoverTab[125544]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1064
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1064
			// _ = "end of CoverTab[125539]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1064
			_go_fuzz_dep_.CoverTab[125540]++
												if i+width(value[i]) == len(value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1065
				_go_fuzz_dep_.CoverTab[125545]++
													trailing_space = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1066
				// _ = "end of CoverTab[125545]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1067
				_go_fuzz_dep_.CoverTab[125546]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1067
				// _ = "end of CoverTab[125546]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1067
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1067
			// _ = "end of CoverTab[125540]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1067
			_go_fuzz_dep_.CoverTab[125541]++
												if previous_break {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1068
				_go_fuzz_dep_.CoverTab[125547]++
													break_space = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1069
				// _ = "end of CoverTab[125547]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1070
				_go_fuzz_dep_.CoverTab[125548]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1070
				// _ = "end of CoverTab[125548]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1070
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1070
			// _ = "end of CoverTab[125541]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1070
			_go_fuzz_dep_.CoverTab[125542]++
												previous_space = true
												previous_break = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1072
			// _ = "end of CoverTab[125542]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1073
			_go_fuzz_dep_.CoverTab[125549]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1073
			if is_break(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1073
				_go_fuzz_dep_.CoverTab[125550]++
													line_breaks = true
													if i == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1075
					_go_fuzz_dep_.CoverTab[125554]++
														leading_break = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1076
					// _ = "end of CoverTab[125554]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1077
					_go_fuzz_dep_.CoverTab[125555]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1077
					// _ = "end of CoverTab[125555]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1077
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1077
				// _ = "end of CoverTab[125550]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1077
				_go_fuzz_dep_.CoverTab[125551]++
													if i+width(value[i]) == len(value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1078
					_go_fuzz_dep_.CoverTab[125556]++
														trailing_break = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1079
					// _ = "end of CoverTab[125556]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1080
					_go_fuzz_dep_.CoverTab[125557]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1080
					// _ = "end of CoverTab[125557]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1080
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1080
				// _ = "end of CoverTab[125551]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1080
				_go_fuzz_dep_.CoverTab[125552]++
													if previous_space {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1081
					_go_fuzz_dep_.CoverTab[125558]++
														space_break = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1082
					// _ = "end of CoverTab[125558]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1083
					_go_fuzz_dep_.CoverTab[125559]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1083
					// _ = "end of CoverTab[125559]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1083
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1083
				// _ = "end of CoverTab[125552]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1083
				_go_fuzz_dep_.CoverTab[125553]++
													previous_space = false
													previous_break = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1085
				// _ = "end of CoverTab[125553]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1086
				_go_fuzz_dep_.CoverTab[125560]++
													previous_space = false
													previous_break = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1088
				// _ = "end of CoverTab[125560]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1089
			// _ = "end of CoverTab[125549]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1089
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1089
		// _ = "end of CoverTab[125514]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1089
		_go_fuzz_dep_.CoverTab[125515]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1092
		preceded_by_whitespace = is_blankz(value, i)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1092
		// _ = "end of CoverTab[125515]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1093
	// _ = "end of CoverTab[125493]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1093
	_go_fuzz_dep_.CoverTab[125494]++

										emitter.scalar_data.multiline = line_breaks
										emitter.scalar_data.flow_plain_allowed = true
										emitter.scalar_data.block_plain_allowed = true
										emitter.scalar_data.single_quoted_allowed = true
										emitter.scalar_data.block_allowed = true

										if leading_space || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		_go_fuzz_dep_.CoverTab[125561]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		return leading_break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		// _ = "end of CoverTab[125561]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		_go_fuzz_dep_.CoverTab[125562]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		return trailing_space
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		// _ = "end of CoverTab[125562]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		_go_fuzz_dep_.CoverTab[125563]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		return trailing_break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		// _ = "end of CoverTab[125563]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1101
		_go_fuzz_dep_.CoverTab[125564]++
											emitter.scalar_data.flow_plain_allowed = false
											emitter.scalar_data.block_plain_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1103
		// _ = "end of CoverTab[125564]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1104
		_go_fuzz_dep_.CoverTab[125565]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1104
		// _ = "end of CoverTab[125565]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1104
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1104
	// _ = "end of CoverTab[125494]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1104
	_go_fuzz_dep_.CoverTab[125495]++
										if trailing_space {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1105
		_go_fuzz_dep_.CoverTab[125566]++
											emitter.scalar_data.block_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1106
		// _ = "end of CoverTab[125566]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1107
		_go_fuzz_dep_.CoverTab[125567]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1107
		// _ = "end of CoverTab[125567]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1107
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1107
	// _ = "end of CoverTab[125495]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1107
	_go_fuzz_dep_.CoverTab[125496]++
										if break_space {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1108
		_go_fuzz_dep_.CoverTab[125568]++
											emitter.scalar_data.flow_plain_allowed = false
											emitter.scalar_data.block_plain_allowed = false
											emitter.scalar_data.single_quoted_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1111
		// _ = "end of CoverTab[125568]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1112
		_go_fuzz_dep_.CoverTab[125569]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1112
		// _ = "end of CoverTab[125569]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1112
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1112
	// _ = "end of CoverTab[125496]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1112
	_go_fuzz_dep_.CoverTab[125497]++
										if space_break || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1113
		_go_fuzz_dep_.CoverTab[125570]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1113
		return special_characters
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1113
		// _ = "end of CoverTab[125570]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1113
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1113
		_go_fuzz_dep_.CoverTab[125571]++
											emitter.scalar_data.flow_plain_allowed = false
											emitter.scalar_data.block_plain_allowed = false
											emitter.scalar_data.single_quoted_allowed = false
											emitter.scalar_data.block_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1117
		// _ = "end of CoverTab[125571]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1118
		_go_fuzz_dep_.CoverTab[125572]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1118
		// _ = "end of CoverTab[125572]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1118
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1118
	// _ = "end of CoverTab[125497]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1118
	_go_fuzz_dep_.CoverTab[125498]++
										if line_breaks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1119
		_go_fuzz_dep_.CoverTab[125573]++
											emitter.scalar_data.flow_plain_allowed = false
											emitter.scalar_data.block_plain_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1121
		// _ = "end of CoverTab[125573]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1122
		_go_fuzz_dep_.CoverTab[125574]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1122
		// _ = "end of CoverTab[125574]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1122
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1122
	// _ = "end of CoverTab[125498]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1122
	_go_fuzz_dep_.CoverTab[125499]++
										if flow_indicators {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1123
		_go_fuzz_dep_.CoverTab[125575]++
											emitter.scalar_data.flow_plain_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1124
		// _ = "end of CoverTab[125575]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1125
		_go_fuzz_dep_.CoverTab[125576]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1125
		// _ = "end of CoverTab[125576]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1125
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1125
	// _ = "end of CoverTab[125499]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1125
	_go_fuzz_dep_.CoverTab[125500]++
										if block_indicators {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1126
		_go_fuzz_dep_.CoverTab[125577]++
											emitter.scalar_data.block_plain_allowed = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1127
		// _ = "end of CoverTab[125577]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1128
		_go_fuzz_dep_.CoverTab[125578]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1128
		// _ = "end of CoverTab[125578]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1128
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1128
	// _ = "end of CoverTab[125500]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1128
	_go_fuzz_dep_.CoverTab[125501]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1129
	// _ = "end of CoverTab[125501]"
}

// Check if the event data is valid.
func yaml_emitter_analyze_event(emitter *yaml_emitter_t, event *yaml_event_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1133
	_go_fuzz_dep_.CoverTab[125579]++

										emitter.anchor_data.anchor = nil
										emitter.tag_data.handle = nil
										emitter.tag_data.suffix = nil
										emitter.scalar_data.value = nil

										switch event.typ {
	case yaml_ALIAS_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1141
		_go_fuzz_dep_.CoverTab[125581]++
											if !yaml_emitter_analyze_anchor(emitter, event.anchor, true) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1142
			_go_fuzz_dep_.CoverTab[125590]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1143
			// _ = "end of CoverTab[125590]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1144
			_go_fuzz_dep_.CoverTab[125591]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1144
			// _ = "end of CoverTab[125591]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1144
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1144
		// _ = "end of CoverTab[125581]"

	case yaml_SCALAR_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1146
		_go_fuzz_dep_.CoverTab[125582]++
											if len(event.anchor) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1147
			_go_fuzz_dep_.CoverTab[125592]++
												if !yaml_emitter_analyze_anchor(emitter, event.anchor, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1148
				_go_fuzz_dep_.CoverTab[125593]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1149
				// _ = "end of CoverTab[125593]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1150
				_go_fuzz_dep_.CoverTab[125594]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1150
				// _ = "end of CoverTab[125594]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1150
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1150
			// _ = "end of CoverTab[125592]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1151
			_go_fuzz_dep_.CoverTab[125595]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1151
			// _ = "end of CoverTab[125595]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1151
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1151
		// _ = "end of CoverTab[125582]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1151
		_go_fuzz_dep_.CoverTab[125583]++
											if len(event.tag) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
			_go_fuzz_dep_.CoverTab[125596]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
			return (emitter.canonical || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
				_go_fuzz_dep_.CoverTab[125597]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
				return (!event.implicit && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
					_go_fuzz_dep_.CoverTab[125598]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
					return !event.quoted_implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
					// _ = "end of CoverTab[125598]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
				// _ = "end of CoverTab[125597]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
			// _ = "end of CoverTab[125596]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1152
			_go_fuzz_dep_.CoverTab[125599]++
												if !yaml_emitter_analyze_tag(emitter, event.tag) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1153
				_go_fuzz_dep_.CoverTab[125600]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1154
				// _ = "end of CoverTab[125600]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1155
				_go_fuzz_dep_.CoverTab[125601]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1155
				// _ = "end of CoverTab[125601]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1155
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1155
			// _ = "end of CoverTab[125599]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1156
			_go_fuzz_dep_.CoverTab[125602]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1156
			// _ = "end of CoverTab[125602]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1156
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1156
		// _ = "end of CoverTab[125583]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1156
		_go_fuzz_dep_.CoverTab[125584]++
											if !yaml_emitter_analyze_scalar(emitter, event.value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1157
			_go_fuzz_dep_.CoverTab[125603]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1158
			// _ = "end of CoverTab[125603]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1159
			_go_fuzz_dep_.CoverTab[125604]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1159
			// _ = "end of CoverTab[125604]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1159
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1159
		// _ = "end of CoverTab[125584]"

	case yaml_SEQUENCE_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1161
		_go_fuzz_dep_.CoverTab[125585]++
											if len(event.anchor) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1162
			_go_fuzz_dep_.CoverTab[125605]++
												if !yaml_emitter_analyze_anchor(emitter, event.anchor, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1163
				_go_fuzz_dep_.CoverTab[125606]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1164
				// _ = "end of CoverTab[125606]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1165
				_go_fuzz_dep_.CoverTab[125607]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1165
				// _ = "end of CoverTab[125607]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1165
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1165
			// _ = "end of CoverTab[125605]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1166
			_go_fuzz_dep_.CoverTab[125608]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1166
			// _ = "end of CoverTab[125608]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1166
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1166
		// _ = "end of CoverTab[125585]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1166
		_go_fuzz_dep_.CoverTab[125586]++
											if len(event.tag) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
			_go_fuzz_dep_.CoverTab[125609]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
			return (emitter.canonical || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
				_go_fuzz_dep_.CoverTab[125610]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
				return !event.implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
				// _ = "end of CoverTab[125610]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
			// _ = "end of CoverTab[125609]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1167
			_go_fuzz_dep_.CoverTab[125611]++
												if !yaml_emitter_analyze_tag(emitter, event.tag) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1168
				_go_fuzz_dep_.CoverTab[125612]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1169
				// _ = "end of CoverTab[125612]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1170
				_go_fuzz_dep_.CoverTab[125613]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1170
				// _ = "end of CoverTab[125613]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1170
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1170
			// _ = "end of CoverTab[125611]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1171
			_go_fuzz_dep_.CoverTab[125614]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1171
			// _ = "end of CoverTab[125614]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1171
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1171
		// _ = "end of CoverTab[125586]"

	case yaml_MAPPING_START_EVENT:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1173
		_go_fuzz_dep_.CoverTab[125587]++
											if len(event.anchor) > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1174
			_go_fuzz_dep_.CoverTab[125615]++
												if !yaml_emitter_analyze_anchor(emitter, event.anchor, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1175
				_go_fuzz_dep_.CoverTab[125616]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1176
				// _ = "end of CoverTab[125616]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1177
				_go_fuzz_dep_.CoverTab[125617]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1177
				// _ = "end of CoverTab[125617]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1177
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1177
			// _ = "end of CoverTab[125615]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1178
			_go_fuzz_dep_.CoverTab[125618]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1178
			// _ = "end of CoverTab[125618]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1178
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1178
		// _ = "end of CoverTab[125587]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1178
		_go_fuzz_dep_.CoverTab[125588]++
											if len(event.tag) > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
			_go_fuzz_dep_.CoverTab[125619]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
			return (emitter.canonical || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
				_go_fuzz_dep_.CoverTab[125620]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
				return !event.implicit
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
				// _ = "end of CoverTab[125620]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
			// _ = "end of CoverTab[125619]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1179
			_go_fuzz_dep_.CoverTab[125621]++
												if !yaml_emitter_analyze_tag(emitter, event.tag) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1180
				_go_fuzz_dep_.CoverTab[125622]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1181
				// _ = "end of CoverTab[125622]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1182
				_go_fuzz_dep_.CoverTab[125623]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1182
				// _ = "end of CoverTab[125623]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1182
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1182
			// _ = "end of CoverTab[125621]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1183
			_go_fuzz_dep_.CoverTab[125624]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1183
			// _ = "end of CoverTab[125624]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1183
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1183
		// _ = "end of CoverTab[125588]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1183
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1183
		_go_fuzz_dep_.CoverTab[125589]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1183
		// _ = "end of CoverTab[125589]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1184
	// _ = "end of CoverTab[125579]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1184
	_go_fuzz_dep_.CoverTab[125580]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1185
	// _ = "end of CoverTab[125580]"
}

// Write the BOM character.
func yaml_emitter_write_bom(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1189
	_go_fuzz_dep_.CoverTab[125625]++
										if !flush(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1190
		_go_fuzz_dep_.CoverTab[125627]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1191
		// _ = "end of CoverTab[125627]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1192
		_go_fuzz_dep_.CoverTab[125628]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1192
		// _ = "end of CoverTab[125628]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1192
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1192
	// _ = "end of CoverTab[125625]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1192
	_go_fuzz_dep_.CoverTab[125626]++
										pos := emitter.buffer_pos
										emitter.buffer[pos+0] = '\xEF'
										emitter.buffer[pos+1] = '\xBB'
										emitter.buffer[pos+2] = '\xBF'
										emitter.buffer_pos += 3
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1198
	// _ = "end of CoverTab[125626]"
}

func yaml_emitter_write_indent(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1201
	_go_fuzz_dep_.CoverTab[125629]++
										indent := emitter.indent
										if indent < 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1203
		_go_fuzz_dep_.CoverTab[125633]++
											indent = 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1204
		// _ = "end of CoverTab[125633]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1205
		_go_fuzz_dep_.CoverTab[125634]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1205
		// _ = "end of CoverTab[125634]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1205
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1205
	// _ = "end of CoverTab[125629]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1205
	_go_fuzz_dep_.CoverTab[125630]++
										if !emitter.indention || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		_go_fuzz_dep_.CoverTab[125635]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		return emitter.column > indent
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		// _ = "end of CoverTab[125635]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		_go_fuzz_dep_.CoverTab[125636]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		return (emitter.column == indent && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
			_go_fuzz_dep_.CoverTab[125637]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
			return !emitter.whitespace
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
			// _ = "end of CoverTab[125637]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		// _ = "end of CoverTab[125636]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1206
		_go_fuzz_dep_.CoverTab[125638]++
											if !put_break(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1207
			_go_fuzz_dep_.CoverTab[125639]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1208
			// _ = "end of CoverTab[125639]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1209
			_go_fuzz_dep_.CoverTab[125640]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1209
			// _ = "end of CoverTab[125640]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1209
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1209
		// _ = "end of CoverTab[125638]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1210
		_go_fuzz_dep_.CoverTab[125641]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1210
		// _ = "end of CoverTab[125641]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1210
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1210
	// _ = "end of CoverTab[125630]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1210
	_go_fuzz_dep_.CoverTab[125631]++
										for emitter.column < indent {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1211
		_go_fuzz_dep_.CoverTab[125642]++
											if !put(emitter, ' ') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1212
			_go_fuzz_dep_.CoverTab[125643]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1213
			// _ = "end of CoverTab[125643]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1214
			_go_fuzz_dep_.CoverTab[125644]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1214
			// _ = "end of CoverTab[125644]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1214
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1214
		// _ = "end of CoverTab[125642]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1215
	// _ = "end of CoverTab[125631]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1215
	_go_fuzz_dep_.CoverTab[125632]++
										emitter.whitespace = true
										emitter.indention = true
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1218
	// _ = "end of CoverTab[125632]"
}

func yaml_emitter_write_indicator(emitter *yaml_emitter_t, indicator []byte, need_whitespace, is_whitespace, is_indention bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1221
	_go_fuzz_dep_.CoverTab[125645]++
										if need_whitespace && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1222
		_go_fuzz_dep_.CoverTab[125648]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1222
		return !emitter.whitespace
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1222
		// _ = "end of CoverTab[125648]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1222
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1222
		_go_fuzz_dep_.CoverTab[125649]++
											if !put(emitter, ' ') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1223
			_go_fuzz_dep_.CoverTab[125650]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1224
			// _ = "end of CoverTab[125650]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1225
			_go_fuzz_dep_.CoverTab[125651]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1225
			// _ = "end of CoverTab[125651]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1225
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1225
		// _ = "end of CoverTab[125649]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1226
		_go_fuzz_dep_.CoverTab[125652]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1226
		// _ = "end of CoverTab[125652]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1226
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1226
	// _ = "end of CoverTab[125645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1226
	_go_fuzz_dep_.CoverTab[125646]++
										if !write_all(emitter, indicator) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1227
		_go_fuzz_dep_.CoverTab[125653]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1228
		// _ = "end of CoverTab[125653]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1229
		_go_fuzz_dep_.CoverTab[125654]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1229
		// _ = "end of CoverTab[125654]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1229
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1229
	// _ = "end of CoverTab[125646]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1229
	_go_fuzz_dep_.CoverTab[125647]++
										emitter.whitespace = is_whitespace
										emitter.indention = (emitter.indention && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1231
		_go_fuzz_dep_.CoverTab[125655]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1231
		return is_indention
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1231
		// _ = "end of CoverTab[125655]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1231
	}())
										emitter.open_ended = false
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1233
	// _ = "end of CoverTab[125647]"
}

func yaml_emitter_write_anchor(emitter *yaml_emitter_t, value []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1236
	_go_fuzz_dep_.CoverTab[125656]++
										if !write_all(emitter, value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1237
		_go_fuzz_dep_.CoverTab[125658]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1238
		// _ = "end of CoverTab[125658]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1239
		_go_fuzz_dep_.CoverTab[125659]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1239
		// _ = "end of CoverTab[125659]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1239
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1239
	// _ = "end of CoverTab[125656]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1239
	_go_fuzz_dep_.CoverTab[125657]++
										emitter.whitespace = false
										emitter.indention = false
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1242
	// _ = "end of CoverTab[125657]"
}

func yaml_emitter_write_tag_handle(emitter *yaml_emitter_t, value []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1245
	_go_fuzz_dep_.CoverTab[125660]++
										if !emitter.whitespace {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1246
		_go_fuzz_dep_.CoverTab[125663]++
											if !put(emitter, ' ') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1247
			_go_fuzz_dep_.CoverTab[125664]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1248
			// _ = "end of CoverTab[125664]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1249
			_go_fuzz_dep_.CoverTab[125665]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1249
			// _ = "end of CoverTab[125665]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1249
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1249
		// _ = "end of CoverTab[125663]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1250
		_go_fuzz_dep_.CoverTab[125666]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1250
		// _ = "end of CoverTab[125666]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1250
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1250
	// _ = "end of CoverTab[125660]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1250
	_go_fuzz_dep_.CoverTab[125661]++
										if !write_all(emitter, value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1251
		_go_fuzz_dep_.CoverTab[125667]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1252
		// _ = "end of CoverTab[125667]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1253
		_go_fuzz_dep_.CoverTab[125668]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1253
		// _ = "end of CoverTab[125668]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1253
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1253
	// _ = "end of CoverTab[125661]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1253
	_go_fuzz_dep_.CoverTab[125662]++
										emitter.whitespace = false
										emitter.indention = false
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1256
	// _ = "end of CoverTab[125662]"
}

func yaml_emitter_write_tag_content(emitter *yaml_emitter_t, value []byte, need_whitespace bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1259
	_go_fuzz_dep_.CoverTab[125669]++
										if need_whitespace && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1260
		_go_fuzz_dep_.CoverTab[125672]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1260
		return !emitter.whitespace
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1260
		// _ = "end of CoverTab[125672]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1260
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1260
		_go_fuzz_dep_.CoverTab[125673]++
											if !put(emitter, ' ') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1261
			_go_fuzz_dep_.CoverTab[125674]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1262
			// _ = "end of CoverTab[125674]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1263
			_go_fuzz_dep_.CoverTab[125675]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1263
			// _ = "end of CoverTab[125675]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1263
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1263
		// _ = "end of CoverTab[125673]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1264
		_go_fuzz_dep_.CoverTab[125676]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1264
		// _ = "end of CoverTab[125676]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1264
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1264
	// _ = "end of CoverTab[125669]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1264
	_go_fuzz_dep_.CoverTab[125670]++
										for i := 0; i < len(value); {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1265
		_go_fuzz_dep_.CoverTab[125677]++
											var must_write bool
											switch value[i] {
		case ';', '/', '?', ':', '@', '&', '=', '+', '$', ',', '_', '.', '~', '*', '\'', '(', ')', '[', ']':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1268
			_go_fuzz_dep_.CoverTab[125679]++
												must_write = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1269
			// _ = "end of CoverTab[125679]"
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1270
			_go_fuzz_dep_.CoverTab[125680]++
												must_write = is_alpha(value, i)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1271
			// _ = "end of CoverTab[125680]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1272
		// _ = "end of CoverTab[125677]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1272
		_go_fuzz_dep_.CoverTab[125678]++
											if must_write {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1273
			_go_fuzz_dep_.CoverTab[125681]++
												if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1274
				_go_fuzz_dep_.CoverTab[125682]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1275
				// _ = "end of CoverTab[125682]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1276
				_go_fuzz_dep_.CoverTab[125683]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1276
				// _ = "end of CoverTab[125683]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1276
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1276
			// _ = "end of CoverTab[125681]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1277
			_go_fuzz_dep_.CoverTab[125684]++
												w := width(value[i])
												for k := 0; k < w; k++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1279
				_go_fuzz_dep_.CoverTab[125685]++
													octet := value[i]
													i++
													if !put(emitter, '%') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1282
					_go_fuzz_dep_.CoverTab[125690]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1283
					// _ = "end of CoverTab[125690]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1284
					_go_fuzz_dep_.CoverTab[125691]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1284
					// _ = "end of CoverTab[125691]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1284
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1284
				// _ = "end of CoverTab[125685]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1284
				_go_fuzz_dep_.CoverTab[125686]++

													c := octet >> 4
													if c < 10 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1287
					_go_fuzz_dep_.CoverTab[125692]++
														c += '0'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1288
					// _ = "end of CoverTab[125692]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1289
					_go_fuzz_dep_.CoverTab[125693]++
														c += 'A' - 10
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1290
					// _ = "end of CoverTab[125693]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1291
				// _ = "end of CoverTab[125686]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1291
				_go_fuzz_dep_.CoverTab[125687]++
													if !put(emitter, c) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1292
					_go_fuzz_dep_.CoverTab[125694]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1293
					// _ = "end of CoverTab[125694]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1294
					_go_fuzz_dep_.CoverTab[125695]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1294
					// _ = "end of CoverTab[125695]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1294
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1294
				// _ = "end of CoverTab[125687]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1294
				_go_fuzz_dep_.CoverTab[125688]++

													c = octet & 0x0f
													if c < 10 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1297
					_go_fuzz_dep_.CoverTab[125696]++
														c += '0'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1298
					// _ = "end of CoverTab[125696]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1299
					_go_fuzz_dep_.CoverTab[125697]++
														c += 'A' - 10
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1300
					// _ = "end of CoverTab[125697]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1301
				// _ = "end of CoverTab[125688]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1301
				_go_fuzz_dep_.CoverTab[125689]++
													if !put(emitter, c) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1302
					_go_fuzz_dep_.CoverTab[125698]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1303
					// _ = "end of CoverTab[125698]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1304
					_go_fuzz_dep_.CoverTab[125699]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1304
					// _ = "end of CoverTab[125699]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1304
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1304
				// _ = "end of CoverTab[125689]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1305
			// _ = "end of CoverTab[125684]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1306
		// _ = "end of CoverTab[125678]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1307
	// _ = "end of CoverTab[125670]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1307
	_go_fuzz_dep_.CoverTab[125671]++
										emitter.whitespace = false
										emitter.indention = false
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1310
	// _ = "end of CoverTab[125671]"
}

func yaml_emitter_write_plain_scalar(emitter *yaml_emitter_t, value []byte, allow_breaks bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1313
	_go_fuzz_dep_.CoverTab[125700]++
										if !emitter.whitespace {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1314
		_go_fuzz_dep_.CoverTab[125704]++
											if !put(emitter, ' ') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1315
			_go_fuzz_dep_.CoverTab[125705]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1316
			// _ = "end of CoverTab[125705]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1317
			_go_fuzz_dep_.CoverTab[125706]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1317
			// _ = "end of CoverTab[125706]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1317
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1317
		// _ = "end of CoverTab[125704]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1318
		_go_fuzz_dep_.CoverTab[125707]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1318
		// _ = "end of CoverTab[125707]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1318
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1318
	// _ = "end of CoverTab[125700]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1318
	_go_fuzz_dep_.CoverTab[125701]++

										spaces := false
										breaks := false
										for i := 0; i < len(value); {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1322
		_go_fuzz_dep_.CoverTab[125708]++
											if is_space(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1323
			_go_fuzz_dep_.CoverTab[125709]++
												if allow_breaks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				_go_fuzz_dep_.CoverTab[125711]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				return !spaces
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				// _ = "end of CoverTab[125711]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				_go_fuzz_dep_.CoverTab[125712]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				return emitter.column > emitter.best_width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				// _ = "end of CoverTab[125712]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				_go_fuzz_dep_.CoverTab[125713]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				return !is_space(value, i+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				// _ = "end of CoverTab[125713]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1324
				_go_fuzz_dep_.CoverTab[125714]++
													if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1325
					_go_fuzz_dep_.CoverTab[125716]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1326
					// _ = "end of CoverTab[125716]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1327
					_go_fuzz_dep_.CoverTab[125717]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1327
					// _ = "end of CoverTab[125717]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1327
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1327
				// _ = "end of CoverTab[125714]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1327
				_go_fuzz_dep_.CoverTab[125715]++
													i += width(value[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1328
				// _ = "end of CoverTab[125715]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1329
				_go_fuzz_dep_.CoverTab[125718]++
													if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1330
					_go_fuzz_dep_.CoverTab[125719]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1331
					// _ = "end of CoverTab[125719]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1332
					_go_fuzz_dep_.CoverTab[125720]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1332
					// _ = "end of CoverTab[125720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1332
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1332
				// _ = "end of CoverTab[125718]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1333
			// _ = "end of CoverTab[125709]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1333
			_go_fuzz_dep_.CoverTab[125710]++
												spaces = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1334
			// _ = "end of CoverTab[125710]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1335
			_go_fuzz_dep_.CoverTab[125721]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1335
			if is_break(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1335
				_go_fuzz_dep_.CoverTab[125722]++
													if !breaks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1336
					_go_fuzz_dep_.CoverTab[125725]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1336
					return value[i] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1336
					// _ = "end of CoverTab[125725]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1336
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1336
					_go_fuzz_dep_.CoverTab[125726]++
														if !put_break(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1337
						_go_fuzz_dep_.CoverTab[125727]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1338
						// _ = "end of CoverTab[125727]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1339
						_go_fuzz_dep_.CoverTab[125728]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1339
						// _ = "end of CoverTab[125728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1339
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1339
					// _ = "end of CoverTab[125726]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1340
					_go_fuzz_dep_.CoverTab[125729]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1340
					// _ = "end of CoverTab[125729]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1340
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1340
				// _ = "end of CoverTab[125722]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1340
				_go_fuzz_dep_.CoverTab[125723]++
													if !write_break(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1341
					_go_fuzz_dep_.CoverTab[125730]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1342
					// _ = "end of CoverTab[125730]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1343
					_go_fuzz_dep_.CoverTab[125731]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1343
					// _ = "end of CoverTab[125731]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1343
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1343
				// _ = "end of CoverTab[125723]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1343
				_go_fuzz_dep_.CoverTab[125724]++
													emitter.indention = true
													breaks = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1345
				// _ = "end of CoverTab[125724]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1346
				_go_fuzz_dep_.CoverTab[125732]++
													if breaks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1347
					_go_fuzz_dep_.CoverTab[125735]++
														if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1348
						_go_fuzz_dep_.CoverTab[125736]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1349
						// _ = "end of CoverTab[125736]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1350
						_go_fuzz_dep_.CoverTab[125737]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1350
						// _ = "end of CoverTab[125737]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1350
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1350
					// _ = "end of CoverTab[125735]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1351
					_go_fuzz_dep_.CoverTab[125738]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1351
					// _ = "end of CoverTab[125738]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1351
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1351
				// _ = "end of CoverTab[125732]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1351
				_go_fuzz_dep_.CoverTab[125733]++
													if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1352
					_go_fuzz_dep_.CoverTab[125739]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1353
					// _ = "end of CoverTab[125739]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1354
					_go_fuzz_dep_.CoverTab[125740]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1354
					// _ = "end of CoverTab[125740]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1354
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1354
				// _ = "end of CoverTab[125733]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1354
				_go_fuzz_dep_.CoverTab[125734]++
													emitter.indention = false
													spaces = false
													breaks = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1357
				// _ = "end of CoverTab[125734]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1358
			// _ = "end of CoverTab[125721]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1358
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1358
		// _ = "end of CoverTab[125708]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1359
	// _ = "end of CoverTab[125701]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1359
	_go_fuzz_dep_.CoverTab[125702]++

										emitter.whitespace = false
										emitter.indention = false
										if emitter.root_context {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1363
		_go_fuzz_dep_.CoverTab[125741]++
											emitter.open_ended = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1364
		// _ = "end of CoverTab[125741]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1365
		_go_fuzz_dep_.CoverTab[125742]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1365
		// _ = "end of CoverTab[125742]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1365
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1365
	// _ = "end of CoverTab[125702]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1365
	_go_fuzz_dep_.CoverTab[125703]++

										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1367
	// _ = "end of CoverTab[125703]"
}

func yaml_emitter_write_single_quoted_scalar(emitter *yaml_emitter_t, value []byte, allow_breaks bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1370
	_go_fuzz_dep_.CoverTab[125743]++

										if !yaml_emitter_write_indicator(emitter, []byte{'\''}, true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1372
		_go_fuzz_dep_.CoverTab[125747]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1373
		// _ = "end of CoverTab[125747]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1374
		_go_fuzz_dep_.CoverTab[125748]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1374
		// _ = "end of CoverTab[125748]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1374
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1374
	// _ = "end of CoverTab[125743]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1374
	_go_fuzz_dep_.CoverTab[125744]++

										spaces := false
										breaks := false
										for i := 0; i < len(value); {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1378
		_go_fuzz_dep_.CoverTab[125749]++
											if is_space(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1379
			_go_fuzz_dep_.CoverTab[125750]++
												if allow_breaks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				_go_fuzz_dep_.CoverTab[125752]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				return !spaces
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				// _ = "end of CoverTab[125752]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				_go_fuzz_dep_.CoverTab[125753]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				return emitter.column > emitter.best_width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				// _ = "end of CoverTab[125753]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				_go_fuzz_dep_.CoverTab[125754]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				return i > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				// _ = "end of CoverTab[125754]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				_go_fuzz_dep_.CoverTab[125755]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				return i < len(value)-1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				// _ = "end of CoverTab[125755]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				_go_fuzz_dep_.CoverTab[125756]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				return !is_space(value, i+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				// _ = "end of CoverTab[125756]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1380
				_go_fuzz_dep_.CoverTab[125757]++
													if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1381
					_go_fuzz_dep_.CoverTab[125759]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1382
					// _ = "end of CoverTab[125759]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1383
					_go_fuzz_dep_.CoverTab[125760]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1383
					// _ = "end of CoverTab[125760]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1383
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1383
				// _ = "end of CoverTab[125757]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1383
				_go_fuzz_dep_.CoverTab[125758]++
													i += width(value[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1384
				// _ = "end of CoverTab[125758]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1385
				_go_fuzz_dep_.CoverTab[125761]++
													if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1386
					_go_fuzz_dep_.CoverTab[125762]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1387
					// _ = "end of CoverTab[125762]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1388
					_go_fuzz_dep_.CoverTab[125763]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1388
					// _ = "end of CoverTab[125763]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1388
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1388
				// _ = "end of CoverTab[125761]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1389
			// _ = "end of CoverTab[125750]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1389
			_go_fuzz_dep_.CoverTab[125751]++
												spaces = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1390
			// _ = "end of CoverTab[125751]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1391
			_go_fuzz_dep_.CoverTab[125764]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1391
			if is_break(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1391
				_go_fuzz_dep_.CoverTab[125765]++
													if !breaks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1392
					_go_fuzz_dep_.CoverTab[125768]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1392
					return value[i] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1392
					// _ = "end of CoverTab[125768]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1392
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1392
					_go_fuzz_dep_.CoverTab[125769]++
														if !put_break(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1393
						_go_fuzz_dep_.CoverTab[125770]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1394
						// _ = "end of CoverTab[125770]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1395
						_go_fuzz_dep_.CoverTab[125771]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1395
						// _ = "end of CoverTab[125771]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1395
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1395
					// _ = "end of CoverTab[125769]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1396
					_go_fuzz_dep_.CoverTab[125772]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1396
					// _ = "end of CoverTab[125772]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1396
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1396
				// _ = "end of CoverTab[125765]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1396
				_go_fuzz_dep_.CoverTab[125766]++
													if !write_break(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1397
					_go_fuzz_dep_.CoverTab[125773]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1398
					// _ = "end of CoverTab[125773]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1399
					_go_fuzz_dep_.CoverTab[125774]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1399
					// _ = "end of CoverTab[125774]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1399
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1399
				// _ = "end of CoverTab[125766]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1399
				_go_fuzz_dep_.CoverTab[125767]++
													emitter.indention = true
													breaks = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1401
				// _ = "end of CoverTab[125767]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1402
				_go_fuzz_dep_.CoverTab[125775]++
													if breaks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1403
					_go_fuzz_dep_.CoverTab[125779]++
														if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1404
						_go_fuzz_dep_.CoverTab[125780]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1405
						// _ = "end of CoverTab[125780]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1406
						_go_fuzz_dep_.CoverTab[125781]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1406
						// _ = "end of CoverTab[125781]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1406
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1406
					// _ = "end of CoverTab[125779]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1407
					_go_fuzz_dep_.CoverTab[125782]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1407
					// _ = "end of CoverTab[125782]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1407
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1407
				// _ = "end of CoverTab[125775]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1407
				_go_fuzz_dep_.CoverTab[125776]++
													if value[i] == '\'' {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1408
					_go_fuzz_dep_.CoverTab[125783]++
														if !put(emitter, '\'') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1409
						_go_fuzz_dep_.CoverTab[125784]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1410
						// _ = "end of CoverTab[125784]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1411
						_go_fuzz_dep_.CoverTab[125785]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1411
						// _ = "end of CoverTab[125785]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1411
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1411
					// _ = "end of CoverTab[125783]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1412
					_go_fuzz_dep_.CoverTab[125786]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1412
					// _ = "end of CoverTab[125786]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1412
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1412
				// _ = "end of CoverTab[125776]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1412
				_go_fuzz_dep_.CoverTab[125777]++
													if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1413
					_go_fuzz_dep_.CoverTab[125787]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1414
					// _ = "end of CoverTab[125787]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1415
					_go_fuzz_dep_.CoverTab[125788]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1415
					// _ = "end of CoverTab[125788]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1415
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1415
				// _ = "end of CoverTab[125777]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1415
				_go_fuzz_dep_.CoverTab[125778]++
													emitter.indention = false
													spaces = false
													breaks = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1418
				// _ = "end of CoverTab[125778]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1419
			// _ = "end of CoverTab[125764]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1419
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1419
		// _ = "end of CoverTab[125749]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1420
	// _ = "end of CoverTab[125744]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1420
	_go_fuzz_dep_.CoverTab[125745]++
										if !yaml_emitter_write_indicator(emitter, []byte{'\''}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1421
		_go_fuzz_dep_.CoverTab[125789]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1422
		// _ = "end of CoverTab[125789]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1423
		_go_fuzz_dep_.CoverTab[125790]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1423
		// _ = "end of CoverTab[125790]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1423
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1423
	// _ = "end of CoverTab[125745]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1423
	_go_fuzz_dep_.CoverTab[125746]++
										emitter.whitespace = false
										emitter.indention = false
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1426
	// _ = "end of CoverTab[125746]"
}

func yaml_emitter_write_double_quoted_scalar(emitter *yaml_emitter_t, value []byte, allow_breaks bool) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1429
	_go_fuzz_dep_.CoverTab[125791]++
										spaces := false
										if !yaml_emitter_write_indicator(emitter, []byte{'"'}, true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1431
		_go_fuzz_dep_.CoverTab[125795]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1432
		// _ = "end of CoverTab[125795]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1433
		_go_fuzz_dep_.CoverTab[125796]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1433
		// _ = "end of CoverTab[125796]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1433
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1433
	// _ = "end of CoverTab[125791]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1433
	_go_fuzz_dep_.CoverTab[125792]++

										for i := 0; i < len(value); {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1435
		_go_fuzz_dep_.CoverTab[125797]++
											if !is_printable(value, i) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
			_go_fuzz_dep_.CoverTab[125798]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
			return (!emitter.unicode && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
				_go_fuzz_dep_.CoverTab[125799]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
				return !is_ascii(value, i)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
				// _ = "end of CoverTab[125799]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
			// _ = "end of CoverTab[125798]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
			_go_fuzz_dep_.CoverTab[125800]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1436
			return is_bom(value, i)
												// _ = "end of CoverTab[125800]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1437
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1437
			_go_fuzz_dep_.CoverTab[125801]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1437
			return is_break(value, i)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1437
			// _ = "end of CoverTab[125801]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1437
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1437
			_go_fuzz_dep_.CoverTab[125802]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1437
			return value[i] == '"'
												// _ = "end of CoverTab[125802]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1438
		}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1438
			_go_fuzz_dep_.CoverTab[125803]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1438
			return value[i] == '\\'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1438
			// _ = "end of CoverTab[125803]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1438
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1438
			_go_fuzz_dep_.CoverTab[125804]++

												octet := value[i]

												var w int
												var v rune
												switch {
			case octet&0x80 == 0x00:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1445
				_go_fuzz_dep_.CoverTab[125810]++
													w, v = 1, rune(octet&0x7F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1446
				// _ = "end of CoverTab[125810]"
			case octet&0xE0 == 0xC0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1447
				_go_fuzz_dep_.CoverTab[125811]++
													w, v = 2, rune(octet&0x1F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1448
				// _ = "end of CoverTab[125811]"
			case octet&0xF0 == 0xE0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1449
				_go_fuzz_dep_.CoverTab[125812]++
													w, v = 3, rune(octet&0x0F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1450
				// _ = "end of CoverTab[125812]"
			case octet&0xF8 == 0xF0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1451
				_go_fuzz_dep_.CoverTab[125813]++
													w, v = 4, rune(octet&0x07)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1452
				// _ = "end of CoverTab[125813]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1452
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1452
				_go_fuzz_dep_.CoverTab[125814]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1452
				// _ = "end of CoverTab[125814]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1453
			// _ = "end of CoverTab[125804]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1453
			_go_fuzz_dep_.CoverTab[125805]++
												for k := 1; k < w; k++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1454
				_go_fuzz_dep_.CoverTab[125815]++
													octet = value[i+k]
													v = (v << 6) + (rune(octet) & 0x3F)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1456
				// _ = "end of CoverTab[125815]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1457
			// _ = "end of CoverTab[125805]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1457
			_go_fuzz_dep_.CoverTab[125806]++
												i += w

												if !put(emitter, '\\') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1460
				_go_fuzz_dep_.CoverTab[125816]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1461
				// _ = "end of CoverTab[125816]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1462
				_go_fuzz_dep_.CoverTab[125817]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1462
				// _ = "end of CoverTab[125817]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1462
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1462
			// _ = "end of CoverTab[125806]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1462
			_go_fuzz_dep_.CoverTab[125807]++

												var ok bool
												switch v {
			case 0x00:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1466
				_go_fuzz_dep_.CoverTab[125818]++
													ok = put(emitter, '0')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1467
				// _ = "end of CoverTab[125818]"
			case 0x07:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1468
				_go_fuzz_dep_.CoverTab[125819]++
													ok = put(emitter, 'a')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1469
				// _ = "end of CoverTab[125819]"
			case 0x08:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1470
				_go_fuzz_dep_.CoverTab[125820]++
													ok = put(emitter, 'b')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1471
				// _ = "end of CoverTab[125820]"
			case 0x09:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1472
				_go_fuzz_dep_.CoverTab[125821]++
													ok = put(emitter, 't')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1473
				// _ = "end of CoverTab[125821]"
			case 0x0A:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1474
				_go_fuzz_dep_.CoverTab[125822]++
													ok = put(emitter, 'n')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1475
				// _ = "end of CoverTab[125822]"
			case 0x0b:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1476
				_go_fuzz_dep_.CoverTab[125823]++
													ok = put(emitter, 'v')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1477
				// _ = "end of CoverTab[125823]"
			case 0x0c:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1478
				_go_fuzz_dep_.CoverTab[125824]++
													ok = put(emitter, 'f')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1479
				// _ = "end of CoverTab[125824]"
			case 0x0d:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1480
				_go_fuzz_dep_.CoverTab[125825]++
													ok = put(emitter, 'r')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1481
				// _ = "end of CoverTab[125825]"
			case 0x1b:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1482
				_go_fuzz_dep_.CoverTab[125826]++
													ok = put(emitter, 'e')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1483
				// _ = "end of CoverTab[125826]"
			case 0x22:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1484
				_go_fuzz_dep_.CoverTab[125827]++
													ok = put(emitter, '"')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1485
				// _ = "end of CoverTab[125827]"
			case 0x5c:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1486
				_go_fuzz_dep_.CoverTab[125828]++
													ok = put(emitter, '\\')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1487
				// _ = "end of CoverTab[125828]"
			case 0x85:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1488
				_go_fuzz_dep_.CoverTab[125829]++
													ok = put(emitter, 'N')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1489
				// _ = "end of CoverTab[125829]"
			case 0xA0:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1490
				_go_fuzz_dep_.CoverTab[125830]++
													ok = put(emitter, '_')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1491
				// _ = "end of CoverTab[125830]"
			case 0x2028:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1492
				_go_fuzz_dep_.CoverTab[125831]++
													ok = put(emitter, 'L')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1493
				// _ = "end of CoverTab[125831]"
			case 0x2029:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1494
				_go_fuzz_dep_.CoverTab[125832]++
													ok = put(emitter, 'P')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1495
				// _ = "end of CoverTab[125832]"
			default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1496
				_go_fuzz_dep_.CoverTab[125833]++
													if v <= 0xFF {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1497
					_go_fuzz_dep_.CoverTab[125835]++
														ok = put(emitter, 'x')
														w = 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1499
					// _ = "end of CoverTab[125835]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1500
					_go_fuzz_dep_.CoverTab[125836]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1500
					if v <= 0xFFFF {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1500
						_go_fuzz_dep_.CoverTab[125837]++
															ok = put(emitter, 'u')
															w = 4
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1502
						// _ = "end of CoverTab[125837]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1503
						_go_fuzz_dep_.CoverTab[125838]++
															ok = put(emitter, 'U')
															w = 8
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1505
						// _ = "end of CoverTab[125838]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1506
					// _ = "end of CoverTab[125836]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1506
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1506
				// _ = "end of CoverTab[125833]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1506
				_go_fuzz_dep_.CoverTab[125834]++
													for k := (w - 1) * 4; ok && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1507
					_go_fuzz_dep_.CoverTab[125839]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1507
					return k >= 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1507
					// _ = "end of CoverTab[125839]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1507
				}(); k -= 4 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1507
					_go_fuzz_dep_.CoverTab[125840]++
														digit := byte((v >> uint(k)) & 0x0F)
														if digit < 10 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1509
						_go_fuzz_dep_.CoverTab[125841]++
															ok = put(emitter, digit+'0')
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1510
						// _ = "end of CoverTab[125841]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1511
						_go_fuzz_dep_.CoverTab[125842]++
															ok = put(emitter, digit+'A'-10)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1512
						// _ = "end of CoverTab[125842]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1513
					// _ = "end of CoverTab[125840]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1514
				// _ = "end of CoverTab[125834]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1515
			// _ = "end of CoverTab[125807]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1515
			_go_fuzz_dep_.CoverTab[125808]++
												if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1516
				_go_fuzz_dep_.CoverTab[125843]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1517
				// _ = "end of CoverTab[125843]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1518
				_go_fuzz_dep_.CoverTab[125844]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1518
				// _ = "end of CoverTab[125844]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1518
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1518
			// _ = "end of CoverTab[125808]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1518
			_go_fuzz_dep_.CoverTab[125809]++
												spaces = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1519
			// _ = "end of CoverTab[125809]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1520
			_go_fuzz_dep_.CoverTab[125845]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1520
			if is_space(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1520
				_go_fuzz_dep_.CoverTab[125846]++
													if allow_breaks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					_go_fuzz_dep_.CoverTab[125848]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					return !spaces
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					// _ = "end of CoverTab[125848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					_go_fuzz_dep_.CoverTab[125849]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					return emitter.column > emitter.best_width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					// _ = "end of CoverTab[125849]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					_go_fuzz_dep_.CoverTab[125850]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					return i > 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					// _ = "end of CoverTab[125850]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
				}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					_go_fuzz_dep_.CoverTab[125851]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					return i < len(value)-1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					// _ = "end of CoverTab[125851]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
				}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1521
					_go_fuzz_dep_.CoverTab[125852]++
														if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1522
						_go_fuzz_dep_.CoverTab[125855]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1523
						// _ = "end of CoverTab[125855]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1524
						_go_fuzz_dep_.CoverTab[125856]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1524
						// _ = "end of CoverTab[125856]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1524
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1524
					// _ = "end of CoverTab[125852]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1524
					_go_fuzz_dep_.CoverTab[125853]++
														if is_space(value, i+1) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1525
						_go_fuzz_dep_.CoverTab[125857]++
															if !put(emitter, '\\') {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1526
							_go_fuzz_dep_.CoverTab[125858]++
																return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1527
							// _ = "end of CoverTab[125858]"
						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1528
							_go_fuzz_dep_.CoverTab[125859]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1528
							// _ = "end of CoverTab[125859]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1528
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1528
						// _ = "end of CoverTab[125857]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1529
						_go_fuzz_dep_.CoverTab[125860]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1529
						// _ = "end of CoverTab[125860]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1529
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1529
					// _ = "end of CoverTab[125853]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1529
					_go_fuzz_dep_.CoverTab[125854]++
														i += width(value[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1530
					// _ = "end of CoverTab[125854]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1531
					_go_fuzz_dep_.CoverTab[125861]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1531
					if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1531
						_go_fuzz_dep_.CoverTab[125862]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1532
						// _ = "end of CoverTab[125862]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1533
						_go_fuzz_dep_.CoverTab[125863]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1533
						// _ = "end of CoverTab[125863]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1533
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1533
					// _ = "end of CoverTab[125861]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1533
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1533
				// _ = "end of CoverTab[125846]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1533
				_go_fuzz_dep_.CoverTab[125847]++
													spaces = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1534
				// _ = "end of CoverTab[125847]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1535
				_go_fuzz_dep_.CoverTab[125864]++
													if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1536
					_go_fuzz_dep_.CoverTab[125866]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1537
					// _ = "end of CoverTab[125866]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1538
					_go_fuzz_dep_.CoverTab[125867]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1538
					// _ = "end of CoverTab[125867]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1538
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1538
				// _ = "end of CoverTab[125864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1538
				_go_fuzz_dep_.CoverTab[125865]++
													spaces = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1539
				// _ = "end of CoverTab[125865]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1540
			// _ = "end of CoverTab[125845]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1540
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1540
		// _ = "end of CoverTab[125797]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1541
	// _ = "end of CoverTab[125792]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1541
	_go_fuzz_dep_.CoverTab[125793]++
										if !yaml_emitter_write_indicator(emitter, []byte{'"'}, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1542
		_go_fuzz_dep_.CoverTab[125868]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1543
		// _ = "end of CoverTab[125868]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1544
		_go_fuzz_dep_.CoverTab[125869]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1544
		// _ = "end of CoverTab[125869]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1544
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1544
	// _ = "end of CoverTab[125793]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1544
	_go_fuzz_dep_.CoverTab[125794]++
										emitter.whitespace = false
										emitter.indention = false
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1547
	// _ = "end of CoverTab[125794]"
}

func yaml_emitter_write_block_scalar_hints(emitter *yaml_emitter_t, value []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1550
	_go_fuzz_dep_.CoverTab[125870]++
										if is_space(value, 0) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1551
		_go_fuzz_dep_.CoverTab[125874]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1551
		return is_break(value, 0)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1551
		// _ = "end of CoverTab[125874]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1551
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1551
		_go_fuzz_dep_.CoverTab[125875]++
											indent_hint := []byte{'0' + byte(emitter.best_indent)}
											if !yaml_emitter_write_indicator(emitter, indent_hint, false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1553
			_go_fuzz_dep_.CoverTab[125876]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1554
			// _ = "end of CoverTab[125876]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1555
			_go_fuzz_dep_.CoverTab[125877]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1555
			// _ = "end of CoverTab[125877]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1555
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1555
		// _ = "end of CoverTab[125875]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1556
		_go_fuzz_dep_.CoverTab[125878]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1556
		// _ = "end of CoverTab[125878]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1556
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1556
	// _ = "end of CoverTab[125870]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1556
	_go_fuzz_dep_.CoverTab[125871]++

										emitter.open_ended = false

										var chomp_hint [1]byte
										if len(value) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1561
		_go_fuzz_dep_.CoverTab[125879]++
											chomp_hint[0] = '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1562
		// _ = "end of CoverTab[125879]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1563
		_go_fuzz_dep_.CoverTab[125880]++
											i := len(value) - 1
											for value[i]&0xC0 == 0x80 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1565
			_go_fuzz_dep_.CoverTab[125882]++
												i--
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1566
			// _ = "end of CoverTab[125882]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1567
		// _ = "end of CoverTab[125880]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1567
		_go_fuzz_dep_.CoverTab[125881]++
											if !is_break(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1568
			_go_fuzz_dep_.CoverTab[125883]++
												chomp_hint[0] = '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1569
			// _ = "end of CoverTab[125883]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1570
			_go_fuzz_dep_.CoverTab[125884]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1570
			if i == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1570
				_go_fuzz_dep_.CoverTab[125885]++
													chomp_hint[0] = '+'
													emitter.open_ended = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1572
				// _ = "end of CoverTab[125885]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1573
				_go_fuzz_dep_.CoverTab[125886]++
													i--
													for value[i]&0xC0 == 0x80 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1575
					_go_fuzz_dep_.CoverTab[125888]++
														i--
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1576
					// _ = "end of CoverTab[125888]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1577
				// _ = "end of CoverTab[125886]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1577
				_go_fuzz_dep_.CoverTab[125887]++
													if is_break(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1578
					_go_fuzz_dep_.CoverTab[125889]++
														chomp_hint[0] = '+'
														emitter.open_ended = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1580
					// _ = "end of CoverTab[125889]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1581
					_go_fuzz_dep_.CoverTab[125890]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1581
					// _ = "end of CoverTab[125890]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1581
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1581
				// _ = "end of CoverTab[125887]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1582
			// _ = "end of CoverTab[125884]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1582
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1582
		// _ = "end of CoverTab[125881]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1583
	// _ = "end of CoverTab[125871]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1583
	_go_fuzz_dep_.CoverTab[125872]++
										if chomp_hint[0] != 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1584
		_go_fuzz_dep_.CoverTab[125891]++
											if !yaml_emitter_write_indicator(emitter, chomp_hint[:], false, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1585
			_go_fuzz_dep_.CoverTab[125892]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1586
			// _ = "end of CoverTab[125892]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1587
			_go_fuzz_dep_.CoverTab[125893]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1587
			// _ = "end of CoverTab[125893]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1587
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1587
		// _ = "end of CoverTab[125891]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1588
		_go_fuzz_dep_.CoverTab[125894]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1588
		// _ = "end of CoverTab[125894]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1588
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1588
	// _ = "end of CoverTab[125872]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1588
	_go_fuzz_dep_.CoverTab[125873]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1589
	// _ = "end of CoverTab[125873]"
}

func yaml_emitter_write_literal_scalar(emitter *yaml_emitter_t, value []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1592
	_go_fuzz_dep_.CoverTab[125895]++
										if !yaml_emitter_write_indicator(emitter, []byte{'|'}, true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1593
		_go_fuzz_dep_.CoverTab[125900]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1594
		// _ = "end of CoverTab[125900]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1595
		_go_fuzz_dep_.CoverTab[125901]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1595
		// _ = "end of CoverTab[125901]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1595
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1595
	// _ = "end of CoverTab[125895]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1595
	_go_fuzz_dep_.CoverTab[125896]++
										if !yaml_emitter_write_block_scalar_hints(emitter, value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1596
		_go_fuzz_dep_.CoverTab[125902]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1597
		// _ = "end of CoverTab[125902]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1598
		_go_fuzz_dep_.CoverTab[125903]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1598
		// _ = "end of CoverTab[125903]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1598
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1598
	// _ = "end of CoverTab[125896]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1598
	_go_fuzz_dep_.CoverTab[125897]++
										if !put_break(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1599
		_go_fuzz_dep_.CoverTab[125904]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1600
		// _ = "end of CoverTab[125904]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1601
		_go_fuzz_dep_.CoverTab[125905]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1601
		// _ = "end of CoverTab[125905]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1601
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1601
	// _ = "end of CoverTab[125897]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1601
	_go_fuzz_dep_.CoverTab[125898]++
										emitter.indention = true
										emitter.whitespace = true
										breaks := true
										for i := 0; i < len(value); {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1605
		_go_fuzz_dep_.CoverTab[125906]++
											if is_break(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1606
			_go_fuzz_dep_.CoverTab[125907]++
												if !write_break(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1607
				_go_fuzz_dep_.CoverTab[125909]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1608
				// _ = "end of CoverTab[125909]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1609
				_go_fuzz_dep_.CoverTab[125910]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1609
				// _ = "end of CoverTab[125910]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1609
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1609
			// _ = "end of CoverTab[125907]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1609
			_go_fuzz_dep_.CoverTab[125908]++
												emitter.indention = true
												breaks = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1611
			// _ = "end of CoverTab[125908]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1612
			_go_fuzz_dep_.CoverTab[125911]++
												if breaks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1613
				_go_fuzz_dep_.CoverTab[125914]++
													if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1614
					_go_fuzz_dep_.CoverTab[125915]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1615
					// _ = "end of CoverTab[125915]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1616
					_go_fuzz_dep_.CoverTab[125916]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1616
					// _ = "end of CoverTab[125916]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1616
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1616
				// _ = "end of CoverTab[125914]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1617
				_go_fuzz_dep_.CoverTab[125917]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1617
				// _ = "end of CoverTab[125917]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1617
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1617
			// _ = "end of CoverTab[125911]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1617
			_go_fuzz_dep_.CoverTab[125912]++
												if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1618
				_go_fuzz_dep_.CoverTab[125918]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1619
				// _ = "end of CoverTab[125918]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1620
				_go_fuzz_dep_.CoverTab[125919]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1620
				// _ = "end of CoverTab[125919]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1620
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1620
			// _ = "end of CoverTab[125912]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1620
			_go_fuzz_dep_.CoverTab[125913]++
												emitter.indention = false
												breaks = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1622
			// _ = "end of CoverTab[125913]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1623
		// _ = "end of CoverTab[125906]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1624
	// _ = "end of CoverTab[125898]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1624
	_go_fuzz_dep_.CoverTab[125899]++

										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1626
	// _ = "end of CoverTab[125899]"
}

func yaml_emitter_write_folded_scalar(emitter *yaml_emitter_t, value []byte) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1629
	_go_fuzz_dep_.CoverTab[125920]++
										if !yaml_emitter_write_indicator(emitter, []byte{'>'}, true, false, false) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1630
		_go_fuzz_dep_.CoverTab[125925]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1631
		// _ = "end of CoverTab[125925]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1632
		_go_fuzz_dep_.CoverTab[125926]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1632
		// _ = "end of CoverTab[125926]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1632
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1632
	// _ = "end of CoverTab[125920]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1632
	_go_fuzz_dep_.CoverTab[125921]++
										if !yaml_emitter_write_block_scalar_hints(emitter, value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1633
		_go_fuzz_dep_.CoverTab[125927]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1634
		// _ = "end of CoverTab[125927]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1635
		_go_fuzz_dep_.CoverTab[125928]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1635
		// _ = "end of CoverTab[125928]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1635
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1635
	// _ = "end of CoverTab[125921]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1635
	_go_fuzz_dep_.CoverTab[125922]++

										if !put_break(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1637
		_go_fuzz_dep_.CoverTab[125929]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1638
		// _ = "end of CoverTab[125929]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1639
		_go_fuzz_dep_.CoverTab[125930]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1639
		// _ = "end of CoverTab[125930]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1639
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1639
	// _ = "end of CoverTab[125922]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1639
	_go_fuzz_dep_.CoverTab[125923]++
										emitter.indention = true
										emitter.whitespace = true

										breaks := true
										leading_spaces := true
										for i := 0; i < len(value); {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1645
		_go_fuzz_dep_.CoverTab[125931]++
											if is_break(value, i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1646
			_go_fuzz_dep_.CoverTab[125932]++
												if !breaks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
				_go_fuzz_dep_.CoverTab[125935]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
				return !leading_spaces
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
				// _ = "end of CoverTab[125935]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
				_go_fuzz_dep_.CoverTab[125936]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
				return value[i] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
				// _ = "end of CoverTab[125936]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1647
				_go_fuzz_dep_.CoverTab[125937]++
													k := 0
													for is_break(value, k) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1649
					_go_fuzz_dep_.CoverTab[125939]++
														k += width(value[k])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1650
					// _ = "end of CoverTab[125939]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1651
				// _ = "end of CoverTab[125937]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1651
				_go_fuzz_dep_.CoverTab[125938]++
													if !is_blankz(value, k) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1652
					_go_fuzz_dep_.CoverTab[125940]++
														if !put_break(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1653
						_go_fuzz_dep_.CoverTab[125941]++
															return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1654
						// _ = "end of CoverTab[125941]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1655
						_go_fuzz_dep_.CoverTab[125942]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1655
						// _ = "end of CoverTab[125942]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1655
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1655
					// _ = "end of CoverTab[125940]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1656
					_go_fuzz_dep_.CoverTab[125943]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1656
					// _ = "end of CoverTab[125943]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1656
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1656
				// _ = "end of CoverTab[125938]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1657
				_go_fuzz_dep_.CoverTab[125944]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1657
				// _ = "end of CoverTab[125944]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1657
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1657
			// _ = "end of CoverTab[125932]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1657
			_go_fuzz_dep_.CoverTab[125933]++
												if !write_break(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1658
				_go_fuzz_dep_.CoverTab[125945]++
													return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1659
				// _ = "end of CoverTab[125945]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1660
				_go_fuzz_dep_.CoverTab[125946]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1660
				// _ = "end of CoverTab[125946]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1660
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1660
			// _ = "end of CoverTab[125933]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1660
			_go_fuzz_dep_.CoverTab[125934]++
												emitter.indention = true
												breaks = true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1662
			// _ = "end of CoverTab[125934]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1663
			_go_fuzz_dep_.CoverTab[125947]++
												if breaks {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1664
				_go_fuzz_dep_.CoverTab[125950]++
													if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1665
					_go_fuzz_dep_.CoverTab[125952]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1666
					// _ = "end of CoverTab[125952]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1667
					_go_fuzz_dep_.CoverTab[125953]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1667
					// _ = "end of CoverTab[125953]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1667
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1667
				// _ = "end of CoverTab[125950]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1667
				_go_fuzz_dep_.CoverTab[125951]++
													leading_spaces = is_blank(value, i)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1668
				// _ = "end of CoverTab[125951]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1669
				_go_fuzz_dep_.CoverTab[125954]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1669
				// _ = "end of CoverTab[125954]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1669
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1669
			// _ = "end of CoverTab[125947]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1669
			_go_fuzz_dep_.CoverTab[125948]++
												if !breaks && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				_go_fuzz_dep_.CoverTab[125955]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				return is_space(value, i)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				// _ = "end of CoverTab[125955]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				_go_fuzz_dep_.CoverTab[125956]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				return !is_space(value, i+1)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				// _ = "end of CoverTab[125956]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
			}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				_go_fuzz_dep_.CoverTab[125957]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				return emitter.column > emitter.best_width
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				// _ = "end of CoverTab[125957]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1670
				_go_fuzz_dep_.CoverTab[125958]++
													if !yaml_emitter_write_indent(emitter) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1671
					_go_fuzz_dep_.CoverTab[125960]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1672
					// _ = "end of CoverTab[125960]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1673
					_go_fuzz_dep_.CoverTab[125961]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1673
					// _ = "end of CoverTab[125961]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1673
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1673
				// _ = "end of CoverTab[125958]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1673
				_go_fuzz_dep_.CoverTab[125959]++
													i += width(value[i])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1674
				// _ = "end of CoverTab[125959]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1675
				_go_fuzz_dep_.CoverTab[125962]++
													if !write(emitter, value, &i) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1676
					_go_fuzz_dep_.CoverTab[125963]++
														return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1677
					// _ = "end of CoverTab[125963]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1678
					_go_fuzz_dep_.CoverTab[125964]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1678
					// _ = "end of CoverTab[125964]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1678
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1678
				// _ = "end of CoverTab[125962]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1679
			// _ = "end of CoverTab[125948]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1679
			_go_fuzz_dep_.CoverTab[125949]++
												emitter.indention = false
												breaks = false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1681
			// _ = "end of CoverTab[125949]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1682
		// _ = "end of CoverTab[125931]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1683
	// _ = "end of CoverTab[125923]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1683
	_go_fuzz_dep_.CoverTab[125924]++
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1684
	// _ = "end of CoverTab[125924]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1685
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/emitterc.go:1685
var _ = _go_fuzz_dep_.CoverTab
