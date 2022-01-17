#!/usr/bin/env python

import sys
from argparse import ArgumentParser
import json
from jsonschema import validate
import yaml


class Validator(object):

    def run(self):
        self.buildParser()
        self.buildOptions()
        self.parseOptions()

        fd = open(self.args.schema)
        schema = json.load(fd)

        if self.args.fmt:
            print(json.dumps(schema, indent=2))
            sys.exit(0)

        fd = open(self.args.config)
        config = yaml.safe_load(fd)

        validate(instance=config, schema=schema)

    def buildParser(self):
        self.parser = ArgumentParser(description='Validate YAML file against JSON-schema defintion')

    def buildOptions(self):
        self.parser.add_argument('--schema', default="trapmux.json",
                    help='JSON schema file')
        self.parser.add_argument('--config', default="../tools/trapmux_min.yml",
                    help='YAML configuration file')

        self.parser.add_argument('--fmt', default=False, action='store_true',
                    help='Format schema file and exit')

    def parseOptions(self):
        self.args = self.parser.parse_args()


if __name__ == '__main__':
    cmd = Validator()
    cmd.run()

