package main
import(
	"fmt"
)
func main(){
	// var num1 int = 9
	// var num2 int =4
	// answer := num1+num2// for addition
	// fmt.Println("sum is", answer)
//WHAT IF VARIABLE ARE OF DIFFERENT DATA TYPE
	//var num1 float32 = 9
	//var num2 int = 4
	//answer:= num1/num2//for division
	//fmt.Println("resultant is ", answer)
//OP: invalid operation: num1 / num2 (mismatched types float32 and int)
	//answer := num1/ float32(num2)// num2 is converted into float
	//fmt.Println(answer)
	//OP: 2.25 {if we convert suppose 6.7 into int it will give 6}
//NOTE : WE FOLLOW BEDMAS 
// B: BRACKET
// E: EXPONENTS
// D: division
// M: MULTIPLICATION
// A: ADDITION
// S: SUBSTRACTION
//PROGRAM TO RETURN THE REMINDER FORM THE USER INPUT
var num1 int
var num2 int
fmt.Println("Enter the numbers")
fmt.Scanln(&num1)// Scanln takes the input and &variableName stores the input into variable
fmt.Scanln(&num2)
result := num1 % num2
fmt.Println("reminder is : ", result)
}