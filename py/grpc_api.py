# apt-get update && apt install python3-pip -y
# pip install grpcio tensorflow tensorflow_serving_api

import grpc
import tensorflow as tf
from tensorflow_serving.apis import predict_pb2
from tensorflow_serving.apis import prediction_service_pb2_grpc


MODEL_NAME = "bert"
HOST = "localhost:8500"
MODEL_SIGNATURE_NAME = "serving_default"
TIMEOUT = 5.0

channel = grpc.insecure_channel(HOST)
stub = prediction_service_pb2_grpc.PredictionServiceStub(channel)
request = predict_pb2.PredictRequest()
request.model_spec.name = MODEL_NAME
request.model_spec.signature_name = MODEL_SIGNATURE_NAME
request.inputs['input_1'].CopyFrom(tf.make_tensor_proto(["hello world"]))
result_future = stub.Predict.future(request, TIMEOUT)  
result = result_future.result()
print(result)
