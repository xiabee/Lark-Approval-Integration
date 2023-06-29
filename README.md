# Lark-Approval-Integration
Integration of Lark (AKA Feishu) approval and AWS Lambda



## 环境变量设置

| Key           | **Value**                                           |
| ------------- | --------------------------------------------------- |
| APPROVAL_CODE | 飞书审批 ID                                         |
| APP_ID        | 飞书应用 ID                                         |
| APP_SECRET    | 飞书应用密钥                                        |
| WEBHOOK_KEY   | 飞书群机器人 webhook 最后的哈希串（用于发送群消息） |



## 设计简介

### 飞书订阅流程

获取凭证（tenantAccessToken） :arrow_right: 订阅审批 :arrow_right: 监听订阅事件

参考

* https://open.feishu.cn/document/server-docs/approval-v4/event/subscription-steps
* https://open.feishu.cn/document/server-docs/approval-v4/event/event-interface/subscribe
* https://open.feishu.cn/document/server-docs/approval-v4/instance/get



### 程序执行流程

飞书程序绑定 APIGateway Endpoint（飞书发送 url_verification 同时程序发送订阅申请） :arrow_right: 订阅后飞书发送审批消息给 APIGateway :arrow_right: Lambda 函数处理消息



## Refer

* https://github.com/larksuite/oapi-sdk-go/tree/v3_main
* SSO 授权程序：https://github.com/xiabee/AWS_SSO_Authorization
