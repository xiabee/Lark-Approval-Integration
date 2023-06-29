package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"main/lib"
	"os"
)

var appID = os.Getenv("APP_ID")
var appSecret = os.Getenv("APP_SECRET")
var approvalCode = os.Getenv("APPROVAL_CODE")

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var eventData struct {
	//	Event struct {
	//		InstanceID string `json:"instance_id"`
	//		// 其他审批事件字段...
	//	} `json:"event"`
	//	// 其他字段...
	//}
	//if err := json.Unmarshal([]byte(request.Body), &eventData); err != nil {
	//	log.Printf("Failed to unmarshal event data: %s", err)
	//	return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	//}
	//
	//// 打印审批事件的InstanceID
	//fmt.Println("InstanceID:", eventData.Event.InstanceID)
	//SendMsg("InstanceID:" + eventData.Event.InstanceID)

	instanceID, err := lib.GetInstanceID(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}
	lib.SendMsg(instanceID)
	// 返回成功响应
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
	//var client = lark.NewClient(appID, appSecret,
	//	lark.WithLogLevel(larkcore.LogLevelDebug),
	//	lark.WithReqTimeout(3*time.Second),
	//	lark.WithEnableTokenCache(true),
	//	lark.WithHttpClient(http.DefaultClient))
	//
	//err := larkAPI.Subscribe(appID, appSecret, approvalCode, client)
	//if err != nil {
	//	log.Println(err)
	//}

}
