//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:1
package main

//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:1
import (
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:1
)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:1
import (
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:1
)

import "fmt"

func testSwitch(x int) {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:5
	_go_fuzz_dep_.CoverTab[9200]++
														switch x {
	case 1:
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:7
		_go_fuzz_dep_.CoverTab[530290]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:7
		_go_fuzz_dep_.CoverTab[9201]++
															fmt.Println("One")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:8
		// _ = "end of CoverTab[9201]"
	case 2:
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:9
		_go_fuzz_dep_.CoverTab[530291]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:9
		_go_fuzz_dep_.CoverTab[9202]++
															fmt.Println("Two")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:10
		// _ = "end of CoverTab[9202]"
	default:
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:11
		_go_fuzz_dep_.CoverTab[530292]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:11
		_go_fuzz_dep_.CoverTab[9203]++
															fmt.Println("Other")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:12
		// _ = "end of CoverTab[9203]"
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:13
	// _ = "end of CoverTab[9200]"
}

func testTypeSwitch(x interface{}) {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:16
	_go_fuzz_dep_.CoverTab[9204]++
														switch x.(type) {
	case int:
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:18
		_go_fuzz_dep_.CoverTab[530293]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:18
		_go_fuzz_dep_.CoverTab[9205]++
															fmt.Println("Integer")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:19
		// _ = "end of CoverTab[9205]"
	case string:
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:20
		_go_fuzz_dep_.CoverTab[530294]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:20
		_go_fuzz_dep_.CoverTab[9206]++
															fmt.Println("String")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:21
		// _ = "end of CoverTab[9206]"
	case bool:
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:22
		_go_fuzz_dep_.CoverTab[530295]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:22
		_go_fuzz_dep_.CoverTab[9207]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:22
		// _ = "end of CoverTab[9207]"
	default:
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:23
		_go_fuzz_dep_.CoverTab[530296]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:23
		_go_fuzz_dep_.CoverTab[9208]++
															fmt.Println("Unknown")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:24
		// _ = "end of CoverTab[9208]"
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:25
	// _ = "end of CoverTab[9204]"
}

func testLoops() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:28
	_go_fuzz_dep_.CoverTab[9209]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:28
	_go_fuzz_dep_.CoverTab[786749] = 0
														for i := 0; i < 5; i++ {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:29
		if _go_fuzz_dep_.CoverTab[786749] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:29
			_go_fuzz_dep_.CoverTab[530303]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:29
		} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:29
			_go_fuzz_dep_.CoverTab[530304]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:29
		}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:29
		_go_fuzz_dep_.CoverTab[786749] = 1
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:29
		_go_fuzz_dep_.CoverTab[9214]++
															fmt.Println("Loop", i)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:30
		// _ = "end of CoverTab[9214]"
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:31
	if _go_fuzz_dep_.CoverTab[786749] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:31
		_go_fuzz_dep_.CoverTab[530305]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:31
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:31
		_go_fuzz_dep_.CoverTab[530306]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:31
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:31
	// _ = "end of CoverTab[9209]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:31
	_go_fuzz_dep_.CoverTab[9210]++

														numbers := []int{1, 2, 3}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:33
	_go_fuzz_dep_.CoverTab[786750] = 0
														for _, num := range numbers {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:34
		if _go_fuzz_dep_.CoverTab[786750] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:34
			_go_fuzz_dep_.CoverTab[530307]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:34
		} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:34
			_go_fuzz_dep_.CoverTab[530308]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:34
		}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:34
		_go_fuzz_dep_.CoverTab[786750] = 1
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:34
		_go_fuzz_dep_.CoverTab[9215]++
															fmt.Println("Range", num)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:35
		// _ = "end of CoverTab[9215]"
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:36
	if _go_fuzz_dep_.CoverTab[786750] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:36
		_go_fuzz_dep_.CoverTab[530309]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:36
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:36
		_go_fuzz_dep_.CoverTab[530310]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:36
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:36
	// _ = "end of CoverTab[9210]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:36
	_go_fuzz_dep_.CoverTab[9211]++

														goto end
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:38
	// _ = "end of CoverTab[9211]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:38
	_go_fuzz_dep_.CoverTab[9212]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:38
	_go_fuzz_dep_.CoverTab[786751] = 0

														for i := 0; i < 5; i++ {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:40
		if _go_fuzz_dep_.CoverTab[786751] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:40
			_go_fuzz_dep_.CoverTab[530311]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:40
		} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:40
			_go_fuzz_dep_.CoverTab[530312]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:40
		}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:40
		_go_fuzz_dep_.CoverTab[786751] = 1
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:40
		_go_fuzz_dep_.CoverTab[9216]++
															fmt.Println("Loop", i)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:41
		// _ = "end of CoverTab[9216]"
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:42
	if _go_fuzz_dep_.CoverTab[786751] == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:42
		_go_fuzz_dep_.CoverTab[530313]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:42
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:42
		_go_fuzz_dep_.CoverTab[530314]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:42
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:42
	// _ = "end of CoverTab[9212]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:42
	_go_fuzz_dep_.CoverTab[9213]++

end:
														;
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:45
	// _ = "end of CoverTab[9213]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:45
}

func testIfElse(x int) {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:47
	_go_fuzz_dep_.CoverTab[9217]++
														if x < 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:48
		_go_fuzz_dep_.CoverTab[530297]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:48
		_go_fuzz_dep_.CoverTab[9218]++
															fmt.Println("Negative")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:49
		// _ = "end of CoverTab[9218]"
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:50
		_go_fuzz_dep_.CoverTab[530298]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:50
		_go_fuzz_dep_.CoverTab[9219]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:50
		if x == 0 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:50
			_go_fuzz_dep_.CoverTab[530299]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:50
			_go_fuzz_dep_.CoverTab[9220]++
																fmt.Println("Zero")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:51
			// _ = "end of CoverTab[9220]"
		} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:52
			_go_fuzz_dep_.CoverTab[530300]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:52
			_go_fuzz_dep_.CoverTab[9221]++
																fmt.Println("Positive")
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:53
			// _ = "end of CoverTab[9221]"
		}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:54
		// _ = "end of CoverTab[9219]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:54
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:54
	// _ = "end of CoverTab[9217]"
}

func testGoto() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:57
	_go_fuzz_dep_.CoverTab[9222]++
														i := 0
loop:
	if i < 5 {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:60
		_go_fuzz_dep_.CoverTab[530301]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:60
		_go_fuzz_dep_.CoverTab[9223]++
															fmt.Println("Goto", i)
															i++
															goto loop
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:63
		// _ = "end of CoverTab[9223]"
	} else {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:64
		_go_fuzz_dep_.CoverTab[530302]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:64
		_go_fuzz_dep_.CoverTab[9224]++
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:64
		// _ = "end of CoverTab[9224]"
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:64
	}
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:64
	// _ = "end of CoverTab[9222]"
}

func main() {
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:67
	_go_fuzz_dep_.InstrumentMain()
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:67
	_go_fuzz_dep_.CoverTab[9225]++
														testSwitch(1)
														testTypeSwitch("hello")
														testLoops()
														testIfElse(10)
														testGoto()
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:72
	// _ = "end of CoverTab[9225]"
}

//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:73
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/Desktop/proj/oran-sc/ric-app-kpimon-go/go-fuzz/go-fuzz-build/build-test/main.go:73
var _ = _go_fuzz_dep_.CoverTab
