package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RestResponse struct {
	Status int
	Error  error
	Body   []byte
}

type RestClient struct {
	httpClient *http.Client
}

func (self *RestClient) SetHttpClient(httpClient *http.Client) {
	self.httpClient = httpClient
}

func (self *RestClient) GetObject(url string, obj interface{}) RestResponse {
	return self.executeRequest("GET", url, obj)
}

func (self *RestClient) executeRequest(method string, url string, obj interface{}) RestResponse {
	client := self.httpClient
	if client == nil {
		client = &http.Client{}
	}
	ret := RestResponse{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		ret.Error = err
		return ret
	}
	resp, err := client.Do(req)
	if err != nil {
		ret.Error = err
		return ret
	}
	ret.Status = resp.StatusCode
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ret.Error = err
		return ret
	}
	ret.Body = data
	if obj != nil {
		json.Unmarshal(data, obj)
	}
	return ret
}
