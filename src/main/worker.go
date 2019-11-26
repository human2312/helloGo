package main

import (
	"fmt"
	"time"
)

//工作池
func worker(id int,jobs <- chan int ,result chan<- int)  {
	for j := range jobs {
		fmt.Println("worker",id,"processing job",j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main()  {
	jobs := make(chan int,100)
	results := make(chan int,100)

	//启动三个worker工作池
	for w:=1;w <= 3;w++ {
		go worker(1,jobs,results)
	}

	//发送9个jobs任务
	for j := 1;j <=9 ;j++{
		jobs <- j
	}
	close(jobs)

	//输出jobs任务结果
	for a := 1;a <= 9; a++ {
		<-results
	}
}
