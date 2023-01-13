package main

import "fmt"

func functions_returns(x int, b int) int {

	// If you want the function to return a value,
	// you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function

	// Syntax: Function_name(parameter1 type, parameter2 type) type{ return ouput }

	return x + b

}

// only return 

func named_return_values(a int, b int) (result int) {
    result = b - a
    return
}


func store_value_inVariable(i string, j string)(result string){
    
    result = i + " " + j
    return
}


func multiple_returns(x int, y string) (result int, txt1 string) {

    result = x + x
    txt1 = y + " World!"
    return
}


func omit_results(x int, s string)(result int, txt1 string){

    result = x + x
    txt1 = s + " World!"
    return
}




func main() {

	fmt.Println(functions_returns(1, 2))
    fmt.Println(named_return_values(10, 60))
    full_name := store_value_inVariable("Homero", "Simpson")
    fmt.Println(full_name)
    fmt.Println(multiple_returns(5, "Hello"))
    _, b := omit_results(5, "hello")
    fmt.Println(b)
}
