package sdk

import (
	"fmt"

	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
)

type App struct {
	AppId          string
	AppSecret      string
	AppAccessToken string
}

type Tenant struct {
	TenantAccessToken string
	Expire            int64
}

type User struct {
	UserAccessToken string
}

func BuildInternalApp(appId, appSecret, logLevel string) (*App, error) {
	err := logger.InitLogger(logLevel)
	if err != nil {
		fmt.Printf("logger init failed:%v\n", err)
		return nil, err
	}
	resp, err := GetAppAccessTokenInternal(appId, appSecret)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("req err, code: %d, msg: %s", resp.Code, resp.Msg)
	}
	return &App{
		AppId:          appId,
		AppSecret:      appSecret,
		AppAccessToken: appSecret,
	}, nil
}

func BuildApp(appId, appSecret, appTicket, logLevel string) (*App, error) {
	err := logger.InitLogger(logLevel)
	if err != nil {
		fmt.Printf("logger init failed:%v\n", err)
		return nil, err
	}
	resp, err := GetAppAccessToken(appId, appSecret, appTicket)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("req err, code: %d, msg: %s", resp.Code, resp.Msg)
	}
	return &App{
		AppId:          appId,
		AppSecret:      appSecret,
		AppAccessToken: resp.AppAccessToken,
	}, nil
}

func BuildTenantInternal(appId, appSecret, logLevel string) (*Tenant, error) {
	err := logger.InitLogger(logLevel)
	if err != nil {
		fmt.Printf("logger init failed:%v\n", err)
		return nil, err
	}
	resp, err := GetTenantAccessTokenInternal(appId, appSecret)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("req err, code: %d, msg: %s", resp.Code, resp.Msg)
	}
	return &Tenant{
		TenantAccessToken: resp.TenantAccessToken,
		Expire:            resp.Expire,
	}, nil
}

func BuildTenant(appAccessToken, tenantKey, logLevel string) (*Tenant, error) {
	err := logger.InitLogger(logLevel)
	if err != nil {
		fmt.Printf("logger init failed:%v\n", err)
		return nil, err
	}
	resp, err := GetTenantAccessToken(appAccessToken, tenantKey)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("req err, code: %d, msg: %s", resp.Code, resp.Msg)
	}
	return &Tenant{
		TenantAccessToken: resp.TenantAccessToken,
		Expire:            resp.Expire,
	}, nil
}
