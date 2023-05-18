package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/17 14:49
// @Update  WeiDa  2023/5/17 14:49

type User struct {
	Name string
	Age  int
}

type UserIterator struct {
	Users []*User //迭代数组
	Index int     //当前索引
}

func (u *UserIterator) HasNext() bool {
	if u.Index >= len(u.Users) {
		return false
	}
	return true
}

func (u *UserIterator) GetNext() *User {
	if u.Index >= len(u.Users) {
		return nil
	}
	res := u.Users[u.Index]
	u.Index++
	return res
}
