package main

func Switch(status int) string {
	switch status {
	case 0:

		return "初始化"

	case 1:
		return "执行中"
	case 2:
		return "重试"
	case 3:
		return "执行失败"
	default:
		return "异常"

	}
}
