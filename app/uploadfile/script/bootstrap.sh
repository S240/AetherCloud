#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/uploadfile"
exec "$CURDIR/bin/uploadfile"
