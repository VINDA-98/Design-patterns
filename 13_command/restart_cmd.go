package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 11:30
// @Update  WeiDa  2023/5/16 11:30

type RestartCmd struct {
	Dev Device
}

func (o *RestartCmd) Execute() {
	o.Dev.Restart()
}
