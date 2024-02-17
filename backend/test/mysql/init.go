package main

import "webserver/db"

var (
	dbName  = "testa"
	adminDB *db.MysqlDB
	tables  []interface{}
)

// InitAdminDB 初始化数据库
func InitAdminDB(dsn string) error {
	var err error
	adminDB, err = db.InitMysql(dsn, dbName, tables)
	if err != nil {
		return err
	}

	return nil
}
