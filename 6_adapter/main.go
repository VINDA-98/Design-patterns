package main

// @Title  __adapter
// @Description  MyGO
// @Author  WeiDa  2023/5/8 15:44
// @Update  WeiDa  2023/5/8 15:44

func main() {
	c := &Client{}

	m := &MacComputer{}
	c.InsertUsbToComputer(m)

	// type C 的接口 通过 适配器 转换成可以通过 USB 插入到 电脑中
	oth := &Other{}
	ad := &Adapter{OtherSvr: oth}
	c.InsertUsbToComputer(ad)
}
