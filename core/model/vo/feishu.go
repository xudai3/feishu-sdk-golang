package vo

type Feishu struct {
	Type      string      `json:"type"`
	Token     string      `json:"token"`
	Event     FeishuEvent `json:"event"`
	MsgType   string      `json:"msg_type"`
	Challenge string      `json:"challenge"`
}

type FeishuInteractive struct {
	OpenId        string            `json:"open_id"`
	UserId        string            `json:"user_id"`
	OpenMessageId string            `json:"open_message_id"`
	TenantKey     string            `json:"tenant_key"`
	Action        InteractiveAction `json:"action"`
	Challenge     string            `json:"challenge"`
}
type FeishuEvent struct {
	Type             string `json:"type"`
	AppId            string `json:"app_id"`
	TenantKey        string `json:"tenant_key"`
	RootId           string `json:"root_id"`
	ParentId         string `json:"parent_id"`
	OpenChatId       string `json:"open_chat_id"`
	ChatType         string `json:"chat_type"`
	MsgType          string `json:"msg_type"`
	OpenId           string `json:"open_id"`
	OpenMessageId    string `json:"open_message_id"`
	IsMention        bool   `json:"is_mention"`
	Text             string `json:"text"`
	TextWithoutAtBot string `json:"text_without_at_bot"`
}

type InteractiveAction struct {
	Value  interface{} `json:"value"`
	Tag    string      `json:"tag"`
	Option string      `json:"option"`
}
