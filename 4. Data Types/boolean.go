package main

import ("fmt")

func booleans() {

    var b1 bool = true // typed declaration with initial value
    var b2 = true // untyped declaration with initial value
    var b3 bool // typed declaration without initial value
    b4 := true // untyped declaration with initial value

    fmt.Println(b1) // true
    fmt.Println(b2) // true
    fmt.Println(b3) // false
    fmt.Println(b4) // true

}
