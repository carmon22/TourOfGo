package main

import (
  "fmt"
  "math"
)

func sqrt(x float64) string {
  /* if statements need not be surrounded by parentheses but the braces { } are required */
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v  
  }
  return lim
}

func pow1(x, n, lim float64) float64 {
  /* Variables declared inside an if short statement are also available in any of the else blocks */
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
  /*
    Unlike other languages, there are no parentheses surrounding the three components of the for statement. Braces are always required.
  */
  sum := 0
  for i:= 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)

  /*
    The init and post statements are optional
  */  
  sum2 := 1
  for ; sum2 < 1000; {
    sum2 += sum2
  }
  fmt.Println(sum2)

  /*
    C's while is spelled "for" in Go. Kinda odd but For is Go's "while"
  */
  sum3 := 1
  for sum3 < 1000 {
    sum3 += sum3
  }
  fmt.Println(sum3)

  fmt.Println(sqrt(2), sqrt(-4))

  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
    pow(3, 2, 5),
    )
 
  fmt.Println(
    pow1(3, 2, 10),
    pow1(3, 3, 20),
    )
}
