package main

import (
	"fmt"
) 

func main() {
	//VARIABLE TYPE INTEGER...................
	var number uint8 = 50

	fmt.Println("NUmber is : ", number)
	//uint range = 0-255
	var num1 uint16 = 270
	num1 = num1 + 100

	fmt.Println("num afteraddition is : ", num1)
	//uint16 range 0-65535
	//SIGNED INTEGER................
	var negative_num int16=-50
	negative_num= negative_num+10

	fmt.Println("num is : ",negative_num)
	//int16 range -128 to 127
	//int32 range -2147483648 to 2147483648
	//VARIABLE TYPE FLOAT.........................
	var flt_num float32 = 779
	flt_num= flt_num+0.5895
	fmt.Println("float : ", flt_num)
	//VARIABLE STRING.................
	var myName string="alam"

	myName=myName+" well its me"+" hi"
	// in string + is used to concat
	fmt.Println("ashar",myName)
	 
	var num = 2000
	fmt.Printf("the number is : %T",num)
	//OP: int
	// implicit variable coz data type of the variable is not decleared. 
	//%T gives data type
	var num12 = 2000.89
	fmt.Printf("the number is : %T",num12)
	//OP: float64
	//FASTER VARIABLE DECLARATION
	//can not declare empty var like this : cgpa :=
	my_name := "Ashar"
	cgpa := 7.5

	fmt.Println("my name is : ",my_name)
	fmt.Println("my cgpa is :  ",cgpa)
	fmt.Printf("%T ",cgpa)
//DECLAREING THE VARIABLE WITHOUT VALUE TO CHECK FOR DEFAULT VALUE...........
	var age uint64
	var bi bool
	fmt.Println(bi,age)
	//OP: age default value is 0 and bi value is false
//FLOATING POINTS.............................
	fmt.Printf("number is : %f \n" , 2.987654321)
//OP: number is : 2.987654 {upto 6 dicimal places will be printed}
	fmt.Printf("number is : %g ",2.987654321)
	//OP : number is : 2.987654321 {whole number will be printed}
	fmt.Printf("num is : %.2f \n %.3f \n %.f \n",2.987654321, 2.987654321, 2.987654321)
//OP: num is 2.99    2.988   3 {%.2f : upto 2 decimal %.3f : 3decimal and %.f wii round the num}
//SPRINTF...........
	var out string = fmt.Sprintf("num is %05d",45)// var out stores the value of Sprint
	fmt.Println(out)
//OP: num is 00045 {%05d means print 0s upto 4places and print the values at 5th place onwards}
	fmt.Printf("value is : %9d ", 71)
//OP: value is :        71 
//USER INPUT..............................................................
	var fullN string
	fmt.Scanln(&fullN)
	fmt.Println(fullN)
}
