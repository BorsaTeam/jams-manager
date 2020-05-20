#!/bin/sh

cd /home/application

gotestsum --format=short-verbose --junitfile "$TEST_RESULTS_DIR"/gotestsum-report.xml -- -p 2 -coverprofile=coverage.txt $(go list ./... | grep -v vendor/)

testStatus=$?
if [ $testStatus -ne 0 ]; then
    echo "Tests failed"
    exit 1
fi