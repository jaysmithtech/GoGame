// character.go
package character

import "fmt"

// Character is an exported struct
type Character struct {
	Name        string
	Class       string
	Level       int
	Health      int
	AttackPower int
}

// NewCharacter is an exported function that creates a new character
func NewCharacter(name, class string, level int, health int, attackpower int) Character {
	return Character{
		Name:        name,
		Class:       class,
		Level:       level,
		Health:      health,
		AttackPower: attackpower,
	}
}

// Display is a method that returns a formatted string describing the character
func (c Character) Display() string {
	return fmt.Sprintf("%s is a level %d %s with %d health and %d attack power.", c.Name, c.Level, c.Class, c.Health, c.AttackPower)
}

// Attack method that decreases the opponent's health
func (c *Character) Attack(opponent *Character) {
	fmt.Printf("%s attacks %s for %d damage!\n", c.Name, opponent.Name, c.AttackPower)
	opponent.Health -= c.AttackPower
	if opponent.Health < 0 {
		opponent.Health = 0
	}
}

func (c Character) IsAlive() bool {
	return c.Health > 0
}
