package main
import (
    "fmt"
    "math/rand"
    "time"
    "strconv")


func gcd(a, b int) int {
    if b == 0 {
        if a < 0 {
            return -a
        } else {
            return a
        }
    } else {
        return gcd(b, a % b)
    }
}


func lcm(a, b int) int {
    res := (a / gcd(a, b)) * b
    if res < 0 {
        return -res
    } else {
        return res
    }
}


// Calculate the totient function λ(n)
// where n = p * q and p and q are prime.
func totient(p, q int) int {
    return lcm(p - 1, q - 1)
}


// Return a pseudo random number in the range [min, max).
func rand_range(min int, max int) int {
    return min + rand.Intn(max - min)
}


// Pick a random exponent e in the range [3, λ_n)
// such that gcd(e, λ_n) = 1.
func random_exponent(λ_n int) int {
    for {
        e := rand_range(3, λ_n)
        if gcd(e, λ_n) == 1 {
            return e
        }
    }
}


func inverse_mod(a, n int) int{
    t, newt := 0, 1
    r, newr := n, a
    for newr != 0 {
        quotient := r / newr
        t, newt = newt, t - (quotient * newt)
        r, newr = newr, r - (quotient * newr)
    }
    if r > 1 {
        panic("a is not invertible")
    }
    if t < 0 {
        t = t + n
    }
    return t
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


func is_probably_prime(p, num_tests int) bool {
    for i := 0; i < num_tests; i++ {
        n := rand_range(2, p)
        np := fast_exp_mod(n, p - 1, p)
        if np != 1 { return false }
    }
    return true
}


//finds a probable prime in the range [min, max)
func find_prime(min, max, num_tests int) int {
    for {
        p := rand_range(min, max)
        p |= 1 //increments p by 1 if p is even
        if is_probably_prime(p, num_tests) {
            return p
        }
    }
}


func main() {
    rand.Seed(time.Now().UnixNano())
    p, q := find_prime(10000, 50000, 20), find_prime(10000, 50000, 20)
    for p == q {
        q = find_prime(10000, 50000, 20)
    }
    n := p * q
    λ_n := totient(p, q)
    e := random_exponent(λ_n)
    d := inverse_mod(e, λ_n)

    fmt.Println("*** Public ***")
    fmt.Println("Public key modulus (n):", n)
    fmt.Println("Public key exponent (e):", e)
    fmt.Println()
    fmt.Println("*** Private ***")
    fmt.Println("Primes (p, q):", p, q)
    fmt.Println("λ(n):", λ_n)
    fmt.Println("d:", d)
    fmt.Println()

    for {
        var message string
        fmt.Printf("Enter a number in range [0, %d) which will be the message:", n)
        fmt.Scanln(&message)
        if len(message) == 0 { break }
        m, _ := strconv.Atoi(message)
        //if m < 0 || m >= n { break }
        ciphertext := fast_exp_mod(m, e, n)
        plaintext := fast_exp_mod(ciphertext, d, n)
        fmt.Println("Ciphertext:", ciphertext)
        fmt.Println("Plaintext:", plaintext)
    }
}
