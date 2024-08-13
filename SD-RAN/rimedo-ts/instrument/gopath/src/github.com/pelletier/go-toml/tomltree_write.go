//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:1
)

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/big"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type valueComplexity int

const (
	valueSimple	valueComplexity	= iota + 1
	valueComplex
)

type sortNode struct {
	key		string
	complexity	valueComplexity
}

// Encodes a string to a TOML-compliant multi-line string value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:28
// This function is a clone of the existing encodeTomlString function, except that whitespace characters
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:28
// are preserved. Quotation marks and backslashes are also not escaped.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:31
func encodeMultilineTomlString(value string, commented string) string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:31
	_go_fuzz_dep_.CoverTab[124190]++
												var b bytes.Buffer
												adjacentQuoteCount := 0

												b.WriteString(commented)
												for i, rr := range value {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:36
		_go_fuzz_dep_.CoverTab[124192]++
													if rr != '"' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:37
			_go_fuzz_dep_.CoverTab[124194]++
														adjacentQuoteCount = 0
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:38
			// _ = "end of CoverTab[124194]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:39
			_go_fuzz_dep_.CoverTab[124195]++
														adjacentQuoteCount++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:40
			// _ = "end of CoverTab[124195]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:41
		// _ = "end of CoverTab[124192]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:41
		_go_fuzz_dep_.CoverTab[124193]++
													switch rr {
		case '\b':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:43
			_go_fuzz_dep_.CoverTab[124196]++
														b.WriteString(`\b`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:44
			// _ = "end of CoverTab[124196]"
		case '\t':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:45
			_go_fuzz_dep_.CoverTab[124197]++
														b.WriteString("\t")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:46
			// _ = "end of CoverTab[124197]"
		case '\n':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:47
			_go_fuzz_dep_.CoverTab[124198]++
														b.WriteString("\n" + commented)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:48
			// _ = "end of CoverTab[124198]"
		case '\f':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:49
			_go_fuzz_dep_.CoverTab[124199]++
														b.WriteString(`\f`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:50
			// _ = "end of CoverTab[124199]"
		case '\r':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:51
			_go_fuzz_dep_.CoverTab[124200]++
														b.WriteString("\r")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:52
			// _ = "end of CoverTab[124200]"
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:53
			_go_fuzz_dep_.CoverTab[124201]++
														if adjacentQuoteCount >= 3 || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:54
				_go_fuzz_dep_.CoverTab[124204]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:54
				return i == len(value)-1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:54
				// _ = "end of CoverTab[124204]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:54
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:54
				_go_fuzz_dep_.CoverTab[124205]++
															adjacentQuoteCount = 0
															b.WriteString(`\"`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:56
				// _ = "end of CoverTab[124205]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:57
				_go_fuzz_dep_.CoverTab[124206]++
															b.WriteString(`"`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:58
				// _ = "end of CoverTab[124206]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:59
			// _ = "end of CoverTab[124201]"
		case '\\':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:60
			_go_fuzz_dep_.CoverTab[124202]++
														b.WriteString(`\`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:61
			// _ = "end of CoverTab[124202]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:62
			_go_fuzz_dep_.CoverTab[124203]++
														intRr := uint16(rr)
														if intRr < 0x001F {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:64
				_go_fuzz_dep_.CoverTab[124207]++
															b.WriteString(fmt.Sprintf("\\u%0.4X", intRr))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:65
				// _ = "end of CoverTab[124207]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:66
				_go_fuzz_dep_.CoverTab[124208]++
															b.WriteRune(rr)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:67
				// _ = "end of CoverTab[124208]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:68
			// _ = "end of CoverTab[124203]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:69
		// _ = "end of CoverTab[124193]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:70
	// _ = "end of CoverTab[124190]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:70
	_go_fuzz_dep_.CoverTab[124191]++
												return b.String()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:71
	// _ = "end of CoverTab[124191]"
}

// Encodes a string to a TOML-compliant string value
func encodeTomlString(value string) string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:75
	_go_fuzz_dep_.CoverTab[124209]++
												var b bytes.Buffer

												for _, rr := range value {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:78
		_go_fuzz_dep_.CoverTab[124211]++
													switch rr {
		case '\b':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:80
			_go_fuzz_dep_.CoverTab[124212]++
														b.WriteString(`\b`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:81
			// _ = "end of CoverTab[124212]"
		case '\t':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:82
			_go_fuzz_dep_.CoverTab[124213]++
														b.WriteString(`\t`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:83
			// _ = "end of CoverTab[124213]"
		case '\n':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:84
			_go_fuzz_dep_.CoverTab[124214]++
														b.WriteString(`\n`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:85
			// _ = "end of CoverTab[124214]"
		case '\f':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:86
			_go_fuzz_dep_.CoverTab[124215]++
														b.WriteString(`\f`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:87
			// _ = "end of CoverTab[124215]"
		case '\r':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:88
			_go_fuzz_dep_.CoverTab[124216]++
														b.WriteString(`\r`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:89
			// _ = "end of CoverTab[124216]"
		case '"':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:90
			_go_fuzz_dep_.CoverTab[124217]++
														b.WriteString(`\"`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:91
			// _ = "end of CoverTab[124217]"
		case '\\':
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:92
			_go_fuzz_dep_.CoverTab[124218]++
														b.WriteString(`\\`)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:93
			// _ = "end of CoverTab[124218]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:94
			_go_fuzz_dep_.CoverTab[124219]++
														intRr := uint16(rr)
														if intRr < 0x001F {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:96
				_go_fuzz_dep_.CoverTab[124220]++
															b.WriteString(fmt.Sprintf("\\u%0.4X", intRr))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:97
				// _ = "end of CoverTab[124220]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:98
				_go_fuzz_dep_.CoverTab[124221]++
															b.WriteRune(rr)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:99
				// _ = "end of CoverTab[124221]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:100
			// _ = "end of CoverTab[124219]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:101
		// _ = "end of CoverTab[124211]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:102
	// _ = "end of CoverTab[124209]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:102
	_go_fuzz_dep_.CoverTab[124210]++
												return b.String()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:103
	// _ = "end of CoverTab[124210]"
}

func tomlTreeStringRepresentation(t *Tree, ord MarshalOrder) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:106
	_go_fuzz_dep_.CoverTab[124222]++
												var orderedVals []sortNode
												switch ord {
	case OrderPreserve:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:109
		_go_fuzz_dep_.CoverTab[124225]++
													orderedVals = sortByLines(t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:110
		// _ = "end of CoverTab[124225]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:111
		_go_fuzz_dep_.CoverTab[124226]++
													orderedVals = sortAlphabetical(t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:112
		// _ = "end of CoverTab[124226]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:113
	// _ = "end of CoverTab[124222]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:113
	_go_fuzz_dep_.CoverTab[124223]++

												var values []string
												for _, node := range orderedVals {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:116
		_go_fuzz_dep_.CoverTab[124227]++
													k := node.key
													v := t.values[k]

													repr, err := tomlValueStringRepresentation(v, "", "", ord, false)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:121
			_go_fuzz_dep_.CoverTab[124229]++
														return "", err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:122
			// _ = "end of CoverTab[124229]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:123
			_go_fuzz_dep_.CoverTab[124230]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:123
			// _ = "end of CoverTab[124230]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:123
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:123
		// _ = "end of CoverTab[124227]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:123
		_go_fuzz_dep_.CoverTab[124228]++
													values = append(values, quoteKeyIfNeeded(k)+" = "+repr)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:124
		// _ = "end of CoverTab[124228]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:125
	// _ = "end of CoverTab[124223]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:125
	_go_fuzz_dep_.CoverTab[124224]++
												return "{ " + strings.Join(values, ", ") + " }", nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:126
	// _ = "end of CoverTab[124224]"
}

func tomlValueStringRepresentation(v interface{}, commented string, indent string, ord MarshalOrder, arraysOneElementPerLine bool) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:129
	_go_fuzz_dep_.CoverTab[124231]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:132
	tv, ok := v.(*tomlValue)
	if ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:133
		_go_fuzz_dep_.CoverTab[124235]++
													v = tv.value
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:134
		// _ = "end of CoverTab[124235]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:135
		_go_fuzz_dep_.CoverTab[124236]++
													tv = &tomlValue{}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:136
		// _ = "end of CoverTab[124236]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:137
	// _ = "end of CoverTab[124231]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:137
	_go_fuzz_dep_.CoverTab[124232]++

												switch value := v.(type) {
	case uint64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:140
		_go_fuzz_dep_.CoverTab[124237]++
													return strconv.FormatUint(value, 10), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:141
		// _ = "end of CoverTab[124237]"
	case int64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:142
		_go_fuzz_dep_.CoverTab[124238]++
													return strconv.FormatInt(value, 10), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:143
		// _ = "end of CoverTab[124238]"
	case float64:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:144
		_go_fuzz_dep_.CoverTab[124239]++

													bits := 64

													if !math.IsNaN(value) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:148
			_go_fuzz_dep_.CoverTab[124253]++

														_, acc := big.NewFloat(value).Float32()
														if acc == big.Exact {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:151
				_go_fuzz_dep_.CoverTab[124254]++
															bits = 32
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:152
				// _ = "end of CoverTab[124254]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:153
				_go_fuzz_dep_.CoverTab[124255]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:153
				// _ = "end of CoverTab[124255]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:153
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:153
			// _ = "end of CoverTab[124253]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:154
			_go_fuzz_dep_.CoverTab[124256]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:154
			// _ = "end of CoverTab[124256]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:154
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:154
		// _ = "end of CoverTab[124239]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:154
		_go_fuzz_dep_.CoverTab[124240]++
													if math.Trunc(value) == value {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:155
			_go_fuzz_dep_.CoverTab[124257]++
														return strings.ToLower(strconv.FormatFloat(value, 'f', 1, bits)), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:156
			// _ = "end of CoverTab[124257]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:157
			_go_fuzz_dep_.CoverTab[124258]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:157
			// _ = "end of CoverTab[124258]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:157
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:157
		// _ = "end of CoverTab[124240]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:157
		_go_fuzz_dep_.CoverTab[124241]++
													return strings.ToLower(strconv.FormatFloat(value, 'f', -1, bits)), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:158
		// _ = "end of CoverTab[124241]"
	case string:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:159
		_go_fuzz_dep_.CoverTab[124242]++
													if tv.multiline {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:160
			_go_fuzz_dep_.CoverTab[124259]++
														if tv.literal {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:161
				_go_fuzz_dep_.CoverTab[124260]++
															b := strings.Builder{}
															b.WriteString("'''\n")
															b.Write([]byte(value))
															b.WriteString("\n'''")
															return b.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:166
				// _ = "end of CoverTab[124260]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:167
				_go_fuzz_dep_.CoverTab[124261]++
															return "\"\"\"\n" + encodeMultilineTomlString(value, commented) + "\"\"\"", nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:168
				// _ = "end of CoverTab[124261]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:169
			// _ = "end of CoverTab[124259]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:170
			_go_fuzz_dep_.CoverTab[124262]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:170
			// _ = "end of CoverTab[124262]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:170
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:170
		// _ = "end of CoverTab[124242]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:170
		_go_fuzz_dep_.CoverTab[124243]++
													return "\"" + encodeTomlString(value) + "\"", nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:171
		// _ = "end of CoverTab[124243]"
	case []byte:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:172
		_go_fuzz_dep_.CoverTab[124244]++
													b, _ := v.([]byte)
													return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:174
		// _ = "end of CoverTab[124244]"
	case bool:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:175
		_go_fuzz_dep_.CoverTab[124245]++
													if value {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:176
			_go_fuzz_dep_.CoverTab[124263]++
														return "true", nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:177
			// _ = "end of CoverTab[124263]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:178
			_go_fuzz_dep_.CoverTab[124264]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:178
			// _ = "end of CoverTab[124264]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:178
		// _ = "end of CoverTab[124245]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:178
		_go_fuzz_dep_.CoverTab[124246]++
													return "false", nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:179
		// _ = "end of CoverTab[124246]"
	case time.Time:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:180
		_go_fuzz_dep_.CoverTab[124247]++
													return value.Format(time.RFC3339), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:181
		// _ = "end of CoverTab[124247]"
	case LocalDate:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:182
		_go_fuzz_dep_.CoverTab[124248]++
													return value.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:183
		// _ = "end of CoverTab[124248]"
	case LocalDateTime:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:184
		_go_fuzz_dep_.CoverTab[124249]++
													return value.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:185
		// _ = "end of CoverTab[124249]"
	case LocalTime:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:186
		_go_fuzz_dep_.CoverTab[124250]++
													return value.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:187
		// _ = "end of CoverTab[124250]"
	case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:188
		_go_fuzz_dep_.CoverTab[124251]++
													return tomlTreeStringRepresentation(value, ord)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:189
		// _ = "end of CoverTab[124251]"
	case nil:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:190
		_go_fuzz_dep_.CoverTab[124252]++
													return "", nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:191
		// _ = "end of CoverTab[124252]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:192
	// _ = "end of CoverTab[124232]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:192
	_go_fuzz_dep_.CoverTab[124233]++

												rv := reflect.ValueOf(v)

												if rv.Kind() == reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:196
		_go_fuzz_dep_.CoverTab[124265]++
													var values []string
													for i := 0; i < rv.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:198
			_go_fuzz_dep_.CoverTab[124268]++
														item := rv.Index(i).Interface()
														itemRepr, err := tomlValueStringRepresentation(item, commented, indent, ord, arraysOneElementPerLine)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:201
				_go_fuzz_dep_.CoverTab[124270]++
															return "", err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:202
				// _ = "end of CoverTab[124270]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:203
				_go_fuzz_dep_.CoverTab[124271]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:203
				// _ = "end of CoverTab[124271]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:203
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:203
			// _ = "end of CoverTab[124268]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:203
			_go_fuzz_dep_.CoverTab[124269]++
														values = append(values, itemRepr)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:204
			// _ = "end of CoverTab[124269]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:205
		// _ = "end of CoverTab[124265]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:205
		_go_fuzz_dep_.CoverTab[124266]++
													if arraysOneElementPerLine && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:206
			_go_fuzz_dep_.CoverTab[124272]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:206
			return len(values) > 1
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:206
			// _ = "end of CoverTab[124272]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:206
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:206
			_go_fuzz_dep_.CoverTab[124273]++
														stringBuffer := bytes.Buffer{}
														valueIndent := indent + `  `

														stringBuffer.WriteString("[\n")

														for _, value := range values {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:212
				_go_fuzz_dep_.CoverTab[124275]++
															stringBuffer.WriteString(valueIndent)
															stringBuffer.WriteString(commented + value)
															stringBuffer.WriteString(`,`)
															stringBuffer.WriteString("\n")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:216
				// _ = "end of CoverTab[124275]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:217
			// _ = "end of CoverTab[124273]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:217
			_go_fuzz_dep_.CoverTab[124274]++

														stringBuffer.WriteString(indent + commented + "]")

														return stringBuffer.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:221
			// _ = "end of CoverTab[124274]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:222
			_go_fuzz_dep_.CoverTab[124276]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:222
			// _ = "end of CoverTab[124276]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:222
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:222
		// _ = "end of CoverTab[124266]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:222
		_go_fuzz_dep_.CoverTab[124267]++
													return "[" + strings.Join(values, ", ") + "]", nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:223
		// _ = "end of CoverTab[124267]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:224
		_go_fuzz_dep_.CoverTab[124277]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:224
		// _ = "end of CoverTab[124277]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:224
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:224
	// _ = "end of CoverTab[124233]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:224
	_go_fuzz_dep_.CoverTab[124234]++
												return "", fmt.Errorf("unsupported value type %T: %v", v, v)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:225
	// _ = "end of CoverTab[124234]"
}

func getTreeArrayLine(trees []*Tree) (line int) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:228
	_go_fuzz_dep_.CoverTab[124278]++

												line = int(^uint(0) >> 1)

												for _, tv := range trees {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:232
		_go_fuzz_dep_.CoverTab[124280]++
													if tv.position.Line < line || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:233
			_go_fuzz_dep_.CoverTab[124281]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:233
			return line == 0
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:233
			// _ = "end of CoverTab[124281]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:233
		}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:233
			_go_fuzz_dep_.CoverTab[124282]++
														line = tv.position.Line
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:234
			// _ = "end of CoverTab[124282]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:235
			_go_fuzz_dep_.CoverTab[124283]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:235
			// _ = "end of CoverTab[124283]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:235
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:235
		// _ = "end of CoverTab[124280]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:236
	// _ = "end of CoverTab[124278]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:236
	_go_fuzz_dep_.CoverTab[124279]++
												return
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:237
	// _ = "end of CoverTab[124279]"
}

func sortByLines(t *Tree) (vals []sortNode) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:240
	_go_fuzz_dep_.CoverTab[124284]++
												var (
		line	int
		lines	[]int
		tv	*Tree
		tom	*tomlValue
		node	sortNode
	)
	vals = make([]sortNode, 0)
	m := make(map[int]sortNode)

	for k := range t.values {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:251
		_go_fuzz_dep_.CoverTab[124287]++
													v := t.values[k]
													switch v.(type) {
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:254
			_go_fuzz_dep_.CoverTab[124289]++
														tv = v.(*Tree)
														line = tv.position.Line
														node = sortNode{key: k, complexity: valueComplex}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:257
			// _ = "end of CoverTab[124289]"
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:258
			_go_fuzz_dep_.CoverTab[124290]++
														line = getTreeArrayLine(v.([]*Tree))
														node = sortNode{key: k, complexity: valueComplex}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:260
			// _ = "end of CoverTab[124290]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:261
			_go_fuzz_dep_.CoverTab[124291]++
														tom = v.(*tomlValue)
														line = tom.position.Line
														node = sortNode{key: k, complexity: valueSimple}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:264
			// _ = "end of CoverTab[124291]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:265
		// _ = "end of CoverTab[124287]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:265
		_go_fuzz_dep_.CoverTab[124288]++
													lines = append(lines, line)
													vals = append(vals, node)
													m[line] = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:268
		// _ = "end of CoverTab[124288]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:269
	// _ = "end of CoverTab[124284]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:269
	_go_fuzz_dep_.CoverTab[124285]++
												sort.Ints(lines)

												for i, line := range lines {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:272
		_go_fuzz_dep_.CoverTab[124292]++
													vals[i] = m[line]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:273
		// _ = "end of CoverTab[124292]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:274
	// _ = "end of CoverTab[124285]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:274
	_go_fuzz_dep_.CoverTab[124286]++

												return vals
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:276
	// _ = "end of CoverTab[124286]"
}

func sortAlphabetical(t *Tree) (vals []sortNode) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:279
	_go_fuzz_dep_.CoverTab[124293]++
												var (
		node		sortNode
		simpVals	[]string
		compVals	[]string
	)
	vals = make([]sortNode, 0)
	m := make(map[string]sortNode)

	for k := range t.values {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:288
		_go_fuzz_dep_.CoverTab[124297]++
													v := t.values[k]
													switch v.(type) {
		case *Tree, []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:291
			_go_fuzz_dep_.CoverTab[124299]++
														node = sortNode{key: k, complexity: valueComplex}
														compVals = append(compVals, node.key)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:293
			// _ = "end of CoverTab[124299]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:294
			_go_fuzz_dep_.CoverTab[124300]++
														node = sortNode{key: k, complexity: valueSimple}
														simpVals = append(simpVals, node.key)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:296
			// _ = "end of CoverTab[124300]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:297
		// _ = "end of CoverTab[124297]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:297
		_go_fuzz_dep_.CoverTab[124298]++
													vals = append(vals, node)
													m[node.key] = node
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:299
		// _ = "end of CoverTab[124298]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:300
	// _ = "end of CoverTab[124293]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:300
	_go_fuzz_dep_.CoverTab[124294]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:303
	sort.Strings(simpVals)
	i := 0
	for _, key := range simpVals {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:305
		_go_fuzz_dep_.CoverTab[124301]++
													vals[i] = m[key]
													i++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:307
		// _ = "end of CoverTab[124301]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:308
	// _ = "end of CoverTab[124294]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:308
	_go_fuzz_dep_.CoverTab[124295]++

												sort.Strings(compVals)
												for _, key := range compVals {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:311
		_go_fuzz_dep_.CoverTab[124302]++
													vals[i] = m[key]
													i++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:313
		// _ = "end of CoverTab[124302]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:314
	// _ = "end of CoverTab[124295]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:314
	_go_fuzz_dep_.CoverTab[124296]++

												return vals
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:316
	// _ = "end of CoverTab[124296]"
}

func (t *Tree) writeTo(w io.Writer, indent, keyspace string, bytesCount int64, arraysOneElementPerLine bool) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:319
	_go_fuzz_dep_.CoverTab[124303]++
												return t.writeToOrdered(w, indent, keyspace, bytesCount, arraysOneElementPerLine, OrderAlphabetical, "  ", false, false)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:320
	// _ = "end of CoverTab[124303]"
}

func (t *Tree) writeToOrdered(w io.Writer, indent, keyspace string, bytesCount int64, arraysOneElementPerLine bool, ord MarshalOrder, indentString string, compactComments, parentCommented bool) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:323
	_go_fuzz_dep_.CoverTab[124304]++
												var orderedVals []sortNode

												switch ord {
	case OrderPreserve:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:327
		_go_fuzz_dep_.CoverTab[124307]++
													orderedVals = sortByLines(t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:328
		// _ = "end of CoverTab[124307]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:329
		_go_fuzz_dep_.CoverTab[124308]++
													orderedVals = sortAlphabetical(t)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:330
		// _ = "end of CoverTab[124308]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:331
	// _ = "end of CoverTab[124304]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:331
	_go_fuzz_dep_.CoverTab[124305]++

												for _, node := range orderedVals {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:333
		_go_fuzz_dep_.CoverTab[124309]++
													switch node.complexity {
		case valueComplex:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:335
			_go_fuzz_dep_.CoverTab[124310]++
														k := node.key
														v := t.values[k]

														combinedKey := quoteKeyIfNeeded(k)
														if keyspace != "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:340
				_go_fuzz_dep_.CoverTab[124317]++
															combinedKey = keyspace + "." + combinedKey
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:341
				// _ = "end of CoverTab[124317]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:342
				_go_fuzz_dep_.CoverTab[124318]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:342
				// _ = "end of CoverTab[124318]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:342
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:342
			// _ = "end of CoverTab[124310]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:342
			_go_fuzz_dep_.CoverTab[124311]++

														switch node := v.(type) {

			case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:346
				_go_fuzz_dep_.CoverTab[124319]++
															tv, ok := t.values[k].(*Tree)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:348
					_go_fuzz_dep_.CoverTab[124325]++
																return bytesCount, fmt.Errorf("invalid value type at %s: %T", k, t.values[k])
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:349
					// _ = "end of CoverTab[124325]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:350
					_go_fuzz_dep_.CoverTab[124326]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:350
					// _ = "end of CoverTab[124326]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:350
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:350
				// _ = "end of CoverTab[124319]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:350
				_go_fuzz_dep_.CoverTab[124320]++
															if tv.comment != "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:351
					_go_fuzz_dep_.CoverTab[124327]++
																comment := strings.Replace(tv.comment, "\n", "\n"+indent+"#", -1)
																start := "# "
																if strings.HasPrefix(comment, "#") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:354
						_go_fuzz_dep_.CoverTab[124329]++
																	start = ""
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:355
						// _ = "end of CoverTab[124329]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:356
						_go_fuzz_dep_.CoverTab[124330]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:356
						// _ = "end of CoverTab[124330]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:356
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:356
					// _ = "end of CoverTab[124327]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:356
					_go_fuzz_dep_.CoverTab[124328]++
																writtenBytesCountComment, errc := writeStrings(w, "\n", indent, start, comment)
																bytesCount += int64(writtenBytesCountComment)
																if errc != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:359
						_go_fuzz_dep_.CoverTab[124331]++
																	return bytesCount, errc
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:360
						// _ = "end of CoverTab[124331]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:361
						_go_fuzz_dep_.CoverTab[124332]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:361
						// _ = "end of CoverTab[124332]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:361
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:361
					// _ = "end of CoverTab[124328]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:362
					_go_fuzz_dep_.CoverTab[124333]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:362
					// _ = "end of CoverTab[124333]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:362
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:362
				// _ = "end of CoverTab[124320]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:362
				_go_fuzz_dep_.CoverTab[124321]++

															var commented string
															if parentCommented || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
					_go_fuzz_dep_.CoverTab[124334]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
					return t.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
					// _ = "end of CoverTab[124334]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
					_go_fuzz_dep_.CoverTab[124335]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
					return tv.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
					// _ = "end of CoverTab[124335]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
				}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:365
					_go_fuzz_dep_.CoverTab[124336]++
																commented = "# "
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:366
					// _ = "end of CoverTab[124336]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:367
					_go_fuzz_dep_.CoverTab[124337]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:367
					// _ = "end of CoverTab[124337]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:367
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:367
				// _ = "end of CoverTab[124321]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:367
				_go_fuzz_dep_.CoverTab[124322]++
															writtenBytesCount, err := writeStrings(w, "\n", indent, commented, "[", combinedKey, "]\n")
															bytesCount += int64(writtenBytesCount)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:370
					_go_fuzz_dep_.CoverTab[124338]++
																return bytesCount, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:371
					// _ = "end of CoverTab[124338]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:372
					_go_fuzz_dep_.CoverTab[124339]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:372
					// _ = "end of CoverTab[124339]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:372
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:372
				// _ = "end of CoverTab[124322]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:372
				_go_fuzz_dep_.CoverTab[124323]++
															bytesCount, err = node.writeToOrdered(w, indent+indentString, combinedKey, bytesCount, arraysOneElementPerLine, ord, indentString, compactComments, parentCommented || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
					_go_fuzz_dep_.CoverTab[124340]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
					return t.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
					// _ = "end of CoverTab[124340]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
				}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
					_go_fuzz_dep_.CoverTab[124341]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
					return tv.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
					// _ = "end of CoverTab[124341]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:373
				}())
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:374
					_go_fuzz_dep_.CoverTab[124342]++
																return bytesCount, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:375
					// _ = "end of CoverTab[124342]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:376
					_go_fuzz_dep_.CoverTab[124343]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:376
					// _ = "end of CoverTab[124343]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:376
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:376
				// _ = "end of CoverTab[124323]"
			case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:377
				_go_fuzz_dep_.CoverTab[124324]++
															for _, subTree := range node {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:378
					_go_fuzz_dep_.CoverTab[124344]++
																var commented string
																if parentCommented || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
						_go_fuzz_dep_.CoverTab[124347]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
						return t.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
						// _ = "end of CoverTab[124347]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
						_go_fuzz_dep_.CoverTab[124348]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
						return subTree.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
						// _ = "end of CoverTab[124348]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
					}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:380
						_go_fuzz_dep_.CoverTab[124349]++
																	commented = "# "
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:381
						// _ = "end of CoverTab[124349]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:382
						_go_fuzz_dep_.CoverTab[124350]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:382
						// _ = "end of CoverTab[124350]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:382
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:382
					// _ = "end of CoverTab[124344]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:382
					_go_fuzz_dep_.CoverTab[124345]++
																writtenBytesCount, err := writeStrings(w, "\n", indent, commented, "[[", combinedKey, "]]\n")
																bytesCount += int64(writtenBytesCount)
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:385
						_go_fuzz_dep_.CoverTab[124351]++
																	return bytesCount, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:386
						// _ = "end of CoverTab[124351]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:387
						_go_fuzz_dep_.CoverTab[124352]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:387
						// _ = "end of CoverTab[124352]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:387
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:387
					// _ = "end of CoverTab[124345]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:387
					_go_fuzz_dep_.CoverTab[124346]++

																bytesCount, err = subTree.writeToOrdered(w, indent+indentString, combinedKey, bytesCount, arraysOneElementPerLine, ord, indentString, compactComments, parentCommented || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
						_go_fuzz_dep_.CoverTab[124353]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
						return t.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
						// _ = "end of CoverTab[124353]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
					}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
						_go_fuzz_dep_.CoverTab[124354]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
						return subTree.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
						// _ = "end of CoverTab[124354]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:389
					}())
																if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:390
						_go_fuzz_dep_.CoverTab[124355]++
																	return bytesCount, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:391
						// _ = "end of CoverTab[124355]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:392
						_go_fuzz_dep_.CoverTab[124356]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:392
						// _ = "end of CoverTab[124356]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:392
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:392
					// _ = "end of CoverTab[124346]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:393
				// _ = "end of CoverTab[124324]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:394
			// _ = "end of CoverTab[124311]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:395
			_go_fuzz_dep_.CoverTab[124312]++
														k := node.key
														v, ok := t.values[k].(*tomlValue)
														if !ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:398
				_go_fuzz_dep_.CoverTab[124357]++
															return bytesCount, fmt.Errorf("invalid value type at %s: %T", k, t.values[k])
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:399
				// _ = "end of CoverTab[124357]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:400
				_go_fuzz_dep_.CoverTab[124358]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:400
				// _ = "end of CoverTab[124358]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:400
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:400
			// _ = "end of CoverTab[124312]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:400
			_go_fuzz_dep_.CoverTab[124313]++

														var commented string
														if parentCommented || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
				_go_fuzz_dep_.CoverTab[124359]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
				return t.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
				// _ = "end of CoverTab[124359]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
			}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
				_go_fuzz_dep_.CoverTab[124360]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
				return v.commented
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
				// _ = "end of CoverTab[124360]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
			}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:403
				_go_fuzz_dep_.CoverTab[124361]++
															commented = "# "
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:404
				// _ = "end of CoverTab[124361]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:405
				_go_fuzz_dep_.CoverTab[124362]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:405
				// _ = "end of CoverTab[124362]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:405
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:405
			// _ = "end of CoverTab[124313]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:405
			_go_fuzz_dep_.CoverTab[124314]++
														repr, err := tomlValueStringRepresentation(v, commented, indent, ord, arraysOneElementPerLine)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:407
				_go_fuzz_dep_.CoverTab[124363]++
															return bytesCount, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:408
				// _ = "end of CoverTab[124363]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:409
				_go_fuzz_dep_.CoverTab[124364]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:409
				// _ = "end of CoverTab[124364]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:409
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:409
			// _ = "end of CoverTab[124314]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:409
			_go_fuzz_dep_.CoverTab[124315]++

														if v.comment != "" {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:411
				_go_fuzz_dep_.CoverTab[124365]++
															comment := strings.Replace(v.comment, "\n", "\n"+indent+"#", -1)
															start := "# "
															if strings.HasPrefix(comment, "#") {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:414
					_go_fuzz_dep_.CoverTab[124368]++
																start = ""
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:415
					// _ = "end of CoverTab[124368]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:416
					_go_fuzz_dep_.CoverTab[124369]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:416
					// _ = "end of CoverTab[124369]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:416
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:416
				// _ = "end of CoverTab[124365]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:416
				_go_fuzz_dep_.CoverTab[124366]++
															if !compactComments {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:417
					_go_fuzz_dep_.CoverTab[124370]++
																writtenBytesCountComment, errc := writeStrings(w, "\n")
																bytesCount += int64(writtenBytesCountComment)
																if errc != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:420
						_go_fuzz_dep_.CoverTab[124371]++
																	return bytesCount, errc
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:421
						// _ = "end of CoverTab[124371]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:422
						_go_fuzz_dep_.CoverTab[124372]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:422
						// _ = "end of CoverTab[124372]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:422
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:422
					// _ = "end of CoverTab[124370]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:423
					_go_fuzz_dep_.CoverTab[124373]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:423
					// _ = "end of CoverTab[124373]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:423
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:423
				// _ = "end of CoverTab[124366]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:423
				_go_fuzz_dep_.CoverTab[124367]++
															writtenBytesCountComment, errc := writeStrings(w, indent, start, comment, "\n")
															bytesCount += int64(writtenBytesCountComment)
															if errc != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:426
					_go_fuzz_dep_.CoverTab[124374]++
																return bytesCount, errc
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:427
					// _ = "end of CoverTab[124374]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:428
					_go_fuzz_dep_.CoverTab[124375]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:428
					// _ = "end of CoverTab[124375]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:428
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:428
				// _ = "end of CoverTab[124367]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:429
				_go_fuzz_dep_.CoverTab[124376]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:429
				// _ = "end of CoverTab[124376]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:429
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:429
			// _ = "end of CoverTab[124315]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:429
			_go_fuzz_dep_.CoverTab[124316]++

														quotedKey := quoteKeyIfNeeded(k)
														writtenBytesCount, err := writeStrings(w, indent, commented, quotedKey, " = ", repr, "\n")
														bytesCount += int64(writtenBytesCount)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:434
				_go_fuzz_dep_.CoverTab[124377]++
															return bytesCount, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:435
				// _ = "end of CoverTab[124377]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:436
				_go_fuzz_dep_.CoverTab[124378]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:436
				// _ = "end of CoverTab[124378]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:436
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:436
			// _ = "end of CoverTab[124316]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:437
		// _ = "end of CoverTab[124309]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:438
	// _ = "end of CoverTab[124305]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:438
	_go_fuzz_dep_.CoverTab[124306]++

												return bytesCount, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:440
	// _ = "end of CoverTab[124306]"
}

// quote a key if it does not fit the bare key format (A-Za-z0-9_-)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:443
// quoted keys use the same rules as strings
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:445
func quoteKeyIfNeeded(k string) string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:445
	_go_fuzz_dep_.CoverTab[124379]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
	if len(k) >= 2 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
		_go_fuzz_dep_.CoverTab[124383]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
		return k[0] == '"'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
		// _ = "end of CoverTab[124383]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
		_go_fuzz_dep_.CoverTab[124384]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
		return k[len(k)-1] == '"'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
		// _ = "end of CoverTab[124384]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:449
		_go_fuzz_dep_.CoverTab[124385]++
													return k
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:450
		// _ = "end of CoverTab[124385]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:451
		_go_fuzz_dep_.CoverTab[124386]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:451
		// _ = "end of CoverTab[124386]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:451
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:451
	// _ = "end of CoverTab[124379]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:451
	_go_fuzz_dep_.CoverTab[124380]++
												isBare := true
												for _, r := range k {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:453
		_go_fuzz_dep_.CoverTab[124387]++
													if !isValidBareChar(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:454
			_go_fuzz_dep_.CoverTab[124388]++
														isBare = false
														break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:456
			// _ = "end of CoverTab[124388]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:457
			_go_fuzz_dep_.CoverTab[124389]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:457
			// _ = "end of CoverTab[124389]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:457
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:457
		// _ = "end of CoverTab[124387]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:458
	// _ = "end of CoverTab[124380]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:458
	_go_fuzz_dep_.CoverTab[124381]++
												if isBare {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:459
		_go_fuzz_dep_.CoverTab[124390]++
													return k
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:460
		// _ = "end of CoverTab[124390]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:461
		_go_fuzz_dep_.CoverTab[124391]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:461
		// _ = "end of CoverTab[124391]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:461
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:461
	// _ = "end of CoverTab[124381]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:461
	_go_fuzz_dep_.CoverTab[124382]++
												return quoteKey(k)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:462
	// _ = "end of CoverTab[124382]"
}

func quoteKey(k string) string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:465
	_go_fuzz_dep_.CoverTab[124392]++
												return "\"" + encodeTomlString(k) + "\""
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:466
	// _ = "end of CoverTab[124392]"
}

func writeStrings(w io.Writer, s ...string) (int, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:469
	_go_fuzz_dep_.CoverTab[124393]++
												var n int
												for i := range s {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:471
		_go_fuzz_dep_.CoverTab[124395]++
													b, err := io.WriteString(w, s[i])
													n += b
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:474
			_go_fuzz_dep_.CoverTab[124396]++
														return n, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:475
			// _ = "end of CoverTab[124396]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:476
			_go_fuzz_dep_.CoverTab[124397]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:476
			// _ = "end of CoverTab[124397]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:476
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:476
		// _ = "end of CoverTab[124395]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:477
	// _ = "end of CoverTab[124393]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:477
	_go_fuzz_dep_.CoverTab[124394]++
												return n, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:478
	// _ = "end of CoverTab[124394]"
}

// WriteTo encode the Tree as Toml and writes it to the writer w.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:481
// Returns the number of bytes written in case of success, or an error if anything happened.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:483
func (t *Tree) WriteTo(w io.Writer) (int64, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:483
	_go_fuzz_dep_.CoverTab[124398]++
												return t.writeTo(w, "", "", 0, false)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:484
	// _ = "end of CoverTab[124398]"
}

// ToTomlString generates a human-readable representation of the current tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:487
// Output spans multiple lines, and is suitable for ingest by a TOML parser.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:487
// If the conversion cannot be performed, ToString returns a non-nil error.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:490
func (t *Tree) ToTomlString() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:490
	_go_fuzz_dep_.CoverTab[124399]++
												b, err := t.Marshal()
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:492
		_go_fuzz_dep_.CoverTab[124401]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:493
		// _ = "end of CoverTab[124401]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:494
		_go_fuzz_dep_.CoverTab[124402]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:494
		// _ = "end of CoverTab[124402]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:494
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:494
	// _ = "end of CoverTab[124399]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:494
	_go_fuzz_dep_.CoverTab[124400]++
												return string(b), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:495
	// _ = "end of CoverTab[124400]"
}

// String generates a human-readable representation of the current tree.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:498
// Alias of ToString. Present to implement the fmt.Stringer interface.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:500
func (t *Tree) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:500
	_go_fuzz_dep_.CoverTab[124403]++
												result, _ := t.ToTomlString()
												return result
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:502
	// _ = "end of CoverTab[124403]"
}

// ToMap recursively generates a representation of the tree using Go built-in structures.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
// The following types are used:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - bool
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - float64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - int64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - string
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - uint64
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - time.Time
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - map[string]interface{} (where interface{} is any of this list)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:505
//   - []interface{} (where interface{} is any of this list)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:516
func (t *Tree) ToMap() map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:516
	_go_fuzz_dep_.CoverTab[124404]++
												result := map[string]interface{}{}

												for k, v := range t.values {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:519
		_go_fuzz_dep_.CoverTab[124406]++
													switch node := v.(type) {
		case []*Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:521
			_go_fuzz_dep_.CoverTab[124407]++
														var array []interface{}
														for _, item := range node {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:523
				_go_fuzz_dep_.CoverTab[124411]++
															array = append(array, item.ToMap())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:524
				// _ = "end of CoverTab[124411]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:525
			// _ = "end of CoverTab[124407]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:525
			_go_fuzz_dep_.CoverTab[124408]++
														result[k] = array
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:526
			// _ = "end of CoverTab[124408]"
		case *Tree:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:527
			_go_fuzz_dep_.CoverTab[124409]++
														result[k] = node.ToMap()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:528
			// _ = "end of CoverTab[124409]"
		case *tomlValue:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:529
			_go_fuzz_dep_.CoverTab[124410]++
														result[k] = tomlValueToGo(node.value)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:530
			// _ = "end of CoverTab[124410]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:531
		// _ = "end of CoverTab[124406]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:532
	// _ = "end of CoverTab[124404]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:532
	_go_fuzz_dep_.CoverTab[124405]++
												return result
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:533
	// _ = "end of CoverTab[124405]"
}

func tomlValueToGo(v interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:536
	_go_fuzz_dep_.CoverTab[124412]++
												if tree, ok := v.(*Tree); ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:537
		_go_fuzz_dep_.CoverTab[124416]++
													return tree.ToMap()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:538
		// _ = "end of CoverTab[124416]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:539
		_go_fuzz_dep_.CoverTab[124417]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:539
		// _ = "end of CoverTab[124417]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:539
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:539
	// _ = "end of CoverTab[124412]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:539
	_go_fuzz_dep_.CoverTab[124413]++

												rv := reflect.ValueOf(v)

												if rv.Kind() != reflect.Slice {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:543
		_go_fuzz_dep_.CoverTab[124418]++
													return v
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:544
		// _ = "end of CoverTab[124418]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:545
		_go_fuzz_dep_.CoverTab[124419]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:545
		// _ = "end of CoverTab[124419]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:545
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:545
	// _ = "end of CoverTab[124413]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:545
	_go_fuzz_dep_.CoverTab[124414]++
												values := make([]interface{}, rv.Len())
												for i := 0; i < rv.Len(); i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:547
		_go_fuzz_dep_.CoverTab[124420]++
													item := rv.Index(i).Interface()
													values[i] = tomlValueToGo(item)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:549
		// _ = "end of CoverTab[124420]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:550
	// _ = "end of CoverTab[124414]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:550
	_go_fuzz_dep_.CoverTab[124415]++
												return values
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:551
	// _ = "end of CoverTab[124415]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:552
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_write.go:552
var _ = _go_fuzz_dep_.CoverTab
