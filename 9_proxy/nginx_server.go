package main

import (
	"errors"
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/10 15:59
// @Update  WeiDa  2023/5/10 15:59

type Nginx struct {
	AppLi       *App
	MaxAllowReq int            //最大允许请求次数
	RateLimit   map[string]int //限流请求
}

func NewNginx() *Nginx {
	return &Nginx{
		AppLi:       &App{},
		MaxAllowReq: 1,
		RateLimit:   map[string]int{},
	}
}

// checkParams 检查参数
func (n *Nginx) checkParams(url string) error {
	if n.RateLimit[url] == 0 {
		n.RateLimit[url] = 1
	} else {
		n.RateLimit[url]++
	}
	//超过最大请求次数，抛出错误
	if n.RateLimit[url] > n.MaxAllowReq {
		log.Println("Beyond the limit number ... ")
		return errors.New("Beyond the limit number ... ")
	}

	return nil
}

// HandlerReq 实现Nginx的请求头处理方法
func (n *Nginx) HandlerReq(url, method string) (int, string) {
	if err := n.checkParams(url); err != nil {
		return 500, err.Error()
	}

	return n.AppLi.HandlerReq(url, method)
}
