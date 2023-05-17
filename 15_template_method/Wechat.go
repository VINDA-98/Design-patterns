package main

import (
	"fmt"
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/17 12:16
// @Update  WeiDa  2023/5/17 12:16

type Wechat struct {
	AutoMsg string
}

func (w *Wechat) GetMsg() string {
	return w.AutoMsg
}

func (w *Wechat) SaveMsg(string) {
	log.Printf("[%s] save msg to cache ...  \n", "wechat")
}

func (w *Wechat) GenerateReply(msg string) string {
	return fmt.Sprintf("[%s] %s", "wechat", msg)
}

func (w *Wechat) SendMsg(msg string) {
	log.Printf("send msg:[%s] to %s ... \n", msg, "wechat")
}
