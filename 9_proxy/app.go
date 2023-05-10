package main

import "log"

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/10 15:55
// @Update  WeiDa  2023/5/10 15:55

type App struct {
}

func (a *App) HandlerReq(url, method string) (statusCode int, msg string) {
	log.Println("url ==", url, "  method ==", method)
	// 设置允许被请求成功的访问路径
	if url == "/get_info" && method == "GET" {
		return 200, "success"
	}

	if url == "/edit_info" && method == "POST" {
		return 200, "success"
	}

	return 400, "failed  ..."
}
