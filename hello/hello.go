package main

import (
    "fmt"

    "github.com/Cotter45/first_GO_round/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
