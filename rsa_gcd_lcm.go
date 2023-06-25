package main
import (
    "fmt"
    "strconv")


var debug = false
var debug_calls_gcd = 0


func gcd(a, b int) int {
    if debug {
        debug_calls_gcd += 1
        fmt.Printf("gcd(%d, %d)\n", a, b)
    }
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


func main() {
    /*
    debug = true
    fmt.Println(gcd(12, 18))
    fmt.Println()
    fmt.Println(gcd(18, 12))
    fmt.Println()
    fmt.Println(gcd(-12, 18))
    fmt.Println()
    fmt.Println(gcd(-18, 12))
    fmt.Println()
    fmt.Println(gcd(12, -18))
    fmt.Println()
    fmt.Println(gcd(18, -12))
    fmt.Println()
    fmt.Println(gcd(-12, -18))
    fmt.Println()
    fmt.Println(gcd(-18, -12))
    fmt.Println()
    */

    //assumes inputs are empty or valid integers
    for {
        var astr, bstr string
        fmt.Printf("A:")
        fmt.Scanln(&astr)
        fmt.Printf("B:")
        fmt.Scanln(&bstr)
        if len(astr) == 0 || len(bstr) == 0 { break }
        a, _ := strconv.Atoi(astr)
        b, _ := strconv.Atoi(bstr)
        fmt.Printf("GCD(%d, %d) = %d\n", a, b, gcd(a,b))
        fmt.Printf("LCM(%d, %d) = %d\n", a, b, lcm(a,b))
    }
}
