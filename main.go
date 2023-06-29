package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"main/larkAPI"
	"main/lib"
	"net/http"
	"os"
	"strconv"
	"time"
)

var approvalCode = os.Getenv("APPROVAL_CODE")

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var appID = os.Getenv("APP_ID")
	var appSecret = os.Getenv("APP_SECRET")

	var client = lark.NewClient(appID, appSecret,
		lark.WithLogLevel(larkcore.LogLevelDebug),
		lark.WithReqTimeout(3*time.Second),
		lark.WithEnableTokenCache(true),
		lark.WithHttpClient(http.DefaultClient))

	instanceID, err := larkAPI.GetInstanceID(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	instanceDetail, err := larkAPI.GetInstanceDetails(appID, appSecret, instanceID, client)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	lib.SendMsg(strconv.Itoa(instanceDetail.StatusCode) + instanceDetail.Msg)

	// return success
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
