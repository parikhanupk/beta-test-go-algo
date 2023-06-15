package main
import (
    "fmt"
    "math/rand"
    "time"
    "strconv")


type Customer struct {
    id string;
    num_purchases int;
}


func make_random_array(num_items, max int) []Customer {
    var arr = make([]Customer, num_items);
    for i := 0; i < num_items; i++ {
        arr[i].id = "C" + strconv.Itoa(i);
        arr[i].num_purchases = rand.Intn(max);
    }
    return arr;
}


func print_array(arr []Customer, num_items int) {
    if len(arr) <= num_items {
        fmt.Println(arr);
    } else {
        fmt.Println(arr[:num_items]);
    }
}


func check_sorted(arr []Customer) {
    for i := 1; i < len(arr); i++ {
        if arr[i - 1].num_purchases > arr[i].num_purchases {
            fmt.Println("The array is NOT sorted!");
            return;
        }
    }
    fmt.Println("The array is sorted");
}


func counting_sort(arr []Customer, max int) []Customer {
    var counts = make([]int, max); //all values will be zero in go
    for i := 0; i < len(arr); i++ {
        counts[arr[i].num_purchases] += 1;
    }
    for i := 1; i < max; i++ {
        counts[i] += counts[i - 1];
    }
    var rarr = make([]Customer, len(arr));
    for i := len(arr) - 1; i >= 0; i-- {
        counts[arr[i].num_purchases]--;
        rarr[counts[arr[i].num_purchases]] = arr[i];
    }
    return rarr;
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
    sorted := counting_sort(arr, max)
    print_array(sorted, 40)

    // Verify that it's sorted.
    check_sorted(sorted)
}
