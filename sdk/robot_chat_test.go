package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
	"gotest.tools/assert"
	"testing"
)

func TestTenant_CreateChat(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.CreateChat(vo.CreateChatReqVo{
		Name:        "test",
		Description: "hhhh",
		OpenIds: []string{
			"ou_433f2b1abf3e0ec316fd9e60d4cda654",
			"ou_87f1b2210acad10a90cc3690802626d7",
		},
	})
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.ChatList(0, "")
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatInfo(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.ChatInfo("oc_95f13644b4182501420139ccee26c989")
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_UpdateChat(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.UpdateChat(vo.UpdateChatReqVo{
		ChatId:      "",
		OwnerUserId: "",
		OwnerOpenId: "",
		Name:        "",
		I18nNames:   nil,
	})
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_AddChatUser(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.AddChatUser(vo.UpdateChatMemberReqVo{
		ChatId:  "",
		UserIds: nil,
		OpenIds: nil,
	})
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)

	resp1, err1 := tenant.RemoveChatUser(vo.UpdateChatMemberReqVo{
		ChatId:  "",
		UserIds: nil,
		OpenIds: nil,
	})
	logger.Info(json.ToJsonIgnoreError(resp1), err1)
	assert.Equal(t, err1, nil)
	assert.Equal(t, resp1.Code, 0)
}

func TestTenant_DisbandChat(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.DisbandChat(vo.UpdateChatData{ChatId: ""})
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_AddBot(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.AddBot(vo.UpdateChatData{ChatId: "oc_95f13644b4182501420139ccee26c989"})
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
