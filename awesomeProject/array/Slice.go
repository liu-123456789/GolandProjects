package main

import "fmt"

func shareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1:%v,len:%d,cap%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap%d\n", s2, len(s2), cap(s2))
	s2[0] = 99
	fmt.Printf("s1:%v,len:%d,cap%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap%d\n", s2, len(s2), cap(s2))
	s2 = append(s2, 199)
	fmt.Printf("s1:%v,len:%d,cap%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap%d\n", s2, len(s2), cap(s2))
	s1[0] = 199
	fmt.Printf("s1:%v,len:%d,cap%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap%d\n", s2, len(s2), cap(s2))
}
