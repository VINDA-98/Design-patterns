package main

// @Title  _1_appearance
// @Description  MyGO
// @Author  WeiDa  2023/5/15 11:31
// @Update  WeiDa  2023/5/15 11:31

func main() {

	//新建银行业务
	bank := NewBank("VinDa-98", "9898")

	//存款
	bank.SaveMoney("VinDa-98", "9999", 1000000000)
	bank.SaveMoney("VinDa-98", "9898", 1000000000)

	//取款
	bank.WithdrawMoney("VinDa-98", "9898", 1)

	//取款
	bank.WithdrawMoney("VinDa-98", "9899", 1)

	//存款
	bank.SaveMoney("VinDa-98", "9898", 2)
}
