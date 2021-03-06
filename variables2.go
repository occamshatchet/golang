package main
import "fmt"
func main() {

    // var declars 1 or more variables.
    var a string = "initial"
    fmt.Println(a)

    // You can declare multiple variables at once.
    var c, b int = 1, 2
    fmt.Println(b,c)

    // go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)

    // variables without initialization are zero values. 
    var e int
    fmt.Println(e)

    // := is shorthand for declaring and initalizing a variable.
    f := "short"
    fmt.Println(f)
}

