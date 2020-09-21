package main

import (
	"github.com/gorhill/cronexpr"
	"log"
	"time"
)

func proRun() {
	db := connectDB()
	now := time.Now()

	// 查询需要执行的存储过程,nexttime 过期的任务
	rows, _ := db.Query("SELECT jobno,nexttime,interval,what FROM ds_job where nexttime <$1", now)

	//执行存储过程
	for rows.Next() {
		var (
			jobno    int
			nexttime time.Time
			interval string
			what     string
		)

		rows.Scan(&jobno, &nexttime, &interval, &what)

		go func() {
			//打印日志
			log.Printf("开始执行存储过程 %s", what)

			//fmt.Println(jobno)
			//fmt.Println(nexttime)
			//fmt.Println(interval)
			//fmt.Println(what)

			// 执行存储过程
			callPro(db, what)
			// 执行之后更新该job的nexttime
			nexttime = getNextTime(interval)
			// fmt.Printf("now is %s next time is %s", time.Now(), nexttime)
			stmt, err := db.Prepare("update ds_job set nexttime=$1 WHERE jobno=$2")
			if err != nil {
				log.Println("fail to update:%v", err)
			}
			if _, err = stmt.Exec(nexttime, jobno); err != nil {
				log.Println(err)
			}

		}()

	}

}

func getNextTime(interval string) time.Time {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)
	now = time.Now()
	if expr, err = cronexpr.Parse(interval); err != nil {
		log.Printf("任务日期格式错误", interval)
		// 如果日期格式错误，则设置任务的nexttime为一个比较大的时间
		nextTime, _ = time.Parse("2006-01-02 15:04:05", "2999-12-31 00:00:00")
		return nextTime
	}

	nextTime = expr.Next(now)
	return nextTime
}
