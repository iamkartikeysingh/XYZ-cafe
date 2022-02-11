package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var order string
	var qty int
	var sales []int
	menu := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	fmt.Println("Welcome to XYZ Cafe. Kindly refer from the menu")
	fmt.Println("Menu: ")
	fmt.Println("C: coffee, 40rs\nD: dosa, 80 rs\nT: tomato soup, 20rs\nI : idli , 20rs\nV : vada, 25rs.\nB: bhature &chhole, 30rs\nP: paneer pakoda, 30rs\nM: manchurian, 90rs\nH: hakka noodle, 70rs.\nF: French fries, 30rs\nJ: jalebi, 30rs\nL; lemonade, 15rs\nS: spring roll, 20rs\n")

	for true {
		fmt.Println("Enter your choice and quantity. Type 'END' after entering all choices")
		fmt.Scan(&order)
		order = strings.ToUpper(order)
		if order != "END" {

			fmt.Scan(&qty)
			amount := qty * menu[order]
			fmt.Println("Total cost of this item: ", amount)
			sales = append(sales, amount)
		} else {
			billtotal(sales)
		}
	}
}
func billtotal(sale []int) {
	var sum int = 0
	for i := 0; i < len(sale); i++ {
		sum = sale[i] + sum
	}
	fmt.Println("Total bill is : ", sum)
	fmt.Println("Have a nice day")
	os.Exit(0)
}
