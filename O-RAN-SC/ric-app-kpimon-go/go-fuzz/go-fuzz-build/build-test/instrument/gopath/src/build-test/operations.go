//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:1
package main

//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:1
import (
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:1
)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:1
import (
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:1
)

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func networkOperation() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:11
	_go_fuzz_dep_.CoverTab[9226]++
															conn, err := net.Dial("tcp", "localhost:8080")
															if err != nil {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:13
		_go_fuzz_dep_.CoverTab[530315]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:13
		_go_fuzz_dep_.CoverTab[9228]++
																fmt.Println("Error:", err)
																return
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:15
		// _ = "end of CoverTab[9228]"
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:16
		_go_fuzz_dep_.CoverTab[530316]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:16
		_go_fuzz_dep_.CoverTab[9229]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:16
		// _ = "end of CoverTab[9229]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:16
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:16
	// _ = "end of CoverTab[9226]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:16
	_go_fuzz_dep_.CoverTab[9227]++
															defer conn.Close()

//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:20
	fmt.Fprintf(conn, "Hello, Server!\n")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:20
	// _ = "end of CoverTab[9227]"
}

func fileOperations(filename string) {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:23
	_go_fuzz_dep_.CoverTab[9230]++
															file, err := os.Create(filename)
															if err != nil {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:25
		_go_fuzz_dep_.CoverTab[530317]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:25
		_go_fuzz_dep_.CoverTab[9234]++
																panic(err)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:26
		// _ = "end of CoverTab[9234]"
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:27
		_go_fuzz_dep_.CoverTab[530318]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:27
		_go_fuzz_dep_.CoverTab[9235]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:27
		// _ = "end of CoverTab[9235]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:27
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:27
	// _ = "end of CoverTab[9230]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:27
	_go_fuzz_dep_.CoverTab[9231]++
															defer file.Close()

															writer := bufio.NewWriter(file)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:30
	_go_fuzz_dep_.CoverTab[786752] = 0
															for i := 0; i < 5; i++ {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:31
		if _go_fuzz_dep_.CoverTab[786752] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:31
			_go_fuzz_dep_.CoverTab[530327]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:31
		} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:31
			_go_fuzz_dep_.CoverTab[530328]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:31
		}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:31
		_go_fuzz_dep_.CoverTab[786752] = 1
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:31
		_go_fuzz_dep_.CoverTab[9236]++
																writer.WriteString(fmt.Sprintf("Writing line %d\n", i))
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:32
		// _ = "end of CoverTab[9236]"
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:33
	if _go_fuzz_dep_.CoverTab[786752] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:33
		_go_fuzz_dep_.CoverTab[530329]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:33
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:33
		_go_fuzz_dep_.CoverTab[530330]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:33
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:33
	// _ = "end of CoverTab[9231]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:33
	_go_fuzz_dep_.CoverTab[9232]++
															writer.Flush()

															data, err := ioutil.ReadFile(filename)
															if err != nil {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:37
		_go_fuzz_dep_.CoverTab[530319]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:37
		_go_fuzz_dep_.CoverTab[9237]++
																panic(err)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:38
		// _ = "end of CoverTab[9237]"
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:39
		_go_fuzz_dep_.CoverTab[530320]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:39
		_go_fuzz_dep_.CoverTab[9238]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:39
		// _ = "end of CoverTab[9238]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:39
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:39
	// _ = "end of CoverTab[9232]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:39
	_go_fuzz_dep_.CoverTab[9233]++
															fmt.Println(string(data))
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:40
	// _ = "end of CoverTab[9233]"
}

func memoryOperations() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:43
	_go_fuzz_dep_.CoverTab[9239]++
															data := make(map[string][]int)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:44
	_go_fuzz_dep_.CoverTab[786753] = 0
															for i := 0; i < 5; i++ {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:45
		if _go_fuzz_dep_.CoverTab[786753] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:45
			_go_fuzz_dep_.CoverTab[530331]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:45
		} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:45
			_go_fuzz_dep_.CoverTab[530332]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:45
		}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:45
		_go_fuzz_dep_.CoverTab[786753] = 1
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:45
		_go_fuzz_dep_.CoverTab[9241]++
																key := fmt.Sprintf("key_%d", i)
																data[key] = append(data[key], i, i+1, i+2)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:47
		// _ = "end of CoverTab[9241]"
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:48
	if _go_fuzz_dep_.CoverTab[786753] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:48
		_go_fuzz_dep_.CoverTab[530333]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:48
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:48
		_go_fuzz_dep_.CoverTab[530334]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:48
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:48
	// _ = "end of CoverTab[9239]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:48
	_go_fuzz_dep_.CoverTab[9240]++

															if v, exists := data["key_2"]; exists {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:50
		_go_fuzz_dep_.CoverTab[530321]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:50
		_go_fuzz_dep_.CoverTab[9242]++
																fmt.Println("Found:", v)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:51
		// _ = "end of CoverTab[9242]"
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:52
		_go_fuzz_dep_.CoverTab[530322]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:52
		_go_fuzz_dep_.CoverTab[9243]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:52
		// _ = "end of CoverTab[9243]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:52
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:52
	// _ = "end of CoverTab[9240]"
}

func operations() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:55
	_go_fuzz_dep_.CoverTab[9244]++
															if len(os.Args) > 1 && func() bool {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:56
		_go_fuzz_dep_.CoverTab[9245]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:56
		return os.Args[1] == "net"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:56
		// _ = "end of CoverTab[9245]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:56
	}() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:56
		_go_fuzz_dep_.CoverTab[530323]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:56
		_go_fuzz_dep_.CoverTab[9246]++
																networkOperation()
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:57
		// _ = "end of CoverTab[9246]"
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
		_go_fuzz_dep_.CoverTab[530324]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
		_go_fuzz_dep_.CoverTab[9247]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
		if len(os.Args) > 1 && func() bool {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
			_go_fuzz_dep_.CoverTab[9248]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
			return os.Args[1] == "file"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
			// _ = "end of CoverTab[9248]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
		}() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
			_go_fuzz_dep_.CoverTab[530325]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:58
			_go_fuzz_dep_.CoverTab[9249]++
																	fileOperations("test.txt")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:59
			// _ = "end of CoverTab[9249]"
		} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:60
			_go_fuzz_dep_.CoverTab[530326]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:60
			_go_fuzz_dep_.CoverTab[9250]++
																	memoryOperations()
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:61
			// _ = "end of CoverTab[9250]"
		}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:62
		// _ = "end of CoverTab[9247]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:62
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:62
	// _ = "end of CoverTab[9244]"
}

//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:63
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/operations.go:63
var _ = _go_fuzz_dep_.CoverTab
