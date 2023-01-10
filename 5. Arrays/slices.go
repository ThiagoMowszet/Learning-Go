package main 

import "fmt"

func slices_arrays(){

    /*

    slices are also used to store multiple values of the same type in a single variable.
    However, unlike arrays, the length of a slice can grow and shrink as you see fit

    In Go, there are several ways to create a slice:
    Using the []datatype{values} format Create a slice from an array Using the make() function

    Syntax

    slice_name := []datatype{values}
    common way : myslice := []int{}

    */



    /*

    In Go, there are two functions that can be used to return the length and capacity of a slice:

    len() function - returns the length of the slice (the number of elements in the slice)
    cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

    */


    myslice1 := []int{}
    fmt.Println(len(myslice1)) // 0
    fmt.Println(cap(myslice1)) // 0
    fmt.Println(myslice1) // []

    myslice4 := []string{"Go", "Slices", "Are", "Powerful"}
    fmt.Println(len(myslice4)) // 4
    fmt.Println(cap(myslice4)) // 4
    fmt.Println(myslice4) // [Go Slices Are Powerful]



    // Crete a slice from an array

    // var myarray = [length]datatype{values} // An array
    // myslice := myarray[start:end] // A slice made from the array



    arr1 := [6]int{10, 11, 12, 13, 14,15}
    myslice := arr1[2:4]

    fmt.Printf("myslice = %v\n", myslice) // myslice = [12 13]
    fmt.Printf("length = %d\n", len(myslice)) // lenght = 2
    fmt.Printf("capacity = %d\n", cap(myslice)) // capacity = 4




    // Create a slice with the make() function


    // Syntax: slice_name := make([]type, length, capacity)


    myslice5 := make([]int, 5, 10)
    fmt.Printf("myslice1 = %v\n", myslice5)
    fmt.Printf("length = %d\n", len(myslice5))
    fmt.Printf("capacity = %d\n", cap(myslice5))


    // with omitted capacity
    myslice8 := make([]int, 5)
    fmt.Printf("myslice4 = %v\n", myslice8)
    fmt.Printf("length = %d\n", len(myslice8))
    fmt.Printf("capacity = %d\n", cap(myslice8))

}
