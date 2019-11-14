package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/justericgg/irepair/infra/adpter/api"
	"github.com/justericgg/irepair/infra/repository/ddb"
	"log"
	"time"
)

type Event struct {
	RequestContext events.APIGatewayWebsocketProxyRequestContext `json:"requestContext"`
	Action         string                                        `json:"action"`
	Data           string                                        `json:"data"`
}

type Item struct {
	ConnectionId string
}

type Payload struct {
	Id      string `json:"id"`
	Author  string `json:"author"`
	Avatar  string `json:"avatar"`
	Message string `json:"message"`
	Images  string `json:"images"`
	Time    int    `json:"time"`
}

type postData struct {
	Data Payload `json:"data"`
}

func getBotMessage(message string) ([]byte, error) {

	var botMessage []byte

	if message == "誰最漂亮" {
		botData := buildBotData("BOT001", "魔鏡", "是9N唷~")
		return json.Marshal(botData)
	}

	if message == "我誰" {
		botData := buildBotData("BOT002", "蛇丸寶貝", "我瘋子~")
		return json.Marshal(botData)
	}

	return botMessage, nil
}

func buildBotData(id, author, message string) Payload {

	return Payload{
		Id:      id,
		Author:  author,
		Avatar:  "avatar-2.png",
		Message: message,
		Images:  "",
		Time:    int(time.Now().Unix()),
	}
}

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	var data postData
	err := json.Unmarshal([]byte(request.Body), &data)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	botMessage, err := getBotMessage(data.Data.Message)

	if err != nil {
		log.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	message, err := json.Marshal(data.Data)
	if err != nil {
		log.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	connectionIds, err := ddb.GetConnections()
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: "DB error", StatusCode: 500}, nil
	}

	endpoint := fmt.Sprintf("https://%s/%s", request.RequestContext.DomainName, request.RequestContext.Stage)
	apiConn, err := api.GetConnection()

	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{Body: "call back API connection error", StatusCode: 500}, nil
	}

	for _, connectionId := range connectionIds {

		connId := string(connectionId)

		_, err := apiConn.Post(endpoint, connId, message)

		if err != nil {
			log.Println(err)
		}

		if len(botMessage) > 0 {

			_, err := apiConn.Post(endpoint, connId, botMessage)

			if err != nil {
				log.Println(err)
			}
		}
	}

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
