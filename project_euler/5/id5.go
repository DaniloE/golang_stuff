package main

import (
    "fmt"
    "time"
)

/*
https://projecteuler.net/problem=5
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/


func lcm(a, b int64) int64 {
    return a * b / gcd(a, b)
}

func gcd(a, b int64) int64 {
    c := a % b
    if c == 0 {
        return b;       
    }
    return gcd(b, c)
}

func main() {
    start := time.Now()
    
    var number, i int64 = 1, 2
    var ende int64  = 20

    for ; i <= ende; i++ {
        number = lcm(number, i)
    }
    fmt.Println("Zahl: ", number)
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 232.792.560
*/
