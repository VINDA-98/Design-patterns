package main

import (
	"log"
)

// @Title  main
// @Description 实现Item具体的订阅细节步骤
// @Author  WeiDa  2023/5/18 18:07
// @Update  WeiDa  2023/5/18 18:07

type CanTee struct {
	Obs  []Observe
	Name string
}

func (c *CanTee) Register(o Observe) {
	c.Obs = append(c.Obs, o)
	log.Println("添加成功新的订阅账户:" + o.GetEmID())
}

func (c *CanTee) Deregister(o Observe) {
	for index, v := range c.Obs {
		if v.GetEmID() == o.GetEmID() {
			c.Obs = append(c.Obs[:index], c.Obs[index+1:]...)
			log.Println("Deregister Observe  ", o.GetEmID())
			return
		}
	}
}

func (c *CanTee) NoticeAll() {
	for _, v := range c.Obs {
		v.Update()
	}
}
