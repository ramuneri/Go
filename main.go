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
    fmt.Println("\n\n")

    mybill.addItem("soup", 2.50)
    mybill.addItem("chicken", 5.50)
    mybill.addItem("salad", 1.00)

    mybill.updateTip(2)

    fmt.Println(mybill.format())


}