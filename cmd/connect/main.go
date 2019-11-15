package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/justericgg/irepair/chat/application/usecase"
	"github.com/justericgg/irepair/chat/infra/repository/ddb"
	"log"
)

type Event struct {
	RequestContext events.APIGatewayWebsocketProxyRequestContext `json:"requestContext"`
}

type Item struct {
	ConnectionId string `json:"connectionId"`
}

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	svc := usecase.NewJoinRoomSvc(&ddb.RoomRepository{})
	err := svc.Join(request.RequestContext.ConnectionID)

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: "error", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
