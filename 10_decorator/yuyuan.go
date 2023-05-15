package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 10:36
// @Update  WeiDa  2023/5/15 10:36

type YuYuan struct {
	MilkTea IMilkTea
}

// GetPrice 加了芋圆的价格
func (m *YuYuan) GetPrice() int {
	return m.MilkTea.GetPrice() + 2
}
