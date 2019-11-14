package ddb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "iRepairChatRoom"

type ConnectionId string

type Item struct {
	ConnectionId string `json:"connectionId"`
}

func connect() (*dynamodb.DynamoDB, error) {
	dynamodbSession, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		return nil, err
	}
	ddb := dynamodb.New(dynamodbSession)

	return ddb, nil
}

func Put(connectionId string) error {

	db, err := connect()
	if err != nil {
		return err
	}

	item := Item{ConnectionId: connectionId}
	attr, err := dynamodbattribute.MarshalMap(item)
	input := &dynamodb.PutItemInput{
		Item:      attr,
		TableName: aws.String(tableName),
	}
	_, err = db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func Delete(connectionId string) error {
	db, err := connect()
	if err != nil {
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"connectionId": {
				S: aws.String(connectionId),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err = db.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}

func GetConnections() ([]ConnectionId, error) {

	db, err := connect()
	if err != nil {
		return nil, err
	}

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := db.Scan(params)
	if err != nil {
		return nil, err
	}

	item := Item{}
	var connectionIds []ConnectionId
	for _, i := range result.Items {
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return nil, err
		}
		connectionIds = append(connectionIds, ConnectionId(item.ConnectionId))
	}

	return connectionIds, nil
}
