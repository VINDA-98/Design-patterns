package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 11:40
// @Update  WeiDa  2023/5/15 11:40

type VerifyCode struct {
	Code string `json:"code,omitempty"`
}

func NewVerifyCode(code string) *VerifyCode {
	return &VerifyCode{Code: code}
}

func (v *VerifyCode) CheckCode(code string) bool {
	return v.Code == code
}
