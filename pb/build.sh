#!/bin/bash

#编译proto文件生成go文件
protoc.exe --go_out=. *.proto