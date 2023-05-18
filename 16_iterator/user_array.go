package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/17 14:56
// @Update  WeiDa  2023/5/17 14:56

type UserArray struct {
	users []*User
}

func (u *UserArray) createIterator() Iterator {
	return &UserIterator{Users: u.users}
}
