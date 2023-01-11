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
    fmt.Printf("myslice5 = %v\n", myslice5) // [0 0 0 0 0]
    fmt.Printf("length = %d\n", len(myslice5))// 5
    fmt.Printf("capacity = %d\n", cap(myslice5)) // 10


    // with omitted capacity
    myslice8 := make([]int, 5)
    fmt.Printf("myslice4 = %v\n", myslice8) // [0 0 0 0 0]
    fmt.Printf("length = %d\n", len(myslice8)) // 5
    fmt.Printf("capacity = %d\n", cap(myslice8)) // 5 - if the capacity parameter is not defined, it will be equal to the length.




// Access elements of a slice


    prices := []int{10, 20, 30} 
    fmt.Println(prices[0]) // 10
    fmt.Println(prices[2]) // 30
    


// Change elements of a slices 

    prices2 := []int{10, 50, 10}
    prices2[2] = 5
    fmt.Println(prices2[2]) // 5



// Append elements to a slice
// Syntax : slice_name = append(slice_name, element1, element2, ...)

    myslice_append := []int{1, 2, 3, 4, 5, 6}
    fmt.Printf("myslice1 = %v\n", myslice_append) // 
    fmt.Printf("length = %d\n", len(myslice_append))
    fmt.Printf("capacity = %d\n", cap(myslice_append))

    myslice_append = append(myslice1, 20, 21)
    fmt.Printf("myslice1 = %v\n", myslice_append)
    fmt.Printf("length = %d\n", len(myslice_append))
    fmt.Printf("capacity = %d\n", cap(myslice_append))

}
