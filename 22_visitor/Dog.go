package main

import (
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:54
// @Update  WeiDa  2023/5/22 11:54

type Dog struct {
}

func (d *Dog) GetType() {
	log.Println(" Dog ... ")
}

func (d *Dog) Accept(v IVisitor) {
	v.VsDog(d) //调用IVisitor接口中跟狗狗有关的访问方法
}
