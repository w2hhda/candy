package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/go-redis/redis"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	user := beego.AppConfig.String("sqluser")
	password := beego.AppConfig.String("sqlpass")
	dbName := beego.AppConfig.String("sqldb")
	orm.RegisterDataBase("default", "mysql", user+":"+password+"@/"+dbName+"?charset=utf8")

	orm.RegisterModel(new(Candy), new(Record), new(Game), new(Token), new(User), new(GameCandy))
}

func DB() *sql.DB {
	user := beego.AppConfig.String("sqluser")
	password := beego.AppConfig.String("sqlpass")
	host := beego.AppConfig.String("host")
	dbName := beego.AppConfig.String("sqldb")

	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+host+":3306)/"+dbName)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis_addr"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
