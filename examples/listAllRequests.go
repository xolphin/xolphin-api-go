package main

import (
    "xolphin-api-go"
    "fmt"
)

func main() {
    client := xolphin.NewClient("<username>", "<password>")
    requests, err := client.Request().All()
    if err != nil {
        fmt.Println(err)
    }
    for _, request := range requests {
        fmt.Println(request.Id)
    }
    
}