package main

import "fmt"

func main() {
//create array
    var a [5]int
    fmt.Println("emp:", a)

//create slice
    s := make([]string, 3)
    fmt.Println("empL", s)
    
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

// copy array "c" into s
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    l := c[2:5]
    fmt.Println("sl1:", l)

    l = s[2:]
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    var r = make([]string, 3)
    r[0], r[1], r[2] = "g", "h", "i"
    fmt.Println("dcl:", r)
    
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}

