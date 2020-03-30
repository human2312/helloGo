package pipeline

import "fmt"

/**
 * @Author: Lemyhello
 * @Description:
 * @File:  pipeline
 * @Version: X.X.X
 * @Date: 2020/3/30 上午11:14
 */

//RoutineSquare 常规平方
func RoutineSquare(intslice ...int)  {
	for _,v := range intslice {
		fmt.Println(v * v)
	}
}

//Producer 生产者
func Producer(intslice ...int) (chan int) {
	res := make(chan int)
	go func() {
		defer close(res)
		for _,v := range intslice {
			res <- v
		}
	}()
	return res
}

//Handle 执行
func Handle(intCh chan int) (chan int) {
	res := make(chan int)
	go func() {
		defer close(res)
		for v := range intCh {
			res <- (v * v)
		}
	}()
	return res
}

//Consumer 消费者
func Consumer(outCh chan int)  {
	for v := range outCh {
		fmt.Println(v)
	}
}