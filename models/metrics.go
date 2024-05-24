package models

type Metric struct {
    Name       			string            		`json:"name"`
    Type       			string            		`json:"type"`
    Value      			float64           		`json:"value"`
	Timestamp			int64			  		`json:"timeStamp"`
    Attributes 			map[string]string 		`json:"attributes"`
}

type Payload struct {
    Metrics []Metric `json:"metrics"`
}

type MetricResponse struct{
	RequestId 		string 		`json:"requestid"`
}