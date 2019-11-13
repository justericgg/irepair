package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Event struct {
	RequestContext events.APIGatewayWebsocketProxyRequestContext `json:"requestContext"`
}

type Item struct {
	ConnectionId string `json:"connectionId"`
}

func HandleRequest(ctx context.Context, event Event) (events.APIGatewayProxyResponse, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Session Error", StatusCode: 500}, nil
	}

	item := Item{ConnectionId: event.RequestContext.ConnectionID}

	ddb := dynamodb.New(sess)
	attr, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      attr,
		TableName: aws.String("iRepairChatRoom"),
	}
	_, err = ddb.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: "ok", StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
