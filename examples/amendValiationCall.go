package main

import (
    "xolphin-api-go"
    "fmt"
)

func main() {
    client := xolphin.NewClient("<username>", "<password>")

    callDetails := map[string]string{
        "Timezone" : "Europe/Amsterdam",
        "Action" : "ScheduledCallback", // available actions: ScheduledCallback, ManualCallback, ReplacePhone, replaceEmailAddress, sendCallbackEmail
        "PhoneNumber" : "132465",
        "ExtensionNumber" : "12",
        "EmailAddress" : "email@client.com",
        "Language" : "en-us",
        "Date" : "2023-11-27", // have to be in future
        "Time" : "14:00",
        "Comments" : "",
    }
    request, err := client.Request().ScheduleValidationCall(123456,callDetails)
    if err != nil {
        fmt.Println("err")
        fmt.Println(err)
    }
    fmt.Println(request.Message)
}