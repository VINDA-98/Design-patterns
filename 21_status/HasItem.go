package main

import (
	"errors"
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:26
// @Update  WeiDa  2023/5/22 11:26

type HasItem struct {
	Machine *ShopMachine
}

func (h *HasItem) AddItem(num int) error {
	err := h.Machine.IncItem(num)
	if err != nil {
		return err
	}
	return nil
}

func (h *HasItem) ReqItem() error {
	if h.Machine.ItemCnt == 0 {
		h.Machine.SetState(h.Machine.NoItemState)
		return errors.New("no item ... ")
	}
	log.Println("req item ... ")
	h.Machine.SetState(h.Machine.ReqItemState)
	return nil
}

func (h *HasItem) InsertMoney(int) error {
	return errors.New("please select item ... ")
}

func (h *HasItem) GetItem() error {
	return errors.New("please select item ... ")
}
