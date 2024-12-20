package main

import "fmt"

func main() {
    name := "Ramune"
    age := 20

    // adds spaces itself :0
    fmt.Println("My name is", name, "and my age is", age)

    // with formated strings
    fmt.Printf("My name is %v and my age is %v \n", name, age)
    fmt.Printf("My name is %q and my age is %v \n", name, age) // q adds quotes
    fmt.Printf("name is of type %T \n", name) // gives type
    fmt.Printf("You scored %0.2f points\n", 98.79799) // directly

    var info = fmt.Sprintf("My name is %v and my age is %v \n", name, age) // with Sprintf
    fmt.Println("\nSaved info string is: ", info)











    // // String
    // var nameOne string = "Ramune"
    // var nameTwo = "Ruta"
    // var nameFour string // no val, so will give an empty string

    // fmt.Println(nameOne, nameTwo, nameFour)

    // // or shorter way:
    // nameThree := "Marijus"
    // fmt.Println(nameThree)


    // // Int
    // var ageOne int = 20
    // var ageTwo = 17
    // ageThree := 15
    // fmt.Println(ageOne, ageTwo, ageThree)

}