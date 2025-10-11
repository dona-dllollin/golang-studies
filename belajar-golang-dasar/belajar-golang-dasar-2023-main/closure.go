// package main

// import "fmt"

// func main() {
// 	counter := 0

// 	increment := func() {
// 		fmt.Println("Increment")
// 		counter++
// 	}

// 	increment()
// 	increment()
// 	increment()

// 	fmt.Println(counter)
// }


package main

import "fmt"

func main() {
    // fungsi yang menghasilkan closure
    counter := func() func() int {
        x := 0
        return func() int {
            x++     // x "diingat" oleh closure ini
            return x
        }
    }()

    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}


