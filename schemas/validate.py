#!/usr/bin/env python

import json
from jsonschema import validate
import yaml


schema_file = "trapmux.json"
config_file = "../tools/trapmux_min.yml"

fd = open(schema_file)
schema = json.load(fd)
print(json.dumps(schema, indent=4))

fd = open(config_file)
data = yaml.safe_load(fd)

validate(instance=data, schema=schema)

