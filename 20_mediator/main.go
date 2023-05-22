package main

// @Title  _0_mediator
// @Description  MyGO
// @Author  WeiDa  2023/5/22 10:29
// @Update  WeiDa  2023/5/22 10:29

func main() {
	h := &HosService{
		Queue:  make([]IPeople, 0),
		IsFree: true,
	}

	p1 := &Programmer{Name: "VinDa1", Mgr: h}
	p2 := &Programmer{Name: "VinDa2", Mgr: h}

	m1 := &Manager{Name: "VinDa3", Mgr: h}
	m2 := &Manager{Name: "VinDa4", Mgr: h}

	p1.Enter()

	p2.Enter()
	m1.Enter()
	m2.Enter()

	p2.Leave()
	m1.Leave()
}
