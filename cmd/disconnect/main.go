package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/justericgg/irepair/infra/repository/ddb"
	"log"
)

type Event struct {
	RequestContext events.APIGatewayWebsocketProxyRequestContext `json:"requestContext"`
}

type Item struct {
	ConnectionId string `json:"connectionId"`
}

func HandleRequest(ctx context.Context, event Event) (events.APIGatewayProxyResponse, error) {

	err := ddb.Delete(event.RequestContext.ConnectionID)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: "DB error", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
