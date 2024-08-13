# xApp Testing Procedures

This guide demonstrates an example of testing an xApp. Before start, make sure you follow the [installation guide](../README.md) and have all components built and ready. 

### Deploy RAN Simulator
The `run_report_loop` function in `e2sim/e2sm_examples/kpm_e2sm/src/kpm/kpm_callbacks.cpp` file has many variations, depending on the test input targeting different components. By default, the version with the `// kpm sm` comment on top should be uncommented. Make sure that is the case, and build the RAN Simulator following the [installation guide](../README.md). 

For O-RAN-SC, the RAN node has to be deployed prior to the xApp to successfully establish the connection between the xApp and RAN. 

```
cd e2sim/e2sm_examples/kpm_e2sm
make run
```
The E2T should return a successful outcome message, indicating a successful E2 setup procedure.

To terminate the E2 simulator, you can run (in a separate terminal) ```make stop```

### Onboard and Deploy xApp
Next step is to deploy the test target xApp. Make sure the xApp image with instrumentation is built and ready to use. 
```
cd ric-app-kpimon-go

#  create a local helm repository to host xApp configurations
make helm

# onboard the xApp using the xApp onboarder
make onboard

# deploy xApp to RIC
make install
```

After a successful deployment, you can view the status of the xApp in Kubernetes:
```
# verify the status of the pod
kubectl get po -n ricxapp

# get run-time log of the xApp (replace the pod name)
kubectl logs -n ricxapp ricxapp-kpimon-go-7744d5f76b-45pnh -f
```

In addition, please verify that the RAN simulator shows a succssful subscription to the xApp

### Start Testing
With both the xApp and RAN simulator deployed, we can start the testing process:

#### set up port-forwarding for code coverage feedback
The xApp is already instrumented with code coverage feedback and runtime monitor for collecting the feedback. The main fuzzer needs to access that information. Therefore, we need to open a port (we use 19999) from the Kubernetes deployment so the main fuzzer can read the feedback. On a seperate terminal, run the following command:

```
# set up port-forwarding for code coverage feedback (replace the pod name with the running xApp pod)
kubectl port-forward -n ricxapp ricxapp-kpimon-go-7744d5f76b-8xn7l 19999:19999

```

#### build main fuzzer and start testing
First, make sure the main fuzzer is using the correct mutator:

In `oran-input-gen/go-fuzz/go-fuzz/asn1_mutator/shared_mem.go`, make sure the following line has the kpm mutator enabled (should be enabled by default):
```
// Invoke the subprocess, passing the file descriptors for the pipes.
cmd := exec.Command("mutator/kpm/mutator", sharedFile.Name())
```

The fuzzer needs to use the kpm corpus to mutate on and generate targeted inputs:

```
cd oran-input-gen
rm -r corpus
cp -r corpus_dir/kpm corpus
```

Finally, build and run fuzzer
```
make fuzzer
make run-fuzzer
```

### End Testing
To terminate the testing process, use the following commands:
```
# first, stop the fuzzer process using ctrl+C
# terminate port-forward using ctrl+C

# uninstall the xApp:
cd ric-app-kpimon-go
make uninstall

# uninstall RAN simulator:
cd e2sim/e2sm_examples/kpm_e2sm
make stop
```