package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/19 11:16
// @Update  WeiDa  2023/5/19 11:16
import "fmt"

type Contract struct {
	next IDepartment
}

func (f *Contract) Execute(e *Employee) {
	if e.ContractDone {
		fmt.Errorf("Contract has serverd ... ")
	}
	fmt.Println("start Contract ... ")
	e.ContractDone = true
	f.next.Execute(e)
}

func (f *Contract) SetNext(d IDepartment) {
	f.next = d
}
