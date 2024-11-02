package main // test the compiler that all the files should be copiled

import "fmt" //importing packages that contains methods and fmt contains methods to manipulate strings

func main(){
	// all functions start with func and main is our entry point and only one in entire dir
	fmt.Println("Benchod")

	//variables 
	var firstName string = "Munakala"
	var middleName = "";
	lastName := "Partha Saradhi"; // we can only use this inside a function

	fmt.Println(firstName,middleName,lastName);

	var n1 int8 = 34;
	var n2 int16 = 3423;
	n3 := 45676
	fmt.Println(n1,n2,n3);

	//how to use fmt library
	fmt.Print("Partha \n")
	fmt.Print("Saradhi \n") //pritn doesn't add new line and use \n to add new line

	name := "Partha Saradhi"
	leetcode := 1799.905453463;
	codeforces := 444;

	fmt.Printf("my name is %v and my leetcode rating is %0.4f and my codeforces rating is %vv\n",name,leetcode,codeforces)

	//sPrintf()
	sprint := fmt.Sprintf("my name is %q",name);
	fmt.Println(sprint)
}