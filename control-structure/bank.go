package main


import (
	"fmt"
	"example.com/bank/fileops"
)

const accountBalanceFile = "./balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 1000.0)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		// panic("Can't continue, sorry")
	}
	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice was:", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)

		case 2:
			fmt.Print("How much do you want to deposit? ")

			var depositValue float64
			fmt.Scan(&depositValue)

			if depositValue <= 0 {
				fmt.Println("Invalid Value. Must be greater than 0")
				continue
			}

			accountBalance += depositValue
			fmt.Println("Value deposited successfully.")
			fmt.Println("Account Balance: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("How much do you want withdraw from your account? ")

			var withdrawValue float64
			fmt.Scan(&withdrawValue)

			if withdrawValue > accountBalance || withdrawValue <= 0 {
				fmt.Println("Invalid Value or you dont have this value to withdraw")
				continue
			}

			accountBalance -= withdrawValue
			fmt.Println("Value withdrawed successfully.")
			fmt.Println("Account Balance: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			return
			// break não funcionará no for por conta do switch case
		}

		fmt.Println()
	}
}
