package main

import (
	"log"
)

func proRun() {
	db := &dbObj{}
	db.connectDB()
	defer db.close()

	var rows []job
	rows = queryJob(db)

	//执行存储过程
	for _, job := range rows {

		go func() {

			dbf := &dbObj{}
			dbf.connectDB()
			defer dbf.close()

			//打印日志
			log.Printf("开始执行存储过程 %s", job.what)

			job.setStat(dbf, 1)

			// 执行存储过程
			dbf.callPro(job.what)
			// 执行之后更新该job的nexttime
			job.setNextTime(dbf)
			// fmt.Printf("now is %s next time is %s", time.Now(), job.nexttime)

			job.setStat(dbf, 0)

		}()
	}

}
