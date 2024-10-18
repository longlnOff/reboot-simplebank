package initialize

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
	"github.com/longln/reboot-simplebank/global"
)

func InitDatabase() {
	dataSource := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		global.Config.DatabaseConfig.Engine,
		global.Config.DatabaseConfig.Info.UserName,
		global.Config.DatabaseConfig.Info.Password,
		global.Config.DatabaseConfig.Info.Host,
		global.Config.DatabaseConfig.Info.Port,
		global.Config.DatabaseConfig.Info.DatabaseName,
	)

	Conn, err := sql.Open(global.Config.DatabaseConfig.Driver, dataSource)
	if err != nil {
		panic(err)
	}
	global.Database = Conn
	SetPool()
	global.Logger.Info("init database success")
}

func SetPool() {
	global.Database.SetConnMaxLifetime(time.Duration(global.Config.DatabaseConfig.Info.ConnMaxLifetime))
	global.Database.SetMaxIdleConns(global.Config.DatabaseConfig.Info.MaxIdleConns)
	global.Database.SetMaxOpenConns(global.Config.DatabaseConfig.Info.MaxOpenConns)
}