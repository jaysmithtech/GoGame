package main

import (
	"fmt"
	"os"

	"github.com/jaysmithtech/learngo/character"
	//	"github.com/jaysmithtech/learngo/character"
)

func fight(char1, char2 *character.Character) {
	fmt.Printf("%s and %s are fighting!\n", char1.Name, char2.Name)

	for char1.IsAlive() && char2.IsAlive() {
		char1.Attack(char2)
		if char2.IsAlive() {
			char2.Attack(char1)
		}
	}

	if char1.IsAlive() {
		fmt.Printf("%s wins the battle!\n", char1.Name)
	} else {
		fmt.Printf("%s wins the battle!\n", char2.Name)
	}
}

func main() {
	// Create characters with Health and AttackPower
	aragorn := character.NewCharacter("Aragorn", "Ranger", 10, 100, 15)
	sprint := character.NewCharacter("Sprint", "Warrior", 12, 120, 12)
	lest := character.NewCharacter("Lest", "Mage", 8, 80, 20)

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Display Aragorn")
		fmt.Println("2. Display Sprint")
		fmt.Println("3. Display Lest")
		fmt.Println("4. Aragorn vs Sprint")
		fmt.Println("5. Aragorn vs Lest")
		fmt.Println("6. Sprint vs Lest")
		fmt.Println("7. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Character Details for Aragorn:")
			fmt.Println(aragorn.Display())
		case 2:
			fmt.Println("Character Details for Sprint:")
			fmt.Println(sprint.Display())
		case 3:
			fmt.Println("Character Details for Lest:")
			fmt.Println(lest.Display())
		case 4:
			fmt.Println("Aragorn vs Sprint!")
			fight(&aragorn, &sprint)
		case 5:
			fmt.Println("Aragorn vs Lest!")
			fight(&aragorn, &lest)
		case 6:
			fmt.Println("Sprint vs Lest!")
			fight(&sprint, &lest)
		case 7:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		// Reset health after each battle (if you want to reset characters after every fight)
		aragorn.Health = 100
		sprint.Health = 120
		lest.Health = 80
	}
}
