package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 10:36
// @Update  WeiDa  2023/5/15 10:36

type Milk struct {
	MilkTea IMilkTea
}

// GetPrice 加了牛奶的价格
func (m *Milk) GetPrice() int {
	return m.MilkTea.GetPrice() + 5
}
