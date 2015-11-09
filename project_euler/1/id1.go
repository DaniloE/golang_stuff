package main

import (
    "fmt"
    "time"
    "math"
    "strconv"
)

/*
https://projecteuler.net/problem=1
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000. 
*/

func main() {

    sum:=0
    end:=1000
    start := time.Now()

    for int := 1; int < end; int++ { 
        if math.Mod(float64(int),3) == 0 {
            sum += int
        } else if math.Mod(float64(int),5) == 0 {
            sum += int      
        }
    }
    fmt.Println("Sum: ", strconv.Itoa(sum))
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 233168
*/
