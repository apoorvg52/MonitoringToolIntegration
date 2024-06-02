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
