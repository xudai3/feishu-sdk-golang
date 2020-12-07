package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
)

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessage(msg vo.MsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiRobotSendMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessageBatch(msg vo.BatchMsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiRobotSendBatchMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// ---------------
// 自行封装的一些方法

func (t Tenant) SendTextMessage(chatId string, text string) {
	if chatId == "" {
		logger.Debugf("chatId empty, send msg failed")
		return
	}
	msg := vo.MsgVo{
		MsgType: consts.MsgTypeText,
		ChatId: chatId,
		Content: &vo.MsgContent{
			Text: text,
		},
	}
	_, err := t.SendMessage(msg)
	if err != nil {
		logger.Errorf("send message:%s failed:%s", text, err)
	}
}

func (t Tenant) SendImageMessage(chatId string, imageKey string) {
	if chatId == "" {
		logger.Debugf("chatId empty, send msg failed")
		return
	}
	req := vo.MsgVo{
		MsgType: consts.MsgTypeImage,
		ChatId: chatId,
		Content: &vo.MsgContent{
			ImageKey:imageKey,
		}}
	rsp, err := t.SendMessage(req)
	if err != nil {
		logger.Errorf("send image message failed rsp:%v err:%v", rsp, err)
		return
	}
}

func (t Tenant) SendCardMessage(chatId string, card *vo.Card, update bool) {
	req := vo.MsgVo{
		MsgType: consts.MsgTypeInteractive,
		ChatId:  chatId,
		UpdateMulti: update,
		Card:        card,
	}
	rsp, err := t.SendMessage(req)
	if err != nil {
		logger.Errorf("send card message failed rsp:%v err:%v", rsp, err)
		return
	}
}
