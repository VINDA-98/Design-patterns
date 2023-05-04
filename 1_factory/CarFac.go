package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/4 10:51
// @Update  WeiDa  2023/5/4 10:51

import "errors"

// 具体的车厂
func getCar(band string) (ICar, error) {
	if band == "BYD" {
		return newBYD(), nil
	}
	return nil, errors.New("not support car band ... ")
}
