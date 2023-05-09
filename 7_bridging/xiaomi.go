package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/9 11:14
// @Update  WeiDa  2023/5/9 11:14

import "fmt"

type Xiaomi struct {
}

func (h *Xiaomi) Show(content string) {
	fmt.Printf("[%s] Xiaomi show ... \n", content)
}
