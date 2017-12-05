#!/bin/sh

thrift -out . --gen go:thrift_import=github.com/apache/thrift/lib/go/thrift echo.thrift


