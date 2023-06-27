package main
import (
    "fmt"
    "math"
    "time"
    "runtime/debug")


// Build a sieve of Eratosthenes.
func sieve_of_eratosthenes(max int) []bool {
    is_prime := make([]bool, max + 1)
    if max >= 2 {
        is_prime[2] = true
        for i := 3; i <= max; i += 2 { is_prime[i] = true }
        for i := 3; i <= int(math.Ceil(math.Sqrt(float64(max)))); i += 2 {
            if is_prime[i] {
                for j := i * i; j <= max; j += i {
                    is_prime[j] = false
                }
            }
        }
    }
    return is_prime
}


func sieve_of_euler(max int) []bool {
    is_prime := make([]bool, max + 1)
    if max >= 2 {
        is_prime[2] = true
        for i := 3; i <= max; i += 2 { is_prime[i] = true }
        for i := 3; i <= int(math.Ceil(math.Sqrt(float64(max)))); i += 2 {
            if is_prime[i] {
                maxj := max / i
                if maxj & 1 == 0 { maxj -= 1 }
                for j := maxj; j >= i; j -= 2 {
                    if is_prime[j] {
                        is_prime[i * j] = false
                    }
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


func benchmark(gen_sieve func(int) []bool, max int) ([]bool, float64) {
    start := time.Now()
    sieve := gen_sieve(max)
    return sieve, time.Since(start).Seconds()
}


func main() {
    //check your system memory before setting pow2
    //pow2 = 32 will allocate 2 ^ 32 bytes (approx. 4 GiB)
    //pow2 = 33 will allocate 2 ^ 33 bytes (approx. 8 GiB) and so on
    max, sleep_duration, pow2 := 1, 0 * time.Second, 30

    fmt.Printf("%12s   %12s   %12s   %12s   %12s\n", "Max", "Eratosthenes", "Euler", "Ratio", "Comparison")

    for i := 0; i <= pow2; i++ {
        debug.FreeOSMemory()
        time.Sleep(sleep_duration)
        eratosthenes_sieve, eratosthenes_time := benchmark(sieve_of_eratosthenes, max)

        debug.FreeOSMemory()
        time.Sleep(sleep_duration)
        euler_sieve, euler_time := benchmark(sieve_of_euler, max)

        comparison := "ERROR"
        if len(eratosthenes_sieve) == len(euler_sieve) {
            j := 0
            for ; j < len(eratosthenes_sieve); j++ {
                if eratosthenes_sieve[j] != euler_sieve[j] {
                    break
                }
            }
            if j == len(eratosthenes_sieve) {
                comparison = "OK"
            }
        }

        fmt.Printf("%12d   %12.6f   %12.6f   %12.6f   %12s\n", max, eratosthenes_time, euler_time, eratosthenes_time / euler_time, comparison)

        max <<= 1
    }
}
