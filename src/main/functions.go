//接口体验
package main

import "fmt"

//Computer 定义一台计算机基本结构
type Computer struct {
	id       string
	mouse    string
	keyboard string
	screen   string
	cpu      string
	price    float64
}

//Order 一个订单基本操作
type Order interface {
	GetId() string
	Buy()
	Cancel()
}

//GetId 获取电脑编号
func (c Computer) GetId() string {
	return c.id
}

//Buy 下单
func (c Computer) Buy() {
	fmt.Println(c.GetId() + "电脑已经被下单")
}

//Cancel 取消订单
func (c Computer) Cancel() {
	fmt.Println(c.GetId() + "电脑订单已经被取消")
}

func main() {
	//赋值一台电脑
	mac := Computer{id: "1", mouse: "apple mouse", keyboard: "apple keyboard", screen: "15in", cpu: "intel i7", price: 15999.00}
	fmt.Println(mac.price)
	//实例化接口
	var o Order
	o = mac
	o.Buy()    //下单操作
	o.Cancel() //取消订单
}
