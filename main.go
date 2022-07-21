package main

import "fmt"

func main() {

	e1 := elements{1, 2, 3, 4, 5, 6, 7, 8, 9}
	e2 := elements{2, 4, 6, 8, 10}
	e3 := elements{1, 3, 5, 7, 9}

	sum1 := sumElements(e1)
	sum2 := sumElements(e2)
	sum3 := sumElements(e3)

	e1.print()
	fmt.Println("sum1:", sum1)

	e2.print()
	fmt.Println("sum2:", sum2)

	e3.print()
	fmt.Println("sum3:", sum3)
}
