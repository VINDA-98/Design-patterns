package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 10:44
// @Update  WeiDa  2023/5/22 10:44

type Programmer struct {
	Name string
	Mgr  IOrderMgr
}

func (p *Programmer) Enter() {
	if !p.Mgr.CanEnter(p) {
		log.Println(" [Block]queue block , " + p.Name + " 排队成功 ... ")
		return
	}
	log.Println(p.Name + " 已在队列中 ... ")
}

func (p *Programmer) Leave() {
	log.Println(p.Name + " 离开队列或办公室 ")
	p.Mgr.NoticePeople()
}

func (p *Programmer) PermitEnter() {
	log.Println(p.Name + " 进入办公室  ... ")
	p.Enter()
}
