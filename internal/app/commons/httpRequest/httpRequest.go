package httpRequest

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type HttpHeaders map[string]string

func GetHTTPResponse(method string, url string, jsondatastring string, arrheader map[string]string) (string, error) {

	finalresult := ""

	var ioreader io.Reader
	if jsondatastring != "" {
		ioreader = bytes.NewBuffer([]byte(jsondatastring))
	}

	req, err := http.NewRequest(method, url, ioreader)

	if len(arrheader) > 0 {
		for k, v := range arrheader {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		finalresult = ""
		return finalresult, err
	}
	defer resp.Body.Close()

	respbody, _ := ioutil.ReadAll(resp.Body)
	finalresult = string(respbody)
	// fmt.Println("response Body:", finalresult)

	return finalresult, nil

}

func PostData(url string, body []byte, headers HttpHeaders, timeoutInSec int) (code int, response []byte, err error) {
	clientHTTP := http.Client{}

	clientHTTP.Timeout = time.Second * 60 //set timeout to 1 minute (default)
	if timeoutInSec != 0 {
		clientHTTP.Timeout = time.Second * time.Duration(timeoutInSec) //set timeout to custom timeout
	}

	req, err := http.NewRequest("POST",
		url, bytes.NewReader(body))

	if err != nil {
		return 0, nil, err
	}

	//add header
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	//execute http post
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return 0, nil, nil
	}
	defer resp.Body.Close()

	//read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, responseBody, nil
}

func PostDataWithBasicAuth(url string, username, password string, body []byte, headers HttpHeaders, timeoutInSec int) (code int, response []byte, err error) {
	//build basic auth token
	authToken := base64.URLEncoding.EncodeToString([]byte(username + ":" + password))
	if headers == nil {
		headers = make(map[string]string)
	}

	//add Authorization header
	headers["Authorization"] = "Basic " + authToken

	//check if Client-Id header is not exists
	//this header is used in some authorization server
	if _, ok := headers["Client-Id"]; !ok {
		headers["Client-Id"] = username
	}

	clientHTTP := http.Client{}

	clientHTTP.Timeout = time.Second * 60 //set timeout to 1 minute (default)
	if timeoutInSec != 0 {
		clientHTTP.Timeout = time.Second * time.Duration(timeoutInSec) //set timeout to custom timeout
	}

	req, err := http.NewRequest("POST",
		url, bytes.NewReader(body))

	if err != nil {
		return 0, nil, err
	}

	//add header
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	//execute http post
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return 0, nil, nil
	}
	//read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, responseBody, nil
}

// Creates a new file upload http request with optional extra params
func MultiFileUploadRequest(uri string,
	params map[string]string,
	headers map[string]string,
	paramName string, paths []string) (*http.Request, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(paramName, filepath.Base(path))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	for key, headerVal := range headers {
		req.Header.Add(key, headerVal)
	}

	return req, err
}

//MultiFileUploadRequestWithBasicAuth do multifile upload request with basic auth method
func MultiFileUploadRequestWithBasicAuth(uri string,
	username, password string,
	params map[string]string,
	headers map[string]string,
	paramName string, paths []string) (*http.Request, error) {

	//build basic auth token
	authToken := base64.URLEncoding.EncodeToString([]byte(username + ":" + password))
	if headers == nil {
		headers = make(map[string]string)
	}

	//add Authorization header
	headers["Authorization"] = "Basic " + authToken

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(paramName, filepath.Base(path))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	for key, headerVal := range headers {
		req.Header.Add(key, headerVal)
	}

	return req, err
}
