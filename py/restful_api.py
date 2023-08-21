import requests
import json
from time import time

MODEL_NAME = "bert"
HOST = "localhost:8501/"
DATA = json.dumps({"signature_name": "serving_default", "inputs": ["hello world"]})

headers = {"content-type": "application/json"}

try:
    start_time = time()
    json_response = requests.post(HOST + 'v1/models/{}:predict'.format(MODEL_NAME), data=DATA, headers=headers)
    print("predict: ", json_response.json())
    end_time = time()
    seconds_elapsed = end_time - start_time
    print("TIME: {}".format(seconds_elapsed))
except requests.exceptions.HTTPError as err:
    print(err)
