package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "os"
)

var pool = "abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ:|?$%@][{}#&/()*"

func randomString(l int) string {

    bytes := make([]byte, l)

    for i := 0; i < l; i++ {
        bytes[i] = pool[rand.Intn(len(pool))]
    }

    return string(bytes)
}

func main() {
    for _, arg := range os.Args {
        switch arg {
            case "-h", "-help","--help":
                fmt.Println("Ok.. I can help you..")
                fmt.Println("Usage: ./gen {len of pass}")
                fmt.Println("Use -h -help or --help to see help again")
                fmt.Println("See you later)")
        }
        
    }

    if len(os.Args) == 1 {
		fmt.Println("Ok.. I can help you..")
        fmt.Println("Usage: ./gen {len of pass}")
        fmt.Println("Use -h -help or --help to see help again")
        fmt.Println("See you later)")
        return
    }

	heightOfPass, _ := strconv.Atoi(os.Args[1])
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomString(heightOfPass))
}
