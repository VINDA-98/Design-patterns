package main

import "log"

type Car struct {
	Name  string
	Price int64
}

func NewCar(name string, price int64) *Car {
	return &Car{Name: name, Price: price}
}

// Run 实现方法
func (c *Car) Run() {
	log.Printf("%s is running,price is %d", c.Name, c.Price)
}

//func (c *Car) SetName(name string) {
//	c.Name = name
//}
//func (c *Car) GetName() string {
//	return c.Name
//}
//func (c *Car) SetPrice(price int64) {
//	c.Price = price
//}
//func (c *Car) GetPrice() int64 {
//	return c.Price
//}
