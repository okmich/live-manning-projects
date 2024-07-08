package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Employee struct {
	name    string
	phone   string
	deleted bool
}

type LinearProbingHashTable struct {
	capacity  int
	employees []*Employee
}

// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func NewLinearProbingHashTable(capacity int) *LinearProbingHashTable {
	employees := make([]*Employee, capacity)
	hashTable := LinearProbingHashTable{capacity: capacity, employees: employees}
	return &hashTable
}

func (hashTable *LinearProbingHashTable) dump() {
	for idx, emp := range hashTable.employees {
		if emp == nil {
			fmt.Printf("%2d: ---\n", idx)
		} else if emp.deleted {
			fmt.Printf("%2d: xxx\n", idx)
		} else {
			fmt.Printf("%2d: %-15s\t %v\n", idx, emp.name, emp.phone)
		}
	}
	fmt.Println()
}

// Return the key's index or where it would be if present and
// the probe sequence length.
// If the key is not present and the table is full, return -1 for the index.
func (hashTable *LinearProbingHashTable) find(name string) (int, int) {
	hash := hash(name)
	index := hash % hashTable.capacity
	deletedIndex := -1
	probeSequenceLength := 0
	for probeSequenceLength < hashTable.capacity {
		index = (hash + probeSequenceLength) % hashTable.capacity
		if hashTable.employees[index] == nil {
			if deletedIndex >= 0 {
				return deletedIndex, probeSequenceLength + 1
			} else {
				return index, probeSequenceLength + 1
			}
		} else if hashTable.employees[index].deleted {
			if deletedIndex < 0 {
				deletedIndex = index
			}
		}
		if hashTable.employees[index].name == name {
			return index, probeSequenceLength + 1
		}
		probeSequenceLength++
	}

	return -1, probeSequenceLength + 1
}

// Add an item to the hash table.
func (hashTable *LinearProbingHashTable) set(name string, phone string) {
	index, _ := hashTable.find(name)
	if index < 0 {
		panic("Cannot find employee with name " + name)
	}
	if hashTable.employees[index] == nil || hashTable.employees[index].deleted {
		emp := Employee{name: name, phone: phone}
		hashTable.employees[index] = &emp
	} else {
		hashTable.employees[index].phone = phone
	}
}

// Return an item from the hash table.
func (hashTable *LinearProbingHashTable) get(name string) string {
	index, _ := hashTable.find(name)
	if index < 0 {
		return ""
	} else if hashTable.employees[index] == nil || hashTable.employees[index].deleted {
		return ""
	} else {
		return hashTable.employees[index].phone
	}
}

// Return true if the person is in the hash table.
func (hashTable *LinearProbingHashTable) contains(name string) bool {
	index, _ := hashTable.find(name)
	if index < 0 {
		return false
	} else if hashTable.employees[index] == nil || hashTable.employees[index].deleted {
		return false
	} else {
		return true
	}
}

// Delete this key's entry.
func (hashTable *LinearProbingHashTable) delete(name string) {
	// See where the key belongs.
	index, _ := hashTable.find(name)

	if index >= 0 && hashTable.employees[index] != nil && !hashTable.employees[index].deleted {
		hashTable.employees[index].deleted = true
	}
}

// Make a display showing whether each array entry is nil.
func (hashTable *LinearProbingHashTable) dumpConcise() {
	// Loop through the array.
	for i, employee := range hashTable.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else if employee.deleted {
			// Display this entry.
			fmt.Printf("x")
		} else {
			// Display this entry.
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Return the average probe sequence length for the items in the table.
func (hashTable *LinearProbingHashTable) aveProbeSequenceLength() float32 {
	totalLength := 0
	numValues := 0
	for _, employee := range hashTable.employees {
		if employee != nil {
			_, probeLength := hashTable.find(employee.name)
			totalLength += probeLength
			numValues++
		}
	}
	return float32(totalLength) / float32(numValues)
}

// Show this key's probe sequence.
func (hashTable *LinearProbingHashTable) probe(name string) int {
	// Hash the key.
	hash := hash(name) % hashTable.capacity
	fmt.Printf("Probing %s (%d)\n", name, hash)

	// Keep track of a deleted spot if we find one.
	deletedIndex := -1

	// Probe up to hashTable.capacity times.
	for i := 0; i < hashTable.capacity; i++ {
		index := (hash + i) % hashTable.capacity

		fmt.Printf("    %d: ", index)
		if hashTable.employees[index] == nil {
			fmt.Printf("---\n")
		} else if hashTable.employees[index].deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%s\n", hashTable.employees[index].name)
		}

		// If this spot is empty, the value isn't in the table.
		if hashTable.employees[index] == nil {
			// If we found a deleted spot, return its index.
			if deletedIndex >= 0 {
				fmt.Printf("    Returning deleted index %d\n", deletedIndex)
				return deletedIndex
			}

			// Return this index, which holds nil.
			fmt.Printf("    Returning nil index %d\n", index)
			return index
		}

		// If this spot is deleted, remember where it is.
		if hashTable.employees[index].deleted {
			if deletedIndex < 0 {
				deletedIndex = index
			}
		} else if hashTable.employees[index].name == name {
			// If this cell holds the key, return its data.
			fmt.Printf("    Returning found index %d\n", index)
			return index
		}

		// Otherwise continue the loop.
	}

	// If we get here, then the key is not
	// in the table and the table is full.

	// If we found a deleted spot, return it.
	if deletedIndex >= 0 {
		fmt.Printf("    Returning deleted index %d\n", deletedIndex)
		return deletedIndex
	}

	// There's nowhere to put a new entry.
	fmt.Printf("    Table is full\n")
	return -1
}

func main() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101", false},
		Employee{"Bob Baker", "202-555-0102", false},
		Employee{"Cindy Cant", "202-555-0103", false},
		Employee{"Dan Deever", "202-555-0104", false},
		Employee{"Edwina Eager", "202-555-0105", false},
		Employee{"Fred Franklin", "202-555-0106", false},
		Employee{"Gina Gable", "202-555-0107", false},
	}

	hashTable := NewLinearProbingHashTable(10)
	for _, employee := range employees {
		hashTable.set(employee.name, employee.phone)
	}
	hashTable.dump()

	hashTable.probe("Hank Hardy")
	fmt.Printf("Table contains Sally Owens: %t\n", hashTable.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hashTable.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hashTable.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hashTable.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
	hashTable.dump()

	hashTable.probe("Ann Archer")
	hashTable.probe("Bob Baker")
	hashTable.probe("Cindy Cant")
	hashTable.probe("Dan Deever")
	hashTable.probe("Edwina Eager")
	hashTable.probe("Fred Franklin")
	hashTable.probe("Gina Gable")
	hashTable.set("Hank Hardy", "202-555-0108")
	hashTable.probe("Hank Hardy")

	// Look at clustering.
	fmt.Println(time.Now())                   // Print the time so it will compile if we use a fixed seed.
	random := rand.New(rand.NewSource(12345)) // Initialize with a fixed seed
	// random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	bigCapacity := 1009
	bigHashTable := NewLinearProbingHashTable(bigCapacity)
	numItems := int(float32(bigCapacity) * 0.9)
	for i := 0; i < numItems; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		bigHashTable.set(str, str)
	}
	bigHashTable.dumpConcise()
	fmt.Printf("Average probe sequence length: %f\n",
		bigHashTable.aveProbeSequenceLength())
}
