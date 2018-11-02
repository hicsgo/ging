package model

import (
	"log"
)

import (
	"github.com/jinzhu/gorm"
	"github.com/hicsgo/glib"
)

import (
	"ging/sharing"
	"ging/setting_config"
)

/* ================================================================================
 * 数据模型相关信息
 * email   : golang123@outlook.com
 * author  : hicsgo
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 基础数据模型(Id应该直接放入basemodel，业务限制现在不使用)
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type BaseModel struct {
	//Id       string   `gorm:"primary_key;column:id"`
	DbMap     *gorm.DB `msgpack:"-" sql:"-" json:"-"`
	DBName    string   `msgpack:"-" sql:"-" json:"-"`
	TableName string   `msgpack:"-" sql:"-" json:"-"`
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化DbMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (baseModel *BaseModel) ModelDbMap(s sharing.ISharing, isWriteDB bool) {

	if isWriteDB {
		baseModel.DBName = s.GetWriteDBName()
		baseModel.TableName = s.GetWriteTableName()
	} else {
		baseModel.DBName = s.GetReadDBName()
		baseModel.TableName = s.GetReadTableName()
	}

	dbMap := sharing.GetDatabaseMap(baseModel.DBName, *setting_config.GlobalSetting)
	baseModel.DbMap = dbMap
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 注册回调钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) RegisterCallback() {
	gorm.DefaultCallback.Update().Replace("gorm:after_update", baseModel.AfterUpdate)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * AfterCreate钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) AfterCreate(scope *gorm.Scope) {
	tableName := scope.TableName()
	log.Printf("AfterCreate Hook: %s, id: %s", tableName, baseModel.Id)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * AfterUpdate钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) AfterUpdate(scope *gorm.Scope) {
	tableName := scope.TableName()
	log.Printf("AfterUpdate Hook: %s, id: %s", tableName, baseModel.Id)

	RemoveFromCache(baseModel.Id)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * AfterDelete钩子
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
/*
func (baseModel *BaseModel) AfterDelete(scope *gorm.Scope) {
	tableName := scope.TableName()
	log.Printf("AfterDelete Hook: %s, id: %s", tableName, baseModel.Id)

	RemoveFromCache(baseModel.Id)
}
*/

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 事务
 * fun: 回调函数，接受事务DbMap
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (baseModel *BaseModel) Transactions(s sharing.ISharing, isWriteDB bool, fun func(dbMap *gorm.DB)) error {
	if isWriteDB {
		baseModel.DBName = s.GetWriteDBName()
		baseModel.TableName = s.GetWriteTableName()
	} else {
		baseModel.DBName = s.GetReadDBName()
		baseModel.TableName = s.GetReadTableName()
	}

	var tranDbMap *gorm.DB = nil
	defer func() {
		if tranDbMap != nil {
			tranDbMap.Close()
			log.Printf("Trans Close")
		}
	}()
	err := glib.Capture2(
		func() {
			log.Printf("Trans Begin")
			tranDbMap = sharing.GetDatabaseMap(baseModel.DBName, *setting_config.GlobalSetting).Begin()
			baseModel.DbMap = tranDbMap

			fun(tranDbMap)

			tranDbMap.Commit()
			log.Printf("Trans Commit")
		}, func(e interface{}) {
			tranDbMap.Rollback()
			log.Printf("Trans Rollback %v", e)
		})
	return err
}
