package main

import (
    "fmt"
    "time"
    "math/big"
)

/*
https://projecteuler.net/problem=7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10001st prime number?
*/

func main() {
    start := time.Now()
    var max int64 = 10001
    var k,i int64
    for k<max {
        i++
        newint := big.NewInt(i)
        if newint.ProbablyPrime(5) == true {
            k++
        }
    }
    

    fmt.Println("Zahl: ", i)
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 104743 
*/
