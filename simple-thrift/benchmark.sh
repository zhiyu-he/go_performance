#!/bin/sh

thrift -out . --gen go:thrift_import=github.com/apache/thrift/lib/go/thrift,package_prefix=github.com/zhiyu-he/go_performance/simple-thrift/ echo.thrift
go test -v -bench=. benchmark_test.go -benchmem -cpuprofile profile.out

