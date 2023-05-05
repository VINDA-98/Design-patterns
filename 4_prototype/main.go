package main

// @Title  __prototype
// @Description  MyGO
// @Author  WeiDa  2023/5/4 16:57
// @Update  WeiDa  2023/5/4 16:57

func main() {
	f1 := &pFile{"文件1"}
	f2 := &pFile{"文件2"}

	dir1 := &pDir{Name: "目录1", Child: []INode{f1}}
	dir2 := &pDir{Name: "目录2", Child: []INode{f2, dir1, f1}}

	dir1.Show("  ")
	dir2.Show("  ")

	dir1.Clone()
	dir2.Clone()
}
