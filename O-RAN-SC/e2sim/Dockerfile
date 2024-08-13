# FROM ubuntu:22.04

# RUN apt-get update
# RUN apt-get install -y apt-transport-https ca-certificates gnupg software-properties-common wget g++ libsctp-dev
# RUN wget -O - https://apt.kitware.com/keys/kitware-archive-latest.asc 2>/dev/null | apt-key add -
# RUN apt-add-repository 'deb https://apt.kitware.com/ubuntu/ bionic main'
# RUN apt-get update
# RUN apt-get install -y cmake

FROM nexus3.o-ran-sc.org:10002/o-ran-sc/bldr-ubuntu20-c-go:1.0.0 as ubuntu

WORKDIR /e2sim
COPY . /e2sim
RUN rm -rf build
RUN mkdir build
RUN cd build && cmake .. && make package && cmake .. -DDEV_PKG=1 && make package
