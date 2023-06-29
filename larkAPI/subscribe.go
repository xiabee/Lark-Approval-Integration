package larkAPI

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/approval/v4"
)

func Subscribe(client *lark.Client, approvalCode string) {

	req := larkapproval.NewSubscribeApprovalReqBuilder().ApprovalCode(approvalCode).Build()

	resp, err := client.Approval.Approval.Subscribe(context.Background(), req, larkcore.WithTenantAccessToken("t-g1046tkNGT3ZDRP7YAEZVFCLRUP3SDGG6QF2SSZS"))

	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
