package main


func main(){
	var a,b = 1,2

	ret := max(a,b)
	println(ret)
	c,d := swap("1","2")
	println(c,d)
}


func max(num1,num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
多个返回值
*/
func swap(x,y string) (string,string){
	return x,y
}