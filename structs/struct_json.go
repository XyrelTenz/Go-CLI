package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 1. The Base Struct
// This struct holds common fields.
type User struct {
	// 'json:"name"' tells Go: "When converting to JSON, rename 'ID' to 'user_id'"
	ID    int    `json:"user_id"`
	Name  string `json:"full_name"`
	Email string `json:"email_address,omitempty"` // omitempty: hide if empty
}

// 2. The Embedding Struct
// This struct embeds 'User' and adds more fields.
type Admin struct {
	// We embed the User struct directly here.
	// This is like "inheriting" all User fields.
	User

	// These are specific to Admin
	Permissions []string `json:"permissions"`
	Level       int      `json:"admin_level"`
}

func main() {
	// --- A. CREATING THE STRUCT ---
	fmt.Println("--- 1. Struct Creation ---")

	// We can initialize the embedded struct directly inside.
	myAdmin := Admin{
		User: User{
			ID:    101,
			Name:  "Mai Long",
			Email: "mai@example.com",
		},
		Permissions: []string{"ban_users", "delete_posts"},
		Level:       99,
	}

	// Because of embedding, we can access User fields directly on Admin!
	// We don't need to type myAdmin.User.Name (though we can).
	fmt.Printf("Admin Name: %s\n", myAdmin.Name) // Accessed directly!
	fmt.Printf("Admin Level: %d\n", myAdmin.Level)

	// --- B. STRUCT TO JSON (Marshaling) ---
	fmt.Println("\n--- 2. Struct -> JSON ---")

	jsonData, err := json.MarshalIndent(myAdmin, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Notice: The User fields are NOT nested inside a "User" object.
	// They are "flattened" into the main JSON object.
	fmt.Println(string(jsonData))

	// --- C. JSON TO STRUCT (Unmarshaling) ---
	fmt.Println("\n--- 3. JSON -> Struct ---")

	rawJson := `
	{
		"user_id": 500,
		"full_name": "New Admin",
		"permissions": ["view_only"],
		"admin_level": 1
	}`

	var newAdmin Admin
	// The JSON decoder is smart enough to map "user_id" -> Admin.User.ID
	err = json.Unmarshal([]byte(rawJson), &newAdmin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded Admin: %s (ID: %d)\n", newAdmin.Name, newAdmin.ID)
}
