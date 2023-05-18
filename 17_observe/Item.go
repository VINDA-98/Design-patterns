package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/18 18:05
// @Update  WeiDa  2023/5/18 18:05

type Item interface { //订阅的事件行为接口
	Register(Observe)   //注册
	Deregister(Observe) //注销
	NoticeAll()         //通知所有订阅者
}
