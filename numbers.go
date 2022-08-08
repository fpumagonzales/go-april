package main

import "fmt"

type elements []int

func sumElements(e elements) int {
	sum := 0
	for _, item := range e {
		sum = sum + item
	}

	return sum
}

func (e elements) print() {
	for _, item := range e {
		fmt.Println(item)
	}

}

func avgElements(e elements) int {
	sum := sumElements(e)
	return sum / len(e)
}
