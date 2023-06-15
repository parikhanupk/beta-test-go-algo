package main
import (
    "fmt"
    "math/rand"
    "time"
    "strconv")


func partition(arr []int) int {
    var lo, hi = 0, len(arr) - 1;
    var pivot = arr[hi];
    var i = lo - 1;
    for j := lo; j < hi; j++ {
        if arr[j] <= pivot {
            i = i + 1;
            arr[i], arr[j] = arr[j], arr[i];
        }
    }
    i = i + 1;
    arr[i], arr[hi] = arr[hi], arr[i];
    return i;
}


func quicksort(arr []int) {
    if len(arr) < 2 {
        return;
    }
    var p = partition(arr[0 :]);
    quicksort(arr[: p])
    quicksort(arr[p + 1 :])
}


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


func binary_search(arr []int, target int) (index, num_tests int) {
    var lo, hi = 0, len(arr) - 1
    for lo <= hi {
        num_tests += 1;
        index = (lo + hi) / 2;
        if arr[index] == target {
            return;
        } else {
            if arr[index] < target {
                lo = index + 1;
            } else {
                hi = index - 1;
            }
        }
    }
    return -1, num_tests;
}


func main() {
    rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make an array and display it after sorting.
    arr := make_random_array(num_items, max)
    quicksort(arr);
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
        var index, num_tests = binary_search(arr, target);
        if index < 0 || index >= len(arr) {
            fmt.Printf("Target %d not found, %d tests\n", target, num_tests);
        } else {
            fmt.Printf("arr[%d] = %d, %d tests\n", index, arr[index], num_tests);
        }
    }
}
