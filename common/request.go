package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"

	"io"
)

var (
	ServerURL, AuraServer, AuraPort string
)

type Result struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}


//AccessAPI 基础http访问接口
/**
requestMethod:请求方法，GET/POST/DELETE/PUT
requestAPI: 请求地址，如http://business-public.aurayou/api/v1/1/province
projectName: 系统项目名
postBody: post数据主体
headerMap: 请求头数据
formMap: 表单数据
 */
func AccessAPI(requestMethod, requestAPI, postBody, projectName string, headerData map[string]string,formData string) (result Result, err error) {
	client := &http.Client{}
	AuraServer = viper.GetString(projectName + ".host")
	AuraPort = viper.GetString(projectName + ".port")
	ServerURL := "http://" + AuraServer + ":" + AuraPort
	requestURL := ServerURL + requestAPI
	fmt.Println(requestURL)
	var contentType string
	var body io.Reader
	var userName string = "Aura-SIT tester"
	if len(strings.TrimSpace(postBody)) != 0 {
		body = strings.NewReader(postBody)
		contentType = "application/json"
	} else if len(strings.TrimSpace(formData)) != 0 {
		body = strings.NewReader(formData)
		contentType = "application/x-www-form-urlencoded"
	}
	req, err := http.NewRequest(requestMethod, requestURL, body)
	if err != nil {
		logrus.Warnf("Request ", requestMethod, " URL error. Please check the Aura server URL: %s", requestURL)
		return result, fmt.Errorf("Request " + requestMethod + " URL error.Please check the Aura server URL: %s", requestURL)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("UserName", userName)
	//头信息
	if headerData != nil && len(headerData) != 0 {
		for k, v := range headerData {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Warnf("Cannot connect to the Aura server. Please check the Aura server URL: %s", requestURL)
		return result, fmt.Errorf("Cannot connect to the Aura server. Please check the Aura server URL: %s", requestURL)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Warnf("Cannot read the reponse body")
		return result, fmt.Errorf("Cannot read the reponse body")
	}
	bodyTextString := string(bodyText)
	bodyTextStringClean := strings.Replace(bodyTextString, `\`, "", -1)
	result = Result{}
	err = json.Unmarshal([]byte(bodyTextStringClean), &result)
	if err != nil {
		logrus.Warnf("Cannot extract response status JSON")
		return result, fmt.Errorf("Cannot extract response status JSON")
	}
	return result, nil
}

func AccessAPIWithPostBody(requestMethod, requestAPI, postBody, projectName string) (Result, error) {
	return AccessAPI(requestMethod, requestAPI, postBody, projectName, nil, "")
}


func AccessAPIWithDataAndBody(requestMethod, requestAPI, postBody, projectName string, headerData map[string]string, formData string) (Result, error) {
	return AccessAPI(requestMethod, requestAPI, postBody, projectName, headerData, formData)
}