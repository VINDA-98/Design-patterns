package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 12:17
// @Update  WeiDa  2023/5/16 12:17

type FIFO struct {
}

func (f *FIFO) Evict(c *Cache) {
	log.Println("FIFO evict cache")
}
