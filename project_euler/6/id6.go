package main

import (
    "fmt"
    "time"
)

/*
https://projecteuler.net/problem=6
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

func main() {
    start := time.Now()

    var end int = 100
    var number, s_square, sum_square, i int 
    
    for ; i<= end; i++ {
        s_square+= i*i
    }

    sum:=(end*end)/2    
    sum_square=sum*sum
    number=sum_square-s_square
    
    fmt.Println("Zahl: ", number)
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 25164150
*/
