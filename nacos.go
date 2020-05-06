package conf

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)
// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
var endpoint = "acm.aliyun.com"+ ":8080"
var namespaceId = "713b7bf0-708c-4772-b325-5d38c5b7544d"

func GetNacos(module,env,ak,sk string) string {
	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint,
		NamespaceId:    namespaceId,
		AccessKey:      ak,
		SecretKey:      sk,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: module,
		Group:  env})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Get config：" + content)
	return content
}