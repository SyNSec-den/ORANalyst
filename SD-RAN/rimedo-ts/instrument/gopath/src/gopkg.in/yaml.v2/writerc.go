//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:1
package yaml

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:1
)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:1
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:1
)

// Set the writer error and return false.
func yaml_emitter_set_writer_error(emitter *yaml_emitter_t, problem string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:4
	_go_fuzz_dep_.CoverTab[127840]++
									emitter.error = yaml_WRITER_ERROR
									emitter.problem = problem
									return false
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:7
	// _ = "end of CoverTab[127840]"
}

// Flush the output buffer.
func yaml_emitter_flush(emitter *yaml_emitter_t) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:11
	_go_fuzz_dep_.CoverTab[127841]++
										if emitter.write_handler == nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:12
		_go_fuzz_dep_.CoverTab[127845]++
											panic("write handler not set")
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:13
		// _ = "end of CoverTab[127845]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:14
		_go_fuzz_dep_.CoverTab[127846]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:14
		// _ = "end of CoverTab[127846]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:14
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:14
	// _ = "end of CoverTab[127841]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:14
	_go_fuzz_dep_.CoverTab[127842]++

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:17
	if emitter.buffer_pos == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:17
		_go_fuzz_dep_.CoverTab[127847]++
											return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:18
		// _ = "end of CoverTab[127847]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:19
		_go_fuzz_dep_.CoverTab[127848]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:19
		// _ = "end of CoverTab[127848]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:19
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:19
	// _ = "end of CoverTab[127842]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:19
	_go_fuzz_dep_.CoverTab[127843]++

										if err := emitter.write_handler(emitter, emitter.buffer[:emitter.buffer_pos]); err != nil {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:21
		_go_fuzz_dep_.CoverTab[127849]++
											return yaml_emitter_set_writer_error(emitter, "write error: "+err.Error())
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:22
		// _ = "end of CoverTab[127849]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:23
		_go_fuzz_dep_.CoverTab[127850]++
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:23
		// _ = "end of CoverTab[127850]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:23
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:23
	// _ = "end of CoverTab[127843]"
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:23
	_go_fuzz_dep_.CoverTab[127844]++
										emitter.buffer_pos = 0
										return true
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:25
	// _ = "end of CoverTab[127844]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:26
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/yaml.v2@v2.4.0/writerc.go:26
var _ = _go_fuzz_dep_.CoverTab
