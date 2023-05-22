package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:52
// @Update  WeiDa  2023/5/22 11:52

type IAnimal interface {
	GetType()
	Accept(v IVisitor) //同意访问方法
}
