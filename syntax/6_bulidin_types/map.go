package main

func Map() {
	//初始化map
	m1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}
	println(m1["key1"])
	//make方法，预估容量
	m2 := make(map[string]string, 4)
	m2["key2"] = "value2"

	//读取元素；两个返回值
	value1, ok := m1["key1"]
	//value1 true
	println(value1, ok)

	value2, ok := m1["key2"]
	//" ",false
	println(value2, ok)

	val2 := m2["key2"]
	println(val2)

	val3 := m2["key3"]
	println(val3) //返回空字符串

	//长度len
	println("len用法")
	println(len(m1)) //1

	//遍历，随机的，不固定
	println("第一次遍历")
	for k, v := range m1 {
		println(k, v)
	}
	println("第二次遍历")
	for k := range m1 {
		println(k, m1[k])
	}
	println("第三次遍历")
	for k, v := range m1 {
		println(k, v)
	}
	//delete删除元素
	delete(m1, "key1")
	println("删除key1后")
	for k, v := range m1 {
		println(k, v)
	}
}
