package main

import (
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/9 18:42
// @Update  WeiDa  2023/5/9 18:42

type Folder struct {
	Com  []IComponent
	Name string
}

func (f *Folder) Search(content string) {
	log.Printf("strat search folder , name:[%s] content:[%s] \n", f.Name, content)
	for _, v := range f.Com {
		v.Search(content)
	}
}

func (f *Folder) AddComponent(c IComponent) {
	f.Com = append(f.Com, c)
}
