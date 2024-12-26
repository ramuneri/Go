package main

import (
    "fmt"
)

func main() {
    menu := map[string]float64 {
        "soup": 1.99,
        "pie": 3.99,
        "salad": 4.99,
        "pizza": 7.99,
    }

    fmt.Println(menu)
    fmt.Println(menu["pie"])

    for k, v := range menu {
        fmt.Println(k, "-", v)
    }


    phoneBook := map[int]string {
        1234: "Ramune",
        5678: "Ruta",
        9101: "Marijus",
    }

    fmt.Println(phoneBook)
    fmt.Println(phoneBook[5678])

    phoneBook[1234] = "Ramunele"
    fmt.Println(phoneBook)


}