package main

import (
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/8 15:53
// @Update  WeiDa  2023/5/8 15:53

type Adapter struct {
	OtherSvr *Other
}

func (a *Adapter) InsertUsb() {
	log.Println("**********  Convert TypeC interface to Usb ... *********")
	a.OtherSvr.InsertTypeC()
	log.Println("**********  Convert HDMI interface to Usb ... *********")
	a.OtherSvr.InsertHDMI()
}
