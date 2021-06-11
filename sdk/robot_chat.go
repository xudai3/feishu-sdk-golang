package sdk

import (
	"fmt"

	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
)

//创建群 https://open.feishu.cn/document/ukTMukTMukTM/ukDO5QjL5gTO04SO4kDN
func (t Tenant) CreateChat(msg vo.CreateChatReqVo) (*vo.CreateChatRespVo, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiCreateChat, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.CreateChatRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取群列表 https://open.feishu.cn/document/ukTMukTMukTM/uITO5QjLykTO04iM5kDN
func (t Tenant) ChatList(pageSize int, pageToken string) (*vo.GroupListRespVo, error) {
	queryParams := map[string]interface{}{}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiChatList, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.GroupListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取群信息 https://open.feishu.cn/document/ukTMukTMukTM/uMTO5QjLzkTO04yM5kDN
func (t Tenant) ChatInfo(chatId string) (*vo.ChatInfoRespVo, error) {
	queryParams := map[string]interface{}{
		"chat_id": chatId,
	}

	respBody, err := http.Get(consts.ApiChatInfo, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.ChatInfoRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//更新群信息 https://open.feishu.cn/document/ukTMukTMukTM/uYTO5QjL2kTO04iN5kDN
func (t Tenant) UpdateChat(msg vo.UpdateChatReqVo) (*vo.UpdateChatRespVo, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiUpdateChat, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.UpdateChatRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//拉用户进群 https://open.feishu.cn/document/ukTMukTMukTM/ucTO5QjL3kTO04yN5kDN
func (t Tenant) AddChatUser(msg vo.UpdateChatMemberReqVo) (*vo.UpdateChatMemberRespVo, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiAddChatUser, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.UpdateChatMemberRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//移除用户出群 https://open.feishu.cn/document/ukTMukTMukTM/uADMwUjLwADM14CMwATN
func (t Tenant) RemoveChatUser(msg vo.UpdateChatMemberReqVo) (*vo.UpdateChatMemberRespVo, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiRemoveChatUser, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.UpdateChatMemberRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//解散群 https://open.feishu.cn/document/ukTMukTMukTM/uUDN5QjL1QTO04SN0kDN
func (t Tenant) DisbandChat(msg vo.UpdateChatData) (*vo.CommonVo, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiDisbandChat, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//拉机器人进群 https://open.feishu.cn/document/ukTMukTMukTM/uYDO04iN4QjL2gDN
func (t Tenant) AddBot(msg vo.UpdateChatData) (*vo.CommonVo, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiAddBot, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// ---------------
// 自行封装的一些方法

func (t Tenant) GetGroupIdByName(groupName string) (string, error) {
	var groupId string
	groupData, err := t.ChatList(100, "")
	if err != nil {
		logger.Errorf("get chat list failed:%s", err)
		return "", err
	}
	if groupData == nil {
		err = fmt.Errorf("group:%s's groupData empty", groupName)
		logger.Errorf(err.Error())
		return "", err
	} else {
		for _, group := range groupData.Data.Groups {
			if group.Name == groupName {
				groupId = group.ChatId
			}
		}
	}
	if groupId == "" {
		err = fmt.Errorf("未找到该群:%s", groupName)
		logger.Error(err.Error())
		return "", err
	}
	return groupId, nil
}

func (t Tenant) ListUserOpenIdsFromGroup(groupName string) ([]string, error) {
	var openIds []string
	groupId, err := t.GetGroupIdByName(groupName)
	if err != nil {
		return openIds, err
	}
	if groupId == "" {
		return openIds, fmt.Errorf("groupId empty")
	}

	memberData, err := t.ChatInfo(groupId)
	if err != nil {
		logger.Errorf("get member list failed:%s", err)
		return openIds, err
	}
	openIds = make([]string, len(memberData.Data.Members))
	for idx, member := range memberData.Data.Members {
		openIds[idx] = member.OpenId
	}
	return openIds, nil
}

func (t Tenant) ListUserInfosFromGroup(groupName string) ([]vo.UserDetailInfo, error) {
	openIds, err := t.ListUserOpenIdsFromGroup(groupName)
	if err != nil {
		return nil, err
	}

	userDetails, err := t.ListUsersByOpenIds(openIds)
	if err != nil {
		logger.Errorf("get user details failed:%s", err)
		return nil, err
	}
	return userDetails, nil
}
