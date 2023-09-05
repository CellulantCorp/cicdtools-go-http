package requests

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var response *http.Response
var requestError error
var client = &http.Client{
	Timeout: time.Second * 10,
}

func CheckRateLimit(req *http.Request) {
	// loop until the request is successful
	for requestError != nil {
		// logrus.Error("Error: ", err.Error())
		time.Sleep(5 * time.Second)
		fmt.Println(" .. .. ")
		response, requestError = client.Do(req)
	}

	// As long as the response.StatusCode is 429, keep retrying until the request is successful
	for response.StatusCode == 429 && response != nil {
		time.Sleep(5 * time.Second)
		fmt.Println(" .. ")
		response, requestError = client.Do(req)
	}
}

func SendGetRequest(url string, authToken string) (*http.Response, []byte) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", authToken))

	// send the request and get the response
	response, requestError = client.Do(req)
	CheckRateLimit(req)
	body, _ := ioutil.ReadAll(response.Body)
	// return response.Header, body
	return response, body
}

func SendPostRequest(url string, authToken string, requestBody io.Reader) (*http.Response, []byte) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	//req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", authToken))

	// send the request and get the response
	response, requestError = client.Do(req)
	CheckRateLimit(req)
	body, _ := ioutil.ReadAll(response.Body)

	// Check for 200 response code before returning the body
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if statusOK {
		return response, body
	} else {
		logrus.Errorln(response.StatusCode, " - Error sending POST request. ", string(body))
		return response, nil
	}

	//return response.Header, body
}

func SendPutRequest(url string, authToken string, requestBody io.Reader) (*http.Response, []byte) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("PUT", url, requestBody)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", authToken))

	// send the request and get the response
	response, requestError = client.Do(req)
	CheckRateLimit(req)
	body, _ := ioutil.ReadAll(response.Body)
	// return response.Header, body
	return response, body
}

func SendDeleteRequest(url string, authToken string, requestBody io.Reader) (*http.Response, []byte) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("DELETE", url, requestBody)
	if err != nil {
		//return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("user-agent", "GitlabScannerAPI")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Keep-Alive", "timeout=30, max=60")
	req.Header.Add("Authorization", fmt.Sprintf("%s%s", "Bearer ", authToken))

	// send the request and get the response
	response, requestError = client.Do(req)
	CheckRateLimit(req)
	body, _ := ioutil.ReadAll(response.Body)
	// return response.Header, body
	return response, body
}
