package main

import (
  "fmt"
  "strconv"
)


func consecutive(binary[]string, nos int, check string) int {
  // 10110
  count := 0
  for idx, _ := range binary {
    if nos >= len(binary) {
      break
    } else {
        window := binary[idx:nos]
        for _, v := range window {
          if v == check {
            count++
            continue
          } else { break }
        }
      idx++
      nos++
    }
  }
  return count
}

func checkBinary(binary[]string) {
  c1 := consecutive(binary, 3, "0")
  c2 := consecutive(binary, 3, "0")
  c3 := consecutive(binary, 5, "1")
  fmt.Println(binary)

  switch {
    case c1 >= 3 && binary[len(binary)-1] == "1":
      fmt.Println("000-ODD")

    case c2 < 3 && binary[len(binary)-1] == "1":
      fmt.Println(">3 0s-ODD")

    case c3 <= 5 && binary[len(binary)-1] == "0":
      fmt.Println("5 1s - EVEN")

    default:
      fmt.Println("No matches!")
  }
}

func main() {
  var num int64 = 21 
  bin := strconv.FormatInt(num, 2)
  binary := make([]string, len(bin))
  var bit string = ""

  // strconv.FormatInt converts to binary
  for idx, val := range bin {
    bit = string(val)
    binary[idx] = bit
  }
  checkBinary(binary)
}
