# Xolphin API wrapper for Go

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

#### Getting list of requests

```go
requests, err := client.Request().All()
if err != nil {
    fmt.Println(err)
}
for _, request := range requests {
    fmt.Println(request.Id)
}
```

### Getting request by ID

```go
request, err := client.Request().Get(1234)
if err != nil {
    fmt.Println(err)
}
fmt.Println(request.Embedded.Product.Brand)
```

### Request certificate

```go
ccr := client.Request().Create(24, 1, `<csr_string>`, "EMAIL")
ccr.Address = "Address"
ccr.ApproverFirstName = "FirstName"
ccr.ApproverLastName = "LastName"
ccr.ApproverPhone = "+1234567890"
ccr.ZIPCode = "123456"
ccr.City = "City"
ccr.Company = "Company"
ccr.ApproverEmail = "email@domain.com"
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
