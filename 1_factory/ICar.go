package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 10:49
// @Update  WeiDa  2023/5/4 10:49

// ICar 实现的接口
type ICar interface {
	Run()
	SetName(string)
	GetName() string
	SetPrice(int64)
	GetPrice() int64
}
