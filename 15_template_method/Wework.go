package main

import (
	"fmt"
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/17 12:16
// @Update  WeiDa  2023/5/17 12:16

type Wework struct {
	AutoMsg string
}

func (w *Wework) GetMsg() string {
	return w.AutoMsg
}

func (w *Wework) SaveMsg(string) {
	log.Printf("[%s] save msg to cache ...  \n", "wechat")
}

func (w *Wework) GenerateReply(msg string) string {
	return fmt.Sprintf("[%s] %s", "wechat", msg)
}

func (w *Wework) SendMsg(msg string) {
	log.Printf("send msg:[%s] to %s ... \n", msg, "wechat")
}
