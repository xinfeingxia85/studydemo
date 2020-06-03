package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
)

func main() {
	var endpoint = "acm.aliyun.com"
	var namespaceId = "7e2bcf2a-71e0-4369-a54c-398dcffa9650"
	var accessKey = "LTAI4Fn2FCiq143LfHrL2utj"
	var secretKey = "MbrmaRQtSZbr4NGlWQmsa2hwF8UEq2"
	status := make(chan bool)

	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	// Initialize client.
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "cn.sjk.zeus.sms.enable"
	var group = "DEFAULT_GROUP"

	// Get plain content from ACM.
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	fmt.Println(err)

	fmt.Println("Get config：" + content)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			status <- true
			fmt.Println(data)
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case <-status:
			fmt.Println("退出le ")
			os.Exit(-1)
		}
	}
}