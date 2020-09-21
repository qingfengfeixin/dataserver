package main

import (
	"time"
)

func init() {
	//程序启动修改所有任务的状态字段为0
	job := &job{}
	job.intiJob()
}

func main() {

	//每隔一段时间（1 min）执行一次
	for {
		proRun()
		// 睡眠 1min（不让它占用过多cpu）
		select {
		case <-time.NewTimer(60 * 1000 * time.Millisecond).C: //将在1 min 可读，返回
		}
	}

}
