# SD-RAN Testing
The following guide is tested on Ubuntu 18.04.6. __Super user permissions are assumed for all commands.__


### SD-RAN Installation and Deployment
This part follows [SD-RAN's Installation Guide](https://docs.sd-ran.org/master/sdran-in-a-box/README.html) with minimal modification. Please refer to that guide for a more detailed explanation of the installation and deployment process. 

```
cd sdran-in-a-box
make riab opt=ransim ver=stable
```
Wait until all kubernetes pods conditions are met and the script outputs Done

Make sure that pods in riab namespace are running normally: ```kubectl get po -A```

You can view the logs of a component using command like ```kubectl logs -n riab onos-e2t-5798f554b7-2827r -f``` (replace namespace and pod name accordingly). 

### Uninstall SD-RAN

To uninstall the Kubernetes deployment of O-RAN-SC
```
cd sdran-in-a-box
make clean
```

### Deploy the Local Docker Registry
After the RIC installation, we deploy a local docker registry to host our modified docker images.
```
sudo docker run -d -p 5001:5000 --restart=always --name ric registry:2
```

### Build Component Images for Testing
Below we use a sample xApp as testing target and demonstrate the steps to test it using ORANalyst. You can optionally test other components using a similar method. 

All components are built as docker images and managed by kubernetes.

#### Build Main Fuzzer
```
cd input-generator
make images
make docker-push
```
Optionally, a pre-built image can be found at [ty3gx/input-generator](https://hub.docker.com/repository/docker/ty3gx/input-generator/general)

#### Build Main Fuzzer
```
cd input-generator
make images
make docker-push
```
Alternatively, a pre-built image can be found at [ty3gx/input-generator](https://hub.docker.com/repository/docker/ty3gx/input-generator/general)


#### Build RAN Simulator
```
cd ran-simulator
make images
make push-ransim
```
Alternatively, a pre-built image can be found at [ty3gx/ran-simulator](https://hub.docker.com/repository/docker/ty3gx/ran-simulator/general)

#### Build Sample xApp
An xApp (rimedo-ts) is instrumented and provided as sample.
```
cd rimedo-ts
make images
make docker-push-latest
```
Alternatively, a pre-built image can be found at [ty3gx/rimedo-ts](https://hub.docker.com/repository/docker/ty3gx/rimedo-ts/general)



### Deploy and Start Testing
With all the images built, we can start deploying the xApp and the fuzzer to start the fuzzing process. The deployment configuration is specified in `sdran-helm-charts` directory under each component. Specifically, the xApp is going to be deployed as a pod communicating with other RIC components. The fuzzer and RAN simulator is going to be deployed in a single pod running on two seperate containers. The fuzzer generates test inputs and send to RAN simulator, and RAN simulator forwards the test to the RIC, gets routed by E2T to eventually reach the target. The code coverage is then collected and directed back to the fuzzer for feedback. 

#### Deploy xApp
```
cd rimedo-ts
make install-xapp
```
To view the deployment status of the xApp, run `kubectl get po -n riab` and find the pod name starting with `rimedo-ts`. To view the runtime logs, `kubectl logs -n riab rimedo-ts-77fc59d77-td99b -c rimedo-ts -f` (replace the pod name to your corresponding running pod)

#### Deploy Fuzzer and RAN Simulator
```
cd ran-simulator
make install-ransim
```
To view the deployment status of the fuzer and ran simulator, run `kubectl get po -n riab` and find the pod name starting with `ran-simulator`. To view the runtime logs of the ran-simulator, `kubectl logs -n riab ran-simulator-79b46bcd5-f5mn8 -c ran-simulator -f` (replace the pod name to your corresponding running pod). To view the runtime logs of the main fuzzer, replace the container name to `input-generator`, i.e., `kubectl logs -n riab ran-simulator-79b46bcd5-f5mn8 -c input-generator -f`

Once both pods are deployed, you should se the fuzzing inputs reaching the xApp and the testing process starting.

#### Stopping the Testing
To stop the testing, you need to remove the deployment of the xApp, fuzzer, and ran simulator:

To remove xApp:
```
cd rimedo-ts
make delete-xapp
```

To remove fuzzer and ran simulator:
```
cd ran-simulator
make delete-ransim
```


### Testing Other Targets
The main fuzzer (input-generator) provides testing of other service models and messages. Replace the mutator in `main.go` to other mutators and replace the corpus folder to other messages and then rebuild the input-generator image. 

