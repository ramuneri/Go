package main

import ("fmt"
        "strings"
        "sort"
    )




func main() {

    // 6-----------------------
    greeting := "Hello dino, this is me!"

    fmt.Println(strings.Contains(greeting, "Hello")) // true
    fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi")) // but does not change the original
    fmt.Println(strings.Split(greeting, " ")) // gives a slice [Hello dino, this is me!]
    fmt.Println("\n")


    ages := []int{20, 17, 23, 15, 14}
    sort.Ints(ages)
    fmt.Println(ages)

    index := sort.SearchInts(ages, 20) // must sort before using this
    fmt.Println(index) // will search in sorted list

    names := []string{"Ramune", "Ruta", "Marijus", "Zopa"}
    sort.Strings(names)
    fmt.Println(names)
    fmt.Println(sort.SearchStrings(names, "Tomas")) // 3 bc Tomas would fit on on 3












//     // 5----------------------
//     // var ages [3]int = [3]int{20, 17, 15}
//     var ages = [3]int{20, 17, 15}

//     names := [4]string{"Ramune", "Ruta", "Marijus", "Zopa"}

//     fmt.Println(ages, len(ages))
//     fmt.Println(names, len(names))


//     // slices - no fixed size (slices >> arrays)
//     var scores = []int{100, 94, 96}
//     fmt.Println(scores)

//     scores[1] = 99
//     scores = append(scores, 80) // adding new val
//     fmt.Println(scores)

//     // slices ranges
//     rangeOne := names[0:3]
//     fmt.Println(rangeOne)

//     rangeTwo := names[2:] // from two up to the end (works and other way)
//     fmt.Println(rangeTwo)



    // 4-------------------
    // name := "Ramune"
    // age := 20

    // // adds spaces itself :0
    // fmt.Println("My name is", name, "and my age is", age)

    // // with formated strings
    // fmt.Printf("My name is %v and my age is %v \n", name, age)
    // fmt.Printf("My name is %q and my age is %v \n", name, age) // q adds quotes
    // fmt.Printf("name is of type %T \n", name) // gives type
    // fmt.Printf("You scored %0.2f points\n", 98.79799) // directly

    // var info = fmt.Sprintf("My name is %v and my age is %v \n", name, age) // with Sprintf
    // fmt.Println("\nSaved info string is: ", info)

 
    // 3--------------------
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