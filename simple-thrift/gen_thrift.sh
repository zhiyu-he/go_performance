#!/bin/sh

thrift -out . -r --gen go:thrift_import=github.com/ThoseFlowers/thrift/lib/go/thrift echo.thrift


