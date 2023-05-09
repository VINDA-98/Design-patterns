package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/9 11:10
// @Update  WeiDa  2023/5/9 11:10

type Win struct {
	Model string
	Dis   IDisplay
}

// SetDisplay 设置win的显示器
func (m *Win) SetDisplay(dis IDisplay) {
	m.Dis = dis
}

func (m *Win) Show() {
	m.Dis.Show("win " + m.Model)
}
