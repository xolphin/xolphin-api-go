package xolphin

type BaseResponse struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`

	Page  int `json:"page"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

func (self BaseResponse) isError() bool {
	return len(self.Message) > 0 || len(self.Errors) > 0
}

type RequestsResponse struct {
	BaseResponse
	Embedded struct {
		Requests []RequestResponse `json:"requests"`
	} `json:"_embedded"`
}

type RequestValidationResponse struct {
	Status        bool   `json:"status"`
	StatusDetail  int    `json:"statusDetail"`
	StatusMessage string `json:"statusMessage"`
	Domains       []struct {
		Domain        string      `json:"domain"`
		DCVType       string      `json:"dcvType"`
		DCVEmail      string      `json:"dcvEmail"`
		Status        interface{} `json:"status"`       // TODO: what is proper type for this - boolean or integer?
		StatusDetail  interface{} `json:"statusDetail"` // TODO: what is proper type for this - boolean or integer?
		StatusMessage string      `json:"statusMessage"`
		Md5           string      `json:"md5"`
		Sha1          string      `json:"sha1"`
	}
}

type ProductsResponse struct {
	BaseResponse
	Embedded struct {
		Products []ProductResponse `json:"products"`
	} `json:"_embedded"`
}

type ProductResponse struct {
	BaseResponse
	Id             int    `json:"id"`
	Brand          string `json:"brand"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Validation     string `json:"validation"`
	IncludeDomains int    `json:"includeDomains"`
	MaxDomains     int    `json:"maxDomains"`
	Prices         []struct {
		Years              int     `json:"years"`
		Price              float64 `json:"price,string"`
		PriceExtra         float64 `json:"priceExtra,string"`
		PriceExtraWildcard float64 `json:"priceExtraWildcard,string"`
	} `json:"prices"`
}

type RequestResponse struct {
	BaseResponse
	Id                      int        `json:"id"`
	DomainName              string     `json:"domainName"`
	SubjectAlternativeNames []string   `json:"subjectAlternativeNames"`
	Years                   int        `json:"years"`
	Company                 string     `json:"company"`
	DateOrdered             CustomTime `json:"dateOrdered"`
	Validations             struct {
		Docs         RequestValidationResponse `json:"docs"`
		Organization RequestValidationResponse `json:"organization"`
		Phone        RequestValidationResponse `json:"phone"`
		Whois        RequestValidationResponse `json:"whois"`
		DCV          RequestValidationResponse `json:"dcv"`
		Request      RequestValidationResponse `json:"request"`
	} `json:"validations"`
	Department        string `json:"department"`
	Address           string `json:"address"`
	ZIPCode           string `json:"zipCode"`
	City              string `json:"city"`
	Province          string `json:"province"`
	Country           string `json:"country"`
	Reference         string `json:"reference"`
	ApproverFirstName string `json:"approverFirstName"`
	ApproverLastName  string `json:"approverLastName"`
	ApproverEmail     string `json:"approverEmail"`
	ApproverPhone     string `json:"approverPhone"`
	RequiresAction    bool   `json:"requiresAction"`
	BrandValidation   bool   `json:"brandValidation"`
	KVK               string `json:"kvk"`
	Embedded          struct {
		Product ProductResponse `json:"product"`
	} `json:"_embedded"`
}

type RequestEEResponse struct {
	BaseResponse
	Id          int        `json:"id"`
	DateOrdered CustomTime `json:"dateOrdered"`
	Pkcs7       string     `json:"pkcs7"`
	Certificate string     `json:"crt"`
}

type CertificatesResponse struct {
	BaseResponse
	Embedded struct {
		Certificates []CertificateResponse `json:"certificates"`
	} `json:"_embedded"`
}

type CertificateResponse struct {
	BaseResponse
	Id                      int        `json:"id"`
	DomainName              string     `json:"domainName"`
	SubjectAlternativeNames []string   `json:"subjectAlternativeNames"`
	DateIssued              CustomTime `json:"dateIssued"`
	DateExpired             CustomTime `json:"dateExpired"`
	Company                 string     `json:"company"`
	CustomerId              int        `json:"customerId"`
	Embedded                struct {
		Product ProductResponse `json:"product"`
	} `json:"_embedded"`
}

type CSRResponse struct {
	BaseResponse
	Type     string   `json:"type"`
	Size     int      `json:"size"`
	Company  string   `json:"company"`
	CN       string   `json:"cn"`
	State    string   `json:"state"`
	City     string   `json:"city"`
	Country  string   `json:"country"`
	AltNames []string `json:"altNames"`
}

type NotesResponse struct {
	BaseResponse
	Embedded struct {
		Notes []NoteResponse `json:"notes"`
	} `json:"_embedded"`
}

type NoteResponse struct {
	BaseResponse
	Contact   string     `json:"contact"`
	Staff     string     `json:"staff"`
	Date      string     `json:"date"`
	Time      string     `json:"time"`
	Message   string     `json:"message"`
	CreatedAt CustomTime `json:"createdAt"`
}
