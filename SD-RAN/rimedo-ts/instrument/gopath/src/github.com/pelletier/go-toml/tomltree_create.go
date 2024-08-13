//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:1
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:1
)

import (
	"fmt"
	"reflect"
	"time"
)

var kindToType = [reflect.String + 1]reflect.Type{
	reflect.Bool:		reflect.TypeOf(true),
	reflect.String:		reflect.TypeOf(""),
	reflect.Float32:	reflect.TypeOf(float64(1)),
	reflect.Float64:	reflect.TypeOf(float64(1)),
	reflect.Int:		reflect.TypeOf(int64(1)),
	reflect.Int8:		reflect.TypeOf(int64(1)),
	reflect.Int16:		reflect.TypeOf(int64(1)),
	reflect.Int32:		reflect.TypeOf(int64(1)),
	reflect.Int64:		reflect.TypeOf(int64(1)),
	reflect.Uint:		reflect.TypeOf(uint64(1)),
	reflect.Uint8:		reflect.TypeOf(uint64(1)),
	reflect.Uint16:		reflect.TypeOf(uint64(1)),
	reflect.Uint32:		reflect.TypeOf(uint64(1)),
	reflect.Uint64:		reflect.TypeOf(uint64(1)),
}

// typeFor returns a reflect.Type for a reflect.Kind, or nil if none is found.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:26
// supported values:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:26
// string, bool, int64, uint64, float64, time.Time, int, int8, int16, int32, uint, uint8, uint16, uint32, float32
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:29
func typeFor(k reflect.Kind) reflect.Type {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:29
	_go_fuzz_dep_.CoverTab[124125]++
												if k > 0 && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:30
		_go_fuzz_dep_.CoverTab[124127]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:30
		return int(k) < len(kindToType)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:30
		// _ = "end of CoverTab[124127]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:30
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:30
		_go_fuzz_dep_.CoverTab[124128]++
													return kindToType[k]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:31
		// _ = "end of CoverTab[124128]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:32
		_go_fuzz_dep_.CoverTab[124129]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:32
		// _ = "end of CoverTab[124129]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:32
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:32
	// _ = "end of CoverTab[124125]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:32
	_go_fuzz_dep_.CoverTab[124126]++
												return nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:33
	// _ = "end of CoverTab[124126]"
}

func simpleValueCoercion(object interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:36
	_go_fuzz_dep_.CoverTab[124130]++
												switch original := object.(type) {
	case string, bool, int64, uint64, float64, time.Time:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:38
		_go_fuzz_dep_.CoverTab[124131]++
													return original, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:39
		// _ = "end of CoverTab[124131]"
	case int:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:40
		_go_fuzz_dep_.CoverTab[124132]++
													return int64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:41
		// _ = "end of CoverTab[124132]"
	case int8:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:42
		_go_fuzz_dep_.CoverTab[124133]++
													return int64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:43
		// _ = "end of CoverTab[124133]"
	case int16:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:44
		_go_fuzz_dep_.CoverTab[124134]++
													return int64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:45
		// _ = "end of CoverTab[124134]"
	case int32:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:46
		_go_fuzz_dep_.CoverTab[124135]++
													return int64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:47
		// _ = "end of CoverTab[124135]"
	case uint:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:48
		_go_fuzz_dep_.CoverTab[124136]++
													return uint64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:49
		// _ = "end of CoverTab[124136]"
	case uint8:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:50
		_go_fuzz_dep_.CoverTab[124137]++
													return uint64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:51
		// _ = "end of CoverTab[124137]"
	case uint16:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:52
		_go_fuzz_dep_.CoverTab[124138]++
													return uint64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:53
		// _ = "end of CoverTab[124138]"
	case uint32:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:54
		_go_fuzz_dep_.CoverTab[124139]++
													return uint64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:55
		// _ = "end of CoverTab[124139]"
	case float32:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:56
		_go_fuzz_dep_.CoverTab[124140]++
													return float64(original), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:57
		// _ = "end of CoverTab[124140]"
	case fmt.Stringer:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:58
		_go_fuzz_dep_.CoverTab[124141]++
													return original.String(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:59
		// _ = "end of CoverTab[124141]"
	case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:60
		_go_fuzz_dep_.CoverTab[124142]++
													value := reflect.ValueOf(original)
													length := value.Len()
													arrayValue := reflect.MakeSlice(value.Type(), 0, length)
													for i := 0; i < length; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:64
			_go_fuzz_dep_.CoverTab[124145]++
														val := value.Index(i).Interface()
														simpleValue, err := simpleValueCoercion(val)
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:67
				_go_fuzz_dep_.CoverTab[124147]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:68
				// _ = "end of CoverTab[124147]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:69
				_go_fuzz_dep_.CoverTab[124148]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:69
				// _ = "end of CoverTab[124148]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:69
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:69
			// _ = "end of CoverTab[124145]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:69
			_go_fuzz_dep_.CoverTab[124146]++
														arrayValue = reflect.Append(arrayValue, reflect.ValueOf(simpleValue))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:70
			// _ = "end of CoverTab[124146]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:71
		// _ = "end of CoverTab[124142]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:71
		_go_fuzz_dep_.CoverTab[124143]++
													return arrayValue.Interface(), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:72
		// _ = "end of CoverTab[124143]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:73
		_go_fuzz_dep_.CoverTab[124144]++
													return nil, fmt.Errorf("cannot convert type %T to Tree", object)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:74
		// _ = "end of CoverTab[124144]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:75
	// _ = "end of CoverTab[124130]"
}

func sliceToTree(object interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:78
	_go_fuzz_dep_.CoverTab[124149]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:85
	value := reflect.ValueOf(object)
	insideType := value.Type().Elem()
	length := value.Len()
	if length > 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:88
		_go_fuzz_dep_.CoverTab[124154]++
													insideType = reflect.ValueOf(value.Index(0).Interface()).Type()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:89
		// _ = "end of CoverTab[124154]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:90
		_go_fuzz_dep_.CoverTab[124155]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:90
		// _ = "end of CoverTab[124155]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:90
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:90
	// _ = "end of CoverTab[124149]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:90
	_go_fuzz_dep_.CoverTab[124150]++
												if insideType.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:91
		_go_fuzz_dep_.CoverTab[124156]++

													tablesArray := make([]*Tree, 0, length)
													for i := 0; i < length; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:94
			_go_fuzz_dep_.CoverTab[124158]++
														table := value.Index(i)
														tree, err := toTree(table.Interface())
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:97
				_go_fuzz_dep_.CoverTab[124160]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:98
				// _ = "end of CoverTab[124160]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:99
				_go_fuzz_dep_.CoverTab[124161]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:99
				// _ = "end of CoverTab[124161]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:99
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:99
			// _ = "end of CoverTab[124158]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:99
			_go_fuzz_dep_.CoverTab[124159]++
														tablesArray = append(tablesArray, tree.(*Tree))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:100
			// _ = "end of CoverTab[124159]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:101
		// _ = "end of CoverTab[124156]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:101
		_go_fuzz_dep_.CoverTab[124157]++
													return tablesArray, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:102
		// _ = "end of CoverTab[124157]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:103
		_go_fuzz_dep_.CoverTab[124162]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:103
		// _ = "end of CoverTab[124162]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:103
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:103
	// _ = "end of CoverTab[124150]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:103
	_go_fuzz_dep_.CoverTab[124151]++

												sliceType := typeFor(insideType.Kind())
												if sliceType == nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:106
		_go_fuzz_dep_.CoverTab[124163]++
													sliceType = insideType
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:107
		// _ = "end of CoverTab[124163]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:108
		_go_fuzz_dep_.CoverTab[124164]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:108
		// _ = "end of CoverTab[124164]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:108
	// _ = "end of CoverTab[124151]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:108
	_go_fuzz_dep_.CoverTab[124152]++

												arrayValue := reflect.MakeSlice(reflect.SliceOf(sliceType), 0, length)

												for i := 0; i < length; i++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:112
		_go_fuzz_dep_.CoverTab[124165]++
													val := value.Index(i).Interface()
													simpleValue, err := simpleValueCoercion(val)
													if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:115
			_go_fuzz_dep_.CoverTab[124167]++
														return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:116
			// _ = "end of CoverTab[124167]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:117
			_go_fuzz_dep_.CoverTab[124168]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:117
			// _ = "end of CoverTab[124168]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:117
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:117
		// _ = "end of CoverTab[124165]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:117
		_go_fuzz_dep_.CoverTab[124166]++
													arrayValue = reflect.Append(arrayValue, reflect.ValueOf(simpleValue))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:118
		// _ = "end of CoverTab[124166]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:119
	// _ = "end of CoverTab[124152]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:119
	_go_fuzz_dep_.CoverTab[124153]++
												return &tomlValue{value: arrayValue.Interface(), position: Position{}}, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:120
	// _ = "end of CoverTab[124153]"
}

func toTree(object interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:123
	_go_fuzz_dep_.CoverTab[124169]++
												value := reflect.ValueOf(object)

												if value.Kind() == reflect.Map {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:126
		_go_fuzz_dep_.CoverTab[124173]++
													values := map[string]interface{}{}
													keys := value.MapKeys()
													for _, key := range keys {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:129
			_go_fuzz_dep_.CoverTab[124175]++
														if key.Kind() != reflect.String {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:130
				_go_fuzz_dep_.CoverTab[124178]++
															if _, ok := key.Interface().(string); !ok {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:131
					_go_fuzz_dep_.CoverTab[124179]++
																return nil, fmt.Errorf("map key needs to be a string, not %T (%v)", key.Interface(), key.Kind())
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:132
					// _ = "end of CoverTab[124179]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:133
					_go_fuzz_dep_.CoverTab[124180]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:133
					// _ = "end of CoverTab[124180]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:133
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:133
				// _ = "end of CoverTab[124178]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:134
				_go_fuzz_dep_.CoverTab[124181]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:134
				// _ = "end of CoverTab[124181]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:134
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:134
			// _ = "end of CoverTab[124175]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:134
			_go_fuzz_dep_.CoverTab[124176]++

														v := value.MapIndex(key)
														newValue, err := toTree(v.Interface())
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:138
				_go_fuzz_dep_.CoverTab[124182]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:139
				// _ = "end of CoverTab[124182]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:140
				_go_fuzz_dep_.CoverTab[124183]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:140
				// _ = "end of CoverTab[124183]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:140
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:140
			// _ = "end of CoverTab[124176]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:140
			_go_fuzz_dep_.CoverTab[124177]++
														values[key.String()] = newValue
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:141
			// _ = "end of CoverTab[124177]"
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:142
		// _ = "end of CoverTab[124173]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:142
		_go_fuzz_dep_.CoverTab[124174]++
													return &Tree{values: values, position: Position{}}, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:143
		// _ = "end of CoverTab[124174]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:144
		_go_fuzz_dep_.CoverTab[124184]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:144
		// _ = "end of CoverTab[124184]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:144
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:144
	// _ = "end of CoverTab[124169]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:144
	_go_fuzz_dep_.CoverTab[124170]++

												if value.Kind() == reflect.Array || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:146
		_go_fuzz_dep_.CoverTab[124185]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:146
		return value.Kind() == reflect.Slice
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:146
		// _ = "end of CoverTab[124185]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:146
	}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:146
		_go_fuzz_dep_.CoverTab[124186]++
													return sliceToTree(object)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:147
		// _ = "end of CoverTab[124186]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:148
		_go_fuzz_dep_.CoverTab[124187]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:148
		// _ = "end of CoverTab[124187]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:148
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:148
	// _ = "end of CoverTab[124170]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:148
	_go_fuzz_dep_.CoverTab[124171]++

												simpleValue, err := simpleValueCoercion(object)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:151
		_go_fuzz_dep_.CoverTab[124188]++
													return nil, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:152
		// _ = "end of CoverTab[124188]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:153
		_go_fuzz_dep_.CoverTab[124189]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:153
		// _ = "end of CoverTab[124189]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:153
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:153
	// _ = "end of CoverTab[124171]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:153
	_go_fuzz_dep_.CoverTab[124172]++
												return &tomlValue{value: simpleValue, position: Position{}}, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:154
	// _ = "end of CoverTab[124172]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:155
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/tomltree_create.go:155
var _ = _go_fuzz_dep_.CoverTab
