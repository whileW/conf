package conf

type Mysql struct {
	Host 		string 	`json:"host"`
	Port 		string	`json:"port"`
	Database	string	`json:"database"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}
type Xcx struct {
	AppId     string 	`json:"app_id"`
	AppSecret string 	`json:"app_secret"`
}

type config struct {
	ZyjyMysql		Mysql	`json:"zyjy_mysql"`
	ZyjyXcx 		Xcx 	`json:"zyjy_xcx"`
}