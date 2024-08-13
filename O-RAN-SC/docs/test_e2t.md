# E2 Termination Testing Procedures

This guide demonstrates steps to test an RIC component using the E2T component in O-RAN-SC's RIC as an example. Before start, make sure you follow the [installation guide](../README.md) and have all components built and ready. 

### Build Modified E2T
Build the instrumented E2T to collect code feedback and communicate with the main fuzzer:
```
cd ric-plt-e2/RIC-E2-TERMINATION
make image
```
Alternatively, you can also find a pre-built image at [ty3gx/ric-plt-e2:latest](https://hub.docker.com/repository/docker/ty3gx/ric-plt-e2/general)

### Build original xApp
When testing the E2T, we can choose to onboard some xApp to make sure the southbound interface of E2T is properly setup. To build the original xApp (not instrumented):
```
cd ric-app-kpimon-go
make build-original
```

### Install RIC with Modified E2T
Now we need to replace the original E2T image with our modified one in deployment. The easiest way is to change the pulled image from deployment script. First if the RIC is still deployed, uninstall it follow the [installation guide](../README.md). Then redeploy the ric using the modified script:

```
./install -f ../RECIPE_EXAMPLE/example_recipe_oran_i_release_fuzz_e2t.yaml 
```

After the deployment complete, you can view the running status of the E2T pod and monitor the logs during testing:
```
# verify the status of the pod
kubectl get po -n ricplt

# get run-time log of the xApp (replace the pod name with the running E2T pod)
kubectl logs -n ricplt deployment-ricplt-e2term-alpha-65b4fd79fb-gngt5 -f
```

### Deploy RAN Simulator
The `run_report_loop` function in `e2sim/e2sm_examples/kpm_e2sm/src/kpm/kpm_callbacks.cpp` file has many variations, depending on the test input targeting different components. By default, the version with the `// kpm sm` comment on top should be uncommented. However, we want to test the E2T, hence comment the current `run_report_loop` function and uncomment the `run_report_loop` function with comment `// e2ap` on top. Build the RAN Simulator following the [installation guide](../README.md). 

For O-RAN-SC, the RAN node has to be deployed prior to the xApp to successfully establish the connection between the xApp and RAN. 

```
cd e2sim/e2sm_examples/kpm_e2sm
make run
```
The E2T should return a successful outcome message, indicating a successful E2 setup procedure.

To terminate the E2 simulator, you can run (in a separate terminal) ```make stop```

### Onboard and Deploy xApp
Next step is to deploy the unmodified xApp to set up the southbound interface of E2T. 
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
Now with the E2T instrumented and properly deployed and the RAN simulator running with the right message type, we can start the testing process:

#### set up port-forwarding for code coverage feedback
The E2T is instrumented with code coverage feedback and runtime monitor for collecting the feedback. The main fuzzer needs to access that information. Therefore, we need to open a port (we use 19999) from the Kubernetes deployment so the main fuzzer can read the feedback. On a seperate terminal, run the following command:

```
# set up port-forwarding for code coverage feedback (replace the pod name with the running E2T pod)
kubectl port-forward -n ricplt deployment-ricplt-e2term-alpha-65b4fd79fb-gngt5 19999:19999
```

#### build main fuzzer and start testing
First, make sure the main fuzzer is using the correct mutator:

In `oran-input-gen/go-fuzz/go-fuzz/asn1_mutator/shared_mem.go`, make sure the following line has the e2ap mutator enabled. By default, the kpm mutator is in use, replace the following line:
```
// Invoke the subprocess, passing the file descriptors for the pipes.
cmd := exec.Command("mutator/kpm/mutator", sharedFile.Name())
```
to use the e2ap mutator:
```
// Invoke the subprocess, passing the file descriptors for the pipes.
cmd := exec.Command("mutator/e2ap/mutator", sharedFile.Name())
```

The fuzzer needs to use the e2ap corpus to mutate on and generate targeted inputs:

```
cd oran-input-gen
rm -r corpus
cp -r corpus_dir/e2ap corpus
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