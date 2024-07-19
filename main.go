package main

import "fmt"

func helloworld() {
	fmt.Println("helloWOrld")
}

func hellodd() {
	fmt.Println("helloWOrld")
}

func main() {
	fmt.Println("demo")
	//类型定义
	type NewInt int

	//类型别名
	type MyInt = int

	func main() {
		var a NewInt
		var b MyInt

		fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
		fmt.Printf("type of b:%T\n", b) //type of b:int
	}
}
