package main

import ("fmt")

func consts() {

const SOLARSYSTEM_PLANETS int = 8 

fmt.Println(SOLARSYSTEM_PLANETS) // 8

// There are 2 types of constants: Typed constants and Untyped constants.

    // Typed constants

    const A int = 1
    fmt.Println(A)

    // Untyped constants

    const B = 2 // In this case the type of constant is inferred from the value, means the compiler decides the type of the constant, based on the value



// Multiple Const declaration

    const (
        PI float32 = 3.14
        I = "Hi!"
    )

    fmt.Println(PI)
    fmt.Println(I)
}
