
# import logging
# from logging.config import fileConfig
import re
from datetime import datetime

from flask import Flask, request, abort, render_template, jsonify, Response
from flask_cors import CORS


# -- Flask setup  --------
CONTAINER_ID = open('/etc/hostname').read().strip()
# VERSION = open('/app/version').read().strip()
VERSION = "0.0.1"

app = Flask(__name__)
cors = CORS(app)

birth_time = datetime.now()

# -- Routes --------
@app.route('/')
@app.route('/index')
def index():
    return render_template('index.html', version=VERSION)

def getFile(source, path):
    sane = pathClean(source)
    if sane is None:
        return abort(400, "Unable to parse source data name: {source}")
    filename = f'src/{path}/{sane}.json'
    try:
        with open(filename, 'r') as fd:
            data = fd.read()
    except Exception as ex:
        return abort(400, "Unable to read JSON configs: %s" % ex)

    resp = Response(data, mimetype='application/json')
    return resp

def pathClean(base):
    try:
        filename = base.rsplit('/', 1)[-1]
        return re.sub('[^\w]', '', filename)
    except Exception:
        pass

@app.route('/schema/<schema>', methods=['GET'])
def getSchema(schema):
    return getFile(schema, path='schema')

@app.route('/load/<config>', methods=['GET'])
def getConfig(config):
    return getFile(config, path='data')

@app.route('/save/<config>', methods=['POST'])
def saveConfig(config):
    try:
        raw_data = request.json
    except Exception as ex:
        return abort(400, f"Corrupted JSON message: {ex}")

    if not raw_data:
        return abort(400, "Empty JSON message")

    saneConfig = pathClean(config)
    filename = f'src/data/{saneConfig}.json'
    try:
        configData = request.data.decode('utf-8')
        with open(filename, 'w') as fd:
            fd.write(configData)
        status = {config: config, status: "Saved file"}
    except Exception as ex:
        return abort(500, f"Unable to save JSON config {filename}: {ex}")
    return jsonify(status)

@app.route('/liveness', methods=['GET'])
def lineness():
    uptime = datetime.now() - birth_time
    data = jsonify(dict(uptime=uptime))
    mystatus = 200 # Use 5xxx if we have issues
    response = Response(data, status=mystatus, mimetype='application/json')
    return response

@app.route('/readiness', methods=['GET'])
def readiness():
    uptime = datetime.now() - birth_time
    data = jsonify(dict(uptime=uptime))
    mystatus = 200 # Use 5xxx if we have issues
    response = Response(data, status=mystatus, mimetype='application/json')
    return response

@app.route('/health', methods=['GET'])
def health():
    uptime = datetime.now() - birth_time
    return render_template('health.html', uptime=uptime)

@app.route('/version', methods=['GET'])
def show_version():
    return VERSION


if __name__ == '__main__':
    cors.run(host='0.0.0.0', port=8080)

