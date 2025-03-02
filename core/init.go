package core

import (
	"database/sql"
	"sync"

	"transfy/config/database"
	"transfy/logs"
	//"github.com/redis/go-redis/v9"
)

var (
	once        sync.Once
	//RedisClient *redis.Client
	MysqlDb     *sql.DB
	ErroMysql   error
)

func InitConnection() {
	logs.Init()
	once.Do(func() {
		logs.Info("Init connexion")
		MysqlDb, ErroMysql = database.ConnectMysql()
		//RedisClient = database.ConnectRedis()

	})
}
