package main

// @Title  main
// @Description  OwnPart.go
// @Author  WeiDa  2023/5/19 10:26
// @Update  WeiDa  2023/5/19 10:26
import "fmt"

type OwnPart struct {
	next IDepartment
}

func (f *OwnPart) Execute(e *Employee) {
	if e.EnterSpecPartDone {
		fmt.Errorf("OwnPart has serverd ... ")
	}
	fmt.Println("Start OwnPart ... ")
	e.EnterSpecPartDone = true
}

func (f *OwnPart) SetNext(d IDepartment) {
	fmt.Printf("employee has been entered department ...")
}
