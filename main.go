package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	tf_framework "grpc-client/internal/tensorflow/tensorflow/tensorflow/core/framework"
	tfs_api_pb "grpc-client/internal/tensorflow/tensorflow/tensorflow_serving/apis"

	"google.golang.org/grpc"
)

const (
	HOST                 = "localhost:8500"
	MODEL_NAME           = "bert"
	MODEL_SIGNATURE_NAME = "serving_default"
)

func main() {
	conn, err := grpc.Dial(HOST, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := tfs_api_pb.NewPredictionServiceClient(conn)
	// channel.unary_unary(
	// '/tensorflow.serving.PredictionService/Predict',
	// request_serializer=tensorflow__serving_dot_apis_dot_predict__pb2.PredictRequest.SerializeToString,
	// response_deserializer=tensorflow__serving_dot_apis_dot_predict__pb2.PredictResponse.FromString,
	// )
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	predictRequest := tfs_api_pb.PredictRequest{
		ModelSpec: &tfs_api_pb.ModelSpec{
			Name:          MODEL_NAME,
			SignatureName: MODEL_SIGNATURE_NAME,
		},
		Inputs: map[string]*tf_framework.TensorProto{
			"input_1": &tf_framework.TensorProto{
				Dtype: tf_framework.DataType_DT_STRING,
				TensorShape: &tf_framework.TensorShapeProto{
					Dim: []*tf_framework.TensorShapeProto_Dim{
						{
							Size: 1,
						},
						{
							Size: 1,
						},
					},
				},
				StringVal: [][]byte{[]byte("hello world")},
			},
		},
	}

	for key, value := range predictRequest.Inputs {
		log.Printf("Input %s", key)

		for _, element := range value.FloatVal {
			log.Printf("\t%f", element)
		}
	}

	predictResponse, err := c.Predict(ctx, &predictRequest)
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}

	jsonResponse, err := json.Marshal(predictResponse)
	if err != nil {
		log.Fatalf("could not marshal: %v", err)
	}
	log.Printf("Response: %s", jsonResponse)

	for key, value := range predictResponse.Outputs {
		log.Printf("Output %s", key)

		for _, element := range value.FloatVal {
			log.Printf("\t%f", element)
		}
	}
}
