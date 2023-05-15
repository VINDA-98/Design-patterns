package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 17:26
// @Update  WeiDa  2023/5/15 17:26

type Game struct {
	APlayers []*Player
	BPlayers []*Player
}

func NewGame() *Game {
	return &Game{
		APlayers: make([]*Player, 0, 2),
		BPlayers: make([]*Player, 0, 2),
	}
}

func (g *Game) AddAPlayer() {
	g.APlayers = append(g.APlayers, NewPlayer(APlay))
}

func (g *Game) AddBPlayer() {
	g.BPlayers = append(g.BPlayers, NewPlayer(BPlay))
}
