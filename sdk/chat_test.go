package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
	"gotest.tools/assert"
	"testing"
)

func TestTenant_GroupList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.GroupList("u-OVRWRbISgBZK6j9pu2ApJg", 0, "")
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatMembers(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.ChatMembers("", "", 0, "")
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatSearch(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651", consts.Debug)
	t.Log(e)

	resp, err := tenant.ChatSearch("u-OVRWRbISgBZK6j9pu2ApJg", "北极星测试企业。", 0, "")
	logger.Info(json.ToJsonIgnoreError(resp), err)
	logger.Info(111, resp.Data)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
