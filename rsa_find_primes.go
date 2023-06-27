package main
import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "strconv")


const num_tests = 20


// Return a pseudo random number in the range [min, max).
func rand_range(min int, max int) int {
    return min + rand.Intn(max - min)
}


func fast_exp_mod(num, pow, mod int) int {
    result := 1
    for pow > 0 {
        if pow & 1 == 1 { result = (result * num) % mod }
        pow >>= 1
        num = (num * num) % mod
    }
    return result
}


func is_probably_prime(p int) bool {
    for i := 0; i < num_tests; i++ {
        n := rand_range(2, p)
        np := fast_exp_mod(n, p - 1, p)
        if np != 1 { return false }
    }
    return true
}


//assumes digits will be in the range [1, 15] (64 bit machines)
func find_prime(digits int) int {
    for {
        digits_max := int(math.Pow(float64(10), float64(digits)))
        digits_min := digits_max / 10
        if digits_min == 1 { digits_min = 2 }
        p := rand_range(digits_min, digits_max)
        p |= 1 //increments p by 1 if p is even
        if is_probably_prime(p) {
            return p
        }
    }
}


func test_known_values() {
    primes := []int {
        10009, 11113, 11699, 12809, 14149,
        15643, 17107, 17881, 19301, 19793,
    }
    composites := []int {
        10323, 11397, 12212, 13503, 14599,
        16113, 17547, 17549, 18893, 19999,
    }

    fmt.Printf("Probability: %9.6f%%\n\n", (1 - math.Pow(0.5, num_tests)) * 100)
    fmt.Printf("Primes:\n")
    for _, p := range primes {
        if is_probably_prime(p) {
            fmt.Printf("%d  Prime\n", p)
        } else {
            fmt.Printf("%d  Composite\n", p)
        }
    }
    fmt.Printf("\nComposites:\n")
    for _, p := range composites {
        if is_probably_prime(p) {
            fmt.Printf("%d  Prime\n", p)
        } else {
            fmt.Printf("%d  Composite\n", p)
        }
    }
}


func main() {
    rand.Seed(time.Now().UnixNano())
    test_known_values()
    fmt.Println()

    for {
        var digits_str string
        fmt.Printf("# Digits: ")
        fmt.Scanln(&digits_str)
        if len(digits_str) == 0 { break }
        digits, _ := strconv.Atoi(digits_str)
        if digits < 1 || digits > 15 { break }
        fmt.Println("Prime:", find_prime(digits))
    }
}
