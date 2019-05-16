package api

// Response REST API
type Response struct {
	Response interface{} `json:"response"`
	Data     interface{} `json:"data,omitemp"`
}
