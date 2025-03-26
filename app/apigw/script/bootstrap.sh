#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=apigw
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}