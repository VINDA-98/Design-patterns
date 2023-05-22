package main

// @Title  _2_visitor
// @Description  MyGO
// @Author  WeiDa  2023/5/22 11:50
// @Update  WeiDa  2023/5/22 11:50

func main() {
	dog := &Dog{}
	vs := &ColorVs{}
	dog.Accept(vs)
}
