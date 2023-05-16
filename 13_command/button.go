package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/16 11:27
// @Update  WeiDa  2023/5/16 11:27

type Button struct {
	Cmd Command
}

func (b *Button) Press() {
	//Press the button
	b.Cmd.Execute()
}
