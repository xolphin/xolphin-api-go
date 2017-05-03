package main

import (
	"xolphin"
	"fmt"
)

func main() {
	client := xolphin.NewTestClient("nikita@xolphin.nl", "test1234")

	ccr := client.Request().CreateEE()
	ccr.CSR = `-----BEGIN CERTIFICATE REQUEST-----
MIICuDCCAaACADB0MQswCQYDVQQGEwJOTDEWMBQGA1UEAxMNd3d3LnNzbGRldi5u
bDEWMBQGA1UEBxMNSGVlcmh1Z293YWFyZDEQMA4GA1UEChMHWG9scGhpbjEWMBQG
A1UECBMNTm9vcmQtSG9sbGFuZDELMAkGA1UECxMCSVQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQDkTFMNKhWObiUH1xDeRcqYiROKvJZEI7njDlQbXeVC
qSaBSBClJD98OSgE0+arnVyJ5Fnjkovgb9aC8HbWYbiD1tbWD+hKC1wPODEvknl7
jfRpO36cIglIhftHn5p0k+f86uS2Nz1yr+5qwAEQ7C0UQZxMuBbzK6Uv8m0TbAVV
CQ+i+uJfuXSsAT8LZbyoJsu50ZySG3uEhCSISh1cS/t/M2INHbXfzGA3GWH2FXCZ
Qd9eLSBAVEanIr3TlRpNU3f0IDwJpm56BVTNtNL7GI2NilfUF9oSo9PSlPbAkPsN
BrVdOiZgiOEO74rXE8fjA5Zm8uic6MyFRo4FuTUBWubrAgMBAAGgADANBgkqhkiG
9w0BAQsFAAOCAQEAHuB6643oI7oPw5SbFQAcfTfQbyRTXNqwYOO8YRcCkDIu7aX9
pupGvf+cSa5IDcZ6Dz22/Khp5rYlcutAtEN7d8MNBwBlcVzbWL+DuTU5ZWAxcRi4
nmiw/C+RxFF+WsUcjwb+dNrTYkz03t7voMg/0NS7RRU/oTu8heDDcH+Ffam1bZJQ
zi21dz/AsjG9Jc6GJqZs0ImEwCQlBhJYnXPj3FB8U/mzSWSfq502fdtmQrvIgHmI
RLuxVO9QX0YLQ0ew1x2z+eFfBsfSI1+DkF4+5TrfJP6jQyHBR+VDdGYBO2d8rBw7
ITIJC9tt2F4GjaPAI1xY9eoq/QfZxzNoNufMuw==
-----END CERTIFICATE REQUEST-----`
	ccr.ApproverFirstName = "FirstName"
	ccr.ApproverLastName = "LastName"
	ccr.ApproverPhone = "+1234567890"
	ccr.ApproverEmail = "email@domain.com"
	ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test1.domain.com")
	ccr.SubjectAlternativeNames = append(ccr.SubjectAlternativeNames, "test2.domain.com")
	ccr.DCVType = "FILE"

	result, err := client.Request().SendEE(ccr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result.Certificate)

}
