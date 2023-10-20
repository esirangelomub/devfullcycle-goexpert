package main

//Implement the findAllHobbyists function that takes a hobby, and a map consisting of people's names mapped to their respective hobbies. The function should return a slice containing the names of the people (in any order) who enjoy the hobby.
//For example, the following code should display the name 'Chad'.
//
//hobbies := map[string] []string {
//	"Steve": Ilstring{"Fashion",
//	"Patty": []string{"Drama", "Magic",
//	"Chad": Ilstring{"Puzzles", "Pets", "Yoga"},
//}
//fmt.PrintIn(findAllHobbyists ("Yoga", hobbies))

import "fmt"

// findAllHobbyists is a function that searches for all the people who have a specific hobby.
// It takes a string 'hobby' and a map 'hobbies' where the keys are people's names and the values are slices of their hobbies.
// The function returns a slice containing the names of the people who have the specified hobby.
func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
	// An empty slice to store the names of people who have the desired hobby.
	var hobbyists []string

	// Loop through each person and their hobbies in the map.
	for person, personHobbies := range hobbies {
		// Loop through each hobby of the current person.
		for _, h := range personHobbies {
			// Check if the current hobby matches the desired hobby.
			if h == hobby {
				// If it matches, append the person's name to the hobbyists slice.
				hobbyists = append(hobbyists, person)
				// Break out of the inner loop since we found a match for the current person.
				break
			}
		}
	}
	return hobbyists
}

func main() {
	// A sample map of people and their hobbies.
	hobbies := map[string][]string{
		"Steve": []string{"Fashion", "Piano", "Reading"},
		"Patty": []string{"Drama", "Magic", "Pets"},
		"Chad":  []string{"Puzzles", "Pets", "Yoga"},
	}

	// Call the findAllHobbyists function with the hobby "Yoga" and print the result.
	// Expected output: ["Chad"].
	fmt.Println(findAllHobbyists("Yoga", hobbies))
}
