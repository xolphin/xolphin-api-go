package xolphin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

type SupportEndpoints struct {
	c *Client
}

func (self *SupportEndpoints) ApproverEmailAddresses(domain string) ([]string, error) {
	result := []string{}
	payload := url.Values{}
	payload.Set("domain", domain)

	data, err := self.c.Get("approver-email-addresses", payload)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (self *SupportEndpoints) DecodeCSR(csr string) (CSRResponse, error) {
	result := CSRResponse{}

	payload := url.Values{}
	payload.Set("csr", csr)

	data, err := self.c.Post("decode-csr", payload)
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

func (self *SupportEndpoints) Products() ([]ProductResponse, error) {
	result := []ProductResponse{}

	data, err := self.c.Get("products", url.Values{"page": []string{"1"}})
	if err != nil {
		return result, err
	}
	var prodsData ProductsResponse
	err = json.Unmarshal(data, &prodsData)
	if err != nil {
		return result, err
	}
	if !prodsData.isError() {
		result = prodsData.Embedded.Products
		for {
			if prodsData.Page == prodsData.Pages {
				return result, nil
			}

			data, err = self.c.Get("products", url.Values{"page": []string{fmt.Sprintf("%d", prodsData.Page+1)}})
			if err != nil {
				return result, err
			}
			prodsData = ProductsResponse{}
			err = json.Unmarshal(data, &prodsData)
			if err != nil {
				return result, err
			}
			if !prodsData.isError() {
				result = append(result, prodsData.Embedded.Products...)
			} else {
				return result, errors.New(prodsData.Message)
			}
		}
	} else {
		return result, errors.New(prodsData.Message)
	}
}

func (self *SupportEndpoints) Product(id int) (ProductResponse, error) {
	result := ProductResponse{}
	data, err := self.c.Get(fmt.Sprintf("products/%d", id), url.Values{})
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




