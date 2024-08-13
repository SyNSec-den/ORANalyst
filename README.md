# ORANalyst

ORANalyst is a systematic testing framework designed for O-RAN's RAN Intelligent Controller (RIC) implementations. For details, please check out our paper, ["ORANalyst: Systematic Testing Framework for Open RAN Implementations"](https://www.usenix.org/conference/usenixsecurity24/presentation/yang-tianchang) (USENIX Security '24).

#### Instructions
You can read on [instructions on testing with ORANalyst](ORANalyst/README.md). For a more detailed example of testing with ORANalyst, you can find more detailed instructions on testing with ORANalyst under each directory:
- [Testing O-RAN-SC](O-RAN-SC/README.md)
- [Testing SD-RAN](SD-RAN/README.md)

#### System
__Ubuntu 18.04.6__ is strongly recommended for the best compatibility deploying O-RAN-SC and SD-RAN. 

The instructions are tested in a system with the following configuration:
```
Memory: 16GB DDR4
CPU: Intel i7-9750H
OS: Ubuntu 18.04.6 LTS
GNOME: 3.28.2
```

#### Software Version
go version go1.20.14 

__Note: latest Go version 1.22.6 is known to be uncamptible with constraint_collector__

__Some README files are still under construction and may miss critical information__