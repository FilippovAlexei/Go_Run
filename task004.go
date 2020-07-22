/*
Вводится количество вводимых чисел.
Вводятся сами числа.
Требуется найти из них максимальное
*/
package main

import "fmt"

func main() {
	var number, num int
	fmt.Println("Введите количество водим чисел")
	fmt.Scan(&number)
	fmt.Println("Сами числа")
	fmt.Scan(&num)

	if number > num {

		fmt.Print("Ввел больше чисел")

	} else {
		fmt.Println("sfd")
	}
	fmt.Println("sd")
}
