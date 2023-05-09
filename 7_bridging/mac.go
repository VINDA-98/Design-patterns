package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/9 11:10
// @Update  WeiDa  2023/5/9 11:10

type Mac struct {
	Model string
	Dis   IDisplay
}

func (m *Mac) SetDisplay(dis IDisplay) {
	m.Dis = dis
}

func (m *Mac) Show() {
	m.Dis.Show("mac " + m.Model)
}
