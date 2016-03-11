package xolphin

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type DCVRequest struct {
	Domain        string `json:"domain"`
	DCVType       string `json:"dcvType"`
	ApproverEmail string `json:"approverEmail"`
}

type CertificateCreationRequest struct {
	Product                 int
	Years                   int
	CSR                     string
	DCVType                 string
	SubjectAlternativeNames []string
	DCV                     []DCVRequest
	Company                 string
	Department              string
	Address                 string
	ZIPCode                 string
	City                    string
	ApproverFirstName       string
	ApproverLastName        string
	ApproverEmail           string
	ApproverPhone           string
	KVK                     string
	Reference               string
}

func (self CertificateCreationRequest) ToURLValues() url.Values {
	values := url.Values{}
	values.Set("product", fmt.Sprintf("%d", self.Product))
	values.Set("years", fmt.Sprintf("%d", self.Years))
	values.Set("csr", self.CSR)
	values.Set("dcvType", self.DCVType)
	if len(self.SubjectAlternativeNames) > 0 {
		values.Set("subjectAlternativeNames", strings.Join(self.SubjectAlternativeNames, ","))
	}
	if len(self.DCV) > 0 {
		d, e := json.Marshal(self.DCV)
		if e == nil {
			values.Set("dcv", string(d))
		}
	}
	if self.Company != "" {
		values.Set("company", self.Company)
	}
	if self.Department != "" {
		values.Set("department", self.Department)
	}
	if self.Address != "" {
		values.Set("address", self.Address)
	}
	if self.ZIPCode != "" {
		values.Set("zipcode", self.ZIPCode)
	}
	if self.City != "" {
		values.Set("city", self.City)
	}
	if self.ApproverFirstName != "" {
		values.Set("approverFirstName", self.ApproverFirstName)
	}
	if self.ApproverLastName != "" {
		values.Set("approverLastName", self.ApproverLastName)
	}
	if self.ApproverEmail != "" {
		values.Set("approverEmail", self.ApproverEmail)
	}
	if self.ApproverPhone != "" {
		values.Set("approverPhone", self.ApproverPhone)
	}
	if self.KVK != "" {
		values.Set("kvk", self.KVK)
	}
	if self.Reference != "" {
		values.Set("reference", self.Reference)
	}
	return values
}

type CertificateReissueRequest struct {
	CSR                     string
	DCVType                 string
	SubjectAlternativeNames []string
	DCV                     []DCVRequest
	Company                 string
	Department              string
	Address                 string
	ZIPCode                 string
	City                    string
	ApproverFirstName       string
	ApproverLastName        string
	ApproverEmail           string
	ApproverPhone           string
	KVK                     string
	Reference               string
}

func (self CertificateReissueRequest) ToURLValues() url.Values {
	values := url.Values{}
	values.Set("csr", self.CSR)
	values.Set("dcvType", self.DCVType)
	if len(self.SubjectAlternativeNames) > 0 {
		values.Set("subjectAlternativeNames", strings.Join(self.SubjectAlternativeNames, ","))
	}
	if len(self.DCV) > 0 {
		d, e := json.Marshal(self.DCV)
		if e == nil {
			values.Set("dcv", string(d))
		}
	}
	if self.Company != "" {
		values.Set("company", self.Company)
	}
	if self.Department != "" {
		values.Set("department", self.Department)
	}
	if self.Address != "" {
		values.Set("address", self.Address)
	}
	if self.ZIPCode != "" {
		values.Set("zipcode", self.ZIPCode)
	}
	if self.City != "" {
		values.Set("city", self.City)
	}
	if self.ApproverFirstName != "" {
		values.Set("approverFirstName", self.ApproverFirstName)
	}
	if self.ApproverLastName != "" {
		values.Set("approverLastName", self.ApproverLastName)
	}
	if self.ApproverEmail != "" {
		values.Set("approverEmail", self.ApproverEmail)
	}
	if self.ApproverPhone != "" {
		values.Set("approverPhone", self.ApproverPhone)
	}
	if self.KVK != "" {
		values.Set("kvk", self.KVK)
	}
	if self.Reference != "" {
		values.Set("reference", self.Reference)
	}
	return values
}


type CertificateRenewRequest struct {
	Product                 int
	Years                   int
	CSR                     string
	DCVType                 string
	SubjectAlternativeNames []string
	DCV                     []DCVRequest
	Company                 string
	Department              string
	Address                 string
	ZIPCode                 string
	City                    string
	ApproverFirstName       string
	ApproverLastName        string
	ApproverEmail           string
	ApproverPhone           string
	KVK                     string
	Reference               string
}

func (self CertificateRenewRequest) ToURLValues() url.Values {
	values := url.Values{}
	values.Set("product", fmt.Sprintf("%d", self.Product))
	values.Set("years", fmt.Sprintf("%d", self.Years))
	values.Set("csr", self.CSR)
	values.Set("dcvType", self.DCVType)
	if len(self.SubjectAlternativeNames) > 0 {
		values.Set("subjectAlternativeNames", strings.Join(self.SubjectAlternativeNames, ","))
	}
	if len(self.DCV) > 0 {
		d, e := json.Marshal(self.DCV)
		if e == nil {
			values.Set("dcv", string(d))
		}
	}
	if self.Company != "" {
		values.Set("company", self.Company)
	}
	if self.Department != "" {
		values.Set("department", self.Department)
	}
	if self.Address != "" {
		values.Set("address", self.Address)
	}
	if self.ZIPCode != "" {
		values.Set("zipcode", self.ZIPCode)
	}
	if self.City != "" {
		values.Set("city", self.City)
	}
	if self.ApproverFirstName != "" {
		values.Set("approverFirstName", self.ApproverFirstName)
	}
	if self.ApproverLastName != "" {
		values.Set("approverLastName", self.ApproverLastName)
	}
	if self.ApproverEmail != "" {
		values.Set("approverEmail", self.ApproverEmail)
	}
	if self.ApproverPhone != "" {
		values.Set("approverPhone", self.ApproverPhone)
	}
	if self.KVK != "" {
		values.Set("kvk", self.KVK)
	}
	if self.Reference != "" {
		values.Set("reference", self.Reference)
	}
	return values
}