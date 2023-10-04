#!/bin/bash
export LANG=zh_CN.UTF-8

dep:
	cd src; go mod tidy; cd -

gen-model:
	gentool -c etc/gen.yml