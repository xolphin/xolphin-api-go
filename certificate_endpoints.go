package xolphin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type CertificateEndpoints struct {
	c *Client
}

func (self *CertificateEndpoints) All() ([]CertificateResponse, error) {
	result := []CertificateResponse{}

	data, err := self.c.Get("certificates", url.Values{"page": []string{"1"}})
	if err != nil {
		return result, err
	}
	var certsData CertificatesResponse
	err = json.Unmarshal(data, &certsData)
	if err != nil {
		return result, err
	}
	if !certsData.isError() {
		result = certsData.Embedded.Certificates
		for {
			if certsData.Page == certsData.Pages {
				return result, nil
			}

			data, err = self.c.Get("certificates", url.Values{"page": []string{fmt.Sprintf("%d", certsData.Page+1)}})
			if err != nil {
				return result, err
			}
			certsData = CertificatesResponse{}
			err = json.Unmarshal(data, &certsData)
			if err != nil {
				return result, err
			}
			if !certsData.isError() {
				result = append(result, certsData.Embedded.Certificates...)
			} else {
				return result, errors.New(certsData.Message)
			}
		}
	} else {
		return result, errors.New(certsData.Message)
	}
}

func (self *CertificateEndpoints) Get(id int) (CertificateResponse, error) {
	result := CertificateResponse{}
	data, err := self.c.Get(fmt.Sprintf("certificates/%d", id), url.Values{})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	if result.isError() {
		return result, errors.New(result.Message)
	}
	return result, nil
}

func (self *CertificateEndpoints) Download(id int, format string) ([]byte, error) {
	payload := url.Values{}
	payload.Set("format", format)

	return self.c.Get(fmt.Sprintf("certificates/%d/download", id), payload)
}

func (self *CertificateEndpoints) Reissue(id int, request CertificateReissueRequest) (RequestResponse, error) {
	result := RequestResponse{}
	data, err := self.c.Post(fmt.Sprintf("certificates/%d/reissue", id), request.ToURLValues())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	if result.isError() {
		return result, errors.New(result.Message)
	}
	return result, nil
}

func (self *CertificateEndpoints) Renew(id int, request CertificateRenewRequest) (RequestResponse, error) {
	result := RequestResponse{}
	data, err := self.c.Post(fmt.Sprintf("certificates/%d/renew", id), request.ToURLValues())
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	if result.isError() {
		return result, errors.New(result.Message)
	}
	return result, nil
}

func (self *CertificateEndpoints) Cancel(id int, reason string, revoke bool) (BaseResponse, error) {
	result := BaseResponse{}

	payload := url.Values{}
	payload.Set("reason", reason)
	if revoke {
		payload.Set("revoke", "true")
	}

	data, err := self.c.Post(fmt.Sprintf("certificates/%d/cancel", id), payload)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	if result.isError() {
		return result, errors.New(result.Message)
	}
	return result, nil
}
