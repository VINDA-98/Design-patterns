package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/15 11:40
// @Update  WeiDa  2023/5/15 11:40

type Recorder struct {
	money int
}

func NewRecorder() *Recorder {
	return &Recorder{money: 0}
}

// GetData 用户当前余额
func (r *Recorder) GetData() int {
	return r.money
}

func (r *Recorder) SaveData(money int) {
	r.money = money
}
