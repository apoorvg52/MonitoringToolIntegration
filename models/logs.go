package models

type LogData struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	LogLevel  string `json:"logLevel"`
	Service   string `json:"service"`
}