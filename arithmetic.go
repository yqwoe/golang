package main


/*
+ 加
- 减
* 乘
/ 除
% 取余
++ 自增
-- 自减
*/


func main(){
	var a,b = 21,10
	var c int
	c = a + b
	println("a + b = ?",c)
	c = a - b
	println("a - b = ?",c)
	c = a * b
	println("a * b = ?",c)
	c = a / b
	println("a / b = ?",c)
	c = a % b
	println("a % b = ?",c)
	a ++
	println("a ++",a)
	a --
	println("a --",a)
}