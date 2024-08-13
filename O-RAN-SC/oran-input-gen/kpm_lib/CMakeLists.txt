cmake_minimum_required(VERSION 3.10)

project(decode VERSION 1.0)

if(CMAKE_CXX_COMPILER_ID MATCHES "GNU" OR CMAKE_CXX_COMPILER_ID MATCHES "Clang")
  set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wno-error")
  set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wno-error")
endif()

# Set global flags for build of generated asn.1 bindings.
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -DASN_DISABLE_OER_SUPPORT=1")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -DASN_DISABLE_OER_SUPPORT=1")

# Include directories
include_directories(
    /home/tianchang/Desktop/proj/oaic/srsRAN-e2/e2_bindings/E2SM-KPM
    # /home/tianchang/Desktop/proj/oaic/srsRAN-e2/e2_bindings/E2AP-v01.01
    # /usr/local/share/asn1c
)

# Gather all source files from the specified directory
file(GLOB KPM_SOURCES "/home/tianchang/Desktop/proj/oaic/srsRAN-e2/e2_bindings/E2SM-KPM/*.c")
# file(GLOB E2AP_SOURCES "/home/tianchang/Desktop/proj/oaic/srsRAN-e2/e2_bindings/E2AP-v01.01/*.c")

# Compile source files into static libraries
add_library(KPM STATIC ${KPM_SOURCES})
# add_library(E2AP STATIC ${E2AP_SOURCES})
add_library(Mutator STATIC ../cpp/mutator.h ../cpp/mutator.c)

# Link static libraries with the main executable
add_executable(mutator ../cpp/mutator.h ../cpp/mutator.c)
target_link_libraries(mutator KPM)
