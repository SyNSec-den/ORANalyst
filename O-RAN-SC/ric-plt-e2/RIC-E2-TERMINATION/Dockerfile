##############################################################################
#
#   Copyright (c) 2019 AT&T Intellectual Property.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
##############################################################################

#
#   This source code is part of the near-RT RIC (RAN Intelligent Controller)
#   platform project (RICP).
#
FROM nexus3.o-ran-sc.org:10002/o-ran-sc/bldr-ubuntu20-c-go:1.0.0 as ubuntu

WORKDIR /opt/e2/

ARG BUILD_TYPE="Release"
RUN apt-get update
RUN apt-get install -y lcov
RUN mkdir -p /opt/e2/RIC-E2-TERMINATION/ \
    && mkdir -p /opt/e2/RIC-E2-TERMINATION/TEST/T1 \
    && mkdir -p /opt/e2/RIC-E2-TERMINATION/TEST/T2 \
    && mkdir -p /opt/e2/RIC-E2-TERMINATION/3rdparty


RUN wget https://apt.llvm.org/llvm.sh && \
    chmod +x llvm.sh && \
    ./llvm.sh 10
RUN apt-get update && \
    apt-get install -y clang


RUN git clone https://github.com/AFLplusplus/AFLplusplus.git
RUN cd AFLplusplus && make clean all && \
    make install && cd instrumentation && LLVM_CONFIG=llvm-config-9 make 


COPY . /opt/e2/RIC-E2-TERMINATION/
RUN cd RIC-E2-TERMINATION/go-fuzz/go-fuzz-run-c && go build
RUN cp RIC-E2-TERMINATION/go-fuzz/go-fuzz-run-c/go-fuzz-run /opt/e2/

RUN apt-get install -y libgtest-dev
RUN cd /usr/src/gtest && cmake CMakeLists.txt && make && cp lib/*.a /usr/lib
RUN apt-get install -y google-mock
RUN cd /usr/src/googletest/googlemock  && cmake CMakeLists.txt && make && cp lib/*.a /usr/lib
RUN cp -rf /usr/src/googletest/googlemock/include/gmock /usr/include/
 
RUN apt-get install -y libboost-all-dev
RUN mv /opt/e2/RIC-E2-TERMINATION/CMakeLists.txt /opt/e2/ && cat /opt/e2/RIC-E2-TERMINATION/config/config.conf \
    && wget --content-disposition https://github.com/cgreen-devs/cgreen/releases/download/1.2.0/cgreen_1.2.0_amd64.deb \
    && dpkg -i cgreen_1.2.0_amd64.deb \
    && wget --content-disposition https://packagecloud.io/o-ran-sc/release/packages/debian/stretch/rmr_4.9.4_amd64.deb/download.deb \
    && dpkg -i rmr_4.9.4_amd64.deb \
    && wget --content-disposition https://packagecloud.io/o-ran-sc/release/packages/debian/stretch/rmr-dev_4.9.4_amd64.deb/download.deb \
    && dpkg -i rmr-dev_4.9.4_amd64.deb \
    && wget --content-disposition https://packagecloud.io/o-ran-sc/release/packages/debian/stretch/mdclog_0.1.1-1_amd64.deb/download.deb \
    && dpkg -i mdclog_0.1.1-1_amd64.deb \
    && wget --content-disposition https://packagecloud.io/o-ran-sc/release/packages/debian/stretch/mdclog-dev_0.1.1-1_amd64.deb/download.deb \
    && dpkg -i mdclog-dev_0.1.1-1_amd64.deb \
    && apt-get update && apt-get install -y libcurl4-gnutls-dev gawk libtbb-dev libtbb-doc \
    && apt-get update && apt-get -y install libtbb2 \
    && apt-get install  -y python3 python3-pip python3-setuptools python3-wheel ninja-build \
    && pip3 install  meson \
    && cd /opt/e2/RIC-E2-TERMINATION/3rdparty && git clone https://github.com/pistacheio/pistache.git  && cd pistache && git checkout 0.0.5 \
    && meson setup build  \
       --buildtype=release -DPISTACHE_USE_SSL=false -DPISTACHE_BUILD_EXAMPLES=false  -DPISTACHE_BUILD_TESTS=false \
        -DPISTACHE_BUILD_DOCS=false  --prefix=/usr/local \
    && meson compile -C build  && meson install -C build  && ldconfig
RUN cd /opt/e2/RIC-E2-TERMINATION/3rdparty && git clone -v https://github.com/jupp0r/prometheus-cpp.git \
    && cd prometheus-cpp && git checkout v1.1.0 && git submodule init && git submodule update && mkdir build && cd build \
    && cmake .. -DBUILD_SHARED_LIBS=OFF && make -j 4  && make install && ldconfig \
    && cd /opt/e2/RIC-E2-TERMINATION/3rdparty && git clone https://github.com/jarro2783/cxxopts.git \
    && cd /opt/e2/RIC-E2-TERMINATION/3rdparty && git clone https://github.com/Tencent/rapidjson.git \
    && cd /opt/e2/RIC-E2-TERMINATION/3rdparty && git clone https://github.com/zeux/pugixml.git \
    && cd /opt/e2/ && git clone https://github.com/bilke/cmake-modules.git \
    && cd /opt/e2/ \
    # && cmake -D CMAKE_BUILD_TYPE=$BUILD_TYPE . \
    && cmake -D CMAKE_C_COMPILER=afl-clang -D CMAKE_CXX_COMPILER=afl-clang++ -D CMAKE_BUILD_TYPE=$BUILD_TYPE . \
    && make \
    && echo "3" > /opt/e2/rmr.verbose



RUN if [$BUILD_TYPE == "Debug"] ; then make e2_coverage ; fi 

#    && git clone http://gerrit.o-ran-sc.org/r/ric-plt/tracelibcpp \
#    && cd tracelibcpp && mkdir build && cd build  \
#    &&  sed -i '19iset\(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O3"\)' ../CMakeLists.txt \
#    &&  sed -i '19iset\(CMAKE_CXX_STANDARD 17\)' ../CMakeLists.txt\
#    && /usr/local/bin/cmake .. && make && cd .. && cp include/tracelibcpp/tracelibcpp.hpp . && cd .. \
#    && mkdir nlohmann && cd nlohmann && find / -type f -name "json.hpp" -exec cp {} json.hpp \; && cd .. \
#    && find / -type d -name "opentracing" -exec cp -r {} . \; \
#    && cd /usr/local/lib/ && find / -type f -name "libyaml-cpp.a" -exec cp {} libyaml-cpp.a \; \
#    && find / -type f -name "libopentracing.a" -exec cp {} libopentracing.a \; && cd /opt/e2/RIC-E2-TERMINATION && ls nlohmann  \

FROM ubuntu:20.04


ENV DEBIAN_FRONTEND=noninteractive
RUN echo 'Etc/UTC' > /etc/timezone && \
    ln -s /usr/share/zoneinfo/Etc/UTC /etc/localtime && \
    apt-get update && \
    apt-get install -y tzdata

RUN apt-get update && apt-get install -y net-tools iputils-ping curl tcpdump gdb && mkdir -p /opt/e2/config

COPY --from=ubuntu /opt/e2 /opt/e2
COPY --from=ubuntu /opt/e2/rmr.verbose /tmp/rmr.verbose
COPY --from=ubuntu /opt/e2/RIC-E2-TERMINATION/dockerRouter.txt /opt/e2/dockerRouter.txt
COPY --from=ubuntu /opt/e2/RIC-E2-TERMINATION/config/config.conf /opt/e2/config/config.conf
COPY --from=ubuntu /opt/e2/RIC-E2-TERMINATION/startup.sh /opt/e2/startup.sh
#COPY --from=ubuntu /opt/e2/RIC-E2-TERMINATION/tracelibcpp/build/libtracelibcpp.so /usr/local/lib/libtracelibcpp.so
#COPY --from=ubuntu /opt/e2/RIC-E2-TERMINATION/tracelibcpp/build/libtracelibcpp.so.0 /usr/local/lib/libtracelibcpp.so.0
#COPY --from=ubuntu /opt/e2/RIC-E2-TERMINATION/tracelibcpp/build/libtracelibcpp.so.0.0.2 /usr/local/lib/libtracelibcpp.so.0.0.2


COPY --from=ubuntu /usr/local/lib/librmr_si.so.4.9.4 /usr/local/lib/librmr_si.so.4.9.4
RUN ln -s librmr_si.so.4.9.4  /usr/local/lib/librmr_si.so.4
RUN ln -s librmr_si.so.4 /usr/local/lib/librmr_si.so

COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libmdclog.so /usr/local/lib/libmdclog.so
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libmdclog.so.0 /usr/local/lib/libmdclog.so.0
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libmdclog.so.0.1.7 /usr/local/lib/libmdclog.so.0.1.7

COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libcgreen.so /usr/local/libcgreen.so
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libcgreen.so.1 /usr/local/libcgreen.so.1
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libcgreen.so.1.2.0 /usr/local/libcgreen.so.1.2.0

COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libtbb.so /usr/local/libtbb.so
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libtbb.so.2 /usr/local/libtbb.so.2

COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libtbbmalloc.so /usr/local/libtbbmalloc.so
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libtbbmalloc.so.2 /usr/local/libtbbmalloc.so.2

COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libtbbmalloc_proxy.so /usr/local/libtbbmalloc_proxy.so
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libtbbmalloc_proxy.so.2 /usr/local/libtbbmalloc_proxy.so.2

COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_log_setup.so /usr/lib/libboost_log_setup.so
COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_log_setup.so.1.71.0 /usr/lib/libboost_log_setup.so.1.71.0
#COPY --from=ubuntu /usr/local/lib/libpistache.so  /usr/local/lib/libpistache.so
#COPY --from=ubuntu /usr/local/lib/libpistache.so.0  /usr/local/lib/libpistache.so.0
#COPY --from=ubuntu /usr/local/lib/libpistache.so.0.0.001-git20191031  /usr/local/lib/libpistache.so.0.0.001-git20191031

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ='Asia/Kolkata'
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/* /usr/lib/

COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/ /tmp/lib-files/
COPY move_lib_files.sh /tmp/
RUN chmod +x /tmp/move_lib_files.sh
RUN /tmp/move_lib_files.sh

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_log.so /usr/lib/libboost_log.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_log.so.1.71.0 /usr/lib/libboost_log.so.1.71.0

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_thread.so /usr/lib/libboost_thread.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_thread.so.1.71.0 /usr/lib/libboost_thread.so.1.71.0

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_filesystem.so /usr/lib/libboost_filesystem.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_filesystem.so.1.71.0 /usr/lib/libboost_filesystem.so.1.71.0

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_regex.so /usr/lib/libboost_regex.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_regex.so.1.71.0 /usr/lib/libboost_regex.so.1.71.0

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_system.so /usr/lib/libboost_system.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost_system.so.1.71.0 /usr/lib/libboost_system.so.1.71.0

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost*.so.1.71.0 /usr/lib/
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libboost*.so /usr/lib/

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicui18n.so /usr/lib/libicui18n.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicui18n.so.66.1 /usr/lib/libicui18n.so.66.1
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicui18n.so.66 /usr/lib/libicui18n.so.66

# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicuuc.so /usr/lib/libicuuc.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicuuc.so.66.1 /usr/lib/libicuuc.so.66.1
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicuuc.so.66 /usr/lib/libicuuc.so.66
 
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicudata.so /usr/lib/libicudata.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicudata.so.66.1 /usr/lib/libicudata.so.66.1
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libicudata.so.66 /usr/lib/libicudata.so.66
 
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libstdc++.so.6 /usr/lib/libstdc++.so.6
 
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libcurl-gnutls.so.4.6.0 /usr/lib/libcurl-gnutls.so.4.6.0
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libcurl-gnutls.so.4 /usr/lib/libcurl-gnutls.so.4
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libcurl-gnutls.so /usr/lib/libcurl-gnutls.so
# COPY --from=ubuntu /usr/lib/x86_64-linux-gnu/libsctp.so.1 /usr/lib/libsctp.so.1
COPY --from=ubuntu /usr/local/bin/rmr_probe /opt/e2/rmr_probe

WORKDIR /opt/e2/
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV RMR_SEED_RT=dockerRouter.txt
ENV E2TERM_POD_NAME=e2term
RUN chmod +x /opt/e2/startup.sh /opt/e2/rmr_probe
EXPOSE 38000
EXPOSE 19999
# CMD ["sh", "-c", "./startup.sh"]
CMD ["/bin/bash", "-c", "./go-fuzz-run"]