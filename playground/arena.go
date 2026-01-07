// Package playground
package playground

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//TODO: Add Inventory
//TODO: Add Random Enemies
//TODO: Store
//TODO: Add Items from Inventory

type Player struct {
	Name      string
	HP        int
	MaxHP     int
	Gold      int
	Potions   int
	Inventory []Product
	Damage    int
}

type Enemy struct {
	Name   string
	HP     int
	MaxHP  int
	Damage int
	Exp    int
}

type Product struct {
	Name   string
	Price  int
	Type   string
	Effect int
}

type Stores struct {
	Name     string
	Products []Product
}

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
// TODO: Implement STORE
func LoadingScreen(p *Player, s Stores) {

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

func Inventory(p *Player) {

}

func Store(p *Player, s Stores) {

	for {

		fmt.Printf("Welcome %s", p.Name)
		fmt.Printf("Your Gold is %d", p.Gold)

		for i, item := range s.Products {
			fmt.Printf("%d, %s - %d Gold", i+1, item.Name, item.Price)
		}

		fmt.Println("0. Exit Shop")

		fmt.Print("Buy Item ID: ")
		var choice int

		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Please input Valid Operations")
		}

		if choice == 0 {
			return
		}

		index := choice - 1

		if index >= 0 && index < len(s.Products) {

			targetItem := s.Products[index]

			if p.Gold >= targetItem.Price {
				p.Gold -= targetItem.Price

				p.Inventory = append(p.Inventory, targetItem)

				fmt.Printf("You bought %s \n", targetItem.Name)

			} else {

				fmt.Println("Not enough Gold")
			}
		}

	}

}

// TODO:
func OpenInventory(p *Player) {

	item := p.Inventory

	fmt.Print(item)
}

// TODO: Adventure
func Adventure(p *Player) {

	enemy := []Enemy{
		{Name: "Goblin", HP: 50, MaxHP: 50, Damage: 10},
		{Name: "Orc", HP: 100, MaxHP: 100, Damage: 20},
		{Name: "Dragon", HP: 300, MaxHP: 300, Damage: 50},
	}

	// Randomize Enemy
	RandomEnemy := rand.Intn(len(enemy))

	// Read the Memoery of enemy and get its Value
	enemies := &enemy[RandomEnemy]

	fmt.Printf("You've Encounter An %s", enemies.Name)

	for {

		if p.HP <= 0 {
			fmt.Println(" You're Dead ")
		}

		if enemies.HP <= 0 {
			enemies.HP = 0
			fmt.Printf("You defeated the %s", enemies.Name)
		}

		// Exit Screen if Dead
		if p.HP <= 0 {
			p.HP = 0
			fmt.Println("Return to Spawn Point")
			return
		}

		// Logic for repeating adventure after previous enemy is defated then randomize new enemy encounter
		if enemies.HP <= 0 {

			enemies = &enemy[rand.Intn(len(enemy))]

			fmt.Printf("You've Encounter An %s \n", enemies.Name)

			return
		}

		fmt.Printf("\n[%s HP: %d] vs [%s HP: %d]\n", p.Name, p.HP, enemies.Name, enemies.HP)

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
			fmt.Printf("You hit the %s for %d damage!\n", enemies.Name, p.Damage)
			enemies.HP -= p.Damage
			p.HP -= enemies.Damage
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

	rand.New(rand.NewSource(time.Now().UnixNano()))
	hero := &Player{
		Name:    "Xyrel",
		HP:      75,
		MaxHP:   100,
		Potions: 3,
		Damage:  15,
	}

	generalStore := Stores{
		Name: "Town Market",
		Products: []Product{
			{Name: "Small Potion", Price: 20, Type: "Consumable", Effect: 20},
			{Name: "Mega Potion", Price: 50, Type: "Consumable", Effect: 100},
			{Name: "Iron Dagger", Price: 100, Type: "Weapon", Effect: 5},
			{Name: "Steel Sword", Price: 250, Type: "Weapon", Effect: 15},
		},
	}
	LoadingScreen(hero, generalStore)

}
