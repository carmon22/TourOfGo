package main

import (
  "fmt"
  "math/cmplx"
  "math"
)

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return ///< this is called a "naked" return. It should only be used in short functions as it can harm readability in longer functions
}

/* var statement declares a list of variables; as in function argument lists, the type is last. It can be at package or function level. */
var c, python, java bool

/* Declares variables i and j of type 'int' and initializes them to 1 and 2. If initializer present, the type can be omitted as the variable will take the type of the initializer */
var i, j int = 1, 2

var (
  ToBe bool       = false
  MaxInt uint64   = 1<<64 - 1
  z     complex128 = cmplx.Sqrt(-5 + 12i)
)

/* Constants cannot be declared using the ":=" syntax */
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
  fmt.Println(split(17))

  var i_local int ///< function level var -- default value to 0

  // short assignment statement can be used in place of a var declaration with implicit type
  // outside a function, every statement beings with a keyword (var, func, etc) so the ":=" construct is not available
  k := 3

  fmt.Println(i, j, k, c, python, java, i_local)

  /* Variable declarations may be "factored" into blocks, as with import statements */
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)

  var i int
  var f float64
  var b bool
  var s string
  /* Default values wihtout explicit initial value are given their "zero" value.
    0 for numeric types,
    false for the boolean type,
    "" (the empty string) for strings.
  */
  fmt.Printf("%v %v %v %q\n", i, f, b, s)

  var x, y int = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f2)
	fmt.Println(x, y, z)

  v := 42 // change me!
  fmt.Printf("v is of type %T\n", v)

  const World = "世界"
  fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

  fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
