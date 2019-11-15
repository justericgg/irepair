package ddb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/justericgg/irepair/chat/domain/model/room"
)

const tableName = "iRepairChatRoom"

type Item struct {
	ConnectionId string `json:"connectionId"`
}

func connect() (*dynamodb.DynamoDB, error) {
	dynamodbSession, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		return nil, err
	}
	db := dynamodb.New(dynamodbSession)

	return db, nil
}

type RoomRepository struct{}

func (repo *RoomRepository) Save(theRoom room.Room) error {

	db, err := connect()
	if err != nil {
		return err
	}

	item := Item{ConnectionId: theRoom.GetFirstUserConnId()}
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

func (repo *RoomRepository) Delete(theRoom room.Room) error {

	db, err := connect()
	if err != nil {
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"connectionId": {
				S: aws.String(theRoom.GetFirstUserConnId()),
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
