package main

import ("fmt")

func multiple_variables() {

// in GO it's possible to declare multiple variables in the same line

    var a, b, c, d int = 1, 2, 3, 4

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)

// if the type keyword it's not specified, we can declare different types of variables in the same line

    var e, f = 6, "Hello"
    g, h := 7, "World!"

    fmt.Println(e) // 6
    fmt.Println(f) // Hello
    fmt.Println(g) // 7
    fmt.Println(h) // World!



}
