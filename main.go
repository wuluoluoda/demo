package main

import "fmt"

type student struct{
	ID int
	Name string
}
type class struct{
	Title string
	Students []student

}
func Newname(id int,name string)student{
	return student{
		ID: id,
		Name:name,

	}
}
//创建一个班级变量c1
func main(){
	//创建一个班级变量c1
	c1:=class{
		Title: "火箭101",
		Students: make([]student,0,20),
	}
	//为它添加内容
	for i:=0;i<10;i++{
		temstdent:=(i,fmt.	Sprintf("format:stus%02ds,i"))
		c1.Students=append(c1.Students,temstdent)

	}
	fmt.Printf(format:"%#v\n",c1)
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