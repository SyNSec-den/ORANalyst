//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:1
package homedir

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:1
)

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// DisableCache will disable caching of the home directory. Caching is enabled
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:15
// by default.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:17
var DisableCache bool

var homedirCache string
var cacheLock sync.RWMutex

// Dir returns the home directory for the executing user.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:22
//
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:22
// This uses an OS-specific method for discovering the home directory.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:22
// An error is returned if a home directory cannot be detected.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:26
func Dir() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:26
	_go_fuzz_dep_.CoverTab[107507]++
											if !DisableCache {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:27
		_go_fuzz_dep_.CoverTab[107511]++
												cacheLock.RLock()
												cached := homedirCache
												cacheLock.RUnlock()
												if cached != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:31
			_go_fuzz_dep_.CoverTab[107512]++
													return cached, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:32
			// _ = "end of CoverTab[107512]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:33
			_go_fuzz_dep_.CoverTab[107513]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:33
			// _ = "end of CoverTab[107513]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:33
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:33
		// _ = "end of CoverTab[107511]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:34
		_go_fuzz_dep_.CoverTab[107514]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:34
		// _ = "end of CoverTab[107514]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:34
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:34
	// _ = "end of CoverTab[107507]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:34
	_go_fuzz_dep_.CoverTab[107508]++

											cacheLock.Lock()
											defer cacheLock.Unlock()

											var result string
											var err error
											if runtime.GOOS == "windows" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:41
		_go_fuzz_dep_.CoverTab[107515]++
												result, err = dirWindows()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:42
		// _ = "end of CoverTab[107515]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:43
		_go_fuzz_dep_.CoverTab[107516]++

												result, err = dirUnix()
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:45
		// _ = "end of CoverTab[107516]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:46
	// _ = "end of CoverTab[107508]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:46
	_go_fuzz_dep_.CoverTab[107509]++

											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:48
		_go_fuzz_dep_.CoverTab[107517]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:49
		// _ = "end of CoverTab[107517]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:50
		_go_fuzz_dep_.CoverTab[107518]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:50
		// _ = "end of CoverTab[107518]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:50
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:50
	// _ = "end of CoverTab[107509]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:50
	_go_fuzz_dep_.CoverTab[107510]++
											homedirCache = result
											return result, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:52
	// _ = "end of CoverTab[107510]"
}

// Expand expands the path to include the home directory if the path
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:55
// is prefixed with `~`. If it isn't prefixed with `~`, the path is
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:55
// returned as-is.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:58
func Expand(path string) (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:58
	_go_fuzz_dep_.CoverTab[107519]++
											if len(path) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:59
		_go_fuzz_dep_.CoverTab[107524]++
												return path, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:60
		// _ = "end of CoverTab[107524]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:61
		_go_fuzz_dep_.CoverTab[107525]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:61
		// _ = "end of CoverTab[107525]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:61
	// _ = "end of CoverTab[107519]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:61
	_go_fuzz_dep_.CoverTab[107520]++

											if path[0] != '~' {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:63
		_go_fuzz_dep_.CoverTab[107526]++
												return path, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:64
		// _ = "end of CoverTab[107526]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:65
		_go_fuzz_dep_.CoverTab[107527]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:65
		// _ = "end of CoverTab[107527]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:65
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:65
	// _ = "end of CoverTab[107520]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:65
	_go_fuzz_dep_.CoverTab[107521]++

											if len(path) > 1 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
		_go_fuzz_dep_.CoverTab[107528]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
		return path[1] != '/'
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
		// _ = "end of CoverTab[107528]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
	}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
		_go_fuzz_dep_.CoverTab[107529]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
		return path[1] != '\\'
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
		// _ = "end of CoverTab[107529]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:67
		_go_fuzz_dep_.CoverTab[107530]++
												return "", errors.New("cannot expand user-specific home dir")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:68
		// _ = "end of CoverTab[107530]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:69
		_go_fuzz_dep_.CoverTab[107531]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:69
		// _ = "end of CoverTab[107531]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:69
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:69
	// _ = "end of CoverTab[107521]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:69
	_go_fuzz_dep_.CoverTab[107522]++

											dir, err := Dir()
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:72
		_go_fuzz_dep_.CoverTab[107532]++
												return "", err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:73
		// _ = "end of CoverTab[107532]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:74
		_go_fuzz_dep_.CoverTab[107533]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:74
		// _ = "end of CoverTab[107533]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:74
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:74
	// _ = "end of CoverTab[107522]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:74
	_go_fuzz_dep_.CoverTab[107523]++

											return filepath.Join(dir, path[1:]), nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:76
	// _ = "end of CoverTab[107523]"
}

// Reset clears the cache, forcing the next call to Dir to re-detect
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:79
// the home directory. This generally never has to be called, but can be
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:79
// useful in tests if you're modifying the home directory via the HOME
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:79
// env var or something.
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:83
func Reset() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:83
	_go_fuzz_dep_.CoverTab[107534]++
											cacheLock.Lock()
											defer cacheLock.Unlock()
											homedirCache = ""
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:86
	// _ = "end of CoverTab[107534]"
}

func dirUnix() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:89
	_go_fuzz_dep_.CoverTab[107535]++
											homeEnv := "HOME"
											if runtime.GOOS == "plan9" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:91
		_go_fuzz_dep_.CoverTab[107541]++

												homeEnv = "home"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:93
		// _ = "end of CoverTab[107541]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:94
		_go_fuzz_dep_.CoverTab[107542]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:94
		// _ = "end of CoverTab[107542]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:94
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:94
	// _ = "end of CoverTab[107535]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:94
	_go_fuzz_dep_.CoverTab[107536]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:97
	if home := os.Getenv(homeEnv); home != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:97
		_go_fuzz_dep_.CoverTab[107543]++
												return home, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:98
		// _ = "end of CoverTab[107543]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:99
		_go_fuzz_dep_.CoverTab[107544]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:99
		// _ = "end of CoverTab[107544]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:99
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:99
		// _ = "end of CoverTab[107536]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:99
		_go_fuzz_dep_.CoverTab[107537]++

												var stdout bytes.Buffer

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:104
	if runtime.GOOS == "darwin" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:104
		_go_fuzz_dep_.CoverTab[107545]++
													cmd := exec.Command("sh", "-c", `dscl -q . -read /Users/"$(whoami)" NFSHomeDirectory | sed 's/^[^ ]*: //'`)
													cmd.Stdout = &stdout
													if err := cmd.Run(); err == nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:107
			_go_fuzz_dep_.CoverTab[107546]++
														result := strings.TrimSpace(stdout.String())
														if result != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:109
				_go_fuzz_dep_.CoverTab[107547]++
															return result, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:110
				// _ = "end of CoverTab[107547]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:111
				_go_fuzz_dep_.CoverTab[107548]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:111
				// _ = "end of CoverTab[107548]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:111
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:111
			// _ = "end of CoverTab[107546]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:112
			_go_fuzz_dep_.CoverTab[107549]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:112
			// _ = "end of CoverTab[107549]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:112
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:112
		// _ = "end of CoverTab[107545]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:113
		_go_fuzz_dep_.CoverTab[107550]++
													cmd := exec.Command("getent", "passwd", strconv.Itoa(os.Getuid()))
													cmd.Stdout = &stdout
													if err := cmd.Run(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:116
			_go_fuzz_dep_.CoverTab[107551]++

														if err != exec.ErrNotFound {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:118
				_go_fuzz_dep_.CoverTab[107552]++
															return "", err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:119
				// _ = "end of CoverTab[107552]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:120
				_go_fuzz_dep_.CoverTab[107553]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:120
				// _ = "end of CoverTab[107553]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:120
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:120
			// _ = "end of CoverTab[107551]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:121
			_go_fuzz_dep_.CoverTab[107554]++
														if passwd := strings.TrimSpace(stdout.String()); passwd != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:122
				_go_fuzz_dep_.CoverTab[107555]++

															passwdParts := strings.SplitN(passwd, ":", 7)
															if len(passwdParts) > 5 {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:125
					_go_fuzz_dep_.CoverTab[107556]++
																return passwdParts[5], nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:126
					// _ = "end of CoverTab[107556]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:127
					_go_fuzz_dep_.CoverTab[107557]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:127
					// _ = "end of CoverTab[107557]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:127
				}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:127
				// _ = "end of CoverTab[107555]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:128
				_go_fuzz_dep_.CoverTab[107558]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:128
				// _ = "end of CoverTab[107558]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:128
			}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:128
			// _ = "end of CoverTab[107554]"
		}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:129
		// _ = "end of CoverTab[107550]"
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:130
	// _ = "end of CoverTab[107537]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:130
	_go_fuzz_dep_.CoverTab[107538]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:133
	stdout.Reset()
	cmd := exec.Command("sh", "-c", "cd && pwd")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:136
		_go_fuzz_dep_.CoverTab[107559]++
													return "", err
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:137
		// _ = "end of CoverTab[107559]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:138
		_go_fuzz_dep_.CoverTab[107560]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:138
		// _ = "end of CoverTab[107560]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:138
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:138
	// _ = "end of CoverTab[107538]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:138
	_go_fuzz_dep_.CoverTab[107539]++

												result := strings.TrimSpace(stdout.String())
												if result == "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:141
		_go_fuzz_dep_.CoverTab[107561]++
													return "", errors.New("blank output when reading home directory")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:142
		// _ = "end of CoverTab[107561]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:143
		_go_fuzz_dep_.CoverTab[107562]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:143
		// _ = "end of CoverTab[107562]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:143
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:143
	// _ = "end of CoverTab[107539]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:143
	_go_fuzz_dep_.CoverTab[107540]++

												return result, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:145
	// _ = "end of CoverTab[107540]"
}

func dirWindows() (string, error) {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:148
	_go_fuzz_dep_.CoverTab[107563]++

												if home := os.Getenv("HOME"); home != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:150
		_go_fuzz_dep_.CoverTab[107567]++
													return home, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:151
		// _ = "end of CoverTab[107567]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:152
		_go_fuzz_dep_.CoverTab[107568]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:152
		// _ = "end of CoverTab[107568]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:152
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:152
	// _ = "end of CoverTab[107563]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:152
	_go_fuzz_dep_.CoverTab[107564]++

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:155
	if home := os.Getenv("USERPROFILE"); home != "" {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:155
		_go_fuzz_dep_.CoverTab[107569]++
													return home, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:156
		// _ = "end of CoverTab[107569]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:157
		_go_fuzz_dep_.CoverTab[107570]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:157
		// _ = "end of CoverTab[107570]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:157
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:157
	// _ = "end of CoverTab[107564]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:157
	_go_fuzz_dep_.CoverTab[107565]++

												drive := os.Getenv("HOMEDRIVE")
												path := os.Getenv("HOMEPATH")
												home := drive + path
												if drive == "" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:162
		_go_fuzz_dep_.CoverTab[107571]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:162
		return path == ""
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:162
		// _ = "end of CoverTab[107571]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:162
	}() {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:162
		_go_fuzz_dep_.CoverTab[107572]++
													return "", errors.New("HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:163
		// _ = "end of CoverTab[107572]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:164
		_go_fuzz_dep_.CoverTab[107573]++
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:164
		// _ = "end of CoverTab[107573]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:164
	}
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:164
	// _ = "end of CoverTab[107565]"
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:164
	_go_fuzz_dep_.CoverTab[107566]++

												return home, nil
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:166
	// _ = "end of CoverTab[107566]"
}

//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:167
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/mitchellh/go-homedir@v1.1.0/homedir.go:167
var _ = _go_fuzz_dep_.CoverTab
