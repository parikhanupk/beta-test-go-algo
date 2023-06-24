package main
import (
    "fmt")


// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
    hash := 5381
    for _, ch := range value {
        hash = ((hash << 5) + hash) + int(ch)
    }

    // Make sure the result is non-negative.
    if hash < 0 { hash = -hash }
    return hash
}


type Employee struct {
    name string
    phone string
}


type ChainingHashTable struct {
    num_buckets int
    num_entries int
    buckets [][]*Employee
}


// Initialize a ChainingHashTable and return a pointer to it.
func NewChainingHashTable(num_buckets int) *ChainingHashTable {
    return &ChainingHashTable { num_buckets, 0, make([][]*Employee, num_buckets)}
}


// Display the hash table's contents.
func (hash_table *ChainingHashTable) dump() {
    for index, bucket := range hash_table.buckets {
        fmt.Printf("Bucket %d:\n", index)
        for _, employee := range bucket {
            fmt.Printf("\t%s: %s\n", employee.name, employee.phone)
        }
    }
}


// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (hash_table *ChainingHashTable) find(name string) (int, int) {
    bi := hash(name)
    bi %= hash_table.num_buckets
    for ei, employee := range hash_table.buckets[bi] {
        if employee.name == name {
            return bi, ei
        }
    }
    return bi, -1
}


// Add an item to the hash table.
func (hash_table *ChainingHashTable) set(name string, phone string) {
    bi, ei := hash_table.find(name)
    if ei >= 0 {
        hash_table.buckets[bi][ei].phone = phone
    } else {
        hash_table.buckets[bi] = append(hash_table.buckets[bi], &Employee { name, phone })
        hash_table.num_entries += 1
    }
}


// Return an item from the hash table.
func (hash_table *ChainingHashTable) get(name string) (string) {
    bi, ei := hash_table.find(name)
    if ei >= 0 {
        return hash_table.buckets[bi][ei].phone
    } else {
        return ""
    }
}


// Return true if the person is in the hash table.
func (hash_table *ChainingHashTable) contains(name string) (bool) {
    _, ei := hash_table.find(name)
    return ei >= 0
}


// Delete this key's entry.
func (hash_table *ChainingHashTable) delete(name string) {
    bi, ei := hash_table.find(name)
    if ei >= 0 {
        hash_table.buckets[bi] = append(hash_table.buckets[bi][:ei], hash_table.buckets[bi][ei + 1:]...)
        hash_table.num_entries -= 1
    }
}


func main() {
    // Make some names.
    employees := []Employee {
        Employee { "Ann Archer",    "202-555-0101" },
        Employee { "Bob Baker",     "202-555-0102" },
        Employee { "Cindy Cant",    "202-555-0103" },
        Employee { "Dan Deever",    "202-555-0104" },
        Employee { "Edwina Eager",  "202-555-0105" },
        Employee { "Fred Franklin", "202-555-0106" },
        Employee { "Gina Gable",    "202-555-0107" },
        Employee { "Herb Henshaw",  "202-555-0108" },
        Employee { "Ida Iverson",   "202-555-0109" },
        Employee { "Jeb Jacobs",    "202-555-0110" },
    }

    hash_table := NewChainingHashTable(10)
    for _, employee := range employees {
        hash_table.set(employee.name, employee.phone)
    }
    hash_table.dump()

    fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
    fmt.Println("Deleting Dan Deever")
    hash_table.delete("Dan Deever")
    fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
    fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
    fmt.Println("Changing Fred Franklin")
    hash_table.set("Fred Franklin", "202-555-0100")
    fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
}
