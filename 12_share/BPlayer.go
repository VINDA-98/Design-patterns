package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 14:50
// @Update  WeiDa  2023/5/15 14:50

type BPlayer struct {
	Color string
}

func NewBPlayer() *BPlayer {
	return &BPlayer{Color: "blue"}
}

func (a *BPlayer) GetColor() string {
	return a.Color
}
