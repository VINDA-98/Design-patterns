package main

import "log"

type pDir struct {
	Name  string
	Child []INode
}

func (dir *pDir) Show(sep string) {
	log.Println("\n" + sep + dir.Name)
	for _, v := range dir.Child {
		v.Show(sep + sep)
	}
}

func (dir *pDir) Clone() INode {
	cloneChild := make([]INode, len(dir.Child))
	for i, v := range dir.Child {
		cloneChild[i] = v.Clone()
	}
	return &pDir{Name: dir.Name + "_clone", Child: cloneChild}
}
