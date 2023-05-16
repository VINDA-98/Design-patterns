package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 12:28
// @Update  WeiDa  2023/5/16 12:28

type EvictionAlgo interface {
	Evict(c *Cache) //EvictionAlgo 清理缓存算法
}
