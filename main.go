package main

import ("fmt"
        "strings"
        // "sort"
        "math"
    )

    func sayGreeting(n string) {
        fmt.Printf("Good morning %v \n", n)
    }


    func sayBye(n string) {
        fmt.Printf("Bye  %v \n", n)
    }


    func cycleNames(n []string, f func(string)) {
        for _, v := range n {
            f(v)
        }
    }

    func circleArea(r float64) float64 {
        return math.Pi * r * r
    }


    func getInitials(n string) (string, string) {
        s := strings.ToUpper(n)
        names := strings.Split(s, " ")

        var initials []string
        for _, v := range names {
            initials = append(initials, v[:1])
        }

        if len(initials) > 1 {
            return initials[0], initials[1]
        }
        return initials[0], "_"
    }


func main() {
    // 9. --------------------------
    // sayGreeting("Ramune")
    // sayBye("Ruta")

    // cycleNames([]string{"ramunele", "rami", "rama"}, sayGreeting)

    // a1 := circleArea(10.5)
    // a2 := circleArea(15)

    // fmt.Println(a1, a2)
    // fmt.Printf("a1 = %0.3f, a2 = %0.3f", a1, a2)

    // fn, ln := getInitials("ramune Jan")
    // fmt.Println(fn, ln)

    // // 7--------------------------
    // x := 0
    // for x < 5 {
    //     fmt.Println("x = ", x)
    //     x++
    // }

    // names := []string{"Ramune", "Ruta", "Marijus", "Zopa"}
    // for i := 0; i < len(names); i++ {
    //     fmt.Println(names[i])
    // }

    // for index, value := range names {
    //     fmt.Printf("the positions at index %v is %v \n", index, value)
    // }

    // for _, value := range names { // if no need for index
    //     fmt.Printf("the name is %v \n", value)
    // }


    // // 6-----------------------
    // greeting := "Hello dino, this is me!"

    // fmt.Println(strings.Contains(greeting, "Hello")) // true
    // fmt.Println(strings.ReplaceAll(greeting, "Hello", "hi")) // but does not change the original
    // fmt.Println(strings.Split(greeting, " ")) // gives a slice [Hello dino, this is me!]
    // fmt.Println("\n")


    // ages := []int{20, 17, 23, 15, 14}
    // sort.Ints(ages)
    // fmt.Println(ages)

    // index := sort.SearchInts(ages, 20) // must sort before using this
    // fmt.Println(index) // will search in sorted list

    // names := []string{"Ramune", "Ruta", "Marijus", "Zopa"}
    // sort.Strings(names)
    // fmt.Println(names)
    // fmt.Println(sort.SearchStrings(names, "Tomas")) // 3 bc Tomas would fit on on 3


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