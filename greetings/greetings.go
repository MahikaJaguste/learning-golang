package greetings

import (
    "fmt"
    "errors"
    "math/rand"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("Name cannot be empty")
    }
    greetingFormat := randomFormat()
    message := fmt.Sprintf(greetingFormat, name)
    return message, nil
}

func randomFormat() string {
    greetingsList := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    } 
    return greetingsList[rand.Intn(len(greetingsList))]
}

func Hellos(names []string) (map[string]string, error) {
    greetingsMp := make(map[string]string)
    for _, v := range names {
        greetingStr, err := Hello(v)
        if err != nil {
            return nil, err 
        }
        greetingsMp[v] = greetingStr
    }
    return greetingsMp, nil
}