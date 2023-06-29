package larkAPI

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/approval/v4"
	"log"
	"main/lib"
	"strconv"
)

func Subscribe(appID string, appSecret string, approvalCode string, client *lark.Client) error {

	token, err := getTenantAccessToken(appID, appSecret)
	if err != nil {
		return err
	}

	//client := lark.NewClient(appID, appSecret, lark.WithEnableTokenCache(false))
	req := larkapproval.NewSubscribeApprovalReqBuilder().ApprovalCode(approvalCode).Build()
	resp, err := client.Approval.Approval.Subscribe(context.Background(), req, larkcore.WithTenantAccessToken(token))

	if err != nil {
		return err
	}

	lib.SendMsg(string(resp.RawBody))

	// Server Error Handling
	if !resp.Success() {
		log.Println(resp.Code, resp.Msg, resp.RequestId())
		lib.SendMsg(strconv.Itoa(resp.Code) + " " + resp.Msg + "" + resp.RequestId())
		return err
	}
	lib.SendMsg(string(resp.RawBody))
	return nil
}
