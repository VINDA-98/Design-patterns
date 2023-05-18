package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/18 18:00
// @Update  WeiDa  2023/5/18 18:00

// Observe 订阅接口
type Observe interface {
	Update()
	GetEmID() string //得到员工ID
}
