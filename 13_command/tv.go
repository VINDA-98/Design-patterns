package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 11:34
// @Update  WeiDa  2023/5/16 11:34

type TV struct {
	Switch string
}

func (t *TV) Open() {
	t.Switch = "enable"
	log.Println("tv on ... ")
}

func (t *TV) Close() {
	t.Switch = "enable"
	log.Println("tv close ... ")
}

func (t *TV) Restart() {
	t.Switch = "restart"
	log.Println("tv restart ... ")
}
