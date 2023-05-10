package main

import (
	"log"
)

// @Title  __proxy
// @Description  MyGO
// @Author  WeiDa  2023/5/10 15:55
// @Update  WeiDa  2023/5/10 15:55

func main() {
	ng := NewNginx()
	log.Println(ng.HandlerReq("/get_info", "GET"))
	log.Println(ng.HandlerReq("/edit_info", "POST"))
	log.Println(ng.HandlerReq("/edit_info", "POST")) //限流访问，只允许当前请求1次
	log.Println(ng.HandlerReq("/edit_info1", "POST"))
}
