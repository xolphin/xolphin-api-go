package xolphin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Client struct {
	BaseURI string

	username string
	password string

	client *http.Client
}

func (self Client) Get(method string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(self.BaseURI, method)+"?"+data.Encode(), nil)
	if err != nil {
		return []byte{}, err
	}
	req.SetBasicAuth(self.username, self.password)
    req.Header.Add("Accept", "application/json")
    req.Header.Set("User-Agent", "xolphin-api-go/v1.5.0")
	resp, err := self.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (self Client) Post(method string, data url.Values) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range data {
		if k == "document" {
			file, err := os.Open(v[0])
			if err != nil {
				return nil, err
			}
			defer file.Close()

			part, err := writer.CreateFormFile("document", filepath.Base(v[0]))
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return nil, err
			}
		} else {
			writer.WriteField(k, v[0])
		}
	}
	err := writer.Close()
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(self.BaseURI, method), body)
	if err != nil {
		return []byte{}, err
	}
	req.SetBasicAuth(self.username, self.password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "xolphin-api-go/v1.5.0")
	resp, err := self.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (self Client) Download(method string, data url.Values) ([]byte, error) {
	return self.Get(method, data)
}

func (self *Client) Request() *RequestEndpoints {
	return &RequestEndpoints{
		c: self,
	}
}

func (self *Client) Certificate() *CertificateEndpoints {
	return &CertificateEndpoints{
		c: self,
	}
}

func (self *Client) Support() *SupportEndpoints {
	return &SupportEndpoints{
		c: self,
	}
}

func GetClient(username, password string, test bool) *Client{
	var uri = "https://api.xolphin.com/v1/%s"
	if(test){
		uri = "https://test-api.xolphin.com/v1/%s"
	}

	c := &Client{
		BaseURI:  uri,
		username: username,
		password: password,
		client: &http.Client{},
	}
	return c
}

func NewClient(username, password string) *Client {
	return GetClient(username, password, false)
}

func NewTestClient(username, password string) *Client {
	return GetClient(username, password, true)
}
