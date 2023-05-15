package main

import "log"

func main() {
	xiCha := &XiCha{}
	log.Println("喜茶惊爆价:", xiCha.GetPrice())

	//加芋圆的价格
	yy := &YuYuan{MilkTea: xiCha}
	log.Println("芋圆喜茶惊爆价:", yy.GetPrice())

	//加芋圆牛奶的价格
	nn := &Milk{MilkTea: yy}
	log.Println("芋圆牛奶喜茶惊爆价:", nn.GetPrice())

	//加牛奶的价格
	nn.MilkTea = xiCha
	log.Println("芋圆牛奶喜茶惊爆价:", nn.GetPrice())

}
