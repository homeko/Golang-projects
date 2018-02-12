package main

import "fmt"

func main() {
  fmt.Println("Hello World")

  var age int = 28
  fmt.Println("my age is", age)

  calcBill(50,3);
}

func calcBill(price, no int) int {
  var totalPrice = price * no
  return totalPrice
}
