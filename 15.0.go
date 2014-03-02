package main
import "fmt"

func fib(ben int) {
    x := make([]int, ben) // We create an array "x" to hold integers up to the value given in the function call.
    x[0], x[1] = 1, 1 // Assign value 1 to x[0], 1 to x[1]
    for n := 2; n < ben; n++ {
        x[n] = x[n-1] + x[n-2]
    }
    fmt.Println(x)
}



func main() {
    fib(10)
}
