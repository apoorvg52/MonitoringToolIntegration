package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"monitoringtool/api/models"
	"monitoringtool/api/utils"
    "errors"

)

func SendMetricToNewRelic(apiKey string, payload []models.Payload) (string, error) {

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", constants.METRICS_URL, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }
    
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Api-Key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
        return "", fmt.Errorf("failed to send data to New Relic: %s", resp.Status)
    }

	var responsePayload models.MetricResponse
	// Read the response body
    body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &responsePayload)
    if err != nil {
        return "", fmt.Errorf("failed to receive the response from the server %s", resp.Status)
    }

    return responsePayload.RequestId, nil
}

func SendLogsToNewRelic(apiKey string, payload models.LogData) (string, error) {

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", constants.LOGS_URL, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }
    
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Api-Key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
        return "", fmt.Errorf("failed to send data to New Relic: %s", resp.Status)
    }

	var responsePayload models.MetricResponse
	// Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("failed to unmarshal response body %s", resp.Status)
    }

	err = json.Unmarshal(body, &responsePayload)
    if err != nil {
        return "", fmt.Errorf("failed to receive the response from the server %s", resp.Status)
    }

    return responsePayload.RequestId, nil
}

func SendEventsToNewRelic(apiKey string, compressedData []byte) (models.EventsResponse, error) {

    var responsePayload models.EventsResponse

     // Create a new HTTP request
     req, err := http.NewRequest("POST", constants.GetEventsURL(), bytes.NewBuffer(compressedData))
     if err != nil {
        fmt.Println("Error creating HTTP request:", err)
        return responsePayload, err
     }
 
     // Set headers
     req.Header.Set("Content-Type", "application/json")
     req.Header.Set("Api-Key", apiKey)
     req.Header.Set("Content-Encoding", "gzip")
 
     // Create a new HTTP client and send the request
     client := &http.Client{}
     resp, err := client.Do(req)
     if err != nil {
        fmt.Println("Error sending request:", err)
        return responsePayload, err
     }
     defer resp.Body.Close()

     if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error: received status code %d\n", resp.StatusCode)
        return responsePayload, err
    }
	// Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return responsePayload, fmt.Errorf("failed to unmarshal response body %s", resp.Status)
    }

	err = json.Unmarshal(body, &responsePayload)
    if err != nil {
        return responsePayload, fmt.Errorf("failed to receive the response from the server %s", resp.Status)
    }

    return responsePayload, nil
}

func SendTraceToNewRelic(apiKey string, payload models.TraceData)  (string, error){

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    // Create a new HTTP request
    req, err := http.NewRequest("POST", constants.TRACE_URL, bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error creating HTTP request:", err)
        return "", err
    }

    // Set headers
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Api-Key", apiKey)

    // Create a new HTTP client and send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return "", err
    }
    defer resp.Body.Close()

    // Check the response status
    if resp.StatusCode != http.StatusAccepted {
        fmt.Printf("Error: received status code %d\n", resp.StatusCode)
        return "", errors.New("Response is not 200.")
    }

    var responsePayload models.MetricResponse
	// Read the response body
    body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &responsePayload)
    if err != nil {
        return "", fmt.Errorf("failed to receive the response from the server %s", resp.Status)
    }

    return responsePayload.RequestId, nil

}
