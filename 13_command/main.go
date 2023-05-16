package main

// @Title  _3_command
// @Description  MyGO
// @Author  WeiDa  2023/5/16 10:40
// @Update  WeiDa  2023/5/16 10:40

func main() {
	//实现思路：
	//	提供按钮的命令接口（interface）
	//	设置命令驱动的设备接口，包含设备可以支持的行为信息，例如开机、关机、播放和重启等。
	//	实现命令接口的结构体，结构体要包含对应的驱动设备信息，并且包括需要被实现的命令方法。
	//	提供实现对象结构体（TV），实现Device接口的方法

	//声明开关按钮
	openB := new(Button)
	closeB := new(Button)

	tv := new(TV)

	//设置按钮驱动的电视对象
	openCmd := &OpenCmd{
		Dev: tv,
	}
	closeCmd := &CloseCmd{
		Dev: tv,
	}

	//设置按钮驱动的命令
	openB.Cmd = openCmd
	closeB.Cmd = closeCmd

	//按下开关按钮
	openB.Press()
	closeB.Press()
}
