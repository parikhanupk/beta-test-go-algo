package main
import (
    "fmt")


type Node struct {
    data string
    left *Node
    right *Node
}


func (node *Node) insert_value(value string) {
    value_node := &Node { value, nil, nil }
    for {
        if value <= node.data {
            if node.left == nil {
                node.left = value_node
                return
            } else {
                node = node.left
            }
        } else {
            if node.right == nil {
                node.right = value_node
                return
            } else {
                node = node.right
            }
        }
    }
}


func (node *Node) find_value(target string) *Node {
    for node != nil {
        if node.data == target {
            return node
        }
        if node.data < target {
            node = node.right
        } else {
            node = node.left
        }
    }
    return nil
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


func main() {
    // Make a root node to act as sentinel.
    root := Node { "", nil, nil }

    // Add some values.
    root.insert_value("I")
    root.insert_value("G")
    root.insert_value("C")
    root.insert_value("E")
    root.insert_value("B")
    root.insert_value("K")
    root.insert_value("S")
    root.insert_value("Q")
    root.insert_value("M")

    // Add F.
    root.insert_value("F")

    // Display the values in sorted order.
    fmt.Printf("Sorted values: %s\n", root.right.inorder())

    // Let the user search for values.
    for {
        // Get the target value.
        target := ""
        fmt.Printf("String: ")
        fmt.Scanln(&target)
        if len(target) == 0 { break }

        // Find the value's node.
        node := root.find_value(target)
        if node == nil {
            fmt.Printf("%s not found\n", target)
        } else {
            fmt.Printf("Found value %s\n", target)
        }
    }
}
