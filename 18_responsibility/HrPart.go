package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/19 11:21
// @Update  WeiDa  2023/5/19 11:21

import "fmt"

type HrPart struct {
	next IDepartment
}

func (f *HrPart) Execute(e *Employee) {
	if e.HrDone {
		fmt.Errorf("HrPart has serverd ... ")
	}
	fmt.Println("start HrPart ... ")
	e.HrDone = true
	f.next.Execute(e)
}
func (f *HrPart) SetNext(d IDepartment) {
	f.next = d
}
