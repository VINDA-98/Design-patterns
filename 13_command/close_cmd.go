package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 11:30
// @Update  WeiDa  2023/5/16 11:30

type CloseCmd struct {
	Dev Device
}

func (o *CloseCmd) Execute() {
	o.Dev.Close()
}
