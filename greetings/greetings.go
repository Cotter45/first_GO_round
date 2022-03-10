package greetings

import (
    // "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, string) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "blank", "empty name"
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, "nil"
}