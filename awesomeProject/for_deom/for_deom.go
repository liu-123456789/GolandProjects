package main

func main() {
	//forman()
	//forArray()
	//forSlice()
	//loopbug()
	var str string
	str = Switch(10)
	println(str)
}

func forArray() {
	println("遍历数组")
	arr := [3]string{"a", "b", "c"}
	for idx, val := range arr {
		println(idx, &val)
	}
}
func forSlice() {
	println("遍历切片")
	arr := []string{"a1", "b1", "c1"}
	for idx, _ := range arr {
		println(idx, arr[idx])
	}
}

func forman() {
	println("遍历map")
	m := map[string]string{
		"1": "A",
		"2": "B",
		"3": "c",
	}

	for key, value := range m {
		println(key, value)
	}
}

func loopbug() {
	users := []User{
		{
			name: "tom",
		},
		{
			name: "jeey",
		},
	}
	m := make(map[string]*User)
	for _, u := range users {
		m[u.name] = &u
	}
	for name, u := range m {
		println(name, u.name)
	}

}

type User struct {
	name string
}
