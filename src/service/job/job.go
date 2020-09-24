package job

import (
	"github.com/gorhill/cronexpr"
	"log"
	"time"
	"utils/db"
)

type JOB struct {
	jobno    int
	nexttime time.Time
	interval string
	what     string
	stat     int
}

type jober interface {
	queryJob(db *db.DBOBJ)
	setNextTime(db *db.DBOBJ)
	setStat(db *db.DBOBJ, stat int)
}

func  IntiJob(d *db.DBOBJ) {
	// 设置任务状态为0

	if _, err := d.DB.Exec("update ds_job set stat = $1", 0); err != nil {
		log.Println(err)
	}
}

func queryJob(db *db.DBOBJ) []JOB {

	now := time.Now()

	// 查询需要执行的存储过程,nexttime 过期的任务
	rows, _ := db.DB.Query("SELECT jobno,nexttime,interval,what FROM ds_job where stat=0 and nexttime <$1", now)
	defer rows.Close()
	var jobs []JOB
	//执行存储过程
	for rows.Next() {

		job := &JOB{}
		rows.Scan(&job.jobno, &job.nexttime, &job.interval, &job.what)
		jobs = append(jobs, *job)
	}
	return jobs

}

func (this *JOB) setNextTime(db *db.DBOBJ) {
	this.nexttime = getNextTime(this.interval)

	stmt, err := db.DB.Prepare("update ds_job set nexttime=$1 WHERE jobno=$2")
	if err != nil {
		log.Println("fail to update:%v", err)
	}
	if _, err = stmt.Exec(this.nexttime, this.jobno); err != nil {
		log.Println(err)
	}

}

func (this *JOB) setStat(db *db.DBOBJ, stat int) {
	this.stat = stat

	stmt, err := db.DB.Prepare("update ds_job set stat=$1 WHERE jobno=$2")
	if err != nil {
		log.Println("fail to update:%v", err)
	}
	if _, err = stmt.Exec(this.stat, this.jobno); err != nil {
		log.Println(err)
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
