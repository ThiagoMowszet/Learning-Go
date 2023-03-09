package main 

import("fmt")

func integers() { 

    /*
    The integer data type has two categories:

    Signed integers - can store both positive and negative values
    Unsigned integers - can only store non-negative values
    */


    // Signed integers -  declared with one of the int keywords, can store both positive and negative values

    var x int = 500
    var y int = -4500

    fmt.Printf("Type: %T, value: %v", x, x)
    fmt.Printf("Type: %T, value: %v", y, y)


    /*

    Go has five keywords/types of signed integers:

    int = 32 bits in 32 bit systems and 64 bit in 64 bit systems \ Range : -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

    int 8 = 8 bits/1 byte \ Range: 	-128 to 127

    int16 = 16 bits/2 byte \ Range:	-32768 to 32767

    int32 = 32 bits/4 byte \ Range: -2147483648 to 2147483647

    int64 = 64 bits/8 byte \ Range: -9223372036854775808 to 9223372036854775807

    */



    // Unsigned integers - declared with one of the uint keywords, can only store non-negative values


    var z uint = 500 
    var w uint = 4500

    fmt.Printf("Type: %T, value: %v", z, z)
    fmt.Printf("Type: %T, value: %v", w, w)


    /*

    Go has five keywords/types of unsigned integers: 

    int = 32 bits in 32 bit systems and 64 bit in 64 bit systems \ Range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

    int 8 = 8 bits/1 byte \ Range: 	0 to 255

    int16 = 16 bits/2 byte \ Range:	0 to 65535

    int32 = 32 bits/4 byte \ Range: 0 to 4294967295

    int64 = 64 bits/8 byte \ Range: 0 to 18446744073709551615

    */


}
