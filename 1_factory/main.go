package main

import "fmt"

// @Title  __factory
// @Description  MyGO
// @Author  WeiDa  2023/5/4 10:42
// @Update  WeiDa  2023/5/4 10:42

// 写一个工厂，可以生产车
func main() {
	byd, _ := getCar("BYD")
	fmt.Println(byd.GetName(), byd.GetPrice())
	byd.Run()

}
