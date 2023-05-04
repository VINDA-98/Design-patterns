package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 11:40
// @Update  WeiDa  2023/5/4 11:40

// HXEK 鸿星尔克
type HXEK struct {
}

func (*HXEK) MakeCloth() ICloth {
	return &Cloth{
		Brand: "鸿星尔克品牌服装",
		Color: "黑色",
		Size:  "M",
	}
}

func (*HXEK) MakePant() IPant {
	return &Pant{
		Brand: "鸿星尔克品牌裤子",
		Color: "橙色",
		Size:  "M",
	}
}
