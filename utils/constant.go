package constants

import (
    "bytes"
    "compress/gzip"
    "io/ioutil"
	"monitoringtool/api/models"
	"time"
	"fmt"
)

const (
	// NEW_RELIC_API_KEY    = "YOUR_NEW_RELIC_API_KEY"
	NEW_RELIC_API_KEY    = "eu01xx889db64394229f7f835a639368249aNRAL"
	METRICS_URL = "https://metric-api.eu.newrelic.com/metric/v1"
	LOGS_URL = "https://log-api.eu.newrelic.com/log/v1"
	TRACE_URL = "https://trace-api.eu.newrelic.com/trace/v1"
	NEW_RELIC_ACCOUNT_ID = "4455474"
)

func GetEventsURL() string {
	return fmt.Sprintf("https://insights-collector.eu01.nr-data.net/v1/accounts/%d/events", NEW_RELIC_ACCOUNT_ID)
}
// Function to read and compress JSON data from a file
func ReadAndCompressJSON(filePath string) ([]byte, error) {
    fileData, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var buffer bytes.Buffer
    writer := gzip.NewWriter(&buffer)
    if _, err := writer.Write(fileData); err != nil {
        return nil, err
    }
    if err := writer.Close(); err != nil {
        return nil, err
    }

    return buffer.Bytes(), nil
}

func ProcessPayload(payload []models.Payload) ([]models.Payload) {
	for index, _ := range payload{
		for metricsIndex, _ := range payload[index].Metrics {
			payload[index].Metrics[metricsIndex].Timestamp = time.Now().Unix()
		}
	}

	return payload
}
