package main

import "fmt"

func main() {
	// Create a map to store student grades
	grades := map[string]int{
		"Alice":   85,
		"Bob":     90,
		"Charlie": 78,
	}

	// Display all students and their grades
	fmt.Println("Student Grades:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}

	// Retrieve a student's grade
	var studentName string
	fmt.Print("\nEnter a student's name: ")
	fmt.Scanln(&studentName)

	if grade, exists := grades[studentName]; exists {
		fmt.Printf("%s's grade: %d\n", studentName, grade)
	} else {
		fmt.Println("Student not found.")
	}

	// Update or remove a student
	var option string
	fmt.Print("\nChoose an option (update/remove/none): ")
	fmt.Scanln(&option)

	if option == "update" {
		fmt.Print("Enter the new grade: ")
		var newGrade int
		fmt.Scanln(&newGrade)
		grades[studentName] = newGrade
		fmt.Printf("Updated %s's grade to %d\n", studentName, newGrade)
	} else if option == "remove" {
		delete(grades, studentName)
		fmt.Printf("Removed %s from the records.\n", studentName)
	} else {
		fmt.Println("No changes made.")
	}

	// Display the final list of grades
	fmt.Println("\nFinal Student Grades:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
	fmt.Println("enter the name:")
}
