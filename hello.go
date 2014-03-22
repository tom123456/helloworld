package main

import (
    "fmt"
    "time"
    "math/rand"
    "math"
)

//import "code.google.com/p/go-tour/pic"

func add(x int, y int) int {
    return x + y
}
/*
func Pic(dx, dy int) [][]uint8 {
}
*/

func main() {
    fmt.Println("Hello, world. 你好，世界！")
    fmt.Println(time.Now().Unix())
    rand.Seed(time.Now().Unix())
    fmt.Println("My favorite number is", rand.Intn(3))
    fmt.Printf("Now you have %g problems. \n",
                        math.Nextafter(2, 3))
    fmt.Println(add(42,13))
    //pic.Show(Pic)
}



