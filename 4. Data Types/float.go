package main

import("fmt")

func float() {
    
/*

The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.
The float data type has two keywords:

float32 = 32 bits	| Range: -3.4e+38 to 3.4e+38.

float64 = 64 bits | Range: -1.7e+308 to +1.7e+308.

*/


// float32

  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)


// float64

  var i float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", i, i)

}
