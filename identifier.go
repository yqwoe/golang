package main


var a = "菜鸟教程"
var b string ="runoob.com"
var c bool
var ff float32

// 因式分解,用于声明全局变量
var (
	e int
	f string
	g bool
)

var h,i string= "a","b"
var aa,bb int

/*
声明变量,如果不指定类型,默认会匹配类型
如果指定类型,则会静态检查类型是否一致
int 类型默认值 0,bool默认值false,string为""
float 默认为+0.000000e+000

int,float,bool,string 属于值类型
var j = 0
var i = j => 实际上是在内存中将j的值进行了拷贝
省略var声明只能在函数中声明
省略var和类型,会交由编译器自动推断
以下声明形式不能在当前作用域中重复声明
		j,k := 123,"abc"
		j,k := 234,"def" => 编译器提示报错
		&j返回j的内存地址

	age := 19 // => 如果block中声明age并没有使用,编译器也会提示报错
	全局变量是允许声明但不使用
*/

func main(){
	var j = 1
	var i = &j
	j = 13


	j,k := 123,"abc"
	println(a,b,c,ff)
	println(e,f,g)
	println(h,i)
	println(aa,bb)
	println(j,k)
	println(j,i)
}