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
	"math/rand"
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
	GetReadDBKeyName() string
	GetReadTableName() string
	GetWriteDBKeyName() string
	GetWriteTableName() string
	GetProjectName() string
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取ReadDatabaseMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetReadDatabaseMap(dbKey, projectName string, setting setting.Setting) *gorm.DB {
	var currentDatabase *setting.DatabaseConnectionOption
	isLog := true
	for i, dbOption := range setting.DatabaseConfig.DatabaseOptions {
		if dbOption.ProjectName == projectName {
			for _, database := range setting.DatabaseConfig.DatabaseOptions[i].ReadDBConns {
				if database.Key == dbKey {
					currentDatabase = database
					isLog = database.IsLog
					break
				}
			}
		}
	}
	dbMap, err := getDatabaseConnection(*currentDatabase, isLog)
	if err != nil {
		panic(fmt.Sprintf("database connection fault: %s", err.Error()))
	}

	return dbMap
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取WriteDatabaseMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetWriteDatabaseMap(dbKey, projectName string, setting setting.Setting) *gorm.DB {
	var currentDatabase *setting.DatabaseConnectionOption
	isLog := true
	for i, dbOption := range setting.DatabaseConfig.DatabaseOptions {
		if dbOption.ProjectName == projectName {
			for _, database := range setting.DatabaseConfig.DatabaseOptions[i].WirteDBConns {
				if database.Key == dbKey {
					currentDatabase = database
					isLog = database.IsLog
					break
				}
			}
		}
	}
	dbMap, err := getDatabaseConnection(*currentDatabase, isLog)
	if err != nil {
		panic(fmt.Sprintf("database connection fault: %s", err.Error()))
	}

	return dbMap
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据项目名称获取获取读库的key
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetReadDBKey(projectName string) string {
	dbKey := ""
	for i, dbOption := range setting.DatabaseConfig.DatabaseOptions {
		if dbOption.ProjectName == projectName {
			readbConnCount := len(setting.DatabaseConfig.DatabaseOptions[i].ReadDBConns)
			if readbConnCount == 0 {
				break
			} else {
				index := rand.Intn(readbConnCount) //随机拉取一个数据库
				dbKey = dbOption.ReadDBConns[index].Key
			}
		}
	}
	return dbKey
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 根据项目名称获取获取写库的key
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetWirteDBKey(projectName string) string {
	dbKey := ""
	for i, dbOption := range setting.DatabaseConfig.DatabaseOptions {
		if dbOption.ProjectName == projectName {
			readbConnCount := len(setting.DatabaseConfig.DatabaseOptions[i].WirteDBConns)
			if readbConnCount == 0 {
				break
			} else {
				index := rand.Intn(readbConnCount) //随机拉取一个数据库
				dbKey = dbOption.WirteDBConns[index].Key
			}
		}
	}
	return dbKey
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
