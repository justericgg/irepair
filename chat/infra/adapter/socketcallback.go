package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type Connection struct {
	Conn *apigatewaymanagementapi.ApiGatewayManagementApi
}

func GetConnection() (*Connection, error) {

	var config *aws.Config
	apiSession, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	conn := apigatewaymanagementapi.New(apiSession)

	return &Connection{Conn: conn}, nil
}

func (apiConn *Connection) Post(endpoint string, connectionId string, data []byte) (*apigatewaymanagementapi.PostToConnectionOutput, error) {

	apiConn.Conn.Endpoint = endpoint

	res, err := apiConn.Conn.PostToConnection(
		&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: &connectionId,
			Data:         data,
		})

	if err != nil {
		return nil, err
	}

	return res, nil
}
