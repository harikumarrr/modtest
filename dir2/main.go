package main

import (
	"github.com/google/uuid"
	"fmt"
)

var (
	Number string
)

func main() {
	fmt.Println("Dir-2", uuid.New().String())
	Number = uuid.New().String()
}