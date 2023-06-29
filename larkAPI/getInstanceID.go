package larkAPI

import (
	"encoding/json"
)

type Event struct {
	AppID        string `json:"app_id"`
	ApprovalCode string `json:"approval_code"`
	CustomKey    string `json:"custom_key"`
	DefKey       string `json:"def_key"`
	GenerateType string `json:"generate_type"`
	InstanceCode string `json:"instance_code"`
	OpenID       string `json:"open_id"`
	OperateTime  string `json:"operate_time"`
	Status       string `json:"status"`
	TaskID       string `json:"task_id"`
	TenantKey    string `json:"tenant_key"`
	Type         string `json:"type"`
	UserID       string `json:"user_id"`
}

type Data struct {
	UUID  string `json:"uuid"`
	Event Event  `json:"event"`
}

func GetInstanceID(jsonStr string) (string, error) {

	var data Data
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	return data.Event.InstanceCode, nil
}
