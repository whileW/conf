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
	loadXnMysql(mode,ak,sk)
	loadSysMsgMysql(mode,ak,sk)
	loadRedis(mode,ak,sk)
	loadTxTRTC(mode,ak,sk)
	loadTx(mode,ak,sk)
	loadDx(mode,ak,sk)
	loadEtcd(mode,ak,sk)
	loadWps(mode,ak,sk)
	loadFile(mode,ak,sk)
	loadNsq(mode,ak,sk)
}
//加载西牛数据库
func loadXnMysql(mode,ak,sk string) {
	confString := GetNacos("xn-mysql",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.XnMysql)
	if err != nil {
		panic(err)
	}
}
//加载消息数据库
func loadSysMsgMysql(mode,ak,sk string)  {
	confString := GetNacos("sys-msg-mysql",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.SysMsgMysql)
	if err != nil {
		panic(err)
	}
}
//加载redis数据库
func loadRedis(mode,ak,sk string) {
	confString := GetNacos("redis",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.Redis)
	if err != nil {
		panic(err)
	}
}
//加载etcd配置
func loadEtcd(mode,ak,sk string)  {
	confString := GetNacos("etcd",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.Etcd)
	if err != nil {
		panic(err)
	}
}
//加载腾讯语音服务密钥
func loadTxTRTC(mode,ak,sk string)  {
	confString := GetNacos("tx-trtc",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.TxTRTC)
	if err != nil {
		panic(err)
	}
}
//加载腾讯服务密钥
func loadTx(mode,ak,sk string)  {
	confString := GetNacos("tx",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.Tx)
	if err != nil {
		panic(err)
	}
}
//加载顶象服务密钥
func loadDx(mode,ak,sk string)  {
	confString := GetNacos("dx",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.Dx)
	if err != nil {
		panic(err)
	}
}
//加载wps密钥
func loadWps(mode,ak,sk string)  {
	confString := GetNacos("wps",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.Wps)
	if err != nil {
		panic(err)
	}
}
//加载file文件
func loadFile(mode,ak,sk string)  {
	confString := GetNacos("file_conf",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.FileConf)
	if err != nil {
		panic(err)
	}
}
//加载wps密钥
func loadNsq(mode,ak,sk string)  {
	confString := GetNacos("nsq",mode,ak,sk)
	if confString == ""{
		panic("get nacos config error")
	}
	err := json.Unmarshal([]byte(confString),&Conf.Nsq)
	if err != nil {
		panic(err)
	}
}
