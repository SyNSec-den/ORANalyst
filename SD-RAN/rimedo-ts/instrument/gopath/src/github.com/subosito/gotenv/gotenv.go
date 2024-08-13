//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:1
// Package gotenv provides functionality to dynamically load the environment variables
package gotenv

//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:2
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:2
)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:2
import (
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:2
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:2
)

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	// Pattern for detecting valid line format
	linePattern	= `\A\s*(?:export\s+)?([\w\.]+)(?:\s*=\s*|:\s+?)('(?:\'|[^'])*'|"(?:\"|[^"])*"|[^#\n]+)?\s*(?:\s*\#.*)?\z`

	// Pattern for detecting valid variable within a value
	variablePattern	= `(\\)?(\$)(\{?([A-Z0-9_]+)?\}?)`
)

// Env holds key/value pair of valid environment variable
type Env map[string]string

/*
Load is a function to load a file or multiple files and then export the valid variables into environment variables if they do not exist.
When it's called with no argument, it will load `.env` file on the current path and set the environment variables.
Otherwise, it will loop over the filenames parameter and set the proper environment variables.
*/
func Load(filenames ...string) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:29
	_go_fuzz_dep_.CoverTab[128146]++
											return loadenv(false, filenames...)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:30
	// _ = "end of CoverTab[128146]"
}

/*
OverLoad is a function to load a file or multiple files and then export and override the valid variables into environment variables.
*/
func OverLoad(filenames ...string) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:36
	_go_fuzz_dep_.CoverTab[128147]++
											return loadenv(true, filenames...)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:37
	// _ = "end of CoverTab[128147]"
}

/*
Must is wrapper function that will panic when supplied function returns an error.
*/
func Must(fn func(filenames ...string) error, filenames ...string) {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:43
	_go_fuzz_dep_.CoverTab[128148]++
											if err := fn(filenames...); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:44
		_go_fuzz_dep_.CoverTab[128149]++
												panic(err.Error())
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:45
		// _ = "end of CoverTab[128149]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:46
		_go_fuzz_dep_.CoverTab[128150]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:46
		// _ = "end of CoverTab[128150]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:46
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:46
	// _ = "end of CoverTab[128148]"
}

/*
Apply is a function to load an io Reader then export the valid variables into environment variables if they do not exist.
*/
func Apply(r io.Reader) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:52
	_go_fuzz_dep_.CoverTab[128151]++
											return parset(r, false)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:53
	// _ = "end of CoverTab[128151]"
}

/*
OverApply is a function to load an io Reader then export and override the valid variables into environment variables.
*/
func OverApply(r io.Reader) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:59
	_go_fuzz_dep_.CoverTab[128152]++
											return parset(r, true)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:60
	// _ = "end of CoverTab[128152]"
}

func loadenv(override bool, filenames ...string) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:63
	_go_fuzz_dep_.CoverTab[128153]++
											if len(filenames) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:64
		_go_fuzz_dep_.CoverTab[128156]++
												filenames = []string{".env"}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:65
		// _ = "end of CoverTab[128156]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:66
		_go_fuzz_dep_.CoverTab[128157]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:66
		// _ = "end of CoverTab[128157]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:66
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:66
	// _ = "end of CoverTab[128153]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:66
	_go_fuzz_dep_.CoverTab[128154]++

											for _, filename := range filenames {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:68
		_go_fuzz_dep_.CoverTab[128158]++
												f, err := os.Open(filename)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:70
			_go_fuzz_dep_.CoverTab[128161]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:71
			// _ = "end of CoverTab[128161]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:72
			_go_fuzz_dep_.CoverTab[128162]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:72
			// _ = "end of CoverTab[128162]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:72
		}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:72
		// _ = "end of CoverTab[128158]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:72
		_go_fuzz_dep_.CoverTab[128159]++

												err = parset(f, override)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:75
			_go_fuzz_dep_.CoverTab[128163]++
													return err
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:76
			// _ = "end of CoverTab[128163]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:77
			_go_fuzz_dep_.CoverTab[128164]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:77
			// _ = "end of CoverTab[128164]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:77
		}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:77
		// _ = "end of CoverTab[128159]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:77
		_go_fuzz_dep_.CoverTab[128160]++

												f.Close()
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:79
		// _ = "end of CoverTab[128160]"
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:80
	// _ = "end of CoverTab[128154]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:80
	_go_fuzz_dep_.CoverTab[128155]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:82
	// _ = "end of CoverTab[128155]"
}

// parse and set :)
func parset(r io.Reader, override bool) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:86
	_go_fuzz_dep_.CoverTab[128165]++
											env, err := StrictParse(r)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:88
		_go_fuzz_dep_.CoverTab[128168]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:89
		// _ = "end of CoverTab[128168]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:90
		_go_fuzz_dep_.CoverTab[128169]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:90
		// _ = "end of CoverTab[128169]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:90
	// _ = "end of CoverTab[128165]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:90
	_go_fuzz_dep_.CoverTab[128166]++

											for key, val := range env {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:92
		_go_fuzz_dep_.CoverTab[128170]++
												setenv(key, val, override)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:93
		// _ = "end of CoverTab[128170]"
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:94
	// _ = "end of CoverTab[128166]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:94
	_go_fuzz_dep_.CoverTab[128167]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:96
	// _ = "end of CoverTab[128167]"
}

func setenv(key, val string, override bool) {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:99
	_go_fuzz_dep_.CoverTab[128171]++
											if override {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:100
		_go_fuzz_dep_.CoverTab[128172]++
												os.Setenv(key, val)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:101
		// _ = "end of CoverTab[128172]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:102
		_go_fuzz_dep_.CoverTab[128173]++
												if _, present := os.LookupEnv(key); !present {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:103
			_go_fuzz_dep_.CoverTab[128174]++
													os.Setenv(key, val)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:104
			// _ = "end of CoverTab[128174]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:105
			_go_fuzz_dep_.CoverTab[128175]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:105
			// _ = "end of CoverTab[128175]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:105
		}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:105
		// _ = "end of CoverTab[128173]"
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:106
	// _ = "end of CoverTab[128171]"
}

// Parse is a function to parse line by line any io.Reader supplied and returns the valid Env key/value pair of valid variables.
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:109
// It expands the value of a variable from the environment variable but does not set the value to the environment itself.
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:109
// This function is skipping any invalid lines and only processing the valid one.
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:112
func Parse(r io.Reader) Env {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:112
	_go_fuzz_dep_.CoverTab[128176]++
											env, _ := StrictParse(r)
											return env
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:114
	// _ = "end of CoverTab[128176]"
}

// StrictParse is a function to parse line by line any io.Reader supplied and returns the valid Env key/value pair of valid variables.
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:117
// It expands the value of a variable from the environment variable but does not set the value to the environment itself.
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:117
// This function is returning an error if there are any invalid lines.
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:120
func StrictParse(r io.Reader) (Env, error) {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:120
	_go_fuzz_dep_.CoverTab[128177]++
											env := make(Env)
											scanner := bufio.NewScanner(r)

											i := 1
											bom := string([]byte{239, 187, 191})

											for scanner.Scan() {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:127
		_go_fuzz_dep_.CoverTab[128179]++
												line := scanner.Text()

												if i == 1 {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:130
			_go_fuzz_dep_.CoverTab[128181]++
													line = strings.TrimPrefix(line, bom)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:131
			// _ = "end of CoverTab[128181]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:132
			_go_fuzz_dep_.CoverTab[128182]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:132
			// _ = "end of CoverTab[128182]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:132
		}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:132
		// _ = "end of CoverTab[128179]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:132
		_go_fuzz_dep_.CoverTab[128180]++

												i++

												err := parseLine(line, env)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:137
			_go_fuzz_dep_.CoverTab[128183]++
													return env, err
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:138
			// _ = "end of CoverTab[128183]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:139
			_go_fuzz_dep_.CoverTab[128184]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:139
			// _ = "end of CoverTab[128184]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:139
		}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:139
		// _ = "end of CoverTab[128180]"
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:140
	// _ = "end of CoverTab[128177]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:140
	_go_fuzz_dep_.CoverTab[128178]++

											return env, nil
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:142
	// _ = "end of CoverTab[128178]"
}

func parseLine(s string, env Env) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:145
	_go_fuzz_dep_.CoverTab[128185]++
											rl := regexp.MustCompile(linePattern)
											rm := rl.FindStringSubmatch(s)

											if len(rm) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:149
		_go_fuzz_dep_.CoverTab[128189]++
												return checkFormat(s, env)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:150
		// _ = "end of CoverTab[128189]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:151
		_go_fuzz_dep_.CoverTab[128190]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:151
		// _ = "end of CoverTab[128190]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:151
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:151
	// _ = "end of CoverTab[128185]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:151
	_go_fuzz_dep_.CoverTab[128186]++

											key := rm[1]
											val := rm[2]

//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:157
	hdq := strings.HasPrefix(val, `"`)

//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:160
	hsq := strings.HasPrefix(val, `'`)

//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:163
	val = strings.Trim(val, " ")

//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:166
	rq := regexp.MustCompile(`\A(['"])(.*)(['"])\z`)
	val = rq.ReplaceAllString(val, "$2")

	if hdq {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:169
		_go_fuzz_dep_.CoverTab[128191]++
												val = strings.Replace(val, `\n`, "\n", -1)
												val = strings.Replace(val, `\r`, "\r", -1)

//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:174
		re := regexp.MustCompile(`\\([^$])`)
												val = re.ReplaceAllString(val, "$1")
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:175
		// _ = "end of CoverTab[128191]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:176
		_go_fuzz_dep_.CoverTab[128192]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:176
		// _ = "end of CoverTab[128192]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:176
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:176
	// _ = "end of CoverTab[128186]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:176
	_go_fuzz_dep_.CoverTab[128187]++

											rv := regexp.MustCompile(variablePattern)
											fv := func(s string) string {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:179
		_go_fuzz_dep_.CoverTab[128193]++
												return varReplacement(s, hsq, env)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:180
		// _ = "end of CoverTab[128193]"
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:181
	// _ = "end of CoverTab[128187]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:181
	_go_fuzz_dep_.CoverTab[128188]++

											val = rv.ReplaceAllStringFunc(val, fv)
											val = parseVal(val, env)

											env[key] = val
											return nil
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:187
	// _ = "end of CoverTab[128188]"
}

func parseExport(st string, env Env) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:190
	_go_fuzz_dep_.CoverTab[128194]++
											if strings.HasPrefix(st, "export") {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:191
		_go_fuzz_dep_.CoverTab[128196]++
												vs := strings.SplitN(st, " ", 2)

												if len(vs) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:194
			_go_fuzz_dep_.CoverTab[128197]++
													if _, ok := env[vs[1]]; !ok {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:195
				_go_fuzz_dep_.CoverTab[128198]++
														return fmt.Errorf("line `%s` has an unset variable", st)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:196
				// _ = "end of CoverTab[128198]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:197
				_go_fuzz_dep_.CoverTab[128199]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:197
				// _ = "end of CoverTab[128199]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:197
			}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:197
			// _ = "end of CoverTab[128197]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:198
			_go_fuzz_dep_.CoverTab[128200]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:198
			// _ = "end of CoverTab[128200]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:198
		}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:198
		// _ = "end of CoverTab[128196]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:199
		_go_fuzz_dep_.CoverTab[128201]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:199
		// _ = "end of CoverTab[128201]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:199
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:199
	// _ = "end of CoverTab[128194]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:199
	_go_fuzz_dep_.CoverTab[128195]++

											return nil
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:201
	// _ = "end of CoverTab[128195]"
}

func varReplacement(s string, hsq bool, env Env) string {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:204
	_go_fuzz_dep_.CoverTab[128202]++
											if strings.HasPrefix(s, "\\") {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:205
		_go_fuzz_dep_.CoverTab[128207]++
												return strings.TrimPrefix(s, "\\")
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:206
		// _ = "end of CoverTab[128207]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:207
		_go_fuzz_dep_.CoverTab[128208]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:207
		// _ = "end of CoverTab[128208]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:207
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:207
	// _ = "end of CoverTab[128202]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:207
	_go_fuzz_dep_.CoverTab[128203]++

											if hsq {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:209
		_go_fuzz_dep_.CoverTab[128209]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:210
		// _ = "end of CoverTab[128209]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:211
		_go_fuzz_dep_.CoverTab[128210]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:211
		// _ = "end of CoverTab[128210]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:211
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:211
	// _ = "end of CoverTab[128203]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:211
	_go_fuzz_dep_.CoverTab[128204]++

											sn := `(\$)(\{?([A-Z0-9_]+)\}?)`
											rn := regexp.MustCompile(sn)
											mn := rn.FindStringSubmatch(s)

											if len(mn) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:217
		_go_fuzz_dep_.CoverTab[128211]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:218
		// _ = "end of CoverTab[128211]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:219
		_go_fuzz_dep_.CoverTab[128212]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:219
		// _ = "end of CoverTab[128212]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:219
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:219
	// _ = "end of CoverTab[128204]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:219
	_go_fuzz_dep_.CoverTab[128205]++

											v := mn[3]

											replace, ok := env[v]
											if !ok {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:224
		_go_fuzz_dep_.CoverTab[128213]++
												replace = os.Getenv(v)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:225
		// _ = "end of CoverTab[128213]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:226
		_go_fuzz_dep_.CoverTab[128214]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:226
		// _ = "end of CoverTab[128214]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:226
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:226
	// _ = "end of CoverTab[128205]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:226
	_go_fuzz_dep_.CoverTab[128206]++

											return replace
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:228
	// _ = "end of CoverTab[128206]"
}

func checkFormat(s string, env Env) error {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:231
	_go_fuzz_dep_.CoverTab[128215]++
											st := strings.TrimSpace(s)

											if (st == "") || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:234
		_go_fuzz_dep_.CoverTab[128218]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:234
		return strings.HasPrefix(st, "#")
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:234
		// _ = "end of CoverTab[128218]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:234
	}() {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:234
		_go_fuzz_dep_.CoverTab[128219]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:235
		// _ = "end of CoverTab[128219]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:236
		_go_fuzz_dep_.CoverTab[128220]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:236
		// _ = "end of CoverTab[128220]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:236
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:236
	// _ = "end of CoverTab[128215]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:236
	_go_fuzz_dep_.CoverTab[128216]++

											if err := parseExport(st, env); err != nil {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:238
		_go_fuzz_dep_.CoverTab[128221]++
												return err
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:239
		// _ = "end of CoverTab[128221]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:240
		_go_fuzz_dep_.CoverTab[128222]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:240
		// _ = "end of CoverTab[128222]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:240
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:240
	// _ = "end of CoverTab[128216]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:240
	_go_fuzz_dep_.CoverTab[128217]++

											return fmt.Errorf("line `%s` doesn't match format", s)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:242
	// _ = "end of CoverTab[128217]"
}

func parseVal(val string, env Env) string {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:245
	_go_fuzz_dep_.CoverTab[128223]++
											if strings.Contains(val, "=") {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:246
		_go_fuzz_dep_.CoverTab[128225]++
												if !(val == "\n" || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:247
			_go_fuzz_dep_.CoverTab[128226]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:247
			return val == "\r"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:247
			// _ = "end of CoverTab[128226]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:247
		}()) {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:247
			_go_fuzz_dep_.CoverTab[128227]++
													kv := strings.Split(val, "\n")

													if len(kv) == 1 {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:250
				_go_fuzz_dep_.CoverTab[128229]++
														kv = strings.Split(val, "\r")
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:251
				// _ = "end of CoverTab[128229]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:252
				_go_fuzz_dep_.CoverTab[128230]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:252
				// _ = "end of CoverTab[128230]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:252
			}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:252
			// _ = "end of CoverTab[128227]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:252
			_go_fuzz_dep_.CoverTab[128228]++

													if len(kv) > 1 {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:254
				_go_fuzz_dep_.CoverTab[128231]++
														val = kv[0]

														for i := 1; i < len(kv); i++ {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:257
					_go_fuzz_dep_.CoverTab[128232]++
															parseLine(kv[i], env)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:258
					// _ = "end of CoverTab[128232]"
				}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:259
				// _ = "end of CoverTab[128231]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:260
				_go_fuzz_dep_.CoverTab[128233]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:260
				// _ = "end of CoverTab[128233]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:260
			}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:260
			// _ = "end of CoverTab[128228]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:261
			_go_fuzz_dep_.CoverTab[128234]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:261
			// _ = "end of CoverTab[128234]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:261
		}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:261
		// _ = "end of CoverTab[128225]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:262
		_go_fuzz_dep_.CoverTab[128235]++
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:262
		// _ = "end of CoverTab[128235]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:262
	}
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:262
	// _ = "end of CoverTab[128223]"
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:262
	_go_fuzz_dep_.CoverTab[128224]++

											return val
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:264
	// _ = "end of CoverTab[128224]"
}

//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:265
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/subosito/gotenv@v1.2.0/gotenv.go:265
var _ = _go_fuzz_dep_.CoverTab
