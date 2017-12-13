package main


/*
+ 加
- 减
* 乘
/ 除
% 取余
++ 自增
-- 自减

== 检查值是否相等
!= 检查值不相等
> 检查值是否大于
< 检查值是否小于
>= 检查值是否大于等于
<= 检查值是否小于等于


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

	println("=======================================================")

	println("a == b",a == b)
	println("a != b",a != b)
	println("a > b",a > b)
	println("a < b", a < b)
	println("a >= b",a >= b)
	println("a <= b", a <= b)

	if a > 10 || b < 10{
		println("a > 10 || b < 10")
	}else{
		println(a,b)
	}

	// 0011 1100 => 60

	println( 0 * 2 + 0) // 从第一位开始
	println( 0 * 2 + 0)
	println( 0 * 2 + 1)
	println( 1 * 2 + 1)
	println( 3 * 2 + 1)
	println( 7 * 2 + 1)
	println( 15 * 2 + 0)
	println( 30 * 2 + 0) // 0011 1100 => 60

	println( 60 << 2)

	println(0 * 2 + 1)
	println(1 * 2 + 1)
	println(3 * 2 + 1)
	println(7 * 2 + 1)
	println(15 * 2 + 0)
	println(30 * 2 + 0)
	println(60 * 2 + 0)
	println(120 * 2 + 0) // 60 = 0011 1100  << 2 => 1111 0000 => 240

	// 0011 1100 & 0000 1101 => 0000 1100
	// 0011 1100 & 1111 0000 => 0011 0000

	println(0 * 2 + 0)
	println(0 * 2 + 0)
	println(0 * 2 + 1)
	println(1 * 2 + 1)
	println(3 * 2 + 0)
	println(6 * 2 + 0)
	println(12 * 2 + 0)
	println(24 * 2 + 0) //=> 176

	println(60 & 240)

	// 1000001001

	println(0 * 2 + 1)
	println(1 * 2 + 0)
	println(2 * 2 + 0)
	println(4 * 2 + 0)
	println(8 * 2 + 0)
	println(16 * 2 + 0)
	println(32 * 2 + 1)
	println(65 * 2 + 0)
	println(130 * 2 + 0)
	println(260 * 2 + 1)

	// 521 => 1000001001

	println(521 / 2) // 1
	println(260 / 2) // 0
	println(130 / 2) // 0
	println(65 / 2) // 1
	println(32 / 2) // 0
	println(16 / 2) // 0
	println(8 / 2) // 0
	println(4 / 2) // 0
	println(2 / 2) // 0
	println(1 / 2) // 1


	// 0011 1100 | 0000 1101 => 0011 1101 => 61

	// 0011 1100 ^ 0000 1101 => 0011 0001  => 按位异或运算,加入对数相异 结果为1,否则为0

	// 0011 1100 << 2 => 1111 0000 右移两位

	println(60 << 3) // 111100000

	println(0 * 2 + 1)
	println(1 * 2 + 1)
	println(3 * 2 + 1)
	println(7 * 2 + 1)
	println(15 * 2 + 0)
	println(30 * 2 + 0)
	println(60 * 2 + 0)
	println(120 * 2 + 0)
	println(240 * 2 + 0)

	println(60 >> 5) // 00000001

	println(0 * 2 + 0)
	println(0 * 2 + 0)
	println(0 * 2 + 0)
	println(0 * 2 + 0)
	println(0 * 2 + 0)
	println(0 * 2 + 0)
	println(0 * 2 + 0)
	println(0 * 2 + 1)


	println(-3 << 2) // 100000011

	var ptr *int

	ptr = &a
	println(*ptr)

}