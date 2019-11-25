//协程 实操
package main

import (
	"sync"
)

var wg sync.WaitGroup

//定义一个执行方法
func sendMsg(str string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- str
	}
	wg.Done()
}

func main() {
	wg.Add(2) //等待两个协程结束

	//sendMsg("direct run")
	ch := make(chan string) //实例化一个管道

	go sendMsg("goroutine run", ch)

	go sendMsg("gotoutine run2", ch)

	i := 0

	for {
		i++
		str := <-ch
		println(str)

		if i >= 10 {
			close(ch)
			break
		}
	}

	wg.Wait()
}
