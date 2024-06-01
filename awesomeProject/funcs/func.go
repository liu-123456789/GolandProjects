package main

func main() {
	//str1, i1 := f0("str1", 123)
	//println(str1, i1)
	//j1, c1 := f1("str2", 7788)
	//println(j1, c1)
	_, _ = f3(123, "s3")
	//println()
}

func f0(str1 string, i int) (string, int) {
	return str1, i
}

func f1(string2 string, int2 int) (s1 string, i1 int) {
	return string2, int2
}

func f3(int2 int, string2 string) (string, int) {
	return "s2", 12345
}
