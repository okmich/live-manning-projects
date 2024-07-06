package main

import "fmt"

type Employee struct {
	name  string
	phone string
}

type ChainingHashTable struct {
	numBuckets int
	buckets    [][]*Employee
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

func NewChainingHashTable(numBuckets int) *ChainingHashTable {
	buckets := make([][]*Employee, numBuckets)
	hashTable := ChainingHashTable{numBuckets: numBuckets, buckets: buckets}
	return &hashTable
}

// Display the hash table's contents.
func (hashTable *ChainingHashTable) dump() {
	for idx, bucket := range hashTable.buckets {
		fmt.Printf("Bucket %v:\n", idx)
		for _, employeePointer := range bucket {
			fmt.Printf("\t%v : %v\n", employeePointer.name, employeePointer.phone)
		}
	}
}

// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (hashTable *ChainingHashTable) find(name string) (int, int) {
	bucketIdx := hash(name) % hashTable.numBuckets
	bucket := hashTable.buckets[bucketIdx]
	for idx, emp := range bucket {
		if emp.name == name {
			return bucketIdx, idx
		}
	}

	return bucketIdx, -1
}

// Add an item to the hash table.
func (hashTable *ChainingHashTable) set(name string, phone string) {
	bucketIdx, index := hashTable.find(name)
	bucket := hashTable.buckets[bucketIdx]
	if index < 0 {
		emp := Employee{name: name, phone: phone}
		bucket := append(bucket, &emp)
		hashTable.buckets[bucketIdx] = bucket
	} else {
		bucket[index].phone = phone
	}
}

// Return an item from the hash table.
func (hashTable *ChainingHashTable) get(name string) string {
	bucketIdx, index := hashTable.find(name)
	bucket := hashTable.buckets[bucketIdx]
	if index < 0 {
		return ""
	} else {
		return bucket[index].phone
	}
}

// Return true if the person is in the hash table.
func (hashTable *ChainingHashTable) contains(name string) bool {
	_, index := hashTable.find(name)
	return index >= 0
}

// Delete this key's entry.
func (hashTable *ChainingHashTable) delete(name string) {
	bucketIdx, index := hashTable.find(name)
	if index > 0 {
		bucket := hashTable.buckets[bucketIdx]
		bucket = append(bucket[:index], bucket[index+1:]...)
	}
}

func main() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101"},
		Employee{"Bob Baker", "202-555-0102"},
		Employee{"Cindy Cant", "202-555-0103"},
		Employee{"Dan Deever", "202-555-0104"},
		Employee{"Edwina Eager", "202-555-0105"},
		Employee{"Fred Franklin", "202-555-0106"},
		Employee{"Gina Gable", "202-555-0107"},
		Employee{"Herb Henshaw", "202-555-0108"},
		Employee{"Ida Iverson", "202-555-0109"},
		Employee{"Jeb Jacobs", "202-555-0110"},
	}

	hashTable := NewChainingHashTable(10)
	for _, employee := range employees {
		hashTable.set(employee.name, employee.phone)
	}
	hashTable.dump()

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
}
