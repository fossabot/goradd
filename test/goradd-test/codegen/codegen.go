package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/goradd/goradd/codegen/generator"
	"github.com/goradd/goradd/pkg/orm/db"
	_ "github.com/goradd/goradd/pkg/page/control/generator"
	_ "github.com/goradd/goradd/pkg/bootstrap/generator"
	_ "goradd-project/config" // Initialize required variables
	_ "goradd-tmp/template"
)

func main() {
	config()
	generator.Generate()
}

func config() {
	initDatabases()

}

func initDatabases() {

	cfg := mysql.NewConfig()

	cfg.DBName = "goradd"
	cfg.User = "root"
	cfg.Passwd = "12345"
	cfg.ParseTime = true

	key := "goradd"

	db1 := db.NewMysql5(key, "", cfg)

	db.AddDatabase(db1, key)

	db.AnalyzeDatabases()

}