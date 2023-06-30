package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"log"
	"main/larkAPI"
	"main/lib"
	"net/http"
	"os"
	"time"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var appID = os.Getenv("APP_ID")
	var appSecret = os.Getenv("APP_SECRET")
	var approvalCode = os.Getenv("APPROVAL_CODE")

	var client = lark.NewClient(appID, appSecret,
		lark.WithLogLevel(larkcore.LogLevelDebug),
		lark.WithReqTimeout(3*time.Second),
		lark.WithEnableTokenCache(true),
		lark.WithHttpClient(http.DefaultClient))

	type FeishuEvent struct {
		Type    string          `json:"type"`
		Event   json.RawMessage `json:"event"`
		Token   string          `json:"token"`
		Encrypt bool            `json:"encrypt"`
	}

	var event FeishuEvent
	if err := json.Unmarshal([]byte(request.Body), &event); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}
	log.Println("event.Type: " + event.Type)

	RES, FLAG, err := lib.IsChallenge(ctx, request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// validation request
	if FLAG {
		// subscribe approval by the way
		err := larkAPI.Subscribe(appID, appSecret, approvalCode, client)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}
		return RES, nil
	}

	// if is not the validation request
	instanceID, err := larkAPI.GetInstanceID(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	instanceDetail, err := larkAPI.GetInstanceDetails(appID, appSecret, instanceID, client)
	log.Println(instanceDetail)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	sta, err := lib.GetApprovalStatus(instanceDetail)
	form, err := lib.GetApprovalForms(instanceDetail)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	lib.SendMsg(sta + ": " + form)

	// return success
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
