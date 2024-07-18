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
		temstdent:=(i,fmt.	Sprintf("format:stu%02ds,i"))
		c1.Students=append(c1.Students,temstdent)

	}
	fmt.Printf(format:"%#v\n",c1)
}