package main

import (
	"fmt"
	"hellogo/pipeline/pipeline"
)

/**
 * @Author: Lemyhello
 * @Description: Golang并发模型：流水线模型
 * @Question: 设计一个程序：计算一个整数切片中元素的平方值并把它打印出来。非并发的方式是使用for遍历整个切片，然后计算平方，打印结果。
 * @Version: X.X.X
 * @Date: 2020/3/30 上午10:48
 */

func main()  {
	fmt.Println("常规做法")
	pipeline.RoutineSquare(1,2,3,4,5,6,7,8,9,10,11,12,13,14)
	fmt.Println("流水线模型做法")
	//生产者
	pch := pipeline.Producer(1,2,3,4,5,6,7,8,9,10,11,12,13,14)
	//处理
	hCh := pipeline.Handle(pch)
	//消费
	pipeline.Consumer(hCh)
}