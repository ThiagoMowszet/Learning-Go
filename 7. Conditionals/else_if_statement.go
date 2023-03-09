package main

import("fmt")

func else_if_statements(){

    a := 5 
    b := 15

    if a > b {
        fmt.Printf("%v is greather than %v \n", a, b)
    } else if b > a {
        fmt.Printf("%v is greather than %v \n", b, a)
    } else {
        fmt.Println("there equals")
    }

}


func main() {
    else_if_statements()
}
