package conf

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)
// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
var endpoint = "acm.aliyun.com"+ ":8080"
var namespaceId = "c6dce47e-b299-4543-b163-4cdacbb93875"

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
		DataId: "com.xn.nacos."+module+".properties",
		Group:  env})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Get config：" + content)
	return content
}