package main
import "fmt"

func main() {
    //create an integer array that will hold five values.
    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("len:", len(a))

    //declare and initialize an array in one line
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    //compse two arrays to create a multi-demensional data structure
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
//        fmt.Println(i)
        for j := 0; j <3; j++ {
            fmt.Println("ielement", i)
            fmt.Println("jelement", j)
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
