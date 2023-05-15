package main

import "errors"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 14:59
// @Update  WeiDa  2023/5/15 14:59

const (
	APlay = "APlayer"
	BPlay = "BPlayer"
)

type DressFactory struct {
	Dm map[string]Dress
}

var DressInstance = &DressFactory{
	Dm: map[string]Dress{},
}

// GetDressInstance 得到服装实例
func GetDressInstance() *DressFactory {
	return DressInstance
}

func (d *DressFactory) GetDressType(pType string) (Dress, error) {
	if d.Dm[pType] != nil {
		return d.Dm[pType], nil
	}
	// d.Dm[pType] 为 nil 的情况
	switch pType {
	case APlay:
		d.Dm[pType] = NewAPlayer()
		return d.Dm[pType], nil
	case BPlay:
		d.Dm[pType] = NewBPlayer()
		return d.Dm[pType], nil
	default:
		return nil, errors.New("暂无支持的服装类型")
	}

}
