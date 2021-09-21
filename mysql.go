package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int    `json:"port"`
}

//获取mysql配置

func GetMysqlFromConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	mysqlConfig := &MysqlConfig{}
	return mysqlConfig, config.Get(path...).Scan(mysqlConfig)
}
