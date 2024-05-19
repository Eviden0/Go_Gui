package test

import (
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
)

// 阿里机器翻译SDK测试
func TestAliTrans(t *testing.T) {
	// 创建ecsClient实例
	alimtClient, err := alimt.NewClientWithAccessKey(
		"cn-beijing", // 地域ID
		"A",          // 您的Access Key ID
		"B")          // 您的Access Key Secret
	if err != nil {
		// 异常处理
		panic(err)
	}
	// 创建API请求并设置参数
	request := alimt.CreateTranslateECommerceRequest()
	// 等价于 request.PageSize = "10"
	request.Method = "POST"                       //设置请求
	request.FormatType = "text"                   //翻译文本的格式
	request.SourceLanguage = "en"                 //源语言
	request.SourceText = "hello my name is xiao." //原文
	request.TargetLanguage = "zh"                 //目标语言
	request.Scene = "title"                       //目标语言
	// 发起请求并处理异常
	response, err := alimtClient.TranslateECommerce(request)
	if err != nil {
		// 异常处理
		panic(err)
	}
	str := response.Data.Translated
	fmt.Println(str)
}
