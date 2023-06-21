// Fibonacci numbers
package main
import (
    "strconv"
    "fmt")


func fibonacci(n int64) int64 {
    if n < 2 {
        return n
    } else {
        return fibonacci(n - 1) + fibonacci(n - 2)
    }
}


func main() {
    for {
        // Get n as a string.
        var n_string string
        fmt.Printf("N: ")
        fmt.Scanln(&n_string)

        // If the n string is blank, break out of the loop.
        if len(n_string) == 0 { break }

        // Convert to int and calculate the Fibonacci number.
        n, _ := strconv.ParseInt(n_string, 10, 64)
        fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
    }
}
