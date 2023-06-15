package main
import (
    "fmt"
    "math/rand"
    "time")


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


func check_sorted(arr []int) {
    for i := 1; i < len(arr); i++ {
        if arr[i - 1] > arr[i] {
            fmt.Println("The array is NOT sorted!");
            return;
        }
    }
    fmt.Println("The array is sorted");
}


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

    // Sort and display the result.
    quicksort(arr)
    print_array(arr, 40)

    // Verify that it's sorted.
    check_sorted(arr)
}
