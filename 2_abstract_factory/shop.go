package main

import "errors"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 11:50
// @Update  WeiDa  2023/5/4 11:50

type Shop interface {
	MakeCloth() ICloth
	MakePant() IPant
}

// NewShop 写一个工厂，可以生产商品
func NewShop(shopType string) (Shop, error) {
	switch shopType {
	case "HHGZ":
		return &HHGZ{}, nil
	case "HXEK":
		return &HXEK{}, nil
	default:
		return nil, errors.New("没有该品牌的商品")
	}
}
