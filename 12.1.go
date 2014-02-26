package main
import "fmt"

func main() {
    fmt.Println(order(40, 24))
}

func order(a, b int) (int, int) {
    if a > b {
        return b,a
    }
    return a,b
}
