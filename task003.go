package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		A, B int
	)

	fmt.Fscan(os.Stdin, &A)
	fmt.Fscan(os.Stdin, &B)
	resulr := (A  B)
	s := resulr * B
	fmt.Println(s)
}
