package gtools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gtools/configs"
	"gtools/util"
	"io"
	"net/http"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

func (a *App) Translation(query string) *util.Resp {
	trans := *a.Trans
	return trans.translation(query)
}

type ITrans interface {
	translation(string) *util.Resp
}

type TxTransResult struct {
	Response TxResponse `json:"response"`
}

type TxResponse struct {
	TargetText string `json:"TargetText"`
	Source     string `json:"Source"`
	Target     string `json:"Target"`
	RequestId  string `json:"RequestId"`
}

type BdTransResult struct {
	From        string      `json:"from"`
	To          string      `json:"to"`
	TransResult []transItem `json:"trans_result"`
}
type transItem struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type BdTransClient struct {
	ConfigMap map[string]string
	Log       *logrus.Logger
}

type TxTransClient struct {
	ConfigMap map[string]string
	Log       *logrus.Logger
}

type AliTransClient struct {
	ConfigMap map[string]string
	Log       *logrus.Logger
}

func (t *BdTransClient) translation(query string) *util.Resp {
	var appid = t.ConfigMap["appid"]
	var salt = t.ConfigMap["salt"]
	var secret = t.ConfigMap["secret"]
	var str = appid + query + salt + secret
	h := md5.New()
	h.Write([]byte(str))
	var sign = hex.EncodeToString(h.Sum(nil))
	req, _ := http.NewRequest("GET", configs.BdTransUrl, nil)
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("from", "auto")
	q.Add("to", "zh")
	q.Add("appid", appid)
	q.Add("salt", salt)
	q.Add("sign", sign)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log.Errorf(configs.BdTransRequestErr, err.Error())
		return util.Error("翻译失败，请检查百度翻译设置项")
	}
	defer resp.Body.Close()
	var result []byte
	result, _ = io.ReadAll(resp.Body)
	if strings.Contains(string(result), "error_code") {
		return util.Error("翻译失败，请检查百度翻译设置项")
	}
	var resultJson = new(BdTransResult)
	err = json.Unmarshal(result, &resultJson)
	if err != nil {
		return util.Error("翻译结果解析失败，请重试")
	}
	var resultStr = ""
	for _, v := range resultJson.TransResult {
		resultStr += (v.Dst + "\n")
	}
	return util.Success(resultStr)
}

func (t *TxTransClient) translation(query string) *util.Resp {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		t.ConfigMap["secretid"],
		t.ConfigMap["secretkey"],
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = configs.Endpoint
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := tmt.NewClient(credential, t.ConfigMap["region"], cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tmt.NewTextTranslateRequest()

	request.SourceText = common.StringPtr(query)
	request.Source = common.StringPtr(t.ConfigMap["from"])
	request.Target = common.StringPtr(t.ConfigMap["to"])
	request.ProjectId = common.Int64Ptr(0)

	// 返回的resp是一个TextTranslateResponse的实例，与请求对象对应
	response, err := client.TextTranslate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		t.Log.Error(fmt.Sprintf(configs.TxTransRequestErr, err))
		return util.Error("腾讯翻译API异常")
	}
	if err != nil {
		return util.Error("腾讯翻译失败")
	}
	return util.Success(response.Response.TargetText)
}

func (t *AliTransClient) translation(query string) *util.Resp {
	// 创建ecsClient实例
	t.Log.Info(t.ConfigMap["mtAccessKeyId"])
	alimtClient, err := alimt.NewClientWithAccessKey(
		t.ConfigMap["region"],            // 地域ID
		t.ConfigMap["mtAccessKeyId"],     // 您的Access Key ID
		t.ConfigMap["mtAccessKeySecret"]) // 您的Access Key Secret
	if err != nil {
		// 异常处理
		t.Log.Errorf(configs.AliTransCreateErr, err.Error())
		return util.Error("阿里翻译初始化失败，请检查配置")
	}
	// 创建API请求并设置参数
	request := alimt.CreateTranslateECommerceRequest()
	// 等价于 request.PageSize = "10"
	request.Method = t.ConfigMap["method"]       //设置请求
	request.FormatType = t.ConfigMap["type"]     //翻译文本的格式
	request.SourceLanguage = t.ConfigMap["from"] //源语言
	request.SourceText = query                   //原文
	request.TargetLanguage = t.ConfigMap["to"]   //目标语言
	request.Scene = t.ConfigMap["scene"]         //目标语言
	// 发起请求并处理异常
	response, err := alimtClient.TranslateECommerce(request)
	if err != nil {
		// 异常处理
		t.Log.Errorf(configs.AliTransRequestErr, err.Error())
		return util.Error("阿里翻译请求失败，请检查配置")
	}
	str := response.Data.Translated
	return util.Success(str)
}
