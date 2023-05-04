package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 11:40
// @Update  WeiDa  2023/5/4 11:40

// HH GZ 花花公子
type HHGZ struct {
}

func (*HHGZ) MakeCloth() ICloth {
	return &Cloth{
		Brand: "花花公子品牌服装",
		Color: "花色",
		Size:  "M",
	}
}

func (*HHGZ) MakePant() IPant {
	return &Pant{
		Brand: "花花公子品牌裤子",
		Color: "蓝色",
		Size:  "M",
	}
}
