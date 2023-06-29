package main

import (
	"context"
	"fmt"
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
	var client = lark.NewClient(appID, appSecret,
		lark.WithLogLevel(larkcore.LogLevelDebug),
		lark.WithReqTimeout(3*time.Second),
		lark.WithEnableTokenCache(true),
		lark.WithHttpClient(http.DefaultClient))

	larkAPI.Subscribe(client, approvalCode)
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	// lambda.Start(Handler)

	str, err := larkAPI.GetTenantAccessToken(appID, appSecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)

}
