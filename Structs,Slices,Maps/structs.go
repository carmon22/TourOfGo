package main

import "fmt"

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
  
  v := Vertex{1,2}
  v.X = 4
  fmt.Println(v.X)

  // Struct fields can be accessed through a struct pointer
  // To access the field X of a struct when we have the struct pointer, p we could write (*p).X
  // But that notation is cumbersome so the language let's use write p.X without explicit 
  p := &v
  p.X = 1e9
  fmt.Println(v)

  // Creating this block with brackets, {}, allows for variable shadowing
  // w/out these blocks golang would throw an error that p is being redefined in same block
  {
    var (
      v1 = Vertex{1, 2}   // has type Vertex
      v2 = Vertex{X: 1}   // Y:0 is implicit
      v3 = Vertex{}       // X:0 and Y:0
      p = &Vertex{1, 2}   // has type *Vertex
    )
    fmt.Println(v1, p, v2, v3)
  }

  var a [2]string // Create an array, "a", of type string with 2 "slots" or memory space
  a[0] = "Hello"
  a[1] = "World!"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  // An array's length is part of its type, so arrays cannot be resized
  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)

  // A "Slice" is a dynamically-sized, flexible view into the elements of an array
  // slices are much more common than arrays
  // The type []T is a slice with elements of type T
  // A slice is formed by specifing two indicies, a low and high bound, seperated by a colon
  // a[low : high]
  var s[]int = primes[1:4]
  fmt.Println(s)

  // Slices are like references to arrays, changing the elements of a slice modifies the corresponding elements of its underlying array
  // other slices that share the underlying array will see those changes
  
  // ======Slicing Defaults======
  /* When slicing, you may omit the high and low bounds and use defaults instead
    For the array -- var a [10]int
    The following slices are equivalent: a[0:10], a[:10], a[0:], a[:]
  */
  {
    var s []int
	  fmt.Println(s, len(s), cap(s))
	  if s == nil {
		  fmt.Println("nil!")
	  }
  }

  // ========== Creating a slice with make ==========
  /* Slices can be created with the built-in "make" function; this is how to create dynamically-sized arrays
    The make function allocates a zeroed array and returns a slice that refers to that array:
  */
  {
    a := make([]int, 5)
	  printSlice("a", a)

	  b := make([]int, 0, 5)
	  printSlice("b", b)

	  c := b[:2]
	  printSlice("c", c)

	  d := c[2:5]
	  printSlice("d", d)
  }
  // To append to a slice use the built-in package "append"
  {
    var s []int
    printSlice1(s)

    // append works on nil slices.
    s = append(s, 0)
    printSlice1(s)

    // The slice can add more than one element at a time.
    s = append(s, 2, 3, 4)
    printSlice1(s)
  }
  // Range
  {
    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

    func power(pow []int){
      for i, v := range pow {
        fmt.Printf("2**%d\n", i, v)
      }
    }
    power(pow)
  }
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice1(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

