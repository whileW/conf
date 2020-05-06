package conf

type Mysql struct {
	Host 		string 	`json:"host"`
	Port 		string	`json:"port"`
	Database	string	`json:"database"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}
type Redis struct {
	Host 		string 	`json:"host"`
	Port 		string	`json:"port"`
}
type Etcd struct {
	Host 		string 	`json:"host"`
	Port 		string	`json:"port"`
}
type TxTRTC struct {
	SDKAppID	int		`json:"sdk_app_id"`
	Secret		string	`json:"secret"`
}
type Tx struct {
	SecretId	string	`json:"secret_id"`
	SecretKey	string	`json:"secret_key"`
}
type Dx struct {
	SecretId	string	`json:"secret_id"`
	SecretKey	string	`json:"secret_key"`
}
type Wps struct {
	AppID	string		`json:"sdk_app_id"`
	APPKEY		string `json:"appkey"`
}
type FileConf struct {
	RootPath 	string		`json:"root_path"`
}
type Nsq struct {
	ProducerHost 		string 	`json:"producer_host"`
	ProducerPort 		string	`json:"producer_port"`
	ConsumerHost		string	`json:"consumer_host"`
	ConsumerPort		string	`json:"consumer_port"`
}
type config struct {
	XnMysql		Mysql	`json:"xn_mysql"`
	SysMsgMysql	Mysql	`json:"sys_msg_mysql"`
	Redis		Redis	`json:"redis"`
	Etcd		Etcd	`json:"etcd"`
	TxTRTC		TxTRTC	`json:"tx_trtc"`
	Tx 			Tx		`json:"tx"`
	Dx			Dx		`json:"dx"`
	Wps			Wps		`json:"wps"`
	FileConf 	FileConf `json:"file_conf"`
	Nsq 		Nsq 	`json:"nsq"`
}