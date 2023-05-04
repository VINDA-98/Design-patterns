package main

import "log"

// @Title  __factory
// @Description  MyGO
// @Author  WeiDa  2023/5/4 10:42
// @Update  WeiDa  2023/5/4 10:42

// 写一个工厂，可以生产车
func main() {
	hhGz, err := NewShop("HHGZ")
	if err != nil {
		log.Print(err)
		return
	}
	hhGzCloth := hhGz.MakeCloth()
	log.Printf("Cloth brand:[%s],color:[%s],size[%s]\n\n", hhGzCloth.GetBrand(), hhGzCloth.GetColor(), hhGzCloth.GetSize())

	hxEk, _ := NewShop("HXEK")
	hxEkPant := hxEk.MakePant()
	log.Printf("Pant brand:[%s],color:[%s],size[%s]\n\n", hxEkPant.GetBrand(), hxEkPant.GetColor(), hxEkPant.GetSize())

}
