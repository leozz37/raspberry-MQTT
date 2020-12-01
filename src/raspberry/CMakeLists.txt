cmake_minimum_required(VERSION 3.13)
project(raspberry)

set(CMAKE_CXX_STANDARD 17)

add_executable(raspberry main.cpp)

find_library(WIRINGPI_LIBRARIES NAMES wiringPi)
target_link_libraries(raspberry ${WIRINGPI_LIBRARIES})