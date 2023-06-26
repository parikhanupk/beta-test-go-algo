package main
import (
    "fmt"
    "math"
    "time")


// Build a sieve of Eratosthenes.
func sieve_of_eratosthenes(max int) []bool {
    is_prime := make([]bool, max + 1)
    if max >= 2 {
        is_prime[2] = true
        for i := 3; i <= max; i++ { is_prime[i] = true }
        for i := 3; i <= int(math.Ceil(math.Sqrt(float64(max)))); i += 2 {
            if is_prime[i] == true {
                for j := i * i; j <= max; j += i {
                    is_prime[j] = false
                }
            }
        }
    }
    return is_prime
}


func print_sieve(is_prime []bool) {
    if len(is_prime) > 2 {
        fmt.Print("2")
        for i := 3; i < len(is_prime); i += 2 {
            if is_prime[i] {
                fmt.Printf(" %d", i)
            }
        }
        fmt.Println()
    }
}


func sieve_to_primes(is_prime []bool) []int {
    if len(is_prime) > 2 {
        primes := make([]int, 1)
        primes[0] = 2
        for i := 3; i < len(is_prime); i += 2 {
            if is_prime[i] {
                primes = append(primes, i)
            }
        }
        return primes
    }
    return nil
}


func main() {
    var max int
    fmt.Printf("Max: ")
    fmt.Scan(&max)

    start := time.Now()
    sieve := sieve_of_eratosthenes(max)
    elapsed := time.Since(start)
    fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

    if max <= 1000 {
        print_sieve(sieve)

        primes := sieve_to_primes(sieve)
        fmt.Println(primes)
    }
}
