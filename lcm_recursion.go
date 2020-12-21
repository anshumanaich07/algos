package main 

import (
  "fmt"
)


var (
  temp = 1
  count = 0
)


func disp(count int) {
  defer fmt.Printf("LCM called %v times\n", count)
}

func LCM(num1, num2 int) int {
  count++
  disp(count)

  if temp % num1 == 0 && temp % num2 == 0 {
    return temp
  }
  temp++
  LCM(num1, num2)
  return temp
}

func main() {
  fmt.Printf("LCM of 28 and 42 is %v", LCM(28, 42))
}
