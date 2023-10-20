package main

import "fmt"

//Scientists have discovered a species of fire-breathing dragons. DNA analysis of the dragon reveals that it is a reptile evolved from a common ancestor of crocodile, hundreds of millions of years ago.
//Even though they are related, the different reptile species cannot cross-breed.
//Researchers would like to develop a lifecycle model of this rare species, in order to better study them. Complete the implementation below so that:
//• The FireDragon species implements the Reptile interface.
//• When a ReptileEgg hatches, a new reptile will be created of the same species that laid the egg.
//• nil is returned if a ReptileEgg tries to hatch more than once.

// Reptile is an interface that defines the behavior of reptilian creatures.
// It requires the reptilian creature to have a Lay method which produces an egg.
type Reptile interface {
	Lay() ReptileEgg
}

// ReptileCreator is a type which defines a function that creates and returns a new reptile.
type ReptileCreator func() Reptile

// ReptileEgg represents an egg from which a reptile can hatch.
// It contains a creation function that defines how to create a specific type of reptile,
// and a boolean flag to track if the egg has been hatched before.
type ReptileEgg struct {
	CreateReptile ReptileCreator
	hatched       bool
}

// Hatch is a method on a ReptileEgg which allows the egg to hatch.
// If the egg has been hatched before, it returns nil.
// Otherwise, it uses the creation function to create a new reptile and returns it.
func (egg *ReptileEgg) Hatch() Reptile {
	if egg.hatched {
		return nil
	}

	egg.hatched = true
	return egg.CreateReptile()
}

// FireDragon is a type which represents a specific species of reptile.
type FireDragon struct {
}

// Lay is a method on FireDragon which produces an egg.
// The egg contains a creation function specific to FireDragon.
func (f FireDragon) Lay() ReptileEgg {
	return ReptileEgg{
		CreateReptile: func() Reptile {
			return FireDragon{}
		},
	}
}

func main() {
	// Create an instance of FireDragon.
	fireDragon := FireDragon{}
	// Use the FireDragon's Lay method to produce an egg.
	egg := fireDragon.Lay()
	// Use the egg's Hatch method to produce a new FireDragon.
	childDragon := egg.Hatch()
	// Print the new FireDragon. This will print something like `main.FireDragon`.
	fmt.Println(childDragon)

	// Try to hatch the egg again.
	secondChild := egg.Hatch()
	// Since the egg has been hatched before, it should return nil.
	if secondChild == nil {
		fmt.Println("The egg has already been hatched!")
	}
}
