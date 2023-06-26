package main
import (
    "fmt"
    "strconv")


func fast_exp(num, pow int) int {
    result := 1
    for pow > 0 {
        if pow & 1 == 1 { result *= num }
        pow >>= 1
        num *= num
    }
    return result
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


func main() {
    //assumes inputs are empty or valid integers
    for {
        var num_str, pow_str, mod_str string
        fmt.Printf("num:")
        fmt.Scanln(&num_str)
        fmt.Printf("pow:")
        fmt.Scanln(&pow_str)
        fmt.Printf("mod:")
        fmt.Scanln(&mod_str)
        if len(num_str) == 0 || len(pow_str) == 0 || len(mod_str) == 0 { break }
        num, _ := strconv.Atoi(num_str)
        pow, _ := strconv.Atoi(pow_str)
        mod, _ := strconv.Atoi(mod_str)
        if num < 1 || pow < 1 || mod < 1 { break }
        fmt.Printf("%d ^ %d = %d\n", num, pow, fast_exp(num, pow))
        fmt.Printf("(%d ^ %d) mod %d = %d\n", num, pow, mod, fast_exp_mod(num, pow, mod))
    }
}
