package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/17 12:12
// @Update  WeiDa  2023/5/17 12:12

type IBasic interface {
	GetMsg() string              //获得消息
	SaveMsg(string)              //保存消息
	GenerateReply(string) string //生成回复
	SendMsg(string)              //发送消息
}

type Mgr struct {
	B IBasic
}

func (m *Mgr) GetAndSendMsg() {
	msg := m.B.GetMsg()
	m.B.SaveMsg(msg)
	reply := m.B.GenerateReply(msg)
	m.B.SendMsg(reply)
}
