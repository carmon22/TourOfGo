package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

// Zero value of a map is nil. A nill map has no keys, nor can keys be added.
// The make function returns a map of the tiven type, initialized and ready for use
var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)

  m["Bell Labs"] = Vertex{ 40.68433, -74.39967, }

  fmt.Println(m["Bell Labs"])
}
