package xolphin

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"net/url"
	"time"
)

type RequestEndpoints struct {
	c *Client
}

func (self *RequestEndpoints) All() ([]RequestResponse, error) {
	result := []RequestResponse{}

	data, err := self.c.Get("requests", url.Values{"page": []string{"1"}})
	if err != nil {
		return result, err
	}
	var requestsData RequestsResponse
	err = json.Unmarshal(data, &requestsData)
	if err != nil {
		return result, err
	}
	if !requestsData.isError() {
		result = requestsData.Embedded.Requests
		for {
			if requestsData.Page == requestsData.Pages {
				return result, nil
			}

			data, err = self.c.Get("requests", url.Values{"page": []string{fmt.Sprintf("%d", requestsData.Page + 1)}})
			if err != nil {
				return result, err
			}
			requestsData = RequestsResponse{}
			err = json.Unmarshal(data, &requestsData)
			if err != nil {
				return result, err
			}
			if !requestsData.isError() {
				result = append(result, requestsData.Embedded.Requests...)
			} else {
				return result, errors.New(requestsData.Message)
			}
		}
	} else {
		return result, errors.New(requestsData.Message)
	}
}

func (self *RequestEndpoints) Create(product int, years int, csr string, dcv_type string) CertificateCreationRequest {
	return CertificateCreationRequest{
		Product: product,
		Years: years,
		CSR: csr,
		DCVType: dcv_type,
	}
}

func (self *RequestEndpoints) Reissue(csr string, dcv_type string) CertificateReissueRequest {
	return CertificateReissueRequest{
		CSR: csr,
		DCVType: dcv_type,
	}
}

func (self *RequestEndpoints) Renew(product int, years int, csr string, dcv_type string) CertificateRenewRequest {
	return CertificateRenewRequest{
		Product: product,
		Years: years,
		CSR: csr,
		DCVType: dcv_type,
	}
}

func (self *RequestEndpoints) CreateEE() EECreationRequest {
	return EECreationRequest{}
}

func (self *RequestEndpoints) Send(request CertificateCreationRequest) (RequestResponse, error) {
	result := RequestResponse{}
	data, err := self.c.Post("requests", request.ToURLValues())
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

func (self *RequestEndpoints) SendEE(request EECreationRequest) (RequestEEResponse, error) {
	result := RequestEEResponse{}
	data, err := self.c.Post("requests/ee", request.ToURLValues())
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

func (self *RequestEndpoints) Get(id int) (RequestResponse, error) {
	result := RequestResponse{}
	data, err := self.c.Get(fmt.Sprintf("requests/%d", id), url.Values{})
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

func (self *RequestEndpoints) UploadDocument(id int, file_path string, description string) (BaseResponse, error) {
	result := BaseResponse{}

	payload := url.Values{}
	payload.Set("document", file_path)
	payload.Set("description", description)

	data, err := self.c.Post(fmt.Sprintf("requests/%d/upload-document", id), payload)
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

func (self *RequestEndpoints) RetryDCV(id int, domain string, dcv_type string, email string) (BaseResponse, error) {
	result := BaseResponse{}

	payload := url.Values{}
	payload.Set("domain", domain)
	payload.Set("dcvType", dcv_type)
	if email != "" {
		payload.Set("email", email)
	}

	data, err := self.c.Post(fmt.Sprintf("requests/%d/retry-dcv", id), payload)
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
func (self *RequestEndpoints) ScheduleValidationCall(id int, callParams interface{}) (BaseResponse, error) {
	result := BaseResponse{}
	payload := url.Values{}

	callParamsType := reflect.TypeOf(callParams)
	name := callParamsType.Name()

	if(name == "Time"){
		callDetails, ok := callParams.(time.Time)
		if(ok){			
			payload.Set("date", callDetails.Format("2006-01-02"))
			payload.Set("time", callDetails.Format("15:04"))
		}
	}else{
		callDetails, ok := callParams.(map[string]string)
		if(ok){			
			payload.Set("timezone",callDetails["Timezone"])
			payload.Set("action",callDetails["Action"])
			payload.Set("phoneNumber",callDetails["PhoneNumber"])
			payload.Set("extensionNumber",callDetails["ExtensionNumber"])
			payload.Set("emailAddress",callDetails["EmailAddress"])
			payload.Set("language",callDetails["Language"])
			payload.Set("comments",callDetails["Comments"])
			payload.Set("date",callDetails["Date"])
			payload.Set("time",callDetails["Time"])
		}
	}
	
	data, err := self.c.Post(fmt.Sprintf("requests/%d/schedule-validation-call", id), payload)

	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil

}



func (self *RequestEndpoints) SendNote(id int, note string) (BaseResponse, error) {
	result := BaseResponse{}

	payload := url.Values{}
	payload.Set("message", note)

	data, err := self.c.Post(fmt.Sprintf("requests/%d/notes", id), payload)
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

func (self *RequestEndpoints) GetNotes(id int) ([]NoteResponse, error) {
	result := []NoteResponse{}

	payload := url.Values{}

	data, err := self.c.Get(fmt.Sprintf("requests/%d/notes", id), payload)

	if err != nil {
		return result, err
	}
	var requestsData NotesResponse
	err = json.Unmarshal(data, &requestsData)
	if err != nil {
		return result, err
	}
	if requestsData.isError() {
		return result, err
	}else{
		return requestsData.Embedded.Notes, err
	}
}

func (self *RequestEndpoints) SendComodoSA(id int, to string, language string) (BaseResponse, error) {
	result := BaseResponse{}

	payload := url.Values{}
	payload.Set("sa_email", to)
	payload.Set("language", language)

	data, err := self.c.Post(fmt.Sprintf("requests/%d/sa", id), payload)
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