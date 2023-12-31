// Exhaustive search
package main
import (
    "fmt"
    "math/rand"
    "time"
    "sort")


const num_items = 250
const min_value = 1
const max_value = 10
const min_weight = 4
const max_weight = 10
var allowed_weight int


type Item struct {
    id, blocked_by  int
    block_list      []int   // Other items that this one blocks.
    value, weight   int
    is_selected     bool
}


func main() {
    items := make_items(num_items, min_value, max_value, min_weight, max_weight)
    allowed_weight = sum_weights(items, true) / 2

    // Display basic parameters.
    fmt.Println("*** Parameters ***")
    fmt.Printf("# items: %d\n", num_items)
    fmt.Printf("Total value: %d\n", sum_values(items, true))
    fmt.Printf("Total weight: %d\n", sum_weights(items, true))
    fmt.Printf("Allowed weight: %d\n", allowed_weight)
    print_items(items, true)
    fmt.Println()

    // Exhaustive search
    if num_items > 25 {    // Only run exhaustive search if num_items <= 25.
        fmt.Println("Too many items for exhaustive search\n")
    } else {
        fmt.Println("*** Exhaustive Search ***")
        run_algorithm(exhaustive_search, items, allowed_weight)
    }

    // Branch and Bound
    fmt.Println()
    if num_items > 45 {    // Only run branch and bound search if num_items <= 45.
        fmt.Println("Too many items for branch and bound search\n")
    } else {
        fmt.Println("*** Branch and Bound Search ***")
        run_algorithm(branch_and_bound, items, allowed_weight)
    }

    // Rod's technique
    fmt.Println()
    if num_items > 85 {    // Only use Rod's technique if num_items <= 85.
        fmt.Println("Too many items for Rod's technique\n")
    } else {
        fmt.Println("*** Rod's technique ***")
        run_algorithm(rods_technique, items, allowed_weight)
    }

    // Rod's technique sorted
    fmt.Println()
    if num_items > 350 {    // Only use Rod's technique sorted if num_items <= 350.
        fmt.Println("Too many items for Rod's technique sorted\n")
    } else {
        fmt.Println("*** Rod's technique sorted***")
        run_algorithm(rods_technique_sorted, items, allowed_weight)
    }
}


// Make some random items.
func make_items(num_items, min_value, max_value, min_weight, max_weight int) []Item {
    // Initialize a pseudorandom number generator.
    //random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
    random := rand.New(rand.NewSource(1337)) // Initialize with a fixed seed

    items := make([]Item, num_items)
    for i := 0; i < num_items; i++ {
        items[i] = Item {
            i,
            -1,
            nil,
            random.Intn(max_value - min_value + 1) + min_value,
            random.Intn(max_weight - min_weight + 1) + min_weight,
            false }
    }
    return items
}


// Return a copy of the items slice.
func copy_items(items []Item) []Item {
    new_items := make([]Item, len(items))
    copy(new_items, items)
    return new_items
}


// Return the total value of the items.
// If add_all is false, only add up the selected items.
func sum_values(items []Item, add_all bool) int {
    total := 0
    for i := 0; i < len(items); i++ {
        if add_all || items[i].is_selected {
            total += items[i].value
        }
    }
    return total
}


// Return the total weight of the items.
// If add_all is false, only add up the selected items.
func sum_weights(items []Item, add_all bool) int {
    total := 0
    for i := 0; i < len(items); i++ {
        if add_all || items[i].is_selected {
            total += items[i].weight
        }
    }
    return total
}


// Return the value of this solution.
func solution_value(items []Item, allowed_weight int) int {
    // If the solution's total weight > allowed_weight,
    // return 0 so we won't use this solution.
    if sum_weights(items, false) > allowed_weight { return -1 }

    // Return the sum of the selected values.
    return sum_values(items, false)
}


// Print the selected items.
func print_items(items []Item, all bool) {
    num_printed := 0
    for i, item := range items {
        if all || item.is_selected {
            fmt.Printf("%d(%d, %d) ", i, item.value, item.weight)
        }
        num_printed += 1
        if num_printed > 100 {
            fmt.Println("...")
            return
        }
    }
    fmt.Println()
}


func run_algorithm(alg func([]Item, int) ([]Item, int, int), items []Item, allowed_weight int) {
    // Copy the items so the run isn't influenced by a previous run.
    test_items := copy_items(items)

    start := time.Now()

    // Run the algorithm.
    solution, total_value, function_calls := alg(test_items, allowed_weight)

    elapsed := time.Since(start)

    fmt.Printf("Elapsed: %f\n", elapsed.Seconds())
    print_items(solution, false)
    fmt.Printf("Value: %d, Weight: %d, Calls: %d\n",
        total_value, sum_weights(solution, false), function_calls)
}


// Recursively assign values in or out of the solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func exhaustive_search(items []Item, allowed_weight int) ([]Item, int, int) {
    return do_exhaustive_search(items, allowed_weight, 0)
}


func do_exhaustive_search(items []Item, allowed_weight, next_index int) ([]Item, int, int) {
    //fmt.Println("---", next_index, items)
    if next_index >= len(items) {
        return copy_items(items), solution_value(items, allowed_weight), 1
    } else {
        items[next_index].is_selected = true
        included_solution, included_value, included_calls := do_exhaustive_search(items, allowed_weight, next_index + 1)
        //fmt.Println("inc---", included_solution, included_value)
        items[next_index].is_selected = false
        excluded_solution, excluded_value, excluded_calls := do_exhaustive_search(items, allowed_weight, next_index + 1)
        //fmt.Println("exc---", excluded_solution, excluded_value)
        if included_value >= excluded_value {
            return included_solution, included_value, included_calls + excluded_calls + 1
        } else {
            return excluded_solution, excluded_value, excluded_calls + included_calls + 1
        }
    }
}


func branch_and_bound(items []Item, allowed_weight int) ([]Item, int, int) {
    return do_branch_and_bound(items, allowed_weight, 0, 0, 0, sum_values(items, true), 0)
}


func do_branch_and_bound(items []Item, allowed_weight, best_value, current_value, current_weight, remaining_value, next_index int) ([]Item, int, int) {
    if next_index >= len(items) {
        return copy_items(items), current_value, 1
    } else {
        if current_value + remaining_value <= best_value {
            return nil, current_value, 1
        }
        included_solution, included_value, included_calls := []Item{}, 0, 1
        if current_weight + items[next_index].weight <= allowed_weight {
            items[next_index].is_selected = true
            included_solution, included_value, included_calls = do_branch_and_bound(items, allowed_weight, best_value, current_value + items[next_index].value, current_weight + items[next_index].weight, remaining_value - items[next_index].value, next_index + 1)
            if included_value > best_value {
                best_value = included_value
            }
        }
        items[next_index].is_selected = false
        excluded_solution, excluded_value, excluded_calls := do_branch_and_bound(items, allowed_weight, best_value, current_value, current_weight, remaining_value - items[next_index].value, next_index + 1)
        if included_value >= excluded_value {
            return included_solution, included_value, included_calls + excluded_calls + 1
        } else {
            return excluded_solution, excluded_value, excluded_calls + included_calls + 1
        }
    }
}


// Build the items' block lists.
func make_block_lists(items []Item) {
    for i := range items {
        items[i].block_list = []int{}
        for j := range items {
            if i != j {
                if items[i].value >= items[j].value && items[i].weight <= items[j].weight {
                    items[i].block_list = append(items[i].block_list, items[j].id)
                }
            }
        }
    }
}


// Block items on this item's blocks list.
func block_items(source Item, items []Item) {
    for _, i := range source.block_list {
        if items[i].blocked_by == -1 {
            items[i].blocked_by = source.id
        }
    }
}


// Unblock items on this item's blocks list.
func unblock_items(source Item, items []Item) {
    for _, i := range source.block_list {
        if items[i].blocked_by == source.id {
            items[i].blocked_by = -1
        }
    }
}


func rods_technique(items []Item, allowed_weight int) ([]Item, int, int) {
    make_block_lists(items)
    return do_rods_technique(items, allowed_weight, 0, 0, 0, sum_values(items, true), 0)
}


func do_rods_technique(items []Item, allowed_weight, best_value, current_value, current_weight, remaining_value, next_index int) ([]Item, int, int) {
    //fmt.Println("---", items, allowed_weight, best_value, current_value, current_weight, remaining_value, next_index)
    if next_index >= len(items) {
        return copy_items(items), current_value, 1
    } else {
        if current_value + remaining_value <= best_value {
            return nil, current_value, 1
        }
        included_solution, included_value, included_calls := []Item{}, 0, 1
        if items[next_index].blocked_by == -1 {
            if current_weight + items[next_index].weight <= allowed_weight {
                items[next_index].is_selected = true
                included_solution, included_value, included_calls = do_rods_technique(items, allowed_weight, best_value, current_value + items[next_index].value, current_weight + items[next_index].weight, remaining_value - items[next_index].value, next_index + 1)
                if included_value > best_value {
                    best_value = included_value
                }
            }
        }
        block_items(items[next_index], items)
        items[next_index].is_selected = false
        excluded_solution, excluded_value, excluded_calls := do_rods_technique(items, allowed_weight, best_value, current_value, current_weight, remaining_value - items[next_index].value, next_index + 1)
        unblock_items(items[next_index], items)
        if included_value >= excluded_value {
            return included_solution, included_value, included_calls + excluded_calls + 1
        } else {
            return excluded_solution, excluded_value, excluded_calls + included_calls + 1
        }
    }
}


func rods_technique_sorted(items []Item, allowed_weight int) ([]Item, int, int) {
    make_block_lists(items)

    // Sort so items with longer blocked lists come first.
    sort.Slice(items, func(i, j int) bool {
        return len(items[i].block_list) > len(items[j].block_list)
    })

    // Reset the items' IDs.
    for i := range items {
        items[i].id = i
    }

    // Rebuild the blocked lists with the new indices.
    make_block_lists(items)

    return do_rods_technique(items, allowed_weight, 0, 0, 0, sum_values(items, true), 0)
}
