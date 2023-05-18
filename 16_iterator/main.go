package main

import (
	"log"
)

// @Title  _6_iterator
// @Description  MyGO
// @Author  WeiDa  2023/5/17 14:42
// @Update  WeiDa  2023/5/17 14:42

func main() {
	ua := UserArray{
		users: []*User{
			&User{Name: "VinDa", Age: 98},
			&User{Name: "VinDa", Age: 99},
			&User{Name: "VinDa", Age: 100},
			&User{Name: "VinDa", Age: 101},
		}}

	it := ua.createIterator()

	for it.HasNext() {
		log.Println(it.GetNext())
	}
}
