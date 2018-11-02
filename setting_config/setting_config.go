package setting_config

import (
	"fmt"
	"time"
)

import (
	"ging/setting"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 全局配置数据(项目中自行覆盖使用)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
var GlobalSetting *setting.Setting


/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 配置初始化
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func init() {
	fmt.Printf("%v global settings init\n", time.Now())
	databaseConnOptions := getDatabaseConnections()
	GlobalSetting = &setting.Setting{
		DatabaseConfig: setting.DatabaseOption{
			Connections: databaseConnOptions,
			IsLog:       true,
		},
		IsPro: false,
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库连接字符串初始化配置
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func getDatabaseConnections() []*setting.DatabaseConnectionOption {
	databaseHost := "localhost:3306"
	databaseName := "xxxxxxxxxxxx"

	if setting.Setting.IsPro {
		databaseHost = "localhost:3306"
		databaseName = "xxxxxxxxxxxx"
	}

	connections := []*setting.DatabaseConnectionOption{{
		Key:      "xxxxxxxxxxxx",
		Username: "xxxxxxxxxxxx",
		Password: "xxxxxxxxxxxx",
		Host:     databaseHost,
		Database: databaseName,
		Dialect:  "mysql",
	}, {
		Key:      "xxxxxxxxxxxx",
		Username: "xxxxxxxxxxxx",
		Password: "xxxxxxxxxxxx",
		Host:     databaseHost,
		Database: "xxxxxxxxxxxx",
		Dialect:  "mysql",
	}}

	return connections
}
