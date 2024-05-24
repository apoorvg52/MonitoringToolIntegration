package models

type TraceData struct {
    TraceID   string `json:"trace.id"`
    SpanID    string `json:"span.id"`
    Timestamp int64  `json:"timestamp"`
    Name      string `json:"name"`
    Duration  int64  `json:"duration"`
    Service   string `json:"service"`
}