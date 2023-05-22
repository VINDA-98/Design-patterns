package main

import (
	"errors"
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:27
// @Update  WeiDa  2023/5/22 11:27

type HasMoney struct {
	Machine *ShopMachine
}

func (h *HasMoney) AddItem(int) error {
	return errors.New("item is processing ... ")
}
func (h *HasMoney) ReqItem() error {
	return errors.New("item is processing ... ")
}
func (h *HasMoney) InsertMoney(int) error {
	return errors.New("item is processing ... ")
}
func (h *HasMoney) GetItem() error {
	log.Println("get item ...")
	h.Machine.ItemCnt -= 1
	if h.Machine.ItemCnt == 0 {
		h.Machine.SetState(h.Machine.NoItemState)
	} else {
		h.Machine.SetState(h.Machine.HasItemState)
	}
	return nil
}
