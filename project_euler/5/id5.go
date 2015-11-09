package main

import (
    "fmt"
    "time"
    "math"
    "strconv"
)

/*
https://projecteuler.net/problem=5
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func main() {

    start := time.Now()
    last:=20
    number:=last-1
    
    for {
        number++
        nope:=false
        for k := last;k>1;k-- {
            if math.Mod(float64(number),float64(k)) != 0 {
                nope=true
                break
            }
        }
        if nope == false {
            break
        }
    }    
    fmt.Println("Lastprime: ", strconv.Itoa(number))
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 232.792.560
*/
