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
	heightOfPass, _ := strconv.Atoi(os.Args[1])
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomString(heightOfPass))
}
