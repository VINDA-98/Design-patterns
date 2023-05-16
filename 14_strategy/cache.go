package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 12:18
// @Update  WeiDa  2023/5/16 12:18

type Cache struct {
	Storage map[string]string
	E       EvictionAlgo
	Cap     int
	MaxCap  int
}

func InitCache(e EvictionAlgo) *Cache {
	return &Cache{
		Storage: make(map[string]string, 0),
		E:       e,
		Cap:     0,
		MaxCap:  3,
	}
}

func (c *Cache) Evict() {
	c.E.Evict(c)
	c.Cap--
}

func (c *Cache) Add(key, value string) {
	if c.Cap >= c.MaxCap {
		c.Evict()
	}
	c.Cap++
	c.Storage[key] = value
	log.Printf("Add cache key :%s, value: %s \n", key, value)
}

func (c *Cache) SetEvi(e EvictionAlgo) {
	c.E = e
}
