package main
import (
    "fmt"
    "strconv"
    "math"
    "time")


//assumes num > 1 as
//0 has infinite factors
//this returns proper prime factors only
func find_factors(num int) []int {
    var factors []int
    for num % 2 == 0 {
        factors = append(factors, 2)
        num /= 2
    }
    for factor := 3; factor * factor <= num; factor += 2 {
        for num % factor == 0 {
            factors = append(factors, factor)
            num /= factor
        }
    }
    if num > 1 {
        factors = append(factors, num)
    }
    return factors
}


func multiply_slice(nums []int) int {
    product := 1
    for _, num := range nums {
        product *= num
    }
    return product
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


var prime_cache []int


//assumes num > 1 as
//0 has infinite factors
//this returns proper prime factors only
//also assumes num <= square root of max prime in prime_cache
func find_factors_sieve(num int) []int {
    var factors []int
    for _, factor := range prime_cache {
        for num % factor == 0 {
            factors = append(factors, factor)
            num /= factor
        }
        if num == 1 {
            break
        }
    }
    if num > 1 {
        factors = append(factors, num)
    }
    return factors
}


func main() {
    prime_cache = sieve_to_primes(sieve_of_euler(2000000000))

    //assumes inputs are empty or valid integers
    for {
        var numstr string
        fmt.Printf("Number to factorize:")
        fmt.Scanln(&numstr)
        if len(numstr) == 0 { break }
        num, _ := strconv.Atoi(numstr)
        if num < 2 { break }

        /*
        factors := find_factors(num)
        fmt.Println(factors)
        fmt.Println("Is factorization by find_factors correct:", multiply_slice(factors) == num)

        factors = find_factors_sieve(num)
        fmt.Println(factors)
        fmt.Println("Is factorization by find_factors_sieve correct:", multiply_slice(factors) == num)
        */

        // Find the factors the slow way.
        start := time.Now()
        factors := find_factors(num)
        elapsed := time.Since(start)
        fmt.Printf("find_factors:       %f seconds\n", elapsed.Seconds())
        //fmt.Println(multiply_slice(factors))
        fmt.Println(factors)
        fmt.Println()

        // Use the Euler's sieve to find the factors.
        start = time.Now()
        factors = find_factors_sieve(num)
        elapsed = time.Since(start)
        fmt.Printf("find_factors_sieve: %f seconds\n", elapsed.Seconds())
        //fmt.Println(multiply_slice(factors))
        fmt.Println(factors)
        fmt.Println()
    }
}
