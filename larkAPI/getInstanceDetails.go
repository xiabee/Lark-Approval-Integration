package larkAPI

import (
	"context"
	"errors"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/approval/v4"
	"strconv"
)

// GetInstanceDetails : Return a GetInstanceResp object for subsequent processing
func GetInstanceDetails(appID string, appSecret string, instanceID string, client *lark.Client) (*larkapproval.GetInstanceResp, error) {

	req := larkapproval.NewGetInstanceReqBuilder().InstanceId(instanceID).Build()
	token, err := getTenantAccessToken(appID, appSecret)
	resp, err := client.Approval.Instance.Get(context.Background(), req, larkcore.WithTenantAccessToken(token))

	if err != nil {
		return nil, err
	}

	// Server error handling
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		err = errors.New(strconv.Itoa(resp.Code) + " " + resp.Msg + " " + resp.RequestId())
		return nil, err
	}

	return resp, nil
}
