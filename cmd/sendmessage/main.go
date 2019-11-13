package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
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

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {

	var postData postData
	err := json.Unmarshal([]byte(request.Body), &postData)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Session Error", StatusCode: 500}, nil
	}
	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("iRepairChatRoom"),
	}

	result, err := svc.Scan(params)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	var config *aws.Config
	apiSession, err := session.NewSession(config)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Session Error", StatusCode: 500}, nil
	}
	apiClient := apigatewaymanagementapi.New(apiSession)
	apiClient.Endpoint = fmt.Sprintf("https://%s/%s", request.RequestContext.DomainName, request.RequestContext.Stage)

	item := Item{}
	for _, i := range result.Items {

		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}
		connectionId := item.ConnectionId

		log.Println(item.ConnectionId)

		message, err := json.Marshal(postData.Data)
		if err != nil {
			log.Println(err.Error())
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}

		_, err = apiClient.PostToConnection(
			&apigatewaymanagementapi.PostToConnectionInput{
				ConnectionId: &connectionId,
				Data:         message,
			})
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		}
	}

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
