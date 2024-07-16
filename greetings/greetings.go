package greetings

import (
    "fmt"
    "errors"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("Name cannot be empty")
    }
    message := fmt.Sprintf("Hello, %v", name)
    return message, nil
}