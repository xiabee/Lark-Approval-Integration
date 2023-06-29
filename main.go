package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

type FeishuEvent struct {
	Type    string          `json:"type"`
	Event   json.RawMessage `json:"event"`
	Token   string          `json:"token"`
	Encrypt bool            `json:"encrypt"`
}

type ApprovalEvent struct {
	ApprovalCode string `json:"approval_code"`
	EventID      string `json:"event_id"`
	UserID       string `json:"user_id"`
	OpenID       string `json:"open_id"`
	Type         string `json:"type"`
	// 其他审批事件字段
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	logger.Println("aaaaaaaaaaaa")
	logger.Printf("Request: %v", request)

	// 解析飞书事件
	var event FeishuEvent
	if err := json.Unmarshal([]byte(request.Body), &event); err != nil {
		logger.Printf("Failed to parse Feishu event: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// 检查是否是验证请求
	if event.Type == "url_verification" {
		SendMsg("url_verification")
		// 提取 CHALLENGE 参数
		var verificationEvent struct {
			Challenge string `json:"challenge"`
		}
		if err := json.Unmarshal([]byte(request.Body), &verificationEvent); err != nil {
			logger.Printf("Failed to parse verification event: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}

		// 构建响应对象
		responseBody := struct {
			Challenge string `json:"challenge"`
		}{
			Challenge: verificationEvent.Challenge,
		}

		// 返回验证响应
		responseJSON, err := json.Marshal(responseBody)
		if err != nil {
			logger.Printf("Failed to marshal response body: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500}, nil
		}

		// 构建合法的JSON响应字符串
		response := events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(responseJSON),
		}

		// 返回响应
		return response, nil
	}

	if event.Type == "event_type_approval" {
		// 解析审批事件
		SendMsg("event_type_approval")
		var approvalEvent ApprovalEvent
		if err := json.Unmarshal(event.Event, &approvalEvent); err != nil {
			logger.Println("Failed to parse approval event: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}
		SendMsg("Approval Code: " + approvalEvent.ApprovalCode)
		// 处理审批事件
		logger.Println("Handling Approval Event")
		logger.Printf("Event ID: %s", approvalEvent.EventID)
		logger.Printf("User ID: %s", approvalEvent.UserID)
		logger.Printf("Open ID: %s", approvalEvent.OpenID)
		logger.Printf("Event Type: %s", approvalEvent.Type)
		// 处理其他审批事件字段

		// 返回成功响应
		return events.APIGatewayProxyResponse{StatusCode: 200}, nil
	}
	// 处理其他飞书事件
	switch event.Type {

	case "event_type_approval_instance_status_change":
		// 处理事件类型1的逻辑
		var approvalEvent ApprovalEvent
		logger.Println("event_type_approval_instance_status_change")
		SendMsg("event_type_approval_instance_status_change")
		if err := json.Unmarshal(event.Event, &approvalEvent); err != nil {
			logger.Println("Failed to parse approval event: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}
		SendMsg("Approval Code: " + approvalEvent.ApprovalCode)

		logger.Printf("Event ID: %s", approvalEvent.EventID)
		logger.Printf("User ID: %s", approvalEvent.UserID)
		logger.Printf("Open ID: %s", approvalEvent.OpenID)
		logger.Printf("Event Type: %s", approvalEvent.Type)

		SendMsg(fmt.Sprintf("Event ID: %s", approvalEvent.EventID))
		SendMsg(fmt.Sprintf("User ID: %s", approvalEvent.UserID))
		SendMsg(fmt.Sprintf("Open ID: %s", approvalEvent.OpenID))
		SendMsg(fmt.Sprintf("Event Type: %s", approvalEvent.Type))

	case "event_type_approval_instance_created":
		SendMsg("event_type_approval_instance_created")
		// 处理事件类型2的逻辑
		logger.Println("event_type_approval_instance_created")
	case "event_type_approval_instance_finished":
		SendMsg("event_type_approval_instance_finished")
		logger.Println("event_type_approval_instance_finished")

	case "event_callback":
		var approvalEvent ApprovalEvent
		if err := json.Unmarshal(event.Event, &approvalEvent); err != nil {
			logger.Println("Failed to parse approval event: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}
		SendMsg("event_callback")
		SendMsg("Instance ID: " + approvalEvent.EventID)
		response := events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       "OK",
		}
		return response, nil
	default:
		var approvalEvent ApprovalEvent
		SendMsg(fmt.Sprintf("Unhandled Feishu event type: %s", event.Type))
		if err := json.Unmarshal(event.Event, &approvalEvent); err != nil {
			logger.Println("Failed to parse approval event: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}
		logger.Println("Unhandled Feishu event type: %s", event.Type)

	}

	// 返回成功响应
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
