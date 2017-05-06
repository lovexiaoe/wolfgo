package base

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func InitDB() error {
	var err error
	Engine, err = xorm.NewEngine("mysql", mysql_url)

	if err != nil {
		log.Println("创建数据库连接失败")
		return err
	}
	Engine.SetMapper(core.SameMapper{})
	Engine.ShowSQL = true   //在控制台打印生成的sql
	Engine.ShowDebug = true //在控制台打印调试信息
	Engine.ShowErr = true   //在控制台打印错误信息

	//	err = engine.CreateTables(&Customers{})
	//	if err != nil {
	//		log.Println("create table:", err.Error())
	//	}
	return nil
}
