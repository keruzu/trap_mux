#!/usr/bin/env bash

for cmd in `ls -1 | grep -v .sh | grep -v Makefile`; do
    echo "===   $cmd  =================================="
    ( cd $cmd ; make )
done

