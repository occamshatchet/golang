package main

import "fmt"

func main() {
    //most basic type with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
    //classic loop 
    for j := 7; j <=9; j++ {
        fmt.Println(j)
    }
    // for loop condition will loop untill return.
    for {
        fmt.Println("loop")
        break
    }
}
