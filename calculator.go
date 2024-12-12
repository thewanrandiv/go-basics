package main

import ("fmt")

func main() {
	var revenue int
	var expenses int
	fmt.Println("Hello, World!")
	fmt.Println("Enter the revenue:")
	fmt.Scanln(&revenue)

	fmt.Println("Enter the expenses:")
	fmt.Scanln(&expenses)
	profit:=revenue-expenses
	fmt.Println("The profit is:",profit)
}
