package lib

import (
	"encoding/json"
)

type ApprovalData struct {
	ApprovalCode string `json:"approval_code"`
	ApprovalName string `json:"approval_name"`
	DepartmentID string `json:"department_id"`
	EndTime      string `json:"end_time"`
	Form         string `json:"form"`
	InstanceCode string `json:"instance_code"`
	OpenID       string `json:"open_id"`
	Reverted     bool   `json:"reverted"`
	SerialNumber string `json:"serial_number"`
	StartTime    string `json:"start_time"`
	Status       string `json:"status"`
	TaskList     []struct {
		EndTime   string `json:"end_time"`
		ID        string `json:"id"`
		NodeID    string `json:"node_id"`
		NodeName  string `json:"node_name"`
		OpenID    string `json:"open_id"`
		StartTime string `json:"start_time"`
		Status    string `json:"status"`
		Type      string `json:"type"`
	} `json:"task_list"`
	Timeline []struct {
		CreateTime string `json:"create_time"`
		Ext        string `json:"ext"`
		NodeKey    string `json:"node_key"`
		OpenID     string `json:"open_id"`
		Type       string `json:"type"`
		UserID     string `json:"user_id"`
	} `json:"timeline"`
	UserID string `json:"user_id"`
	UUID   string `json:"uuid"`
}

// GetApprovalStatus : return the status
func GetApprovalStatus(jsonStr string) (string, error) {
	var data struct {
		Data ApprovalData `json:"data"`
	}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", nil
	}

	return data.Data.Status, nil
}

// GetApprovalForms : return the form
func GetApprovalForms(jsonStr string) (string, error) {
	var data struct {
		Data ApprovalData `json:"data"`
	}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", nil
	}
	return data.Data.Form, nil
}
