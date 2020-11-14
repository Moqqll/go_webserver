package mysql

import (
	"fmt"
	"go_webserver/setting"

	_ "github.com/go-sql-driver/mysql" //	mysql数据库驱动
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	dbConn *sqlx.DB
)

//Init ...
func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBname,
	)
	dbConn, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed, err:%v\n", zap.Error(err))
		return err
	}
	dbConn.SetMaxOpenConns(cfg.MaxOpenConns)
	dbConn.SetMaxIdleConns(cfg.MaxIdleConns)
	return nil
}

//Close ...
func Close() {
	_ = dbConn.Close()
}
