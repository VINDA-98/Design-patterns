package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/9 11:14
// @Update  WeiDa  2023/5/9 11:14

import "fmt"

type Honor struct {
}

// Show 实现IDisplay接口，继承显示
func (h *Honor) Show(content string) {
	fmt.Printf("[%s] Honor show ... \n", content)
}
