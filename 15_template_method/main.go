package main

// @Title  _5_template_method
// @Description  MyGO
// @Author  WeiDa  2023/5/17 12:09
// @Update  WeiDa  2023/5/17 12:09

func main() {
	m := &Mgr{}

	m.B = &Wechat{
		AutoMsg: "wechat auto msg",
	}

	m.GetAndSendMsg()

	m.B = &Wework{
		AutoMsg: "wework auto msg",
	}
	m.GetAndSendMsg()

	m.B = &Lark{
		AutoMsg: "lark auto msg",
	}
	m.GetAndSendMsg()
}
