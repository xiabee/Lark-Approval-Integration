package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"log"
	"main/larkAPI"
	"net/http"
	"os"
	"time"
)

var appID = os.Getenv("APP_ID")
var appSecret = os.Getenv("APP_SECRET")
var approvalCode = os.Getenv("APPROVAL_CODE")

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	// lambda.Start(Handler)
	var client = lark.NewClient(appID, appSecret,
		lark.WithLogLevel(larkcore.LogLevelDebug),
		lark.WithReqTimeout(3*time.Second),
		lark.WithEnableTokenCache(true),
		lark.WithHttpClient(http.DefaultClient))

	err := larkAPI.Subscribe(appID, appSecret, approvalCode, client)
	if err != nil {
		log.Println(err)
	}

}
