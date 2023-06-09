package main

// @Title  __factory
// @Description  MyGO
// @Author  WeiDa  2023/5/4 10:42
// @Update  WeiDa  2023/5/4 10:42

type ICloth interface {
	SetColor(string)
	SetSize(string)
	GetColor() string
	GetSize() string
	GetBrand() string
}

type Cloth struct {
	Brand string
	Color string
	Size  string
}

func NewCloth(color string, size string, brand string) *Cloth {
	return &Cloth{Color: color, Size: size, Brand: brand}
}

func (c *Cloth) SetColor(col string) {
	c.Color = col
}
func (c *Cloth) SetSize(size string) {
	c.Size = size
}
func (c *Cloth) GetColor() string {
	return c.Color
}
func (c *Cloth) GetSize() string {
	return c.Size
}
func (c *Cloth) GetBrand() string {
	return c.Brand
}
