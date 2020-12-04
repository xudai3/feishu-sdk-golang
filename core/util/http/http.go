package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const defaultContentType = "application/json"

var httpClient = &http.Client{}

type HeaderOption struct {
	Name  string
	Value string
}

type QueryParameter struct {
	Key   string
	Value interface{}
}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient = &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}
}

func BuildTokenHeaderOptions(tenantAccessToken string) HeaderOption {
	return HeaderOption{
		Name:  "Authorization",
		Value: "Bearer " + tenantAccessToken,
	}
}

func DeleteRequest(url string, body string, headerOptions ...HeaderOption) (string, error) {
	req, err := http.NewRequest("DELETE", url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", defaultContentType)
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func() {
		if resp != nil{
			if e := resp.Body.Close(); e != nil {
				logger.Error(e)
			}
		}
	}()
	return responseHandle(resp, err)
}

func Delete(url string, params map[string]interface{}, body string, headerOptions ...HeaderOption) (string, error) {
	logger.Infof("请求body %s", body)

	fullUrl := url + ConvertToQueryParams(params)
	return DeleteRequest(fullUrl, body, headerOptions...)
}

func PatchRequest(url string, body string, headerOptions ...HeaderOption) (string, error) {
	req, err := http.NewRequest("PATCH", url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", defaultContentType)
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func() {
		if resp != nil{
			if e := resp.Body.Close(); e != nil {
				logger.Error(e)
			}
		}
	}()
	return responseHandle(resp, err)
}

func Patch(url string, params map[string]interface{}, body string, headerOptions ...HeaderOption) (string, error) {
	logger.Infof("请求body %s", body)

	fullUrl := url + ConvertToQueryParams(params)
	return PatchRequest(fullUrl, body, headerOptions...)
}

func PostRequest(url string, body string, headerOptions ...HeaderOption) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", defaultContentType)
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func() {
		if resp != nil {
			if e := resp.Body.Close(); e != nil {
				logger.Error(e)
			}
		}
	}()
	return responseHandle(resp, err)
}

func Post(url string, params map[string]interface{}, body string, headerOptions ...HeaderOption) (string, error) {
	logger.Infof("请求body %s", body)

	fullUrl := url + ConvertToQueryParams(params)
	return PostRequest(fullUrl, body, headerOptions...)
}

func PostRepetition(url string, params []QueryParameter, body string, headerOptions ...HeaderOption) (string, error) {
	logger.Infof("请求body %s", body)

	fullUrl := url + ConvertToQueryParamsRepetition(params)
	return PostRequest(fullUrl, body, headerOptions...)
}

func GetRequest(url string, headerOptions ...HeaderOption) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	for _, headerOption := range headerOptions {
		req.Header.Set(headerOption.Name, headerOption.Value)
	}
	resp, err := httpClient.Do(req)
	defer func() {
		if resp != nil {
			if e := resp.Body.Close(); e != nil {
				logger.Error(e)
			}
		}
	}()
	return responseHandle(resp, err)
}

func Get(url string, params map[string]interface{}, headerOptions ...HeaderOption) (string, error) {
	fullUrl := url + ConvertToQueryParams(params)
	return GetRequest(fullUrl, headerOptions...)
}

func GetRepetition(url string, params []QueryParameter, headerOptions ...HeaderOption) (string, error) {
	fullUrl := url + ConvertToQueryParamsRepetition(params)
	logger.Debugf("fullUrl:%v", fullUrl)
	return GetRequest(fullUrl, headerOptions...)
}

func responseHandle(resp *http.Response, err error) (string, error) {
	if err != nil {
		logger.Error(err)
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	respBody := string(b)
	logger.Infof("api %s 响应结果: %s", resp.Request.URL, respBody)
	return respBody, nil
}

func ConvertToQueryParams(params map[string]interface{}) string {
	paramsJson := json.ToJsonIgnoreError(params)
	params = map[string]interface{}{}
	_ = json.FromJson(paramsJson, &params)

	if &params == nil || len(params) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for k, v := range params {
		if v == nil {
			continue
		}
		buffer.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}

// 生成key相同但是有很多个不同value的query params
// https://open.feishu.cn/open-apis/contact/v1/department/detail/batch_get?
// department_ids=od-2efe30807a10608754862a63b108828f&department_ids=od-da6427b2adbceb91204d7fa6aeb7e8ff
func ConvertToQueryParamsRepetition(params []QueryParameter) string {
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for _, v := range params {
		if v.Value == nil {
			continue
		}
		buffer.WriteString(fmt.Sprintf("%s=%v&", v.Key, v.Value))
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}
