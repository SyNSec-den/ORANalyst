//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:1
)

import (
	"encoding/base64"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type resolveMapItem struct {
	value	interface{}
	tag	string
}

var resolveTable = make([]byte, 256)
var resolveMap = make(map[string]resolveMapItem)

func init() {
	t := resolveTable
	t[int('+')] = 'S'
	t[int('-')] = 'S'
	for _, c := range "0123456789" {
		t[int(c)] = 'D'
	}
	for _, c := range "yYnNtTfFoO~" {
		t[int(c)] = 'M'
	}
	t[int('.')] = '.'

	var resolveMapList = []struct {
		v	interface{}
		tag	string
		l	[]string
	}{
		{true, yaml_BOOL_TAG, []string{"y", "Y", "yes", "Yes", "YES"}},
		{true, yaml_BOOL_TAG, []string{"true", "True", "TRUE"}},
		{true, yaml_BOOL_TAG, []string{"on", "On", "ON"}},
		{false, yaml_BOOL_TAG, []string{"n", "N", "no", "No", "NO"}},
		{false, yaml_BOOL_TAG, []string{"false", "False", "FALSE"}},
		{false, yaml_BOOL_TAG, []string{"off", "Off", "OFF"}},
		{nil, yaml_NULL_TAG, []string{"", "~", "null", "Null", "NULL"}},
		{math.NaN(), yaml_FLOAT_TAG, []string{".nan", ".NaN", ".NAN"}},
		{math.Inf(+1), yaml_FLOAT_TAG, []string{".inf", ".Inf", ".INF"}},
		{math.Inf(+1), yaml_FLOAT_TAG, []string{"+.inf", "+.Inf", "+.INF"}},
		{math.Inf(-1), yaml_FLOAT_TAG, []string{"-.inf", "-.Inf", "-.INF"}},
		{"<<", yaml_MERGE_TAG, []string{"<<"}},
	}

	m := resolveMap
	for _, item := range resolveMapList {
		for _, s := range item.l {
			m[s] = resolveMapItem{item.v, item.tag}
		}
	}
}

const longTagPrefix = "tag:yaml.org,2002:"

func shortTag(tag string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:61
	_go_fuzz_dep_.CoverTab[126635]++

										if strings.HasPrefix(tag, longTagPrefix) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:63
		_go_fuzz_dep_.CoverTab[126637]++
											return "!!" + tag[len(longTagPrefix):]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:64
		// _ = "end of CoverTab[126637]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:65
		_go_fuzz_dep_.CoverTab[126638]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:65
		// _ = "end of CoverTab[126638]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:65
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:65
	// _ = "end of CoverTab[126635]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:65
	_go_fuzz_dep_.CoverTab[126636]++
										return tag
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:66
	// _ = "end of CoverTab[126636]"
}

func longTag(tag string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:69
	_go_fuzz_dep_.CoverTab[126639]++
										if strings.HasPrefix(tag, "!!") {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:70
		_go_fuzz_dep_.CoverTab[126641]++
											return longTagPrefix + tag[2:]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:71
		// _ = "end of CoverTab[126641]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:72
		_go_fuzz_dep_.CoverTab[126642]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:72
		// _ = "end of CoverTab[126642]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:72
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:72
	// _ = "end of CoverTab[126639]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:72
	_go_fuzz_dep_.CoverTab[126640]++
										return tag
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:73
	// _ = "end of CoverTab[126640]"
}

func resolvableTag(tag string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:76
	_go_fuzz_dep_.CoverTab[126643]++
										switch tag {
	case "", yaml_STR_TAG, yaml_BOOL_TAG, yaml_INT_TAG, yaml_FLOAT_TAG, yaml_NULL_TAG, yaml_TIMESTAMP_TAG:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:78
		_go_fuzz_dep_.CoverTab[126645]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:79
		// _ = "end of CoverTab[126645]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:79
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:79
		_go_fuzz_dep_.CoverTab[126646]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:79
		// _ = "end of CoverTab[126646]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:80
	// _ = "end of CoverTab[126643]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:80
	_go_fuzz_dep_.CoverTab[126644]++
										return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:81
	// _ = "end of CoverTab[126644]"
}

var yamlStyleFloat = regexp.MustCompile(`^[-+]?(\.[0-9]+|[0-9]+(\.[0-9]*)?)([eE][-+]?[0-9]+)?$`)

func resolve(tag string, in string) (rtag string, out interface{}) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:86
	_go_fuzz_dep_.CoverTab[126647]++
										if !resolvableTag(tag) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:87
		_go_fuzz_dep_.CoverTab[126652]++
											return tag, in
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:88
		// _ = "end of CoverTab[126652]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:89
		_go_fuzz_dep_.CoverTab[126653]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:89
		// _ = "end of CoverTab[126653]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:89
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:89
	// _ = "end of CoverTab[126647]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:89
	_go_fuzz_dep_.CoverTab[126648]++

										defer func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:91
		_go_fuzz_dep_.CoverTab[126654]++
											switch tag {
		case "", rtag, yaml_STR_TAG, yaml_BINARY_TAG:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:93
			_go_fuzz_dep_.CoverTab[126656]++
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:94
			// _ = "end of CoverTab[126656]"
		case yaml_FLOAT_TAG:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:95
			_go_fuzz_dep_.CoverTab[126657]++
												if rtag == yaml_INT_TAG {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:96
				_go_fuzz_dep_.CoverTab[126659]++
													switch v := out.(type) {
				case int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:98
					_go_fuzz_dep_.CoverTab[126660]++
														rtag = yaml_FLOAT_TAG
														out = float64(v)
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:101
					// _ = "end of CoverTab[126660]"
				case int:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:102
					_go_fuzz_dep_.CoverTab[126661]++
														rtag = yaml_FLOAT_TAG
														out = float64(v)
														return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:105
					// _ = "end of CoverTab[126661]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:106
				// _ = "end of CoverTab[126659]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:107
				_go_fuzz_dep_.CoverTab[126662]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:107
				// _ = "end of CoverTab[126662]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:107
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:107
			// _ = "end of CoverTab[126657]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:107
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:107
			_go_fuzz_dep_.CoverTab[126658]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:107
			// _ = "end of CoverTab[126658]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:108
		// _ = "end of CoverTab[126654]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:108
		_go_fuzz_dep_.CoverTab[126655]++
											failf("cannot decode %s `%s` as a %s", shortTag(rtag), in, shortTag(tag))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:109
		// _ = "end of CoverTab[126655]"
	}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:110
	// _ = "end of CoverTab[126648]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:110
	_go_fuzz_dep_.CoverTab[126649]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:114
	hint := byte('N')
	if in != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:115
		_go_fuzz_dep_.CoverTab[126663]++
											hint = resolveTable[in[0]]
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:116
		// _ = "end of CoverTab[126663]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:117
		_go_fuzz_dep_.CoverTab[126664]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:117
		// _ = "end of CoverTab[126664]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:117
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:117
	// _ = "end of CoverTab[126649]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:117
	_go_fuzz_dep_.CoverTab[126650]++
										if hint != 0 && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
		_go_fuzz_dep_.CoverTab[126665]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
		return tag != yaml_STR_TAG
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
		// _ = "end of CoverTab[126665]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
		_go_fuzz_dep_.CoverTab[126666]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
		return tag != yaml_BINARY_TAG
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
		// _ = "end of CoverTab[126666]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:118
		_go_fuzz_dep_.CoverTab[126667]++

											if item, ok := resolveMap[in]; ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:120
			_go_fuzz_dep_.CoverTab[126669]++
												return item.tag, item.value
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:121
			// _ = "end of CoverTab[126669]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:122
			_go_fuzz_dep_.CoverTab[126670]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:122
			// _ = "end of CoverTab[126670]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:122
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:122
		// _ = "end of CoverTab[126667]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:122
		_go_fuzz_dep_.CoverTab[126668]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:128
		switch hint {
		case 'M':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:129
			_go_fuzz_dep_.CoverTab[126671]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:129
			// _ = "end of CoverTab[126671]"

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:132
		case '.':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:132
			_go_fuzz_dep_.CoverTab[126672]++

												floatv, err := strconv.ParseFloat(in, 64)
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:135
				_go_fuzz_dep_.CoverTab[126679]++
													return yaml_FLOAT_TAG, floatv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:136
				// _ = "end of CoverTab[126679]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:137
				_go_fuzz_dep_.CoverTab[126680]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:137
				// _ = "end of CoverTab[126680]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:137
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:137
			// _ = "end of CoverTab[126672]"

		case 'D', 'S':
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:139
			_go_fuzz_dep_.CoverTab[126673]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:143
			if tag == "" || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:143
				_go_fuzz_dep_.CoverTab[126681]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:143
				return tag == yaml_TIMESTAMP_TAG
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:143
				// _ = "end of CoverTab[126681]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:143
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:143
				_go_fuzz_dep_.CoverTab[126682]++
													t, ok := parseTimestamp(in)
													if ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:145
					_go_fuzz_dep_.CoverTab[126683]++
														return yaml_TIMESTAMP_TAG, t
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:146
					// _ = "end of CoverTab[126683]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:147
					_go_fuzz_dep_.CoverTab[126684]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:147
					// _ = "end of CoverTab[126684]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:147
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:147
				// _ = "end of CoverTab[126682]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:148
				_go_fuzz_dep_.CoverTab[126685]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:148
				// _ = "end of CoverTab[126685]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:148
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:148
			// _ = "end of CoverTab[126673]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:148
			_go_fuzz_dep_.CoverTab[126674]++

												plain := strings.Replace(in, "_", "", -1)
												intv, err := strconv.ParseInt(plain, 0, 64)
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:152
				_go_fuzz_dep_.CoverTab[126686]++
													if intv == int64(int(intv)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:153
					_go_fuzz_dep_.CoverTab[126687]++
														return yaml_INT_TAG, int(intv)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:154
					// _ = "end of CoverTab[126687]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:155
					_go_fuzz_dep_.CoverTab[126688]++
														return yaml_INT_TAG, intv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:156
					// _ = "end of CoverTab[126688]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:157
				// _ = "end of CoverTab[126686]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:158
				_go_fuzz_dep_.CoverTab[126689]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:158
				// _ = "end of CoverTab[126689]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:158
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:158
			// _ = "end of CoverTab[126674]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:158
			_go_fuzz_dep_.CoverTab[126675]++
												uintv, err := strconv.ParseUint(plain, 0, 64)
												if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:160
				_go_fuzz_dep_.CoverTab[126690]++
													return yaml_INT_TAG, uintv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:161
				// _ = "end of CoverTab[126690]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:162
				_go_fuzz_dep_.CoverTab[126691]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:162
				// _ = "end of CoverTab[126691]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:162
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:162
			// _ = "end of CoverTab[126675]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:162
			_go_fuzz_dep_.CoverTab[126676]++
												if yamlStyleFloat.MatchString(plain) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:163
				_go_fuzz_dep_.CoverTab[126692]++
													floatv, err := strconv.ParseFloat(plain, 64)
													if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:165
					_go_fuzz_dep_.CoverTab[126693]++
														return yaml_FLOAT_TAG, floatv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:166
					// _ = "end of CoverTab[126693]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:167
					_go_fuzz_dep_.CoverTab[126694]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:167
					// _ = "end of CoverTab[126694]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:167
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:167
				// _ = "end of CoverTab[126692]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:168
				_go_fuzz_dep_.CoverTab[126695]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:168
				// _ = "end of CoverTab[126695]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:168
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:168
			// _ = "end of CoverTab[126676]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:168
			_go_fuzz_dep_.CoverTab[126677]++
												if strings.HasPrefix(plain, "0b") {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:169
				_go_fuzz_dep_.CoverTab[126696]++
													intv, err := strconv.ParseInt(plain[2:], 2, 64)
													if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:171
					_go_fuzz_dep_.CoverTab[126698]++
														if intv == int64(int(intv)) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:172
						_go_fuzz_dep_.CoverTab[126699]++
															return yaml_INT_TAG, int(intv)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:173
						// _ = "end of CoverTab[126699]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:174
						_go_fuzz_dep_.CoverTab[126700]++
															return yaml_INT_TAG, intv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:175
						// _ = "end of CoverTab[126700]"
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:176
					// _ = "end of CoverTab[126698]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:177
					_go_fuzz_dep_.CoverTab[126701]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:177
					// _ = "end of CoverTab[126701]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:177
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:177
				// _ = "end of CoverTab[126696]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:177
				_go_fuzz_dep_.CoverTab[126697]++
													uintv, err := strconv.ParseUint(plain[2:], 2, 64)
													if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:179
					_go_fuzz_dep_.CoverTab[126702]++
														return yaml_INT_TAG, uintv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:180
					// _ = "end of CoverTab[126702]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:181
					_go_fuzz_dep_.CoverTab[126703]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:181
					// _ = "end of CoverTab[126703]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:181
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:181
				// _ = "end of CoverTab[126697]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:182
				_go_fuzz_dep_.CoverTab[126704]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:182
				if strings.HasPrefix(plain, "-0b") {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:182
					_go_fuzz_dep_.CoverTab[126705]++
														intv, err := strconv.ParseInt("-"+plain[3:], 2, 64)
														if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:184
						_go_fuzz_dep_.CoverTab[126706]++
															if true || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:185
							_go_fuzz_dep_.CoverTab[126707]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:185
							return intv == int64(int(intv))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:185
							// _ = "end of CoverTab[126707]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:185
						}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:185
							_go_fuzz_dep_.CoverTab[126708]++
																return yaml_INT_TAG, int(intv)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:186
							// _ = "end of CoverTab[126708]"
						} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:187
							_go_fuzz_dep_.CoverTab[126709]++
																return yaml_INT_TAG, intv
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:188
							// _ = "end of CoverTab[126709]"
						}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:189
						// _ = "end of CoverTab[126706]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:190
						_go_fuzz_dep_.CoverTab[126710]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:190
						// _ = "end of CoverTab[126710]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:190
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:190
					// _ = "end of CoverTab[126705]"
				} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:191
					_go_fuzz_dep_.CoverTab[126711]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:191
					// _ = "end of CoverTab[126711]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:191
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:191
				// _ = "end of CoverTab[126704]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:191
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:191
			// _ = "end of CoverTab[126677]"
		default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:192
			_go_fuzz_dep_.CoverTab[126678]++
												panic("resolveTable item not yet handled: " + string(rune(hint)) + " (with " + in + ")")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:193
			// _ = "end of CoverTab[126678]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:194
		// _ = "end of CoverTab[126668]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:195
		_go_fuzz_dep_.CoverTab[126712]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:195
		// _ = "end of CoverTab[126712]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:195
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:195
	// _ = "end of CoverTab[126650]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:195
	_go_fuzz_dep_.CoverTab[126651]++
										return yaml_STR_TAG, in
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:196
	// _ = "end of CoverTab[126651]"
}

// encodeBase64 encodes s as base64 that is broken up into multiple lines
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:199
// as appropriate for the resulting length.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:201
func encodeBase64(s string) string {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:201
	_go_fuzz_dep_.CoverTab[126713]++
										const lineLen = 70
										encLen := base64.StdEncoding.EncodedLen(len(s))
										lines := encLen/lineLen + 1
										buf := make([]byte, encLen*2+lines)
										in := buf[0:encLen]
										out := buf[encLen:]
										base64.StdEncoding.Encode(in, []byte(s))
										k := 0
										for i := 0; i < len(in); i += lineLen {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:210
		_go_fuzz_dep_.CoverTab[126715]++
											j := i + lineLen
											if j > len(in) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:212
			_go_fuzz_dep_.CoverTab[126717]++
												j = len(in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:213
			// _ = "end of CoverTab[126717]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:214
			_go_fuzz_dep_.CoverTab[126718]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:214
			// _ = "end of CoverTab[126718]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:214
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:214
		// _ = "end of CoverTab[126715]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:214
		_go_fuzz_dep_.CoverTab[126716]++
											k += copy(out[k:], in[i:j])
											if lines > 1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:216
			_go_fuzz_dep_.CoverTab[126719]++
												out[k] = '\n'
												k++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:218
			// _ = "end of CoverTab[126719]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:219
			_go_fuzz_dep_.CoverTab[126720]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:219
			// _ = "end of CoverTab[126720]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:219
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:219
		// _ = "end of CoverTab[126716]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:220
	// _ = "end of CoverTab[126713]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:220
	_go_fuzz_dep_.CoverTab[126714]++
										return string(out[:k])
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:221
	// _ = "end of CoverTab[126714]"
}

// This is a subset of the formats allowed by the regular expression
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:224
// defined at http://yaml.org/type/timestamp.html.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:226
var allowedTimestampFormats = []string{
	"2006-1-2T15:4:5.999999999Z07:00",
	"2006-1-2t15:4:5.999999999Z07:00",
	"2006-1-2 15:4:5.999999999",
	"2006-1-2",
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:233
}

// parseTimestamp parses s as a timestamp string and
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:235
// returns the timestamp and reports whether it succeeded.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:235
// Timestamp formats are defined at http://yaml.org/type/timestamp.html
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:238
func parseTimestamp(s string) (time.Time, bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:238
	_go_fuzz_dep_.CoverTab[126721]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:243
	i := 0
	for ; i < len(s); i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:244
		_go_fuzz_dep_.CoverTab[126725]++
											if c := s[i]; c < '0' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:245
			_go_fuzz_dep_.CoverTab[126726]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:245
			return c > '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:245
			// _ = "end of CoverTab[126726]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:245
		}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:245
			_go_fuzz_dep_.CoverTab[126727]++
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:246
			// _ = "end of CoverTab[126727]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:247
			_go_fuzz_dep_.CoverTab[126728]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:247
			// _ = "end of CoverTab[126728]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:247
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:247
		// _ = "end of CoverTab[126725]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:248
	// _ = "end of CoverTab[126721]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:248
	_go_fuzz_dep_.CoverTab[126722]++
										if i != 4 || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
		_go_fuzz_dep_.CoverTab[126729]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
		return i == len(s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
		// _ = "end of CoverTab[126729]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
		_go_fuzz_dep_.CoverTab[126730]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
		return s[i] != '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
		// _ = "end of CoverTab[126730]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:249
		_go_fuzz_dep_.CoverTab[126731]++
											return time.Time{}, false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:250
		// _ = "end of CoverTab[126731]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:251
		_go_fuzz_dep_.CoverTab[126732]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:251
		// _ = "end of CoverTab[126732]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:251
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:251
	// _ = "end of CoverTab[126722]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:251
	_go_fuzz_dep_.CoverTab[126723]++
										for _, format := range allowedTimestampFormats {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:252
		_go_fuzz_dep_.CoverTab[126733]++
											if t, err := time.Parse(format, s); err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:253
			_go_fuzz_dep_.CoverTab[126734]++
												return t, true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:254
			// _ = "end of CoverTab[126734]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:255
			_go_fuzz_dep_.CoverTab[126735]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:255
			// _ = "end of CoverTab[126735]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:255
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:255
		// _ = "end of CoverTab[126733]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:256
	// _ = "end of CoverTab[126723]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:256
	_go_fuzz_dep_.CoverTab[126724]++
										return time.Time{}, false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:257
	// _ = "end of CoverTab[126724]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:258
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/resolve.go:258
var _ = _go_fuzz_dep_.CoverTab
