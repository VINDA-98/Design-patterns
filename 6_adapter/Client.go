package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/8 15:45
// @Update  WeiDa  2023/5/8 15:45

type Client struct {
}

// InsertUsbToComputer Client结构体实现插入计算机Usb方法
func (c *Client) InsertUsbToComputer(com Computer) {
	com.InsertUsb()
}
