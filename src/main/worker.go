package main

import (
	"fmt"
	"time"
)

//工作池
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	num := 1000     //任务数
	worknum := 4000 //进程数

	//启动三个worker工作池
	for w := 1; w <= worknum; w++ {
		go worker(w, jobs, results)
	}

	//发送9个jobs任务
	for j := 1; j <= num; j++ {
		jobs <- j
	}
	close(jobs)

	//输出jobs任务结果
	for a := 1; a <= num; a++ {
		<-results
	}
}
