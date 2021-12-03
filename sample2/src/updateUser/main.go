package updateUser

import (
	"github.com/aws/aws-lambda-go/events"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Update User",
		StatusCode: 200,
	}, nil
}

// func main() {
// 	lambda.Start(handler)
// }
