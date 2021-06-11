package sdk

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
)

//func NewFileUploadRequest(uri string, params map[string]string, paramName, path string) error {
//	file, err := os.Open(path)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//	part, err := writer.CreateFormFile(paramName, path)
//	if err != nil {
//		return err
//	}
//	_, err = io.Copy(part, file)
//	for key, val := range params {
//		_ = writer.WriteField(key, val)
//	}
//	writer.WriteField("image_type", "message")
//	err = writer.Close()
//	if err != nil {
//		return err
//	}
//	request, err := http.NewRequest("POST", uri, body)
//	request.Header.Set("Content-Type", writer.FormDataContentType())
//	request.Header.Set("Authorization", "Bearer t-0476c29d27f63fa39535f9f268acec043b849ab6")
//	client := http.Client{}
//	resp, err := client.Do(request)
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	respBytes, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		logger.Error(err.Error())
//		return err
//	}
//	str := (*string)(unsafe.Pointer(&respBytes))
//	logger.Error(*str)
//	return err
//}

func (t Tenant) UploadPicBytes(picBytes []byte, imageType string) (string, error) {
	var err error
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	image, err := writer.CreateFormFile("image", "pkg-download-qrcode.png")

	if err != nil {
		logger.Errorf("create form field error:%s", err)
		return "", err
	}

	_, err = image.Write(picBytes)
	if err != nil {
		logger.Errorf("write image error:%s", err)
		return "", err
	}

	err = writer.WriteField("image_type", imageType)
	if err != nil {
		logger.Errorf("write image_type error:%s", err)
		return "", err
	}
	err = writer.Close()
	if err != nil {
		logger.Errorf("write close error:%s", err)
		return "", err
	}

	rsp := &vo.UploadImageRsp{}
	resp, err := http.PostImage(consts.ApiUploadImage, body, http.BuildTokenHeaderOptions(t.TenantAccessToken), http.BuildMultiPartHeaderOptions(writer))
	if err != nil {
		logger.Errorf("upload pic failed:%s", err)
		return "", err
	}
	json.FromJsonIgnoreError(resp, rsp)
	return rsp.Data.ImageKey, nil
}

func (t Tenant) UploadPicFile(filePath string, imageType string) (string, error) {
	var err error
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	defer file.Close()
	if err != nil {
		logger.Errorf("open qr code file failed:%v", file)
		return "", err
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	image, err := writer.CreateFormFile("image", filePath)
	if err != nil {
		logger.Errorf("create form file error:%s", err)
		return "", err
	}

	_, err = io.Copy(image, file)
	if err != nil {
		logger.Errorf("write image error:%s", err)
		return "", err
	}

	err = writer.WriteField("image_type", imageType)
	if err != nil {
		logger.Errorf("write image_type error:%s", err)
		return "", err
	}
	err = writer.Close()
	if err != nil {
		logger.Errorf("write close error:%s", err)
		return "", err
	}

	rsp := &vo.UploadImageRsp{}
	resp, err := http.PostImage(consts.ApiUploadImage, body, http.BuildTokenHeaderOptions(t.TenantAccessToken), http.BuildMultiPartHeaderOptions(writer))
	if err != nil {
		logger.Errorf("upload pic failed:%s", err)
		return "", err
	}
	json.FromJsonIgnoreError(resp, rsp)
	return rsp.Data.ImageKey, nil
}
