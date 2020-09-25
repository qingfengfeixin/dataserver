package job

import (
	"log"
	"utils/db"
)

func ProRun() {
	d := &db.DBOBJ{}
	d.ConnectDb()
	defer d.Close()

	var jobs []JOB
	jobs = queryJob(d)

	//执行存储过程
	for _, j := range jobs {

		go func(job JOB) {

			dbf := &db.DBOBJ{}
			dbf.ConnectDb()
			defer dbf.Close()

			//打印日志
			log.Printf("开始执行存储过程 %s", job.what)

			job.setStat(dbf, 1)

			// 执行存储过程
			dbf.CallPro(job.what)
			// 执行之后更新该job的nexttime
			job.setNextTime(dbf)
			// fmt.Printf("now is %s next time is %s", time.Now(), job.nexttime)

			job.setStat(dbf, 0)
			//打印日志
			log.Printf("存储过程 %s 执行结束", job.what)

		}(j)
	}

}
