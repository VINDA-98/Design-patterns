package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/9 18:42
// @Update  WeiDa  2023/5/9 18:42

type File struct {
	Name string
}

func (f *File) Search(content string) {
	log.Printf("search file name :[%s] ,content:[%s]\n", f.Name, content)
}

func (f *File) GetName() string {
	return f.Name
}
