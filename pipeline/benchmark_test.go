package main

import (
	"fmt"
	"hellogo/pipeline/pipeline"
	"testing"
)

/**
 * @Author: Lemyhello
 * @Description:
 * @File:  benchmark_test.go
 * @Version: X.X.X
 * @Date: 2020/3/30 上午11:15
 */

func Benchmark_Routine(b *testing.B) {
	fmt.Println("常规做法")
	pipeline.RoutineSquare(1,2,3,4,5,6,7,8,9,10,11,12,13,14)
}

func Benchmark_PHC(b *testing.B) {
	fmt.Println("流水线模型做法")
	//生产者
	pch := pipeline.Producer(1,2,3,4,5,6,7,8,9,10,11,12,13,14)
	//处理
	hCh := pipeline.Handle(pch)
	//消费
	pipeline.Consumer(hCh)
}