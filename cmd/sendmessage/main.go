package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/justericgg/irepair/chat/application/usecase"
	"github.com/justericgg/irepair/chat/infra/repository/ddb"
	"github.com/justericgg/irepair/chat/infra/service"
	"log"
)

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	endpoint := fmt.Sprintf("https://%s/%s", request.RequestContext.DomainName, request.RequestContext.Stage)

	log.Printf("%+v\n", request)

	roomRepo := &ddb.RoomRepository{}
	broadcastSvc := service.BroadcastSvc{}
	svc := usecase.NewMessageSvc(roomRepo, broadcastSvc)
	err := svc.ProcessMessage(endpoint, request.Body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: "error", StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
