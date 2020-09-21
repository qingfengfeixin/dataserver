package main

import (
	"fmt"
	"testing"
)

func testDb(ch chan int) {
	db := connectDB()
	callPro(db,"call sp_t2()")
	ch <- 1
}
func Testdb(t *testing.T){
	ch := make(chan int)

	go testDb(ch)

	for data := range ch {
		// 打印通道数据
		fmt.Println(data)

		if data == 1 {
			//break
		}
	}

}

