package main

import (
	"fmt"
)

/*
  Computers compute the square root of x using a loop.
  z -= (z*z - x) / (2*z)
  Repeating this adjustment makes the guess better nd better until an answer that is as close
  as possible to the actual square root as can be.
*/
func Sqrt(x float64) float64 {
  // Initialize a floating point value
  z := 1.0

  // Create loop
  for i:= 0; i < 10; i++ {
    z -= (z*z -x) / (2*z)
  }

}

func main() {
	fmt.Println(Sqrt(2))
}
