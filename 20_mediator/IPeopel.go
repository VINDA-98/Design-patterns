package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 10:31
// @Update  WeiDa  2023/5/22 10:31

type IPeople interface {
	Enter()       //进入
	Leave()       //离开
	PermitEnter() //许可进入
}
