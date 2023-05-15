package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 10:46
// @Update  WeiDa  2023/5/15 10:46

type XiCha struct {
}

// GetPrice 得到一杯喜茶的基本价格
func (x *XiCha) GetPrice() int {
	return 25
}
