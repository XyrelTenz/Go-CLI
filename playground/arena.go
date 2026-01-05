// Package playground
package playground

import "fmt"

type Player struct {
	Name    string
	HP      int
	MaxHP   int
	Potions int
}

//TODO: TakeDamage
//TODO: DrinkPotion

func TakeDamage(p *Player, amount int) {

	p.HP -= amount

	if p.HP <= 30 {

		fmt.Println("Please Drink Potion to Heal")

	}
}

func DrinkPotion(p *Player, increaseHP int) {

	if p.Potions > 0 {

		p.Potions--

		p.HP += increaseHP

		if p.HP > p.MaxHP {
			p.HP = p.MaxHP
		}
	} else {
		fmt.Println("No Remaining Potion")
	}
}

func Hero() {
	hero := &Player{
		Name:    "Xyrel",
		HP:      75,
		MaxHP:   100,
		Potions: 3,
	}

	fmt.Printf("Hero Name: %v Hero HP: %v \n", hero.Name, hero.HP)
	TakeDamage(hero, 45)
	fmt.Println("HP Remainaing", hero.HP)
	DrinkPotion(hero, 5)
	DrinkPotion(hero, 5)
	DrinkPotion(hero, 5)
	DrinkPotion(hero, 5)
	fmt.Println("HP Remaining", hero.HP)

}
