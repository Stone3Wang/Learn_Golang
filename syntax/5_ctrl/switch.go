package main

// 不需要break，默认执行完case后，会自动跳出switch语句
func Switch(Status int) string {
	switch Status {
	case 1:
		return "初始化"
	case 2:
		return "运行中"

	case 3:
		return "暂停"
	default:
		return "未知状态"
	}
}

// 要做到case的每个条件是互斥的
func Switch2(age int) string {
	switch {
	case age >= 18:
		return "成年人"
	case age >= 14:
		return "青少年"
	default:
		return "未成年"
	}
}
