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
