package sharing

import (
	"fmt"
	"time"
	"log"
)

import (
	"github.com/jinzhu/gorm"
)

import (
	"ging/setting"
)

/* ================================================================================
 * 连接数据库设置
 * email   : golang123@outlook.com
 * author  : hicsgo
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库(库/表名)接口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type ISharing interface {
	GetReadDBName() string
	GetReadTableName() string
	GetWriteDBName() string
	GetWriteTableName() string
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取DatabaseMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetDatabaseMap(dbKey string, setting setting.Setting) *gorm.DB {
	var currentDatabase *setting.DatabaseConnectionOption
	for _, database := range setting.DatabaseConfig.Connections {
		if database.Key == dbKey {
			currentDatabase = database
			break
		}
	}

	isLog := setting.DatabaseConfig.IsLog
	dbMap, err := getDatabaseConnection(*currentDatabase, isLog)
	if err != nil {
		panic(fmt.Sprintf("database connection fault: %s", err.Error()))
	}

	return dbMap
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库链接
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func getDatabaseConnection(connectionOption setting.DatabaseConnectionOption, isLog bool) (*gorm.DB, error) {
	dsn := connectionOption.Username + ":" + connectionOption.Password + "@tcp(" + connectionOption.Host + ")/" + connectionOption.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbMap, err := gorm.Open(connectionOption.Dialect, dsn)

	if err != nil {
		log.Printf("Error connecting to db: %s", err.Error())
	}
	dbMap.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")
	dbMap.DB().SetMaxIdleConns(16)
	dbMap.DB().SetMaxOpenConns(512)
	dbMap.DB().SetConnMaxLifetime(time.Hour)
	dbMap.LogMode(isLog)

	return dbMap, err
}
