package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/17 14:45
// @Update  WeiDa  2023/5/17 14:45

type Iterator interface {
	HasNext() bool
	GetNext() *User
}
