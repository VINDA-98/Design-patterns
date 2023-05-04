package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 14:55
// @Update  WeiDa  2023/5/4 14:55

type Manager struct {
	Build IBuild
}

func NewManager(build IBuild) *Manager {
	return &Manager{Build: build}
}
func (m *Manager) SetWorker(b IBuild) {
	m.Build = b
}

// GetHouse 设置房子构建步骤
func (m *Manager) GetHouse() *House {
	m.Build.SetFrame()
	m.Build.SetStyle()
	m.Build.SetDoor()
	m.Build.SetBed()
	return m.Build.BuildHouse()
}
