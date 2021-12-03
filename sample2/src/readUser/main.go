package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userId := request.PathParameters["userId"]
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Read User(userId: %s)", userId),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
