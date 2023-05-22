package main

import "fmt"

// @Title  _9_memo
// @Description  MyGO
// @Author  WeiDa  2023/5/19 18:29
// @Update  WeiDa  2023/5/19 18:29

func main() {
	u := &User{MemoList: make([]*Memo, 0)}

	o := &Originator{}

	o.SetState("1")
	u.AddMemo(o.NewMemo(o.GetState()))

	o.SetState("2")
	u.AddMemo(o.NewMemo(o.GetState()))

	o.SetState("3")
	u.AddMemo(o.NewMemo(o.GetState()))

	fmt.Println(u.GetMemo(1)) //拿到第二条备忘录
}
