package main

import (
    "xolphin-api-go"
    "fmt"
)

func main() {
    client := xolphin.NewClient("<username>", "<password>")
    request, err := client.Request().Get(964178701)
    if err != nil {
        fmt.Println("ERROR")
        fmt.Println(err)
    }
    fmt.Println(request.Id)
    fmt.Println(request.Embedded.Product.Brand)
}