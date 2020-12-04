package sdk

import (
	"bytes"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/logger"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"unsafe"
)

func NewFileUploadRequest(uri string, params map[string]string, paramName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.WriteField("image_type", "message")
	err = writer.Close()
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer t-0476c29d27f63fa39535f9f268acec043b849ab6")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		logger.Error(err.Error())
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	logger.Error(*str)
	return err
}
