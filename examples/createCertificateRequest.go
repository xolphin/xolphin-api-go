package main

import (
    "xolphin-api-go"
    "fmt"
)

func main() {
    client := xolphin.NewClient("<username>", "<password>")
    csr := `-----BEGIN CERTIFICATE REQUEST-----
MIICszCCAZsCAQAwbjELMAkGA1UEBhMCR0IxDzANBgNVBAgMBkxvbmRvbjEPMA0G
A1UEBwwGTG9uZG9uMRMwEQYDVQQKDApTU0wyNDcgTHRkMQswCQYDVQQLDAJJVDEb
MBkGA1UEAwwSc3NsMjQ3LXRlc3RpbmcuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAwkpLdjQJohfnqBoNXCMTnsOaXfyz7a4nm3BNZhGF9I4fPdd9
fuOZhOqIdy2tmQUp82rm9OjDtj9NjJHD3YRGLwTUJ2CJfCIaKescGAZYe8iNNyJ8
Fuw/Jc2TD6UKgb/HyD9grqMQEPojD469ZA8ZYjAPax8UsilTdU5a8w0F+U6uF7Q+
IMlo/OZ5pqewQTZ/fqPm/MoGiUYnMEnSzE9aVGUSjMov8nd/s/pEbM3Je7zg4aHX
gFemBN5GY8vFQanZhSvGcIhhjdK3ac7Q5j5Y20sYmaORulvz0YH40ZYoml8KP8GO
PKOlLhpKnade4u+kxo72Ws05cRtm4N78B5JLsQIDAQABoAAwDQYJKoZIhvcNAQEL
BQADggEBADNQWLy+z8HdN3Xz0HJAPK6c5G7ukCAu/P8sYwNkWFnGGBrY7Uvl3XXF
vSBMWS04AzylTn2qqHKuz5dIq68lgi0CTl1tR2Sp8/5vwCFZntJHGG6f/L/DRKcL
Iriu7QW4mi5dQpDX6maJu/lE38XrcX/r2QdCkCWskiOdh+EnpwibdfT7sVqSrvDh
Cgd0E7WtG1ecGB1QTpCsHBHMk604Pj/wDnw5SvNoX01dm8n3PY5/UfMFVrCnx2uT
VDiArcp4X9/aZCiJxrf6UzY/7FH6qXYiP+hlIVqrqZ7Y3qaFQznBO3VWWNvuunXd
1V2zpjfQjv+FuBDHyKUgX0Q4h3M9XTw=
-----END CERTIFICATE REQUEST-----`

    ccr := client.Request().Create(115, 1, csr, "EMAIL")
    ccr.Address = "Address"
    ccr.ApproverRepresentativeFirstName = "FirstName"
    ccr.ApproverRepresentativeLastName = "LastName"
    ccr.ApproverRepresentativePhone = "+1234567890"
    ccr.ApproverRepresentativeEmail = "first.last@company.com"
    ccr.ApproverRepresentativePosition = "IT"
    ccr.ZIPCode = "123456"
    ccr.City = "City"
    ccr.Company = "Company"
    ccr.ApproverEmail = "admin@ssl247-testing.com"
    ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test1.ssl247-testing.com")
    ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test2.ssl247-testing.com")
    ccr.DCV = append(ccr.DCV, xolphin.DCVRequest{Domain: "test1.ssl247-testing.com", DCVType: "FILE"})
    ccr.DCV = append(ccr.DCV, xolphin.DCVRequest{Domain: "test2.ssl247-testing.com", DCVType: "EMAIL",  ApproverEmail: "admin@ssl247-testing.com"})

    result, err := client.Request().Send(ccr)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result.Id)
    
}