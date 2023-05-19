package main

import "log"

// @Title  18_responsibility
// @Description  责任链模式
// @Author  WeiDa  2023/5/19 10:23
// @Update  WeiDa  2023/5/19 10:23

func main() {
	e := &Employee{Name: "VinDa"}
	//  前台报道
	//	人事制度宣讲
	//	合同部合同签订
	//	送入具体部门

	o := &OwnPart{}
	c := &Contract{}
	h := &HrPart{}
	f := &Front{}

	f.SetNext(h) //前台设置下一步为HR

	h.SetNext(c) //HR下一步为签订合同

	c.SetNext(o) //合同下一步为送到具体部门

	f.Execute(e) //执行责任链  从前台开始执行

	log.Println("========================================")

	c.Execute(e) //执行责任链  从签合同开始执行

}
