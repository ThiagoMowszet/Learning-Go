package main

import "fmt"


func main() {

    // in GO there are 2 ways to declare an array

    // With the VAR keyword : var array_name = [length]datatype{values}
    // Or : var array_name = [...]datatype{values} 


    // And the 2nd option is with the := sign
    // array_name := [length]datatype{values}
    // array_name := [...]datatype{values}


    var arr1 = [5]int{1,2,3,4,5}
    fmt.Println(arr1)

    arr2:=[3]string{"Ferrari", "Mercedes", "Red Bull"}
    fmt.Println(arr2)



    var arr3 = [...]int{1,2,3}
    arr4 := [...]int{4,5,6,7,8}

    fmt.Println(arr3)
    fmt.Println(arr4)


    // Access Element of an array


    fmt.Println(arr2[0]) // Ferrari
    fmt.Println(arr4[4]) // 8


    // Change Element of an array

    arr2[1] = "McLaren"
    fmt.Println(arr2)


    // Array Initialization

    // If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
    //Tip: The default value for int is 0, and the default value for string is "".


    arr5 := [5]int{} // not initialized
    arr6 := [3]int{1,2} // partially initialized
    arr7 := [5]int{1,2,3,4,5} // fully initialized


    fmt.Println(arr5)
    fmt.Println(arr6)
    fmt.Println(arr7)



    // initialize Only Specific Elements

    arr8 := [5]int{1:10, 2:20}
    fmt.Println(arr8) // [0 10 20 0 0]



    
    // Find the Length of an Array - The len() function is used to find the length of an array

    fmt.Println(len(arr1)) //5
    length_arr1 := len(arr1)
    fmt.Println("La longitud del array es ", length_arr1)


}
