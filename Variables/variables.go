/* 

Data Types

int, float32, string, bool

declare a varible :  var variablename type = value

*/

package main

import ("fmt"
"reflect" )

func main() {

    var my_first_variable string = "This is my first variable"
    var my_first_int_variable int = 5
    fmt.Println(my_first_variable)
    fmt.Println(my_first_int_variable)


    // also we can declare a varible like this : variablename := value \ This is inferred value, means that the compiler decides the type of variable, based on the value

    status := true
    fmt.Println(status)
    fmt.Println("Type of variable: ", reflect.TypeOf(status)) // bool

    // The diference between var and :=, it with := we can only use inside a function.
    // and with var, the varialbe declaration and value assignment cannot be done separetly (must be done in the same line)


    // Other data types like float

    var pi float32 = 3.15
    fmt.Println(pi)


    // We can declare a variable and later his value

    var a int
    var b string
    var c float32

    a = 15
    b = "different types to declare variables"
    c = 1.1


    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

}
