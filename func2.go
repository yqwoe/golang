package	main


func main(){
	var a int = 100
	var b int = 200

	println("交换前a的值为:",a)
	println("交换前b的值为:",b)

	//swap(a,b)
	//&变量返回引用地址
	swap(&a,&b)

	println("交换后a的值为:",a)
	println("交换后b的值为:",b)
}



func swap(x *int,y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

//func swap(x,y int ) int{
//	var temp int
//
//	temp = x
//	x = y
//	y = temp
//
//	return temp
//}