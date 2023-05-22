package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 10:29
// @Update  WeiDa  2023/5/22 10:29

// IOrderMgr 医院服务者接口
type IOrderMgr interface {
	NoticePeople()                //通知病人的方法
	CanEnter(people IPeople) bool //是否允许进入办公室队列方法
}
