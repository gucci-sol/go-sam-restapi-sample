package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gucci-sol/go-sam-restapi-sample/sample2/updateUser"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	updateUser.Handler()
	return events.APIGatewayProxyResponse{
		Body:       "Create User",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(updateUser.Handler())
}
