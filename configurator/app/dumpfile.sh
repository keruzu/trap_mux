#!/bin/bash

FILENAME="${1:-dummy.json}"

CONFIG=`echo $FILENAME | sed -e 's/\.json//'`
URL=http://localhost:8080/save/$CONFIG

curl $URL -d @$FILENAME -H 'Content-Type: application/json'

# Add a newline to our output
echo

