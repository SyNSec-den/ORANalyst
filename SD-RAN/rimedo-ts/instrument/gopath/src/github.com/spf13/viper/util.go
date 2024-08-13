// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Viper is a application configuration system.
// It believes that applications can be configured a variety of ways
// via flags, ENVIRONMENT variables, configuration files retrieved
// from the file system, or a remote key/value store.

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:11
package viper

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:11
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:11
)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:11
import (
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:11
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:11
)

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"github.com/spf13/afero"
	"github.com/spf13/cast"
	jww "github.com/spf13/jwalterweatherman"
)

// ConfigParseError denotes failing to parse configuration file.
type ConfigParseError struct {
	err error
}

// Error returns the formatted configuration error.
func (pe ConfigParseError) Error() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:32
	_go_fuzz_dep_.CoverTab[129554]++
										return fmt.Sprintf("While parsing config: %s", pe.err.Error())
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:33
	// _ = "end of CoverTab[129554]"
}

// toCaseInsensitiveValue checks if the value is a  map;
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:36
// if so, create a copy and lower-case the keys recursively.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:38
func toCaseInsensitiveValue(value interface{}) interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:38
	_go_fuzz_dep_.CoverTab[129555]++
										switch v := value.(type) {
	case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:40
		_go_fuzz_dep_.CoverTab[129557]++
											value = copyAndInsensitiviseMap(cast.ToStringMap(v))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:41
		// _ = "end of CoverTab[129557]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:42
		_go_fuzz_dep_.CoverTab[129558]++
											value = copyAndInsensitiviseMap(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:43
		// _ = "end of CoverTab[129558]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:44
	// _ = "end of CoverTab[129555]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:44
	_go_fuzz_dep_.CoverTab[129556]++

										return value
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:46
	// _ = "end of CoverTab[129556]"
}

// copyAndInsensitiviseMap behaves like insensitiviseMap, but creates a copy of
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:49
// any map it makes case insensitive.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:51
func copyAndInsensitiviseMap(m map[string]interface{}) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:51
	_go_fuzz_dep_.CoverTab[129559]++
										nm := make(map[string]interface{})

										for key, val := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:54
		_go_fuzz_dep_.CoverTab[129561]++
											lkey := strings.ToLower(key)
											switch v := val.(type) {
		case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:57
			_go_fuzz_dep_.CoverTab[129562]++
												nm[lkey] = copyAndInsensitiviseMap(cast.ToStringMap(v))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:58
			// _ = "end of CoverTab[129562]"
		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:59
			_go_fuzz_dep_.CoverTab[129563]++
												nm[lkey] = copyAndInsensitiviseMap(v)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:60
			// _ = "end of CoverTab[129563]"
		default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:61
			_go_fuzz_dep_.CoverTab[129564]++
												nm[lkey] = v
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:62
			// _ = "end of CoverTab[129564]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:63
		// _ = "end of CoverTab[129561]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:64
	// _ = "end of CoverTab[129559]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:64
	_go_fuzz_dep_.CoverTab[129560]++

										return nm
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:66
	// _ = "end of CoverTab[129560]"
}

func insensitiviseMap(m map[string]interface{}) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:69
	_go_fuzz_dep_.CoverTab[129565]++
										for key, val := range m {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:70
		_go_fuzz_dep_.CoverTab[129566]++
											switch val.(type) {
		case map[interface{}]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:72
			_go_fuzz_dep_.CoverTab[129569]++

												val = cast.ToStringMap(val)
												insensitiviseMap(val.(map[string]interface{}))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:75
			// _ = "end of CoverTab[129569]"
		case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:76
			_go_fuzz_dep_.CoverTab[129570]++

												insensitiviseMap(val.(map[string]interface{}))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:78
			// _ = "end of CoverTab[129570]"
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:79
		// _ = "end of CoverTab[129566]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:79
		_go_fuzz_dep_.CoverTab[129567]++

											lower := strings.ToLower(key)
											if key != lower {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:82
			_go_fuzz_dep_.CoverTab[129571]++

												delete(m, key)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:84
			// _ = "end of CoverTab[129571]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:85
			_go_fuzz_dep_.CoverTab[129572]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:85
			// _ = "end of CoverTab[129572]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:85
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:85
		// _ = "end of CoverTab[129567]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:85
		_go_fuzz_dep_.CoverTab[129568]++

											m[lower] = val
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:87
		// _ = "end of CoverTab[129568]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:88
	// _ = "end of CoverTab[129565]"
}

func absPathify(inPath string) string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:91
	_go_fuzz_dep_.CoverTab[129573]++
										jww.INFO.Println("Trying to resolve absolute path to", inPath)

										if inPath == "$HOME" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:94
		_go_fuzz_dep_.CoverTab[129577]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:94
		return strings.HasPrefix(inPath, "$HOME"+string(os.PathSeparator))
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:94
		// _ = "end of CoverTab[129577]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:94
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:94
		_go_fuzz_dep_.CoverTab[129578]++
											inPath = userHomeDir() + inPath[5:]
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:95
		// _ = "end of CoverTab[129578]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:96
		_go_fuzz_dep_.CoverTab[129579]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:96
		// _ = "end of CoverTab[129579]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:96
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:96
	// _ = "end of CoverTab[129573]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:96
	_go_fuzz_dep_.CoverTab[129574]++

										inPath = os.ExpandEnv(inPath)

										if filepath.IsAbs(inPath) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:100
		_go_fuzz_dep_.CoverTab[129580]++
											return filepath.Clean(inPath)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:101
		// _ = "end of CoverTab[129580]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:102
		_go_fuzz_dep_.CoverTab[129581]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:102
		// _ = "end of CoverTab[129581]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:102
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:102
	// _ = "end of CoverTab[129574]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:102
	_go_fuzz_dep_.CoverTab[129575]++

										p, err := filepath.Abs(inPath)
										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:105
		_go_fuzz_dep_.CoverTab[129582]++
											return filepath.Clean(p)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:106
		// _ = "end of CoverTab[129582]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:107
		_go_fuzz_dep_.CoverTab[129583]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:107
		// _ = "end of CoverTab[129583]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:107
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:107
	// _ = "end of CoverTab[129575]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:107
	_go_fuzz_dep_.CoverTab[129576]++

										jww.ERROR.Println("Couldn't discover absolute path")
										jww.ERROR.Println(err)
										return ""
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:111
	// _ = "end of CoverTab[129576]"
}

// Check if file Exists
func exists(fs afero.Fs, path string) (bool, error) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:115
	_go_fuzz_dep_.CoverTab[129584]++
										stat, err := fs.Stat(path)
										if err == nil {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:117
		_go_fuzz_dep_.CoverTab[129587]++
											return !stat.IsDir(), nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:118
		// _ = "end of CoverTab[129587]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:119
		_go_fuzz_dep_.CoverTab[129588]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:119
		// _ = "end of CoverTab[129588]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:119
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:119
	// _ = "end of CoverTab[129584]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:119
	_go_fuzz_dep_.CoverTab[129585]++
										if os.IsNotExist(err) {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:120
		_go_fuzz_dep_.CoverTab[129589]++
											return false, nil
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:121
		// _ = "end of CoverTab[129589]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:122
		_go_fuzz_dep_.CoverTab[129590]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:122
		// _ = "end of CoverTab[129590]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:122
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:122
	// _ = "end of CoverTab[129585]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:122
	_go_fuzz_dep_.CoverTab[129586]++
										return false, err
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:123
	// _ = "end of CoverTab[129586]"
}

func stringInSlice(a string, list []string) bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:126
	_go_fuzz_dep_.CoverTab[129591]++
										for _, b := range list {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:127
		_go_fuzz_dep_.CoverTab[129593]++
											if b == a {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:128
			_go_fuzz_dep_.CoverTab[129594]++
												return true
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:129
			// _ = "end of CoverTab[129594]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:130
			_go_fuzz_dep_.CoverTab[129595]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:130
			// _ = "end of CoverTab[129595]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:130
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:130
		// _ = "end of CoverTab[129593]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:131
	// _ = "end of CoverTab[129591]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:131
	_go_fuzz_dep_.CoverTab[129592]++
										return false
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:132
	// _ = "end of CoverTab[129592]"
}

func userHomeDir() string {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:135
	_go_fuzz_dep_.CoverTab[129596]++
										if runtime.GOOS == "windows" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:136
		_go_fuzz_dep_.CoverTab[129598]++
											home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
											if home == "" {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:138
			_go_fuzz_dep_.CoverTab[129600]++
												home = os.Getenv("USERPROFILE")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:139
			// _ = "end of CoverTab[129600]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:140
			_go_fuzz_dep_.CoverTab[129601]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:140
			// _ = "end of CoverTab[129601]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:140
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:140
		// _ = "end of CoverTab[129598]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:140
		_go_fuzz_dep_.CoverTab[129599]++
											return home
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:141
		// _ = "end of CoverTab[129599]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:142
		_go_fuzz_dep_.CoverTab[129602]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:142
		// _ = "end of CoverTab[129602]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:142
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:142
	// _ = "end of CoverTab[129596]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:142
	_go_fuzz_dep_.CoverTab[129597]++
										return os.Getenv("HOME")
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:143
	// _ = "end of CoverTab[129597]"
}

func safeMul(a, b uint) uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:146
	_go_fuzz_dep_.CoverTab[129603]++
										c := a * b
										if a > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
		_go_fuzz_dep_.CoverTab[129605]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
		return b > 1
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
		// _ = "end of CoverTab[129605]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
		_go_fuzz_dep_.CoverTab[129606]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
		return c/b != a
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
		// _ = "end of CoverTab[129606]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
	}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:148
		_go_fuzz_dep_.CoverTab[129607]++
											return 0
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:149
		// _ = "end of CoverTab[129607]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:150
		_go_fuzz_dep_.CoverTab[129608]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:150
		// _ = "end of CoverTab[129608]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:150
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:150
	// _ = "end of CoverTab[129603]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:150
	_go_fuzz_dep_.CoverTab[129604]++
										return c
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:151
	// _ = "end of CoverTab[129604]"
}

// parseSizeInBytes converts strings like 1GB or 12 mb into an unsigned integer number of bytes
func parseSizeInBytes(sizeStr string) uint {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:155
	_go_fuzz_dep_.CoverTab[129609]++
										sizeStr = strings.TrimSpace(sizeStr)
										lastChar := len(sizeStr) - 1
										multiplier := uint(1)

										if lastChar > 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:160
		_go_fuzz_dep_.CoverTab[129612]++
											if sizeStr[lastChar] == 'b' || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:161
			_go_fuzz_dep_.CoverTab[129613]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:161
			return sizeStr[lastChar] == 'B'
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:161
			// _ = "end of CoverTab[129613]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:161
		}() {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:161
			_go_fuzz_dep_.CoverTab[129614]++
												if lastChar > 1 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:162
				_go_fuzz_dep_.CoverTab[129615]++
													switch unicode.ToLower(rune(sizeStr[lastChar-1])) {
				case 'k':
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:164
					_go_fuzz_dep_.CoverTab[129616]++
														multiplier = 1 << 10
														sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:166
					// _ = "end of CoverTab[129616]"
				case 'm':
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:167
					_go_fuzz_dep_.CoverTab[129617]++
														multiplier = 1 << 20
														sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:169
					// _ = "end of CoverTab[129617]"
				case 'g':
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:170
					_go_fuzz_dep_.CoverTab[129618]++
														multiplier = 1 << 30
														sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:172
					// _ = "end of CoverTab[129618]"
				default:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:173
					_go_fuzz_dep_.CoverTab[129619]++
														multiplier = 1
														sizeStr = strings.TrimSpace(sizeStr[:lastChar])
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:175
					// _ = "end of CoverTab[129619]"
				}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:176
				// _ = "end of CoverTab[129615]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:177
				_go_fuzz_dep_.CoverTab[129620]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:177
				// _ = "end of CoverTab[129620]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:177
			}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:177
			// _ = "end of CoverTab[129614]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:178
			_go_fuzz_dep_.CoverTab[129621]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:178
			// _ = "end of CoverTab[129621]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:178
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:178
		// _ = "end of CoverTab[129612]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:179
		_go_fuzz_dep_.CoverTab[129622]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:179
		// _ = "end of CoverTab[129622]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:179
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:179
	// _ = "end of CoverTab[129609]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:179
	_go_fuzz_dep_.CoverTab[129610]++

										size := cast.ToInt(sizeStr)
										if size < 0 {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:182
		_go_fuzz_dep_.CoverTab[129623]++
											size = 0
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:183
		// _ = "end of CoverTab[129623]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:184
		_go_fuzz_dep_.CoverTab[129624]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:184
		// _ = "end of CoverTab[129624]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:184
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:184
	// _ = "end of CoverTab[129610]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:184
	_go_fuzz_dep_.CoverTab[129611]++

										return safeMul(uint(size), multiplier)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:186
	// _ = "end of CoverTab[129611]"
}

// deepSearch scans deep maps, following the key indexes listed in the
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:189
// sequence "path".
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:189
// The last value is expected to be another map, and is returned.
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:189
//
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:189
// In case intermediate keys do not exist, or map to a non-map value,
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:189
// a new map is created and inserted, and the search continues from there:
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:189
// the initial map "m" may be modified!
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:196
func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:196
	_go_fuzz_dep_.CoverTab[129625]++
										for _, k := range path {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:197
		_go_fuzz_dep_.CoverTab[129627]++
											m2, ok := m[k]
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:199
			_go_fuzz_dep_.CoverTab[129630]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:202
			m3 := make(map[string]interface{})
												m[k] = m3
												m = m3
												continue
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:205
			// _ = "end of CoverTab[129630]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:206
			_go_fuzz_dep_.CoverTab[129631]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:206
			// _ = "end of CoverTab[129631]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:206
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:206
		// _ = "end of CoverTab[129627]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:206
		_go_fuzz_dep_.CoverTab[129628]++
											m3, ok := m2.(map[string]interface{})
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:208
			_go_fuzz_dep_.CoverTab[129632]++

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:211
			m3 = make(map[string]interface{})
												m[k] = m3
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:212
			// _ = "end of CoverTab[129632]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:213
			_go_fuzz_dep_.CoverTab[129633]++
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:213
			// _ = "end of CoverTab[129633]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:213
		}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:213
		// _ = "end of CoverTab[129628]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:213
		_go_fuzz_dep_.CoverTab[129629]++

											m = m3
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:215
		// _ = "end of CoverTab[129629]"
	}
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:216
	// _ = "end of CoverTab[129625]"
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:216
	_go_fuzz_dep_.CoverTab[129626]++
										return m
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:217
	// _ = "end of CoverTab[129626]"
}

//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:218
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/spf13/viper@v1.9.0/util.go:218
var _ = _go_fuzz_dep_.CoverTab
