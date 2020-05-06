package conf

import "os"

var xn_mode,http_port,grpc_port = "","",""
var config_center_ak,config_center_sk = "",""

//获取环境配置
func GetXN_MODE() string {
	if xn_mode == "" {
		xn_mode = os.Getenv("XN_MODE")
		if xn_mode == "" {
			xn_mode = "local"
		}
	}
	return xn_mode
}
//获取http监听端口
func GetHttpPort() string {
	if http_port == "" {
		http_port = os.Getenv("HTTP_PORT")
		if http_port == "" {
			http_port = "30001"
		}
	}
	return http_port
}
//获取grpc监听端口
func GetGrpcPort() string {
	if grpc_port == "" {
		grpc_port = os.Getenv("GRPC_PORT")
		if grpc_port == "" {
			grpc_port = "40001"
		}
	}
	return grpc_port
}
//获取配置中心ak 和 sk
func getConfigCenterAk() string {
	if config_center_ak == "" {
		config_center_ak = os.Getenv("ConfigCenterAccessKey")
	}
	return config_center_ak
}
func getConfigCenterSK() string {
	if config_center_sk == "" {
		config_center_sk = os.Getenv("ConfigCenterSecretKey")
	}
	return config_center_sk
}