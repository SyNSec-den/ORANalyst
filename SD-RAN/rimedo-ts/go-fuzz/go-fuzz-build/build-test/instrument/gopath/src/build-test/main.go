//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:1
package main

//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:1
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:1
)
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:1
import (
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:1
)

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func networkOperation() {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:11
	_go_fuzz_dep_.CoverTab[8826]++
												conn, err := net.Dial("tcp", "localhost:8080")
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:13
		_go_fuzz_dep_.CoverTab[8828]++
													fmt.Println("Error:", err)
													return
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:15
		// _ = "end of CoverTab[8828]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:16
		_go_fuzz_dep_.CoverTab[8829]++
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:16
		// _ = "end of CoverTab[8829]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:16
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:16
	// _ = "end of CoverTab[8826]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:16
	_go_fuzz_dep_.CoverTab[8827]++
												defer conn.Close()

//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:20
	fmt.Fprintf(conn, "Hello, Server!\n")
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:20
	// _ = "end of CoverTab[8827]"
}

func fileOperations(filename string) {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:23
	_go_fuzz_dep_.CoverTab[8830]++
												file, err := os.Create(filename)
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:25
		_go_fuzz_dep_.CoverTab[8834]++
													panic(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:26
		// _ = "end of CoverTab[8834]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:27
		_go_fuzz_dep_.CoverTab[8835]++
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:27
		// _ = "end of CoverTab[8835]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:27
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:27
	// _ = "end of CoverTab[8830]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:27
	_go_fuzz_dep_.CoverTab[8831]++
												defer file.Close()

												writer := bufio.NewWriter(file)
												for i := 0; i < 5; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:31
		_go_fuzz_dep_.CoverTab[8836]++
													writer.WriteString(fmt.Sprintf("Writing line %d\n", i))
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:32
		// _ = "end of CoverTab[8836]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:33
	// _ = "end of CoverTab[8831]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:33
	_go_fuzz_dep_.CoverTab[8832]++
												writer.Flush()

												data, err := ioutil.ReadFile(filename)
												if err != nil {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:37
		_go_fuzz_dep_.CoverTab[8837]++
													panic(err)
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:38
		// _ = "end of CoverTab[8837]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:39
		_go_fuzz_dep_.CoverTab[8838]++
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:39
		// _ = "end of CoverTab[8838]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:39
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:39
	// _ = "end of CoverTab[8832]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:39
	_go_fuzz_dep_.CoverTab[8833]++
												fmt.Println(string(data))
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:40
	// _ = "end of CoverTab[8833]"
}

func memoryOperations() {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:43
	_go_fuzz_dep_.CoverTab[8839]++
												data := make(map[string][]int)
												for i := 0; i < 5; i++ {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:45
		_go_fuzz_dep_.CoverTab[8841]++
													key := fmt.Sprintf("key_%d", i)
													data[key] = append(data[key], i, i+1, i+2)
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:47
		// _ = "end of CoverTab[8841]"
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:48
	// _ = "end of CoverTab[8839]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:48
	_go_fuzz_dep_.CoverTab[8840]++

												if v, exists := data["key_2"]; exists {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:50
		_go_fuzz_dep_.CoverTab[8842]++
													fmt.Println("Found:", v)
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:51
		// _ = "end of CoverTab[8842]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:52
		_go_fuzz_dep_.CoverTab[8843]++
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:52
		// _ = "end of CoverTab[8843]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:52
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:52
	// _ = "end of CoverTab[8840]"
}

func main() {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:55
	_go_fuzz_dep_.InstrumentMain()
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:55
	_go_fuzz_dep_.CoverTab[8844]++
												if len(os.Args) > 1 && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:56
		_go_fuzz_dep_.CoverTab[8845]++
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:56
		return os.Args[1] == "net"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:56
		// _ = "end of CoverTab[8845]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:56
	}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:56
		_go_fuzz_dep_.CoverTab[8846]++
													networkOperation()
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:57
		// _ = "end of CoverTab[8846]"
	} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:58
		_go_fuzz_dep_.CoverTab[8847]++
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:58
		if len(os.Args) > 1 && func() bool {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:58
			_go_fuzz_dep_.CoverTab[8848]++
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:58
			return os.Args[1] == "file"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:58
			// _ = "end of CoverTab[8848]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:58
		}() {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:58
			_go_fuzz_dep_.CoverTab[8849]++
														fileOperations("test.txt")
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:59
			// _ = "end of CoverTab[8849]"
		} else {
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:60
			_go_fuzz_dep_.CoverTab[8850]++
														memoryOperations()
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:61
			// _ = "end of CoverTab[8850]"
		}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:62
		// _ = "end of CoverTab[8847]"
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:62
	}
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:62
	// _ = "end of CoverTab[8844]"
}

//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/rimedo-ts/go-fuzz/go-fuzz-build/build-test/main.go:63
var _ = _go_fuzz_dep_.CoverTab
