package conf

type Mysql struct {
	Host 		string 	`json:"host"`
	Port 		string	`json:"port"`
	Database	string	`json:"database"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}
type config struct {
	ZyjyMysql		Mysql	`json:"zyjy_mysql"`
}