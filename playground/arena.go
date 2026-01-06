// Package playground
package playground

import (
	"fmt"
	"os"
)

type Player struct {
	Name    string
	HP      int
	MaxHP   int
	Potions int
	Damage  int
}

type Enemy struct {
	Name   string
	HP     int
	MaxHP  int
	Damage int
}

//TODO: Create Monster
//TODO: Loop Choices

func TakeDamage(p *Player, amount int) {

	p.HP -= amount

	if p.HP < 0 {
		p.HP = 0
	}

	fmt.Printf(">> You took %d damage! Current HP: %d/%d \n", amount, p.HP, p.MaxHP)

	if p.HP <= 30 && p.HP > 0 {

		fmt.Println("Please Drink Potion to Heal")

	}
}

func DrinkPotion(p *Player) {

	if p.Potions > 0 {

		healAmount := 20

		p.Potions--

		p.HP += healAmount

		if p.HP > p.MaxHP {
			p.HP = p.MaxHP
		}

		fmt.Printf("You drank a potion and healed %d HP. Current HP: %d/%d \n", healAmount, p.HP, p.MaxHP)
	} else {
		fmt.Println("No Remaining Potion")
	}
}

// LoadingScreen UI
func LoadingScreen(p *Player) {

	for {
		fmt.Println("1. Go Adventure")
		fmt.Println("2. Settings")
		fmt.Println("3. Exit")

		var choice int

		_, err := fmt.Scanln(&choice)

		if err != nil {

			fmt.Println("Please Provide Valid Number")

			var discard string

			fmt.Scanln(&discard)

			continue
		}

		switch choice {
		case 1:
			if p.HP <= 0 {
				fmt.Println("Not Qualified to Adventure")
			} else {
				Adventure(p)
			}
		case 2:
			Settings(p)
		case 3:
			fmt.Println("Exiting The Game")
			os.Exit(0)
		default:
			fmt.Println("Invalid Operations")

		}
	}

}

func Adventure(p *Player) {

	enemy := &Enemy{
		Name:   "Goblin",
		HP:     100,
		MaxHP:  1000,
		Damage: 20,
	}

	fmt.Printf("You've Encounter An %s", enemy.Name)

	for {

		if p.HP <= 0 {
			fmt.Println(" You're Dead ")
		}

		if enemy.HP <= 0 {
			enemy.HP = 0
			fmt.Printf("You defeated the %s", enemy.Name)
		}

		fmt.Printf("\n[%s HP: %d] vs [%s HP: %d]\n", p.Name, p.HP, enemy.Name, enemy.HP)

		fmt.Println("1. Attack")
		fmt.Println("2. Drink Potion")
		fmt.Println("3. Run Away")

		fmt.Print("Action: ")

		var action int

		_, err := fmt.Scanln(&action)

		if err != nil {
			var discard string

			fmt.Scanln(&discard)

			continue
		}

		switch action {
		case 1:
			fmt.Printf("You hit the %s for %d damage!\n", enemy.Name, p.Damage)
			enemy.HP -= p.Damage
			p.HP -= enemy.Damage
		case 2:
			DrinkPotion(p)
		case 3:
			fmt.Println("You ran away safely!")
			return
		default:
			fmt.Println("You hesitated...")
		}

	}
}

func Settings(p *Player) {
	for {
		fmt.Println("\n--- SETTINGS ---")
		fmt.Println("1. Change Name")
		fmt.Println("2. View Full Stats")
		fmt.Println("3. Back to Menu")
		fmt.Print("Select: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter new name: ")
			var newName string
			fmt.Scanln(&newName)
			p.Name = newName
			fmt.Println("Name updated!")
		case 2:
			fmt.Printf("\nName: %s\nHP: %d/%d\nAttack: %d\nPotions: %d\n",
				p.Name, p.HP, p.MaxHP, p.Damage, p.Potions)
		case 3:
			return
		}
	}
}

func Hero() {
	hero := &Player{
		Name:    "Xyrel",
		HP:      75,
		MaxHP:   100,
		Potions: 3,
		Damage:  15,
	}

	LoadingScreen(hero)

}
