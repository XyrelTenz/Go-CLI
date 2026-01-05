package pointers

import "fmt"

// 1. BASIC: The Struct
type Student struct {
	Name string
	GPA  float64
}

// 2. INTERMEDIATE: Modifying via Pointer
func UpgradeGPA(s *Student) {
	s.GPA += 0.5 // Direct mutation of the original data
}

// This is used when you need to change WHERE a pointer points to.
// Common in Linked List insertion or Tree manipulation.
func SwapStudents(s1, s2 **Student) {
	// We are swapping the pointers themselves, not the values inside.
	temp := *s1
	*s1 = *s2
	*s2 = temp
}

func RunDemo() {
	// A. Value vs Pointer
	s1 := &Student{Name: "Xyrel", GPA: 3.0}
	s2 := &Student{Name: "BaoBao", GPA: 2.0}

	fmt.Printf("Before Upgrade: %.1f\n", s1.GPA)
	UpgradeGPA(s1)
	fmt.Printf("After Upgrade: %.1f\n", s1.GPA)

	// B. Swapping Pointers
	fmt.Println("--- Pointer Swap ---")
	fmt.Printf("s1 points to: %s, s2 points to: %s\n", s1.Name, s2.Name)

	// Pass the ADDRESS of the pointers (&s1)
	SwapStudents(&s1, &s2)

	fmt.Printf("s1 points to: %s, s2 points to: %s\n", s1.Name, s2.Name)
}
