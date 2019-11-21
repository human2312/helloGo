//some basic trainning by Lemy 2019年11月21日20:31:13
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const constant string = "constant"

func main()  {
	s,i,f,b := str()
	rand.Seed(time.Now().Unix())
	randnum := rand.Intn(4)
	fmt.Println(s,i,f,b)
	fmt.Println(constant)
	fmt.Println(example(randnum))
	fmt.Println("done...")
}

//str some basic string
func str() (s string,i int,f float64,b bool) {
	s = "Lemy" + " learn golang."
	i = 1+2
	f = 7.0 / 3.1
	b = !true
	return
}

func example(i int) (result string)   {
	if i >= 0 && i <= 2 {
		switch i {
		case 1:
			result = "you chose "+ strconv.Itoa(i)
		case 2:
			result = "you chose mid num " + strconv.Itoa(i)
		default:
			result = "you can't chose " + strconv.Itoa(i)
		}
	} else if i == 3 {
		result = "you chose max value is " + strconv.Itoa(i)
	} else {
		result = "no variable need to back"
	}
	return
}