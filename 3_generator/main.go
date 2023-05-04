package main

import (
	"log"
)

// @Title  __generator
// @Description  MyGO
// @Author  WeiDa  2023/5/4 14:33
// @Update  WeiDa  2023/5/4 14:33

//  建一座建筑，无论你是冰室，还是住房，大体会经历如下步骤
// 1. 搭建框架
// 2. 装修
// 3. 安装门
// 4. 安装床

func main() {
	//构建冰室
	b := GetBuild("ice")
	if b == nil {
		log.Println("not support bType ...")
		return
	}
	//提供管理步骤
	m := NewManager(b)
	iceHouse := m.GetHouse()
	log.Println("iceHouse == ", iceHouse)

	b2 := GetBuild("wood")
	if b2 == nil {
		log.Println("not support bType ...")
		return
	}
	m2 := NewManager(b2)
	woodHouse := m2.GetHouse()
	log.Println("woodHouse == ", woodHouse)

	//设置为构建冰室
	m2.SetWorker(b)
	iHouse1 := m2.GetHouse()
	log.Println("iHouse1 == ", iHouse1)

}
