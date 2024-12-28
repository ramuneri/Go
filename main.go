package main

import (
    "fmt"
)

func updateName(x *string) {
    *x = "Marijus"
}

func main() {
    mybill := newBill("Marijaus bill")

    fmt.Println(mybill.format())


}