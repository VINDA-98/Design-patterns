package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 14:46
// @Update  WeiDa  2023/5/4 14:46

type Ice struct {
	Frame string
	Style string
	Door  string
	Bed   string
}

func NewIce() *Ice {
	return &Ice{}
}
func (i *Ice) SetFrame() {
	i.Frame = "Ice"
}
func (i *Ice) SetStyle() {
	i.Style = "white"
}
func (i *Ice) SetDoor() {
	i.Door = "ice door"
}
func (i *Ice) SetBed() {
	i.Bed = "ice bed"
}

func (i *Ice) BuildHouse() *House {
	return &House{
		Frame: i.Frame,
		Style: i.Style,
		Door:  i.Door,
		Bed:   i.Bed,
	}
}
