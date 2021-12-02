package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}

	var err error
	switch request.HTTPMethod {
	case http.MethodGet:
		if userId, ok := request.PathParameters["userId"]; ok {
			resp.Body = fmt.Sprintf("Read User(userId: %s)", userId)
		} else {
			resp.Body = "Read All Users"
		}
	case http.MethodPost:
		resp.Body = "Create User"
	case http.MethodPatch:
		resp.Body = "Update User"
	case http.MethodDelete:
		resp.Body = "Delete User"
	default:
		err = errors.New("invalid HTTP method")
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return resp, nil
}

func main() {
	lambda.Start(handler)
}
