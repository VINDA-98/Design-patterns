package main

import "errors"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:27
// @Update  WeiDa  2023/5/22 11:27

type NoItem struct {
	Machine *ShopMachine
}

func (n *NoItem) AddItem(num int) error {
	err := n.Machine.IncItem(num)
	if err != nil {
		return err
	}
	n.Machine.SetState(n.Machine.HasItemState)

	return nil
}

func (n *NoItem) ReqItem() error {
	return errors.New("no item ... ")
}

func (n *NoItem) InsertMoney(int) error {
	return errors.New("no item ... ")
}

func (n *NoItem) GetItem() error {
	return errors.New("no item ... ")
}
