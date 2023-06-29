package larkAPI

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type Request struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type Response struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

// getTenantAccessToken : Refers to https://open.feishu.cn/document/server-docs/authentication-management/access-token/tenant_access_token_internal
func getTenantAccessToken(appID string, appSecret string) (string, error) {

	apiUrl := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	data := Request{
		AppId:     appID,
		AppSecret: appSecret,
	}

	// Convert the data to a byte array in JSON format
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Handle the response
	var responseData Response
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return "", err
	}

	// Handle the lark error
	if responseData.Code != 0 {
		err = errors.New("GetTenantAccessToken error: code " + strconv.Itoa(responseData.Code) + " " + responseData.Msg)
		return "", err
	}

	tenantAccessToken := responseData.TenantAccessToken
	return tenantAccessToken, nil
}
