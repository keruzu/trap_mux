#!/bin/bash

FILENAME="${1:-dummy.json}"

URL=http://localhost:8080/write?filename=src/data/$FILENAME
HEADERS='Content-Type: application/json'

curl -X POST $URL -d @$FILENAME -H $HEADERS

# Add a newline to our output
echo

