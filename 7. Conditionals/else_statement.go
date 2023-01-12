package main

import ("fmt")

func else_statements(){

    a := 5
    b := 15

    if a > b {
        fmt.Printf("%v is greather than %v \n", a, b)
    } else {
        fmt.Printf("%v is greather than %v \n", b, a)
    }
}


func main(){

    else_statements()
}
