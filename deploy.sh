#!/bin/sh

##### Functions
update_connect_lambda() {
  cd ./cmd/connect && \
  GOOS=linux go build main.go && \
  zip function.zip main
  aws lambda update-function-code --function-name iRepair-connect --zip-file fileb://function.zip && \
  rm -rf ./function.zip && \
  cd ../../
}

update_sendmessage_lambda() {
  cd ./cmd/sendmessage && \
  GOOS=linux go build main.go && \
  zip function.zip main
  aws lambda update-function-code --function-name iRepair-sendmessage --zip-file fileb://function.zip && \
  rm -rf ./function.zip && \
  cd ../../
}

update_disconnect_lambda() {
  cd ./cmd/disconnect && \
  GOOS=linux go build main.go && \
  zip function.zip main
  aws lambda update-function-code --function-name iRepair-disconnect --zip-file fileb://function.zip && \
  rm -rf ./function.zip && \
  cd ../../
}

create_connect_lambda() {
  cd ./cmd/connect && \
  GOOS=linux go build main.go && \
  zip function.zip main
  aws lambda create-function --function-name iRepair-connect --runtime go1.x --zip-file fileb://function.zip --handler main --role arn:aws:iam::378652145250:role/lambda_invoke_function_assume_apigw_role && \
  rm -rf ./function.zip && \
  cd ../../
}

create_disconnect_lambda() {
  cd ./cmd/disconnect && \
  GOOS=linux go build main.go && \
  zip function.zip main
  aws lambda create-function --function-name iRepair-disconnect --runtime go1.x --zip-file fileb://function.zip --handler main --role arn:aws:iam::378652145250:role/lambda_invoke_function_assume_apigw_role && \
  rm -rf ./function.zip && \
  cd ../../
}

create_sendmessage_lambda() {
  cd ./cmd/sendmessage && \
  GOOS=linux go build main.go && \
  zip function.zip main
  aws lambda create-function --function-name iRepair-sendmessage --runtime go1.x --zip-file fileb://function.zip --handler main --role arn:aws:iam::378652145250:role/lambda_invoke_function_assume_apigw_role && \
  rm -rf ./function.zip && \
  cd ../../
}

if [ "$1" = "create" ]; then
    if [ "$2" = "connect" ]; then
      create_connect_lambda
    elif [ "$2" = "disconnect" ]; then
      create_disconnect_lambda
    elif [ "$2" = "sendmessage" ]; then
      create_sendmessage_lambda
    else
      echo "./deploy [create|update] [connect|disconnect|sendmessage]"
    fi
elif [ "$1" = "update" ]; then
    if [ "$2" = "connect" ]; then
      update_connect_lambda
    elif [ "$2" = "disconnect" ]; then
      update_disconnect_lambda
    elif [ "$2" = "sendmessage" ]; then
      update_sendmessage_lambda
    else
      echo "./deploy [create|update] [connect|disconnect|sendmessage]"
    fi
else
    echo "./deploy [create|update] [connect|disconnect|sendmessage]"
fi