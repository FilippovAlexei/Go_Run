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
<<<<<<< Updated upstream
	resulr := (A + B)
	s := resulr * resulr
=======
	resulr := (A * *B)
	//s := resulr * B
>>>>>>> Stashed changes
	fmt.Println(s)
}
