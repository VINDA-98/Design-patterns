package main

// @Title  main
// @Description  MyGO
// @Author  WeiDa  2023/5/22 10:33
// @Update  WeiDa  2023/5/22 10:33

type HosService struct {
	IsFree bool
	Queue  []IPeople
}

func (h *HosService) NoticePeople() {
	if !h.IsFree {
		h.IsFree = true //允许访问
		// 队列有人等待，取出第一个用户
		if len(h.Queue) > 0 {
			q := h.Queue[0]
			h.Queue = h.Queue[1:] //退出
			q.PermitEnter()       //许可进入
		}

	}
}

func (h *HosService) CanEnter(people IPeople) bool {
	if h.IsFree {
		h.IsFree = false
		return true
	}
	//为true 添加新增等待病人
	h.Queue = append(h.Queue, people)
	return false
}
