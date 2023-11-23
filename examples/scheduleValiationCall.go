package main

import (
    "xolphin-api-go"
    "fmt"
    "time"
)

func main() {
    client := xolphin.NewClient("<username>", "<password>")
    call_date_time := time.Now().Add(time.Hour * 72) // have to be in future and not on weekends
    request, err := client.Request().ScheduleValidationCall(123456,call_date_time)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(request.Message)
}