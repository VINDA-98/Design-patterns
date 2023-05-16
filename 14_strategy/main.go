package main

// @Title  _4_strategy
// @Description  MyGO
// @Author  WeiDa  2023/5/16 12:17
// @Update  WeiDa  2023/5/16 12:17

//LRU（Least Recently Used）：清理最近使用的最少的那些缓存数据
//FIFO（First In First Out）：清理最早放入缓存的那些数据
//LFU（Least Frequently Used）：清理使用频率最低的那部分数据

func main() {
	cache := InitCache(&Lru{})
	cache.Add("VinDa", "9898")
	cache.Evict()

	cache.SetEvi(&FIFO{})
	cache.Add("VinDa", "9800")
	cache.Evict()

	cache.SetEvi(&Lfu{})
	cache.Add("VinDa", "9811")
	cache.Add("VinDa", "9822")
	cache.Evict()

}
