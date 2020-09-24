package db

import (
	"database/sql"
	"github.com/Unknwon/goconfig"
	_ "github.com/bmizerany/pq"
	"log"
)

type DBOBJ struct {
	DB *sql.DB
}

type dber interface {
	ConnectDb() *sql.DB
	close()
	callPro(sql string)
}

func (this *DBOBJ) ConnectDb() *sql.DB {
	var (
		cfg      *goconfig.ConfigFile
		err      error
		psqlInfo string
	)

	if cfg, err = goconfig.LoadConfigFile("conf/ds.conf"); err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	if psqlInfo, err = cfg.GetValue("postgres", "db"); err != nil {
		log.Fatalf("无法获取数据库配置信息: %s", err)
	}
	if this.DB, err = sql.Open("postgres", psqlInfo); err != nil {
		log.Fatalf("无法打开数据库连接:%s", err)
	}
	if err = this.DB.Ping(); err != nil {
		log.Fatalf("数据库无法登录:%s", err)
	}
	return this.DB
}

func (this *DBOBJ) Close() {
	this.DB.Close()
}

func (this *DBOBJ) CallPro(sql string) {

	if _, err := this.DB.Exec(sql); err != nil {
		log.Print(err.Error())
	}

}
