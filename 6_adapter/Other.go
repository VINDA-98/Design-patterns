package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/8 15:51
// @Update  WeiDa  2023/5/8 15:51

type Other struct {
}

func (o *Other) InsertTypeC() {
	log.Println("insert type C ... ")
}

func (o *Other) InsertHDMI() {
	log.Println("insert HDMI ... ")
}
