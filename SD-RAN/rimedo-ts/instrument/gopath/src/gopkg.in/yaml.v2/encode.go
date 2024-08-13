//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:1
)

import (
	"encoding"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// jsonNumber is the interface of the encoding/json.Number datatype.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:16
// Repeating the interface here avoids a dependency on encoding/json, and also
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:16
// supports other libraries like jsoniter, which use a similar datatype with
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:16
// the same interface. Detecting this interface is useful when dealing with
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:16
// structures containing json.Number, which is a string under the hood. The
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:16
// encoder should prefer the use of Int64(), Float64() and string(), in that
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:16
// order, when encoding this type.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:23
type jsonNumber interface {
	Float64() (float64, error)
	Int64() (int64, error)
	String() string
}

type encoder struct {
	emitter	yaml_emitter_t
	event	yaml_event_t
	out	[]byte
	flow	bool
	// doneInit holds whether the initial stream_start_event has been
	// emitted.
	doneInit	bool
}

func newEncoder() *encoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:39
	_go_fuzz_dep_.CoverTab[125965]++
									e := &encoder{}
									yaml_emitter_initialize(&e.emitter)
									yaml_emitter_set_output_string(&e.emitter, &e.out)
									yaml_emitter_set_unicode(&e.emitter, true)
									return e
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:44
	// _ = "end of CoverTab[125965]"
}

func newEncoderWithWriter(w io.Writer) *encoder {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:47
	_go_fuzz_dep_.CoverTab[125966]++
									e := &encoder{}
									yaml_emitter_initialize(&e.emitter)
									yaml_emitter_set_output_writer(&e.emitter, w)
									yaml_emitter_set_unicode(&e.emitter, true)
									return e
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:52
	// _ = "end of CoverTab[125966]"
}

func (e *encoder) init() {
	if e.doneInit {
		return
	}
	yaml_stream_start_event_initialize(&e.event, yaml_UTF8_ENCODING)
	e.emit()
	e.doneInit = true
}

func (e *encoder) finish() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:64
	_go_fuzz_dep_.CoverTab[125967]++
									e.emitter.open_ended = false
									yaml_stream_end_event_initialize(&e.event)
									e.emit()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:67
	// _ = "end of CoverTab[125967]"
}

func (e *encoder) destroy() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:70
	_go_fuzz_dep_.CoverTab[125968]++
									yaml_emitter_delete(&e.emitter)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:71
	// _ = "end of CoverTab[125968]"
}

func (e *encoder) emit() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:74
	_go_fuzz_dep_.CoverTab[125969]++

									e.must(yaml_emitter_emit(&e.emitter, &e.event))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:76
	// _ = "end of CoverTab[125969]"
}

func (e *encoder) must(ok bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:79
	_go_fuzz_dep_.CoverTab[125970]++
									if !ok {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:80
		_go_fuzz_dep_.CoverTab[125971]++
										msg := e.emitter.problem
										if msg == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:82
			_go_fuzz_dep_.CoverTab[125973]++
											msg = "unknown problem generating YAML content"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:83
			// _ = "end of CoverTab[125973]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:84
			_go_fuzz_dep_.CoverTab[125974]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:84
			// _ = "end of CoverTab[125974]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:84
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:84
		// _ = "end of CoverTab[125971]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:84
		_go_fuzz_dep_.CoverTab[125972]++
										failf("%s", msg)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:85
		// _ = "end of CoverTab[125972]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:86
		_go_fuzz_dep_.CoverTab[125975]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:86
		// _ = "end of CoverTab[125975]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:86
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:86
	// _ = "end of CoverTab[125970]"
}

func (e *encoder) marshalDoc(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:89
	_go_fuzz_dep_.CoverTab[125976]++
									e.init()
									yaml_document_start_event_initialize(&e.event, nil, nil, true)
									e.emit()
									e.marshal(tag, in)
									yaml_document_end_event_initialize(&e.event, true)
									e.emit()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:95
	// _ = "end of CoverTab[125976]"
}

func (e *encoder) marshal(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:98
	_go_fuzz_dep_.CoverTab[125977]++
									if !in.IsValid() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
		_go_fuzz_dep_.CoverTab[125980]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
		return in.Kind() == reflect.Ptr && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
			_go_fuzz_dep_.CoverTab[125981]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
			return in.IsNil()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
			// _ = "end of CoverTab[125981]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
		// _ = "end of CoverTab[125980]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:99
			_go_fuzz_dep_.CoverTab[125982]++
											e.nilv()
											return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:101
		// _ = "end of CoverTab[125982]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:102
		_go_fuzz_dep_.CoverTab[125983]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:102
		// _ = "end of CoverTab[125983]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:102
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:102
	// _ = "end of CoverTab[125977]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:102
	_go_fuzz_dep_.CoverTab[125978]++
										iface := in.Interface()
										switch m := iface.(type) {
	case jsonNumber:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:105
		_go_fuzz_dep_.CoverTab[125984]++
											integer, err := m.Int64()
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:107
			_go_fuzz_dep_.CoverTab[125994]++

												in = reflect.ValueOf(integer)
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:110
			// _ = "end of CoverTab[125994]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:111
			_go_fuzz_dep_.CoverTab[125995]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:111
			// _ = "end of CoverTab[125995]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:111
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:111
		// _ = "end of CoverTab[125984]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:111
		_go_fuzz_dep_.CoverTab[125985]++
											float, err := m.Float64()
											if err == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:113
			_go_fuzz_dep_.CoverTab[125996]++

												in = reflect.ValueOf(float)
												break
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:116
			// _ = "end of CoverTab[125996]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:117
			_go_fuzz_dep_.CoverTab[125997]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:117
			// _ = "end of CoverTab[125997]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:117
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:117
		// _ = "end of CoverTab[125985]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:117
		_go_fuzz_dep_.CoverTab[125986]++

											in = reflect.ValueOf(m.String())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:119
		// _ = "end of CoverTab[125986]"
	case time.Time, *time.Time:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:120
		_go_fuzz_dep_.CoverTab[125987]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:120
		// _ = "end of CoverTab[125987]"

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:125
	case Marshaler:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:125
		_go_fuzz_dep_.CoverTab[125988]++
											v, err := m.MarshalYAML()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:127
			_go_fuzz_dep_.CoverTab[125998]++
												fail(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:128
			// _ = "end of CoverTab[125998]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:129
			_go_fuzz_dep_.CoverTab[125999]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:129
			// _ = "end of CoverTab[125999]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:129
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:129
		// _ = "end of CoverTab[125988]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:129
		_go_fuzz_dep_.CoverTab[125989]++
											if v == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:130
			_go_fuzz_dep_.CoverTab[126000]++
												e.nilv()
												return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:132
			// _ = "end of CoverTab[126000]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:133
			_go_fuzz_dep_.CoverTab[126001]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:133
			// _ = "end of CoverTab[126001]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:133
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:133
		// _ = "end of CoverTab[125989]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:133
		_go_fuzz_dep_.CoverTab[125990]++
											in = reflect.ValueOf(v)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:134
		// _ = "end of CoverTab[125990]"
	case encoding.TextMarshaler:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:135
		_go_fuzz_dep_.CoverTab[125991]++
											text, err := m.MarshalText()
											if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:137
			_go_fuzz_dep_.CoverTab[126002]++
												fail(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:138
			// _ = "end of CoverTab[126002]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:139
			_go_fuzz_dep_.CoverTab[126003]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:139
			// _ = "end of CoverTab[126003]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:139
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:139
		// _ = "end of CoverTab[125991]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:139
		_go_fuzz_dep_.CoverTab[125992]++
											in = reflect.ValueOf(string(text))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:140
		// _ = "end of CoverTab[125992]"
	case nil:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:141
		_go_fuzz_dep_.CoverTab[125993]++
											e.nilv()
											return
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:143
		// _ = "end of CoverTab[125993]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:144
	// _ = "end of CoverTab[125978]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:144
	_go_fuzz_dep_.CoverTab[125979]++
										switch in.Kind() {
	case reflect.Interface:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:146
		_go_fuzz_dep_.CoverTab[126004]++
											e.marshal(tag, in.Elem())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:147
		// _ = "end of CoverTab[126004]"
	case reflect.Map:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:148
		_go_fuzz_dep_.CoverTab[126005]++
											e.mapv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:149
		// _ = "end of CoverTab[126005]"
	case reflect.Ptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:150
		_go_fuzz_dep_.CoverTab[126006]++
											if in.Type() == ptrTimeType {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:151
			_go_fuzz_dep_.CoverTab[126015]++
												e.timev(tag, in.Elem())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:152
			// _ = "end of CoverTab[126015]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:153
			_go_fuzz_dep_.CoverTab[126016]++
												e.marshal(tag, in.Elem())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:154
			// _ = "end of CoverTab[126016]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:155
		// _ = "end of CoverTab[126006]"
	case reflect.Struct:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:156
		_go_fuzz_dep_.CoverTab[126007]++
											if in.Type() == timeType {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:157
			_go_fuzz_dep_.CoverTab[126017]++
												e.timev(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:158
			// _ = "end of CoverTab[126017]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:159
			_go_fuzz_dep_.CoverTab[126018]++
												e.structv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:160
			// _ = "end of CoverTab[126018]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:161
		// _ = "end of CoverTab[126007]"
	case reflect.Slice, reflect.Array:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:162
		_go_fuzz_dep_.CoverTab[126008]++
											if in.Type().Elem() == mapItemType {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:163
			_go_fuzz_dep_.CoverTab[126019]++
												e.itemsv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:164
			// _ = "end of CoverTab[126019]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:165
			_go_fuzz_dep_.CoverTab[126020]++
												e.slicev(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:166
			// _ = "end of CoverTab[126020]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:167
		// _ = "end of CoverTab[126008]"
	case reflect.String:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:168
		_go_fuzz_dep_.CoverTab[126009]++
											e.stringv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:169
		// _ = "end of CoverTab[126009]"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:170
		_go_fuzz_dep_.CoverTab[126010]++
											if in.Type() == durationType {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:171
			_go_fuzz_dep_.CoverTab[126021]++
												e.stringv(tag, reflect.ValueOf(iface.(time.Duration).String()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:172
			// _ = "end of CoverTab[126021]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:173
			_go_fuzz_dep_.CoverTab[126022]++
												e.intv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:174
			// _ = "end of CoverTab[126022]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:175
		// _ = "end of CoverTab[126010]"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:176
		_go_fuzz_dep_.CoverTab[126011]++
											e.uintv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:177
		// _ = "end of CoverTab[126011]"
	case reflect.Float32, reflect.Float64:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:178
		_go_fuzz_dep_.CoverTab[126012]++
											e.floatv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:179
		// _ = "end of CoverTab[126012]"
	case reflect.Bool:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:180
		_go_fuzz_dep_.CoverTab[126013]++
											e.boolv(tag, in)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:181
		// _ = "end of CoverTab[126013]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:182
		_go_fuzz_dep_.CoverTab[126014]++
											panic("cannot marshal type: " + in.Type().String())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:183
		// _ = "end of CoverTab[126014]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:184
	// _ = "end of CoverTab[125979]"
}

func (e *encoder) mapv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:187
	_go_fuzz_dep_.CoverTab[126023]++
										e.mappingv(tag, func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:188
		_go_fuzz_dep_.CoverTab[126024]++
											keys := keyList(in.MapKeys())
											sort.Sort(keys)
											for _, k := range keys {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:191
			_go_fuzz_dep_.CoverTab[126025]++
												e.marshal("", k)
												e.marshal("", in.MapIndex(k))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:193
			// _ = "end of CoverTab[126025]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:194
		// _ = "end of CoverTab[126024]"
	})
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:195
	// _ = "end of CoverTab[126023]"
}

func (e *encoder) itemsv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:198
	_go_fuzz_dep_.CoverTab[126026]++
										e.mappingv(tag, func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:199
		_go_fuzz_dep_.CoverTab[126027]++
											slice := in.Convert(reflect.TypeOf([]MapItem{})).Interface().([]MapItem)
											for _, item := range slice {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:201
			_go_fuzz_dep_.CoverTab[126028]++
												e.marshal("", reflect.ValueOf(item.Key))
												e.marshal("", reflect.ValueOf(item.Value))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:203
			// _ = "end of CoverTab[126028]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:204
		// _ = "end of CoverTab[126027]"
	})
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:205
	// _ = "end of CoverTab[126026]"
}

func (e *encoder) structv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:208
	_go_fuzz_dep_.CoverTab[126029]++
										sinfo, err := getStructInfo(in.Type())
										if err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:210
		_go_fuzz_dep_.CoverTab[126031]++
											panic(err)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:211
		// _ = "end of CoverTab[126031]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:212
		_go_fuzz_dep_.CoverTab[126032]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:212
		// _ = "end of CoverTab[126032]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:212
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:212
	// _ = "end of CoverTab[126029]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:212
	_go_fuzz_dep_.CoverTab[126030]++
										e.mappingv(tag, func() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:213
		_go_fuzz_dep_.CoverTab[126033]++
											for _, info := range sinfo.FieldsList {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:214
			_go_fuzz_dep_.CoverTab[126035]++
												var value reflect.Value
												if info.Inline == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:216
				_go_fuzz_dep_.CoverTab[126038]++
													value = in.Field(info.Num)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:217
				// _ = "end of CoverTab[126038]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:218
				_go_fuzz_dep_.CoverTab[126039]++
													value = in.FieldByIndex(info.Inline)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:219
				// _ = "end of CoverTab[126039]"
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:220
			// _ = "end of CoverTab[126035]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:220
			_go_fuzz_dep_.CoverTab[126036]++
												if info.OmitEmpty && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:221
				_go_fuzz_dep_.CoverTab[126040]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:221
				return isZero(value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:221
				// _ = "end of CoverTab[126040]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:221
			}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:221
				_go_fuzz_dep_.CoverTab[126041]++
													continue
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:222
				// _ = "end of CoverTab[126041]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:223
				_go_fuzz_dep_.CoverTab[126042]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:223
				// _ = "end of CoverTab[126042]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:223
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:223
			// _ = "end of CoverTab[126036]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:223
			_go_fuzz_dep_.CoverTab[126037]++
												e.marshal("", reflect.ValueOf(info.Key))
												e.flow = info.Flow
												e.marshal("", value)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:226
			// _ = "end of CoverTab[126037]"
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:227
		// _ = "end of CoverTab[126033]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:227
		_go_fuzz_dep_.CoverTab[126034]++
											if sinfo.InlineMap >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:228
			_go_fuzz_dep_.CoverTab[126043]++
												m := in.Field(sinfo.InlineMap)
												if m.Len() > 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:230
				_go_fuzz_dep_.CoverTab[126044]++
													e.flow = false
													keys := keyList(m.MapKeys())
													sort.Sort(keys)
													for _, k := range keys {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:234
					_go_fuzz_dep_.CoverTab[126045]++
														if _, found := sinfo.FieldsMap[k.String()]; found {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:235
						_go_fuzz_dep_.CoverTab[126047]++
															panic(fmt.Sprintf("Can't have key %q in inlined map; conflicts with struct field", k.String()))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:236
						// _ = "end of CoverTab[126047]"
					} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:237
						_go_fuzz_dep_.CoverTab[126048]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:237
						// _ = "end of CoverTab[126048]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:237
					}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:237
					// _ = "end of CoverTab[126045]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:237
					_go_fuzz_dep_.CoverTab[126046]++
														e.marshal("", k)
														e.flow = false
														e.marshal("", m.MapIndex(k))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:240
					// _ = "end of CoverTab[126046]"
				}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:241
				// _ = "end of CoverTab[126044]"
			} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:242
				_go_fuzz_dep_.CoverTab[126049]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:242
				// _ = "end of CoverTab[126049]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:242
			}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:242
			// _ = "end of CoverTab[126043]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:243
			_go_fuzz_dep_.CoverTab[126050]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:243
			// _ = "end of CoverTab[126050]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:243
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:243
		// _ = "end of CoverTab[126034]"
	})
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:244
	// _ = "end of CoverTab[126030]"
}

func (e *encoder) mappingv(tag string, f func()) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:247
	_go_fuzz_dep_.CoverTab[126051]++
										implicit := tag == ""
										style := yaml_BLOCK_MAPPING_STYLE
										if e.flow {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:250
		_go_fuzz_dep_.CoverTab[126053]++
											e.flow = false
											style = yaml_FLOW_MAPPING_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:252
		// _ = "end of CoverTab[126053]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:253
		_go_fuzz_dep_.CoverTab[126054]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:253
		// _ = "end of CoverTab[126054]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:253
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:253
	// _ = "end of CoverTab[126051]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:253
	_go_fuzz_dep_.CoverTab[126052]++
										yaml_mapping_start_event_initialize(&e.event, nil, []byte(tag), implicit, style)
										e.emit()
										f()
										yaml_mapping_end_event_initialize(&e.event)
										e.emit()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:258
	// _ = "end of CoverTab[126052]"
}

func (e *encoder) slicev(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:261
	_go_fuzz_dep_.CoverTab[126055]++
										implicit := tag == ""
										style := yaml_BLOCK_SEQUENCE_STYLE
										if e.flow {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:264
		_go_fuzz_dep_.CoverTab[126058]++
											e.flow = false
											style = yaml_FLOW_SEQUENCE_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:266
		// _ = "end of CoverTab[126058]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:267
		_go_fuzz_dep_.CoverTab[126059]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:267
		// _ = "end of CoverTab[126059]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:267
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:267
	// _ = "end of CoverTab[126055]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:267
	_go_fuzz_dep_.CoverTab[126056]++
										e.must(yaml_sequence_start_event_initialize(&e.event, nil, []byte(tag), implicit, style))
										e.emit()
										n := in.Len()
										for i := 0; i < n; i++ {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:271
		_go_fuzz_dep_.CoverTab[126060]++
											e.marshal("", in.Index(i))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:272
		// _ = "end of CoverTab[126060]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:273
	// _ = "end of CoverTab[126056]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:273
	_go_fuzz_dep_.CoverTab[126057]++
										e.must(yaml_sequence_end_event_initialize(&e.event))
										e.emit()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:275
	// _ = "end of CoverTab[126057]"
}

// isBase60 returns whether s is in base 60 notation as defined in YAML 1.1.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:278
//
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:278
// The base 60 float notation in YAML 1.1 is a terrible idea and is unsupported
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:278
// in YAML 1.2 and by this package, but these should be marshalled quoted for
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:278
// the time being for compatibility with other parsers.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:283
func isBase60Float(s string) (result bool) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:283
	_go_fuzz_dep_.CoverTab[126061]++

										if s == "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:285
		_go_fuzz_dep_.CoverTab[126064]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:286
		// _ = "end of CoverTab[126064]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:287
		_go_fuzz_dep_.CoverTab[126065]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:287
		// _ = "end of CoverTab[126065]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:287
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:287
	// _ = "end of CoverTab[126061]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:287
	_go_fuzz_dep_.CoverTab[126062]++
										c := s[0]
										if !(c == '+' || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		_go_fuzz_dep_.CoverTab[126066]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		return c == '-'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		// _ = "end of CoverTab[126066]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		_go_fuzz_dep_.CoverTab[126067]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		return c >= '0' && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
			_go_fuzz_dep_.CoverTab[126068]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
			return c <= '9'
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
			// _ = "end of CoverTab[126068]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		// _ = "end of CoverTab[126067]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
	}()) || func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		_go_fuzz_dep_.CoverTab[126069]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		return strings.IndexByte(s, ':') < 0
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		// _ = "end of CoverTab[126069]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
	}() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:289
		_go_fuzz_dep_.CoverTab[126070]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:290
		// _ = "end of CoverTab[126070]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:291
		_go_fuzz_dep_.CoverTab[126071]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:291
		// _ = "end of CoverTab[126071]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:291
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:291
	// _ = "end of CoverTab[126062]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:291
	_go_fuzz_dep_.CoverTab[126063]++

										return base60float.MatchString(s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:293
	// _ = "end of CoverTab[126063]"
}

// From http://yaml.org/type/float.html, except the regular expression there
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:296
// is bogus. In practice parsers do not enforce the "\.[0-9_]*" suffix.
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:298
var base60float = regexp.MustCompile(`^[-+]?[0-9][0-9_]*(?::[0-5]?[0-9])+(?:\.[0-9_]*)?$`)

func (e *encoder) stringv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:300
	_go_fuzz_dep_.CoverTab[126072]++
										var style yaml_scalar_style_t
										s := in.String()
										canUsePlain := true
										switch {
	case !utf8.ValidString(s):
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:305
		_go_fuzz_dep_.CoverTab[126075]++
											if tag == yaml_BINARY_TAG {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:306
			_go_fuzz_dep_.CoverTab[126080]++
												failf("explicitly tagged !!binary data must be base64-encoded")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:307
			// _ = "end of CoverTab[126080]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:308
			_go_fuzz_dep_.CoverTab[126081]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:308
			// _ = "end of CoverTab[126081]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:308
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:308
		// _ = "end of CoverTab[126075]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:308
		_go_fuzz_dep_.CoverTab[126076]++
											if tag != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:309
			_go_fuzz_dep_.CoverTab[126082]++
												failf("cannot marshal invalid UTF-8 data as %s", shortTag(tag))
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:310
			// _ = "end of CoverTab[126082]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:311
			_go_fuzz_dep_.CoverTab[126083]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:311
			// _ = "end of CoverTab[126083]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:311
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:311
		// _ = "end of CoverTab[126076]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:311
		_go_fuzz_dep_.CoverTab[126077]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:314
		tag = yaml_BINARY_TAG
											s = encodeBase64(s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:315
		// _ = "end of CoverTab[126077]"
	case tag == "":
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:316
		_go_fuzz_dep_.CoverTab[126078]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:320
		rtag, _ := resolve("", s)
		canUsePlain = rtag == yaml_STR_TAG && func() bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
			_go_fuzz_dep_.CoverTab[126084]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
			return !isBase60Float(s)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
			// _ = "end of CoverTab[126084]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
		}()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
		// _ = "end of CoverTab[126078]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
		_go_fuzz_dep_.CoverTab[126079]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:321
		// _ = "end of CoverTab[126079]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:322
	// _ = "end of CoverTab[126072]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:322
	_go_fuzz_dep_.CoverTab[126073]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:326
	switch {
	case strings.Contains(s, "\n"):
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:327
		_go_fuzz_dep_.CoverTab[126085]++
											style = yaml_LITERAL_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:328
		// _ = "end of CoverTab[126085]"
	case canUsePlain:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:329
		_go_fuzz_dep_.CoverTab[126086]++
											style = yaml_PLAIN_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:330
		// _ = "end of CoverTab[126086]"
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:331
		_go_fuzz_dep_.CoverTab[126087]++
											style = yaml_DOUBLE_QUOTED_SCALAR_STYLE
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:332
		// _ = "end of CoverTab[126087]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:333
	// _ = "end of CoverTab[126073]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:333
	_go_fuzz_dep_.CoverTab[126074]++
										e.emitScalar(s, "", tag, style)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:334
	// _ = "end of CoverTab[126074]"
}

func (e *encoder) boolv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:337
	_go_fuzz_dep_.CoverTab[126088]++
										var s string
										if in.Bool() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:339
		_go_fuzz_dep_.CoverTab[126090]++
											s = "true"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:340
		// _ = "end of CoverTab[126090]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:341
		_go_fuzz_dep_.CoverTab[126091]++
											s = "false"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:342
		// _ = "end of CoverTab[126091]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:343
	// _ = "end of CoverTab[126088]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:343
	_go_fuzz_dep_.CoverTab[126089]++
										e.emitScalar(s, "", tag, yaml_PLAIN_SCALAR_STYLE)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:344
	// _ = "end of CoverTab[126089]"
}

func (e *encoder) intv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:347
	_go_fuzz_dep_.CoverTab[126092]++
										s := strconv.FormatInt(in.Int(), 10)
										e.emitScalar(s, "", tag, yaml_PLAIN_SCALAR_STYLE)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:349
	// _ = "end of CoverTab[126092]"
}

func (e *encoder) uintv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:352
	_go_fuzz_dep_.CoverTab[126093]++
										s := strconv.FormatUint(in.Uint(), 10)
										e.emitScalar(s, "", tag, yaml_PLAIN_SCALAR_STYLE)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:354
	// _ = "end of CoverTab[126093]"
}

func (e *encoder) timev(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:357
	_go_fuzz_dep_.CoverTab[126094]++
										t := in.Interface().(time.Time)
										s := t.Format(time.RFC3339Nano)
										e.emitScalar(s, "", tag, yaml_PLAIN_SCALAR_STYLE)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:360
	// _ = "end of CoverTab[126094]"
}

func (e *encoder) floatv(tag string, in reflect.Value) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:363
	_go_fuzz_dep_.CoverTab[126095]++

										precision := 64
										if in.Kind() == reflect.Float32 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:366
		_go_fuzz_dep_.CoverTab[126098]++
											precision = 32
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:367
		// _ = "end of CoverTab[126098]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:368
		_go_fuzz_dep_.CoverTab[126099]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:368
		// _ = "end of CoverTab[126099]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:368
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:368
	// _ = "end of CoverTab[126095]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:368
	_go_fuzz_dep_.CoverTab[126096]++

										s := strconv.FormatFloat(in.Float(), 'g', -1, precision)
										switch s {
	case "+Inf":
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:372
		_go_fuzz_dep_.CoverTab[126100]++
											s = ".inf"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:373
		// _ = "end of CoverTab[126100]"
	case "-Inf":
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:374
		_go_fuzz_dep_.CoverTab[126101]++
											s = "-.inf"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:375
		// _ = "end of CoverTab[126101]"
	case "NaN":
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:376
		_go_fuzz_dep_.CoverTab[126102]++
											s = ".nan"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:377
		// _ = "end of CoverTab[126102]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:377
	default:
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:377
		_go_fuzz_dep_.CoverTab[126103]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:377
		// _ = "end of CoverTab[126103]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:378
	// _ = "end of CoverTab[126096]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:378
	_go_fuzz_dep_.CoverTab[126097]++
										e.emitScalar(s, "", tag, yaml_PLAIN_SCALAR_STYLE)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:379
	// _ = "end of CoverTab[126097]"
}

func (e *encoder) nilv() {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:382
	_go_fuzz_dep_.CoverTab[126104]++
										e.emitScalar("null", "", "", yaml_PLAIN_SCALAR_STYLE)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:383
	// _ = "end of CoverTab[126104]"
}

func (e *encoder) emitScalar(value, anchor, tag string, style yaml_scalar_style_t) {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:386
	_go_fuzz_dep_.CoverTab[126105]++
										implicit := tag == ""
										e.must(yaml_scalar_event_initialize(&e.event, []byte(anchor), []byte(tag), []byte(value), implicit, implicit, style))
										e.emit()
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:389
	// _ = "end of CoverTab[126105]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:390
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/encode.go:390
var _ = _go_fuzz_dep_.CoverTab
