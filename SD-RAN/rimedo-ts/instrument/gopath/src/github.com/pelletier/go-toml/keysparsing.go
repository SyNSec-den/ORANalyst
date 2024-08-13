// Parsing keys handling both bare and quoted keys.

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:3
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:3
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:3
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:3
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:3
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:3
)

import (
	"errors"
	"fmt"
)

// Convert the bare key group string to an array.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:10
// The input supports double quotation and single quotation,
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:10
// but escape sequences are not supported. Lexers must unescape them beforehand.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:13
func parseKey(key string) ([]string, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:13
	_go_fuzz_dep_.CoverTab[122515]++
												runes := []rune(key)
												var groups []string

												if len(key) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:17
		_go_fuzz_dep_.CoverTab[122519]++
													return nil, errors.New("empty key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:18
		// _ = "end of CoverTab[122519]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:19
		_go_fuzz_dep_.CoverTab[122520]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:19
		// _ = "end of CoverTab[122520]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:19
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:19
	// _ = "end of CoverTab[122515]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:19
	_go_fuzz_dep_.CoverTab[122516]++

												idx := 0
												for idx < len(runes) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:22
		_go_fuzz_dep_.CoverTab[122521]++
													for ; idx < len(runes) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:23
			_go_fuzz_dep_.CoverTab[122524]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:23
			return isSpace(runes[idx])
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:23
			// _ = "end of CoverTab[122524]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:23
		}(); idx++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:23
			_go_fuzz_dep_.CoverTab[122525]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:23
			// _ = "end of CoverTab[122525]"

		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:25
		// _ = "end of CoverTab[122521]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:25
		_go_fuzz_dep_.CoverTab[122522]++
													if idx >= len(runes) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:26
			_go_fuzz_dep_.CoverTab[122526]++
														break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:27
			// _ = "end of CoverTab[122526]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:28
			_go_fuzz_dep_.CoverTab[122527]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:28
			// _ = "end of CoverTab[122527]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:28
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:28
		// _ = "end of CoverTab[122522]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:28
		_go_fuzz_dep_.CoverTab[122523]++
													r := runes[idx]
													if isValidBareChar(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:30
			_go_fuzz_dep_.CoverTab[122528]++

														startIdx := idx
														endIdx := -1
														idx++
														for idx < len(runes) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:35
				_go_fuzz_dep_.CoverTab[122531]++
															r = runes[idx]
															if isValidBareChar(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:37
					_go_fuzz_dep_.CoverTab[122532]++
																idx++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:38
					// _ = "end of CoverTab[122532]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:39
					_go_fuzz_dep_.CoverTab[122533]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:39
					if r == '.' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:39
						_go_fuzz_dep_.CoverTab[122534]++
																	endIdx = idx
																	break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:41
						// _ = "end of CoverTab[122534]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:42
						_go_fuzz_dep_.CoverTab[122535]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:42
						if isSpace(r) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:42
							_go_fuzz_dep_.CoverTab[122536]++
																		endIdx = idx
																		for ; idx < len(runes) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:44
								_go_fuzz_dep_.CoverTab[122539]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:44
								return isSpace(runes[idx])
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:44
								// _ = "end of CoverTab[122539]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:44
							}(); idx++ {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:44
								_go_fuzz_dep_.CoverTab[122540]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:44
								// _ = "end of CoverTab[122540]"

							}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:46
							// _ = "end of CoverTab[122536]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:46
							_go_fuzz_dep_.CoverTab[122537]++
																		if idx < len(runes) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:47
								_go_fuzz_dep_.CoverTab[122541]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:47
								return runes[idx] != '.'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:47
								// _ = "end of CoverTab[122541]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:47
							}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:47
								_go_fuzz_dep_.CoverTab[122542]++
																			return nil, fmt.Errorf("invalid key character after whitespace: %c", runes[idx])
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:48
								// _ = "end of CoverTab[122542]"
							} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:49
								_go_fuzz_dep_.CoverTab[122543]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:49
								// _ = "end of CoverTab[122543]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:49
							}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:49
							// _ = "end of CoverTab[122537]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:49
							_go_fuzz_dep_.CoverTab[122538]++
																		break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:50
							// _ = "end of CoverTab[122538]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:51
							_go_fuzz_dep_.CoverTab[122544]++
																		return nil, fmt.Errorf("invalid bare key character: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:52
							// _ = "end of CoverTab[122544]"
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:53
						// _ = "end of CoverTab[122535]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:53
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:53
					// _ = "end of CoverTab[122533]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:53
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:53
				// _ = "end of CoverTab[122531]"
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:54
			// _ = "end of CoverTab[122528]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:54
			_go_fuzz_dep_.CoverTab[122529]++
														if endIdx == -1 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:55
				_go_fuzz_dep_.CoverTab[122545]++
															endIdx = idx
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:56
				// _ = "end of CoverTab[122545]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:57
				_go_fuzz_dep_.CoverTab[122546]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:57
				// _ = "end of CoverTab[122546]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:57
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:57
			// _ = "end of CoverTab[122529]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:57
			_go_fuzz_dep_.CoverTab[122530]++
														groups = append(groups, string(runes[startIdx:endIdx]))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:58
			// _ = "end of CoverTab[122530]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:59
			_go_fuzz_dep_.CoverTab[122547]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:59
			if r == '\'' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:59
				_go_fuzz_dep_.CoverTab[122548]++

															idx++
															startIdx := idx
															for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:63
					_go_fuzz_dep_.CoverTab[122549]++
																if idx >= len(runes) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:64
						_go_fuzz_dep_.CoverTab[122552]++
																	return nil, fmt.Errorf("unclosed single-quoted key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:65
						// _ = "end of CoverTab[122552]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:66
						_go_fuzz_dep_.CoverTab[122553]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:66
						// _ = "end of CoverTab[122553]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:66
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:66
					// _ = "end of CoverTab[122549]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:66
					_go_fuzz_dep_.CoverTab[122550]++
																r = runes[idx]
																if r == '\'' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:68
						_go_fuzz_dep_.CoverTab[122554]++
																	groups = append(groups, string(runes[startIdx:idx]))
																	idx++
																	break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:71
						// _ = "end of CoverTab[122554]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:72
						_go_fuzz_dep_.CoverTab[122555]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:72
						// _ = "end of CoverTab[122555]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:72
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:72
					// _ = "end of CoverTab[122550]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:72
					_go_fuzz_dep_.CoverTab[122551]++
																idx++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:73
					// _ = "end of CoverTab[122551]"
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:74
				// _ = "end of CoverTab[122548]"
			} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:75
				_go_fuzz_dep_.CoverTab[122556]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:75
				if r == '"' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:75
					_go_fuzz_dep_.CoverTab[122557]++

																idx++
																startIdx := idx
																for {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:79
						_go_fuzz_dep_.CoverTab[122558]++
																	if idx >= len(runes) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:80
							_go_fuzz_dep_.CoverTab[122561]++
																		return nil, fmt.Errorf("unclosed double-quoted key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:81
							// _ = "end of CoverTab[122561]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:82
							_go_fuzz_dep_.CoverTab[122562]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:82
							// _ = "end of CoverTab[122562]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:82
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:82
						// _ = "end of CoverTab[122558]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:82
						_go_fuzz_dep_.CoverTab[122559]++
																	r = runes[idx]
																	if r == '"' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:84
							_go_fuzz_dep_.CoverTab[122563]++
																		groups = append(groups, string(runes[startIdx:idx]))
																		idx++
																		break
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:87
							// _ = "end of CoverTab[122563]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:88
							_go_fuzz_dep_.CoverTab[122564]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:88
							// _ = "end of CoverTab[122564]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:88
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:88
						// _ = "end of CoverTab[122559]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:88
						_go_fuzz_dep_.CoverTab[122560]++
																	idx++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:89
						// _ = "end of CoverTab[122560]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:90
					// _ = "end of CoverTab[122557]"
				} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:91
					_go_fuzz_dep_.CoverTab[122565]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:91
					if r == '.' {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:91
						_go_fuzz_dep_.CoverTab[122566]++
																	idx++
																	if idx >= len(runes) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:93
							_go_fuzz_dep_.CoverTab[122568]++
																		return nil, fmt.Errorf("unexpected end of key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:94
							// _ = "end of CoverTab[122568]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:95
							_go_fuzz_dep_.CoverTab[122569]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:95
							// _ = "end of CoverTab[122569]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:95
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:95
						// _ = "end of CoverTab[122566]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:95
						_go_fuzz_dep_.CoverTab[122567]++
																	r = runes[idx]
																	if !isValidBareChar(r) && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							_go_fuzz_dep_.CoverTab[122570]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							return r != '\''
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							// _ = "end of CoverTab[122570]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
						}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							_go_fuzz_dep_.CoverTab[122571]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							return r != '"'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							// _ = "end of CoverTab[122571]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
						}() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							_go_fuzz_dep_.CoverTab[122572]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							return r != ' '
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							// _ = "end of CoverTab[122572]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
						}() {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:97
							_go_fuzz_dep_.CoverTab[122573]++
																		return nil, fmt.Errorf("expecting key part after dot")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:98
							// _ = "end of CoverTab[122573]"
						} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:99
							_go_fuzz_dep_.CoverTab[122574]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:99
							// _ = "end of CoverTab[122574]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:99
						}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:99
						// _ = "end of CoverTab[122567]"
					} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:100
						_go_fuzz_dep_.CoverTab[122575]++
																	return nil, fmt.Errorf("invalid key character: %c", r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:101
						// _ = "end of CoverTab[122575]"
					}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:102
					// _ = "end of CoverTab[122565]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:102
				}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:102
				// _ = "end of CoverTab[122556]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:102
			}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:102
			// _ = "end of CoverTab[122547]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:102
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:102
		// _ = "end of CoverTab[122523]"
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:103
	// _ = "end of CoverTab[122516]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:103
	_go_fuzz_dep_.CoverTab[122517]++
												if len(groups) == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:104
		_go_fuzz_dep_.CoverTab[122576]++
													return nil, fmt.Errorf("empty key")
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:105
		// _ = "end of CoverTab[122576]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:106
		_go_fuzz_dep_.CoverTab[122577]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:106
		// _ = "end of CoverTab[122577]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:106
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:106
	// _ = "end of CoverTab[122517]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:106
	_go_fuzz_dep_.CoverTab[122518]++
												return groups, nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:107
	// _ = "end of CoverTab[122518]"
}

func isValidBareChar(r rune) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:110
	_go_fuzz_dep_.CoverTab[122578]++
												return isAlphanumeric(r) || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
		_go_fuzz_dep_.CoverTab[122579]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
		return r == '-'
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
		// _ = "end of CoverTab[122579]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
	}() || func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
		_go_fuzz_dep_.CoverTab[122580]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
		return isDigit(r)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
		// _ = "end of CoverTab[122580]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:111
	// _ = "end of CoverTab[122578]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:112
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/keysparsing.go:112
var _ = _go_fuzz_dep_.CoverTab
