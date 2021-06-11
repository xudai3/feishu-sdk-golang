package sdk

import (
	"fmt"

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

func (t Tenant) SendTextMessage(openId string, chatId string, text string) error {
	var err error
	if chatId == "" && openId == "" {
		err = fmt.Errorf("openId and chatId empty, send msg failed")
		logger.Debug(err)
		return err
	}
	msg := vo.MsgVo{
		MsgType: consts.MsgTypeText,
		OpenId:  openId,
		ChatId:  chatId,
		Content: &vo.MsgContent{
			Text: text,
		},
	}
	_, err = t.SendMessage(msg)
	if err != nil {
		logger.Errorf("send message:%s failed:%s", text, err)
		return err
	}
	return nil
}

func (t Tenant) SendImageMessage(openId string, chatId string, imageKey string) error {
	var err error
	if chatId == "" && openId == "" {
		err = fmt.Errorf("openId and chatId empty, send msg failed")
		logger.Debug(err)
		return err
	}
	req := vo.MsgVo{
		MsgType: consts.MsgTypeImage,
		OpenId:  openId,
		ChatId:  chatId,
		Content: &vo.MsgContent{
			ImageKey: imageKey,
		}}
	rsp, err := t.SendMessage(req)
	if err != nil {
		logger.Errorf("send image message failed rsp:%v err:%v", rsp, err)
		return err
	}
	return nil
}

func (t Tenant) SendCardMessage(openId string, chatId string, card *vo.Card, update bool) error {
	var err error
	if chatId == "" && openId == "" {
		err = fmt.Errorf("openId and chatId empty, send msg failed")
		logger.Debug(err)
		return err
	}
	req := vo.MsgVo{
		MsgType:     consts.MsgTypeInteractive,
		OpenId:      openId,
		ChatId:      chatId,
		UpdateMulti: update,
		Card:        card,
	}
	rsp, err := t.SendMessage(req)
	if err != nil {
		logger.Errorf("send card message failed rsp:%v err:%v", rsp, err)
		return err
	}
	return nil
}
