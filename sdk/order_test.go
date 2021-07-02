package sdk

import (
	"testing"

	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
	"gotest.tools/assert"
)

func TestTenant_CheckUser(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2deaa88e790f575e", consts.Debug)
	t.Log(e)

	resp, err := tenant.CheckUser(vo.CheckUserReq{
		OpenId: "ou_53681e85d8dfae37c76849145c3cf98f",
		//UserId: "",
	})
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetOrderList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2deaa88e790f575e", consts.Debug)
	t.Log(e)

	resp, err := tenant.GetOrderList(vo.GetOrderListReq{
		PageSize: 10,
	})
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetOrderInfo(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, consts.Debug)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2deaa88e790f575e", consts.Debug)
	t.Log(e)

	resp, err := tenant.GetOrderInfo("6833281668439097346")
	logger.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
