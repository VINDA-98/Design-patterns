package main

import "fmt"

type pFile struct {
	Name string
}

func (f *pFile) Show(sep string) {
	fmt.Println(sep + f.Name)
}

func (f *pFile) Clone() INode {
	return &pFile{Name: f.Name + "_clone"}
}
