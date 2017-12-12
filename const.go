package main

import "unsafe"

/*
const a = 1
可以省略常量类型,编译器自动推断;
常量不可更改,编译器会提示错误;
iota,特殊常量,可以被编译器修改;
iota初始为0,下一个const自动加1

字符串类型在 go 里是个结构
包含指向底层数组的指针和长度
这两部分每部分都是 8 个字节
所以字符串类型大小为 16 个字节
*/

const abc = 2


const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
	d = iota
	e = "abc"
	f = "ff"
	g
	h = "h"
	i
)

func main(){

	const (
		Unkonw = 0
		Female = 1
		Male = 2
	)

	println(abc,Unkonw,Female,Male)
	println(a,b,c,d,e,f,g,h,i)
}