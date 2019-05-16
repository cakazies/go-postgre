package api

type Response struct {
	Response interface{} `json:"response"`
	Data     interface{} `json:"data,omitemp"`
}
