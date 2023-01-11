package main

import ("fmt")


func operathors_go(){

/*

+ Addition
- Substraction
* Multiplication
/ Division
% Modulus
++ Increment
-- Decrement

*/

a := 5
b := 10

var add int = a + b
fmt.Println(add) // 15

var subsc = a - b
fmt.Printf("the type of subsc is: %T\n" , subsc)
fmt.Println(subsc) // -5


var multip = a * b
fmt.Println(multip) // 50


var divis = b / a
fmt.Println(divis) // 2


var mod = b % a
fmt.Println(mod) // 0


}


func main() { 

operathors_go()

}
