// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	switch {
	case r.Name == "":
		return false
	case r.Address == nil:
		return false
	case r.Address["street"] == "":
		return false
	default:
		return true
	}
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Age = 0
	r.Name = ""
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var count int = 0
	for _, r := range residents {
		if r.HasRequiredInfo() {
			count++
		}
	}
	return count
}
