package main

import (
    "fmt"
    "time"
    "math"
    "math/big"
    "strconv"
)

/*
https://projecteuler.net/problem=3
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func isprime(x int) bool {
    if math.Mod(float64(x),2) == 0 {
        return false
    } else if math.Mod(float64(x),3) == 0 {
        return false
    }
    return true
}


func main() {

    start := time.Now()
    var lastprime int64
    var i int64 
    var number int64 =600851475143
    
    for number > 1{
        i++
        testint:=big.NewInt(i)
        if testint.ProbablyPrime(1) == true {
            if math.Mod(float64(number),float64(i)) == 0 {
                number = number/i
            }
        }
        lastprime=i
    }    
    fmt.Println("Lastprime: ", strconv.FormatInt(lastprime,10))
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 6857
*/
