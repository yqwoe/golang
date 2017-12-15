package main

/*
js: function getSequence(){
		let i = 0
		let func=()=>{
			return i+=1
		}
		return func
	}

	let nextNumber = getSequence()
	nextNumber() => 1
ruby: def get_sequence
		i = 0
		return ->{return i+=1}
	  end
	  next_number = get_sequence
	  next_number.call => 1
*/
func getSequence() func() int {
	i :=0
	return func() int{
		i+=1
		return i
	}
}


func main(){
	nextNumber := getSequence()

	println(nextNumber())
	println(nextNumber())
	println(nextNumber())
	println(nextNumber())
	println(nextNumber())

	nextNumber1 := getSequence()

	println(nextNumber1())
	println(nextNumber1())
	println(nextNumber1())
	println(nextNumber1())
	println(nextNumber1())
	println(nextNumber1())
	println(nextNumber1())

}
