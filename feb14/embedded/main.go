package main

import "fmt"

type SSN string
type Income int

func (i Income) String() string {
	return fmt.Sprintf("$%d.00", i)
}

func (s SSN) String() string {
	return "XXX-XX-" + string(s)[7:]
}

func (s SSN) GetRegion() string {
	return "USA"
}

type Person struct {
	name string
	// Propogate all SSN methods up to Person as well
	SSN
	Income
}

// Explicit implementation of .String() for this type
func (p Person) String() string {
	return p.name + "(" + p.SSN.String() + ") earns " + p.Income.String()
}

func main() {
	var sufi = Person{name: "Sufi", SSN: "867-53-0909", Income: 350}

	fmt.Println(sufi.GetRegion())
	fmt.Println(sufi)
}
