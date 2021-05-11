package main

import (
	"bufio" // package to take input from the user
	"fmt"
	"os"
	"strconv" // string converter/type casting
	"strings"
	//WHATEVER YOU IMPORT YOU HAVE TO USE IT
)

func main() {
	/*var fullN string
	fmt.Println("Enter your full name")
	fmt.Scanln(&fullN)//it ask input for the value of fullN
	fmt.Println(fullN)*/
	//IP: ashar alam
	//OP: ashar {coz it doesnt take input after a space}
	/*	var myName string = "ashar"
		var a, _ = fmt.Println(myName)
		fmt.Println(a)*/
	//OP: ashar    6
	//TAKING USER INPUT WITH BUFIO...................
	// reader := bufio.NewReader(os.Stdin)// reader reads the values from os.Stdio
	// fmt.Print("Enter your full name")
	// my_name, _ := reader.ReadString('\n')
	// fmt.Println(my_name)
	//OP :md ashar alam {read the whole line}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating")
	rating, _ := reader.ReadString('\n')
	myRating, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64) 
	//TrimSpace is used to remove space,it takes one more parameter(size ex: 64,32)
	// means i want it to read till it hits (\n or enter)
	//NOTE: so whenever taking input you are pressing enter so the value of enter is also taken
	// that is why we use variable, _ := to store the value of inter in _
	fmt.Println(myRating + 2)
}
