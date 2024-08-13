# ORANalyst's Main Fuzzer

This contains the main fuzzer component of ORANalyst, and is built on top of [go-fuzz](https://github.com/dvyukov/go-fuzz.git)

### Instrumentation

To instrument a program, follow the example in [ric-app-kpimon-go](../ric-app-kpimon-go). First, you need to instrument the main input receiving loop using `gofuzzdep.LoopPos()` function. This is a dummy function that tells the instrumentor when a new input is ready to be received, an information critical for knowing the end-of-processing of the previous input. See an example of this instrumentation for [the kpimon xApp](../ric-app-kpimon-go/control/control_loop.go).

In addition, you need to include this module in your `go.mod` file for the component. [See an example](../ric-app-kpimon-go/go.mod)

### Running the fuzzer
The instrumentation will provide an instrumented binary with the metadata collected during the information for fuzz input generation. The main fuzzer reads the metadata to generate targeted test input. Structure-aware mutator is invoked in the [mutator wrapper](go-fuzz/go-fuzz/asn1_mutator/mutator.go). [go-fuzz-run](go-fuzz/go-fuzz-run) is responsible for running and monitoring the instrumented binary, maintaining the code coverage information, and communicating with the main fuzzer (default port is 19999). The fuzzer side has a client for communicating with the remote monitor and receiving feedback. [See an example here](main.go). At the fuzzer side, the fuzzer utilizes this client to communicate with the remote target and generate test inputs. See the test example for an end-to-end set-up of this fuzzing framework. 