package main

import (
	"fmt"
	"log"
)

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 11:56
// @Update  WeiDa  2023/5/15 11:56

type Bank struct {
	Ac *Account
	Co *VerifyCode
	Re *Recorder
	No *NoticeSys
}

func NewBank(account, code string) *Bank {
	return &Bank{
		Ac: NewAccount(account),
		Co: NewVerifyCode(code),
		Re: NewRecorder(),
		No: NewNotice("[VinDa Bank Business]"),
	}
}

// SaveMoney 存款业务
func (b *Bank) SaveMoney(name, code string, money int) {
	// 检查账户
	if !b.Ac.CheckAccount(name) {
		log.Println("存款:用户信息校验失败")
		return
	}
	// 检查验证码
	if !b.Co.CheckCode(code) {
		log.Println("存款:验证码校验失败", code)
		return
	}
	// 统计入库
	sum := b.Re.GetData() + money
	b.Re.SaveData(sum)
	// 发送通知 其实可以异步 例如转账时间过久，可以提前通知
	b.No.SendNotice(fmt.Sprintf("存款成功,当前余额:[%d]", sum))
}

// WithdrawMoney 取款
func (b *Bank) WithdrawMoney(name, code string, money int) {
	// 检查账户
	if !b.Ac.CheckAccount(name) {
		log.Println("取款:用户信息校验失败")
		return
	}
	// 检查验证码
	if !b.Co.CheckCode(code) {
		log.Println("取款:验证码校验失败:", code)
		return
	}
	// 写入数据
	nowMoney := b.Re.GetData()
	if nowMoney > money {
		nowMoney -= money
	} else {
		log.Println("取款失败,余额不足")
		return
	}

	b.Re.SaveData(nowMoney)
	// 通知客户
	b.No.SendNotice(fmt.Sprintf("取款成功,当前余额:[%d]", nowMoney))
}
