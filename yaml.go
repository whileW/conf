package conf

import (
	"encoding/json"
)

var Conf = &config{}
//初始化
func init()  {
	loadConfig()
}
//加载配置数据
func loadConfig() {
	mode := GetXN_MODE()
	ak := getConfigCenterAk()
	sk := getConfigCenterSK()
	if ak == "" || sk == "" {
		panic("config center ak or sk is nil")
	}
	loadZyjyMysql(mode,ak,sk)
	loadZyjyXcx(mode,ak,sk)
}
//加载职业教育数据库
func loadZyjyMysql(mode,ak,sk string) {
	confString := GetNacos("zyjy-mysql",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.ZyjyMysql)
	if err != nil {
		panic(err)
	}
}
//加载职业教育小程序
func loadZyjyXcx(mode,ak,sk string) {
	confString := GetNacos("zyjy-xcx",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.ZyjyXcx)
	if err != nil {
		panic(err)
	}
}