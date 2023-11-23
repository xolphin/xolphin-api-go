package main

import (
    "xolphin-api-go"
    "fmt"
)

func main() {
    client := xolphin.NewClient("<username>", "<password>")
    products, err := client.Support().Products()
    if err != nil {
        fmt.Println(err)
    }
    for _, product := range products {
        fmt.Println(product.Id, product.Brand, product.Name)
    }
}