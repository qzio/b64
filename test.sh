#!/bin/sh
set -eu

go build

foobarb64=$(echo 'foobar' | ./b64)
foobar=$(echo 'foobar' | base64)

if [ $foobarb64 != $foobar ]; then
  echo "failed to encode foobar"
  exit 1
fi

foobarb64Decode=$(echo $foobarb64 | ./b64 -d)
foobarDecode=$(echo $foobar | base64 -d)

if [ $foobarb64Decode != $foobarDecode ]; then
  echo "failed to decode foobar"
  exit 1
fi

echo "passed"
