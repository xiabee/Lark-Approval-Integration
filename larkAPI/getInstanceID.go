package larkAPI

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkapproval "github.com/larksuite/oapi-sdk-go/v3/service/approval/v4"
)

func GetInstanceID(approvalCode string, client *lark.Client) (string, error) {

	//client := lark.NewClient("appID", "appSecret", lark.WithEnableTokenCache(false))

	// create request object
	req := larkapproval.NewListInstanceReqBuilder().PageSize(100).ApprovalCode(approvalCode).
		StartTime(`1088038907281`).
		EndTime(`1988038927281`).
		Build()

	resp, err := client.Approval.Instance.List(context.Background(), req, larkcore.WithTenantAccessToken("t-g1046tkNGT3ZDRP7YAEZVFCLRUP3SDGG6QF2SSZS"))

	if err != nil {
		return "", err
	}

	// Server Error Handling
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return "", nil
	}

	// business processing
	fmt.Println(larkcore.Prettify(resp))
	return "", nil
}
