package main

// @Title  __factory
// @Description  MyGO
// @Author  WeiDa  2023/5/4 10:42
// @Update  WeiDa  2023/5/4 10:42

type IPant interface {
	SetColor(string)
	SetSize(string)
	GetColor() string
	GetSize() string
	GetBrand() string
}

type Pant struct {
	Color string
	Size  string
	Brand string
}

func NewPant(color string, size string, brand string) *Pant {
	return &Pant{Color: color, Size: size, Brand: brand}
}

func (c *Pant) SetColor(col string) {
	c.Color = col
}
func (c *Pant) SetSize(size string) {
	c.Size = size
}
func (c *Pant) GetColor() string {
	return c.Color
}
func (c *Pant) GetSize() string {
	return c.Size
}
func (c *Pant) GetBrand() string {
	return c.Brand
}
