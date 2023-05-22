package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/19 18:30
// @Update  WeiDa  2023/5/19 18:30

type Memo struct {
	State string
}

func (m *Memo) GetMemoState() string {
	return m.State
}
