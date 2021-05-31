package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("My favorite number is", rand.Intn(10))
}