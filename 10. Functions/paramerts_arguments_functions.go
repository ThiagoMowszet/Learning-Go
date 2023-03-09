package main

import ("fmt")

func examples(name string) {

    // Function with 1 parameter

    fmt.Println("Hello, my name is ", name)

}


func examples_2(name string, age int, coutry string){

    fmt.Println("Hello", age, "years old" , name, "from", coutry)
}

func main() {

    // Syntax: FunctionName(param1 type, param2 type, param3 type) {code to be executed}

    // "Thiago" is our argument: 

    examples("Thiago")
    examples_2("Thiago", 21, "Argentina")
}
