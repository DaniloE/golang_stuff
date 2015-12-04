package main

import (
    "fmt"
    "time"
    "strconv"
)

/*
https://projecteuler.net/problem=1
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000. 
*/

func get_multi(mod int,max int) int {
    
    var k,i,modi int
    for {
        i++
        modi=mod*i
        if modi < max {
            k += modi
        } else {
            break
        }
    }
    return k
}


func main() {

    var last int
    end:=1000
    start := time.Now()

    last= (get_multi(3,end) + get_multi(5,end)) - get_multi(15,end)

    fmt.Println("Sum: ", strconv.Itoa(last))
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 233168
*/
