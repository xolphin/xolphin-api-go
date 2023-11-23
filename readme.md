# Xolphin API wrapper for Go
xolphin-go-api is a library which allows quick integration of the [Xolphin REST API](https://api.xolphin.com) in Go to automated ordering, issuance and installation of SSL Certificates.

## About Xolphin
[Xolphin](https://www.xolphin.nl/) is the largest supplier of [SSL Certificates](https://www.sslcertificaten.nl) and [Digital Signatures](https://www.digitalehandtekeningen.nl) in the Netherlands. Xolphin has  
a professional team providing reliable support and rapid issuance of SSL Certificates at an affordable price from industry leading brands such as Sectigo, GeoTrust, GlobalSign, Thawte and Symantec.
## Library installation

Library can be installed go package manager

```
go get github.com/xolphin/xolphin-api-go
```

And updated via

```
go get -u github.com/xolphin/xolphin-api-go
```

## Usage

### Client initialization

```go
package main

import (
	"github.com/xolphin/xolphin-api-go"
	"fmt"
)

func main() {
	client := xolphin.NewClient("<username>", "<password>")
	fmt.Println(client.BaseURI)
}
```

### Requests

#### Get list of requests

```go
requests, err := client.Request().All()
if err != nil {
    fmt.Println(err)
}
for _, request := range requests {
    fmt.Println(request.Id)
}
```

#### Get request by ID

```go
request, err := client.Request().Get(1234)
if err != nil {
    fmt.Println(err)
}
fmt.Println(request.Embedded.Product.Brand)
```

#### Request certificate

```go
ccr := client.Request().Create(24, 1, `<csr_string>`, "EMAIL")
ccr.Address = "Address"
ccr.ApproverRepresentativeFirstName = "FirstName"
ccr.ApproverRepresentativeLastName = "LastName"
ccr.ApproverRepresentativePhone = "+1234567890"
ccr.ApproverRepresentativeEmail = "email@domain.com"
ccr.ApproverRepresentativePosition = "IT"
ccr.ZIPCode = "123456"
ccr.City = "City"
ccr.Company = "Company"
ccr.ApproverEmail = "admin@domain.com"
ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test1.domain.com")
ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test2.domain.com")
ccr.DCV = append(ccr.DCV, xolphin.DCVRequest{Domain: "test1.domain.com", DCVType: "EMAIL",  ApproverEmail: "test1@domain.com"})
ccr.DCV = append(ccr.DCV, xolphin.DCVRequest{Domain: "test2.domain.com", DCVType: "EMAIL",  ApproverEmail: "test2@domain.com"})

result, err := client.Request().Send(ccr)
if err != nil {
    fmt.Println(err)
}
fmt.Println(result.Id)
```

#### Create a note

```go
result, err := client.Request().SendNote(1234, "My message")
	if err != nil {
		fmt.Println(err)
	}
fmt.Println(result)
```

#### Get list of notes

```go
result, err := client.Request().GetNotes(960002428)
if err != nil {
    fmt.Println(err)
}
for _,element := range result {
    fmt.Println(element)
}
```

#### Send a "Comodo Subscriber Agreement" email

```go
//currently available languages: en, de, fr, nl
result, err := client.Request().SendComodoSA(1234, "mail@example.com","en")
if err != nil {
    fmt.Println(err)
}
fmt.Println(result)
```

#### Request an "Encryption Everywhere" certificate
```go
ccr := client.Request().CreateEE()
ccr.CSR = "<csr_string>"
ccr.ApproverRepresentativeFirstName = "FirstName"
ccr.ApproverRepresentativeLastName = "LastName"
ccr.ApproverRepresentativePhone = "+1234567890"
ccr.ApproverRepresentativeEmail = "user@domain.com"
ccr.ApproverRepresentativePosition = "IT"
ccr.ApproverEmail = "admin@domain.com"
ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test1.domain.com")
ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test2.domain.com")
ccr.DCVType = "FILE"

result, err := client.Request().SendEE(ccr)
```

#### Schedule a validation callback
```go
now := time.Now()
result, err := client.Request().ScheduleValidationCall(1234,now)

```
Or
```go
callDetails := map[string]string{
    "Timezone" : "Europe/Amsterdam",
    "Date" : "2024-01-01",
    "Time" : "14:00",
    "Action" : "ScheduledCallback",
    "PhoneNumber" : "132465",
    "ExtensionNumber" : "12",
    "EmailAddress" : "email@client.com",
    "Language" : "en-us",
    "Comments" : "",
}
request, err := client.Request().ScheduleValidationCall(1234,callDetails)
```
where action can be:

"ScheduledCallback" Schedule automatic Callback for customer 
"ManualCallback" Notify support that customer is requesting manual callback 
"ReplacePhone" Notify support that current phone number is invalid and suggest a new one, or that the callback method needs to be changed from Email to Telephone. 
"replaceEmailAddress" Notify support that the callback method needs to be changed from Telephone to Email. 
"sendCallbackEmail" Request re-sending callback email to customer

### Certificate

#### Certificates list and expirations

```go
certificates, err := client.Certificate().All()
if err != nil {
    fmt.Println(err)
}
for _, cert := range certificates {
    fmt.Println(cert.Id)
}
```

#### Download certificate

```go
crt, err := client.Certificate().Download(1234, "CRT")
if err != nil {
    fmt.Println(err)
}
ioutil.WriteFile("crt.crt", crt, 0777)
```

### Support

#### Products list

```go
products, err := client.Support().Products()
if err != nil {
    fmt.Println(err)
}
for _, product := range products {
    fmt.Println(product.Id, product.Brand, product.Name)
}
```

#### Decode CSR

```go
dc, err := client.Support().DecodeCSR(csr)
if err != nil {
    fmt.Println(err)
}
fmt.Println(dc.Type, dc.Size)
```
