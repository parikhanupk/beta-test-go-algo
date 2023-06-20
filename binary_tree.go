package main
import (
    "fmt"
    "strings")


type Node struct {
    data string
    left *Node
    right *Node
}


func build_tree() *Node {
    j := &Node { "J", nil, nil }
    i := &Node { "I", nil, nil }
    h := &Node { "H", i, j }
    g := &Node { "G", nil, nil }
    f := &Node { "F", h, nil }
    e := &Node { "E", g, nil }
    d := &Node { "D", nil, nil }
    c := &Node { "C", nil, f }
    b := &Node { "B", d, e }
    a := &Node { "A", b, c }
    return a
}


func (node *Node) display_indented(indent string, depth int) string {
    result := strings.Repeat(indent, depth) + node.data + "\n"
    if node.left != nil {
        result += node.left.display_indented(indent, depth + 1)
    }
    if node.right != nil {
        result += node.right.display_indented(indent, depth + 1)
    }
    return result
}


func (node *Node) preorder() string {
    result := node.data
    if node.left != nil {
        result += " " + node.left.preorder()
    }
    if node.right != nil {
        result += " " + node.right.preorder()
    }
    return result
}


func (node *Node) inorder() string {
    result := ""
    if node.left != nil {
        result += node.left.inorder() + " "
    }
    result += node.data
    if node.right != nil {
        result += " " + node.right.inorder()
    }
    return result
}


func (node *Node) postorder() string {
    result := ""
    if node.left != nil {
        result += node.left.postorder() + " "
    }
    if node.right != nil {
        result += node.right.postorder() + " "
    }
    result += node.data
    return result
}


type Cell struct {
    data *Node
    prev *Cell
    next *Cell
}


type Queue struct {
    head_sentinel *Cell
    tail_sentinel *Cell
}


func make_queue() *Queue {
    queue := &Queue {}
    queue.head_sentinel = &Cell { &Node { "head_sentinel", nil, nil }, nil, nil }
    queue.tail_sentinel = &Cell { &Node { "tail_sentinel", nil, nil }, queue.head_sentinel, nil }
    queue.head_sentinel.next = queue.tail_sentinel
    return queue
}


func (queue *Queue) enqueue(node *Node) {
    cell := &Cell { node, queue.tail_sentinel.prev, queue.tail_sentinel }
    queue.tail_sentinel.prev.next = cell
    queue.tail_sentinel.prev = cell
}


func (queue *Queue) dequeue() *Node {
    node := queue.head_sentinel.next
    queue.head_sentinel.next = node.next
    queue.head_sentinel.next.prev = queue.head_sentinel
    return node.data
}


func (node *Node) breadth_first() string {
    queue := make_queue()
    if node != nil {
        queue.enqueue(node)
    }
    result := ""
    for queue.head_sentinel.next != queue.tail_sentinel {
        cur := queue.dequeue()
        result += cur.data
        if cur.left != nil {
            queue.enqueue(cur.left)
        }
        if cur.right != nil {
            queue.enqueue(cur.right)
        }
        if queue.head_sentinel.next != queue.tail_sentinel {
            result += " "
        }
    }
    return result
}


func main() {
    // Build a tree.
    a_node := build_tree()

    // Display with indentation.
    fmt.Println(a_node.display_indented("  ", 0))

    // Display traversals.
    fmt.Println("Preorder:     ", a_node.preorder())
    fmt.Println("Inorder:      ", a_node.inorder())
    fmt.Println("Postorder:    ", a_node.postorder())
    fmt.Println("Breadth first:", a_node.breadth_first())
}
