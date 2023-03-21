package main

import "fmt"

func change()(int,int) {
	return 21,45
}
func main() {
	var name string = "zhangsan"
	fmt.Println(name)

	a ,_:= change()
	println(a)
	
}
