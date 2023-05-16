package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 12:17
// @Update  WeiDa  2023/5/16 12:17

type Lfu struct {
}

func (f *Lfu) Evict(c *Cache) {
	log.Println("Lfu evict cache")
}
