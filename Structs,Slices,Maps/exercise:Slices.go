package main

import "golang.org/x/tour/pic"

/*
  Implement Pic. It should return a slice of length dy, each element of which is a slice
  of dx 8-bit unsigned integers. When you run the program, it will display your picture, 
  interpreting the integers as grayscale (well, bluescale) values.

  The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
  (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
  (Use uint8(intValue) to convert between types.)
*/

/*
 * Function 
 * @param dx slice
 * @param dy length to slice the slice
 * @return slice of length dy
*/
func Pic(dx, dy int) [][]uint8 {
    // Create slice of [][]uint8 of length dy
    p := make([][]uint8, dy)
    // For each row of p, create a []uint8
    for y := range p {
        p[y] = make([]uint8, dx)
        // For each row in p, insert calculated value to index p[y][x]
        for x := range p[y] {
        	p[y][x] = uint8(x*y)
        }
    }
    // Return the filled in [][]uint8 slice
    return p
}

func main() {
	pic.Show(Pic)
}

