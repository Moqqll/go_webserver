package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //	mysql数据库驱动
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	dbConn *sqlx.DB
)

//Init ...
func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	dbConn, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed, err:%v\n", zap.Error(err))
		return err
	}
	dbConn.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	dbConn.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return nil
}

//Close ...
func Close() {
	_ = dbConn.Close()
}
