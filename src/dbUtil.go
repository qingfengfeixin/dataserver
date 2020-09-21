package main

import (
	"database/sql"
	"github.com/Unknwon/goconfig"
	_ "github.com/bmizerany/pq"
	"log"
)

func connectDB() *sql.DB {
	var (
		cfg      *goconfig.ConfigFile
		err      error
		psqlInfo string
		db       *sql.DB
	)

	if cfg, err = goconfig.LoadConfigFile("conf/ds.conf"); err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	if psqlInfo, err = cfg.GetValue("postgres", "db"); err != nil {
		log.Fatalf("无法获取数据库配置信息: %s",err)
	}
	if db, err = sql.Open("postgres", psqlInfo); err != nil {
		log.Fatalf("无法打开数据库连接:%s",err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("数据库无法登录:%s",err)
	}
	return db
}

func callPro(db *sql.DB,sql string) {

	if _, err := db.Exec(sql); err != nil {
		log.Print(err.Error())
	}

}
