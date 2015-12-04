package main

import (
    "fmt"
    "time"
    "strconv"
)

/*
https://projecteuler.net/problem=4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func ispal(x int) bool {
    charint:=strconv.Itoa(x)
    charint_rune:=[]rune(charint)
    charint_reverse:=charint_rune
    for i, j := 0, len(charint_reverse)-1; i < len(charint_reverse)/2; i, j = i+1, j-1 {
        charint_reverse[i], charint_reverse[j] = charint_reverse[j], charint_reverse[i]
    }
    if charint == string(charint_reverse) {
        return true
    } else {
        return false
    }
}


func main() {

    start := time.Now()
    var biggest int

    for number1:= 100; number1 <= 999; number1++ {
        for number2:=100;number2 <= 999; number2++ {
            number3:=number1*number2
            if ispal(number3) == true {
                if number3 > biggest {
                    biggest=number3
                }
            }
        }

    }

    fmt.Println("Biggest: ", strconv.Itoa(biggest))
    fmt.Println("Laufzeit: ", time.Since(start))
}

/*
Solution: 906609
*/
