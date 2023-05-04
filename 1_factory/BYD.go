package main

type BYD struct {
	Car
}

func (B *BYD) SetName(s string) {
	B.Name = s
}

func (B *BYD) GetName() string {
	return B.Name
}

func (B *BYD) SetPrice(i int64) {
	B.Price = i
}

func (B *BYD) GetPrice() int64 {
	return B.Price
}

func newBYD() ICar {
	return &BYD{
		Car: Car{
			Name:  "比亚迪-汉",
			Price: 250000,
		},
	}
}
