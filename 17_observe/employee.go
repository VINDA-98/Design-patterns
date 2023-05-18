package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/18 18:01
// @Update  WeiDa  2023/5/18 18:01

type Employee struct {
	Name string
	ID   string
}

// Update 订阅消息更新给员工
func (e *Employee) Update() {
	log.Println("Send sms to ", e.Name)
}

func (e *Employee) GetEmID() string {
	log.Printf("员工%s 的ID为: %s \n", e.Name, e.ID)
	return e.ID
}
