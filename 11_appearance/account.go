package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 11:55
// @Update  WeiDa  2023/5/15 11:55

type Account struct {
	Name string
}

func NewAccount(name string) *Account {
	return &Account{Name: name}
}

func (a *Account) CheckAccount(name string) bool {
	return a.Name == name
}
