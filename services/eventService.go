package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"monitoringtool/api/models"
	"monitoringtool/api/utils"
)

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
