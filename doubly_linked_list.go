package main
import (
    "fmt")


type Cell struct {
    data string
    next *Cell
    prev *Cell
}


type DoublyLinkedList struct {
    top_sentinel *Cell
    bottom_sentinel *Cell
}


func make_doubly_linked_list() DoublyLinkedList {
    list := DoublyLinkedList {}
    list.top_sentinel = &Cell { "top_sentinel", nil, nil }
    list.bottom_sentinel = &Cell { "bottom_sentinel", nil, list.top_sentinel }
    list.top_sentinel.next = list.bottom_sentinel
    return list
}


// Add a cell after me.
func (me *Cell) add_after(after *Cell) {
    after.next = me.next
    after.prev = me
    me.next = after
    after.next.prev = after
}


// Add a cell before me.
func (me *Cell) add_before(before *Cell) {
    me.prev.add_after(before)
}


// Delete me.
func (me *Cell) delete() *Cell {
    if me == nil {
        panic("Trying to delete a non-existent cell")
    }
    me.prev.next = me.next
    me.next.prev = me.prev
    me.next = nil
    me.prev = nil
    return me
}


// Append values from an array to the linked list
func (list *DoublyLinkedList) add_range(values []string) {
    for i := 0; i < len(values); i++ {
        value_cell := &Cell { values[i], nil, nil }
        list.bottom_sentinel.add_before(value_cell)
    }
}


func (list *DoublyLinkedList) to_string(separator string) string {
    str := ""
    cell := list.top_sentinel.next
    if cell != list.bottom_sentinel {
        str = cell.data
        cell = cell.next
    }
    for cell != list.bottom_sentinel {
        str += (separator + cell.data)
        cell = cell.next
    }
    return str
}


func (list *DoublyLinkedList) length() (count int) {
    for cell := list.top_sentinel.next; cell != list.bottom_sentinel; cell = cell.next {
        count += 1
    }
    return
}


func (list *DoublyLinkedList) is_empty() bool {
    return list.top_sentinel.next == list.bottom_sentinel
    //for doubly linked list list.bottom_sentinel.prev == list.top_sentinel could also be returned
}


func (list *DoublyLinkedList) contains(value string) bool {
    for cell := list.top_sentinel.next; cell != list.bottom_sentinel; cell = cell.next {
        if cell.data == value {
            return true
        }
    }
    return false
}


// Returns the cell containing the value
// Returns nil if the list doesn't contain the value
func (list *DoublyLinkedList) find(value string) *Cell {
    for cell := list.top_sentinel.next; cell != list.bottom_sentinel; {
        if cell.data == value {
            return cell
        }
        cell = cell.next
    }
    return nil
}


func (list *DoublyLinkedList) remove(value string) bool {
    cell := list.find(value)
    if cell != nil {
        cell.delete()
        return true
    }
    return false
}


func (list *DoublyLinkedList) push_top(value string) {
    list.top_sentinel.add_after(&Cell { value, nil, nil })
}


func (list *DoublyLinkedList) pop_top() string {
    if list.is_empty() {
        panic("Trying to pop from an empty list")
    }
    return list.top_sentinel.next.delete().data
}


func (list *DoublyLinkedList) push_bottom(value string) {
    list.bottom_sentinel.add_before(&Cell { value, nil, nil })
}


func (list *DoublyLinkedList) pop_bottom() string {
    if list.is_empty() {
        panic("Trying to pop from an empty list")
    }
    return list.bottom_sentinel.prev.delete().data
}


func (list *DoublyLinkedList) enqueue(value string) {
    list.push_top(value)
}


func (list *DoublyLinkedList) dequeue() string {
    return list.pop_bottom()
}


func main() {
    // Test queue functions.
    fmt.Printf("*** Queue Functions ***\n")
    queue := make_doubly_linked_list()
    queue.enqueue("Agate")
    queue.enqueue("Beryl")
    fmt.Printf("%s ", queue.dequeue())
    queue.enqueue("Citrine")
    fmt.Printf("%s ", queue.dequeue())
    fmt.Printf("%s ", queue.dequeue())
    queue.enqueue("Diamond")
    queue.enqueue("Emerald")
    for !queue.is_empty() {
        fmt.Printf("%s ", queue.dequeue())
    }
    fmt.Printf("\n\n")

    // Test deque functions. Names starting
    // with F have a fast pass.
    fmt.Printf("*** Deque Functions ***\n")
    deque := make_doubly_linked_list()
    deque.push_top("Ann")
    deque.push_top("Ben")
    fmt.Printf("%s ", deque.pop_bottom())
    deque.push_bottom("F-Cat")
    fmt.Printf("%s ", deque.pop_bottom())
    fmt.Printf("%s ", deque.pop_bottom())
    deque.push_bottom("F-Dan")
    deque.push_top("Eva")
    for !deque.is_empty() {
        fmt.Printf("%s ", deque.pop_bottom())
    }
    fmt.Printf("\n")
}
