package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:24
// @Update  WeiDa  2023/5/22 11:24

type ShopMachine struct {
	HasItemState IState
	NoItemState  IState
	ReqItemState IState
	HasMoney     IState

	CurState IState

	ItemCnt   int
	ItemPrice int
}

func NewShopMachine(itemCnt, itemPrice int) *ShopMachine {
	s := &ShopMachine{ItemCnt: itemCnt, ItemPrice: itemPrice}

	h := &HasItem{Machine: s}
	n := &NoItem{Machine: s}
	r := &ReqItem{Machine: s}
	has := &HasMoney{Machine: s}

	s.HasItemState = h
	s.NoItemState = n
	s.ReqItemState = r
	s.HasMoney = has

	s.CurState = s.HasItemState

	return s
}

func (s *ShopMachine) SetState(state IState) {
	s.CurState = state
}

func (s *ShopMachine) ReqItem() error {
	return s.CurState.ReqItem()
}

func (s *ShopMachine) AddItem(itemNum int) error {
	return s.CurState.AddItem(itemNum)
}

func (s *ShopMachine) IncItem(itemNum int) error {
	s.ItemCnt += itemNum
	log.Printf("add %d item\n", itemNum)
	return nil
}

func (s *ShopMachine) InsertMoney(money int) error {
	return s.CurState.InsertMoney(money)
}

func (s *ShopMachine) GetItem() error {
	return s.CurState.GetItem()
}
