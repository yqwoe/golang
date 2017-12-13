package main

func main() {
	a := 10
	if a > 1 {
		println("a > 1")
	} else if a < 10 {
		println(" a < 10 ")
	}

	b := "abc"

	switch b {
	case "a":
		println("a")
	case "b", "abc":
		println("b")
	case "c":
		println("c")
	default:
		println("abc")
	}

	var x interface{}

	println(x == nil) // true

	switch i := x.(type) {
	case nil:
		println(i)
	case int:
		println(i)
	case float64:
		println(i)
	case func(int) float64:
		println(i)
	case bool, string:
		println(i)
	default:
		println(i)
	}


	for a := 0; a < 10; a++{
		println(a)
	}

	var c int = 15
	var d int

	for d < c {
		d++
		println(d)
	}
	//numbers := []int{1,2,3,4,6,5,56,7}
	for i,x := range []int{1,2,3,4,6,5,56,7}{
		println(i," => ",x)
	}


	var i,j int

	 for i = 2; i < 100; i++{
		for j = 2; j<= (i/j);j++{
			if i%j==0 {
				break;
			}
		}
		if j > (i / j){
			println(i,"是素数")
		}
	}
}
