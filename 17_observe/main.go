package main

// @Title  _7_observe
// @Description  MyGO
// @Author  WeiDa  2023/5/18 17:57
// @Update  WeiDa  2023/5/18 17:57

func main() {
	o1 := &Employee{ID: "10001", Name: "VinDa1"}
	o2 := &Employee{ID: "10002", Name: "VinDa2"}
	o3 := &Employee{ID: "10002", Name: "VinDa3"}

	c := &CanTee{Name: " big CanTee "}
	c.Register(o1)
	c.Register(o2)
	c.Register(o3)

	// 通知所有人
	c.NoticeAll()
}
