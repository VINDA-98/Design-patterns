package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 16:57
// @Update  WeiDa  2023/5/4 16:57

type INode interface {
	Show(string)
	Clone() INode
}
