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

func GetInstanceDetails(appID string, appSecret string, instanceID string, client *lark.Client) (string, error) {

	req := larkapproval.NewGetInstanceReqBuilder().InstanceId(instanceID).Build()
	token, err := getTenantAccessToken(appID, appSecret)
	resp, err := client.Approval.Instance.Get(context.Background(), req, larkcore.WithTenantAccessToken(token))

	if err != nil {
		return "", err
	}

	// Server error handling
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		err = errors.New(strconv.Itoa(resp.Code) + " " + resp.Msg + " " + resp.RequestId())
		return "", err
	}

	// Processing
	fmt.Println(larkcore.Prettify(resp))
	return string(resp.RawBody), nil
}
