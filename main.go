package main

import (
	"fmt"
)

func split(sum int) (x, y int, z float32) {
	x = sum * 4 / 9
	x *= 2
	y = sum - x
	y += 3
	x += 5
	z = 1.

	//dasfsdafasdffsdasdafsdfsdaffsdaasdfafsdsdfa
	//dd

	//asdfsdaf

	return
}

func main() {
	fmt.Println(split(17))
}
