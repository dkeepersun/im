package db

import (
	"database/sql"
	"im/config"
	"im/pkg/logger"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

var DBCli *sql.DB

func init() {
	var err error
	DBCli, err = sql.Open("mysql", config.Logic.MySQL)
	if err != nil {
		logger.Sugar.Error("mysql err ", zap.Error(err))
	}
}

func InitMysql(dataSourceName string) {
	var err error
	DBCli, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
