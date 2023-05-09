package main

// @Title  __bridging
// @Description  MyGO
// @Author  WeiDa  2023/5/9 10:18
// @Update  WeiDa  2023/5/9 10:18

/* 桥接模式
可将业务逻辑或一个大类拆分为不同的层次结构， 从而能独立地进行开发
例如不同的显示器要接入计算机，那么计算机可以独立开发，显示器也可以独立开发
无论计算机里面有 windows 的，还是 mac 的 仍然支持设置 显示器
*/

func main() {
	h := &Honor{}
	x := &Xiaomi{}

	m := &Mac{Model: "mac-air"}
	w := &Win{Model: "R9000P"}

	m.SetDisplay(h)
	m.Show()

	m.SetDisplay(x)
	m.Show()

	w.SetDisplay(h)
	w.Show()

	w.SetDisplay(x)
	w.Show()
}
