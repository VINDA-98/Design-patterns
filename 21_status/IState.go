package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:21
// @Update  WeiDa  2023/5/22 11:21

type IState interface {
	AddItem(int) error     //加货
	GetItem() error        //有货
	InsertMoney(int) error //投币
	ReqItem() error        //请求商品信息
}
