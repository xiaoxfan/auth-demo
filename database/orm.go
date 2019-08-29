package database

import (
	"auth-demo/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"time"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = conn()
	if err != nil {
		panic(err)
	}
}
func conn() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(config.Config.Database.Driver, config.Config.DBUrl())
	if err != nil {
		panic(err)
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "")
	engine.SetTableMapper(tbMapper)
	engine.SetColumnMapper(core.GonicMapper{})
	location, _ := time.LoadLocation("Asia/Shanghai")
	engine.TZLocation = location
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	//注册动态SQL模板配置，可选功能，如应用中无需使用SqlTemplate，可无需初始化
	//此处注册动态SQL模板配置，使用Pongo2模板引擎，配置文件根目录为"../sql"，配置文件后缀为".stpl"
	err = engine.RegisterSqlTemplate(xorm.Pongo2("../sql", ".stpl"))
	if err != nil {
		panic(err)
	}
	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	err = engine.StartFSWatcher()
	if err != nil {
		fmt.Println(err)
	}
	return engine, nil
}
