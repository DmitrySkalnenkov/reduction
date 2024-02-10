#!/bin/bash
shortenertest -test.v -test.run=^TestIteration[1-7]$ -binary-path=cmd/shortener/shortener -server-port=8080 -file-storage-path=/tmp/temp.txt -source-path=.