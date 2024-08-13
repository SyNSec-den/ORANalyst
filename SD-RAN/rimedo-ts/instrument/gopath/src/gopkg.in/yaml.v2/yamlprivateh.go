//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:1
)

const (
	// The size of the input raw buffer.
	input_raw_buffer_size	= 512

	// The size of the input buffer.
	// It should be possible to decode the whole raw buffer.
	input_buffer_size	= input_raw_buffer_size * 3

	// The size of the output buffer.
	output_buffer_size	= 128

	// The size of the output raw buffer.
	// It should be possible to encode the whole output buffer.
	output_raw_buffer_size	= (output_buffer_size*2 + 2)

	// The size of other stacks and queues.
	initial_stack_size	= 16
	initial_queue_size	= 16
	initial_string_size	= 16
)

// Check if the character at the specified position is an alphabetical
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:24
// character, a digit, '_', or '-'.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:26
func is_alpha(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:26
	_go_fuzz_dep_.CoverTab[128032]++
										return b[i] >= '0' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		_go_fuzz_dep_.CoverTab[128033]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		return b[i] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		// _ = "end of CoverTab[128033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		_go_fuzz_dep_.CoverTab[128034]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		return b[i] >= 'A' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
			_go_fuzz_dep_.CoverTab[128035]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
			return b[i] <= 'Z'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
			// _ = "end of CoverTab[128035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		// _ = "end of CoverTab[128034]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		_go_fuzz_dep_.CoverTab[128036]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		return b[i] >= 'a' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
			_go_fuzz_dep_.CoverTab[128037]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
			return b[i] <= 'z'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
			// _ = "end of CoverTab[128037]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		// _ = "end of CoverTab[128036]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		_go_fuzz_dep_.CoverTab[128038]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		return b[i] == '_'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		// _ = "end of CoverTab[128038]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		_go_fuzz_dep_.CoverTab[128039]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		return b[i] == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
		// _ = "end of CoverTab[128039]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:27
	// _ = "end of CoverTab[128032]"
}

// Check if the character at the specified position is a digit.
func is_digit(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:31
	_go_fuzz_dep_.CoverTab[128040]++
										return b[i] >= '0' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:32
		_go_fuzz_dep_.CoverTab[128041]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:32
		return b[i] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:32
		// _ = "end of CoverTab[128041]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:32
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:32
	// _ = "end of CoverTab[128040]"
}

// Get the value of a digit.
func as_digit(b []byte, i int) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:36
	_go_fuzz_dep_.CoverTab[128042]++
										return int(b[i]) - '0'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:37
	// _ = "end of CoverTab[128042]"
}

// Check if the character at the specified position is a hex-digit.
func is_hex(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:41
	_go_fuzz_dep_.CoverTab[128043]++
										return b[i] >= '0' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		_go_fuzz_dep_.CoverTab[128044]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		return b[i] <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		// _ = "end of CoverTab[128044]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		_go_fuzz_dep_.CoverTab[128045]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		return b[i] >= 'A' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
			_go_fuzz_dep_.CoverTab[128046]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
			return b[i] <= 'F'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
			// _ = "end of CoverTab[128046]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		// _ = "end of CoverTab[128045]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		_go_fuzz_dep_.CoverTab[128047]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		return b[i] >= 'a' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
			_go_fuzz_dep_.CoverTab[128048]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
			return b[i] <= 'f'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
			// _ = "end of CoverTab[128048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
		// _ = "end of CoverTab[128047]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:42
	// _ = "end of CoverTab[128043]"
}

// Get the value of a hex-digit.
func as_hex(b []byte, i int) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:46
	_go_fuzz_dep_.CoverTab[128049]++
										bi := b[i]
										if bi >= 'A' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:48
		_go_fuzz_dep_.CoverTab[128052]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:48
		return bi <= 'F'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:48
		// _ = "end of CoverTab[128052]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:48
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:48
		_go_fuzz_dep_.CoverTab[128053]++
											return int(bi) - 'A' + 10
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:49
		// _ = "end of CoverTab[128053]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:50
		_go_fuzz_dep_.CoverTab[128054]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:50
		// _ = "end of CoverTab[128054]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:50
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:50
	// _ = "end of CoverTab[128049]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:50
	_go_fuzz_dep_.CoverTab[128050]++
										if bi >= 'a' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:51
		_go_fuzz_dep_.CoverTab[128055]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:51
		return bi <= 'f'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:51
		// _ = "end of CoverTab[128055]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:51
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:51
		_go_fuzz_dep_.CoverTab[128056]++
											return int(bi) - 'a' + 10
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:52
		// _ = "end of CoverTab[128056]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:53
		_go_fuzz_dep_.CoverTab[128057]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:53
		// _ = "end of CoverTab[128057]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:53
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:53
	// _ = "end of CoverTab[128050]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:53
	_go_fuzz_dep_.CoverTab[128051]++
										return int(bi) - '0'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:54
	// _ = "end of CoverTab[128051]"
}

// Check if the character is ASCII.
func is_ascii(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:58
	_go_fuzz_dep_.CoverTab[128058]++
										return b[i] <= 0x7F
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:59
	// _ = "end of CoverTab[128058]"
}

// Check if the character at the start of the buffer can be printed unescaped.
func is_printable(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:63
	_go_fuzz_dep_.CoverTab[128059]++
										return ((b[i] == 0x0A) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:64
		_go_fuzz_dep_.CoverTab[128060]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:64
		return (b[i] >= 0x20 && func() bool {
												_go_fuzz_dep_.CoverTab[128061]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:65
			return b[i] <= 0x7E
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:65
			// _ = "end of CoverTab[128061]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:65
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:65
		// _ = "end of CoverTab[128060]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:65
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:65
		_go_fuzz_dep_.CoverTab[128062]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:65
		return (b[i] == 0xC2 && func() bool {
												_go_fuzz_dep_.CoverTab[128063]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:66
			return b[i+1] >= 0xA0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:66
			// _ = "end of CoverTab[128063]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:66
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:66
		// _ = "end of CoverTab[128062]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:66
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:66
		_go_fuzz_dep_.CoverTab[128064]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:66
		return (b[i] > 0xC2 && func() bool {
												_go_fuzz_dep_.CoverTab[128065]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:67
			return b[i] < 0xED
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:67
			// _ = "end of CoverTab[128065]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:67
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:67
		// _ = "end of CoverTab[128064]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:67
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:67
		_go_fuzz_dep_.CoverTab[128066]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:67
		return (b[i] == 0xED && func() bool {
												_go_fuzz_dep_.CoverTab[128067]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:68
			return b[i+1] < 0xA0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:68
			// _ = "end of CoverTab[128067]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:68
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:68
		// _ = "end of CoverTab[128066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:68
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:68
		_go_fuzz_dep_.CoverTab[128068]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:68
		return (b[i] == 0xEE)
											// _ = "end of CoverTab[128068]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:69
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:69
		_go_fuzz_dep_.CoverTab[128069]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:69
		return (b[i] == 0xEF && func() bool {
												_go_fuzz_dep_.CoverTab[128070]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:70
			return !(b[i+1] == 0xBB && func() bool {
													_go_fuzz_dep_.CoverTab[128071]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:71
				return b[i+2] == 0xBF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:71
				// _ = "end of CoverTab[128071]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:71
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:71
			// _ = "end of CoverTab[128070]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:71
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:71
			_go_fuzz_dep_.CoverTab[128072]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:71
			return !(b[i+1] == 0xBF && func() bool {
													_go_fuzz_dep_.CoverTab[128073]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
				return (b[i+2] == 0xBE || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
					_go_fuzz_dep_.CoverTab[128074]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
					return b[i+2] == 0xBF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
					// _ = "end of CoverTab[128074]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
				}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
				// _ = "end of CoverTab[128073]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
			}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
			// _ = "end of CoverTab[128072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
		}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
		// _ = "end of CoverTab[128069]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
	}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:72
	// _ = "end of CoverTab[128059]"
}

// Check if the character at the specified position is NUL.
func is_z(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:76
	_go_fuzz_dep_.CoverTab[128075]++
										return b[i] == 0x00
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:77
	// _ = "end of CoverTab[128075]"
}

// Check if the beginning of the buffer is a BOM.
func is_bom(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:81
	_go_fuzz_dep_.CoverTab[128076]++
										return b[0] == 0xEF && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
		_go_fuzz_dep_.CoverTab[128077]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
		return b[1] == 0xBB
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
		// _ = "end of CoverTab[128077]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
		_go_fuzz_dep_.CoverTab[128078]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
		return b[2] == 0xBF
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
		// _ = "end of CoverTab[128078]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:82
	// _ = "end of CoverTab[128076]"
}

// Check if the character at the specified position is space.
func is_space(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:86
	_go_fuzz_dep_.CoverTab[128079]++
										return b[i] == ' '
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:87
	// _ = "end of CoverTab[128079]"
}

// Check if the character at the specified position is tab.
func is_tab(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:91
	_go_fuzz_dep_.CoverTab[128080]++
										return b[i] == '\t'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:92
	// _ = "end of CoverTab[128080]"
}

// Check if the character at the specified position is blank (space or tab).
func is_blank(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:96
	_go_fuzz_dep_.CoverTab[128081]++

										return b[i] == ' ' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:98
		_go_fuzz_dep_.CoverTab[128082]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:98
		return b[i] == '\t'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:98
		// _ = "end of CoverTab[128082]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:98
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:98
	// _ = "end of CoverTab[128081]"
}

// Check if the character at the specified position is a line break.
func is_break(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:102
	_go_fuzz_dep_.CoverTab[128083]++
										return (b[i] == '\r' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:103
		_go_fuzz_dep_.CoverTab[128084]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:103
		return b[i] == '\n'
											// _ = "end of CoverTab[128084]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:104
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:104
		_go_fuzz_dep_.CoverTab[128085]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:104
		return b[i] == 0xC2 && func() bool {
												_go_fuzz_dep_.CoverTab[128086]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:105
			return b[i+1] == 0x85
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:105
			// _ = "end of CoverTab[128086]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:105
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:105
		// _ = "end of CoverTab[128085]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:105
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:105
		_go_fuzz_dep_.CoverTab[128087]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:105
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128088]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
			// _ = "end of CoverTab[128088]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
			_go_fuzz_dep_.CoverTab[128089]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
			return b[i+2] == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
			// _ = "end of CoverTab[128089]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
		// _ = "end of CoverTab[128087]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
		_go_fuzz_dep_.CoverTab[128090]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:106
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128091]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
			// _ = "end of CoverTab[128091]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
			_go_fuzz_dep_.CoverTab[128092]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
			return b[i+2] == 0xA9
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
			// _ = "end of CoverTab[128092]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
		// _ = "end of CoverTab[128090]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
	}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:107
	// _ = "end of CoverTab[128083]"
}

func is_crlf(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:110
	_go_fuzz_dep_.CoverTab[128093]++
										return b[i] == '\r' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:111
		_go_fuzz_dep_.CoverTab[128094]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:111
		return b[i+1] == '\n'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:111
		// _ = "end of CoverTab[128094]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:111
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:111
	// _ = "end of CoverTab[128093]"
}

// Check if the character is a line break or NUL.
func is_breakz(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:115
	_go_fuzz_dep_.CoverTab[128095]++

										return (b[i] == '\r' || func() bool {
											_go_fuzz_dep_.CoverTab[128096]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:118
		return b[i] == '\n'
											// _ = "end of CoverTab[128096]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:119
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:119
		_go_fuzz_dep_.CoverTab[128097]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:119
		return b[i] == 0xC2 && func() bool {
												_go_fuzz_dep_.CoverTab[128098]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:120
			return b[i+1] == 0x85
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:120
			// _ = "end of CoverTab[128098]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:120
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:120
		// _ = "end of CoverTab[128097]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:120
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:120
		_go_fuzz_dep_.CoverTab[128099]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:120
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128100]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
			// _ = "end of CoverTab[128100]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
			_go_fuzz_dep_.CoverTab[128101]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
			return b[i+2] == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
			// _ = "end of CoverTab[128101]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
		// _ = "end of CoverTab[128099]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
		_go_fuzz_dep_.CoverTab[128102]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:121
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128103]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
			// _ = "end of CoverTab[128103]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
			_go_fuzz_dep_.CoverTab[128104]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
			return b[i+2] == 0xA9
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
			// _ = "end of CoverTab[128104]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
		// _ = "end of CoverTab[128102]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
		_go_fuzz_dep_.CoverTab[128105]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:122
		return b[i] == 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:124
		// _ = "end of CoverTab[128105]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:124
	}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:124
	// _ = "end of CoverTab[128095]"
}

// Check if the character is a line break, space, or NUL.
func is_spacez(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:128
	_go_fuzz_dep_.CoverTab[128106]++

										return (b[i] == ' ' || func() bool {
											_go_fuzz_dep_.CoverTab[128107]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:131
		return b[i] == '\r'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:133
		// _ = "end of CoverTab[128107]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:133
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:133
		_go_fuzz_dep_.CoverTab[128108]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:133
		return b[i] == '\n'
											// _ = "end of CoverTab[128108]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:134
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:134
		_go_fuzz_dep_.CoverTab[128109]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:134
		return b[i] == 0xC2 && func() bool {
												_go_fuzz_dep_.CoverTab[128110]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:135
			return b[i+1] == 0x85
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:135
			// _ = "end of CoverTab[128110]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:135
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:135
		// _ = "end of CoverTab[128109]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:135
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:135
		_go_fuzz_dep_.CoverTab[128111]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:135
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128112]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
			// _ = "end of CoverTab[128112]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
			_go_fuzz_dep_.CoverTab[128113]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
			return b[i+2] == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
			// _ = "end of CoverTab[128113]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
		// _ = "end of CoverTab[128111]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
		_go_fuzz_dep_.CoverTab[128114]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:136
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128115]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
			// _ = "end of CoverTab[128115]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
			_go_fuzz_dep_.CoverTab[128116]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
			return b[i+2] == 0xA9
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
			// _ = "end of CoverTab[128116]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
		// _ = "end of CoverTab[128114]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
		_go_fuzz_dep_.CoverTab[128117]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:137
		return b[i] == 0
											// _ = "end of CoverTab[128117]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:138
	}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:138
	// _ = "end of CoverTab[128106]"
}

// Check if the character is a line break, space, tab, or NUL.
func is_blankz(b []byte, i int) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:142
	_go_fuzz_dep_.CoverTab[128118]++

										return (b[i] == ' ' || func() bool {
											_go_fuzz_dep_.CoverTab[128119]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:145
		return b[i] == '\t'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:145
		// _ = "end of CoverTab[128119]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:145
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:145
		_go_fuzz_dep_.CoverTab[128120]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:145
		return b[i] == '\r'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:147
		// _ = "end of CoverTab[128120]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:147
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:147
		_go_fuzz_dep_.CoverTab[128121]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:147
		return b[i] == '\n'
											// _ = "end of CoverTab[128121]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:148
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:148
		_go_fuzz_dep_.CoverTab[128122]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:148
		return b[i] == 0xC2 && func() bool {
												_go_fuzz_dep_.CoverTab[128123]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:149
			return b[i+1] == 0x85
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:149
			// _ = "end of CoverTab[128123]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:149
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:149
		// _ = "end of CoverTab[128122]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:149
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:149
		_go_fuzz_dep_.CoverTab[128124]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:149
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128125]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
			// _ = "end of CoverTab[128125]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
			_go_fuzz_dep_.CoverTab[128126]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
			return b[i+2] == 0xA8
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
			// _ = "end of CoverTab[128126]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
		// _ = "end of CoverTab[128124]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
		_go_fuzz_dep_.CoverTab[128127]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:150
		return b[i] == 0xE2 && func() bool {
												_go_fuzz_dep_.CoverTab[128128]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
			return b[i+1] == 0x80
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
			// _ = "end of CoverTab[128128]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
		}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
			_go_fuzz_dep_.CoverTab[128129]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
			return b[i+2] == 0xA9
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
			// _ = "end of CoverTab[128129]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
		// _ = "end of CoverTab[128127]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
		_go_fuzz_dep_.CoverTab[128130]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:151
		return b[i] == 0
											// _ = "end of CoverTab[128130]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:152
	}())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:152
	// _ = "end of CoverTab[128118]"
}

// Determine the width of the character.
func width(b byte) int {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:156
	_go_fuzz_dep_.CoverTab[128131]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:159
	if b&0x80 == 0x00 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:159
		_go_fuzz_dep_.CoverTab[128136]++
											return 1
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:160
		// _ = "end of CoverTab[128136]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:161
		_go_fuzz_dep_.CoverTab[128137]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:161
		// _ = "end of CoverTab[128137]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:161
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:161
	// _ = "end of CoverTab[128131]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:161
	_go_fuzz_dep_.CoverTab[128132]++
										if b&0xE0 == 0xC0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:162
		_go_fuzz_dep_.CoverTab[128138]++
											return 2
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:163
		// _ = "end of CoverTab[128138]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:164
		_go_fuzz_dep_.CoverTab[128139]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:164
		// _ = "end of CoverTab[128139]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:164
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:164
	// _ = "end of CoverTab[128132]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:164
	_go_fuzz_dep_.CoverTab[128133]++
										if b&0xF0 == 0xE0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:165
		_go_fuzz_dep_.CoverTab[128140]++
											return 3
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:166
		// _ = "end of CoverTab[128140]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:167
		_go_fuzz_dep_.CoverTab[128141]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:167
		// _ = "end of CoverTab[128141]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:167
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:167
	// _ = "end of CoverTab[128133]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:167
	_go_fuzz_dep_.CoverTab[128134]++
										if b&0xF8 == 0xF0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:168
		_go_fuzz_dep_.CoverTab[128142]++
											return 4
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:169
		// _ = "end of CoverTab[128142]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:170
		_go_fuzz_dep_.CoverTab[128143]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:170
		// _ = "end of CoverTab[128143]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:170
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:170
	// _ = "end of CoverTab[128134]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:170
	_go_fuzz_dep_.CoverTab[128135]++
										return 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:171
	// _ = "end of CoverTab[128135]"

}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:173
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/yamlprivateh.go:173
var _ = _go_fuzz_dep_.CoverTab
