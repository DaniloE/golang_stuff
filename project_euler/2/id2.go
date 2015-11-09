package main

import (
    "fmt"
    "time"
    "math"
    "strconv"
)

/*
https://projecteuler.net/problem=2
Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/




func main() {

    sum:=0
    end:=4000000
    first:=1
    second:=2
    var fib int
    start := time.Now()

    // 2 is the first number in the row and even...we have to add it
    sum+=second
    for {
        fib=first+second
        if fib > end {
            break
        }
        if math.Mod(float64(fib),2) == 0 {
            sum += fib 
        }
        first=second
        second=fib
    }
    fmt.Println("Sum: ", strconv.Itoa(sum))
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 4613732
*/