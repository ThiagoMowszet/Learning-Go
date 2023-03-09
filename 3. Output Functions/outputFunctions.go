package main 

import("fmt")

func outputs() {

    // Print vs Println vs Printf


    // Print() function prints its arguments with their default format.

    var i,j string = "Hello", "World"

    fmt.Print(i)
    fmt.Print(j)

    // Output ⬆ = HelloWorld

    // If we want to print the arguments in new lines, we need to use \n

    fmt.Print(i, "\n")
    fmt.Print(j, "\n")

    // Output ⬆ = Hello
    // Output ⬆ = World

    // Or we can only use one Print for printing multiples lines

    fmt.Print(i, "\n", j)

    // Output ⬆ = Hello
    // Output ⬆ = World

    // If we want to add space between arguments, we need to use " "

    fmt.Print(i, " ", j) // Hello World


    // But if neither of the arguments are string, print() inserts a space between the arguments

    var a,b = 10, 20
    fmt.Print(a,b)

    // Output ⬆ = 10 20



    // Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:

    var c,d = 5, 15
    fmt.Println(c,d) 

    // Output ⬆ = 5
    // Output ⬆ = 15



    // Printf() function first formats its argument based on the given formatting verb and then prints them.

    // %v = to print the value of the arguments
    // %T = to print the type of the arguements

    var hi string = "Hello"
    var age int = 15

    fmt.Printf("hi has value: %v and type: %T\n", hi, hi)  
    fmt.Printf("age has value: %v and type: %T", age, age)

}
