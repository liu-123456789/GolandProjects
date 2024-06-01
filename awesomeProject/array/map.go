package main

func mapkv() {
	m1 := map[string]string{
		"key1": "val2",
		"key2": "val2",
		"key3": "val2",
		"key4": "val2",
		"key5": "val2",
		"key6": "val2",
	}
	for k1, v1 := range m1 {
		println(k1, v1)
	}
	delete(m1, "key1")
	println("删除后\n")
	for k, v := range m1 {
		println(k, v)
	}
}
