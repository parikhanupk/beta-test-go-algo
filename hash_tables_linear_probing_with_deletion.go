package main
import (
    "fmt"
    "math/rand")


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
    deleted bool
}


type LinearProbingHashTable struct {
    capacity int
    num_entries int
    employees []*Employee
}


// Initialize a LinearProbingHashTable and return a pointer to it.
func NewLinearProbingHashTable(capacity int) *LinearProbingHashTable {
    return &LinearProbingHashTable { capacity, 0, make([]*Employee, capacity) }
}


// Display the hash table's contents.
func (hash_table *LinearProbingHashTable) dump() {
    for index, employee := range hash_table.employees {
        if employee == nil {
            fmt.Printf("%3d: ---\n", index)
        } else if employee.deleted {
            fmt.Printf("%3d: xxx\n", index)
        } else {
            fmt.Printf("%3d: %-20s %s\n", index, employee.name, employee.phone)
        }
    }
}


func (hash_table *LinearProbingHashTable) find(name string) (int, int) {
    hash := hash(name) % hash_table.capacity
    deleted_index := -1
    for i, index := 0, 0; i < hash_table.capacity; i++ {
        index = (hash + i) % hash_table.capacity
        if hash_table.employees[index] == nil {
            if deleted_index >= 0 {
                return deleted_index, i + 1
            } else {
                return index, i + 1
            }
        }
        if hash_table.employees[index].deleted {
            if deleted_index == -1 {
                deleted_index = index
            }
        } else {
            if hash_table.employees[index].name == name {
                return index, i + 1
            }
        }
    }
    if deleted_index >= 0 {
        return deleted_index, hash_table.capacity
    } else {
        return -1, hash_table.capacity
    }
}


// Add an item to the hash table.
func (hash_table *LinearProbingHashTable) set(name string, phone string) {
    index, _ := hash_table.find(name)
    if index < 0 {
        panic("The hash table is full!")
    }
    if hash_table.employees[index] == nil {
        hash_table.employees[index] = &Employee { name, phone, false }
        hash_table.num_entries += 1
    } else if hash_table.employees[index].deleted {
        hash_table.employees[index] = &Employee { name, phone, false }
    } else {
        hash_table.employees[index].phone = phone
    }
}


// Return an item from the hash table.
func (hash_table *LinearProbingHashTable) get(name string) (string) {
    index, _ := hash_table.find(name)
    if index < 0 || hash_table.employees[index] == nil || hash_table.employees[index].deleted {
        return ""
    } else {
        return hash_table.employees[index].phone
    }
}


// Return true if the person is in the hash table.
func (hash_table *LinearProbingHashTable) contains(name string) (bool) {
    index, _ := hash_table.find(name)
    return index >= 0 && hash_table.employees[index] != nil && hash_table.employees[index].deleted == false
}


func (hash_table *LinearProbingHashTable) delete(name string) {
    index, _ := hash_table.find(name)
    if index >= 0 && hash_table.employees[index] != nil {
        hash_table.employees[index].deleted = true
        hash_table.num_entries -= 1
    }
}


// Show this key's probe sequence.
func (hash_table *LinearProbingHashTable) probe(name string) (int) {
    // Hash the key.
    hash := hash(name) % hash_table.capacity
    fmt.Printf("Probing %s (%d)\n", name, hash)

    // Keep track of a deleted spot if we find one.
    deleted_index := -1

    // Probe up to hash_table.capacity times.
    for i := 0; i < hash_table.capacity; i++ {
        index := (hash + i) % hash_table.capacity

        fmt.Printf("    %d: ", index)
        if hash_table.employees[index] == nil {
            fmt.Printf("---\n")
        } else if hash_table.employees[index].deleted {
            fmt.Printf("xxx\n")
        } else {
            fmt.Printf("%s\n", hash_table.employees[index].name)
        }

        // If this spot is empty, the value isn't in the table.
        if hash_table.employees[index] == nil {
            // If we found a deleted spot, return its index.
            if deleted_index >= 0 {
                fmt.Printf("    Returning deleted index %d\n", deleted_index)
                return deleted_index
            }

            // Return this index, which holds nil. 
            fmt.Printf("    Returning nil index %d\n", index)
            return index
        }

        // If this spot is deleted, remember where it is.
        if hash_table.employees[index].deleted {
            if deleted_index < 0 {
                deleted_index = index
            }
        } else if hash_table.employees[index].name == name {
            // If this cell holds the key, return its data.
            fmt.Printf("    Returning found index %d\n", index)
            return index
        }

        // Otherwise continue the loop.
    }

    // If we get here, then the key is not
    // in the table and the table is full.

    // If we found a deleted spot, return it.
    if deleted_index >= 0 {
        fmt.Printf("    Returning deleted index %d\n", deleted_index)
        return deleted_index
    }

    // There's nowhere to put a new entry.    
    fmt.Printf("    Table is full\n")
    return -1
}


// Make a display showing whether each slice entry is nil.
func (hash_table *LinearProbingHashTable) dump_concise() {
    // Loop through the slice.
    for i, employee := range hash_table.employees {
        if employee == nil {
            // This spot is empty.
            fmt.Printf(".")
        } else if employee.deleted {
            // This spot is a deleted spot
            fmt.Printf("x")
        } else {
            // Display this entry.
            fmt.Printf("O")
        }
        if i % 50 == 49 { fmt.Println() }
    }
    fmt.Println()
}


// Return the average probe sequence length for the items in the table.
func (hash_table *LinearProbingHashTable) ave_probe_sequence_length() float32 {
    total_length := 0
    num_values := 0
    for _, employee := range(hash_table.employees) {
        if employee != nil {
            _, probe_length := hash_table.find(employee.name)
            total_length += probe_length
            num_values++
       }
    }
    return float32(total_length) / float32(num_values)
}


func main() {
    // Make some names.
    employees := []Employee {
        Employee { "Ann Archer",    "202-555-0101", false },
        Employee { "Bob Baker",     "202-555-0102", false },
        Employee { "Cindy Cant",    "202-555-0103", false },
        Employee { "Dan Deever",    "202-555-0104", false },
        Employee { "Edwina Eager",  "202-555-0105", false },
        Employee { "Fred Franklin", "202-555-0106", false },
        Employee { "Gina Gable",    "202-555-0107", false },
    }

    hash_table := NewLinearProbingHashTable(10)
    for _, employee := range employees {
        hash_table.set(employee.name, employee.phone)
    }
    hash_table.dump()

    hash_table.probe("Hank Hardy")
    fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
    fmt.Println("Deleting Dan Deever")
    hash_table.delete("Dan Deever")
    fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
    fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
    fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
    fmt.Println("Changing Fred Franklin")
    hash_table.set("Fred Franklin", "202-555-0100")
    fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
    hash_table.dump()

    hash_table.probe("Ann Archer")
    hash_table.probe("Bob Baker")
    hash_table.probe("Cindy Cant")
    hash_table.probe("Dan Deever")
    hash_table.probe("Edwina Eager")
    hash_table.probe("Fred Franklin")
    hash_table.probe("Gina Gable")
    hash_table.set("Hank Hardy", "202-555-0108")
    hash_table.probe("Hank Hardy")

    // Look at clustering.
    rand.Seed(12345)
    big_capacity := 1009
    big_hash_table := NewLinearProbingHashTable(big_capacity)
    num_items := int(float32(big_capacity) * 0.9)
    for i := 0; i < num_items; i++ {
        str := fmt.Sprintf("%d-%d", i, rand.Intn(1000000))
        big_hash_table.set(str, str)
    }
    big_hash_table.dump_concise()
    fmt.Printf("Average probe sequence length: %f\n",
        big_hash_table.ave_probe_sequence_length())
}
