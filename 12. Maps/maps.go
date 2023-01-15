package main 

import("fmt")



/*

Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.

*/



func maps_go(){

    // Syntax:
    // var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
    // b := map[KeyType]ValueType{key1:value1, key2:value2,...}


    var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
    b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

    fmt.Printf("a\t%v\n", a)
    fmt.Printf("b\t%v\n", b)



    // Creating an empty map

    var c = make(map[string]string)
    var d map[string]string

    fmt.Println(c == nil) // True
    fmt.Println(d == nil) // False



    /* 

    Allowed Key Types
    The map key can be of any data type for which the equality operator (==) is defined. These include:

    Booleans
    Numbers
    Strings
    Arrays
    Pointers
    Structs
    Interfaces (as long as the dynamic type supports equality)
    Invalid key types are:

    Slices
    Maps
    Functions
    These types are invalid because the equality operator (==) is not defined for them.


    */



    // Accessing Map Elements
    // Syntax :  value = map_name[key]


    var e = make(map[string]string)
    e["brand"] = "Ford"
    e["model"] = "Mustang"
    e["year"] = "1964"

    fmt.Printf(e["brand"])




    // Updating and Adding Map Elements

    var f = make(map[string]string)
    f["brand"] = "Ford"
    f["model"] = "Mustang"
    f["year"] = "1964"

    fmt.Println(f)

    f["year"] = "1970" // Updating an element
    f["color"] = "red" // Adding an element

    fmt.Println(f)





    // Remove Element from Map
    // Syntax: delete(map_name, key)


    var g = make(map[string]string)
    g["brand"] = "Ford"
    g["model"] = "Mustang"
    g["year"] = "1964"

    fmt.Println(g)

    delete(g,"year")

    fmt.Println(g)




    // Check For Specific Elements in a Map
    // Syntax val, ok :=map_name[key]


    var h = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

    val1, ok1 := h["brand"] // Checking for existing key and its value
    val2, ok2 := h["color"] // Checking for non-existing key and its value
    val3, ok3 := h["day"]   // Checking for existing key and its value
    _, ok4 := h["model"]    // Only checking for existing key and not its value

    fmt.Println(val1, ok1)
    fmt.Println(val2, ok2)
    fmt.Println(val3, ok3)
    fmt.Println(ok4)


    /*

    Maps Are References
    Maps are references to hash tables.

    If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

    */

    var i = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
    j := i

    fmt.Println(i)
    fmt.Println(j)

    j["year"] = "1970"
    fmt.Println("After change to b:")

    fmt.Println(i)
    fmt.Println(j)



}



func iterating_over_Maps(){

    a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

    for k, v := range a {
        fmt.Printf("%v : %v, ", k, v)
    }

}



func iterate_over_maps_specific_order(){
    a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

    var b = []string             // defining the order
    b = append(b, "one", "two", "three", "four")

    for k, v := range a {        // loop with no order
        fmt.Printf("%v : %v, ", k, v)
    }

    fmt.Println()

    for _, element := range b {  // loop with the defined order
        fmt.Printf("%v : %v, ", element, a[element])
    }
}




func main() {

    maps_go()
}
