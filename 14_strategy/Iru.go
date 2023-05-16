package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 12:17
// @Update  WeiDa  2023/5/16 12:17

type Lru struct {
}

func (f *Lru) Evict(c *Cache) {
	log.Println("Lru evict cache")
}
