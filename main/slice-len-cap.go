package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	pritSlice(s) //[2,3,5,7,11,13]
	//截取切片使其长度为0
	s = s[:0]
	pritSlice(s) //[]

	//拓展其长度
	s = s[:4]
	pritSlice(s) //2,3,5,7
	//舍弃前两个值
	s = s[2:]
	pritSlice(s) //5 ,7
}

func pritSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
