package lib

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type FeishuEvent struct {
	Type    string          `json:"type"`
	Event   json.RawMessage `json:"event"`
	Token   string          `json:"token"`
	Encrypt bool            `json:"encrypt"`
}

// IsChallenge to check if it is a verification request
func IsChallenge(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, bool, error) {
	var event FeishuEvent
	if err := json.Unmarshal([]byte(request.Body), &event); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, false, nil
	}

	// Check if it is a verification request
	if event.Type == "url_verification" {
		// Extract CHALLENGE parameters
		var verificationEvent struct {
			Challenge string `json:"challenge"`
		}
		if err := json.Unmarshal([]byte(request.Body), &verificationEvent); err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 400}, true, nil
		}

		// Build the response object
		responseBody := struct {
			Challenge string `json:"challenge"`
		}{
			Challenge: verificationEvent.Challenge,
		}

		// return validation response
		responseJSON, err := json.Marshal(responseBody)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500}, false, nil
		}

		// Construct a legal JSON response string
		response := events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(responseJSON),
		}
		
		SendMsg("url_verification Succeed challenge: " + verificationEvent.Challenge)
		// return response
		return response, true, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, true, nil
}
