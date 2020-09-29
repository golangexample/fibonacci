package main

import ("fmt"; "strconv")

func loop(x int) int {
    y := make([]int, x + 1, x + 2)
    if x < 2 {
        y = y[0:2]
    }
    y[0] = 0
    y[1] = 1
    for i := 2; i <= x; i++ {
        y[i] = y[i - 1] + y[i - 2]
    }
    return y[x]
}

func main() {

    fmt.Println("The first 20 in a Fibonacci sequence:")

    for i := 0; i <= 19; i++ {
        fmt.Print(strconv.Itoa(loop(i)) + "\n")
    }
}
