package sharing

import (
	"fmt"
	"time"
	"log"
	"math/rand"
)

import (
	"github.com/jinzhu/gorm"
	"github.com/hicsgo/ging/setting"
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
	TableName() string
	GetProjectName() string
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取ReadDatabaseMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetReadDatabaseMap(projectName string, s setting.Setting) *gorm.DB {
	var currentDatabase setting.DatabaseConnectionOption
	isLog := true
	for i, dbOption := range s.DatabaseConfig.DatabaseOptions {
		if dbOption.ProjectName == projectName {

			//读库配置数量
			readDBCount := len(s.DatabaseConfig.DatabaseOptions[i].ReadDBConns)
			if readDBCount == 0 {
				break
			} else {
				//随机拉取一个数据库(可以根据权重获取)
				index := rand.Intn(readDBCount)
				currentDatabase = *dbOption.ReadDBConns[index]
				isLog = dbOption.ReadDBConns[index].IsLog
			}
		}
	}
	dbMap, err := getDatabaseConnection(currentDatabase, isLog)
	if err != nil {
		panic(fmt.Sprintf("database connection fault: %s", err.Error()))
	}

	return dbMap
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取WriteDatabaseMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetWriteDatabaseMap(projectName string, s setting.Setting) *gorm.DB {
	var currentDatabase setting.DatabaseConnectionOption
	isLog := true
	for i, dbOption := range s.DatabaseConfig.DatabaseOptions {
		if dbOption.ProjectName == projectName {

			//写库配置数量
			writeDBCount := len(s.DatabaseConfig.DatabaseOptions[i].WirteDBConns)
			if writeDBCount == 0 {
				break
			} else {
				//随机拉取一个数据库(可以根据权重获取)
				index := rand.Intn(writeDBCount)
				currentDatabase = *dbOption.WirteDBConns[index]
				isLog = dbOption.WirteDBConns[index].IsLog
			}
		}
	}
	dbMap, err := getDatabaseConnection(currentDatabase, isLog)
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
