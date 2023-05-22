package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 10:44
// @Update  WeiDa  2023/5/22 10:44

type Manager struct {
	Name string
	Mgr  IOrderMgr
}

func (m *Manager) Enter() {
	// 判断是否能进入办公室
	if !m.Mgr.CanEnter(m) {
		log.Println(" [Block]queue block , " + m.Name + " 排队成功 ... ")
		return
	}
	log.Println(m.Name + " 已在队列中 ... ")
}
func (m *Manager) Leave() {
	log.Println(m.Name + " 离开队列或办公室 ")
	m.Mgr.NoticePeople()
}

func (m *Manager) PermitEnter() {
	log.Println(m.Name + " 进入办公室  ... ")
	m.Enter()
}
