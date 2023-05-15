package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 14:50
// @Update  WeiDa  2023/5/15 14:50

type Player struct {
	pType string
	dress Dress
}

func NewPlayer(pType string) *Player {
	// 得到服务实例，并且传入需要得到的服装类型
	d, _ := GetDressInstance().GetDressType(pType)
	//if d == nil{
	//
	//}
	//
	return &Player{
		pType: pType,
		dress: d,
	}
}
