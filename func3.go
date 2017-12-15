package main

import (
	"math"
)

func main(){
	//计算平方根,匿名函数
	getSquareRoot := func(x float64) float64{
		return math.Sqrt(x)
	}

	println(getSquareRoot(10))
}