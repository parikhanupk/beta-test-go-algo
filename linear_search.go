package main
import (
    "fmt"
    "math/rand"
    "time"
    "strconv")


func make_random_array(num_items, max int) []int {
    var arr = make([]int, num_items);
    for i := 0; i < num_items; i++ {
        arr[i] = rand.Intn(max);
    }
    return arr;
}


func print_array(arr []int, num_items int) {
    if len(arr) <= num_items {
        fmt.Println(arr);
    } else {
        fmt.Println(arr[:num_items]);
    }
}


func linear_search(arr []int, target int) (index, num_tests int) {
    index = -1
    for i := 0; i < len(arr); i++ {
        num_tests += 1;
        if arr[i] == target {
            index = i;
            return
        }
    }
    return
}


func main() {
    rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted array.
    arr := make_random_array(num_items, max)
    print_array(arr, 40)
    fmt.Println()

    for {
        // Get the target as a string.
        var target_string string;
        fmt.Printf("Target: ")
        fmt.Scanln(&target_string)

        // If the target string is blank, break out of the loop.
        if len(target_string) == 0 { break }

        //assumes input will either be blank or a valid non-negative integer (int in this case)
        var target, _ = strconv.Atoi(target_string);
        var index, num_tests = linear_search(arr, target);
        if index < 0 || index >= len(arr) {
            fmt.Printf("Target %d not found, %d tests\n", target, num_tests);
        } else {
            fmt.Printf("arr[%d] = %d, %d tests\n", index, arr[index], num_tests);
        }
    }
}
