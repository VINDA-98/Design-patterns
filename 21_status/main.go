package main

import "log"

// @Title  _1_status
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:21
// @Update  WeiDa  2023/5/22 11:21

func main() {
	shop := NewShopMachine(2, 100)

	err := shop.AddItem(1)
	if err != nil {
		return
	}

	err = shop.ReqItem()
	if err != nil {
		return
	}

	log.Println(shop.InsertMoney(10))

	log.Println(shop.GetItem())

	err = shop.InsertMoney(100)
	if err != nil {
		return
	}

	err = shop.GetItem()
	if err != nil {
		return
	}

	log.Println("----------------------------------")

	err = shop.ReqItem()
	if err != nil {
		return
	}

	log.Println(shop.GetItem())

	err = shop.InsertMoney(200)
	if err != nil {
		return
	}

	err = shop.GetItem()
	if err != nil {
		return
	}
}
