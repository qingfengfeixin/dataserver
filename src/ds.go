package main

import (
	"log"
)

func proRun() {
	jobs := &job{}

	var rows []job
	rows = jobs.queryJob()

	//执行存储过程
	for _, j1 := range rows {
		job := j1

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
