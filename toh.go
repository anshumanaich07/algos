package main

import "fmt"

func TOH(disks int, A, B, C string) {
  if disks > 0 {
    TOH(disks-1, A, C, B)
    fmt.Printf("Moved %v --> %v \n", A, C)
    TOH(disks-1, B, A, C)
  }
}

func main() {
  disks := 5
  TOH(disks, "A", "B", "C")
}
