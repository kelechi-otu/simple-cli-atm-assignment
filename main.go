package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pin string = "0000"
var balance float64 = 0

func main() {
	fmt.Println("welcome to kaycees simple atm")

	// create a new scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("insert your pin: ")

	// prompt user for input
	scanner.Scan()

	// save the result from the user
	pin = scanner.Text()
	fmt.Println(pin)

	for {
		fmt.Println("select one option \n1. change pin \n2. check balance \n3. withdraw \n4. deposit \n5. exit")

		scanner.Scan()
		selectedOption := scanner.Text()


		if selectedOption == "1" {
			fmt.Print("insert your new pin: ")
			// prompt user for input
			scanner.Scan()

			// save the result from the user
			pin = scanner.Text()

			fmt.Print("do you want to perform another transaction y or n: ")
			scanner.Scan()

			if scanner.Text() == "y" {
				continue
			}
			break
		}

		if selectedOption == "2" {
			fmt.Println("this is ur balance: ", balance)

			fmt.Print("do you want to perform another transaction y or n: ")
			scanner.Scan()

			if scanner.Text() == "y" {
				continue
			}
			break
		}

		if selectedOption == "3" {
			fmt.Print("How much do you want to withdraw: ")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(scanner.Text(), 32)
			if amount >= balance {
				fmt.Println("Insufficient funds")
			} else {
				balance = balance - amount
			}

			fmt.Print("do you want to perform another transaction y or n: ")
			scanner.Scan()

			if scanner.Text() == "y" {
				continue
			}
			break
		}

		if selectedOption == "4" {
			fmt.Print("How much do you want to deposit: ")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(scanner.Text(), 32)
			balance = balance + amount
			fmt.Println("Transaction successful")

			fmt.Print("do you want to perform another transaction y or n: ")
			scanner.Scan()

			if scanner.Text() == "y" {
				continue
			}
			break
		}

		if selectedOption == "5" {
			break
		}

	}

}
