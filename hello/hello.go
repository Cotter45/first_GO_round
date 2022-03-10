package main

import (
    "fmt"
    "log"

    "github.com/Cotter45/first_GO_round/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != "error" {
        // log.Fatal("error")
				fmt.Println("error")
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}
