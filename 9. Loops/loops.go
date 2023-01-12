package main 

import("fmt")


func loops_go(){

    // The for loop is the only loop available in Go.

    for i := 0; i <= 10; i++ {
        fmt.Printf("el valor de i es: %v \n", i)
    }


    // Continue statement

    for i := 0; i < 51; i+=10{
        if i == 10 {
            continue
        } 
        fmt.Printf("i es igual a : %v\n", i)
    }



    // Break statement

    for i := 0; i <= 5; i++{
        if i == 3{
            break
        }
        fmt.Printf("%v valor es : %v \n", i, i)
    }



    // Nested loops

    adj := [2]string{"big", "tasty"}
    fruits := [3]string{"apple", "orange", "banana"}
    for i:=0; i < len(adj); i++ {
        for j:=0; j < len(fruits); j++ {
            fmt.Println(adj[i],fruits[j])
        }
    }


    // Range
    // is used to more easily iterate over an array, slice or map. It returns both the index and the value.

    fruits2 := [3]string{"apple", "orange", "banana"}
    for idx, val := range fruits2 {
        fmt.Printf("%v\t%v\n", idx, val)
    }




    // To only show the value or the index, you can omit the other output using an underscore (_).


    fruits3 := [3]string{"apple", "orange", "banana"}
    for _, val := range fruits3 {
        fmt.Printf("%v\n", val)
    }


    // we want to omit the values (idx stores the index, val stores the value)

    fruits4 := [3]string{"apple", "orange", "banana"}

    for idx, _ := range fruits4 {
        fmt.Printf("%v\n", idx)
    }

}




func main(){

    loops_go()
}
