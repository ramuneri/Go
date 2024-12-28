package main

import (
    "fmt"
)

func updateName(x *string) {
    *x = "Marijus"
}

func main() {
    name := "Ramune"

    m := &name
    fmt.Println("memory address is: ", m)
    fmt.Println("value there is: ", *m)

    fmt.Println(name)
    updateName(m)
    fmt.Println(name)


}