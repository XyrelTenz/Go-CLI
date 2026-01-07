package structs

import (
	"fmt"
	"time"
)

// --- 1. The Base Struct ---
// This contains common fields we want to reuse.
type BaseEntity struct {
	ID        int
	CreatedAt time.Time
}

// A method belonging to BaseEntity
func (b BaseEntity) Describe() {
	fmt.Printf("ID: %d | Created At: %s\n", b.ID, b.CreatedAt.Format(time.RFC822))
}

// --- 2. The Embedding Struct ---
// User embeds BaseEntity. It "inherits" ID and CreatedAt.
type User struct {
	BaseEntity // <--- EMBEDDING HAPPENS HERE (No field name provided)
	Username   string
	Email      string
}

// User can add its own methods
func (u User) Greeting() {
	fmt.Printf("Hello, I am %s!\n", u.Username)
}

// --- 3. Another Struct for Comparison ---
// Admin also embeds BaseEntity, but adds a Role.
type Admin struct {
	BaseEntity // Reusing the same base code
	Role       string
}

func main() {
	// A. Initialization
	// Note: We still have to initialize the inner struct explicitly.
	user := User{
		BaseEntity: BaseEntity{
			ID:        101,
			CreatedAt: time.Now(),
		},
		Username: "dev_jane",
		Email:    "jane@example.com",
	}

	// B. Accessing Promoted Fields
	// We can access 'ID' directly on 'user', even though it lives in 'BaseEntity'.
	fmt.Println("Direct Access:", user.ID)
	// You can also access it the long way if you really want to:
	fmt.Println("Explicit Access:", user.BaseEntity.ID)

	// C. Accessing Promoted Methods
	// 'Describe' belongs to BaseEntity, but User can call it directly.
	fmt.Print("Method Promotion: ")
	user.Describe()

	// D. Using the unique User method
	user.Greeting()

	// E. Using Admin (reuses the same base logic)
	admin := Admin{
		BaseEntity: BaseEntity{ID: 999, CreatedAt: time.Now()},
		Role:       "SuperUser",
	}
	fmt.Printf("Admin Role: %s, Admin ID: %d\n", admin.Role, admin.ID)
}
