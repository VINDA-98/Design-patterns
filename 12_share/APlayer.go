package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 14:50
// @Update  WeiDa  2023/5/15 14:50

type APlayer struct {
	Color string
}

func NewAPlayer() *APlayer {
	return &APlayer{Color: "red"}
}

func (a *APlayer) GetColor() string {
	return a.Color
}
