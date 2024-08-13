// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:5
package store

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:5
)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:5
import (
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:5
)

import (
	"strconv"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func put(node interface{}, path interface{}, entry Entry) error {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:15
	_go_fuzz_dep_.CoverTab[193762]++
														keys, err := utils.ParseStringPath(path.(string))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:17
		_go_fuzz_dep_.CoverTab[193766]++
															return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:18
		// _ = "end of CoverTab[193766]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:19
		_go_fuzz_dep_.CoverTab[193767]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:19
		// _ = "end of CoverTab[193767]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:19
	// _ = "end of CoverTab[193762]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:19
	_go_fuzz_dep_.CoverTab[193763]++
														for i := 0; i < len(keys)-1; i++ {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:20
		_go_fuzz_dep_.CoverTab[193768]++
															node, err = search(node, keys[i])
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:22
			_go_fuzz_dep_.CoverTab[193769]++
																return err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:23
			// _ = "end of CoverTab[193769]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:24
			_go_fuzz_dep_.CoverTab[193770]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:24
			// _ = "end of CoverTab[193770]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:24
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:24
		// _ = "end of CoverTab[193768]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:25
	// _ = "end of CoverTab[193763]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:25
	_go_fuzz_dep_.CoverTab[193764]++

														lastKey := keys[len(keys)-1]
														switch node.(type) {
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:29
		_go_fuzz_dep_.CoverTab[193771]++
															node.(map[string]interface{})[lastKey.(string)] = entry.Value
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:30
		// _ = "end of CoverTab[193771]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:31
	// _ = "end of CoverTab[193764]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:31
	_go_fuzz_dep_.CoverTab[193765]++

														return nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:33
	// _ = "end of CoverTab[193765]"
}

func get(node interface{}, path interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:36
	_go_fuzz_dep_.CoverTab[193772]++
														keys, err := utils.ParseStringPath(path.(string))
														if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:38
		_go_fuzz_dep_.CoverTab[193775]++
															return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:39
		// _ = "end of CoverTab[193775]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:40
		_go_fuzz_dep_.CoverTab[193776]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:40
		// _ = "end of CoverTab[193776]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:40
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:40
	// _ = "end of CoverTab[193772]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:40
	_go_fuzz_dep_.CoverTab[193773]++
														for _, key := range keys {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:41
		_go_fuzz_dep_.CoverTab[193777]++
															node, err = search(node, key)
															if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:43
			_go_fuzz_dep_.CoverTab[193778]++
																log.Info(err)
																return nil, err
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:45
			// _ = "end of CoverTab[193778]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:46
			_go_fuzz_dep_.CoverTab[193779]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:46
			// _ = "end of CoverTab[193779]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:46
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:46
		// _ = "end of CoverTab[193777]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:47
	// _ = "end of CoverTab[193773]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:47
	_go_fuzz_dep_.CoverTab[193774]++

														return node, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:49
	// _ = "end of CoverTab[193774]"
}

func search(node interface{}, key interface{}) (interface{}, error) {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:52
	_go_fuzz_dep_.CoverTab[193780]++
														switch node.(type) {
	case []interface{}:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:54
		_go_fuzz_dep_.CoverTab[193783]++
															switch key.(type) {
		case map[string]string:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:56
			_go_fuzz_dep_.CoverTab[193787]++
																keys := key.(map[string]string)
																array := node.([]interface{})
																for k, v := range keys {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:59
				_go_fuzz_dep_.CoverTab[193788]++
																	for index, value := range array {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:60
					_go_fuzz_dep_.CoverTab[193789]++
																		switch vt := value.(type) {
					case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:62
						_go_fuzz_dep_.CoverTab[193790]++
																			valueMap := value.(map[string]interface{})
																			switch valueMap[k].(type) {
						case string:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:65
							_go_fuzz_dep_.CoverTab[193791]++
																				if valueMap[k] == v {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:66
								_go_fuzz_dep_.CoverTab[193794]++
																					node = array[index]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:67
								// _ = "end of CoverTab[193794]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:68
								_go_fuzz_dep_.CoverTab[193795]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:68
								// _ = "end of CoverTab[193795]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:68
							}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:68
							// _ = "end of CoverTab[193791]"
						case float64:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:69
							_go_fuzz_dep_.CoverTab[193792]++
																				floatValue, _ := strconv.ParseFloat(v, 64)
																				if valueMap[k].(float64) == floatValue {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:71
								_go_fuzz_dep_.CoverTab[193796]++
																					node = array[index]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:72
								// _ = "end of CoverTab[193796]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:73
								_go_fuzz_dep_.CoverTab[193797]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:73
								// _ = "end of CoverTab[193797]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:73
							}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:73
							// _ = "end of CoverTab[193792]"
						default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:74
							_go_fuzz_dep_.CoverTab[193793]++
																				return nil, errors.New(errors.NotSupported, "type %v is not supported", vt)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:75
							// _ = "end of CoverTab[193793]"

						}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:77
						// _ = "end of CoverTab[193790]"

					}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:79
					// _ = "end of CoverTab[193789]"
				}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:80
				// _ = "end of CoverTab[193788]"
			}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:81
			// _ = "end of CoverTab[193787]"
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:82
		// _ = "end of CoverTab[193783]"
	case map[string]interface{}:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:83
		_go_fuzz_dep_.CoverTab[193784]++
															key, ok := key.(string)
															if !ok {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:85
			_go_fuzz_dep_.CoverTab[193798]++
																return nil, errors.New(errors.Unknown, "key is not a string")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:86
			// _ = "end of CoverTab[193798]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:87
			_go_fuzz_dep_.CoverTab[193799]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:87
			// _ = "end of CoverTab[193799]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:87
		}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:87
		// _ = "end of CoverTab[193784]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:87
		_go_fuzz_dep_.CoverTab[193785]++
															node = node.(map[string]interface{})[key]
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:88
		// _ = "end of CoverTab[193785]"
	default:
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:89
		_go_fuzz_dep_.CoverTab[193786]++
															return nil, errors.New(errors.NotSupported, "node can only be of types map[string]interface{} or []interface{}")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:90
		// _ = "end of CoverTab[193786]"
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:91
	// _ = "end of CoverTab[193780]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:91
	_go_fuzz_dep_.CoverTab[193781]++

														if node == nil {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:93
		_go_fuzz_dep_.CoverTab[193800]++
															return nil, errors.New(errors.NotFound, "cannot find the node")
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:94
		// _ = "end of CoverTab[193800]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:95
		_go_fuzz_dep_.CoverTab[193801]++
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:95
		// _ = "end of CoverTab[193801]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:95
	}
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:95
	// _ = "end of CoverTab[193781]"
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:95
	_go_fuzz_dep_.CoverTab[193782]++

														return node, nil
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:97
	// _ = "end of CoverTab[193782]"
}

//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:98
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/onosproject/onos-ric-sdk-go@v0.8.9/pkg/config/store/util.go:98
var _ = _go_fuzz_dep_.CoverTab
