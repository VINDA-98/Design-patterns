package main

import (
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 11:55
// @Update  WeiDa  2023/5/15 11:55

type NoticeSys struct {
	Prefix string //发送消息前置内容
}

func NewNotice(prefix string) *NoticeSys {
	return &NoticeSys{Prefix: prefix}
}

func (n *NoticeSys) SendNotice(msg string) {
	log.Printf("[%s]:%s \n", n.Prefix, msg)
}
