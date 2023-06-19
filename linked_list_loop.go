package main
import (
    "fmt")


type Cell struct {
    data string
    next *Cell
}


type LinkedList struct {
    sentinel *Cell
}


func make_linked_list() LinkedList {
    list := LinkedList {}
    list.sentinel = &Cell { "SENTINEL", nil }
    return list
}


// Add a cell after me.
func (me *Cell) add_after(after *Cell) {
    after.next = me.next
    me.next = after
}


// Delete a cell after me.
func (me *Cell) delete_after() *Cell {
    deleted := me.next
    if deleted == nil {
        panic("Trying to delete a non-existent cell")
    }
    me.next = deleted.next
    return deleted
}


// Append values from an array to the linked list
func (list *LinkedList)add_range(values []string) {
    last_cell := list.sentinel
    for ; last_cell.next != nil; last_cell = last_cell.next {}
    for i := 0; i < len(values); i++ {
        value_cell := &Cell { values[i], nil }
        last_cell.add_after(value_cell)
        last_cell = value_cell
    }
}


func (list *LinkedList) to_string(separator string) string {
    str := ""
    cell := list.sentinel.next
    if cell != nil {
        str = cell.data
        cell = cell.next
    }
    for cell != nil {
        str += (separator + cell.data)
        cell = cell.next
    }
    return str
}


func (list *LinkedList) length() (count int) {
    for cell := list.sentinel.next; cell != nil; cell = cell.next {
        count += 1
    }
    return
}


func (list *LinkedList) is_empty() bool {
    return list.sentinel.next == nil
}


func (list *LinkedList) contains(value string) bool {
    for cell := list.sentinel.next; cell != nil; cell = cell.next {
        if cell.data == value {
            return true
        }
    }
    return false
}


// Returns the cell before the cell containing the value
// Returns nil if the list doesn't contain the value
func (list *LinkedList) find(value string) *Cell {
    for cell := list.sentinel; cell.next != nil; {
        if cell.next.data == value {
            return cell
        }
        cell = cell.next
    }
    return nil
}


func (list *LinkedList) remove(value string) bool {
    cell := list.find(value)
    if cell != nil {
        cell.delete_after()
        return true
    }
    return false
}


func (list *LinkedList) push(value string) {
    list.sentinel.add_after(&Cell { value, nil })
}


func (list *LinkedList) pop() string {
    return list.sentinel.delete_after().data
}


func (list *LinkedList) has_loop() bool {
    slow := list.sentinel.next
    fast := list.sentinel.next
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
        if fast == slow {
            return true
        }
    }
    return false
}


func (list *LinkedList) to_string_max(separator string, max int) string {
    str := ""
    count := 0
    cell := list.sentinel.next
    if cell != nil {
        str = cell.data
        cell = cell.next
        count += 1
    }
    for cell != nil {
        str += (separator + cell.data)
        cell = cell.next
        count += 1
        if count >= max {
            break
        }
    }
    return str
}


func main() {
    // Make a list from a slice of values.
    values := []string {
        "0", "1", "2", "3", "4", "5",
    }
    list := make_linked_list()
    list.add_range(values)

    fmt.Println(list.to_string(" "))
    if list.has_loop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 5 point to cell 2.
    list.sentinel.next.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.to_string_max(" ", 10))
    if list.has_loop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
    fmt.Println()

    // Make cell 4 point to cell 2.
    list.sentinel.next.next.next.next.next = list.sentinel.next.next

    fmt.Println(list.to_string_max(" ", 10))
    if list.has_loop() {
        fmt.Println("Has loop")
    } else {
        fmt.Println("No loop")
    }
}
