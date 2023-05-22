package main

import (
	"errors"
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:28
// @Update  WeiDa  2023/5/22 11:28

type ReqItem struct {
	Machine *ShopMachine
}

func (r *ReqItem) AddItem(int) error {
	return errors.New("item is processing ... ")
}

func (r *ReqItem) ReqItem() error {
	return errors.New("error in reqItem ")
}

func (r *ReqItem) InsertMoney(money int) error {
	log.Println("insert money ", money)
	if money < r.Machine.ItemPrice {
		return errors.New("not enough money ... ")
	}
	r.Machine.SetState(r.Machine.HasMoney)
	return nil
}

func (r *ReqItem) GetItem() error {
	return errors.New("please insert money ... ")
}
