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



}



func main(){

    loops_go()
}
