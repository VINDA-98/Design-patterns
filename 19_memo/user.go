package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/19 18:31
// @Update  WeiDa  2023/5/19 18:31

import "fmt"

type User struct {
	MemoList []*Memo
}

func (u *User) AddMemo(m *Memo) {
	u.MemoList = append(u.MemoList, m)
	fmt.Println("AddMemo state : ", m.State)
}

func (u *User) GetMemo(index int) *Memo {
	if index < len(u.MemoList) {
		return u.MemoList[index]
	}
	return nil
}
