package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/19 10:33
// @Update  WeiDa  2023/5/19 10:33

import "fmt"

type Front struct {
	next IDepartment
}

func (f *Front) Execute(e *Employee) {
	if e.FrontDone {
		fmt.Errorf("Front has serverd ... ")
	}
	fmt.Println("start front ... ")
	e.FrontDone = true
	f.next.Execute(e)

}
func (f *Front) SetNext(d IDepartment) {
	f.next = d
}
