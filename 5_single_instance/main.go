package main

// @Title  Design_patterns
// @Description  MyGO
// @Author  WeiDa  2023/5/5 10:14
// @Update  WeiDa  2023/5/5 10:14

import (
	"fmt"
	"log"
	"sync"
)

var (
	lock sync.Mutex
	once sync.Once
)

// Single 实例类
type Single struct {
}

// SingleInstance 实例对象(此时创建的 SingleInstance 默认为零值，也就是 nil)
var SingleInstance *Single

// getOneInstance 显示使用 锁的方式，保证线程安全
func getOneInstance() *Single {
	if SingleInstance == nil {
		//多个协程操作，加锁
		lock.Lock()
		defer lock.Unlock()
		// 此处再继续 判断 SingleInstance 是否为 nil，是因为可能其他协程已经拿到锁，且初始化好了 SingleInstance 后，当前协程才拿到锁
		if SingleInstance == nil {
			SingleInstance = &Single{}
		} else {
			log.Println("1 getOneInstance: SingleInstance is not nil ...")
		}
	} else {
		log.Println("2 getOneInstance: SingleInstance is not nil ...")
	}
	return SingleInstance
}

// getTwoInstance
func getTwoInstance() *Single {
	// 通过 sync.Once 来实现单例模式，用来保证某种行为只会被执行一次。
	if SingleInstance == nil {
		log.Println("getTwoInstance: in SingleInstance == nil ...")
		once.Do(func() {
			// 因此 one.Do 只会执行一次，因此不会出现  getOneInstance 的情况，因此此处无需判断 SingleInstance 是否为 nil
			fmt.Println("getTwoInstance: Create SingleInstance successfully ...")
			SingleInstance = &Single{}
		})
	} else {
		log.Println("getTwoInstance : SingleInstance is not nil ...")
	}
	return SingleInstance
}

func main() {
	for i := 0; i < 20; i++ {
		getTwoInstance()
	}

	_, err := fmt.Scanln() // 输出回车就会结束
	if err != nil {
		return
	}
}
