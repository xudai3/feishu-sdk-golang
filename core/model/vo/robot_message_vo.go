package vo

import "encoding/binary"

//定义参照: https://open.feishu.cn/open-apis/message/v4/send/
type MsgVo struct {
	OpenId      string `json:"open_id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Email       string `json:"email,omitempty"`
	ChatId      string `json:"chat_id,omitempty"`
	MsgType     string `json:"msg_type"`
	RootId      string `json:"root_id,omitempty"`
	UpdateMulti bool   `json:"update_multi"`

	Card    *Card       `json:"card,omitempty"`
	Content *MsgContent `json:"content,omitempty"`
}

type BatchMsgVo struct {
	DepartmentIds []string `json:"department_ids"`
	OpenIds       []string `json:"open_ids"`
	UserIds       []string `json:"user_ids"`
	MsgType       string   `json:"msg_type"`

	Card    *Card       `json:"card,omitempty"`
	Content *MsgContent `json:"content,omitempty"`
}

type MsgContent struct {
	Text     string   `json:"text"`
	ImageKey string   `json:"image_key"`
	Post     *MsgPost `json:"post,omitempty"`
}

type MsgPost struct {
	ZhCn *MsgPostValue `json:"zh_cn,omitempty"`
	EnUs *MsgPostValue `json:"en_us,omitempty"`
	JaJp *MsgPostValue `json:"ja_jp,omitempty"`
}

type MsgPostValue struct {
	Title   string      `json:"title"`
	Content interface{} `json:"content"`
}

type MsgPostContentText struct {
	Tag      string `json:"tag"`
	UnEscape bool   `json:"un_escape"`
	Text     string `json:"text"`
}

type MsgPostContentA struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
	Href string `json:"href"`
}

type MsgPostContentAt struct {
	Tag    string `json:"tag"`
	UserId string `json:"user_id"`
}

type MsgPostContentImage struct {
	Tag      string  `json:"tag"`
	ImageKey string  `json:"image_key"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
}

type MsgResp struct {
	CommonVo

	Data MsgRespData `json:"data"`
}

type MsgRespData struct {
	MessageId string `json:"message_id"`
}

type UpdateImageVo struct {
	Image binary.ByteOrder
}

type UploadImageRsp struct {
	CommonVo
	Data ImageRspData `json:"data"`
}

type ImageRspData struct {
	ImageKey string `json:"image_key"`
}
