// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"
func main() {
	// create a local variable "val" and assign it an int value
val:= 40
	// print the value of the local variable "val"
	fmt.Printf("%d=%T",val,val)
	// print the value of the package-level variable "val"
	printGlobal()
	// update the package-level variable "val" and print it again
	updateGlobal()
	// print the pointer address of the local variable "val"
	var a *int
	a = &val
	fmt.Println(a)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
*a = 443;
fmt.Println(val)	
}
func printGlobal(){
	fmt.Printf("\n%s=%T\n",val,val)
}
func updateGlobal(){
	val = "updated global"
}