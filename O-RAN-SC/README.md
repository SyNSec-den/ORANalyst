# O-RAN-SC Testing
The following guide is tested on Ubuntu 18.04.6. __Super user permissions are assumed for all commands.__


### O-RAN-SC Installation and Deployment
This part follows [O-RAN-SC's I-release installation guide](https://docs.o-ran-sc.org/projects/o-ran-sc-ric-plt-ric-dep/en/latest/installation-guides.html#installing-near-realtime-ric-in-ric-cluster) with minimal modification. Please refer to that guide for a more detailed explanation of the installation and deployment process. 

```
cd ric-dep/bin
./install_k8s_and_helm.sh
./install_common_templates_to_helm.sh
./install -f ../RECIPE_EXAMPLE/example_recipe_oran_i_release.yaml
```
Make sure that pods in ricplt namespace are running normally: ```kubectl get po -A```

You can view the logs of a component using command like ```kubectl logs -n ricplt deployment-ricplt-e2term-alpha-6c69fb8c5f-hcnjt -f``` (replace namespace and pod name accordingly). 

### Uninstall O-RAN-SC

To uninstall the Kubernetes deployment of O-RAN-SC
```
cd ric-dep/bin
./uninstall
```

### Build the xApp Image for Testing
You can test ORANalyst by fuzzing against a sample xApp, provided in ric-app-kpimon-go folder. The xApp can be built as a docker image following the command in Makefile and the build step in Dockerfile. The instrumentation on the xApp is already applied.

```
# cd to the xApp folder
cd ric-app-kpimon-go

# deploy a local docker registry to host docker images
sudo docker run -d -p 5001:5000 --restart=always --name ric registry:2

# build the xApp docker image using the Dockerfile and push to the local docker registry
make build
```

Alternatively, you can find a pre-built xApp docker image at [ty3gx/kpimon-go:latest](https://hub.docker.com/repository/docker/ty3gx/kpimon-go/general)


### xApp Onboarder Installation
xApps can be onboarded using ```dms_cli``` tool:
```
# clone appmgr
git clone "https://gerrit.o-ran-sc.org/r/ric-plt/appmgr"
cd appmgr/xapp_orchestrater/dev/xapp_onboarder

# if pip3 is not installed, install using the following command
yum install python3-pip

# in case dms_cli binary is already installed, it can be uninstalled using following command
pip3 uninstall xapp_onboarder

# install xapp_onboarder using following command
pip3 install ./
```

### RAN Simulator Installation
The RAN simulator is responsible for managing the E2 interface connection with the RIC, and send the test inputs to the RIC. The provided RAN simulator has the KPM service model installed.

```
# cd to the e2sim directory
cd e2sim

# build e2sim with kpm service model
cd e2sm_examples/kpm_e2sm
make build
```

Alternatively, you can find a pre-built docker image at [ty3gx/e2sim:latest](https://hub.docker.com/repository/docker/ty3gx/e2sim/general)

### Building the Main Fuzzer
The main fuzzer uses functionality generated by ASN.1C for message mutation. Therefore, the Go code implementing the fuzzer needs to be compiled with the C-implemented mutator. The mutator is in `oran-input-gen/kpm`. A pre-built library file is in `oran-input-gen/kpm/build`. Depending on your system, you may need to recompile this library file using the provided CMake. 

In addition, in `oran-input-gen` directory, you need to replace the CGo flags specifying the header and library files for the mutator to their correct absolute paths. Replace the paths in `#cgo CFLAGS` and `#cgo LDFLAGS` in the Go files to the correct path to the mutator.
```
// #cgo CFLAGS: -I/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm
// #cgo LDFLAGS: -L/home/tianchang/Desktop/proj/oran-sc/oran-input-gen/kpm/build -lkpm -lm
```

After the mutator is correctly linked, you can compile the main fuzzer:

```
cd oran-input-gen
make fuzzer
```




## Testing Procedures
After building the corresponding components, you can follow the testing procedures:

- [xApp Testing Example](docs/test_xApp.md)
- [RIC Component (E2T) Testing Example](docs/test_e2t.md)
