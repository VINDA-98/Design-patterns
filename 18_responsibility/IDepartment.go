package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/19 10:26
// @Update  WeiDa  2023/5/19 10:26

type IDepartment interface {
	Execute(e *Employee)
	SetNext(d IDepartment)
}
