package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)



func CallExternalAPI(url string, method string, payload interface{}, headers map[string]string, path string) (*http.Response, error) {
    // Convert payload to JSON if it's not nil
    var jsonData []byte
    var err error
    if payload != nil {
        jsonData, err = json.Marshal(payload)
        if err != nil {
            return nil, err
        }
    }

    // Create a new HTTP request
    req_url := url + path
    
    fmt.Println("URL:>", req_url)
    var req *http.Request
    if payload != nil {
        req, err = http.NewRequest(method, req_url, bytes.NewBuffer(jsonData))
    } else {
        req, err = http.NewRequest(method, req_url, nil)
    }
    if err != nil {
        return nil, err
    }

    // Set headers
    for key, value := range headers {
        req.Header.Set(key, value)
    }


    // Create an HTTP client and set a timeout
    client := &http.Client{Timeout: 10 * time.Second}

    // Make the HTTP request
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    // Check for non-200 status codes
    if resp.StatusCode != http.StatusOK {
        log.Printf("/app/utils/callExternalAPI.go := Received code: %d\n" , resp.StatusCode)
        return nil, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
        
    }

    return resp, nil
}