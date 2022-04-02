package main

import(
	"fmt"
	"strings"
)

/*
Test different kinds of deep copy and shallow copy....
case 1:
struct


case 2:
pointer of struct


case 3:
array in struct..

case 4:
map slice in struct...

*/

type Employer struct{
	//var name string	not right
	name string	
}

type Employee struct{
	 no int
	 name string

	 siblings []string
	
	 leaveDuration [3]int
		
	 boss *Employer
}

func (me *Employee) printInfo(){
	fmt.Println("=====Person Info=====")
	fmt.Printf("id %v name %v\n", me.no, me.name)
	fmt.Printf("siblings: %v\n", strings.Join(me.siblings, " "))
	var leaveDays int = 0;
	for _, day := range me.leaveDuration {leaveDays += day}
	fmt.Printf("leave days: %v\n", leaveDays)
	fmt.Printf("boss is: %v\n", me.boss.name)
}

func main(){
	//var boss1 := Employer{"tom"}
	var boss1 = Employer{"tom"}
	var boss2 Employer = Employer{"ben"}
	
	var me = Employee{ 1, "qiang", []string{"ying", "fang"}, [3]int{1, 2, 0}, &boss1}
	
	fmt.Println("first case")
	
	me.printInfo()
	
	fmt.Println("second case")
	
	var me2 = me
	
	me2.no = 2
	
	me2.name = "qiang2"
	
	me2.siblings[1] = "qing"
	
	me2.leaveDuration[2] = 3
	
	me2.boss = &boss2
	
	me2.printInfo()
	
	
	fmt.Println("Third case")
	me.printInfo()
	
	
}