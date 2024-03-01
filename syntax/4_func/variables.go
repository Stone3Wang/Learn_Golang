package main

// 不定参数
func YourName(name string, alias ...string) {
	if len(alias) > 0 {
		println(alias[0])
	}
}

func YourNameInvoke() {
	YourName("A")
	YourName("大B", "B爷")
	YourName("JB", "小B", "大A")
}
